package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/model"
	"io/ioutil"
	"testing"
)

func TestHandler(t *testing.T) {
	names := [...]string{
		"memory",
		"sqlite3",
	}

	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			t.Run("insert", func (t *testing.T) {
				testInsert(t, name)
			})

			t.Run("read", func (t *testing.T) {
				testRead(t, name)
			})

			t.Run("update", func (t *testing.T) {
				testUpdate(t, name)
			})

			t.Run("delete", func (t *testing.T) {
				testDelete(t, name)
			})
		})
	}
}

func storageInstance(name string) (Handler, error) {
	var h Handler

	switch name {
	case "memory":
		memory := &Memory{}
		memory.Init()
		h = memory

	case "sqlite3":
		tmpFile, _ := ioutil.TempFile("", "go_shop2")

		db, err := gorm.Open("sqlite3", tmpFile.Name())
		if err != nil {
			return nil, err
		}

		db.LogMode(true)
		db.AutoMigrate(&model.Shoe{})

		sqLite3 := &Sql{}
		sqLite3.Init(db)
		h = sqLite3
	}

	return h, nil
}

func testInsert(t *testing.T, storageName string) {
	s, err := storageInstance(storageName)
	if err != nil {
		t.Error(err)
	}

	p1Src := model.Shoe{
		Brand: "a",
		Type: "b",
	}

	p1Dst, _ := s.Insert(p1Src)
	if p1Dst.ID != 1 {
		t.Fail()
	}

	if s.Count() != 1 {
		t.Fail()
	}

	p2Src := model.Shoe{
		Brand: "c",
		Type: "d",
	}

	p2Dst, _ := s.Insert(p2Src)
	if p2Dst.ID != 2 {
		t.Fail()
	}

	if s.Count() != 2 {
		t.Fail()
	}
}

func testRead(t *testing.T, storageName string) {
	s, err := storageInstance(storageName)
	if err != nil {
		t.Error(err)
	}

	var p1, p2 model.Shoe

	p1 = model.Shoe{
		Brand: "a",
		Type: "b",
	}
	p2, _ = s.Insert(p1)

	if p2.ID != 1 {
		t.Fail()
	}

	p3 := s.Read(1)
	if p3.ID != 1 {
		t.Fail()
	}

	if s.Count() != 1 {
		t.Fail()
	}
}

func testUpdate(t *testing.T, storageName string) {
	s, err := storageInstance(storageName)
	if err != nil {
		t.Error(err)
	}

	p1Src := model.Shoe{
		Brand: "a",
		Type: "b",
	}
	p1Dst, _ := s.Insert(p1Src)

	if p1Dst.ID != 1 {
		t.Fail()
	}

	if p1Dst.Brand != "a" {
		t.Fail()
	}

	p2Dst, _ := s.Update(
		p1Dst,
		map[string]interface{}{
			"Brand": "c",
		},
	)

	if p2Dst.Brand != "c" {
		t.Errorf("p2Dst.Name; expected = 'c'; current='%s';", p2Dst.Brand)
	}
}

func testDelete(t *testing.T, storageName string) {
	s, err := storageInstance(storageName)
	if err != nil {
		t.Error(err)
	}

	p1Src := model.Shoe{
		Brand: "a",
		Type: "b",
	}
	p1Dst, _ := s.Insert(p1Src)

	p2Src := model.Shoe{
		Brand: "a",
		Type: "b",
	}
	p2Dst, _ := s.Insert(p2Src)

	if s.Count() != 2 {
		t.Fail()
	}

	s.Delete(p2Dst.ID)
	if s.Count() != 1 {
		t.Fail()
	}

	list := s.List()
	if list[0].ID != p1Dst.ID {
		t.Fail()
	}
}
