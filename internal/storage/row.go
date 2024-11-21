package storage

import (
	"lightDB/utils"
	"unsafe"
)

const (
	ColumnUsernameSize = 32
	ColumnEmailSize    = 255
)

type Row struct {
	Id       uint32
	Username [ColumnUsernameSize]byte
	Email    [ColumnEmailSize]byte
}

const (
	IdSize         = uint32(unsafe.Sizeof(Row{}.Id))
	UsernameSize   = uint32(unsafe.Sizeof(Row{}.Username))
	EmailSize      = uint32(unsafe.Sizeof(Row{}.Email))
	IdOffset       = 0
	UsernameOffset = IdOffset + IdSize
	EmailOffset    = UsernameOffset + UsernameSize
	RowSize        = IdSize + UsernameSize + EmailSize
)

func SerializeRow(source *Row, dest []byte) {
	buffer := make([]byte, RowSize)
	copy(buffer[IdOffset:IdOffset+IdSize], utils.Uint32ToBytes(source.Id))
	copy(buffer[UsernameOffset:UsernameOffset+UsernameSize], source.Username[:])
	copy(buffer[EmailOffset:EmailOffset+EmailSize], source.Email[:])
	copy(dest, buffer)
}

func DeserializeRaw(buffer []byte) Row {
	var row Row
	row.Id = utils.BytesToUint32(buffer[IdOffset : IdOffset+IdSize])
	copy(row.Username[:], buffer[UsernameOffset:UsernameOffset+UsernameSize])
	copy(row.Email[:], buffer[EmailOffset:EmailOffset+EmailSize])

	return row
}
