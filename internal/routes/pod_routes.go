package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahilrana27582/k8s_communication/internal/handler"
)

func RegisterPodRoutes(rg *gin.RouterGroup, h *handler.PodHandler) {
	pods := rg.Group("/:namespace/pods")
	{
		pods.GET("", h.GetPods)
		pods.GET("/:name", h.GetPod)
		pods.GET("/:name/logs", h.GetPodLogs)
		pods.POST("", h.CreatePod)
		pods.DELETE("/:name", h.DeletePod)
	}
}

func RegisterHealthChecks(rg *gin.RouterGroup) {
	rg.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service Healthy",
		})
	})
}
