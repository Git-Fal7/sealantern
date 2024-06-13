package readerwriter

import "io"

type ConnReadWrite struct {
	Rdr    io.Reader
	Wtr    io.Writer
	Buffer [16]byte
}
