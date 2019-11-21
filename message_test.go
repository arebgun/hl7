package hl7

import (
	"strings"
	"testing"
)

func TestMessageString(t *testing.T) {
	var message = strings.Replace(`MSH|^~\&|EP^IC|EPICADT|SMS|SMSADT|199912271408|CHARRIS|ADT^A04|1817457|D|2.5|
PID||0493575^^^2^ID 1|454721||DOE^JOHN^^^^|DOE^JOHN^^^^|19480203|M||B|254 MYSTREET AVE^^MYTOWN^OH^44123^USA||(216)123-4567|||M|NON|400003403~1129086|
NK1||ROE^MARIE^^^^|SPO||(216)123-4567||EC|||||||||||||||||||||||||||
TST|
PV1||O|168 ~219~C~PMA^^^^^^^^^||||277^ALLEN MYLASTNAME^BONNIE^^^^|||||||||| ||2688684|||||||||||||||||||||||||199912271408||||||002376853`, "\n", "\r", -1)

	m, _, _ := ParseMessage([]byte(message))

	expected := `MSH|^~\&|EP^IC|EPICADT|SMS|SMSADT|199912271408|CHARRIS|ADT^A04|1817457|D|2.5`
	got := m.Segment("MSH", 0).String()

	// expected := `^~\&`
	// got, _ := m.Query("MSH-2")

	if got != expected {
		t.Errorf("got %s; expected %s", got, expected)
	}
}
