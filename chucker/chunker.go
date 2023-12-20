package chunker

import (
	"github.com/jotfs/fastcdc-go"
)

type Chunk = fastcdc.Chunk

// Chunker is an interface for chunking algorithms.
type Chunker interface {
	Next() (*Chunk, error) // Next returns the next chunk and any error that might occur.
	Reset() error          // Reset resets the Chunker to its initial state.
	Close() error          // Close releases any resources associated with the Chunker.
}
