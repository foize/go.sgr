package sgr

import (
	"fmt"
)

func parsePrintArguments(a []interface{}) {
	for i, b := range a {
		switch b.(type) {
		case string:
			a[i] = (interface{})(MustParse(b.(string)))
		}
	}
}

// Print wraps fmt.Print
// All arguments with type string are parsed before calling fmt.Print
func Print(a ...interface{}) (n int, err error) {
	parsePrintArguments(a)
	return fmt.Print(a...)
}

// Println wraps fmt.Println
// All arguments with type string are parsed before calling fmt.Println
func Println(a ...interface{}) (n int, err error) {
	parsePrintArguments(a)
	return fmt.Println(a...)
}

// Printf wraps fmt.Printf
// The format string is parsed before fmt.Printf is called
func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(MustParse(format), a...)
}
