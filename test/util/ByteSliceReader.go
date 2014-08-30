package util

import (
	"io"
)

type byteSliceReader struct {
	data    []byte
	dataLen int
	pos     int
}

func ReadSlice(data []byte) io.Reader {
	result := &byteSliceReader{
		data:    data,
		dataLen: len(data)}

	return result
}

func (reader *byteSliceReader) Read(p []byte) (n int, err error) {
	available := reader.dataLen - reader.pos

	n = len(p)
	if n > available {
		n = available
	}

	if n > 0 {
		copy(p[0:n], reader.data[reader.pos:reader.pos+n])
	}

	reader.pos += n
	if reader.pos >= reader.dataLen {
		err = io.EOF
	}

	return
}
