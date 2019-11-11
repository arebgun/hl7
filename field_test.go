package hl7

import (
	"testing"
)

func TestFieldSetString(t *testing.T) {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	expectectedLastName := "Wehr"
	expectedBirthdate := "19900915"

	if err := m.SetString("PID-5-1", expectectedLastName); err != nil {
		t.Errorf("we caught: %s", err.Error())
	}

	if err := m.SetString("PID-7", expectedBirthdate); err != nil {
		t.Errorf("we caught: %s", err.Error())
	}

	lastName, err := m.Query("PID-5-1")
	if err != nil {
		t.Errorf("we caught: %s", err.Error())
	}

	birthdate, err := m.Query("PID-7")
	if err != nil {
		t.Errorf("we caught: %s", err.Error())
	}

	if expectectedLastName != lastName {
		t.Errorf("expected %s; got %s", expectectedLastName, lastName)
	}

	if expectedBirthdate != birthdate {
		t.Errorf("expected %s; got %s", expectedBirthdate, birthdate)
	}
}
