package sgr

import (
	"bytes"
	"errors"
	"fmt"
)

var errorInvalidPiece = errors.New("Invalid part. Expecting type `string` or `sgr.Sgr`.")

type sgrBuilder struct {
	buf bytes.Buffer
}

func (sb *sgrBuilder) append(part interface{}) error {
	// Just a simple piece of string to add to the result
	str, strOk := part.(string)
	if strOk {
		sb.appendString(str)
		return nil
	}

	// Add a SGR definition
	sgr, sgrOk := part.(Sgr)
	if sgrOk {
		sb.appendSgr(sgr)
		return nil
	}

	// Invalid type for this piece.
	return errorInvalidPiece
}

func (sb *sgrBuilder) appendString(str string) {
	fmt.Println("'" + str + "'")
	sb.buf.WriteString(str)
}

func (sb *sgrBuilder) appendSgr(sgr Sgr) {
	sb.buf.Write(CSI)
	sb.buf.Write(sgr.sequence())
	sb.buf.WriteByte(SgrEnd)
}

func (sb *sgrBuilder) string() string {
	return sb.buf.String()
}
