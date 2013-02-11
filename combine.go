package sgr

import (
	"bytes"
	"fmt"
)

// Same as Combine, but panics instead of returning an error.
func MustCombine(parts ...interface{}) string {
	str, err := combine(true, parts...)
	if err != nil {
		panic(err)
	}
	return str
}

// Same as CombineWithoutReset, but panics instead of returning an error.
func MustCombineWithoutReset(parts ...interface{}) string {
	str, err := combine(false, parts...)
	if err != nil {
		panic(err)
	}
	return str
}

// Combines a string 
func Combine(parts ...interface{}) (string, error) {
	return combine(true, parts...)
}

func CombineWithoutReset(parts ...interface{}) (string, error) {
	return combine(false, parts...)
}

func combine(reset bool, parts ...interface{}) (string, error) {
	// Builder used to build the colored string.
	sb := new(sgrBuilder)

	// Loop through all the given parts
	for i, part := range parts {
		if err := sb.append(part); err != nil {
			return "", fmt.Errorf("Error compiling part %d. %s", i, err)
		}
	}

	// Reset sequence if needed.
	if reset {
		if err := sb.append(Reset); err != nil {
			panic(err)
		}
	}

	// All done, return resulting string
	return sb.String(), nil
}
