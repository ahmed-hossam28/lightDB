package main

import (
	"fmt"
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
	PrepareSuccess      PrepareResult = "PrepareSuccess"
	PrepareUnrecognized PrepareResult = "PrepareUnrecognized"
	PrepareSyntaxError  PrepareResult = "PrepareSyntaxError"
)

type Statement struct {
	Type        StatementType
	RowToInsert Row
}

func prepareStatement(input string, statement *Statement) PrepareResult {
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
		statement.RowToInsert.id = uint32(res)
		//statement.RowToInsert.username = [32]byte([]byte(in[2]))
		copy(statement.RowToInsert.username[:], in[2])
		//statement.RowToInsert.email = [255]byte([]byte(in[3]))
		copy(statement.RowToInsert.email[:], in[3])
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

func executeInsert(statement *Statement, table *Table) ExecuteResult {
	if table.RowsCount >= TableMaxRows {
		return ExecuteError
	}

	row := &statement.RowToInsert
	SerializeRow(row, table.RowSlot(table.RowsCount))
	table.RowsCount += 1
	return ExecuteSuccess
}
func executeSelect(statement *Statement, table *Table) ExecuteResult {
	var row Row
	for i := uint32(0); i < table.RowsCount; i++ {
		row = DeserializeRaw(table.RowSlot(i))
		printRow(row)
	}
	return ExecuteSuccess
}
func printRow(row Row) {
	fmt.Println(row.id, string(row.username[:]), string(row.email[:]))
}
func executeStatement(statement *Statement, table *Table) ExecuteResult {
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
