package director

import "fmt"

// ConnectionError is raised when the client cannot communicate to the host
type ConnectionError struct {
	Host           string
	TimeoutSeconds int64
	SourceError    error
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("Client could not connect to host '%s' after %d milliseconds. Error: %v",
		e.Host,
		e.TimeoutSeconds,
		e.SourceError,
	)
}

// NonLeaderCmdError is raised when the client sends a command to a non-leader node
type NonLeaderCmdError struct {
	leader string
}

func (e *NonLeaderCmdError) Error() string {
	return fmt.Sprintf("A non leader cannot propose commands. The leader is at %s", e.leader)
}

// InvalidIndexError is raised when an invalid log index is met
type InvalidIndexError struct {
	invalidIndex       int64
	lastestCommitIndex int64
}

func (e *InvalidIndexError) Error() string {
	return fmt.Sprintf("An index of %d is invalid; the latest commit index is %d", e.invalidIndex, e.lastestCommitIndex)
}

// LogParseError is raised when a log line cannot be parsed successfully
type LogParseError struct {
	invalidLine string
	sourceError error
}

func (e *LogParseError) Error() string {
	return fmt.Sprintf("Log could not be parsed successfully, the following line is of mal state: %s. Source: %v", e.invalidLine, e.sourceError)
}

// IOError is raised when the file could not be read/written
type IOError struct {
	filename    string
	sourceError error
}

func (e *IOError) Error() string {
	return fmt.Sprintf("File %s cannot be read/written successfull: %v", e.filename, e.sourceError)
}
