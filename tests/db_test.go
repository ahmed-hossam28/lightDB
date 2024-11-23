package tests

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func setupTestEnvironment() {
	os.MkdirAll("temp", os.ModePerm)
}

func cleanupTestEnvironment() {
	os.RemoveAll("temp")
}
func TestMain(m *testing.M) {
	setupTestEnvironment()
	code := m.Run()
	cleanupTestEnvironment()
	os.Exit(code)
}

func runScript(commands []string) []string {
	cmd := exec.Command("../bin/lightdb")
	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	_ = cmd.Start()

	for _, command := range commands {
		_, _ = stdin.Write([]byte(command + "\n"))
	}
	stdin.Close()

	outputBytes, _ := io.ReadAll(stdout)
	cmd.Wait()

	cleanOutput := removeNonPrintable(string(outputBytes))

	return strings.Split(cleanOutput, "\n")
}

func removeNonPrintable(input string) string {
	// Remove only non-printable characters except newlines and the prompt control sequences
	return strings.Map(func(r rune) rune {
		if r == '\n' || r == '\r' {
			return r // Keep newlines and carriage returns
		}
		if r >= 32 && r <= 126 { // ASCII printable range
			return r
		}
		return -1 // Remove non-printable characters
	}, input)
}

func TestInsertAndSelect(t *testing.T) {
	output := runScript([]string{
		"insert 1 user1 user1@example.com",
		"select",
		".exit",
	})

	expected := []string{
		"db > [32mExecuted.",
		"[0mdb > (1, user1, user1@example.com)",
		"[32mExecuted.",
		"[0mdb > ",
	}

	assert.Equal(t, expected, output)
}
func TestInsertBoundUserNameLengthAndSelect(t *testing.T) {
	output := runScript([]string{
		"insert 1 " + strings.Repeat("a", 32) + " user1@example.com",
		"select",
		".exit",
	})

	expected := []string{
		"db > [32mExecuted.",
		"[0mdb > (1, " + strings.Repeat("a", 32) + ", user1@example.com)",
		"[32mExecuted.",
		"[0mdb > ",
	}

	assert.Equal(t, expected, output)
}

func TestInsertOutOfBoundUserNameLengthAndSelect(t *testing.T) {
	output := runScript([]string{
		"insert 1 " + strings.Repeat("a", 33) + " user1@example.com",
	})
	expected := []string{
		"db > [31mError string is too long.",
		"[0mdb > ",
	}
	assert.Equal(t, expected, output)
}

func TestSyntaxError(t *testing.T) {
	output := runScript([]string{
		"insert 1",
	})
	expected := []string{
		"db > [31mSyntax error could not parse statement insert 1.",
		"[0mdb > ",
	}

	assert.Equal(t, expected, output)
}

func TestNegativeID(t *testing.T) {
	output := runScript([]string{
		"insert -1 user1 user1@example.com",
	})
	expected := []string{
		"db > [31mError negative id not allowed.",
		"[0mdb > ",
	}
	assert.Equal(t, expected, output)
}

func TestStringTooLong(t *testing.T) {
	output := runScript([]string{
		"insert 1 " + strings.Repeat("ab", 35) + " user1@example.com",
	})
	expected := []string{
		"db > [31mError string is too long.",
		"[0mdb > ",
	}
	assert.Equal(t, expected, output)
}

func TestPersistenceToDisk(t *testing.T) {
	runScript([]string{
		"insert 1 ahmad user1@example.com",
		".exitP",
	})
	output2 := runScript([]string{
		"select",
	})
	expected2 := []string{
		"db > (1, ahmad, user1@example.com)",
		"[32mExecuted.",
		"[0mdb > ",
	}
	assert.Equal(t, expected2, output2)
}
