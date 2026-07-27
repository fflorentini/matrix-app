package services

import (
	"matrix-app/go-api/models"

	"gonum.org/v1/gonum/mat"
)

func ComputeQR(matrix [][]float64) (*models.QRResponse, error) {

	dense := toDense(matrix)

	var qr mat.QR
	qr.Factorize(dense)

	var q mat.Dense
	var r mat.Dense

	qr.QTo(&q)
	qr.RTo(&r)

	return &models.QRResponse{
		Q: denseToSlice(&q),
		R: denseToSlice(&r),
	}, nil
}

func toDense(matrix [][]float64) *mat.Dense {

	rows := len(matrix)
	cols := len(matrix[0])

	data := make([]float64, 0, rows*cols)

	for _, row := range matrix {
		data = append(data, row...)
	}

	return mat.NewDense(rows, cols, data)
}

func denseToSlice(m mat.Matrix) [][]float64 {

	rows, cols := m.Dims()

	result := make([][]float64, rows)

	for i := 0; i < rows; i++ {

		result[i] = make([]float64, cols)

		for j := 0; j < cols; j++ {
			result[i][j] = m.At(i, j)
		}
	}

	return result
}