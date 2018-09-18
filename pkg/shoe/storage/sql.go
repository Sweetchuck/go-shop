package storage

import (
	"github.com/jinzhu/gorm"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/model"
)

type Sql struct {
	db *gorm.DB
}

func (s *Sql) Init(db *gorm.DB) {
	s.db = db
}

func (s *Sql) Insert(p model.Shoe) (model.Shoe, error) {
	s.db.Model(p).Create(&p)

	return p, nil
}

func (s *Sql) Read(id uint) (p model.Shoe) {
	s.db.Model(model.Shoe{}).First(&p)

	return
}

func (s *Sql) Update(p model.Shoe, fields map[string]interface{}) (model.Shoe, error) {
	s.db.Model(&p).Updates(fields)

	return p, nil
}

func (s *Sql) Delete(id uint) error {
	p := model.Shoe{
		Model: gorm.Model{
			ID: id,
		},
	}
	s.db.Delete(p)

	return nil
}

func (s *Sql) List() (list []model.Shoe) {
	list = []model.Shoe{}
	s.db.Model(model.Shoe{}).Find(&list)

	return
}

func (s *Sql) Count() int {
	var numOfRecords int
	s.db.Model(model.Shoe{}).Count(numOfRecords)

	return numOfRecords
}
