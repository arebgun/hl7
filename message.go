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

func (m Message) Query(s string) (res string, ok bool, err error) {
	q, err := ParseQuery(s)
	if err != nil {
		return "", false, stackerr.Wrap(err)
	}

	res, ok = q.FromMessage(m)

	return res, ok, nil
}

func (m Message) ToString() string {
	items := []string{}

	for _, s := range m {
		items = append(items, s.ToString())
	}

	return strings.Join(items, segmentSeperator)
}

func (s Segment) Field(index int) Field {
	if index >= len(s) {
		return nil
	}

	return s[index]
}

func (s Segment) ToString() string {
	items := []string{}

	for _, f := range s {
		items = append(items, f.ToString())
	}

	return strings.Join(items, fieldSeperator)
}

func (f Field) FieldItem(index int) FieldItem {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f Field) Component(index int) Component {
	if index >= len(f.FieldItem(0)) {
		return nil
	}

	return f.FieldItem(0)[index]
}

func (f Field) ToString() string {
	items := []string{}

	for _, fi := range f {
		items = append(items, fi.ToString())
	}

	return strings.Join(items, repeatingFieldSeperator)
}

func (f FieldItem) Component(index int) Component {
	if index >= len(f) {
		return nil
	}

	return f[index]
}

func (f FieldItem) ToString() string {
	items := []string{}

	for _, s := range f {
		items = append(items, s.ToString())
	}

	return strings.Join(items, componentSeperator)
}

func (c Component) ToString() string {
	items := []string{}

	for _, s := range c {
		items = append(items, string(s))
	}

	return strings.Join(items, repeatingComponentSeperator)
}
