package utils

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	cases := []string{
		"http://example.com",
		"https://example.com",
		"http://example.com/path/to/resource",
		"https://example.com/path/to/resource?query=param",
		"http://example.com:8080/path/to/resource",
		"https://example.com/path/to/resource?query=param&another=param2",
	}
	for _, v := range cases {
		if match, err := RegexpMatchUrlFromString(v); err != nil {
			t.Fatal(err)
		} else {
			fmt.Println(match)
		}
	}
}
