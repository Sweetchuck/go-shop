package storage

import (
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	"sync"
)

type Memory struct {
	people map[uint]model.Person
	sync.Mutex
	lastID uint
}

func (m *Memory) Init() {
	m.people = map[uint]model.Person{}
}

func (m *Memory) Insert(p model.Person) (model.Person, error) {
	m.Lock()
	m.lastID++
	p.ID = m.lastID
	m.people[p.ID] = p
	m.Unlock()

	return p, nil
}

func (m *Memory) Update(p model.Person, fields map[string]interface{}) (model.Person, error) {
	m.Lock()
	m.people[p.ID] = p
	m.Unlock()

	return p, nil
}

func (m *Memory) Delete(id uint) error {
	delete(m.people, id)

	return nil
}

func (m *Memory) List() ([]model.Person) {
	list := make([]model.Person, len(m.people))
	for _, p := range m.people {
		if p.DeletedAt != nil {
			continue
		}

		list = append(list, p)
	}

	return list
}

func (m *Memory) Read(id uint) (p model.Person) {
	return m.people[id]
}
