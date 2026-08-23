package model

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound    = errors.New("not found")
	ErrInterlock   = errors.New("interlock blocked")
	ErrInvalid     = errors.New("invalid argument")
	ErrConflict    = errors.New("state conflict")
	ErrWindow      = errors.New("outside process window")
	ErrTransition  = errors.New("illegal transition")
	ErrTension     = errors.New("tension out of band")
	ErrSpeed       = errors.New("speed not permitted")
)

type DomainError struct {
	Scope string
	Op    string
	Cause error
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("%s:%s: %v", e.Scope, e.Op, e.Cause)
}

func (e *DomainError) Unwrap() error { return e.Cause }

func Wrap(scope, op string, err error) error {
	if err == nil {
		return nil
	}
	return &DomainError{Scope: scope, Op: op, Cause: err}
}

func Is(err, target error) bool { return errors.Is(err, target) }