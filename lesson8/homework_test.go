package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MultiError struct {
	errs []error
}

func (e *MultiError) Error() string {
	multiError := fmt.Sprintf("%d errors occured:\n", len(e.errs))
	for _, e := range e.errs {
		multiError += fmt.Sprintf("\t* %s", e.Error())
	}
	multiError += "\n"
	return multiError
}

func Append(err error, errs ...error) *MultiError {
	multiError, ok := err.(*MultiError)
	if !ok {
		return &MultiError{
			errs: errs,
		}
	}
	multiError.errs = append(multiError.errs, errs...)
	return multiError
}

func TestMultiError(t *testing.T) {
	var err error
	err = Append(err, errors.New("error 1"))
	err = Append(err, errors.New("error 2"))

	expectedMessage := "2 errors occured:\n\t* error 1\t* error 2\n"
	assert.EqualError(t, err, expectedMessage)
}
