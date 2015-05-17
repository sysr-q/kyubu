package classic

import "github.com/sysr-q/kyubu/packets"

type PositionOrientation struct {
	PacketId   byte
	PlayerId   int8
	X, Y, Z    int16
	Yaw, Pitch byte
}

func (p PositionOrientation) Id() byte {
	return 0x08
}

func (p PositionOrientation) Size() int {
	return packets.ReflectSize(p)
}

func (p PositionOrientation) Bytes() []byte {
	return packets.ReflectBytes(p)
}

func ReadPositionOrientation(b []byte) (packets.Packet, error) {
	var p PositionOrientation
	err := packets.ReflectRead(b, &p)
	return &p, err
}

func NewPositionOrientation(playerId int8, x, y, z int16, yaw, pitch byte) (p *PositionOrientation, err error) {
	p = &PositionOrientation{
		PacketId: 0x08,
		PlayerId: playerId,
		X:        x,
		Y:        y,
		Z:        z,
		Yaw:      yaw,
		Pitch:    pitch,
	}
	return
}

func init() {
	packets.Register(&packets.PacketInfo{
		Id:        0x08,
		Read:      ReadPositionOrientation,
		Size:      packets.ReflectSize(&PositionOrientation{}),
		Direction: packets.Anomalous,
		Name:      "Position/Orientation",
	})
}