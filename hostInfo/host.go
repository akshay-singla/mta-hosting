package hostinfo

import (
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Server represents a server hosting MTAs.
type Server struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	Active   bool   `json:"active"`
}

// Servers represents a collection of servers.
type Servers []Server

// RetrieveInactiveHostnames retrieves hostnames with less than or equal to X inactive IP addresses.
func RetrieveInactiveHostnames(c *gin.Context) {
	threshold := getThreshold()
	inactiveHostnames := getInactiveHostnames(threshold)

	log.Println(inactiveHostnames)
	c.JSON(200, inactiveHostnames)
}

// getThreshold retrieves the threshold value from the environment variable.
// If the environment variable is not set, it defaults to 1.
func getThreshold() int {
	thresholdStr := os.Getenv("THRESHOLD")
	if thresholdStr == "" {
		return 1
	}

	threshold, err := strconv.Atoi(thresholdStr)
	if err != nil {
		log.Printf("Invalid threshold value: %s. Using default value of 1.", thresholdStr)
		return 1
	}

	return threshold
}

// getInactiveHostnames retrieves the hostnames with less than or equal to X inactive IP addresses.
func getInactiveHostnames(threshold int) []string {
	ipConfig := getIPConfig()
	hostnames := make(map[string]int)

	// Count the number of inactive IP addresses for each hostname
	for _, server := range ipConfig {
		if !server.Active {
			hostnames[server.Hostname]++
			if hostnames[server.Hostname] > threshold {
				delete(hostnames, server.Hostname)
			}
		}
	}

	// Collect the hostnames
	result := make([]string, 0, len(hostnames))
	for hostname := range hostnames {
		result = append(result, hostname)
	}

	// sort the hostnames in result
	sort.Strings(result)

	return result
}

// getIPConfig retrieves the IP configuration data (mocked here).
func getIPConfig() Servers {
	return Servers{
		{IP: "127.0.0.1", Hostname: "mta-prod-1", Active: true},
		{IP: "127.0.0.2", Hostname: "mta-prod-1", Active: false},
		{IP: "127.0.0.3", Hostname: "mta-prod-2", Active: true},
		{IP: "127.0.0.4", Hostname: "mta-prod-2", Active: true},
		{IP: "127.0.0.5", Hostname: "mta-prod-2", Active: true},
		{IP: "127.0.0.6", Hostname: "mta-prod-3", Active: false},
	}
}
