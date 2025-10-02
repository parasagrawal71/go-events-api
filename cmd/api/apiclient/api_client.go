package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// APIClient defines a reusable HTTP client
type APIClient struct {
	Client  *http.Client
	Headers map[string]string // optional headers like Authorization
}

// NewClient initializes a new API client
func NewClient(timeout time.Duration, headers map[string]string) *APIClient {
	return &APIClient{
		Client: &http.Client{
			Timeout: timeout,
		},
		Headers: headers,
	}
}

// GET request
func (api *APIClient) Get(url string, response interface{}) error {
	req, err := http.NewRequest(Methods.GET, url, nil)
	if err != nil {
		return err
	}

	// Add headers if any
	for key, value := range api.Headers {
		req.Header.Set(key, value)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decodeResponse(resp, response)
}

// POST request with JSON body
func (api *APIClient) Post(url string, payload interface{}, response interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(Methods.POST, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	for key, value := range api.Headers {
		req.Header.Set(key, value)
	}

	resp, err := api.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return decodeResponse(resp, response)
}

// Helper function to decode JSON response
func decodeResponse(resp *http.Response, out interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", body)
	}

	if out != nil {
		if err := json.Unmarshal(body, out); err != nil {
			return err
		}
	}

	// Print response
	jsonData, err := json.MarshalIndent(out, "", "  ") // pretty print with 2-space indent
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}
	fmt.Println("result: ", string(jsonData))

	return nil
}
