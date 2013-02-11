package sgr

import (
	"fmt"
	"strconv"
	"strings"
)

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
	sb := new(sgrBuilder)

	// position in the parsing process
	pos := 0

	// index of the currenctly processing square bracket start
	idxStart := 0
	idxEnd := 0

	for {
		// Find next square bracket, break loop when none was found.
		relBlockOpen := strings.IndexRune(format[pos:], '[')
		if relBlockOpen == -1 {
			sb.appendString(format[pos:])
			break
		}
		idxStart = pos + relBlockOpen

		// Test for escaped square bracket
		if format[idxStart+1] == '[' {
			sb.appendString(format[pos : idxStart+2])
			pos = idxStart + 2
			continue
		}

		// Add skipped string (if any)
		if idxStart > pos { //idxStart > pos+1 ???
			sb.appendString(format[pos:idxStart])
		}

		// Find square bracket end
		relBlockClose := strings.IndexRune(format[idxStart:], ']')
		if relBlockClose == -1 {
			return "", fmt.Errorf("Opened square bracket never closed at pos %d.", idxStart)
		}
		idxEnd = idxStart + relBlockClose

		// found a block
		block := format[idxStart+1 : idxEnd]
		fields := strings.Fields(block)
		for _, field := range fields {

			switch field {
			// Options
			case "reset":
				sb.appendSgr(Reset)
			case "fg-reset":
				sb.appendSgr(ResetForegroundColor)
			case "bg-reset":
				sb.appendSgr(ResetBackgroundColor)
			case "bold":
				sb.appendSgr(Bold)
			case "boldOff":
				sb.appendSgr(BoldOff)
			case "underline":
				sb.appendSgr(Underline)
			case "underlineOff":
				sb.appendSgr(UnderlineOff)
			case "blink":
				sb.appendSgr(Blink)
			case "blinkOff":
				sb.appendSgr(BlinkOff)
			case "imageNegative":
				sb.appendSgr(ImageNegative)
			case "imagePositive":
				sb.appendSgr(ImagePositive)
			case "framed":
				sb.appendSgr(Framed)
			case "encircled":
				sb.appendSgr(Encircled)
			case "framedEncircledOff":
				sb.appendSgr(FramedEncircledOff)
			case "overlined":
				sb.appendSgr(Overlined)
			case "overlinedOff":
				sb.appendSgr(OverlinedOff)

			// Foreground Colors
			case "fg-black":
				sb.appendSgr(FgBlack)
			case "fg-red":
				sb.appendSgr(FgRed)
			case "fg-green":
				sb.appendSgr(FgGreen)
			case "fg-yellow":
				sb.appendSgr(FgYellow)
			case "fg-blue":
				sb.appendSgr(FgBlue)
			case "fg-magenta":
				sb.appendSgr(FgMagenta)
			case "fg-cyan":
				sb.appendSgr(FgCyan)
			case "fg-grey":
				sb.appendSgr(FgGrey)
			case "fg-white":
				sb.appendSgr(FgWhite)

			// Background Colors
			case "bg-black":
				sb.appendSgr(BgBlack)
			case "bg-red":
				sb.appendSgr(BgRed)
			case "bg-green":
				sb.appendSgr(BgGreen)
			case "bg-yellow":
				sb.appendSgr(BgYellow)
			case "bg-blue":
				sb.appendSgr(BgBlue)
			case "bg-magenta":
				sb.appendSgr(BgMagenta)
			case "bg-cyan":
				sb.appendSgr(BgCyan)
			case "bg-grey":
				sb.appendSgr(BgGrey)
			case "bg-white":
				sb.appendSgr(BgWhite)

			// Not Found
			default:
				isFgColor := strings.HasPrefix(field, "fg-")
				isBgColor := strings.HasPrefix(field, "bg-")
				if isFgColor || isBgColor {
					// Get the color number from the fg- or bg- code
					clr, err := strconv.Atoi(field[3:])
					if err != nil {
						goto invalidBlockCode
					}
					if clr < 0 || clr > 255 {
						return "", fmt.Errorf("Invalid color code %s. Expecting 0-255 or a defined color.", field[3:])
					}

					// Add actual fg or bg color to sgrBuilder
					if isFgColor {
						sb.appendSgr(FgColor(clr))
					} else {
						sb.appendSgr(BgColor(clr))
					}
					continue // next field
				} else {
					// not valid
					goto invalidBlockCode
				}
			}

			continue
		invalidBlockCode:
			return "", fmt.Errorf("Invalid block code '%s' in block at position %d.", field, idxStart)

		}

		// Change starting position for next iteration.
		pos = idxEnd + 1
	}

	if reset {
		sb.appendSgr(Reset)
	}

	if newline {
		sb.appendString("\n")
	}

	return sb.string(), nil
}
