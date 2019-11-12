package hl7

import (
	"strings"
	"testing"
)

func TestPackStrings(t *testing.T) {
	var message = strings.Replace(`MSH|^~\&|EP\S\IC|EPICADT|SMS|SMSADT|199912271408|CHARRIS|ADT^A04|1817457|D|2.5|
PID||0493575^^^2^ID 1|454721||DOE^JOHN^^^^||19480203|M||B|254 MYSTREET AVE^^MYTOWN^OH^44123^USA||(216)123-4567|||M|NON|400003403~1129086|
NK1||ROE^MARIE^^^^|SPO||(216)123-4567||EC|||||||||||||||||||||||||||
PV1||O|168 ~219~C~PMA^^^^^^^^^||||277^ALLEN MYLASTNAME^BONNIE^^^^|||||||||| ||2688684|||||||||||||||||||||||||199912271408||||||002376853`, "\n", "\r", -1)

	m, _, _ := ParseMessage([]byte(message))

	patient := struct {
		First string `hl7:"PID-5-2"`
		Last  string `hl7:"PID-5-1"`
	}{
		"Nathan", "Wehr",
	}

	if err := m.Pack(&patient); err != nil {
		t.Error(err)
	}

	expected := "Wehr"
	got, _ := m.Query("PID-5-1")

	if expected != got {
		t.Errorf("expected %s; got %s", expected, got)
	}
}
