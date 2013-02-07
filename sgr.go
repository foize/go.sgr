package sgr

var CSI = []byte{'\x1b', '['} // ANSI Control Sequence Introducer
const SgrEnd = 'm'            // ANSI end char for SGR (Select Graphic Rendition)

type Sgr interface {
	code() uint8      // Return the ANSI code
	sequence() []byte // Returns the sequence
}
