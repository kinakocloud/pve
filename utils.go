package pve

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GET
func (m *Machine) getQuery(uri string, query map[string]any) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	baseURL, err := url.Parse(m.API)
	if err != nil {
		return nil, fmt.Errorf("invalid API URL: %w", err)
	}

	fullURL, err := baseURL.Parse("/api2/json" + uri)
	if err != nil {
		return nil, fmt.Errorf("invalid URI: %w", err)
	}

	if query != nil {
		values := fullURL.Query()
		for key, value := range query {
			values.Add(key, fmt.Sprintf("%v", value))
		}
		fullURL.RawQuery = values.Encode()
	}

	req, err := http.NewRequest("GET", fullURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("PVEAPIToken=%s=%s", m.User, m.Token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (m *Machine) getQueryJSON(endpoint string, params map[string]any, result any) error {
	data, err := m.getQuery(endpoint, params)
	if err != nil {
		return err
	}

	if result != nil {
		return json.Unmarshal(data, result)
	}

	return nil
}

// POST
func (m *Machine) httpPost(endpoint string, jsonData any) ([]byte, error) {
	fullURL := m.API + endpoint

	var reqBody io.Reader
	if jsonData != nil {
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON: %w", err)
		}
		reqBody = strings.NewReader(string(jsonBytes))
	}

	req, err := http.NewRequest("POST", fullURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("PVEAPIToken=%s@%s!%s=%s",
		m.User, "pam", m.Token, "your-token-secret"))

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

func (m *Machine) PostFormJSON(endpoint string, formData any, result any) error {
	data, err := m.httpPost(endpoint, formData)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, result)
}

// DELETE
func (m *Machine) deleteQuery(endpoint string) error {
	fullURL := m.API + endpoint

	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("PVEAPIToken=%s@%s!%s=%s",
		m.User, "pam", m.Token, "your-token-secret"))

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP error %d", resp.StatusCode)
	}

	return nil
}
