package storage

import (
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	"testing"
)

func TestMemory_Insert(t *testing.T) {
	m := Memory{}
	m.Init()

	var p1, p2 model.Person

	p1 = model.Person{}
	p2, _ = m.Insert(p1)
	if p2.ID != 1 {
		t.Fail()
	}

	p1 = model.Person{}
	p2, _ = m.Insert(p1)
	if p2.ID != 2 {
		t.Fail()
	}
}
