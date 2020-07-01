package hl7

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	q := NewQuery("MSH", 0, 0, 0, 0, 0)
	a.Equal(q, Query{Segment: "MSH"})
}

var longTestMessageContent = strings.Replace(`MSH|^~\&||GA0000||VAERS PROCESSOR|20010331605||ORU^R01|20010422GA03|T|2.3.1|||AL|
PID|||1234^^^^SR~1234-12^^^^LR~00725^^^^MR||Doe^John^Fitzgerald^JR^^^L||20001007|M||2106-3^White^HL70005|123 Peachtree St^APT 3B^Atlanta^GA^30210^^M^^GA067||(678) 555-1212^^PRN|
NK1|1|Jones^Jane^Lee^^RN|VAB^Vaccine administered by (Name)^HL70063|
NK1|2|Jones^Jane^Lee^^RN|FVP^Form completed by (Name)-Vaccine provider^HL70063|101 Main Street^^Atlanta^GA^38765^^O^^GA121||(404) 554-9097^^WPN|
ORC|CN|||||||||||1234567^Welby^Marcus^J^Jr^Dr.^MD^L|||||||||Peachtree Clinic|101 Main Street^^Atlanta^GA^38765^^O^^GA121|(404) 554-9097^^WPN|101 Main Street^^Atlanta^GA^38765^^O^^GA121|
OBR|1|||^CDC VAERS-1 (FDA) Report|||20010316|
OBX|1|NM|21612-7^Reported Patient Age^LN||05|mo^month^ANSI|
OBX|1|TS|30947-6^Date form completed^LN||20010316|
OBX|2|FT|30948-4^Vaccination adverse events and treatment, if any^LN|1|fever of 106F, with vomiting, seizures, persistent crying lasting over 3 hours, loss of appetite|
OBX|3|CE|30949-2^Vaccination adverse event outcome^LN|1|E^required emergency room/doctor visit^NIP005|
OBX|4|CE|30949-2^Vaccination adverse event outcome^LN|1|H^required hospitalization^NIP005|
OBX|5|NM|30950-0^Number of days hospitalized due to vaccination adverse event^LN|1|02|d^day^ANSI|
OBX|6|CE|30951-8^Patient recovered^LN||Y^Yes^ HL70239|
OBX|7|TS|30952-6^Date of vaccination^LN||20010216|
OBX|8|TS|30953-4^Adverse event onset date and time^LN||200102180900|
OBX|9|FT|30954-2^Relevant diagnostic tests/lab data^LN||Electrolytes, CBC, Blood culture|
OBR|2|||30955-9^All vaccines given on date listed in #10^LN|
OBX|1|CE30955-9&30956-7^Vaccine type^LN|1|08^HepB-Adolescent/pediatric^CVX|
OBX|2|CE|30955-9&30957-5^Manufacturer^LN|1|MSD^Merck^MVX|
OBX|3|ST|30955-9&30959-1^Lot number^LN|1|MRK12345|
OBX|4|CE|30955-9&30958-3^ Route^LN|1|IM^Intramuscular ^HL70162|
OBX|5|CE|30955-9&31034-2^Site^LN|1|LA^Left arm^ HL70163|
OBX|6|NM|30955-9&30960-9^Number of previous doses^LN|1|01I
OBX|7|CE|CE|30955-9&30956-7^Vaccine type^LN|2|50^DTaP-Hib^CVX|
OBX|8|CE|30955-9&30957-5^ Manufacturer^LN|2|WAL^Wyeth_Ayerst^MVX|
OBX|9|ST|30955-9&30959-1^Lot number^LN|2|W46932777|
OBX|10|CE|30955-9&30958-3^ Route^LN|2|IM^Intramuscular^HL70162|
OBX|11|CE|30955-9&31034-2^Site^LN|2|LA^Left arm^HL70163|
OBX|12|NM|30955-9&30960-9^Number of previous doses^LN|2|01|
OBR|3|||30961-7^Any other vaccinations within 4 weeks prior to the date listed in #10|
OBX|1|CE|30961-7&30956-7^Vaccine type^LN|1|10^IPV^CVX|
OBX|2|CE|30961-7&30957-5^Manufacturer^LN|1|PMC^Aventis Pasteur ^MVX|
OBX|3|ST|30961-7&30959-1^Lot number^LN|1|PMC123456|
OBX|4|CE|30961-7&30958-3^Route^LN|1|SC^Subcutaneaous^HL70162|
OBX|5|CE|30961-7&31034-2^Site^LN|1|LA^Left arm^HL70163|
OBX|6|NM|30961-7&30960-9^Number of previous doses^LN|1|01|
OBX|7|TS|30961-7&31035-9^date given^LN|1|20001216|
OBX|8|CE|30962-^Vaccinated at^LN||PVT^Private doctor�s office/hospital^NIP009|
OBX|9|CE|30963-3^Vaccine purchased with^LN||PBF^Public funds^NIP008|
OBX|10|FT|30964-1^Other medications^LN||None|
OBX|11|FT|30965-8^Illness at time of vaccination (specify)^LN||None|
OBX|12|FT|30966-6^Pre-existing physician diagnosed allergies, birth defects, medical conditions^LN||Past conditions convulsions|
OBX|13|CE|30967-4^Was adverse event reported previously^LN||N^no^NIP009|
OBR|4||30968-2^Adverse event following prior vaccination in patient^LN|
OBX|1|TX|30968-2&30971-6^Adverse event^LN||None|
OBR|5||30969-0^Adverse event following prior vaccination in brother^LN|
OBX|1|TX||30969-0&30971-6^Adverse event^LN||vomiting, fever, otitis media|
OBX|2|NM||30969-0&30972-4^Onset age^LN||04|mo^month^ANSI|
OBX|3|CE||30969-0&30956-7^Vaccine Type ^LN||10^IPV^CVX|
OBX|4|NM||30969-0&30973-2^Dose number in series^LN||02|
OBR|6|||30970-8^Adverse event following prior vaccination in sister^LN|
OBX|1|TX|30970-8&30971-6^Adverse event^LN||None|
OBR|7||^For children 5 and under|
OBX|1|NM|8339-4^Body weight at birth^LN||82|oz^ounces^ANSI|
OBX|2|NM|30974-0^Number of brothers and sisters^LN||2|
OBR|8|||^Only for reports submitted by manufacturer/immunization project|
OBX|1|ST|30975-7^Mfr./Imm. Proj. report no.^LN||12345678|
OBX|2|TS|30976-5^Date received by manufacturer/immunization project^LN||12345678|
OBX|3|CE|30977-3^15 day report^LN||N^No^HL70136|
OBX|4|CE|30978-1^Report type^LN||IN^Initial^NIP010|
TST|1a&2a&3a&4a^5a^6a~1b&2b&3b&4b^5b^6b
`, "\n", "\r", -1)

