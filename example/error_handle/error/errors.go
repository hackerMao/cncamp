package error

import (
	"errors"
	"fmt"
)

type PathError struct {
	Op   string
	Path string
	Err  error
}

// implemented Error interface
func (e *PathError) Error() string {
	return fmt.Sprintf("%s %s: %s\n", e.Op, e.Path, e.Err)
}

var (
	errNotFound error = errors.New("not found error")
)
