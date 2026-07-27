package models

type QRResponse struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}