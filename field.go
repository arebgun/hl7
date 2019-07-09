package hl7

import (
	"strings"
)

func (f Field) query(q *Query) string {
	if len(f) <= q.FieldItem {
		return f.String()
	}

	if !q.HasComponent {
		if q.HasFieldItem {
			return f.FieldItem(q.FieldItem).String()
		}

		return f.String()
	}

	return f.FieldItem(q.FieldItem).query(q)
}

func (f Field) querySlice(q *Query) []string {
	if !q.HasComponent {
		if q.HasFieldItem {
			return f.FieldItem(q.FieldItem).Components()
		}

		return f.FieldItem(0).Components()
	}

	return f.FieldItem(q.FieldItem).querySlice(q)
}

func (f Field) FieldItem(index int) FieldItem {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f Field) FieldItems() []string {
	items := []string{}

	for _, fi := range f {
		items = append(items, fi.String())
	}

	return items
}

func (f Field) String() string {
	return strings.Join(f.FieldItems(), repeatingFieldSeperator)
}
