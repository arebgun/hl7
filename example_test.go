package hl7

import (
	"fmt"
)

func ExampleQuery_GetString() {
	m, _, _ := ParseMessage([]byte(longTestMessageContent))

	msh9_1, _ := ParseQuery("MSH-9-1")
	msh9_2, _ := ParseQuery("MSH-9-2")
	msh9_1_str, _ := msh9_1.StringFromMessage(m)
	msh9_2_str, _ := msh9_2.StringFromMessage(m)

	fmt.Printf("%s_%s", msh9_1_str, msh9_2_str)
	// Output: ORU_R01
}
