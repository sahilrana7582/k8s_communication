package service

import (
	"context"

	"github.com/sahilrana27582/k8s_communication/internal/k8s"
	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedV1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type PodService struct {
	podClient typedV1.PodsGetter
}

func NewPodService(k8sClient *k8s.K8sClient) *PodService {
	return &PodService{podClient: k8sClient.Client.CoreV1()}
}

func (p *PodService) GetPods(ctx context.Context, ns string) ([]corev1.Pod, error) {
	podList, err := p.podClient.Pods(ns).List(ctx, metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return podList.Items, nil
}
