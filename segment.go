package hl7

import (
	"errors"
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

func (s Segment) setString(q *Query, value string) (Segment, error) {
	if !q.HasField {
		return nil, errors.New("No field defined")
	}

	var err error

	for len(s) < q.Field+2 {
		s = append(s, Field{})
	}

	s[q.Field+1], err = s[q.Field+1].setString(q, value)

	return s, err
}
