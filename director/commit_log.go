package director

import (
	"bufio"
	"errors"
	fmt "fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Stores commands for the state machine; it is up to the state machine to serialize/deserialize commands
type commitLog interface {
	Append(int64, string) error
	Pop() error
	GetEntry(int64) (*logEntry, error)
	GetHighestIndex() int64
	IsEmpty() bool
	Close() error
}

type logEntry struct {
	index   int64
	term    int64
	command string
}

type logFile struct {
	newLineIndices []int64
	file           *os.File
	filename       string
	mutex          sync.RWMutex
}

func (c *logFile) Pop() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.newLineIndices) == 0 {
		return nil
	}

	fileInfo, err := c.file.Stat()
	if err != nil {
		log.Printf("Could not get info for file %s", c.filename)
		return err
	}

	var bytesToTruncate int64
	if len(c.newLineIndices) == 1 {
		bytesToTruncate = c.newLineIndices[0] + 1
	} else {
		bytesToTruncate = c.newLineIndices[len(c.newLineIndices)-1] - c.newLineIndices[len(c.newLineIndices)-2]
	}

	if os.Truncate(c.filename, int64(fileInfo.Size()-bytesToTruncate)) != nil {
		log.Printf("Cannot pop from file %s", c.filename)
		return err
	}
	c.newLineIndices = c.newLineIndices[:len(c.newLineIndices)-1]
	return nil
}

func (c *logFile) GetEntry(index int64) (*logEntry, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	ogIndex := index
	if index < 0 {
		index += c.GetHighestIndex() + 1
	}

	if index >= 0 && index <= c.GetHighestIndex() {
		lineStartIndex := int64(0)
		if index != 0 {
			lineStartIndex = c.newLineIndices[index-1] + 1
		}
		bytes := make([]byte, c.newLineIndices[index]-lineStartIndex) // discard trailing new line
		n, err := c.file.ReadAt(bytes, lineStartIndex)

		if err != nil || n != len(bytes) {
			log.Printf("Problem getting entry for file %s", c.filename)
			return nil, &IOError{filename: c.filename, sourceError: err}
		}
		return parseLogLine(string(bytes))
	}

	return nil, &InvalidIndexError{invalidIndex: ogIndex, lastestCommitIndex: c.GetHighestIndex()}
}

func (c *logFile) Append(term int64, command string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	line := fmt.Sprintf("%d|%d|%s\n", c.GetHighestIndex()+1, term, command)

	startIndex := int64(0)
	if len(c.newLineIndices) > 0 {
		startIndex = c.newLineIndices[len(c.newLineIndices)-1] + 1
	}
	n, err := c.file.WriteAt([]byte(line), startIndex)

	if err != nil || n != len(line) {
		return &IOError{filename: c.filename, sourceError: err}
	}

	c.newLineIndices = append(c.newLineIndices, startIndex-1+int64(len(line)))

	return nil
}

func (c *logFile) GetHighestIndex() int64 {
	return int64(len(c.newLineIndices) - 1)
}

func (c *logFile) IsEmpty() bool {
	return c.GetHighestIndex() == -1
}

func (c *logFile) Close() error {
	// TODO should this have a lock?
	return c.file.Close()
}

func newLogFile(filename string) (commitLog, error) {
	c := &logFile{
		filename:       filename,
		newLineIndices: make([]int64, 0),
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	c.file = file

	scanner := bufio.NewScanner(file)
	index := int64(-1) // start counting before index 0
	for scanner.Scan() {
		index += int64(len(scanner.Text()) + 1) // +1 for new line character
		c.newLineIndices = append(c.newLineIndices, index)
	}

	return c, nil
}

func parseLogLine(line string) (*logEntry, error) {
	// error parsing?
	tokens := strings.Split(line, "|")
	if len(tokens) != 3 {
		log.Printf("Incorrect number of tokens by delimited %s; %d expected but got %d", "|", 3, len(tokens))
		return nil, &LogParseError{invalidLine: line, sourceError: errors.New("Incorrect token length")}
	}
	index, err := strconv.ParseInt(tokens[0], 10, 64)
	if err != nil {
		log.Printf("Cannot parse int64 from %s", tokens[0])
		return nil, &LogParseError{invalidLine: line, sourceError: err}
	}
	term, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		log.Printf("Cannot parse int64 from %s", tokens[1])
		return nil, &LogParseError{invalidLine: line, sourceError: err}
	}
	return &logEntry{
		index:   index,
		term:    term,
		command: tokens[2],
	}, nil
}
