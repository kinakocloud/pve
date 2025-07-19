package pve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// GET
func (m *Machine) getQuery(endpoint string, query map[string]any) ([]byte, error) {
	baseURL, err := url.Parse(m.API)
	if err != nil {
		return nil, fmt.Errorf("invalid API URL: %w", err)
	}

	fullURL, err := baseURL.Parse("/api2/json" + endpoint)
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

	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, resp.Status)
	}

	return body, nil
}

// Auto Retry
func (m *Machine) getQueryRetry(endpoint string, query map[string]any, retries int) ([]byte, error) {
	var err error
	var data []byte

	for range retries {
		data, err = m.getQuery(endpoint, query)
		if err == nil {
			return data, nil
		}
	}

	return nil, fmt.Errorf("failed after %d retries: %w", retries, err)
}

func (m *Machine) getQueryJSON(endpoint string, params map[string]any, result any) error {
	data, err := m.getQueryRetry(endpoint, params, 3)
	if err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	return json.Unmarshal(data, result)
}

// POST
func (m *Machine) httpPost(endpoint string, jsonData any) ([]byte, error) {
	baseURL, err := url.Parse(m.API)
	if err != nil {
		return nil, fmt.Errorf("invalid API URL: %w", err)
	}

	fullURL, err := baseURL.Parse("/api2/json" + endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid URI: %w", err)
	}

	var reqBody io.Reader
	if jsonData != nil {
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON: %w", err)
		}
		reqBody = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequest("POST", fullURL.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("PVEAPIToken=%s=%s", m.User, m.Token))

	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, resp.Status)
	}

	return body, nil
}

func (m *Machine) PostFormJSON(endpoint string, formData any, result any) error {
	data, err := m.httpPost(endpoint, formData)
	if err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	return json.Unmarshal(data, result)
}

// PUT
func (m *Machine) httpPut(endpoint string, jsonData any) ([]byte, error) {
	baseURL, err := url.Parse(m.API)
	if err != nil {
		return nil, fmt.Errorf("invalid API URL: %w", err)
	}

	fullURL, err := baseURL.Parse("/api2/json" + endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid URI: %w", err)
	}

	var reqBody io.Reader
	if jsonData != nil {
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal JSON: %w", err)
		}
		reqBody = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequest("PUT", fullURL.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("PVEAPIToken=%s=%s", m.User, m.Token))

	resp, err := m.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, resp.Status)
	}

	return body, nil
}

func (m *Machine) PutFormJSON(endpoint string, jsonData any, result any) error {
	data, err := m.httpPut(endpoint, jsonData)
	if err != nil {
		return err
	}

	if result == nil {
		return nil
	}

	return json.Unmarshal(data, result)
}

// DELETE
func (m *Machine) deleteQuery(endpoint string) error {
	baseURL, err := url.Parse(m.API)
	if err != nil {
		return fmt.Errorf("invalid API URL: %w", err)
	}

	fullURL, err := baseURL.Parse("/api2/json" + endpoint)
	if err != nil {
		return fmt.Errorf("invalid URI: %w", err)
	}

	req, err := http.NewRequest("DELETE", fullURL.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("PVEAPIToken=%s=%s", m.User, m.Token))

	resp, err := m.Client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("HTTP error %d", resp.StatusCode)
	}

	return nil
}
