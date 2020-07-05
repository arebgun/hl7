package hl7

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustReadFile(f string) []byte {
	d, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return d
}

var (
	allElementsContent = mustReadFile("testdata/all_elements.hl7")
	pdfGeneticsContent = mustReadFile("testdata/pdf_genetics.hl7")
	sampleContent      = mustReadFile("testdata/sample.hl7")
	simpleContent      = mustReadFile("testdata/simple.hl7")
	simpleNohexContent = mustReadFile("testdata/simple_nohex.hl7")
	vaersLongContent   = mustReadFile("testdata/vaers_long.hl7")
)

func TestParseOneSegment(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage([]byte(`MSH|^~\&|IPM|1919|SUPERHOSPITAL|1919|20160101000000||ADT^A08|555544444|D|2.4|||AL|NE`))
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{RepeatedField{Component{"MSH"}}},
			Field{RepeatedField{Component{"|"}}},
			Field{RepeatedField{Component{"^~\\&"}}},
			Field{RepeatedField{Component{"IPM"}}},
			Field{RepeatedField{Component{"1919"}}},
			Field{RepeatedField{Component{"SUPERHOSPITAL"}}},
			Field{RepeatedField{Component{"1919"}}},
			Field{RepeatedField{Component{"20160101000000"}}},
			nil,
			Field{RepeatedField{
				Component{"ADT"},
				Component{"A08"},
			}},
			Field{RepeatedField{Component{"555544444"}}},
			Field{RepeatedField{Component{"D"}}},
			Field{RepeatedField{Component{"2.4"}}},
			nil,
			nil,
			Field{RepeatedField{Component{"AL"}}},
			Field{RepeatedField{Component{"NE"}}},
		},
	}, m)
}

func TestParseTwoSegments(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage([]byte(strings.Join([]string{
		`MSH|^~\&|IPM|1919|SUPERHOSPITAL|1919|20160101000000||ADT^A08|555544444|D|2.4|||AL|NE`,
		`EVN|A08|20160101000001||BATMAN_U|SHBOLTONM^Bolton, Michael^^^^^^|STUFF`,
	}, "\r")))
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{RepeatedField{Component{"MSH"}}},
			Field{RepeatedField{Component{"|"}}},
			Field{RepeatedField{Component{"^~\\&"}}},
			Field{RepeatedField{Component{"IPM"}}},
			Field{RepeatedField{Component{"1919"}}},
			Field{RepeatedField{Component{"SUPERHOSPITAL"}}},
			Field{RepeatedField{Component{"1919"}}},
			Field{RepeatedField{Component{"20160101000000"}}},
			nil,
			Field{RepeatedField{
				Component{"ADT"},
				Component{"A08"},
			}},
			Field{RepeatedField{Component{"555544444"}}},
			Field{RepeatedField{Component{"D"}}},
			Field{RepeatedField{Component{"2.4"}}},
			nil,
			nil,
			Field{RepeatedField{Component{"AL"}}},
			Field{RepeatedField{Component{"NE"}}},
		},
		Segment{
			Field{RepeatedField{Component{"EVN"}}},
			Field{RepeatedField{Component{"A08"}}},
			Field{RepeatedField{Component{"20160101000001"}}},
			nil,
			Field{RepeatedField{Component{"BATMAN_U"}}},
			Field{RepeatedField{
				Component{"SHBOLTONM"},
				Component{"Bolton, Michael"},
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			}},
			Field{RepeatedField{Component{"STUFF"}}},
		},
	}, m)
	evn, err := m.QuerySlice("EVN-5")
	a.NoError(err)
	a.Equal(8, len(evn))
}

func TestParseTwoSegmentsWithLineFeedAndCarriageReturn(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage([]byte("MSH|^~\\&|IPM|1919|SUPERHOSPITAL|1919|20160101000000||ADT^A08|555544444|D|2.4|||AL|NE\r\nEVN|A08|20160101000001||BATMAN_U|SHBOLTONM^Bolton, Michael^^^^^^|STUFF"))
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{RepeatedField{Component{"MSH"}}},
			Field{RepeatedField{Component{"|"}}},
			Field{RepeatedField{Component{"^~\\&"}}},
			Field{RepeatedField{Component{"IPM"}}},
			Field{RepeatedField{Component{"1919"}}},
			Field{RepeatedField{Component{"SUPERHOSPITAL"}}},
			Field{RepeatedField{Component{"1919"}}},
			Field{RepeatedField{Component{"20160101000000"}}},
			nil,
			Field{RepeatedField{
				Component{"ADT"},
				Component{"A08"},
			}},
			Field{RepeatedField{Component{"555544444"}}},
			Field{RepeatedField{Component{"D"}}},
			Field{RepeatedField{Component{"2.4"}}},
			nil,
			nil,
			Field{RepeatedField{Component{"AL"}}},
			Field{RepeatedField{Component{"NE"}}},
		},
		Segment{
			Field{RepeatedField{Component{"EVN"}}},
			Field{RepeatedField{Component{"A08"}}},
			Field{RepeatedField{Component{"20160101000001"}}},
			nil,
			Field{RepeatedField{Component{"BATMAN_U"}}},
			Field{RepeatedField{
				Component{"SHBOLTONM"},
				Component{"Bolton, Michael"},
				nil,
				nil,
				nil,
				nil,
				nil,
				nil,
			}},
			Field{RepeatedField{Component{"STUFF"}}},
		},
	}, m)
	evn, err := m.QuerySlice("EVN-5")
	a.NoError(err)
	a.Equal(8, len(evn))
}

