package models

type QRResultResponse struct {
	Q          [][]float64        `json:"q"`
	R          [][]float64        `json:"r"`
	Statistics StatisticsResponse `json:"statistics"`
}