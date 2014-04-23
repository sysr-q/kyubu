// Package chunk contains all the necessary logic to parse compressed and
// uncompressed chunk data into sane representations.
package chunk

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"fmt"
)

type Chunk struct {
	Length int
	Data []byte
	X, Y, Z int16
}

func (ch *Chunk) TileAt(x, y, z int16) byte {
	if (x < 0 || y < 0 || z < 0 ||
		x >= ch.X || y >= ch.Y || z >= ch.Z) {
		return 0
	}
	b := ch.Data[(y * ch.Y + z) * ch.X + x] & 255
	return b
}

func (ch *Chunk) SetTile(x, y, z int16, blockType byte) {
	if (x < 0 || y < 0 || z < 0 ||
		x >= ch.X || y >= ch.Y || z >= ch.Z) {
		return
	}
	ch.Data[(y * ch.Y + z) * ch.X + x] = blockType
}

func (ch *Chunk) Compress() ([]byte, error) {
	b := bytes.NewBuffer([]byte{})
	gz := gzip.NewWriter(b)
	if err := binary.Write(gz, binary.BigEndian, &ch.Length); err != nil {
		return nil, err
	}
	if _, err := gz.Write(ch.Data); err != nil {
		return nil, err
	}
	gz.Flush()
	return b.Bytes(), nil
}

func Decompress(b []byte, x, y, z int16) (*Chunk, error) {
	buf := bytes.NewReader(b)
	gz, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	var length int32
	lengthBytes := make([]byte, 4)
	if _, err := gz.Read(lengthBytes); err != nil {
		return nil, err
	}
	buf = bytes.NewReader(lengthBytes)
	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return nil, err
	}

	// I don't know /why/ we need to cast back to int first, but if we don't,
	// we'll get x*y*z == 0. Yeah.
	if length != int32(int(x)*int(y)*int(z)) {
		return nil, fmt.Errorf("kyubu/chunk: decoded length %d doesn't match given %d*%d*%d = %d", length, x, y, z, int32(int(x)*int(y)*int(z)))
	}
	
	chunk := &Chunk{
		Length: int(length),
		Data: make([]byte, length),
		X: x,
		Y: y,
		Z: z,
	}
	if _, err = gz.Read(chunk.Data); err != nil {
		return nil, err
	}
	return chunk, nil
}