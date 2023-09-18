package utils

import (
	"net/http"
	"time"
)

// CheckEndpoint checks if an endpoint is available.
func CheckOTELConnectorEndpointHealth(endpoint string, timeout time.Duration) bool {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(endpoint)
	if err != nil {
		return false
	}
	return resp.StatusCode == http.StatusOK
}
