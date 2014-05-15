package cpe

import (
	"fmt"
	"github.com/sysr-q/kyubu/packets"
)

const ExtInfoSize = (packets.ByteSize + packets.StringSize + packets.ShortSize)

type ExtInfo struct {
	PacketId       byte
	AppName        string
	ExtensionCount int16
}

func (p ExtInfo) Id() byte {
	return 0x10
}

func (p ExtInfo) Size() int {
	return ExtInfoSize
}

func (p ExtInfo) Bytes() []byte {
	return packets.ReflectBytes(p)
}

func (p ExtInfo) String() string {
	return "Negotiation"
}

func ReadExtInfo(b []byte) (packets.Packet, error) {
	var p ExtInfo
	err := packets.ReflectRead(b, &p)
	return &p, err
}

func NewExtInfo(appName string, extCount int16) (p *ExtInfo, err error) {
	if len(appName) > packets.StringSize {
		return nil, fmt.Errorf("kyubu/cpe: cannot write over %d bytes in string", packets.StringSize)
	}
	p = &ExtInfo{
		PacketId:       0x10,
		AppName:        appName,
		ExtensionCount: extCount,
	}
	return
}

func init() {
	packets.MustRegister(&packets.PacketInfo{
		Id:   0x10,
		Read: ReadExtInfo,
		Size: ExtInfoSize,
		Type: packets.Both,
		Name: "Ext Info (CPE)",
	})
}
