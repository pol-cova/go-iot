package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pol-cova/terrabyte/internal/api/handlers"
	"github.com/pol-cova/terrabyte/internal/api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	// API routes
	server.GET("/status", healthCheck)
	server.GET("/ping", pong)

	// Telemetry routes
	telemetryHandler := handlers.NewTelemetryHandler()
	api := server.Group("/api/v1")
	{
		api.POST("/telemetry/", middlewares.AuthMiddleware, telemetryHandler.CreateTelemetry)
		api.GET("/telemetry", telemetryHandler.GetTelemetry)
		api.GET("/telemetry/latest", telemetryHandler.GetLatestRecord)
		api.GET("/dashboard", telemetryHandler.ShowDashboard)
		api.GET("/latest-data", telemetryHandler.GetLatestHTML)
	}
}
