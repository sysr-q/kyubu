package cpe

import (
	"fmt"
	"github.com/sysr-q/kyubu/packets"
)

const SetTextHotKeySize = (packets.ByteSize + packets.StringSize + packets.StringSize + packets.IntSize + packets.ByteSize)

type SetTextHotKey struct {
	PacketId byte
	Label    string
	Action   string
	KeyCode  int32
	KeyMods  byte
}

func (p SetTextHotKey) Id() byte {
	return p.PacketId
}

func (p SetTextHotKey) Size() int {
	return SetTextHotKeySize
}

func (p SetTextHotKey) Bytes() []byte {
	return packets.ReflectBytes(p)
}

func (p SetTextHotKey) String() string {
	return "TextHotKey"
}

func ReadSetTextHotKey(b []byte) (packets.Packet, error) {
	var p SetTextHotKey
	err := packets.ReflectRead(b, &p)
	return &p, err
}

func NewSetTextHotKey(label, action string, keyCode int32, keyMods byte) (p *SetTextHotKey, err error) {
	if len(label) > packets.StringSize {
		return nil, fmt.Errorf("kyubu/cpe: cannot write over %d bytes in string", packets.StringSize)
	}
	if len(action) > packets.StringSize {
		return nil, fmt.Errorf("kyubu/cpe: cannot write over %d bytes in string", packets.StringSize)
	}
	p = &SetTextHotKey{
		PacketId: 0x15,
		Label:    label,
		Action:   action,
		KeyCode:  keyCode,
		KeyMods:  keyMods,
	}
	return
}

func init() {
	packets.MustRegister(&packets.PacketInfo{
		Id:   0x15,
		Read: ReadSetTextHotKey,
		Size: SetTextHotKeySize,
		Type: packets.ServerOnly,
		Name: "Set Text Hot Key (CPE)",
	})
}