package hl7

import (
	"testing"
	"time"
)

func TestUnmarshal(t *testing.T) {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	p := struct {
		Name struct {
			First string `hl7:"PID-5-2"`
			Last  string `hl7:"PID-5-1"`
		}
		DOB       time.Time `hl7:"PID-7-1"`
		RandomInt int
	}{}

	if err := Unmarshal(m, &p); err != nil {
		t.Error(err)
	}

	expectations := []struct {
		Expected string
		Got      string
	}{
		{
			"John",
			p.Name.First,
		},
		{
			"Doe",
			p.Name.Last,
		},
		{
			"20001007",
			p.DOB.Format("20060102"),
		},
	}

	for _, e := range expectations {
		if e.Expected != e.Got {
			t.Errorf("expected %s; got %s", e.Expected, e.Got)
		}
	}
}
