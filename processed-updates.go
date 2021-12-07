package main

type ProcessedUpdates struct {
	ids      [500]int
	position int
}

func NewProcessedUpdates() *ProcessedUpdates {
	return &ProcessedUpdates{
		ids:      [500]int{},
		position: 0,
	}
}

func (m *ProcessedUpdates) add(id int) {
	if m.position > len(m.ids) {
		m.position = 0
	} else {
		m.position++
	}
	m.ids[m.position] = id
}

func (m *ProcessedUpdates) exists(id int) bool {
	for _, proccesedId := range m.ids {
		if proccesedId == 0 {
			return false
		} else if proccesedId == id {
			return true
		}
	}
	return false
}
