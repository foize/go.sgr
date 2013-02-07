package tcolor

import (
	"strconv"
)

// Complexer code with extra parameters
type Color uint8

type FgColor Color

var fgColorStart = []byte("38;05;")

func (fgc FgColor) SgrCode() uint8 {
	return 38
}

func (fgc FgColor) Sequence() []byte {
	return append(fgColorStart, []byte(strconv.Itoa(int(fgc)))...)
}

type BgColor Color

var bgColorStart = []byte("48;05;")

func (bgc BgColor) SgrCode() uint8 {
	return 48
}

func (bgc BgColor) Sequence() []byte {
	return append(bgColorStart, []byte(strconv.Itoa(int(bgc)))...)
}

const (
	FgBlack   = FgColor(0)
	FgRed     = FgColor(1)
	FgGreen   = FgColor(2)
	FgYellow  = FgColor(3)
	FgBlue    = FgColor(4)
	FgMagenta = FgColor(5)
	FgCyan    = FgColor(6)
	FgGrey    = FgColor(7)
	FgWhite   = FgColor(255)

	BgBlack   = BgColor(0)
	BgRed     = BgColor(1)
	BgGreen   = BgColor(2)
	BgYellow  = BgColor(3)
	BgBlue    = BgColor(4)
	BgMagenta = BgColor(5)
	BgCyan    = BgColor(6)
	BgGrey    = BgColor(7)
	BgWhite   = BgColor(255)
)
