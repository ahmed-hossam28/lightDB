package executor

import (
	"fmt"
	"lightDB/internal/storage"
	"log"
	"os"
)

type MetaCommandResult string

const (
	MetaCommandSuccess      MetaCommandResult = "MetaSuccess"
	UnrecognizedMetaCommand MetaCommandResult = "UnrecognizedMetaCommand"
)

func PersistToDisk(table *storage.Table) {
	numberOfFullPages := table.RowsCount / storage.RowsPerPage

	for i := range numberOfFullPages {
		page := table.Pager.Page[i]
		if page != nil {
			err := table.Pager.Flush(i, storage.PageSize)
			if err != nil {
				fmt.Printf("Failed to flush pager for page %d: %v\n", i, err)
			}
		}
	}
	if numAdditionalRows := table.RowsCount % storage.RowsPerPage; numAdditionalRows > 0 {
		pageNum := numberOfFullPages
		if page := table.Pager.Page[pageNum]; page != nil {
			err := table.Pager.Flush(pageNum, numAdditionalRows*storage.RowSize)
			if err != nil {
				log.Fatalf("Failed to flush page %d: %v", pageNum, err)
			}
		}
	}

}
func DoMetaCommand(input string, table *storage.Table) MetaCommandResult {
	switch input {
	case ".exit":
		os.Exit(0)
		return MetaCommandSuccess
	case ".save":
		PersistToDisk(table)
		return MetaCommandSuccess
	case ".exitP":
		PersistToDisk(table)
		os.Exit(0)
		return MetaCommandSuccess
	default:
		return UnrecognizedMetaCommand
	}
}
