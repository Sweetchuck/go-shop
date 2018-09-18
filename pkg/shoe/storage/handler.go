package storage

import "gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/model"

type Handler interface {
	Insert(p model.Shoe) (model.Shoe, error)
	Read(id uint) model.Shoe
	Update(p model.Shoe, fields map[string]interface{}) (model.Shoe, error)
	Delete(id uint) error
	List() []model.Shoe
	Count() int
}
