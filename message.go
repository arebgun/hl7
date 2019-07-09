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

func (m Message) Query(query string) (res string, ok bool, err error) {
	q, err := ParseQuery(query)
	if err != nil {
		return "", false, stackerr.Wrap(err)
	}

	res, ok = q.StringFromMessage(m)

	return res, ok, nil
}

func (m Message) QuerySlice(query string) ([]string, bool, error) {
	q, err := ParseQuery(query)
	if err != nil {
		return []string{}, false, stackerr.Wrap(err)
	}

	return m.querySlice(q), true, nil
}

func (m Message) querySlice(q *Query) []string {
	s := m.Segment(q.Segment, q.SegmentOffset)
	return s.querySlice(q)
}

func (s Segment) QuerySlice(query string) ([]string, bool, error) {
	q, err := ParseQuery(query)
	if err != nil {
		return []string{}, false, stackerr.Wrap(err)
	}

	return s.querySlice(q), true, nil
}

func (s Segment) querySlice(q *Query) []string {
	if !q.HasField {
		return s.Fields()
	}

	return s.Field(q.Field + 1).querySlice(q)
}

func (f Field) querySlice(q *Query) []string {
	if !q.HasFieldItem {
		return f.FieldItems()
	}

	return f.FieldItem(q.FieldItem).querySlice(q)
}

func (f FieldItem) querySlice(q *Query) []string {
	if !q.HasComponent {
		return f.Components()
	}

	return f.Component(q.Component + 1).querySlice(q)
}

func (c Component) querySlice(q *Query) []string {
	if !q.HasSubComponent {
		return c.Subcomponents()
	}

	return []string{string(c[q.SubComponent])}
}

func (m Message) String() string {
	items := []string{}

	for _, s := range m {
		items = append(items, s.String())
	}

	return strings.Join(items, segmentSeperator)
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

func (f Field) FieldItem(index int) FieldItem {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f Field) FieldItems() []string {
	items := []string{}

	for _, fi := range f {
		items = append(items, fi.String())
	}

	return items
}

func (f Field) String() string {
	return strings.Join(f.FieldItems(), repeatingFieldSeperator)
}

func (f Field) Component(index int) Component {
	if index >= len(f.FieldItem(0)) {
		return nil
	}

	return f.FieldItem(0)[index]
}

func (f Field) Components() []string {
	items := []string{}

	for _, c := range f[0] {
		items = append(items, c.String())
	}

	return items
}

func (f FieldItem) Component(index int) Component {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f FieldItem) Components() []string {
	items := []string{}

	for _, c := range f {
		items = append(items, c.String())
	}

	return items
}

func (f FieldItem) String() string {
	return strings.Join(f.Components(), componentSeperator)
}

func (c Component) Subcomponents() []string {
	items := []string{}

	for _, s := range c {
		items = append(items, string(s))
	}

	return items
}

func (c Component) String() string {
	return strings.Join(c.Subcomponents(), repeatingComponentSeperator)
}
