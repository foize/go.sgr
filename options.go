package sgr

import (
	"strconv"
)

// Defines a sgr option. Implements the `Sgr` interface.
type sgrOption uint8

func (s sgrOption) code() uint8 {
	return uint8(s)
}

func (s sgrOption) sequence() []byte {
	return []byte(strconv.Itoa(int(s)))
}

const (
	Reset                = sgrOption(0)  // Reset all SGR options.
	ResetForegroundColor = sgrOption(39) // Reset foreground color.
	ResetBackgroundColor = sgrOption(49) // Reset background color.

	Bold    = sgrOption(1)
	BoldOff = sgrOption(22)

	Underline    = sgrOption(4)
	UnderlineOff = sgrOption(24)

	Blink    = sgrOption(5) // Less then 150 per minute.
	BlinkOff = sgrOption(25)

	ImageNegative = sgrOption(7)  // Set reversed-video active (foreground and background negative).
	ImagePositive = sgrOption(27) // Reset reversed-video to normal.

	Framed             = sgrOption(51)
	Encircled          = sgrOption(52)
	FramedEncircledOff = sgrOption(54)

	Overlined    = sgrOption(53)
	OverlinedOff = sgrOption(55)
)
