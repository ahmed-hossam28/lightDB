package utils

import "encoding/binary"

func Uint32ToBytes(value uint32) []byte {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, value)
	return buffer
}

func BytesToUint32(buffer []byte) uint32 {
	return binary.BigEndian.Uint32(buffer)
}
