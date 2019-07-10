package hl7

import (
	"github.com/facebookgo/stackerr"
	"strings"
)

func (s Segment) Query(query string) (string, error) {
	q, err := ParseQuery(query)

	if err != nil {
		return "", stackerr.Wrap(err)
	}

	return s.query(q), nil
}

func (s Segment) query(q *Query) string {
	if len(s) <= q.Field+1 {
		return ""
	}

	if !q.HasField {
		return s.String()
	}

	return s.Field(q.Field + 1).query(q)
}

func (s Segment) QuerySlice(query string) ([]string, error) {
	q, err := ParseQuery(query)

	if err != nil {
		return []string{}, stackerr.Wrap(err)
	}

	return s.querySlice(q), nil
}

func (s Segment) querySlice(q *Query) []string {
	if !q.HasField {
		return s.SliceOfStrigs()
	}

	return s.Field(q.Field + 1).querySlice(q)
}

func (s Segment) Field(index int) Field {
	if index >= len(s) {
		return nil
	}

	return s[index]
}

func (s Segment) String() string {
	return strings.Join(s.SliceOfStrigs(), fieldSeperator)
}

func (s Segment) SliceOfStrigs() []string {
	strs := []string{}

	for _, f := range s {
		strs = append(strs, f.String())
	}

	return strs
}
