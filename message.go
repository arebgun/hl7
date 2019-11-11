package hl7

import (
	_ "fmt"
	_ "github.com/davecgh/go-spew/spew"
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

func (m Message) SegmentPtr(name string, index int) *Segment {
	i := 0
	for _, s := range m {
		if string(s[0][0][0][0]) == name {
			if i == index {
				return &s
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
	return strings.Join(m.SliceOfStrings(), segmentSeperator)
}

func (m Message) SliceOfStrings() []string {
	strs := []string{}

	for _, s := range m {
		strs = append(strs, s.String())
	}

	return strs
}

func (m *Message) SetString(query string, value string) error {
	q, err := ParseQuery(query)

	if err != nil {
		return stackerr.Wrap(err)
	}

	// queryStruct := struct {
	// 	HasSegmentOffet bool
	// 	HasField        bool
	// 	HasFieldItem    bool
	// 	HasComponent    bool
	// 	HasSubcomponent bool
	// }{
	// 	q.HasSegmentOffset,
	// 	q.HasField,
	// 	q.HasFieldItem,
	// 	q.HasComponent,
	// 	q.HasSubComponent,
	// }

	// spew.Dump(queryStruct)

	return m.setString(q, value)
}

func (m *Message) setString(q *Query, value string) error {
	m.SegmentPtr(q.Segment, q.SegmentOffset).setString(q, value)

	return nil
}
