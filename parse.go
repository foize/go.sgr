package sgr

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(format string) (string, error) {
	return parse(true, format)
}

func ParseWithoutReset(format string) (string, error) {
	return parse(false, format)
}

func parse(reset bool, format string) (string, error) {
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
		fmt.Println("block: ", block)
		fields := strings.Fields(block)
		for _, field := range fields {
			isFgColor := strings.HasPrefix(field, "fg-")
			isBgColor := strings.HasPrefix(field, "bg-")
			if isFgColor || isBgColor {
				// Get the color number from the fg- or bg- code
				clr, err := strconv.Atoi(field[3:])
				if err != nil {
					return "", fmt.Errorf("Invalid block code '%s' in block at position %d.", field, idxStart)
				}
				if clr < 0 || clr > 255 {
					return "", fmt.Errorf("Invalid color number %d. Expecting 0-255.", clr)
				}

				// Add actual fg or bg color to sgrBuilder
				if isFgColor {
					sb.appendSgr(FgColor(clr))
				} else {
					sb.appendSgr(BgColor(clr))
				}
				continue // next field
			}

			switch field {
			// Options
			case "reset":
				sb.appendSgr(Reset)
			case "resetForegroundColor":
				sb.appendSgr(ResetForegroundColor)
			case "resetBackgroundColor":
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
			case "fgBlack":
				sb.appendSgr(FgBlack)
			case "fgRed":
				sb.appendSgr(FgRed)
			case "fgGreen":
				sb.appendSgr(FgGreen)
			case "fgYellow":
				sb.appendSgr(FgYellow)
			case "fgBlue":
				sb.appendSgr(FgBlue)
			case "fgMagenta":
				sb.appendSgr(FgMagenta)
			case "fgCyan":
				sb.appendSgr(FgCyan)
			case "fgGrey":
				sb.appendSgr(FgGrey)
			case "fgWhite":
				sb.appendSgr(FgWhite)

			// Background Colors
			case "bgBlack":
				sb.appendSgr(BgBlack)
			case "bgRed":
				sb.appendSgr(BgRed)
			case "bgGreen":
				sb.appendSgr(BgGreen)
			case "bgYellow":
				sb.appendSgr(BgYellow)
			case "bgBlue":
				sb.appendSgr(BgBlue)
			case "bgMagenta":
				sb.appendSgr(BgMagenta)
			case "bgCyan":
				sb.appendSgr(BgCyan)
			case "bgGrey":
				sb.appendSgr(BgGrey)
			case "bgWhite":
				sb.appendSgr(BgWhite)

			// Not Found
			default:
				return "", fmt.Errorf("Invalid block code '%s' in block at position %d.", field, idxStart)
			}
		}

		// Change starting position for next iteration.
		pos = idxEnd + 1
	}

	if reset {
		sb.appendSgr(Reset)
	}

	return sb.string(), nil
}
