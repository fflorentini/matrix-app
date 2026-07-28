package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"matrix-app/go-api/models"
	"net/http"
	"os"
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

	apiURL := os.Getenv("STATISTICS_API_URL")

	if apiURL == "" {
		apiURL = "http://localhost:3000/api/statistics"
	}

	resp, err := http.Post(
		apiURL,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Statistics API Status:", resp.StatusCode)
	fmt.Println("Statistics API Body:", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"statistics api returned status %d",
			resp.StatusCode,
		)
	}

	var result models.StatisticsResponse

	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}