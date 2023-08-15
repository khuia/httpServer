package main

import (
	"encoding/binary"
	"io"
	"net"
)

type Packet struct {
	nonce []byte
	data  []byte
}

func (p *Packet) lenNonce() uint32 {
	return uint32(len(p.nonce))
}
func (p *Packet) lenData() uint32 {
	return uint32(p.lenData())
}
func (p *Packet) lenNonceAndData() uint32 {
	return uint32(len(p.nonce) + len(p.data))
}
func (packet *Packet) pack() []byte {

	var length uint32
	length = packet.lenNonceAndData()

	result := make([]byte, 4+length)
	binary.LittleEndian.PutUint32(result, length)
	copy(result[4:4+packet.lenNonce()], packet.nonce)
	copy(result[4+packet.lenNonce():], packet.data)
	return result

}

func unpack(conn net.Conn) []byte {
	lengthByte := make([]byte, 4)
	io.ReadFull(conn, lengthByte)
	length := binary.LittleEndian.Uint32(lengthByte)
	buf := make([]byte, length)
	io.ReadFull(conn, buf)

	return buf

}
