package storage

import (
	"github.com/jinzhu/gorm"
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/model"
)

type SQL struct {
	db *gorm.DB
}

func (s *SQL) Init(db *gorm.DB) {
	s.db = db
}

func (s *SQL) Insert(p model.Shoe) (model.Shoe, error) {
	pp := &p

	s.db.Model(p).Create(pp)

	return *pp, nil
}

func (s *SQL) Read(id uint) (p model.Shoe) {
	s.db.Model(model.Shoe{}).First(&p)

	return
}

func (s *SQL) Update(p model.Shoe, fields map[string]interface{}) (model.Shoe, error) {
	s.db.Model(&p).Updates(fields)

	return p, nil
}

func (s *SQL) Delete(id uint) error {
	p := model.Shoe{
		Model: gorm.Model{
			ID: id,
		},
	}
	s.db.Delete(p)

	return nil
}

func (s *SQL) List() (list []model.Shoe) {
	list = []model.Shoe{}
	s.db.Model(model.Shoe{}).Find(&list)

	return
}

func (s *SQL) Count() int {
	numOfRecords := new(int)
	s.db.Model(model.Shoe{}).Count(numOfRecords)

	return *numOfRecords
}
