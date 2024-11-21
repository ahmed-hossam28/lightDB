package executor

import "os"

type MetaCommandResult string

const (
	MetaCommandSuccess      MetaCommandResult = "MetaSuccess"
	UnrecognizedMetaCommand MetaCommandResult = "UnrecognizedMetaCommand"
)

func DoMetaCommand(input string) MetaCommandResult {
	if input == ".exit" {
		os.Exit(0)
		return MetaCommandSuccess
	} else {
		return UnrecognizedMetaCommand
	}
}
