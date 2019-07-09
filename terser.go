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

	HasFieldItem bool
	FieldItem    int

	HasComponent bool
	Component    int

	HasSubComponent bool
	SubComponent    int
}

func New(segment string, segmentOffset, field, fieldItem, component, subComponent int) Query {
	return Query{
		Segment:       segment,
		SegmentOffset: max(segmentOffset-1, 0),
		Field:         max(field-1, 0),
		FieldItem:     max(fieldItem-1, 0),
		Component:     max(component-1, 0),
		SubComponent:  max(subComponent-1, 0),
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

	if q.HasFieldItem {
		s += fmt.Sprintf("(%d)", q.FieldItem+1)
	}

	if !q.HasComponent {
		return s
	}

	s += fmt.Sprintf("-%d", q.Component+1)

	if !q.HasSubComponent {
		return s
	}

	s += fmt.Sprintf("-%d", q.SubComponent+1)

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
	if !q.HasFieldItem && !q.HasComponent {
		return len(f)
	}

	if len(f) <= q.FieldItem {
		return 0
	}
	fi := f[q.FieldItem]
	if !q.HasComponent {
		return len(fi)
	}

	if len(fi) <= q.Component {
		return 0
	}
	c := fi[q.Component]
	if !q.HasSubComponent {
		return len(c)
	}

	if len(c) <= q.SubComponent {
		return 0
	}

	return 1
}
