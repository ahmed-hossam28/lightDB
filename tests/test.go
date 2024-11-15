package tests

import "unsafe"

const (
	ColumnUsernameSize = 32
	ColumnEmailSize    = 255
)

type Raw struct {
	id       uint32
	username [ColumnUsernameSize]byte
	email    [ColumnEmailSize]byte
}

const (
	IdSize         = unsafe.Sizeof(Raw{}.id)
	UsernameSize   = unsafe.Sizeof(Raw{}.username)
	EmailSize      = unsafe.Sizeof(Raw{}.email)
	IdOffset       = 0
	UsernameOffset = IdOffset + IdSize
	EmailOffset    = UsernameOffset + UsernameSize
	RowSize        = IdSize + UsernameSize + EmailSize
)
