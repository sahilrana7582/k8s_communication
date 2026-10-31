package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilrana27582/k8s_communication/internal/handler"
)

func RegisterPodRoutes(rg *gin.RouterGroup, h *handler.PodHandler) {
	pods := rg.Group("/pods")
	{
		pods.GET("", h.GetPods)
		pods.GET("/:name", h.GetPod)
		pods.GET("/:name/logs", h.GetPodLogs)
	}
}
