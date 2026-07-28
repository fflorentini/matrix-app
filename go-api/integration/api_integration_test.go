package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestQRWorkflow(t *testing.T) {

	loginPayload := map[string]string{
		"username": "admin",
		"password": "admin123",
	}

	loginBody, _ := json.Marshal(loginPayload)

	loginResp, err := http.Post(
		"http://localhost:8080/login",
		"application/json",
		bytes.NewBuffer(loginBody),
	)

	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected 200, got %d",
			loginResp.StatusCode,
		)
	}

	var loginResult struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(
		loginResp.Body,
	).Decode(&loginResult); err != nil {

		t.Fatalf(
			"could not decode login response: %v",
			err,
		)
	}

	if loginResult.Token == "" {
		t.Fatal("expected JWT token")
	}

	qrPayload := map[string]any{
		"matrix": [][]float64{
			{1, 2},
			{3, 4},
			{5, 6},
		},
	}

	qrBody, _ := json.Marshal(qrPayload)

	req, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:8080/api/qr",
		bytes.NewBuffer(qrBody),
	)

	if err != nil {
		t.Fatalf("request creation failed: %v", err)
	}

	req.Header.Set(
		"Authorization",
		"Bearer "+loginResult.Token,
	)

	req.Header.Set(
		"Content-Type",
		"application/json",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		t.Fatalf(
			"qr request failed: %v",
			err,
		)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected 200, got %d",
			resp.StatusCode,
		)
	}

	var result map[string]any

	if err := json.NewDecoder(
		resp.Body,
	).Decode(&result); err != nil {

		t.Fatalf(
			"could not decode qr response: %v",
			err,
		)
	}

	if _, ok := result["q"]; !ok {
		t.Fatal("missing q matrix")
	}

	if _, ok := result["r"]; !ok {
		t.Fatal("missing r matrix")
	}

	if _, ok := result["statistics"]; !ok {
		t.Fatal("missing statistics")
	}
}