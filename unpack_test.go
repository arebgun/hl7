package hl7

import (
	"testing"
	"time"
)

func TestUnpackStrings(t *testing.T) {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	result := struct {
		Patient struct {
			First string    `hl7:"PID-5-2"`
			Last  string    `hl7:"PID-5-1"`
			DOB   time.Time `hl7:"PID-7"`
		}
	}{}

	if err := m.Unpack(&result); err != nil {
		t.Error(err)
	}

	expectations := []struct {
		Expected string
		Got      string
	}{
		{
			"John",
			result.Patient.First,
		},
		{
			"Doe",
			result.Patient.Last,
		},
		{
			"20001007",
			result.Patient.DOB.Format("20060102"),
		},
	}

	for _, e := range expectations {
		if e.Expected != e.Got {
			t.Errorf("expected %s; got %s", e.Expected, e.Got)
		}
	}
}

func TestUnpackInts(t *testing.T) {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	result := struct {
		Kin struct {
			Int  int  `hl7:"NK1-1"`
			Uint uint `hl7:"NK1-1"`
		}
	}{}

	s := m.Segment("NK1", 0)

	if err := s.Unpack(&result); err != nil {
		t.Error(err)
	}

	expectedInt := int(1)

	if expectedInt != result.Kin.Int {
		t.Errorf("expected %d; got %d", expectedInt, result.Kin.Int)
	}

	expectedUint := uint(1)

	if expectedUint != result.Kin.Uint {
		t.Errorf("expected %d; got %d", expectedUint, result.Kin.Uint)
	}
}

func TestUnpackSlices(t *testing.T) {
	equal := func(a, b []string) bool {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	result := struct {
		Address []string `hl7:"PID-11"`
	}{}

	if err := m.Unpack(&result); err != nil {
		t.Error(err)
	}

	expected := []string{"123 Peachtree St", "APT 3B", "Atlanta", "GA", "30210", "", "M", "", "GA067"}

	if !equal(expected, result.Address) {
		t.Errorf("expected %s; got %s", expected, result.Address)
	}
}
