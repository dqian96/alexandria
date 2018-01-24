package archive

import "fmt"

// InvalidEntryError is raised when a client attempts to create an invalid entry in the archive
type InvalidEntryError struct {
	MaxKeyLength   int
	MaxValueLength int
	KeyLength      int
	ValueLength    int
}

func (e *InvalidEntryError) Error() string {
	return fmt.Sprintf(
		`Invalid input. The max key length should be %d chars and the max value length should be %d chars
		but the key length was %d chars and the value length was %d chars.`,
		e.MaxKeyLength,
		e.MaxValueLength,
		e.KeyLength,
		e.ValueLength,
	)
}
