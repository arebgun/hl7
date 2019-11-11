package hl7

import (
	"fmt"
)

type Query struct {
	Segment string

	HasSegmentOffset bool
	SegmentOffset    int

	HasField bool
	Field    int

	HasRepeatedField bool
	RepeatedField    int

	HasComponent bool
	Component    int

	HasRepeatedComponent bool
	RepeatedComponent    int
}

func NewQuery(segment string, segmentOffset, field, repeatedField, component, repeatedComponent int) Query {
	return Query{
		Segment:           segment,
		SegmentOffset:     max(segmentOffset-1, 0),
		Field:             max(field-1, 0),
		RepeatedField:     max(repeatedField-1, 0),
		Component:         max(component-1, 0),
		RepeatedComponent: max(repeatedComponent-1, 0),
	}
}

func (q Query) String() string {
	s := q.Segment

	if q.HasSegmentOffset {
		s += fmt.Sprintf("(%d)", q.SegmentOffset+1)
	}

	if !q.HasField {
		return s
	}

	s += fmt.Sprintf("-%d", q.Field+1)

	if q.HasRepeatedField {
		s += fmt.Sprintf("(%d)", q.RepeatedField+1)
	}

	if !q.HasComponent {
		return s
	}

	s += fmt.Sprintf("-%d", q.Component+1)

	if !q.HasRepeatedComponent {
		return s
	}

	s += fmt.Sprintf("-%d", q.RepeatedComponent+1)

	return s
}

func (q Query) Count(m Message) int {
	if !q.HasSegmentOffset && !q.HasField {
		return len(m.Segments(q.Segment))
	}

	s := m.Segment(q.Segment, q.SegmentOffset)
	if !q.HasField {
		return len(s)
	}

	if len(s) <= q.Field+1 {
		return 0
	}
	f := s[q.Field+1]
	if !q.HasRepeatedField && !q.HasComponent {
		return len(f)
	}

	if len(f) <= q.RepeatedField {
		return 0
	}
	fi := f[q.RepeatedField]
	if !q.HasComponent {
		return len(fi)
	}

	if len(fi) <= q.Component {
		return 0
	}
	c := fi[q.Component]
	if !q.HasRepeatedComponent {
		return len(c)
	}

	if len(c) <= q.RepeatedComponent {
		return 0
	}

	return 1
}

func (q Query) Spew() {
	println("HasSegmentOffset: ", q.HasSegmentOffset)
	println("HasField: ", q.HasField)
	println("HasRepeatedField: ", q.HasRepeatedField)
	println("HasComponent: ", q.HasComponent)
	println("HasRepeatedComponent: ", q.HasRepeatedComponent)
}
