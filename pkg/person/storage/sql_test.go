package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	"io/ioutil"
	"os"
	"testing"
)

func TestSql_Insert(t *testing.T) {
	tmpFile, _ := ioutil.TempFile("","go_shop2")
	defer os.Remove(tmpFile.Name())

	db, err := gorm.Open("sqlite3", tmpFile.Name())
	if err != nil {
		t.Error(err)
	}

	db.LogMode(true)
	db.AutoMigrate(&model.Person{})

	storageHandler := &Sql{}
	storageHandler.Init(db)

	var pIn, pOut model.Person

	pIn = model.Person{}
	pOut, _ = storageHandler.Insert(pIn)
	if pOut.ID != 1 {
		t.Fail()
	}

	pIn = model.Person{}
	pOut, _ = storageHandler.Insert(pIn)
	if pOut.ID != 2 {
		t.Fail()
	}
}
