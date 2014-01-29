## SGR (Select Graphic Rendition) for Go
The package `sgr` provides methods to use colors and text-decoration in ANSI-escape-sequence compatible terminals. It has support for 256-colors and makes available text-decorations such as bold, underline, negative and blink.

### Installation
Please use go-get to install this package.

`go get github.com/foize/go.sgr`

### Usage
Typical usage is to use Parseln() or MustParseln(), and use the resulting string in succeeding `fmt` or `log` calls.

```go
package main

import (
	"fmt"
	"github.com/foize/go.sgr"
)

func main() {
	sgr.Println("This is an example: [fg-red bold] important text [reset] normal text again.")

	// use sgr.Printf like you're used to with fmt
	// note: you should use `[reset]`  before the newline `\n`
	sgr.Printf("The secret number is [fg-17]%d [reset]\n", 42)

	// you can also parse once and re-use the parsed format string
	// using Parseln eliminates the need for a `[reset]\n` at the end of line
	secretNumberFormat := sgr.MustParseln("The secret number is [bg-17 blink]%d")
	fmt.Printf(secretNumberFormat, 42)

	fmt.Println("This text is normal again, MustParseln puts a reset at the end of the line.")
}
```

### Examples
For more examples, see the [example/example.go file](example/example.go). You can also run the example.

### Escaping
If you want to use an opening square bracket you should escape it like this:

`sgr.MustParseln("foo [[ bar")`

This will print `"foo [ bar"`.

### TODO
Features like the stdlib log package (date time) might be added in a subpackage.

### License
This package is licensed under a [modified BSD 3-clause license](https://github.com/foize/go.sgr/blob/master/LICENSE).

### Godoc
For more information visit [godoc.org/github.com/foize/go.sgr](http://godoc.org/github.com/foize/go.sgr)

### xterm color codes
![ansi sgr color codes](https://raw.github.com/foize/go.sgr/master/xterm_color_chart.png)