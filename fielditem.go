package hl7

import (
	"strings"
)

func (f FieldItem) query(q *Query) string {
	if len(f) <= q.Component {
		return f.String()
	}

	return f.Component(q.Component).query(q)
}

func (f FieldItem) querySlice(q *Query) []string {
	if !q.HasComponent {
		return f.Components()
	}

	return f.Component(q.Component + 1).querySlice(q)
}
func (f FieldItem) Component(index int) Component {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f FieldItem) Components() []string {
	items := []string{}

	for _, c := range f {
		items = append(items, c.String())
	}

	return items
}

func (f FieldItem) String() string {
	return strings.Join(f.Components(), componentSeperator)
}
