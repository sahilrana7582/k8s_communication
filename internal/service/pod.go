package service

import (
	"context"
	"fmt"

	"github.com/sahilrana27582/k8s_communication/internal/k8s"

	corev1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedV1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type PodService struct {
	podClient   typedV1.PodsGetter
	eventClient typedV1.EventsGetter
}

func NewPodService(k8sClient *k8s.K8sClient) *PodService {
	coreClient := k8sClient.Client.CoreV1()

	return &PodService{
		podClient:   coreClient,
		eventClient: coreClient,
	}
}

// GetPods retrieves Pods from a namespace.
// An empty namespace retrieves Pods across all namespaces.
func (p *PodService) GetPods(
	ctx context.Context,
	ns string,
) ([]corev1.Pod, error) {

	podList, err := p.podClient.Pods(ns).List(
		ctx,
		metaV1.ListOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("list pods: %w", err)
	}

	return podList.Items, nil
}

// GetPodsByLabels retrieves Pods matching a label selector.
func (p *PodService) GetPodsByLabels(
	ctx context.Context,
	ns string,
	selector string,
) ([]corev1.Pod, error) {

	podList, err := p.podClient.Pods(ns).List(
		ctx,
		metaV1.ListOptions{
			LabelSelector: selector,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("list pods by labels: %w", err)
	}

	return podList.Items, nil
}
