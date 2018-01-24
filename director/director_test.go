package director

import (
	"bufio"
	fmt "fmt"
	"io/ioutil"
	"os"
	"testing"
)

// Commit Log Tests

func TestGetEntriesFromExistingLog(t *testing.T) {
	// if this test passes then the non-empty log file was indexed properly
	// interaction with it thus reduces to the same as if the file were empty/didn't exist to start with
	// thus, specific append/pop tests for non-empty log files are not necessary
	// as the cases are synonmous to append/pop for empty files

	filename := "test123.txt"

	contents := []byte("0|1|COMMAND1\n1|2|COMMAND2\n")
	err := ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		t.Fatalf("Could not create/write to file")
	}
	defer os.Remove(filename)

	c := createCommitLog(filename, t)
	defer c.Close()

	checkLogEntryEquivalent(t, c, -2, &logEntry{
		index:   0,
		term:    1,
		command: "COMMAND1",
	})

	checkLogEntryEquivalent(t, c, -1, &logEntry{
		index:   1,
		term:    2,
		command: "COMMAND2",
	})
}

func TestAppends(t *testing.T) {
	filename := "test123.txt"
	_, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Could not create file")
	}
	defer os.Remove(filename)

	c := createCommitLog(filename, t)
	defer c.Close()

	appendAndCheck(t, c, 1, "COMMAND1")
	appendAndCheck(t, c, 2, "COMMAND2")

	checkLogEntryEquivalent(t, c, -1, &logEntry{
		index:   1,
		term:    2,
		command: "COMMAND2",
	})

	checkLogEntryEquivalent(t, c, -2, &logEntry{
		index:   0,
		term:    1,
		command: "COMMAND1",
	})
}

func TestPop(t *testing.T) {
	filename := "test123.txt"
	_, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Could not create file")
	}
	defer os.Remove(filename)

	c := createCommitLog(filename, t)
	defer c.Close()

	appendAndCheck(t, c, 1, "COMMAND")
	popAndCheck(t, c)

	_, err = c.GetEntry(0)
	if !c.IsEmpty() || err == nil {
		t.Fatalf("There should be no entry in the log - but there are")
	}
}

func TestAppendPopAppend(t *testing.T) {
	filename := "test123.txt"
	_, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Could not create file")
	}
	defer os.Remove(filename)

	c := createCommitLog(filename, t)
	defer c.Close()

	appendAndCheck(t, c, 1, "COMMAND")

	popAndCheck(t, c)

	appendAndCheck(t, c, 2, "COMMAND2")
	appendAndCheck(t, c, 3, "COMMAND3")

	popAndCheck(t, c)

	appendAndCheck(t, c, 4, "COMMAND4")

	checkLogEntryEquivalent(t, c, -1, &logEntry{
		index:   1,
		term:    4,
		command: "COMMAND4",
	})

	checkLogEntryEquivalent(t, c, 0, &logEntry{
		index:   0,
		term:    2,
		command: "COMMAND2",
	})
}

func createCommitLog(filename string, t *testing.T) commitLog {
	c, err := newLogFile(filename)
	if err != nil {
		t.Fatalf("Could not create commit log")
	}
	return c
}

func checkLogEntryEquivalent(t *testing.T, c commitLog, index int64, expectedEntry *logEntry) {
	logEntry, err := c.GetEntry(index)
	if err != nil {
		t.Fatalf("Could not get entry: %v", err)
	}

	if !isEquivalentEntry(expectedEntry, logEntry) {
		t.Fatalf("'%v' does not match the expected log result of %v", logEntry, expectedEntry)
	}
}

func appendAndCheck(t *testing.T, c commitLog, term int64, command string) {
	if err := c.Append(term, command); err != nil {
		t.Fatalf("Could not append: %v", err)
	}
}

func popAndCheck(t *testing.T, c commitLog) {
	if err := c.Pop(); err != nil {
		t.Fatalf("Could not pop: %v", err)
	}
}

func isEquivalentEntry(entry1 *logEntry, entry2 *logEntry) bool {
	return entry1.index == entry2.index && entry1.term == entry2.term && entry1.command == entry2.command
}

func printFile(filename string) {
	file, _ := os.Open(filename)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
