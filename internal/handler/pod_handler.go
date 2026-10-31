package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahilrana27582/k8s_communication/internal/service"
)

type Handler struct {
	Pod *PodHandler
}

type PodHandler struct {
	srv *service.PodService
}

func NewPodHandler(srv *service.PodService) *PodHandler {
	return &PodHandler{srv: srv}
}

func (h *PodHandler) GetPod(c *gin.Context) {
	namespace := c.Param("namespace")

	pods, err := h.srv.GetPods(c.Request.Context(), namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pods retrieved successfully",
		"data":    pods,
	})
}
