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
		return f.SliceOfStrings()
	}

	return f.Component(q.Component + 1).querySlice(q)
}

func (f FieldItem) Component(index int) Component {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f FieldItem) ComponentPtr(index int) *Component {
	if index >= len(f) {
		return nil
	}

	return &f[index]
}

func (f FieldItem) String() string {
	return strings.Join(f.SliceOfStrings(), componentSeperator)
}

func (f FieldItem) SliceOfStrings() []string {
	strs := []string{}

	for _, c := range f {
		strs = append(strs, c.String())
	}

	return strs
}

func (f FieldItem) setString(q *Query, value string) (FieldItem, error) {
	var err error

	for len(f) < q.Component+1 {
		f = append(f, Component{})
	}

	f[q.Component], err = f[q.Component].setString(q, value)

	return f, err
}
