package api

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

// httpClient is the interface for making HTTP requests
type httpClient interface {
    Do(req *http.Request) (*http.Response, error)
}

// client implements HTTP operations
type client struct {
    httpClient httpClient
    baseURL    string
    token      string
}

// newClient creates a new HTTP client
func newClient(baseURL, token string) *client {
    return &client{
        httpClient: &http.Client{
            Timeout: time.Second * 30,
        },
        baseURL: baseURL,
        token:   token,
    }
}

// get performs a GET request
func (c *client) get(endpoint string) ([]byte, error) {
    req, err := c.newRequest(http.MethodGet, endpoint, nil)
    if err != nil {
        return nil, err
    }
    return c.doRequest(req)
}

// post performs a POST request
func (c *client) post(endpoint string, payload interface{}) ([]byte, error) {
    req, err := c.newRequest(http.MethodPost, endpoint, payload)
    if err != nil {
        return nil, err
    }
    return c.doRequest(req)
}

// put performs a PUT request
func (c *client) put(endpoint string, payload interface{}) ([]byte, error) {
    req, err := c.newRequest(http.MethodPut, endpoint, payload)
    if err != nil {
        return nil, err
    }
    return c.doRequest(req)
}

// delete performs a DELETE request
func (c *client) delete(endpoint string) error {
    req, err := c.newRequest(http.MethodDelete, endpoint, nil)
    if err != nil {
        return err
    }
    _, err = c.doRequest(req)
    return err
}

// newRequest creates a new HTTP request
func (c *client) newRequest(method, endpoint string, payload interface{}) (*http.Request, error) {
    var body io.Reader
    if payload != nil {
        data, err := json.Marshal(payload)
        if err != nil {
            return nil, fmt.Errorf("failed to marshal payload: %w", err)
        }
        body = bytes.NewBuffer(data)
    }

    req, err := http.NewRequest(method, c.baseURL+endpoint, body)
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Authorization", "Bearer "+c.token)
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Accept", "application/json")

    return req, nil
}

// doRequest performs the HTTP request
func (c *client) doRequest(req *http.Request) ([]byte, error) {
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %w", err)
    }

    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(body))
    }

    return body, nil
}
