package hl7

import (
	// "fmt"
	// "reflect"
	"strings"
	"testing"
)

var message = strings.Replace(`MSH|^~\&|EP\S\IC|EPICADT|SMS|SMSADT|199912271408|CHARRIS|ADT^A04|1817457|D|2.5|
PID||0493575^^^2^ID 1|454721||DOE^JOHN^^^^|DOE^JOHN^^^^|19480203|M||B|254 MYSTREET AVE^^MYTOWN^OH^44123^USA||(216)123-4567|||M|NON|400003403~1129086|
NK1||ROE^MARIE^^^^|SPO||(216)123-4567||EC|||||||||||||||||||||||||||
PV1||O|168 ~219~C~PMA^^^^^^^^^||||277^ALLEN MYLASTNAME^BONNIE^^^^|||||||||| ||2688684|||||||||||||||||||||||||199912271408||||||002376853`, "\n", "\r", -1)

func TestString(t *testing.T) {
	m, _, _ := ParseMessage([]byte(message))

	expected := "PID||0493575^^^2^ID 1|454721||DOE^JOHN^^^|DOE^JOHN^^^|19480203|M||B|254 MYSTREET AVE^^MYTOWN^OH^44123^USA||(216)123-4567|||M|NON|400003403~1129086"
	got := m.Segment("PID", 0).String()

	if got != expected {
		t.Errorf("got %s; expected %s", got, expected)
	}
}

func TestSliceOfStrings(t *testing.T) {
	equal := func(a, b []string) bool {
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}

	m, _, _ := ParseMessage([]byte(message))

	expected := []string{"PID", "", "0493575^^^2^ID 1", "454721", "", "DOE^JOHN^^^", "DOE^JOHN^^^", "19480203", "M", "", "B", "254 MYSTREET AVE^^MYTOWN^OH^44123^USA", "", "(216)123-4567", "", "", "M", "NON", "400003403~1129086"}
	got := m.Segment("PID", 0).SliceOfStrigs()

	if !equal(expected, got) {
		t.Errorf("got %s; expected %s", got, expected)
	}
}

// [PID  0493575^^^2^ID 1 454721  DOE^JOHN^^^ DOE^JOHN^^^ 19480203 M  B 254 MYSTREET AVE^^MYTOWN^OH^44123^USA  (216)123-4567   M NON 400003403~1129086]
// [PID  0493575^^^2^ID 1 454721  DOE^JOHN^^^ DOE^JOHN^^^ 19480203 M  B 254 MYSTREET AVE^^MYTOWN^OH^44123^USA  (216)123-4567   M MON 400003403~1129086]
