package executor

import (
	"fmt"
	"lightDB/internal/storage"
	"strconv"
	"strings"
)

type StatementType int

const (
	StatementInsert StatementType = iota
	StatementSelect
	StatementDelete
)

type PrepareResult string

const (
	PrepareSuccess       PrepareResult = "PrepareSuccess"
	PrepareUnrecognized  PrepareResult = "PrepareUnrecognized"
	PrepareSyntaxError   PrepareResult = "PrepareSyntaxError"
	PrepareNegativeError PrepareResult = "PrepareNegativeError"
	PrepareStringTooLong PrepareResult = "PrepareStringTooLong"
)

type Statement struct {
	Type        StatementType
	RowToInsert storage.Row
}

func PrepareStatement(input string, statement *Statement) PrepareResult {
	if input[:min(6, len(input))] == "insert" {
		statement.Type = StatementInsert
		in := strings.Split(strings.TrimSpace(input), " ")
		if len(in) < 4 {
			return PrepareSyntaxError
		}

		res, err := strconv.Atoi(in[1])
		if err != nil {
			return PrepareSyntaxError
		}
		if res < 0 {
			return PrepareNegativeError
		}
		statement.RowToInsert.Id = uint32(res)

		if len(in[2]) > 32 || len(in[3]) > 32 {
			return PrepareStringTooLong
		}
		copy(statement.RowToInsert.Username[:], in[2])

		copy(statement.RowToInsert.Email[:], in[3])
		return PrepareSuccess
	}
	if input == "select" {
		statement.Type = StatementSelect
		return PrepareSuccess
	}
	return PrepareUnrecognized
}

type ExecuteResult string

const (
	ExecuteSuccess   ExecuteResult = "ExecuteSuccess"
	ExecuteTableFull ExecuteResult = "ExecuteTableFull"
	ExecuteError     ExecuteResult = "ExecuteError"
)

func executeInsert(statement *Statement, table *storage.Table) ExecuteResult {
	if table.RowsCount >= storage.TableMaxRows {
		return ExecuteError
	}

	row := &statement.RowToInsert
	cursor := storage.NewCursor(table)
	storage.SerializeRow(row, cursor.End().Value())
	table.RowsCount += 1
	return ExecuteSuccess
}
func executeSelect(statement *Statement, table *storage.Table) ExecuteResult {
	var row storage.Row
	cursor := storage.NewCursor(table).Start()
	printHeader()
	for !cursor.IsEnd() {
		row = storage.DeserializeRaw(cursor.Value())
		printRow(row)
		cursor.Next()
	}
	return ExecuteSuccess
}
func printHeader() {
	fmt.Printf("(ID, Username, Email)\n")
}
func printRow(row storage.Row) {
	fmt.Printf("(%d, %s, %s)\n", row.Id, string(row.Username[:]), string(row.Email[:]))
}
func ExecuteStatement(statement *Statement, table *storage.Table) ExecuteResult {
	switch statement.Type {
	case StatementInsert:
		return executeInsert(statement, table)
		break
	case StatementSelect:
		return executeSelect(statement, table)
		break
	}
	return ExecuteError
}
