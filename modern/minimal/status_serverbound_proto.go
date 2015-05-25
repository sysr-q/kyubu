// ARR, HERE BE DRAGONS! DO NOT EDIT
// protocol_generator -file=status_serverbound.go -direction=serverbound -state=status -package=minimal

package minimal

import (
	"encoding/binary"
	"github.com/kurafuto/kyubu/packets"
	"io"
)

func init() {
	packets.Register(packets.Status, packets.ServerBound, 0x00, func() packets.Packet { return &StatusRequest{} })
	packets.Register(packets.Status, packets.ServerBound, 0x01, func() packets.Packet { return &StatusPing{} })
}

func (t *StatusRequest) Id() byte {
	return 0x00 // 0
}

func (t *StatusRequest) Encode(ww io.Writer) (err error) {
	return
}

func (t *StatusRequest) Decode(rr io.Reader) (err error) {
	return
}

func (t *StatusPing) Id() byte {
	return 0x01 // 1
}

func (t *StatusPing) Encode(ww io.Writer) (err error) {
	// Encoding: Time (int64)
	if err = binary.Write(ww, binary.BigEndian, t.Time); err != nil {
		return
	}

	return
}

func (t *StatusPing) Decode(rr io.Reader) (err error) {
	// Decoding: Time (int64)
	if err = binary.Read(rr, binary.BigEndian, t.Time); err != nil {
		return
	}

	return
}
