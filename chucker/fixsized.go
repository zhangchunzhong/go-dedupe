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
