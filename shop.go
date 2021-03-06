package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/base"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person"
	personModel "gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	personStorage "gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/storage"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe"
	shoeModel "gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/model"
	shoeStorage "gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/storage"
	"net/http"
	"os"
	"strings"
)

// Version number
var Version = "1.0.0"

// GitRev stores the meta data part of the version string
var GitRev = ""

// EnvVarNamePrefix stores the used environment variable names
var EnvVarNamePrefix = "SHOP"

// APIPathPrefix stores the path prefix for the api
var APIPathPrefix = "/api/v1"

// DataSourceURL stores the gorm data source URL.
// The "memory://whatever" also valid
var DataSourceURL = "mysql://root:mysql@tcp(127.0.0.1:3306)/go_shop2?charset=utf8&parseTime=True&loc=Local"

// Address to listen on
var Address = ":8080"

var actionToDo = "listenAndServe"

var personServer person.Server

var shoeServer shoe.Server

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

func writeEnvVarIntoString(name string, dst *string) {
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
		writeEnvVarIntoString("data_source_url", &DataSourceURL)
		writeEnvVarIntoString("api_path_prefix", &APIPathPrefix)
		writeEnvVarIntoString("address", &Address)

		personServer = person.Server{}
		shoeServer = shoe.Server{}

		initStorage()
		router = mux.NewRouter()

		registerStatusRoutes(APIPathPrefix + "/status")

		registerPersonServerRoutes(APIPathPrefix+"/person", personServer)

		registerShoeServerRoutes(APIPathPrefix+"/shoe", shoeServer)

		logger.Infof("starting web server on: %s", Address)
		http.Handle("/", router)
		err := http.ListenAndServe(Address, nil)
		if err != nil {
			panic(err)
		}
	}
}

func initStorage() {
	sqlDialect, sqlArgs, err := parseDataSourceURL(DataSourceURL)
	if err != nil {
		logger.Errorf("Could not parse '%s' as URL\n", DataSourceURL)

		panic(err)
	}

	switch sqlDialect {
	case "memory":
		initStorageMemory()

	case "mysql", "postgres", "sqlite3":
		initStorageSQL(sqlDialect, sqlArgs)

	default:
		panic("Unknown data source: " + DataSourceURL)
	}
}

func initStorageMemory() {
	psh := &personStorage.Memory{}
	psh.Init()
	personServer.Storage = psh

	ssh := &shoeStorage.Memory{}
	ssh.Init()
	shoeServer.Storage = ssh
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

	db.AutoMigrate(&personModel.Person{})
	psh := &personStorage.SQL{}
	psh.Init(db)
	personServer.Storage = psh

	db.AutoMigrate(&shoeModel.Shoe{})
	ssh := &shoeStorage.SQL{}
	ssh.Init(db)
	shoeServer.Storage = ssh
}

func registerPersonServerRoutes(pathPrefix string, server person.Server) {
	router.
		HandleFunc(pathPrefix+"", server.List).
		Methods("GET")

	registerCrudServerRoutes(pathPrefix, server)
}

func registerShoeServerRoutes(pathPrefix string, server shoe.Server) {
	router.
		HandleFunc(pathPrefix+"", server.List).
		Methods("GET")

	registerCrudServerRoutes(pathPrefix, server)
}

func registerCrudServerRoutes(pathPrefix string, server base.CrudServer) {
	router.
		HandleFunc(pathPrefix+"", server.Create).
		Methods("POST")

	router.
		HandleFunc(pathPrefix+"/{id}", server.Read).
		Methods("GET")

	router.
		HandleFunc(pathPrefix+"/{id}", server.Update).
		Methods("PATCH")

	router.
		HandleFunc(pathPrefix+"/{id}", server.Delete).
		Methods("DELETE")
}

func registerStatusRoutes(pathPrefix string) {
	router.
		HandleFunc(
			pathPrefix,
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")

				body := map[string]interface{}{
					"version": fullVersion(),
				}

				encodedBody, err := json.MarshalIndent(body, "", "    ")
				if err != nil {
					logger.Error(err)

					return
				}

				_, err = fmt.Fprintf(w, string(encodedBody))
				if err != nil {
					logger.Error(err)
				}
			},
		).
		Methods("GET")
}

func parseDataSourceURL(dsURL string) (sqlDialect string, sqlArgs []interface{}, err error) {
	parts := strings.SplitN(dsURL, "://", 2)
	if len(parts) != 2 {
		return sqlDialect, sqlArgs, errors.New("invalid data source format")
	}

	sqlDialect = parts[0]
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
