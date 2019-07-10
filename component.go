package hl7

import (
	"strings"
)

func (c Component) query(q *Query) string {
	if len(c) <= q.SubComponent {
		return c.String()
	}

	return c.Subcomponent(q.SubComponent)
}

func (c Component) querySlice(q *Query) []string {
	if !q.HasSubComponent {
		return c.SliceOfStrings()
	}

	return []string{string(c[q.SubComponent])}
}

func (c Component) Subcomponent(index int) string {
	if index >= len(c) {
		return ""
	}

	return string(c[index])
}

func (c Component) String() string {
	return strings.Join(c.SliceOfStrings(), repeatingComponentSeperator)
}

func (c Component) SliceOfStrings() []string {
	strs := []string{}

	for _, s := range c {
		strs = append(strs, string(s))
	}

	return strs
}