func TestParseSampleContent(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage(sampleContent)
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{RepeatedField{Component{"MSH"}}},
			Field{RepeatedField{Component{"|"}}},
			Field{RepeatedField{Component{"^~\\&"}}},
			Field{RepeatedField{Component{"EP^IC"}}},
			Field{RepeatedField{Component{"EPICADT"}}},
			Field{RepeatedField{Component{"SMS"}}},
			Field{RepeatedField{Component{"SMSADT"}}},
			Field{RepeatedField{Component{"199912271408"}}},
			Field{RepeatedField{Component{"CHARRIS"}}},
			Field{RepeatedField{Component{"ADT"}, Component{"A04"}}},
			Field{RepeatedField{Component{"1817457"}}},
			Field{RepeatedField{Component{"D"}}},
			Field{RepeatedField{Component{"2.5"}}},
		},
		Segment{
			Field{RepeatedField{Component{"PID"}}},
			nil,
			Field{RepeatedField{Component{"0493575"}, nil, nil, Component{"2"}, Component{"ID 1"}}},
			Field{RepeatedField{Component{"454721"}}},
			nil,
			Field{RepeatedField{Component{"DOE"}, Component{"JOHN"}, nil, nil, nil, nil}},
			Field{RepeatedField{Component{"DOE"}, Component{"JOHN"}, nil, nil, nil, nil}},
			Field{RepeatedField{Component{"19480203"}}},
			Field{RepeatedField{Component{"M"}}},
			nil,
			Field{RepeatedField{Component{"B"}}},
			Field{RepeatedField{Component{"254 MYSTREET AVE"}, nil, Component{"MYTOWN"}, Component{"OH"}, Component{"44123"}, Component{"USA"}}},
			nil,
			Field{RepeatedField{Component{"(216)123-4567"}}},
			nil,
			nil,
			Field{RepeatedField{Component{"M"}}},
			Field{RepeatedField{Component{"NON"}}},
			Field{RepeatedField{Component{"400003403"}}, RepeatedField{Component{"1129086"}}}},
		Segment{
			Field{RepeatedField{Component{"NK1"}}},
			nil,
			Field{RepeatedField{Component{"ROE"}, Component{"MARIE"}, nil, nil, nil, nil}},
			Field{RepeatedField{Component{"SPO"}}},
			nil,
			Field{RepeatedField{Component{"(216)123-4567"}}},
			nil,
			Field{RepeatedField{Component{"EC"}}},
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		},
		Segment{
			Field{RepeatedField{Component{"PV1"}}},
			nil,
			Field{RepeatedField{Component{"O"}}},
			Field{RepeatedField{Component{"168 "}}, RepeatedField{Component{"219"}}, RepeatedField{Component{"C"}}, RepeatedField{Component{"PMA"}, nil, nil, nil, nil, nil, nil, nil, nil, nil}},
			nil,
			nil,
			nil,
			Field{RepeatedField{Component{"277"}, Component{"ALLEN MYLASTNAME"}, Component{"BONNIE"}, nil, nil, nil, nil}},
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			Field{RepeatedField{Component{" "}}},
			nil,
			Field{RepeatedField{Component{"2688684"}}},
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			Field{RepeatedField{Component{"199912271408"}}},
			nil,
			nil,
			nil,
			nil,
			nil,
			Field{RepeatedField{Component{"002376853\""}}},
		},
	}, m)
}

func TestParseSimpleNohexContent(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage(simpleNohexContent)
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{RepeatedField{Component{"MSH"}}},
			Field{RepeatedField{Component{"|"}}},
			Field{RepeatedField{Component{"^~\\&"}}},
			Field{RepeatedField{Component{"field"}}},
			Field{RepeatedField{
				Component{"\\|~^&HEY"},
			}},
			Field{RepeatedField{
				Component{"component1"},
				Component{"component2"},
			}},
			Field{RepeatedField{
				Component{"subcomponent1a", "subcomponent2a"},
				Component{"subcomponent1b", "subcomponent2b"},
			}},
			Field{
				RepeatedField{Component{"component1a"}, Component{"component2a"}},
				RepeatedField{Component{"component1b"}, Component{"component2b"}},
			},
		},
	}, m)
}

func TestParseSimpleContent(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage(simpleContent)
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{RepeatedField{Component{"MSH"}}},
			Field{RepeatedField{Component{"|"}}},
			Field{RepeatedField{Component{"^~\\&"}}},
			Field{RepeatedField{Component{"field"}}},
			Field{RepeatedField{
				Component{"\\|~^&\\X484559"},
			}},
			Field{RepeatedField{
				Component{"component1"},
				Component{"component2"},
			}},
			Field{RepeatedField{
				Component{"subcomponent1a", "subcomponent2a"},
				Component{"subcomponent1b", "subcomponent2b"},
			}},
			Field{
				RepeatedField{Component{"component1a"}, Component{"component2a"}},
				RepeatedField{Component{"component1b"}, Component{"component2b"}},
			},
		},
	}, m)
}

func TestParseLiteralNewline(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage([]byte("MSH|^~\\&|Newline\nIn\nContent"))
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{RepeatedField{Component{"MSH"}}},
			Field{RepeatedField{Component{"|"}}},
			Field{RepeatedField{Component{"^~\\&"}}},
			Field{RepeatedField{Component{"Newline\nIn\nContent"}}},
		},
	}, m)
}

func TestParseBad(t *testing.T) {
	a := assert.New(t)

	m, d, err := ParseMessage([]byte("MSH00000"))

	a.Error(err)
	a.Nil(m)
	a.Nil(d)
}

func BenchmarkAllElementsContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(allElementsContent)
	}
}

func BenchmarkSampleContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(sampleContent)
	}
}

func BenchmarkSimpleContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(simpleContent)
	}
}

func BenchmarkSimpleNohexContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(simpleNohexContent)
	}
}

func BenchmarkVaersLongContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(vaersLongContent)
	}
}
