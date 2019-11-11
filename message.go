package hl7

import (
	"errors"
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

func (m Message) Segment(name string, offset int) Segment {
	if index, err := m.SegmentIndex(name, offset); err == nil {
		return m[index]
	}

	return nil
}

func (m Message) SegmentIndex(name string, offset int) (int, error) {
	currentOffset := 0
	for index, s := range m {
		if string(s[0][0][0][0]) == name {
			if currentOffset == offset {
				return index, nil
			}

			currentOffset++
		}
	}

	return -1, errors.New("Segment does not exists")
}

func (m *Message) SetSegment(name string, offset int, s Segment) error {
	currentOffset := 0
	for i := range *m {
		if string((*m)[i][0][0][0][0]) == name {
			if currentOffset == offset {
				(*m)[i] = s
				return nil
			}

			currentOffset++
		}
	}

	return errors.New("Cannot set segment; segment does not exist")
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

	index, err := m.SegmentIndex(q.Segment, q.SegmentOffset)

	if err != nil {
		return stackerr.Wrap(err)
	}

	if (*m)[index], err = (*m)[index].setString(q, value); err != nil {
		return stackerr.Wrap(err)
	}

	return nil
}
