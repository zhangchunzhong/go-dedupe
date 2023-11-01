package chunker

import (
	"github.com/jotfs/fastcdc-go"
)

type Chunk = fastcdc.Chunk

// Chunker is an interface for chunking algorithms.
type Chunker interface {
	Next() (*Chunk, error)
}
