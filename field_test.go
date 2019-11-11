package hl7

import (
	"testing"
)

func TestSetLastName(t *testing.T) {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	expectectedLastName := "Wehr"

	if err := m.SetString("PID-5-1", expectectedLastName); err != nil {
		t.Errorf("we caught: %s", err.Error())
	}

	lastName, err := m.Query("PID-5-1")
	if err != nil {
		t.Errorf("we caught: %s", err.Error())
	}

	if expectectedLastName != lastName {
		t.Errorf("expected %s; got %s", expectectedLastName, lastName)
	}
}

func TestSetExternalID(t *testing.T) {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	expectedID := "12345"

	if err := m.SetString("PID-2", expectedID); err != nil {
		t.Errorf("we caught: %s", err.Error())
	}

	ID, err := m.Query("PID-2")
	if err != nil {
		t.Errorf("we caught: %s", err.Error())
	}

	if expectedID != ID {
		t.Errorf("expected %s; got %s", expectedID, ID)
	}
}
