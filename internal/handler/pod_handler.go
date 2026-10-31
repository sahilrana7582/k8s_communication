package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sahilrana27582/k8s_communication/internal/service"
	corev1 "k8s.io/api/core/v1"
)

type Handlers struct {
	Pod *PodHandler
}

type PodHandler struct {
	srv *service.PodService
}

func NewPodHandler(srv *service.PodService) *PodHandler {
	return &PodHandler{srv: srv}
}

func (h *PodHandler) GetPods(c *gin.Context) {
	namespace := c.Param("namespace")

	if strings.TrimSpace(namespace) == "" {
		namespace = corev1.NamespaceDefault
	}

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

func (h *PodHandler) GetPod(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	pod, err := h.srv.GetPod(c.Request.Context(), namespace, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Pods retrieved successfully",
		"data":    pod,
	})
}

type PodLogsRequest struct {
	Container    string `form:"container"`
	TailLines    *int64 `form:"tailLines"`
	SinceSeconds *int64 `form:"sinceSeconds"`
	Timestamps   bool   `form:"timestamps"`
	Previous     bool   `form:"previous"`
}

func (h *PodHandler) GetPodLogs(c *gin.Context) {
	namespace := c.Param("namespace")
	name := c.Param("name")

	var req PodLogsRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid log query parameters",
			"error":   err.Error(),
		})
		return
	}

	if req.TailLines != nil && *req.TailLines < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "tailLines cannot be negative",
		})
		return
	}

	if req.SinceSeconds != nil && *req.SinceSeconds < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "sinceSeconds cannot be negative",
		})
		return
	}

	opts := &corev1.PodLogOptions{
		Container:    req.Container,
		TailLines:    req.TailLines,
		SinceSeconds: req.SinceSeconds,
		Timestamps:   req.Timestamps,
		Previous:     req.Previous,
	}

	podLogs, err := h.srv.GetPodLogs(
		c.Request.Context(),
		namespace,
		name,
		opts,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to fetch pod logs",
			"error":   err.Error(),
		})
		return
	}

	c.Data(http.StatusOK, "text/plain; charset=utf-8", podLogs)
}
