package hl7

import (
	"strings"
)

func (c Component) query(q *Query) string {
	if len(c) <= q.RepeatedComponent {
		return c.String()
	}

	return c.RepeatedComponent(q.RepeatedComponent)
}

func (c Component) querySlice(q *Query) []string {
	if !q.HasRepeatedComponent {
		return c.SliceOfStrings()
	}

	return []string{string(c[q.RepeatedComponent])}
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

func (c Component) RepeatedComponent(index int) string {
	if index >= len(c) {
		return ""
	}

	return string(c[index])
}

func (c Component) setString(q *Query, value string) (Component, error) {
	for len(c) < q.RepeatedComponent+1 {
		c = append(c, RepeatedComponent(""))
	}

	c[q.RepeatedComponent] = RepeatedComponent(value)
	return c, nil
}
