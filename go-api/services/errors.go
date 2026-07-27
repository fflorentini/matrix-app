package services

import "errors"

var (
	ErrEmptyMatrix = errors.New("matrix cannot be empty")

	ErrEmptyRow = errors.New("matrix rows cannot be empty")

	ErrNonRectangularMatrix = errors.New(
		"matrix must be rectangular",
	)
)