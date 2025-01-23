package types

import "time"

// ResourceConfig defines the resource requirements for deployment
type ResourceConfig struct {
    CPU     string `yaml:"cpu"`
    Memory  string `yaml:"memory"`
    Storage string `yaml:"storage"`
}

// DeployResponse represents the response from a deployment request
type DeployResponse struct {
    AppID  string `json:"appId"`
    Status string `json:"status"`
    URL    string `json:"url,omitempty"`
}

// StatusResponse represents the response from a status request
type StatusResponse struct {
    State         string    `json:"state"`
    LastDeployment time.Time `json:"lastDeployment"`
    Resources     struct {
        CPUUsage     float64 `json:"cpuUsage"`
        MemoryUsage  float64 `json:"memoryUsage"`
        StorageUsage float64 `json:"storageUsage"`
    } `json:"resources"`
}

// LogEntry represents a single log entry
type LogEntry struct {
    Timestamp time.Time `json:"timestamp"`
    Message   string    `json:"message"`
}

// LogsResponse represents the response from a logs request
type LogsResponse struct {
    Entries []LogEntry `json:"entries"`
}

// LogOptions represents options for log retrieval
type LogOptions struct {
    Follow bool      `json:"follow"`
    Tail   int       `json:"tail"`
    Since  time.Time `json:"since,omitempty"`
}
