package models

import "errors"

var (
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
	ErrIncorrectUserId = errors.New("incorrect user ID")
)
