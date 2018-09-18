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

	if p1.Name != p2.Name {
		t.Fail()
	}

	if p1.Pass != p2.Pass {
		t.Fail()
	}
}

func TestMemory_Read(t *testing.T) {
	m := Memory{}
	m.Init()

	var p1, p2 model.Person

	p1 = model.Person{
		Name: "a",
		Pass: "b",
	}
	p2, _ = m.Insert(p1)

	if p2.ID != 1 {
		t.Fail()
	}

	p3 := m.Read(1)
	if p3.ID != 1 {
		t.Fail()
	}

	if len(m.people) != 1 {
		t.Fail()
	}
}
