package cartridge

import (
	"os"

	"github.com/pkg/errors"
)

type Cartridge struct {
	buffer []byte
	size   int64
}

func New(filename string) (*Cartridge, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		return nil, errors.Wrap(err, "error on open file")
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "error on fetch stats from file")
	}

	size := stats.Size()

	buffer := make([]byte, size)

	if _, err := file.Read(buffer); err != nil {
		return nil, errors.Wrap(err, "error on read data from file")
	}

	return &Cartridge{
		size:   size,
		buffer: buffer,
	}, nil
}

func (c *Cartridge) Size() int64 {
	return c.size
}

func (c *Cartridge) Buffer() []byte {
	return c.buffer
}