func TestSegmentParsingViaQuery(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage([]byte(longTestMessageContent))
	a.NoError(err)

	pidSegment := m.Segment("PID", 0)
	a.NotEmpty(pidSegment)

	pidSegmentFields := pidSegment.SliceOfStrigs()
	a.Len(pidSegmentFields, 14)

	firstRFieldStr := pidSegment[3].String()
	a.Equal("1234^^^^SR~1234-12^^^^LR~00725^^^^MR", firstRFieldStr)

	firstRFieldSlice := pidSegment[3].SliceOfStrings()
	a.Equal([]string{"1234^^^^SR", "1234-12^^^^LR", "00725^^^^MR"}, firstRFieldSlice)

	pidRFieldStrViaQuery, err := m.Query("PID-3")
	a.NoError(err)
	a.Equal("1234^^^^SR~1234-12^^^^LR~00725^^^^MR", pidRFieldStrViaQuery)

	pidRFieldSliceViaQuery, err := m.QuerySlice("PID-3")
	a.NoError(err)
	a.Equal([]string{"1234^^^^SR", "1234-12^^^^LR", "00725^^^^MR"}, pidRFieldSliceViaQuery)

	pidRFieldFirstRepetitionStrViaQuery, err := m.Query("PID-3-1")
	a.NoError(err)
	a.Equal("1234~1234-12~00725", pidRFieldFirstRepetitionStrViaQuery)

	pidRFieldFirstRepetitionSliceViaQuery, err := m.QuerySlice("PID-3-1")
	a.NoError(err)
	a.Equal([]string{"1234", "1234-12", "00725"}, pidRFieldFirstRepetitionSliceViaQuery)

	pidRFieldLastRepetitionStrViaQuery, err := m.Query("PID-3-5")
	a.NoError(err)
	a.Equal("SR~LR~MR", pidRFieldLastRepetitionStrViaQuery)

	pidRFieldLastRepetitionSliceViaQuery, err := m.QuerySlice("PID-3-5")
	a.NoError(err)
	a.Equal([]string{"SR", "LR", "MR"}, pidRFieldLastRepetitionSliceViaQuery)

	obx11 := m.Segment("OBX", 11)

	obx3Str, err := obx11.Query("OBX-3")
	a.NoError(err)
	a.Equal("30955-9&30957-5^Manufacturer^LN", obx3Str)

	obx3Slice, err := obx11.QuerySlice("OBX-3")
	a.NoError(err)
	a.Equal([]string{"30955-9&30957-5", "Manufacturer", "LN"}, obx3Slice)

	obx31Str, err := obx11.Query("OBX-3-1")
	a.NoError(err)
	a.Equal("30955-9&30957-5", obx31Str)

	obx31Slice, err := obx11.QuerySlice("OBX-3-1")
	a.NoError(err)
	a.Equal([]string{"30955-9&30957-5"}, obx31Slice)

	obx311Str, err := obx11.Query("OBX-3-1-1")
	a.NoError(err)
	a.Equal("30955-9", obx311Str)

	obx311Slice, err := obx11.QuerySlice("OBX-3-1-1")
	a.NoError(err)
	a.Equal([]string{"30955-9"}, obx311Slice)

	obx312Str, err := obx11.Query("OBX-3-1-2")
	a.NoError(err)
	a.Equal("30957-5", obx312Str)

	obx312Slice, err := obx11.QuerySlice("OBX-3-1-2")
	a.NoError(err)
	a.Equal([]string{"30957-5"}, obx312Slice)

	// OBX 3 >1
	obx3Slice, err = obx11.QuerySlice("OBX-3")
	a.NoError(err)
	a.Equal("Manufacturer^LN", strings.Join(obx3Slice[1:], string(d.Component)))

	// OBX >3
	obx3Str, err = obx11.Query("OBX->3")
	a.NoError(err)
	a.Equal("1|MSD^Merck^MVX", obx3Str)

	obx3Slice, err = obx11.QuerySlice("OBX->3")
	a.NoError(err)
	a.Equal([]string{"1", "MSD^Merck^MVX"}, obx3Slice)

	// OBX 3 >1
	obx3Str, err = obx11.Query("OBX-3->1")
	a.NoError(err)
	a.Equal("Manufacturer^LN", obx3Str)

	obx3Slice, err = obx11.QuerySlice("OBX-3->1")
	a.NoError(err)
	a.Equal([]string{"Manufacturer^LN"}, obx3Slice)

	pidRFieldStrViaQuery, err = m.Query("PID-3->1")
	a.NoError(err)
	a.Equal("^^^SR~^^^LR~^^^MR", pidRFieldStrViaQuery)

	pidRFieldSliceViaQuery, err = m.QuerySlice("PID-3->1")
	a.NoError(err)
	a.Equal([]string{"^^^SR", "^^^LR", "^^^MR"}, pidRFieldSliceViaQuery)
}

