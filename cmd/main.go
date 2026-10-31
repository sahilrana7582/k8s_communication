package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sahilrana27582/k8s_communication/internal/config"
	"github.com/sahilrana27582/k8s_communication/internal/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	cfg := config.Load()

	k8sClient, err := k8s.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	podInfo := k8sClient.Client.CoreV1().Pods("taskflow")
	podList, err := podInfo.List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, pod := range podList.Items {
		fmt.Println(pod.Name)
	}

}
