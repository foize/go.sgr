package sgr

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)


var blockCodes = make(map[string]string)

func init() {
	// Options
	blockCodes["reset"] = Reset
	blockCodes["fg-reset"] = ResetForegroundColor
	blockCodes["bg-reset"] = ResetBackgroundColor
	blockCodes["bold"] = Bold
	blockCodes["boldOff"] = BoldOff
	blockCodes["underline"] = Underline
	blockCodes["underlineOff"] = UnderlineOff
	blockCodes["blink"] = Blink
	blockCodes["blinkOff"] = BlinkOff
	blockCodes["imageNegative"] = ImageNegative
	blockCodes["imagePositive"] = ImagePositive
	blockCodes["framed"] = Framed
	blockCodes["encircled"] = Encircled
	blockCodes["framedEncircledOff"] = FramedEncircledOff
	blockCodes["overlined"] = Overlined
	blockCodes["overlinedOff"] = OverlinedOff

	// Foreground Colors
	blockCodes["fg-black"] = FgBlack
	blockCodes["fg-red"] = FgRed
	blockCodes["fg-green"] = FgGreen
	blockCodes["fg-yellow"] = FgYellow
	blockCodes["fg-blue"] = FgBlue
	blockCodes["fg-magenta"] = FgMagenta
	blockCodes["fg-cyan"] = FgCyan
	blockCodes["fg-grey"] = FgGrey
	blockCodes["fg-white"] = FgWhite

	// Background Colors
	blockCodes["bg-black"] = BgBlack
	blockCodes["bg-red"] = BgRed
	blockCodes["bg-green"] = BgGreen
	blockCodes["bg-yellow"] = BgYellow
	blockCodes["bg-blue"] = BgBlue
	blockCodes["bg-magenta"] = BgMagenta
	blockCodes["bg-cyan"] = BgCyan
	blockCodes["bg-grey"] = BgGrey
	blockCodes["bg-white"] = BgWhite
}

func MustParse(format string) string {
	str, err := parse(true, false, format)
	if err != nil {
		panic(err)
	}
	return str
}

func MustParseln(format string) string {
	str, err := parse(true, true, format)
	if err != nil {
		panic(err)
	}
	return str
}

func MustParseWithoutReset(format string) string {
	str, err := parse(false, false, format)
	if err != nil {
		panic(err)
	}
	return str
}

func Parse(format string) (string, error) {
	return parse(true, false, format)
}

func Parseln(format string) (string, error) {
	return parse(true, true, format)
}

func ParseWithoutReset(format string) (string, error) {
	return parse(false, false, format)
}

func parse(reset bool, newline bool, format string) (string, error) {
	// Builder used to build the colored string.
	buf := new(bytes.Buffer)

	// position in the parsing process
	pos := 0

	// index of the currenctly processing square bracket start
	idxStart := 0
	idxEnd := 0

	for {
		// Find next square bracket, break loop when none was found.
		relBlockOpen := strings.IndexRune(format[pos:], '[')
		if relBlockOpen == -1 {
			buf.WriteString(format[pos:])
			break
		}
		idxStart = pos + relBlockOpen

		// Test for escaped square bracket
		if format[idxStart+1] == '[' {
			buf.WriteString(format[pos : idxStart+1])
			pos = idxStart + 2
			continue
		}

		// Add skipped string (if any)
		if idxStart > pos { //idxStart > pos+1 ???
			buf.WriteString(format[pos:idxStart])
		}

		// Find square bracket end
		relBlockClose := strings.IndexRune(format[idxStart:], ']')
		if relBlockClose == -1 {
			return "", fmt.Errorf("Opened square bracket never closed at pos %d. If you want a literal bracket escape it: [[", idxStart)
		}
		idxEnd = idxStart + relBlockClose

		// found a block
		block := format[idxStart+1 : idxEnd]
		fields := strings.Fields(block)
		for _, field := range fields {

			if sgrString, blockCodeExists := blockCodes[field]; blockCodeExists {
				buf.WriteString(sgrString)
				continue
			}

			isFgColor := strings.HasPrefix(field, "fg-")
			isBgColor := strings.HasPrefix(field, "bg-")
			if isFgColor || isBgColor {
				// Check if given number is valid.
				clr, err := strconv.Atoi(field[3:])
				if err != nil {
					return "", fmt.Errorf("Invalid code '%s' in block at position %d.", field, idxStart)
				}
				if clr < 0 || clr > 255 {
					return "", fmt.Errorf("Invalid color code %s. Expecting 0-255 or a defined color.", field[3:])
				}

				// Add color sequence to the buffer
				if isFgColor {
					buf.WriteString(sgrStart + fgColorStart + field[3:] + sgrEnd)
				} else {
					buf.WriteString(sgrStart + bgColorStart + field[3:] + sgrEnd)
				}
				continue // next field
			}

			// Not a valid blockCode and not a fgColor or bgColor
			return "", fmt.Errorf("Invalid code '%s' in block at position %d.", field, idxStart)
		}

		// Change starting position for next iteration.
		pos = idxEnd + 1
	}

	if reset {
		buf.WriteString(Reset)
	}

	if newline {
		buf.WriteString("\n")
	}

	return buf.String(), nil
}
