package main

import (
	"encoding/binary"
	"unsafe"
)

const (
	ColumnUsernameSize = 32
	ColumnEmailSize    = 255
)

type Row struct {
	id       uint32
	username [ColumnUsernameSize]byte
	email    [ColumnEmailSize]byte
}

const (
	IdSize         = uint32(unsafe.Sizeof(Row{}.id))
	UsernameSize   = uint32(unsafe.Sizeof(Row{}.username))
	EmailSize      = uint32(unsafe.Sizeof(Row{}.email))
	IdOffset       = 0
	UsernameOffset = IdOffset + IdSize
	EmailOffset    = UsernameOffset + UsernameSize
	RowSize        = IdSize + UsernameSize + EmailSize
)

func SerializeRow(source *Row, dest []byte) {
	buffer := make([]byte, RowSize)
	copy(buffer[IdOffset:IdOffset+IdSize], uint32ToBytes(source.id))
	copy(buffer[UsernameOffset:UsernameOffset+UsernameSize], source.username[:])
	copy(buffer[EmailOffset:EmailOffset+EmailSize], source.email[:])
	copy(dest, buffer)
}

func DeserializeRaw(buffer []byte) Row {
	var row Row
	row.id = bytesToUint32(buffer[IdOffset : IdOffset+IdSize])
	copy(row.username[:], buffer[UsernameOffset:UsernameOffset+UsernameSize])
	copy(row.email[:], buffer[EmailOffset:EmailOffset+EmailSize])

	return row
}

func uint32ToBytes(value uint32) []byte {
	buffer := make([]byte, 4)
	binary.BigEndian.PutUint32(buffer, value)
	return buffer
}

func bytesToUint32(buffer []byte) uint32 {
	return binary.BigEndian.Uint32(buffer)
}
