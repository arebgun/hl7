package hl7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parserTestPair struct {
	s string
	q Query
}

var parserTestCases = []parserTestPair{
	parserTestPair{"MSH", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		Field:             0,
		RepeatedField:     0,
		Component:         0,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH(1)", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		HasSegmentOffset:  true,
		Field:             0,
		RepeatedField:     0,
		Component:         0,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH(1)-2", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		HasSegmentOffset:  true,
		Field:             1,
		HasField:          true,
		RepeatedField:     0,
		Component:         0,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH(1)-2(3)", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		HasSegmentOffset:  true,
		Field:             1,
		HasField:          true,
		RepeatedField:     2,
		HasRepeatedField:  true,
		Component:         0,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH(1)-2(3)-4", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		HasSegmentOffset:  true,
		Field:             1,
		HasField:          true,
		RepeatedField:     2,
		HasRepeatedField:  true,
		Component:         3,
		HasComponent:      true,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH-2", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		Field:             1,
		HasField:          true,
		RepeatedField:     0,
		Component:         0,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH-2(3)", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		Field:             1,
		HasField:          true,
		RepeatedField:     2,
		HasRepeatedField:  true,
		Component:         0,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH-2(3)-4", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		Field:             1,
		HasField:          true,
		RepeatedField:     2,
		HasRepeatedField:  true,
		Component:         3,
		HasComponent:      true,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH-2-4", Query{
		Segment:           "MSH",
		SegmentOffset:     0,
		Field:             1,
		HasField:          true,
		RepeatedField:     0,
		Component:         3,
		HasComponent:      true,
		RepeatedComponent: 0,
	}},
	parserTestPair{"MSH-2-4-6", Query{
		Segment:              "MSH",
		SegmentOffset:        0,
		Field:                1,
		HasField:             true,
		RepeatedField:        0,
		Component:            3,
		HasComponent:         true,
		RepeatedComponent:    5,
		HasRepeatedComponent: true,
	}},
	{"MSH->4", Query{
		Segment:       "MSH",
		SegmentOffset: 0,
		Field:         3,
		HasField:      true,
		IsRestField:   true,
	}},
	{"MSH-2->1", Query{
		Segment:         "MSH",
		SegmentOffset:   0,
		Field:           1,
		HasField:        true,
		Component:       0,
		HasComponent:    true,
		IsRestComponent: true,
	}},
	{"MSH-2-4->6", Query{
		Segment:                 "MSH",
		SegmentOffset:           0,
		Field:                   1,
		HasField:                true,
		RepeatedField:           0,
		Component:               3,
		HasComponent:            true,
		RepeatedComponent:       5,
		HasRepeatedComponent:    true,
		IsRestRepeatedComponent: true,
	}},
}

func TestParseQuery(t *testing.T) {
	for i := range parserTestCases {
		c := parserTestCases[i]

		t.Run(c.s, func(t *testing.T) {
			a := assert.New(t)

			q, err := ParseQuery(c.s)
			a.NoError(err)

			if a.NotNil(q) {
				a.Equal(c.q, *q, c.s)
			}
		})
	}
}

func BenchmarkQuery(b *testing.B) {
	for i := range parserTestCases {
		c := parserTestCases[i]

		b.Run(c.s, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ParseQuery(c.s)
			}
		})
	}
}
