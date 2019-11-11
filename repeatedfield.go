package hl7

import (
	"strings"
)

func (f RepeatedField) query(q *Query) string {
	if len(f) <= q.Component {
		return f.String()
	}

	return f.Component(q.Component).query(q)
}

func (f RepeatedField) querySlice(q *Query) []string {
	if !q.HasComponent {
		return f.SliceOfStrings()
	}

	return f.Component(q.Component + 1).querySlice(q)
}

func (f RepeatedField) Component(index int) Component {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f RepeatedField) ComponentPtr(index int) *Component {
	if index >= len(f) {
		return nil
	}

	return &f[index]
}

func (f RepeatedField) String() string {
	return strings.Join(f.SliceOfStrings(), componentSeperator)
}

func (f RepeatedField) SliceOfStrings() []string {
	strs := []string{}

	for _, c := range f {
		strs = append(strs, c.String())
	}

	return strs
}

func (f RepeatedField) setString(q *Query, value string) (RepeatedField, error) {
	var err error

	for len(f) < q.Component+1 {
		f = append(f, Component{})
	}

	f[q.Component], err = f[q.Component].setString(q, value)

	return f, err
}
