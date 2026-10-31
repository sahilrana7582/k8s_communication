package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sahilrana27582/k8s_communication/internal/handler"
)

func RegisterRoutes(r *gin.Engine, h *handler.Handlers) {
	api := r.Group("/app/v1")
	RegisterPodRoutes(api, h.Pod)
}
