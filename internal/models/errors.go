package models

import "errors"

var (
	ErrAlreadyExists            = errors.New("already exists")
	ErrNotFound                 = errors.New("not found")
	ErrIncorrectFilter          = errors.New("incorrect filter")
	ErrIncorrectPaging          = errors.New("incorrect paging format")
	ErrIncorrectPostInput       = errors.New("incorrect post input")
	ErrIncorrectCommentInput    = errors.New("incorrect comment input")
	ErrIncorrectSubCommentInput = errors.New("incorrect subcomment input")
	ErrIncorrectIdFormat        = errors.New("incorrect id format")
)
