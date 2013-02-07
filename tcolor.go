package tcolor

var CSI = []byte{'\x1b', '['} // ANSI Control Sequence Introducer
const SgrEnd = 'm'            // ANSI end char for SGR (Select Graphic Rendition)

type Sgr interface {
	SgrCode() uint8   // Return the ANSI code
	Sequence() []byte // Returns the sequence
}
