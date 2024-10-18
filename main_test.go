package main

import (
	// "fmt"
	"testing"
	// "time"
	// "lesson/students"
	// "students"
	// "strings"
	// "sort"
	// "math"
	// "strconv"
	// "testing"
)

func TestLength(t *testing.T) {
	for i, test := range tests {
		size := Length(test.in)
		if size != test.out {
			t.Errorf("#%d: Size(%d)=%s; want %s", i, test.in, size, test.out)
		}
	}
}