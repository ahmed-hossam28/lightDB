package storage

const (
	PageSize      = 4096                        // Size of each page in bytes
	TableMaxPages = 100                         // Maximum number of pages
	RowsPerPage   = PageSize / RowSize          // Rows per page
	TableMaxRows  = RowsPerPage * TableMaxPages // Maximum rows in the table
)

// Table struct definition
type Table struct {
	RowsCount uint32                // Number of rows in the table
	Page      [TableMaxPages][]byte // Pages holding rows (each page is a byte slice)
}

func NewTable() *Table {
	return &Table{}
}
func (table *Table) RowSlot(rowNum uint32) []byte {
	pageNum := rowNum / RowsPerPage
	if table.Page[pageNum] == nil {
		table.Page[pageNum] = make([]byte, PageSize)
	}
	rowOffset := rowNum % RowsPerPage
	bytesOffset := rowOffset * RowSize

	return table.Page[pageNum][bytesOffset : bytesOffset+RowSize]
}
