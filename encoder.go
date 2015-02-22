package main

import (
	"bytes"
	"fmt"
)

// Converts a []byte{0x00, 0xff} into a "0x00, 0xFF"
func bytesToString(data []byte) string {
	buf := new(bytes.Buffer)
	for i, b := range data {
		if i%20 == 0 {
			fmt.Fprintf(buf, "\n\t\t")
		}
		fmt.Fprintf(buf, "0x%02X,", b)
	}

	// remove trailing ","
	return buf.String()
}