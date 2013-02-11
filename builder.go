package sgr

import (
	"bytes"
	"errors"
)

var errorInvalidPiece = errors.New("Invalid part. Expecting type `string` or `sgr.Sgr`.")

type sgrBuilder struct {
	buf bytes.Buffer
}

func (sb *sgrBuilder) append(part interface{}) error {
	// Just a simple piece of string to add to the result
	str, strOk := part.(string)
	if strOk {
		sb.buf.WriteString(str)
		return nil
	}

	// Add a SGR definition
	sgr, sgrOk := part.(Sgr)
	if sgrOk {
		sb.buf.Write(CSI)
		sb.buf.Write(sgr.sequence())
		sb.buf.WriteByte(SgrEnd)
		return nil
	}

	// Invalid type for this piece.
	return errorInvalidPiece
}

func (sb *sgrBuilder) string() string {
	return sb.buf.String()
}
