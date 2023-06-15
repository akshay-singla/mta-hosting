package main

import (
	hostinfo "github.com/akshay-singla/mta-hosting/hostInfo"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Define the /hostnames endpoint
	router.GET("/hostnames", hostinfo.RetrieveInactiveHostnames)

	// Start the server
	router.Run(":8080")
}
