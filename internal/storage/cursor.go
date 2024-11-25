package storage

type Cursor struct {
	Table      *Table
	RowNum     uint32
	EndOfTable bool
}

func NewCursor(table *Table) *Cursor {
	cursor := &Cursor{
		Table:      table,
		RowNum:     0,
		EndOfTable: 0 == table.RowsCount,
	}

	return cursor
}

func (cursor *Cursor) Start() *Cursor {
	cursor.RowNum = 0
	return cursor
}

func (cursor *Cursor) SetRowNum(row uint32) *Cursor {
	cursor.RowNum = row
	return cursor
}

func (cursor *Cursor) End() *Cursor {
	cursor.RowNum = cursor.Table.RowsCount

	return cursor
}

func (cursor *Cursor) IsEnd() bool {
	return cursor.RowNum == cursor.Table.RowsCount
}

func (cursor *Cursor) Next() bool {
	cursor.RowNum++
	if cursor.IsEnd() {
		cursor.EndOfTable = true
		return false
	}
	return true
}

func (cursor *Cursor) Value() []byte {
	rowNum := cursor.RowNum
	pageNum := rowNum / RowsPerPage
	page := cursor.Table.Pager.GetPage(pageNum)
	rowOffset := rowNum % RowsPerPage
	bytesOffset := rowOffset * RowSize

	return page[bytesOffset : bytesOffset+RowSize]
}
