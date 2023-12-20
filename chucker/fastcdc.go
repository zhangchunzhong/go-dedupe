package chunker

import (
	"io"

	"github.com/jotfs/fastcdc-go"
)

type FastCDCChunker = fastcdc.Chunker

// NewFastCDCChunker creates a new FastCDC chunker with the given reader and options.
// The reader is the source of data to be chunked.
// The options specify the configuration for the FastCDC algorithm.
func NewFastCDCChunker(rd io.Reader, opts fastcdc.Options) (*FastCDCChunker, error) {
	return fastcdc.NewChunker(rd, opts)
}

// Reset resets the FastCDC to its initial state.
func (c *FastCDC) Reset() error {
	// For a FastCDC, we might need to reset the underlying data reader.
	// However, the io.Reader interface doesn't provide a Reset method.
	// If your data reader is an *os.File or a *bytes.Buffer, you can cast it and reset it.
	// Otherwise, you might need to provide a way to recreate the data reader.
	return nil
}

// Close releases any resources associated with the FastCDC.
func (c *FastCDC) Close() error {
	// If the FastCDC's data reader also implements the io.Closer interface,
	// we can close it to release resources.
	if closer, ok := c.data.(io.Closer); ok {
		return closer.Close()
	}
	// Otherwise, there's nothing to close.
	return nil
}
