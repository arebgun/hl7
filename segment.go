package hl7

import (
	"fmt"
	"strings"
)

func (s Segment) Query(query string) (string, error) {
	q, err := ParseQuery(query)

	if err != nil {
		return "", err
	}

	return s.query(q)
}

func (s Segment) query(q *Query) (string, error) {
	if !q.HasField {
		return "", fmt.Errorf("field not specified for query %s", q.String())
	}

	if len(s) <= q.Field+1 {
		return "", fmt.Errorf("segment %s does not have field %d for query %s", q.Segment, q.Field+1, q.String())
	}

	if q.IsRestField {
		if len(s) <= q.Field+2 {
			return "", fmt.Errorf("segment %s does not have field %d for query %s", q.Segment, q.Field+2, q.String())
		}
		return strings.Join(s.SliceOfStrigs()[q.Field+2:], string(fieldSeperator)), nil
	}

	return s[q.Field+1].query(q)
}

func (s Segment) QuerySlice(query string) ([]string, error) {
	q, err := ParseQuery(query)

	if err != nil {
		return []string{}, err
	}

	return s.querySlice(q), nil
}

func (s Segment) querySlice(q *Query) []string {
	if !q.HasField {
		return s.SliceOfStrigs()
	}

	if q.IsRestField {
		if len(s) <= q.Field+2 {
			return []string{}
		}
		return s.SliceOfStrigs()[q.Field+2:]
	}
	return s[q.Field+1].querySlice(q)
}

func (s Segment) String() string {
	if s.Tag() == "MSH" {
		return "MSH" + string(fieldSeperator) + strings.Join(s[2:].SliceOfStrigs(), string(fieldSeperator))
	}

	return strings.Join(s.SliceOfStrigs(), string(fieldSeperator))
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
		return nil, fmt.Errorf("No field defined")
	}

	var err error

	for len(s) < q.Field+2 {
		s = append(s, Field{})
	}

	s[q.Field+1], err = s[q.Field+1].setString(q, value)

	return s, err
}

func (s Segment) Tag() string {
	return string(s[0][0][0][0])
}
