package resettablereader

import (
	"bytes"
	"io"
)

// Reader represents an io.Reader which can reset its own content.
type Reader struct {
	io.Reader

	base io.Reader
	buf  *bytes.Buffer
}

// New returns a new *Reader.
func New(r io.Reader) *Reader {
	var buf bytes.Buffer
	return &Reader{
		Reader: io.TeeReader(r, &buf),
		base:   r,
		buf:    &buf,
	}
}

// Reset resets its own content.
func (r *Reader) Reset() {
	r.Reader = io.MultiReader(r.buf, io.TeeReader(r.base, r.buf))
}
