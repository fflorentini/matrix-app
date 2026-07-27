package services

func ValidateMatrix(matrix [][]float64) error {

	if len(matrix) == 0 {
		return ErrEmptyMatrix
	}

	expectedColumns := len(matrix[0])

	if expectedColumns == 0 {
		return ErrEmptyRow
	}

	for _, row := range matrix {

		if len(row) == 0 {
			return ErrEmptyRow
		}

		if len(row) != expectedColumns {
			return ErrNonRectangularMatrix
		}
	}

	return nil
}