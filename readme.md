## SGR (Select Graphic Rendition) for Go
The package go.sgr provides methods to do text-coloring in an ANSI-escape-sequence compatible terminal. It also provides several other options such as bold, underline, negative and blink.

### Installation
Please use go-get to install this package.

`go get github.com/foize/go.sgr`

### Usage
```go
package main

import (
	"fmt"
	"github.com/foize/go.sgr"
)

func main() {
	exampleString := sgr.MustParseln("This is an example: [fgRed bold] important text [reset] normal text again.")
	fmt.Print(exampleString)

	secretNumberFormat := sgr.MustParseln("The secret number is [bg-17 blink]%d")
	fmt.Printf(secretNumberFormat, 42)

	fmt.Println("This text is normal again, MustParseln puts a reset at the end of the line.")
}
```

### xterm color codes
![ansi sgr color codes](https://raw.github.com/foize/go.sgr/master/xterm_color_chart.png)