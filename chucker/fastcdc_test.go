package chunker

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"testing"

	"github.com/jotfs/fastcdc-go"
)

func TestFastCDC(t *testing.T) {

	data := make([]byte, 1000*1024*1024)
	rand.Seed(4542)
	rand.Read(data)
	rd := bytes.NewReader(data)

	chunker, err := NewFastCDCChunker(rd, fastcdc.Options{
		AverageSize: 64 * 1024,
	})
	if err != nil {
		log.Fatal(err)
	}
	m := make(map[string]int64)

	t.Logf("%-32s  %s\n", "CHECKSUM", "CHUNK SIZE")

	for {
		chunk, err := chunker.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		mds := md5.Sum(chunk.Data)
		key := fmt.Sprintf("%x", mds)
		if val, ok := m[key]; ok {
			m[key] = val + 1
		} else {
			m[key] = 1
		}
	}
	for key, count := range m {
		t.Logf("%s: %d\n", key, count)
	}
}
