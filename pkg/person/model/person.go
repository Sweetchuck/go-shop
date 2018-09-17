package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Person struct {
	gorm.Model
	Name string
	Pass string
	Email string
	BirthDate time.Time
}
