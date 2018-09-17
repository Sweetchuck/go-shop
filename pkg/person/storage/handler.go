package storage

import "gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"

type Handler interface {
	Insert(p model.Person) (model.Person, error)
	Update(p model.Person, fields map[string]interface{}) (model.Person, error)
	Delete(id uint) error
	List() []model.Person
	Read(id uint) model.Person
}
