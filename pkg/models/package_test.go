package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUniqDepend(t *testing.T) {
	var (
		expected []string
		result   []string
		input    string
	)

	expected = []string{"libc6"}
	input = "libc6 (>= 2.8)"
	result = getUniqDepends(input)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("\nExpected: %v\nGot: %v", expected, result)
	}

	expected = []string{"python"}
	input = "python (>= 2.7.1.0ubuntu2.1), python (<< 2.8)"
	result = getUniqDepends(input)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("\nExpected: %v\nGot: %v", expected, result)
	}

	expected = []string{"openssl", "debconf", "debconf-2.0"}
	input = "openssl (>= 1.0.0), debconf (>= 0.5) | debconf-2.0"
	result = getUniqDepends(input)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("\nExpected: %v\nGot: %v", expected, result)
	}

	input = "netbase, libc6 (>= 2.11), libgcc1 (>= 1:4.1.1), libncurses5 (>= 5.6+20071006-3), libstdc++6 (>= 4.1.1)"
	expected = []string{"netbase", "libc6", "libgcc1", "libncurses5", "libstdc++6"}
	result = getUniqDepends(input)

	fmt.Println("Result", result)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("\nExpected: %v\nGot: %v", expected, result)
	}
}
