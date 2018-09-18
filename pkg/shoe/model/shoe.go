package model

import (
	"github.com/jinzhu/gorm"
)

type Shoe struct {
	gorm.Model
	Brand string
	Type string
	Size int
	WaterProof bool
}
