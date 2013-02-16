package sgr

import (
	"strconv"
)

const (
	SgrStart     = "\x1b["
	FgColorStart = "38;05;"
	BgColorStart = "48;05;"
	SgrEnd       = "m"
)

// Color strings
const (
	FgBlack   = "\x1b[38;05;0m"
	FgRed     = "\x1b[38;05;1m"
	FgGreen   = "\x1b[38;05;2m"
	FgYellow  = "\x1b[38;05;3m"
	FgBlue    = "\x1b[38;05;4m"
	FgMagenta = "\x1b[38;05;5m"
	FgCyan    = "\x1b[38;05;6m"
	FgGrey    = "\x1b[38;05;7m"
	FgWhite   = "\x1b[38;05;255m"

	BgBlack   = "\x1b[48;05;0m"
	BgRed     = "\x1b[48;05;1m"
	BgGreen   = "\x1b[48;05;2m"
	BgYellow  = "\x1b[48;05;3m"
	BgBlue    = "\x1b[48;05;4m"
	BgMagenta = "\x1b[48;05;5m"
	BgCyan    = "\x1b[48;05;6m"
	BgGrey    = "\x1b[48;05;7m"
	BgWhite   = "\x1b[48;05;255m"
)

func FgColor(num uint8) string {
	return SgrStart + FgColorStart + strconv.Itoa(int(num)) + SgrEnd
}

func BgColor(num uint8) string {
	return SgrStart + BgColorStart + strconv.Itoa(int(num)) + SgrEnd
}

// Option strings
const (
	Reset                = "\x1b[0m"  // Reset all SGR options.
	ResetForegroundColor = "\x1b[39m" // Reset foreground color.
	ResetBackgroundColor = "\x1b[49m" // Reset background color.

	Bold    = "\x1b[1m"
	BoldOff = "\x1b[22m"

	Underline    = "\x1b[4m"
	UnderlineOff = "\x1b[24m"

	Blink    = "\x1b[5m" // Less then 150 per minute.
	BlinkOff = "\x1b[25m"

	ImageNegative = "\x1b[7m"  // Set reversed-video active (foreground and background negative).
	ImagePositive = "\x1b[27m" // Reset reversed-video to normal.

	Framed             = "\x1b[51m"
	Encircled          = "\x1b[52m"
	FramedEncircledOff = "\x1b[54m"

	Overlined    = "\x1b[53m"
	OverlinedOff = "\x1b[55m"
)
