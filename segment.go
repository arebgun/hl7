package hl7

import (
	"github.com/facebookgo/stackerr"
	"strings"
)

func (s Segment) Query(query string) (string, error) {
	if q, err := ParseQuery(query); err != nil {
		return "", stackerr.Wrap(err)
	} else {
		return s.query(q), nil
	}
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
	if q, err := ParseQuery(query); err != nil {
		return []string{}, stackerr.Wrap(err)
	} else {
		return s.querySlice(q), nil
	}
}

func (s Segment) querySlice(q *Query) []string {
	if !q.HasField {
		return s.Fields()
	}

	return s.Field(q.Field + 1).querySlice(q)
}

func (s Segment) Field(index int) Field {
	if index >= len(s) {
		return nil
	}

	return s[index]
}

func (s Segment) Fields() []string {
	items := []string{}

	for _, f := range s {
		items = append(items, f.String())
	}

	return items
}

func (s Segment) String() string {
	return strings.Join(s.Fields(), fieldSeperator)
}
