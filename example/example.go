package main

import (
	"fmt"
	"go.sgr"
)

func main() {
	fmt.Println("normal text")

	// Defined colors
	fmt.Printf(sgr.MustCombine(sgr.FgBlack, "%s\n"), "[fg-black]")
	fmt.Printf(sgr.MustCombine(sgr.FgRed, "%s\n"), "[fg-red]")
	fmt.Printf(sgr.MustCombine(sgr.FgGreen, "%s\n"), "[fg-green]")
	fmt.Printf(sgr.MustCombine(sgr.FgYellow, "%s\n"), "[fg-yellow]")
	fmt.Printf(sgr.MustCombine(sgr.FgBlue, "%s\n"), "[fg-blue]")
	fmt.Printf(sgr.MustCombine(sgr.FgMagenta, "%s\n"), "[fg-magenta]")
	fmt.Printf(sgr.MustCombine(sgr.FgCyan, "%s\n"), "[fg-cyan]")
	fmt.Printf(sgr.MustCombine(sgr.FgGrey, "%s\n"), "[fg-grey]")
	fmt.Printf(sgr.MustCombine(sgr.FgWhite, "%s\n"), "[fg-white]")
	fmt.Printf(sgr.MustCombine(sgr.BgBlack, "%s\n"), "[bg-black]")
	fmt.Printf(sgr.MustCombine(sgr.BgRed, "%s\n"), "[bg-red]")
	fmt.Printf(sgr.MustCombine(sgr.BgGreen, "%s\n"), "[bg-green]")
	fmt.Printf(sgr.MustCombine(sgr.BgYellow, "%s\n"), "[bg-yellow]")
	fmt.Printf(sgr.MustCombine(sgr.BgBlue, "%s\n"), "[bg-blue]")
	fmt.Printf(sgr.MustCombine(sgr.BgMagenta, "%s\n"), "[bg-magenta]")
	fmt.Printf(sgr.MustCombine(sgr.BgCyan, "%s\n"), "[bg-cyan]")
	fmt.Printf(sgr.MustCombine(sgr.BgGrey, "%s\n"), "[bg-grey]")
	fmt.Printf(sgr.MustCombine(sgr.BgWhite, "%s\n"), "[bg-white]")

	// Custom colors
	fmt.Printf(sgr.MustCombine(sgr.FgColor(69), "%s\n"), "[fg-69]")
	fmt.Printf(sgr.MustCombine(sgr.BgColor(90), "%s\n"), "[bg-90]")

	// Options
	fmt.Printf(sgr.MustCombine(sgr.Reset, "%s "), "[reset]")
	fmt.Printf(sgr.MustCombine(sgr.ResetForegroundColor, "%s "), "fg-[reset]")
	fmt.Printf(sgr.MustCombine(sgr.ResetBackgroundColor, "%s "), "bg-[reset]")
	fmt.Printf(sgr.MustCombine(sgr.Bold, "%s "), "[bold]")
	fmt.Printf(sgr.MustCombine(sgr.BoldOff, "%s "), "[boldOff]")
	fmt.Printf(sgr.MustCombine(sgr.Underline, "%s "), "[underline]")
	fmt.Printf(sgr.MustCombine(sgr.UnderlineOff, "%s "), "[underlineOff]")
	fmt.Printf(sgr.MustCombine(sgr.Blink, "%s "), "[blink]")
	fmt.Printf(sgr.MustCombine(sgr.BlinkOff, "%s "), "[blinkOff]")
	fmt.Printf(sgr.MustCombine(sgr.ImageNegative, "%s "), "[imageNegative]")
	fmt.Printf(sgr.MustCombine(sgr.ImagePositive, "%s "), "[imagePositive]")
	fmt.Printf(sgr.MustCombine(sgr.Framed, "%s "), "[framed]")
	fmt.Printf(sgr.MustCombine(sgr.Encircled, "%s "), "[encircled]")
	fmt.Printf(sgr.MustCombine(sgr.FramedEncircledOff, "%s "), "[framedEncircledOff]")
	fmt.Printf(sgr.MustCombine(sgr.Overlined, "%s "), "[overlined]")
	fmt.Printf(sgr.MustCombine(sgr.OverlinedOff, "%s "), "[overlinedOff]")

	exampleString := sgr.MustParseln("This is an example: [fg-red bold] important text [reset] normal text again.")
	fmt.Print(exampleString)

	secretNumberFormat := sgr.MustParseln("The secret number is [bg-17 blink]%d")
	fmt.Printf(secretNumberFormat, 42)

	fmt.Println("This text is normal again, MustParseln puts a reset at the end of the line.")
}
