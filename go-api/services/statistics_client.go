package services

import (
	"bytes"
	"encoding/json"
	"matrix-app/go-api/models"
	"net/http"
)

func FetchStatistics(
	q [][]float64,
	r [][]float64,
) (*models.StatisticsResponse, error) {

	payload := models.StatisticsRequest{
		Q: q,
		R: r,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		"http://localhost:3000/api/statistics",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result models.StatisticsResponse

	err = json.NewDecoder(resp.Body).
		Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}