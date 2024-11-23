package storage

const (
	RowsPerPage  = PageSize / RowSize          // Rows per page
	TableMaxRows = RowsPerPage * TableMaxPages // Maximum rows in the table
)

// Table struct definition
type Table struct {
	Pager     *Pager
	RowsCount uint32 // Number of rows in the table
}

func NewTable(filename string) *Table {
	table := &Table{
		Pager: PagerOpen("temp/" + filename),
	}
	table.RowsCount = uint32(table.Pager.FileLength / uint64(RowSize))
	return table
}

func (table *Table) RowSlot(rowNum uint32) []byte {
	pageNum := rowNum / RowsPerPage
	page := table.Pager.GetPage(pageNum)
	rowOffset := rowNum % RowsPerPage
	bytesOffset := rowOffset * RowSize

	return page[bytesOffset : bytesOffset+RowSize]
}
