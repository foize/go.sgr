package main

import (
	"fmt"
	"go.sgr"
)

func main() {
	fmt.Println("normal text")

	// Defined colors
	fmt.Printf(sgr.MustCombine(sgr.FgBlack, "%s\n"), "FgBlack")
	fmt.Printf(sgr.MustCombine(sgr.FgRed, "%s\n"), "FgRed")
	fmt.Printf(sgr.MustCombine(sgr.FgGreen, "%s\n"), "FgGreen")
	fmt.Printf(sgr.MustCombine(sgr.FgYellow, "%s\n"), "FgYellow")
	fmt.Printf(sgr.MustCombine(sgr.FgBlue, "%s\n"), "FgBlue")
	fmt.Printf(sgr.MustCombine(sgr.FgMagenta, "%s\n"), "FgMagenta")
	fmt.Printf(sgr.MustCombine(sgr.FgCyan, "%s\n"), "FgCyan")
	fmt.Printf(sgr.MustCombine(sgr.FgGrey, "%s\n"), "FgGrey")
	fmt.Printf(sgr.MustCombine(sgr.FgWhite, "%s\n"), "FgWhite")
	fmt.Printf(sgr.MustCombine(sgr.BgBlack, "%s\n"), "BgBlack")
	fmt.Printf(sgr.MustCombine(sgr.BgRed, "%s\n"), "BgRed")
	fmt.Printf(sgr.MustCombine(sgr.BgGreen, "%s\n"), "BgGreen")
	fmt.Printf(sgr.MustCombine(sgr.BgYellow, "%s\n"), "BgYellow")
	fmt.Printf(sgr.MustCombine(sgr.BgBlue, "%s\n"), "BgBlue")
	fmt.Printf(sgr.MustCombine(sgr.BgMagenta, "%s\n"), "BgMagenta")
	fmt.Printf(sgr.MustCombine(sgr.BgCyan, "%s\n"), "BgCyan")
	fmt.Printf(sgr.MustCombine(sgr.BgGrey, "%s\n"), "BgGrey")
	fmt.Printf(sgr.MustCombine(sgr.BgWhite, "%s\n"), "BgWhite")

	// Custom colors
	fmt.Printf(sgr.MustCombine(sgr.FgColor(69), "%s\n"), "FgColor(69)")
	fmt.Printf(sgr.MustCombine(sgr.BgColor(90), "%s\n"), "BgColor(90)")

	// Options
	fmt.Printf(sgr.MustCombine(sgr.Reset, "%s "), "Reset")
	fmt.Printf(sgr.MustCombine(sgr.ResetForegroundColor, "%s "), "ResetForegroundColor")
	fmt.Printf(sgr.MustCombine(sgr.ResetBackgroundColor, "%s "), "ResetBackgroundColor")
	fmt.Printf(sgr.MustCombine(sgr.Bold, "%s "), "Bold")
	fmt.Printf(sgr.MustCombine(sgr.BoldOff, "%s "), "BoldOff")
	fmt.Printf(sgr.MustCombine(sgr.Underline, "%s "), "Underline")
	fmt.Printf(sgr.MustCombine(sgr.UnderlineOff, "%s "), "UnderlineOff")
	fmt.Printf(sgr.MustCombine(sgr.Blink, "%s "), "Blink")
	fmt.Printf(sgr.MustCombine(sgr.BlinkOff, "%s "), "BlinkOff")
	fmt.Printf(sgr.MustCombine(sgr.ImageNegative, "%s "), "ImageNegative")
	fmt.Printf(sgr.MustCombine(sgr.ImagePositive, "%s "), "ImagePositive")
	fmt.Printf(sgr.MustCombine(sgr.Framed, "%s "), "Framed")
	fmt.Printf(sgr.MustCombine(sgr.Encircled, "%s "), "Encircled")
	fmt.Printf(sgr.MustCombine(sgr.FramedEncircledOff, "%s "), "FramedEncircledOff")
	fmt.Printf(sgr.MustCombine(sgr.Overlined, "%s "), "Overlined")
	fmt.Printf(sgr.MustCombine(sgr.OverlinedOff, "%s "), "OverlinedOff")

	str, err := sgr.Parseln("[fgRed] blabla [[ escaped? [fg-12] ]] escaped? [bgYellow] yellow bg")
	if err != nil {
		fmt.Printf("Error on parsing: %s\n", err)
		return
	}
	fmt.Printf("Parsed string: ```%s``` < end parsed string.\n", str)
}
