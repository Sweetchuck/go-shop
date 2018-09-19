package storage

import (
	"github.com/jinzhu/gorm"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
)

type SQL struct {
	db *gorm.DB
}

func (s *SQL) Init(db *gorm.DB) {
	s.db = db
}

func (s *SQL) Insert(p model.Person) (model.Person, error) {
	pp := &p

	s.db.Model(p).Create(pp)

	return *pp, nil
}

func (s *SQL) Read(id uint) (p model.Person) {
	s.db.Model(model.Person{}).First(&p)

	return
}

func (s *SQL) Update(p model.Person, fields map[string]interface{}) (model.Person, error) {
	s.db.Model(&p).Updates(fields)

	return p, nil
}

func (s *SQL) Delete(id uint) error {
	p := model.Person{
		Model: gorm.Model{
			ID: id,
		},
	}
	s.db.Delete(p)

	return nil
}

func (s *SQL) List() (list []model.Person) {
	list = []model.Person{}
	s.db.Model(model.Person{}).Find(&list)

	return
}

func (s *SQL) Count() int {
	numOfRecords := new(int)
	s.db.Model(model.Person{}).Count(numOfRecords)

	return *numOfRecords
}
