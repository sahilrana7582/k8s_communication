package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sahilrana27582/k8s_communication/internal/config"
	"github.com/sahilrana27582/k8s_communication/internal/handler"
	"github.com/sahilrana27582/k8s_communication/internal/k8s"
	"github.com/sahilrana27582/k8s_communication/internal/routes"
	"github.com/sahilrana27582/k8s_communication/internal/service"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Application stopped: %v", err)
		os.Exit(1)
	}
}

func run() error {
	// 1. Load configuration
	cfg := config.Load()

	// 2. Initialize Kubernetes client
	k8sClient, err := k8s.NewClient(cfg)
	if err != nil {
		return fmt.Errorf("initialize kubernetes client: %w", err)
	}

	// 3. Initialize services
	podService := service.NewPodService(k8sClient)

	// 4. Initialize handlers
	handlers := handler.Handlers{
		Pod: handler.NewPodHandler(podService),
	}

	// 5. Initialize Gin router
	router := gin.Default()

	// 6. Register routes
	routes.RegisterRoutes(router, &handlers)

	// 7. Initialize HTTP server
	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// 8. Start server
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("SRE API server starting on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// 9. Listen for termination signals
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("HTTP server failed: %w", err)

	case <-ctx.Done():
		log.Println("Shutdown signal received")
	}

	// 10. Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		_ = server.Close()
		return fmt.Errorf("graceful shutdown failed: %w", err)
	}

	log.Println("SRE API server stopped gracefully")

	return nil
}
