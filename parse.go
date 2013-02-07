package sgr

func Parse(format string) string {
	return parse(true, format)
}

func ParseWithoutReset(format string) string {
	return parse(false, format)
}

func parse(reset bool, format string) string {
	// ++ do parsing
	// ++ actually split options and in the end have a buffer that we can write new text or commands to.

	// basic SGR's:
	// [reset]
	// [resetForegroundColor]
	// [resetBackgroundColor]
	// [bold]
	// [boldOff]
	// [underline]
	// [underlineOff]
	// [blink]
	// [blinkOff]
	// [imageNegative]
	// [imagePositive]
	// [framed]
	// [encircled]
	// [framedEncircledOff]
	// [overlined]
	// [overlinedOff]

	// color SGR examples:
	// [fg-blue] foreground blue
	// [fg-21]   foreground color 21
	// [bg-red] background red
	// [bg-22]  background color 22

	return "TODO colored string"
}
