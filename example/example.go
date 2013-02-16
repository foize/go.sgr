package main

import (
	"fmt"
	"go.sgr"
)

func main() {
	fmt.Println("normal text")

	// Defined colors
	fmt.Printf(sgr.FgBlack+"%s\n", "[fg-black]")
	fmt.Printf(sgr.FgRed+"%s\n", "[fg-red]")
	fmt.Printf(sgr.FgGreen+"%s\n", "[fg-green]")
	fmt.Printf(sgr.FgYellow+"%s\n", "[fg-yellow]")
	fmt.Printf(sgr.FgBlue+"%s\n", "[fg-blue]")
	fmt.Printf(sgr.FgMagenta+"%s\n", "[fg-magenta]")
	fmt.Printf(sgr.FgCyan+"%s\n", "[fg-cyan]")
	fmt.Printf(sgr.FgGrey+"%s\n", "[fg-grey]")
	fmt.Printf(sgr.FgWhite+"%s\n", "[fg-white]")
	fmt.Printf(sgr.BgBlack+"%s\n", "[bg-black]")
	fmt.Printf(sgr.BgRed+"%s\n", "[bg-red]")
	fmt.Printf(sgr.BgGreen+"%s\n", "[bg-green]")
	fmt.Printf(sgr.BgYellow+"%s\n", "[bg-yellow]")
	fmt.Printf(sgr.BgBlue+"%s\n", "[bg-blue]")
	fmt.Printf(sgr.BgMagenta+"%s\n", "[bg-magenta]")
	fmt.Printf(sgr.BgCyan+"%s\n", "[bg-cyan]")
	fmt.Printf(sgr.BgGrey+"%s\n", "[bg-grey]")
	fmt.Printf(sgr.BgWhite+"%s\n", "[bg-white]")

	// Custom colors
	fmt.Printf(sgr.FgColor(69)+"%s\n", "[fg-69]")
	fmt.Printf(sgr.BgColor(90)+"%s\n", "[bg-90]")

	// Options
	fmt.Printf(sgr.Reset+"%s ", "[reset]")
	fmt.Printf(sgr.ResetForegroundColor+"%s ", "fg-[reset]")
	fmt.Printf(sgr.ResetBackgroundColor+"%s ", "bg-[reset]")
	fmt.Printf(sgr.Bold+"%s ", "[bold]")
	fmt.Printf(sgr.BoldOff+"%s ", "[boldOff]")
	fmt.Printf(sgr.Underline+"%s ", "[underline]")
	fmt.Printf(sgr.UnderlineOff+"%s ", "[underlineOff]")
	fmt.Printf(sgr.Blink+"%s ", "[blink]")
	fmt.Printf(sgr.BlinkOff+"%s ", "[blinkOff]")
	fmt.Printf(sgr.ImageNegative+"%s ", "[imageNegative]")
	fmt.Printf(sgr.ImagePositive+"%s ", "[imagePositive]")
	fmt.Printf(sgr.Framed+"%s ", "[framed]")
	fmt.Printf(sgr.Encircled+"%s ", "[encircled]")
	fmt.Printf(sgr.FramedEncircledOff+"%s ", "[framedEncircledOff]")
	fmt.Printf(sgr.Overlined+"%s ", "[overlined]")
	fmt.Printf(sgr.OverlinedOff+"%s ", "[overlinedOff]")

	exampleString := sgr.MustParseln("This is an example: [fg-red bold] important text [reset] normal text again.")
	fmt.Print(exampleString)

	secretNumberFormat := sgr.MustParseln("The secret number is [bg-17 blink]%d")
	fmt.Printf(secretNumberFormat, 42)

	fmt.Println("This text is normal again, MustParseln puts a reset at the end of the line.")
}
