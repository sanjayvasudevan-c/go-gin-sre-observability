package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.New()

	// JSON structured logs
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(
			`{"time":"%s","method":"%s","path":"%s","status":%d,"latency":"%s","client":"%s"}`+"\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
		)
	}))
	r.Use(gin.Recovery())

	// Native Prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().UTC(),
		})
	})

	r.Run(":8080")
}
