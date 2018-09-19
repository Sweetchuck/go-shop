package storage

import (
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/model"
	"sort"
	"sync"
	"time"
)

type Memory struct {
	items map[uint]model.Shoe
	sync.Mutex
	lastID uint
}

func (m *Memory) Init() {
	m.items = map[uint]model.Shoe{}
}

func (m *Memory) Insert(s model.Shoe) (model.Shoe, error) {
	m.Lock()
	m.lastID++
	s.ID = m.lastID
	m.items[s.ID] = s
	m.Unlock()

	return s, nil
}

func (m *Memory) Read(id uint) (p model.Shoe) {
	return m.items[id]
}

func (m *Memory) Update(p model.Shoe, fields map[string]interface{}) (model.Shoe, error) {
	for field, value := range fields {
		switch v := value.(type) {
		case string:
			switch field {
			case "Brand", "brand":
				p.Brand = v

			case "Type", "type":
				p.Type = v
			}
		case int:
			switch field {
			case "Size", "size":
				p.Size = v
			}
		case bool:
			switch field {
			case "WaterProof", "Waterproof", "waterproof":
				p.WaterProof = v
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

func (m *Memory) List() ([]model.Shoe) {
	var list []model.Shoe
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
