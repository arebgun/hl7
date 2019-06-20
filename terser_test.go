package hl7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	q := New("MSH", 0, 0, 0, 0, 0)
	a.Equal(q, Query{Segment: "MSH"})
}
