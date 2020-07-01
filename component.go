package hl7

import (
	"fmt"
	"strings"
)

func (c Component) query(q *Query) (string, error) {
	if len(c) <= q.RepeatedComponent {
		return "", fmt.Errorf("component %d does not have repeated component %d for query %s", q.Component, q.RepeatedComponent, q.String())
	}

	if !q.HasRepeatedComponent {
		return c.String(), nil
	}
	return string(c[q.RepeatedComponent]), nil
}

func (c Component) querySlice(q *Query) []string {
	if !q.HasRepeatedComponent {
		return c.SliceOfStrings()
	}

	return []string{string(c[q.RepeatedComponent])}
}

func (c Component) String() string {
	return strings.Join(c.SliceOfStrings(), string(repeatingComponentSeperator))
}

func (c Component) SliceOfStrings() []string {
	strs := []string{}

	for _, s := range c {
		strs = append(strs, string(s))
	}

	return strs
}

func (c Component) setString(q *Query, value string) (Component, error) {
	for len(c) < q.RepeatedComponent+1 {
		c = append(c, RepeatedComponent(""))
	}

	c[q.RepeatedComponent] = RepeatedComponent(value)
	return c, nil
}
