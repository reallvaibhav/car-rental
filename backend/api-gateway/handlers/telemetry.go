package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestMetric struct {
	Count       int           `json:"count"`
	Errors      int           `json:"errors"`
	TotalTime   time.Duration `json:"total_time_ms"`
	AverageTime time.Duration `json:"average_time_ms"`
	LastAccess  time.Time     `json:"last_access"`
}

type Metrics struct {
	RequestMetrics map[string]*RequestMetric
	StartTime      time.Time
	mu             sync.Mutex
}

func NewMetrics() *Metrics {
	return &Metrics{
		RequestMetrics: make(map[string]*RequestMetric),
		StartTime:      time.Now(),
	}
}

func (m *Metrics) Track() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		key := method + ":" + path

		startTime := time.Now()

		// Process request
		c.Next()

		// Record metrics
		duration := time.Since(startTime)
		statusCode := c.Writer.Status()

		m.mu.Lock()
		defer m.mu.Unlock()

		metric, exists := m.RequestMetrics[key]
		if !exists {
			metric = &RequestMetric{
				LastAccess: time.Now(),
			}
			m.RequestMetrics[key] = metric
		}

		metric.Count++
		if statusCode >= 400 {
			metric.Errors++
		}
		metric.TotalTime += duration
		metric.AverageTime = time.Duration(int64(metric.TotalTime) / int64(metric.Count))
		metric.LastAccess = time.Now()
	}
}

func (m *Metrics) GetMetrics(c *gin.Context) {
	m.mu.Lock()
	defer m.mu.Unlock()

	uptime := time.Since(m.StartTime).String()

	c.JSON(http.StatusOK, gin.H{
		"uptime":   uptime,
		"requests": m.RequestMetrics,
	})
}
