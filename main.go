package main

import (
	"fmt"
	// "strconv"
	// "strings"
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

type Test struct {
	in  string
	outInt  int
	outErr error
}

var tests = []Test{
	{"hello", 5, nil},
	{"Angrey", 6, nil},
	{"", 0, nil},
}

func GetUTFLength(input []byte) (int, error) {
	return 4, nil
}

func TestGetUTFLength(t *testing.T) {
	for _, testCase := range tests {
		gotInt, gotErr := GetUTFLength([]byte(testCase.in))
		if gotInt != testCase.outInt || gotErr != testCase.outErr {
			t.Fatalf("invalid utf8")
		}
	}

}

func main() {
	// ticker := time.Tick(1 * time.Second)
    // fmt.Println("AA")
	// for a := range ticker{
	// 	fmt.Println("EFF", ticker, "grgerg", a)
	// }
	fmt.Println(len("abcdefghigklmnopqrstuvzxcqqwertyuioplkjhgfdsazxcvbnm"))
}