package storage

func Open(filename string) *Table {
	return NewTable(filename)
}
