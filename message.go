package hl7

import (
	"fmt"
	"strings"
)

type Message []Segment
type Segment []Field
type Field []RepeatedField
type RepeatedField []Component
type Component []RepeatedComponent
type RepeatedComponent string

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
	if index, err := m.segmentIndex(name, offset); err == nil {
		return m[index]
	}

	return nil
}

func (m Message) segmentIndex(name string, offset int) (int, error) {
	currentOffset := 0
	for index, s := range m {
		if string(s[0][0][0][0]) == name {
			if currentOffset == offset {
				return index, nil
			}

			currentOffset++
		}
	}

	return -1, fmt.Errorf("segment %s does not exists", name)
}

func (m Message) Query(query string) (res string, err error) {
	q, err := ParseQuery(query)
	if err != nil {
		return "", err
	}

	return m.query(q)
}

func (m Message) query(q *Query) (string, error) {
	index, err := m.segmentIndex(q.Segment, q.SegmentOffset)
	if err != nil {
		return "", err
	}

	return m[index].query(q)
}

func (m Message) QuerySlice(query string) ([]string, error) {
	q, err := ParseQuery(query)
	if err != nil {
		return []string{}, err
	}

	return m.querySlice(q), nil
}

func (m Message) querySlice(q *Query) []string {
	s := m.Segment(q.Segment, q.SegmentOffset)
	return s.querySlice(q)
}

func (m Message) String() string {
	segmentStrings := []string{}

	for _, s := range m {
		segmentStrings = append(segmentStrings, s.String())
	}

	return strings.Join(segmentStrings, segmentSeperator)
}

func (m *Message) SetString(query string, value string) error {
	q, err := ParseQuery(query)
	if err != nil {
		return err
	}

	index, err := m.segmentIndex(q.Segment, q.SegmentOffset)
	if err != nil {
		return err
	}

	if (*m)[index], err = (*m)[index].setString(q, value); err != nil {
		return err
	}

	return nil
}
