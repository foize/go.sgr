package main

import (
	"fmt"
	"go.sgr"
)

func main() {
	fmt.Println("normal text")

	// usage examples:
	exampleString := sgr.MustParseln("This is an example: [fg-red bold] important text [reset] normal text again.")
	fmt.Print(exampleString)

	secretNumberFormat := sgr.MustParseln("The secret number is [bg-17 blink]%d")
	fmt.Printf(secretNumberFormat, 42)

	fmt.Println("This text is normal again, MustParseln puts a reset sequence at the end of the result string.")

	// Defined colors
	fmt.Print(sgr.MustParseln("\n[bold]Colors by name:"))
	fmt.Printf(sgr.MustParseln("setting foregground to black:   [fg-black]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting foregground to red:     [fg-red]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting foregground to green:   [fg-green]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting foregground to yellow:  [fg-yellow]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting foregground to blue:    [fg-blue]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting foregground to magenta: [fg-magenta]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting foregground to cyan:    [fg-cyan]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting foregground to grey:    [fg-grey]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting foregground to white:   [fg-white]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to black:    [bg-black]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to red:      [bg-red]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to green:    [bg-green]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to yellow:   [bg-yellow]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to blue:     [bg-blue]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to magenta:  [bg-magenta]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to cyan:     [bg-cyan]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to grey:     [bg-grey]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("setting background to white:    [bg-white]%s"), "foobar")

	// Custom colors
	fmt.Print(sgr.MustParseln("\n[bold]Colors by number:"))
	fmt.Printf(sgr.MustParseln("Setting foregroud color to 69:  [fg-69]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("Setting background color to 90: [bg-90]%s"), "foobar")

	// Options
	fmt.Print(sgr.MustParseln("\n[bold]Options:"))
	fmt.Printf(sgr.MustParseln("doing Reset                [reset]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing ResetForegroundColor [fg-reset]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing ResetBackgroundColor [bg-reset]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing Bold                 [bold]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing BoldOff              [boldOff]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing Underline            [underline]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing UnderlineOff         [underlineOff]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing Blink                [blink]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing BlinkOff             [blinkOff]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing ImageNegative        [imageNegative]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing ImagePositive        [imagePositive]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing Framed               [framed]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing Encircled            [encircled]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing FramedEncircledOff   [framedEncircledOff]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing Overlined            [overlined]%s"), "foobar")
	fmt.Printf(sgr.MustParseln("doing OverlinedOff         [overlinedOff]%s"), "foobar")
}
