package services

import "testing"

func TestComputeQR_ReturnsMatrices(t *testing.T) {

	matrix := [][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	result, err := ComputeQR(matrix)

	if err != nil {
		t.Fatalf(
			"expected nil error, got %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected result")
	}

	// Q should be 3x3
	if len(result.Q) != 3 {
		t.Fatalf(
			"expected Q rows = 3, got %d",
			len(result.Q),
		)
	}

	if len(result.Q[0]) != 3 {
		t.Fatalf(
			"expected Q cols = 3, got %d",
			len(result.Q[0]),
		)
	}

	// R should be 3x2
	if len(result.R) != 3 {
		t.Fatalf(
			"expected R rows = 3, got %d",
			len(result.R),
		)
	}

	if len(result.R[0]) != 2 {
		t.Fatalf(
			"expected R cols = 2, got %d",
			len(result.R[0]),
		)
	}
}