package config

import (
	"fmt"
	"os"
)

const (
	// APIURLEnvVar is the environment variable name for the API URL
	APIURLEnvVar = "GHAYMAH_API_URL"
	// APITokenEnvVar is the environment variable name for the API token
	APITokenEnvVar = "GHAYMAH_API_TOKEN"
)

// GetAPIURL returns the API URL from environment variables
func GetAPIURL() (string, error) {
	url := os.Getenv(APIURLEnvVar)
	if url == "" {
		return "", fmt.Errorf("environment variable %s is not set", APIURLEnvVar)
	}
	return url, nil
}

// GetAPIToken returns the API token from environment variables
func GetAPIToken() (string, error) {
	token := os.Getenv(APITokenEnvVar)
	if token == "" {
		return "", fmt.Errorf("environment variable %s is not set", APITokenEnvVar)
	}
	return token, nil
}

// ValidateEnv validates that all required environment variables are set
func ValidateEnv() error {
	if _, err := GetAPIURL(); err != nil {
		return err
	}
	if _, err := GetAPIToken(); err != nil {
		return err
	}
	return nil
}

// GetAPIConfig returns the API configuration from environment variables
func GetAPIConfig() (string, string, error) {
	apiURL, err := GetAPIURL()
	if err != nil {
		return "", "", err
	}

	apiToken, err := GetAPIToken()
	if err != nil {
		return "", "", err
	}

	return apiURL, apiToken, nil
}
