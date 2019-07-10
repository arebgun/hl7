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
			return f.FieldItem(q.FieldItem).SliceOfStrings()
		}

		return f.FieldItem(0).SliceOfStrings()
	}

	return f.FieldItem(q.FieldItem).querySlice(q)
}

func (f Field) FieldItem(index int) FieldItem {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f Field) String() string {
	return strings.Join(f.SliceOfStrings(), repeatingFieldSeperator)
}

func (f Field) SliceOfStrings() []string {
	strs := []string{}

	for _, fi := range f {
		strs = append(strs, fi.String())
	}

	return strs
}
