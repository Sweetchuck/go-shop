package storage

import (
	"github.com/jinzhu/gorm"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
)

type Sql struct {
	db *gorm.DB
}

func (s *Sql) Init(db *gorm.DB) {
	s.db = db
}

func (s *Sql) Insert(p model.Person) (model.Person, error) {
	s.db.Model(p).Create(&p)

	return p, nil
}

func (s *Sql) Update(p model.Person, fields map[string]interface{}) (model.Person, error) {
	s.db.Model(&p).Updates(fields)

	return p, nil
}

func (s *Sql) Delete(id uint) error {
	p := model.Person{
		Model: gorm.Model{
			ID: id,
		},
	}
	s.db.Delete(p)

	return nil
}

func (s *Sql) List() (list []model.Person) {
	list = []model.Person{}
	s.db.Model(model.Person{}).Find(&list)

	return
}

func (s *Sql) Read(id uint) (p model.Person) {
	s.db.Model(model.Person{}).First(&p)

	return
}
