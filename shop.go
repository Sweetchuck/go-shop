package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/storage"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var Version = "1.0.0"

var GitRev = ""

var EnvVarNamePrefix = "SHOP"

var ApiPathPrefix = "/api/v1"

var DataSourceUrl = "mysql://root:mysql@tcp(127.0.0.1:3306)/go_shop2?charset=utf8"

var Address = ":8080"

var actionToDo = "listenAndServe"

var personServer person.Server

var logger = logrus.New()

var router *mux.Router

func init() {
	initCliArgsAndOptions()
}

func initCliArgsAndOptions() {
	showVersion := flag.Bool("version", false, "Show the version number")
	flag.Parse()

	if *showVersion {
		actionToDo = "showVersion"
	}
}

func initEnvVar(name string, dst *string) {
	value := os.Getenv(envVarName(name))
	if value != "" {
		*dst = value
	}
}

func main() {
	switch actionToDo {
	case "showVersion":
		fmt.Printf("%s\n", fullVersion())

	default:
		initEnvVar("data_source_url", &DataSourceUrl)
		initEnvVar("api_path_prefix", &ApiPathPrefix)
		initEnvVar("address", &Address)

		personServer = person.Server{}
		initStorage()

		router = mux.NewRouter()
		registerPersonServerRoutes(ApiPathPrefix+"/person", personServer)

		logger.Infof("starting web server on: %s", Address)

		http.Handle("/", router)
		err := http.ListenAndServe(Address, nil)
		if err != nil {
			panic(err)
		}
	}
}

func initStorage() {
	sqlDialect, sqlArgs, err := parseDataSourceUrl(DataSourceUrl)
	if err != nil {
		logger.Errorf("Could not parse '%s' as URL\n", DataSourceUrl)

		panic(err)
	}

	switch sqlDialect {
	case "memory":
		initStorageMemory()

	case "mysql", "postgres", "sqlite3":
		initStorageSQL(sqlDialect, sqlArgs)

	default:
		panic("Unknown data source: " + DataSourceUrl)
	}
}

func initStorageMemory() {
	storageHandler := &storage.Memory{}
	storageHandler.Init()
	personServer.Storage = storageHandler
}

func initStorageSQL(sqlDialect string, sqlArgs []interface{}) {
	db, err := gorm.Open(sqlDialect, sqlArgs...)
	if err != nil {
		logrus.Error(err)

		panic(err)
	}

	logger.Infof("Database connection successfully opened to '%s'", sqlDialect)

	db.SingularTable(true)
	db.LogMode(true)
	db.AutoMigrate(&model.Person{})

	storageHandler := &storage.Sql{}
	storageHandler.Init(db)

	personServer.Storage = storageHandler
}

func registerPersonServerRoutes(pathPrefix string, ps person.Server) {
	router.
		HandleFunc(pathPrefix+"", ps.List).
		Methods("GET")

	router.
		HandleFunc(pathPrefix+"", ps.Insert).
		Methods("POST")

	router.
		HandleFunc(pathPrefix+"/{id}", ps.Read).
		Methods("GET")

	router.
		HandleFunc(pathPrefix+"/{id}", ps.Update).
		Methods("PATCH")

	router.
		HandleFunc(pathPrefix+"/{id}", ps.Delete).
		Methods("DELETE")
}

func parseDataSourceUrl(dsUrl string) (sqlDialect string, sqlArgs []interface{}, err error) {
	var dbUrl *url.URL

	dbUrl, err = url.Parse(dsUrl)
	if err != nil {
		return
	}

	sqlDialect = dbUrl.Scheme

	parts := strings.SplitN(dbUrl.String(), "//", 2)
	for _, sqlArg := range parts[1:] {
		sqlArgs = append(sqlArgs, sqlArg)
	}

	return
}

func envVarName(name string) string {
	return EnvVarNamePrefix + "_" + strings.ToUpper(name)
}

func fullVersion() string {
	v := Version
	if GitRev == "" {
		GitRev = "dev"
	}

	return v + "-" + GitRev
}
