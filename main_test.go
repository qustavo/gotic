package main

import "testing"

func TestBytesToString(t *testing.T) {
	bytes := []byte{0x00, 0xaa, 0xff}
	expected := "0x00,0xAA,0xFF,"
	got := bytesToString(bytes)

	if got != expected {
		t.Errorf("got: '%s', expected: '%s'\n", got, expected)
	}
}
