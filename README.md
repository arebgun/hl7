hl7
===

[![GoDoc](https://godoc.org/fknsrs.biz/p/hl7?status.svg)](https://godoc.org/fknsrs.biz/p/hl7)

Overview
--------

HL7v2 stands for "Health Level 7: Version 2" - a specification for shuttling
clinical data around and between medical institutions. While working on
[Medtasker](http://medtasker.com/) with [Nimblic](https://github.com/nimblic),
I've written this library for reading the protocol and querying the messages
it contains.

I have a [blog post](https://www.fknsrs.biz/blog/golang-hl7-library.html) on
my website if you'd like to hear a bit more of the story.

Install
-------

```
$ go get github.com/nwehr/hl7
```

New Features
------------

This fork adds the ability to `Unmarshal()` a `Message` into a struct using tags

```go
package main

import (
	"fmt"
	"github.com/nwehr/hl7"
	"strings"
	"time"
)

var testMessage = strings.Replace(`MSH|^~\&||GA0000||VAERS PROCESSOR|20010331605||ORU^R01|20010422GA03|T|2.3.1|||AL|
PID|||1234^^^^SR~1234-12^^^^LR~00725^^^^MR||Doe^John^Fitzgerald^JR^^^L||20001007|M||2106-3^White^HL70005|||(678) 555-1212^^PRN|
NK1|1|Jones^Jane^Lee^^RN|VAB^Vaccine administered by (Name)^HL70063|
NK1|2|Jones^Jane^Lee^^RN|FVP^Form completed by (Name)-Vaccine provider^HL70063|||(404) 554-9097^^WPN|
ORC|CN|||||||||||1234567^Welby^Marcus^J^Jr^Dr.^MD^L|||||||||Peachtree Clinic||(404) 554-9097^^WPN||
OBR|1|||^CDC VAERS-1 (FDA) Report|||20010316|
OBX|1|NM|21612-7^Reported Patient Age^LN||05|mo^month^ANSI|
`, "\n", "\r", -1)

func main() {
	m, _, _ := hl7.ParseMessage([]byte(testMessage))

	patient := struct {
		Name struct {
			First string `hl7:"PID-5-2"`
			Last  string `hl7:"PID-5-1"`
		}
		DOB time.Time `hl7:"PID-7-1"`
	}{}

	hl7.Unmarshal(m, &patient)

	fmt.Printf("%+v", patient)
}

```

License
-------

3-clause BSD. A copy is included with the source.
