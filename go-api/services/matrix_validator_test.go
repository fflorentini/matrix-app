package services

import "testing"

func TestValidateMatrix_ValidMatrix(t *testing.T) {

	matrix := [][]float64{
		{1, 2},
		{3, 4},
	}

	err := ValidateMatrix(matrix)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestValidateMatrix_EmptyMatrix(t *testing.T) {

	matrix := [][]float64{}

	err := ValidateMatrix(matrix)

	if err != ErrEmptyMatrix {
		t.Fatalf(
			"expected %v, got %v",
			ErrEmptyMatrix,
			err,
		)
	}
}

func TestValidateMatrix_NonRectangular(t *testing.T) {

	matrix := [][]float64{
		{1, 2},
		{3, 4, 5},
	}

	err := ValidateMatrix(matrix)

	if err != ErrNonRectangularMatrix {
		t.Fatalf(
			"expected %v, got %v",
			ErrNonRectangularMatrix,
			err,
		)
	}
}

func TestValidateMatrix_EmptyRow(t *testing.T) {

	matrix := [][]float64{
		{},
		{},
	}

	err := ValidateMatrix(matrix)

	if err != ErrEmptyRow {
		t.Fatalf(
			"expected %v, got %v",
			ErrEmptyRow,
			err,
		)
	}
}