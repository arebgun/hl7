package hl7

import (
	"github.com/facebookgo/stackerr"
	"strings"
)

type Message []Segment
type Segment []Field
type Field []FieldItem
type FieldItem []Component
type Component []Subcomponent
type Subcomponent string

const (
	segmentSeperator            = "\n"
	fieldSeperator              = "|"
	repeatingFieldSeperator     = "~"
	componentSeperator          = "^"
	repeatingComponentSeperator = "&"
)

func (m Message) Segments(name string) []Segment {
	var a []Segment

	for _, s := range m {
		if string(s[0][0][0][0]) == name {
			a = append(a, s)
		}
	}

	return a
}

func (m Message) Segment(name string, index int) Segment {
	i := 0
	for _, s := range m {
		if string(s[0][0][0][0]) == name {
			if i == index {
				return s
			}

			i++
		}
	}

	return nil
}

func (m Message) Query(query string) (res string, err error) {
	q, err := ParseQuery(query)
	if err != nil {
		return "", stackerr.Wrap(err)
	}

	return m.query(q), nil
}

func (m Message) query(q *Query) string {
	s := m.Segment(q.Segment, q.SegmentOffset)

	return s.query(q)
}

func (m Message) QuerySlice(query string) ([]string, error) {
	q, err := ParseQuery(query)
	if err != nil {
		return []string{}, stackerr.Wrap(err)
	}

	return m.querySlice(q), nil
}

func (m Message) querySlice(q *Query) []string {
	s := m.Segment(q.Segment, q.SegmentOffset)
	return s.querySlice(q)
}

func (m Message) String() string {
	items := []string{}

	for _, s := range m {
		items = append(items, s.String())
	}

	return strings.Join(items, segmentSeperator)
}
