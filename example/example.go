package main

import (
	"fmt"
	"go.sgr"
)

func main() {
	fmt.Println("normal text")

	fmt.Printf(sgr.MustFormat(sgr.FgBlack, "%s\n"), "FgBlack")
	fmt.Printf(sgr.MustFormat(sgr.FgRed, "%s\n"), "FgRed")
	fmt.Printf(sgr.MustFormat(sgr.FgGreen, "%s\n"), "FgGreen")
	fmt.Printf(sgr.MustFormat(sgr.FgYellow, "%s\n"), "FgYellow")
	fmt.Printf(sgr.MustFormat(sgr.FgBlue, "%s\n"), "FgBlue")
	fmt.Printf(sgr.MustFormat(sgr.FgMagenta, "%s\n"), "FgMagenta")
	fmt.Printf(sgr.MustFormat(sgr.FgCyan, "%s\n"), "FgCyan")
	fmt.Printf(sgr.MustFormat(sgr.FgGrey, "%s\n"), "FgGrey")
	fmt.Printf(sgr.MustFormat(sgr.FgWhite, "%s\n"), "FgWhite")
	fmt.Printf(sgr.MustFormat(sgr.BgBlack, "%s\n"), "BgBlack")
	fmt.Printf(sgr.MustFormat(sgr.BgRed, "%s\n"), "BgRed")
	fmt.Printf(sgr.MustFormat(sgr.BgGreen, "%s\n"), "BgGreen")
	fmt.Printf(sgr.MustFormat(sgr.BgYellow, "%s\n"), "BgYellow")
	fmt.Printf(sgr.MustFormat(sgr.BgBlue, "%s\n"), "BgBlue")
	fmt.Printf(sgr.MustFormat(sgr.BgMagenta, "%s\n"), "BgMagenta")
	fmt.Printf(sgr.MustFormat(sgr.BgCyan, "%s\n"), "BgCyan")
	fmt.Printf(sgr.MustFormat(sgr.BgGrey, "%s\n"), "BgGrey")
	fmt.Printf(sgr.MustFormat(sgr.BgWhite, "%s\n"), "BgWhite")

	fmt.Printf(sgr.MustFormat(sgr.Reset, "%s "), "Reset")
	fmt.Printf(sgr.MustFormat(sgr.ResetForegroundColor, "%s "), "ResetForegroundColor")
	fmt.Printf(sgr.MustFormat(sgr.ResetBackgroundColor, "%s "), "ResetBackgroundColor")
	fmt.Printf(sgr.MustFormat(sgr.Bold, "%s "), "Bold")
	fmt.Printf(sgr.MustFormat(sgr.BoldOff, "%s "), "BoldOff")
	fmt.Printf(sgr.MustFormat(sgr.Underline, "%s "), "Underline")
	fmt.Printf(sgr.MustFormat(sgr.UnderlineOff, "%s "), "UnderlineOff")
	fmt.Printf(sgr.MustFormat(sgr.Blink, "%s "), "Blink")
	fmt.Printf(sgr.MustFormat(sgr.BlinkOff, "%s "), "BlinkOff")
	fmt.Printf(sgr.MustFormat(sgr.ImageNegative, "%s "), "ImageNegative")
	fmt.Printf(sgr.MustFormat(sgr.ImagePositive, "%s "), "ImagePositive")
	fmt.Printf(sgr.MustFormat(sgr.Framed, "%s "), "Framed")
	fmt.Printf(sgr.MustFormat(sgr.Encircled, "%s "), "Encircled")
	fmt.Printf(sgr.MustFormat(sgr.FramedEncircledOff, "%s "), "FramedEncircledOff")
	fmt.Printf(sgr.MustFormat(sgr.Overlined, "%s "), "Overlined")
	fmt.Printf(sgr.MustFormat(sgr.OverlinedOff, "%s "), "OverlinedOff")
}
