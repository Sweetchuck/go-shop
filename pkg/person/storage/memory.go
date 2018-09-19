package storage

import (
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/person/model"
	"sort"
	"sync"
	"time"
)

type Memory struct {
	items map[uint]model.Person
	sync.Mutex
	lastID uint
}

func (m *Memory) Init() {
	m.items = map[uint]model.Person{}
}

func (m *Memory) Insert(p model.Person) (model.Person, error) {
	m.Lock()
	m.lastID++
	p.ID = m.lastID
	m.items[p.ID] = p
	m.Unlock()

	return p, nil
}

func (m *Memory) Read(id uint) (p model.Person) {
	return m.items[id]
}

func (m *Memory) Update(p model.Person, fields map[string]interface{}) (model.Person, error) {
	for field, value := range fields {
		switch v := value.(type) {
		case string:
			switch field {
			case "Name", "name":
				p.Name = v

			case "Pass", "pass":
				p.Pass = v

			case "Email", "email":
				p.Email = v
			}
		}
	}

	m.Lock()
	m.items[p.ID] = p
	m.Unlock()

	return p, nil
}

func (m *Memory) Delete(id uint) error {
	p, ok := m.items[id]
	if !ok {
		return nil
	}

	now := time.Now()
	p.DeletedAt = &now
	m.items[id] = p

	return nil
}

func (m *Memory) List() []model.Person {
	var list []model.Person
	for id := range m.items {
		if m.items[id].DeletedAt != nil {
			continue
		}

		list = append(list, m.items[id])
	}

	sort.Slice(
		list,
		func (i, j int) bool {
			return list[i].ID < list[j].ID
		},
	)

	return list
}

func (m *Memory) Count() int {
	numOfItems := 0
	for key := range m.items {
		if m.items[key].DeletedAt != nil {
			continue
		}

		numOfItems++
	}

	return numOfItems
}
