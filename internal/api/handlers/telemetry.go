package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pol-cova/terrabyte/internal/models"
	"github.com/pol-cova/terrabyte/views"
	"net/http"
	"sync"
	"time"
)

type TelemetryHandler struct {
	mu       sync.RWMutex
	readings []models.SensorData
}

func NewTelemetryHandler() *TelemetryHandler {
	return &TelemetryHandler{
		readings: make([]models.SensorData, 0),
	}
}

func (h *TelemetryHandler) CreateTelemetry(c *gin.Context) {
	var data models.SensorData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// timestamp
	data.Timestamp = time.Now()
	data.ID = uint(len(h.readings) + 1)

	// Mac device
	MAC := c.Request.Header.Get("Device-MAC")
	data.DeviceMAC = MAC

	// save data in memory
	h.mu.Lock()
	h.readings = append(h.readings, data)
	h.mu.Unlock()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Data saved",
		"data":    data,
	})
}

func (h *TelemetryHandler) GetTelemetry(c *gin.Context) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.readings) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No data available",
			"data":    []models.SensorData{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data retrieved",
		"data":    h.readings,
	})
}

func (h *TelemetryHandler) GetLatestRecord(c *gin.Context) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.readings) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No data available",
		})
		return
	}
	lastReading := h.readings[len(h.readings)-1]
	c.JSON(http.StatusOK, gin.H{
		"message": "Latest reading",
		"data":    lastReading,
	})
}

func (h *TelemetryHandler) ShowDashboard(c *gin.Context) {
	component := views.Dashboard()
	component.Render(c.Request.Context(), c.Writer)
}

func (h *TelemetryHandler) GetLatestHTML(c *gin.Context) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.readings) == 0 {
		c.Writer.Header().Set("Content-Type", "text/html")
		component := views.Error("No hay datos disponibles")
		component.Render(c.Request.Context(), c.Writer)
		return
	}

	latestReading := h.readings[len(h.readings)-1]
	c.Writer.Header().Set("Content-Type", "text/html")
	component := views.LatestData(latestReading)
	component.Render(c.Request.Context(), c.Writer)
}
