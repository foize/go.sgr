package tcolor

import (
	"strconv"
)

// Simple code like reset and bold
type basicSgr uint8

func (s basicSgr) code() uint8 {
	return uint8(s)
}

func (s basicSgr) sequence() []byte {
	return []byte(strconv.Itoa(int(s)))
}

const (
	Reset                = basicSgr(0)  // Reset all SGR options.
	ResetForegroundColor = basicSgr(39) // Reset foreground color.
	ResetBackgroundColor = basicSgr(49) // Reset background color.

	Bold    = basicSgr(1)
	BoldOff = basicSgr(22)

	Underline    = basicSgr(4)
	UnderlineOff = basicSgr(24)

	Blink    = basicSgr(5) // Less then 150 per minute.
	BlinkOff = basicSgr(25)

	ImageNegative = basicSgr(7)  // Set reversed-video active (foreground and background negative).
	ImagePositive = basicSgr(27) // Reset reversed-video to normal.

	Framed             = basicSgr(51)
	Encircled          = basicSgr(52)
	FramedEncircledOff = basicSgr(54)

	Overlined    = basicSgr(53)
	OverlinedOff = basicSgr(55)
)
