package hostinfo

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetInactiveHostnames(t *testing.T) {

	os.Setenv("THRESHOLD", "")
	threshold := 1
	expectedHostnames := []string{"mta-prod-1", "mta-prod-3"}

	result := getInactiveHostnames(threshold)

	// Check the length of the result
	if len(result) != len(expectedHostnames) {
		t.Errorf("Expected %d inactive hostnames, but got %d", len(expectedHostnames), len(result))
	}

	// Check each individual hostname
	for _, hostname := range expectedHostnames {
		found := false
		for _, resultHostname := range result {
			if hostname == resultHostname {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected hostname %s not found in the result", hostname)
		}
	}
}

func TestIntegration_GetHostsWithActiveIPs(t *testing.T) {
	// Set the THRESHOLD environment variable for testing
	os.Setenv("THRESHOLD", "1")

	// Initialize the Gin router
	router := gin.Default()

	// Define the /hostnames endpoint
	router.GET("/hostnames", RetrieveInactiveHostnames)

	// Set up a test HTTP server
	server := httptest.NewServer(router)
	defer server.Close()

	// Create a new HTTP request to the /hostnames endpoint
	req, err := http.NewRequest("GET", server.URL+"/hostnames", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Send the HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	// Convert the response body to string
	responseBody := string(body)

	// Define the expected response body
	expectedBody := `["mta-prod-1","mta-prod-3"]`

	// Verify the response body
	if responseBody != expectedBody {
		t.Errorf("expected response body %s, but got %s", expectedBody, responseBody)
	}
}
