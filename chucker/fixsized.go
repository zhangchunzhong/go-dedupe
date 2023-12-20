package chunker

import "io"

// FixedSizeChunker implements the Chunker interface using fixed-size chunking.
type FixedSizeChunker struct {
	data      io.Reader
	chunkSize int
	buffer    []byte
}

// NewFixedSizeChunker creates a new instance of FixedSizeChunker.
func NewFixedSizeChunker(data io.Reader, chunkSize int) (*FixedSizeChunker, error) {
	return &FixedSizeChunker{
		data:      data,
		chunkSize: chunkSize,
		buffer:    make([]byte, chunkSize),
	}, nil
}

// Next generates the next chunk using fixed-size chunking.
func (c *FixedSizeChunker) Next() (*Chunk, error) {
	readBytes, err := io.ReadFull(c.data, c.buffer)
	if err == io.ErrUnexpectedEOF || err == io.EOF {
		// Last chunk may be smaller than chunkSize, handle it separately
		if readBytes > 0 {
			return &Chunk{
				Data:   c.buffer[:readBytes],
				Length: readBytes,
			}, nil
		}
		return nil, io.EOF
	} else if err != nil {
		return nil, err
	}

	return &Chunk{
		Data:   c.buffer,
		Length: c.chunkSize,
	}, nil
}

// Reset resets the FixedSizeChunker to its initial state.
func (c *FixedSizeChunker) Reset() error {
	// For a FixedSizeChunker, we might need to reset the underlying data reader.
	// However, the io.Reader interface doesn't provide a Reset method.
	// If your data reader is an *os.File or a *bytes.Buffer, you can cast it and reset it.
	// Otherwise, you might need to provide a way to recreate the data reader.
	return nil
}

// Close releases any resources associated with the FixedSizeChunker.
func (c *FixedSizeChunker) Close() error {
	// If the FixedSizeChunker's data reader also implements the io.Closer interface,
	// we can close it to release resources.
	if closer, ok := c.data.(io.Closer); ok {
		return closer.Close()
	}
	// Otherwise, there's nothing to close.
	return nil
}
