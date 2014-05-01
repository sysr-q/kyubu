package packets

import (
	"bytes"
	"strings"
	"testing"
)

func TestWriteByte(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	pw := PacketWrapper{Buffer: b, Strip: true}
	if err := pw.WriteByte(0x00); err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if !bytes.Equal(b.Bytes(), []byte{0x00}) {
		t.Fatalf("expected {0x00}, got %#v", b.Bytes())
	}
}

func TestWriteSByte(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	pw := PacketWrapper{Buffer: b, Strip: true}
	if err := pw.WriteSByte(0); err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if !bytes.Equal(b.Bytes(), []byte{0x00}) {
		t.Fatalf("expected {0x00}, got %#v", b.Bytes())
	}
}

func TestWriteShort(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	pw := PacketWrapper{Buffer: b, Strip: true}
	if err := pw.WriteShort(0); err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if !bytes.Equal(b.Bytes(), []byte{0x00, 0x00}) {
		t.Fatalf("expected {0x00, 0x00}, got %#v", b.Bytes())
	}
}

func TestWriteString(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	pw := PacketWrapper{Buffer: b, Strip: true}
	if err := pw.WriteString("test"); err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if len(b.Bytes()) != 64 {
		t.Fatalf("expected 64 bytes, got %d", len(b.Bytes()))
	}
	if !bytes.Equal(b.Bytes()[:4], []byte{0x74, 0x65, 0x73, 0x74}) {
		t.Fatalf("expected b.Bytes()[:4] == {0x74, 0x65, 0x73, 0x74}, got %#v", b.Bytes()[:4])
	}
	if err := pw.WriteString(strings.Repeat("t", 65)); err == nil {
		t.Fatal("expected err, got nil")
	}
}

func TestWriteBytes(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	pw := PacketWrapper{Buffer: b, Strip: true}
	if err := pw.WriteBytes([]byte("test")); err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if len(b.Bytes()) != 1024 {
		t.Fatalf("expected 1024 bytes, got %d", len(b.Bytes()))
	}
	if !bytes.Equal(b.Bytes()[:4], []byte{0x74, 0x65, 0x73, 0x74}) {
		t.Fatalf("expected b.Bytes()[:4] == {0x74, 0x65, 0x73, 0x74}, got %#v", b.Bytes()[:4])
	}
	if err := pw.WriteBytes(bytes.Repeat([]byte{0x00}, 1025)); err == nil {
		t.Fatal("expected err, got nil")
	}
}

func TestReadByte(t *testing.T) {
	b := bytes.NewBuffer([]byte{0x00})
	pw := PacketWrapper{Buffer: b, Strip: true}
	i, err := pw.ReadByte()
	if err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if i != 0x00 {
		t.Fatalf("expected 0x00, got %#.2x", i)
	}
}

func TestReadSByte(t *testing.T) {
	b := bytes.NewBuffer([]byte{0x00})
	pw := PacketWrapper{Buffer: b, Strip: true}
	i, err := pw.ReadSByte()
	if err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if i != 0x00 {
		t.Fatalf("expected 0x00, got %#.2x", i)
	}
}

func TestReadShort(t *testing.T) {
	b := bytes.NewBuffer([]byte{0x00, 0x00})
	pw := PacketWrapper{Buffer: b, Strip: true}
	i, err := pw.ReadShort()
	if err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if i != int16(0) {
		t.Fatalf("expected int16(0), got %v", i)
	}
}

func TestReadString(t *testing.T) {
	b := bytes.NewBuffer([]byte(strings.Repeat("a", 64)))
	pw := PacketWrapper{Buffer: b, Strip: true}
	s, err := pw.ReadString()
	if err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if s != strings.Repeat("a", 64) {
		t.Fatalf(`expected "a"*64, got %v`, s)
	}
}

func TestReadBytes(t *testing.T) {
	b := bytes.NewBuffer(bytes.Repeat([]byte{0x01}, 1024))
	pw := PacketWrapper{Buffer: b, Strip: true}
	by, err := pw.ReadBytes()
	if err != nil {
		t.Fatal("expected nil err, got", err)
	}
	if !bytes.Equal(by, bytes.Repeat([]byte{0x01}, 1024)) {
		t.Fatalf("expected {0x01}*1024, got %v", by)
	}
}

func TestPacketTypeString(t *testing.T) {
	serverOnly := ServerOnly.String()
	if serverOnly != "S->C" {
		t.Fatalf(`expected "S->C", got %q`, serverOnly)
	}
	clientOnly := ClientOnly.String()
	if clientOnly != "C->S" {
		t.Fatalf(`expected "C->S", got %q`, clientOnly)
	}
	both := Both.String()
	if both != "C<>S" {
		t.Fatalf(`expected "C<>S", got %q`, both)
	}
	x := PacketType(5).String()
	if x != "????" {
		t.Fatalf(`expected "????", got %q`, x)
	}
}

func TestRegisterPacket(t *testing.T) {
	if ok, err := Register(&PacketInfo{Id: 0x00}); err == nil {
		t.Fatal("round 1: expected an error, got nil and ok:", ok)
	}
}