type countTestCase struct {
	q string
	c int
	m []byte
}

var countTestCases = []countTestCase{
	countTestCase{"MSH", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX", 47, []byte(longTestMessageContent)},
	countTestCase{"WWW", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)", 16, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)", 7, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(2)", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(2)", 6, []byte(longTestMessageContent)},
	countTestCase{"WWW(2)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(30)", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(30)", 6, []byte(longTestMessageContent)},
	countTestCase{"WWW(30)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH-1", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(2)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(2)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(2)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(30)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(30)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(30)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(2)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(2)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(2)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(30)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(30)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(30)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH-1", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-1", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-2", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-3", 3, []byte(longTestMessageContent)},
	countTestCase{"PID-4", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(1)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(2)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(3)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(4)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-3(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(2)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(3)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(4)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5(1)", 7, []byte(longTestMessageContent)},
	countTestCase{"PID-5(2)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5(3)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-5(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-5(2)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"NK1(1)", 4, []byte(longTestMessageContent)},
	countTestCase{"NK1(2)", 7, []byte(longTestMessageContent)},
	countTestCase{"NK1(3)", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1-1", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1-2-3", 0, []byte(longTestMessageContent)},
}

func TestCount(t *testing.T) {
	for i := range countTestCases {
		c := countTestCases[i]

		t.Run(c.q, func(t *testing.T) {
			a := assert.New(t)

			q, err := ParseQuery(c.q)
			a.NoError(err)

			m, _, err := ParseMessage(c.m)
			a.NoError(err)

			if a.NotNil(q) && a.NotNil(m) {
				l := q.Count(m)
				a.Equal(c.c, l, q.String())
			}
		})
	}
}
