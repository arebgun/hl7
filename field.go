package hl7

import (
	"errors"
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

func (f Field) FieldItemPtr(index int) *FieldItem {
	if index >= len(f) {
		return nil
	}

	return &f[index]
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

func (f *Field) setString(q *Query, value string) error {
	if len(*f) <= q.FieldItem {
		return errors.New("Not enough field items")
	}

	if !q.HasFieldItem {
		if q.HasComponent {
			return (*f)[0].ComponentPtr(q.Component).setString(q, value)
		}

		(*f)[0] = FieldItem{Component{Subcomponent(value)}}
		return nil
	}

	return f.FieldItemPtr(q.FieldItem).setString(q, value)
}
