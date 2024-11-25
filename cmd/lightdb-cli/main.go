package main

import (
	"bufio"
	"fmt"
	"lightDB/internal/executor"
	"lightDB/internal/storage"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	input  string
	reader = bufio.NewReader(os.Stdin)
)

func Open(filename string) *storage.Table {
	return storage.Open(filename)
}
func prompt() {
	fmt.Print("db > ")
}

func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") // Windows-specific clear
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear") // Unix-like system clear
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func TerminalCommand() bool {
	switch input {
	case "clear":
		clearScreen()
		return true
	case "ls":
		cmd := exec.Command("ls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		return true
	case "ls -l":
		cmd := exec.Command("ls", "-l")
		cmd.Stdout = os.Stdout
		cmd.Run()
		return true
	}
	return false
}
func main() {
	table := Open("test.db")
	for {
		prompt()
		input, _ = reader.ReadString('\n')
		input = input[:len(input)-1] //removing the \n char
		input = strings.ToLower(input)

		if len(input) == 0 {
			continue
		}
		if TerminalCommand() {
			continue
		}

		if input[0] == '.' {
			switch executor.DoMetaCommand(input, table) {
			case executor.MetaCommandSuccess:
				continue
			case executor.UnrecognizedMetaCommand:
				fmt.Printf("\033[31mUnrecognized command %s\n\033[0m", input)
				continue
			}
		}

		statement := &executor.Statement{}
		switch executor.PrepareStatement(input, statement) {
		case executor.PrepareSuccess:
			break
		case executor.PrepareSyntaxError:
			fmt.Printf("\033[31mSyntax error could not parse statement %s.\n\033[0m", input)
			continue
		case executor.PrepareUnrecognized:
			fmt.Printf("\033[31mUnrecognized keyword at start of %s.\n\033[0m", input)
			continue
		case executor.PrepareNegativeError:
			fmt.Printf("\033[31mError negative id not allowed.\n\033[0m")
			continue
		case executor.PrepareStringTooLong:
			fmt.Printf("\033[31mError string is too long.\n\033[0m")
			continue
		}
		switch executor.ExecuteStatement(statement, table) {
		case executor.ExecuteSuccess:
			fmt.Printf("\033[32mExecuted.\n\033[0m")
			break
		case executor.ExecuteTableFull:
			fmt.Printf("\033[31mError: Table full\n\033[0m")
			break
		}
	}
}
