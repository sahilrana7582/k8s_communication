package main

import (
	"log"

	"github.com/sahilrana27582/k8s_communication/internal/config"
	"github.com/sahilrana27582/k8s_communication/internal/k8s"
)

func main() {
	cfg := config.Load()

	k8sClient, err := k8s.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
