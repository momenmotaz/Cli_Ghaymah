package api

import (
    "encoding/json"
    "fmt"
    "net/url"
    "strconv"
    "time"
    "ghaymah-cli/pkg/config"
    "ghaymah-cli/pkg/types"
)

// GhaymahAPI handles communication with the Ghaymah Cloud API
type GhaymahAPI struct {
    client *client
}

// NewGhaymahAPI creates a new API client
func NewGhaymahAPI(baseURL, token string) *GhaymahAPI {
    return &GhaymahAPI{
        client: newClient(baseURL, token),
    }
}

// Deploy deploys an application to Ghaymah Cloud
func (api *GhaymahAPI) Deploy(config *config.Config) (*types.DeployResponse, error) {
    endpoint := "/apps"
    
    payload := map[string]interface{}{
        "name": config.AppName,
        "image": config.Image,
    }
    
    // Add optional fields if present
    if len(config.EnvVars) > 0 {
        payload["env"] = config.EnvVars
    }
    if config.Region != "" {
        payload["region"] = config.Region
    }
    if config.Resources != (types.ResourceConfig{}) {
        payload["resources"] = config.Resources
    }

    resp, err := api.client.post(endpoint, payload)
    if err != nil {
        return nil, fmt.Errorf("deployment failed: %w", err)
    }

    var deployResp types.DeployResponse
    if err := json.Unmarshal(resp, &deployResp); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }

    return &deployResp, nil
}

// GetStatus gets the status of an application
func (api *GhaymahAPI) GetStatus(appName string) (*types.StatusResponse, error) {
    endpoint := fmt.Sprintf("/apps/status?name=%s", url.QueryEscape(appName))
    
    resp, err := api.client.get(endpoint)
    if err != nil {
        return nil, fmt.Errorf("failed to get status: %w", err)
    }

    var statusResp types.StatusResponse
    if err := json.Unmarshal(resp, &statusResp); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }

    return &statusResp, nil
}

// GetLogs gets the logs of an application
func (api *GhaymahAPI) GetLogs(appName string, options *types.LogOptions) (*types.LogsResponse, error) {
    params := url.Values{}
    params.Add("name", appName)
    
    if options != nil {
        if options.Follow {
            params.Add("follow", "true")
        }
        if options.Tail > 0 {
            params.Add("tail", strconv.Itoa(options.Tail))
        }
        if !options.Since.IsZero() {
            params.Add("since", options.Since.Format(time.RFC3339))
        }
    }

    endpoint := fmt.Sprintf("/apps/logs?%s", params.Encode())
    
    resp, err := api.client.get(endpoint)
    if err != nil {
        return nil, fmt.Errorf("failed to get logs: %w", err)
    }

    var logsResp types.LogsResponse
    if err := json.Unmarshal(resp, &logsResp); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }

    return &logsResp, nil
}

// Helper methods for making HTTP requests
// TODO: Implement these methods
// func (api *GhaymahAPI) get(endpoint string) ([]byte, error)
// func (api *GhaymahAPI) post(endpoint string, payload interface{}) ([]byte, error)
// func (api *GhaymahAPI) put(endpoint string, payload interface{}) ([]byte, error)
// func (api *GhaymahAPI) delete(endpoint string) error
