package utils

import (
	"errors"
	"net/http"
)

type err struct {
	msg  string
	code uint
}

func (e err) Error() string {
	return e.msg
}

const (
	ErrInternal = iota
	ErrNotFound
	ErrBadRequest
)

func NewError(msg string, code uint) error {
	return &err{
		msg:  msg,
		code: code,
	}
}

func ParseErrorToHTTP(in error) (string, int) {
	var e *err
	if !errors.As(in, &e) {
		return "internal error", http.StatusInternalServerError
	}

	switch e.code {
	case ErrInternal:
		return "internal error", http.StatusInternalServerError
	case ErrNotFound:
		return e.msg, http.StatusNoContent
	case ErrBadRequest:
		return e.msg, http.StatusBadRequest
	default:
		return "internal error", http.StatusInternalServerError
	}
}
