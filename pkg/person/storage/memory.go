package storage

import (
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	"sort"
	"sync"
	"time"
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

func (m *Memory) Read(id uint) (p model.Person) {
	return m.people[id]
}

func (m *Memory) Update(p model.Person, fields map[string]interface{}) (model.Person, error) {
	for field, value := range fields {
		switch v := value.(type) {
		case string:
			switch field {
			case "Name", "name":
				fallthrough
			case "Pass", "pass":
				fallthrough
			case "Email", "email":
				p.Name = v
			}
		}
	}

	m.Lock()
	m.people[p.ID] = p
	m.Unlock()

	return p, nil
}

func (m *Memory) Delete(id uint) error {
	p, ok := m.people[id]
	if !ok {
		return nil
	}

	now := time.Now()
	p.DeletedAt = &now
	m.people[id] = p

	return nil
}

func (m *Memory) List() []model.Person {
	var list []model.Person
	for key := range m.people {
		if m.people[key].DeletedAt != nil {
			continue
		}

		list = append(list, m.people[key])
	}

	sort.Slice(
		list,
		func(i, j int) bool {
			return list[i].ID < list[j].ID
		},
	)

	return list
}

func (m *Memory) Count() int {
	numOfItems := 0
	for key := range m.people {
		if m.people[key].DeletedAt != nil {
			continue
		}

		numOfItems++
	}

	return numOfItems
}
