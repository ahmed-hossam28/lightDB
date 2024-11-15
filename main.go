package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	input  string
	reader = bufio.NewReader(os.Stdin)
)

func prompt() {
	fmt.Print("lightdb > ")
}

func main() {
	table := NewTable()
	for {
		prompt()
		input, _ = reader.ReadString('\n')
		input = input[:len(input)-1] //removing the \n char
		if input[0] == '.' {
			switch doMetaCommand(input) {
			case MetaCommandSuccess:
				continue
			case UnrecognizedMetaCommand:
				fmt.Printf("\033[31mUnrecognized command %s\n\033[0m", input)
				continue
			}
		}

		statement := &Statement{}
		switch prepareStatement(input, statement) {
		case PrepareSuccess:
			break
		case PrepareSyntaxError:
			fmt.Printf("\033[31mSyntax error could not parse statement %s\n\033[0m", input)
			continue
		case PrepareUnrecognized:
			fmt.Printf("\033[31mUnrecognized keyword at start of %s\n\033[0m", input)
			continue
		}
		switch executeStatement(statement, table) {
		case ExecuteSuccess:
			fmt.Printf("\033[32mExecuted\n\033[0m")
			break
		case ExecuteTableFull:
			fmt.Printf("\033[31mError: Table full\n\033[0m")
			break
		}
	}
}
