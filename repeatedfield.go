package hl7

import (
	"fmt"
	"strings"
)

func (f RepeatedField) query(q *Query) (string, error) {
	if len(f) <= q.Component {
		return "", fmt.Errorf("repeated field %d does not have component %d for query %s", q.RepeatedField, q.Component, q.String())
	}

	return f[q.Component].query(q)
}

func (f RepeatedField) querySlice(q *Query) []string {
	if !q.HasComponent {
		return f.SliceOfStrings()
	}

	return f[q.Component].querySlice(q)
}

func (f RepeatedField) String() string {
	return strings.Join(f.SliceOfStrings(), string(componentSeperator))
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
