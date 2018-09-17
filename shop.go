package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/storage"
	"net/http"
)

var personServer person.Server

func init() {
	personServer = person.Server{}

	//initPersonServerStorageMemory()
	initPersonServerStorageSQL()

	router = mux.NewRouter()
}

func initPersonServerStorageMemory() {
	storageHandler := &storage.Memory{}
	storageHandler.Init()
	personServer.Storage = storageHandler
}

func initPersonServerStorageSQL() {
	sqlDialect, sqlArgs := ormDataSourceFromConfig()
	db, err := gorm.Open(sqlDialect, sqlArgs...)
	if err != nil {
		logrus.Error(err)

		panic(err)
	}

	db.LogMode(true)
	db.AutoMigrate(&model.Person{})

	storageHandler := &storage.Sql{}
	storageHandler.Init(db)

	personServer.Storage = storageHandler
}

var router *mux.Router

func main() {
	pathPrefix := "/api/v1"
	registerPersonServerRoutes(pathPrefix + "/person", personServer)

	http.Handle("/", router)

	fmt.Println("starting web server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func registerPersonServerRoutes(pathPrefix string, ps person.Server)  {
	router.
		HandleFunc(pathPrefix + "", ps.List).
		Methods("GET")

	router.
		HandleFunc(pathPrefix + "", ps.Insert).
		Methods("POST")

	router.
		HandleFunc(pathPrefix + "/{id}", ps.Read).
		Methods("GET")

	router.
		HandleFunc(pathPrefix + "/{id}", ps.Update).
		Methods("PATCH")

	router.
		HandleFunc(pathPrefix + "/{id}", ps.Delete).
		Methods("DELETE")
}

func ormDataSourceFromConfig() (string, []interface{}) {
	return "mysql",
		[]interface{}{
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?%s",
				"andor",
				"admin",
				"127.0.0.1",
				3311,
				"go_shop2",
				"charset=utf8",
			),
		}
}
