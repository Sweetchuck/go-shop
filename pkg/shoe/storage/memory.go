package storage

import (
	"gitlab.cheppers.com/devops-academy-2018/shop2/pkg/shoe/model"
	"sort"
	"sync"
)

type Memory struct {
	shoes map[uint]model.Shoe
	sync.Mutex
	lastID uint
}

func (m *Memory) Init() {
	m.shoes = map[uint]model.Shoe{}
}

func (m *Memory) Insert(s model.Shoe) (model.Shoe, error) {
	m.Lock()
	m.lastID++
	s.ID = m.lastID
	m.shoes[s.ID] = s
	m.Unlock()

	return s, nil
}

func (m *Memory) Read(id uint) (p model.Shoe) {
	return m.shoes[id]
}

func (m *Memory) Update(p model.Shoe, fields map[string]interface{}) (model.Shoe, error) {
	m.Lock()
	m.shoes[p.ID] = p
	m.Unlock()

	return p, nil
}

func (m *Memory) Delete(id uint) error {
	delete(m.shoes, id)

	return nil
}

func (m *Memory) List() ([]model.Shoe) {
	var list []model.Shoe
	for k, _ := range m.shoes {
		if m.shoes[k].DeletedAt != nil {
			continue
		}

		list = append(list, m.shoes[k])
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
	return len(m.shoes)
}
