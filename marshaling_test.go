package hl7

import (
	"testing"
)

func TestUnmarshal(t *testing.T) {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	patient := struct {
		FirstName string `hl7:"PID-5-2"`
		LastName  string `hl7:"PID-5-1"`
	}{}

	if err := Unmarshal(m, &patient); err != nil {
		t.Error(err)
	}

	expectations := []struct {
		Expected string
		Got      string
	}{
		{
			"John",
			patient.FirstName,
		},
		{
			"Doe",
			patient.LastName,
		},
	}

	for _, e := range expectations {
		if e.Expected != e.Got {
			t.Errorf("expected %s; got %s", e.Expected, e.Got)
		}
	}
}
