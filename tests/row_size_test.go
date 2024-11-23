package tests

import (
	"fmt"
	"testing"
)

func returnSlice(a []byte) []byte {
	a[2] = 12
	return a[1:3]
}
func TestSomeFunctionalities(t *testing.T) {
	page := make([]byte, 100)

	page2 := page
	page2[0] = 10
	page2[1] = 22
	n := returnSlice(page)
	n[0] = 99
	fmt.Println(page[0], page[1], page[2])

}
