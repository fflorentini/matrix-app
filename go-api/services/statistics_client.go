package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"matrix-app/go-api/models"
	"net/http"
	"os"
	"time"
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

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	var resp *http.Response
	var lastErr error

	for attempt := 1; attempt <= 3; attempt++ {

		resp, lastErr = client.Post(
			apiURL,
			"application/json",
			bytes.NewBuffer(body),
		)

		if lastErr == nil && resp.StatusCode == http.StatusOK {
			break
		}

		if resp != nil {
			resp.Body.Close()
		}

		if attempt < 3 {
			time.Sleep(2 * time.Second)
		}
	}

	if lastErr != nil {
		return nil, fmt.Errorf(
			"statistics api unavailable after 3 attempts: %w",
			lastErr,
		)
	}

	if resp == nil {
		return nil, fmt.Errorf(
			"statistics api unavailable after 3 attempts",
		)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"statistics api returned status %d",
			resp.StatusCode,
		)
	}

	var result models.StatisticsResponse

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}