package collyzstandard

import (
	"log"
	"time"

	"github.com/DataDog/zstd"
)

// DefaultCompressionLevel used, where fastest = 1, best = 20, default = 5.
const DefaultCompressionLevel = zstd.BestCompression

// CollyCache is the proposed interface for pluggable cache implementations in Colly.
type CollyCache interface {
	Init() error
	Get(url string) ([]byte, error)
	Put(url string, data []byte) error
	Remove(url string) error
}

var _ CollyCache = &Compressor{}

// Compressor wraps a CollyCache to provide data compression using the Zstandard algorithm.
type Compressor struct {
	Level   int
	Cache   CollyCache
	Logging bool
}

// NewCompressor creates wraps the given CollyCache with a Compressor,
// configured to use the default (maximum) compression level.
func NewCompressor(cache CollyCache) *Compressor {
	c := &Compressor{
		Level: DefaultCompressionLevel,
		Cache: cache,
	}
	return c
}

func (c *Compressor) Init() error {
	return c.Cache.Init()
}

func (c *Compressor) Get(url string) ([]byte, error) {
	b, err := c.Cache.Get(url)
	if err != nil {
		return nil, err
	}
	if len(b) == 0 {
		return nil, nil
	}
	start := time.Now()
	data, err := zstd.Decompress(nil, b)
	if err != nil {
		return nil, err
	}
	if c.Logging {
		dur := time.Since(start)
		log.Printf("decompressed in %s %s", dur, url)
	}
	return data, err
}

func (c *Compressor) Put(url string, data []byte) error {
	if len(data) == 0 {
		return c.Cache.Put(url, nil)
	}
	start := time.Now()
	b, err := zstd.CompressLevel(nil, data, c.Level)
	if err != nil {
		return err
	}
	if c.Logging {
		dur := time.Since(start)
		before := len(data)
		after := len(b)
		ratio := float64(before) / float64(after)
		log.Printf("compressed %d/%d, ratio %.2f, in %s %s", before, after, ratio, dur, url)
	}
	return c.Cache.Put(url, b)
}

func (c *Compressor) Remove(url string) error {
	return c.Cache.Remove(url)
}
