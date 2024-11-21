package tests

import (
	"fmt"
	"testing"
)

func TestRowSize(t *testing.T) {
	fmt.Println(IdSize, IdOffset)
	fmt.Println(UsernameSize, UsernameOffset)
	fmt.Println(EmailSize, EmailOffset)
	fmt.Println(RowSize)
}
