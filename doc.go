/*
Typical usage is to use Parseln() or MustParseln(), and use the resulting string in succeeding `fmt` or `log` calls.

	package main

	import (
		"fmt"
		"github.com/foize/go.sgr"
	)

	func main() {
		exampleString := sgr.MustParseln("This is an example: [fg-red bold] important text [reset] normal text again.")
		fmt.Print(exampleString)

		secretNumberFormat := sgr.MustParseln("The secret number is [bg-17 blink]%d")
		fmt.Printf(secretNumberFormat, 42)

		fmt.Println("This text is normal again, MustParseln puts a reset at the end of the line.")
	}

If you want to use an opening square bracket you should escape it like this:

	sgr.MustParseln("foo [[ bar")

This will print give you "foo [ bar".
*/
package sgr
