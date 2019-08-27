package db

import (
	"fmt"
)

type dbError struct {
	name string
	err  error
}

func (e *dbError) Error() string {
	return fmt.Sprintf("db [%s] error [%s]", e.name, e.err.Error())
}
func NewDBError(name string, err error) error {
	e := &dbError{
		name: name,
		err:  err,
	}
	return e
}
