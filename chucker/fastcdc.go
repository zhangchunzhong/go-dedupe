package chunker

import (
	"io"

	"github.com/jotfs/fastcdc-go"
)

type FastCDCChunker = fastcdc.Chunker

func NewFastCDCChunker(rd io.Reader, opts fastcdc.Options) (*FastCDCChunker, error) {
	return fastcdc.NewChunker(rd, opts)
}
