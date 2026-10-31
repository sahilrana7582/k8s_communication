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
	podClient typedV1.PodsGetter
}

func NewPodService(k8sClient *k8s.K8sClient) *PodService {
	coreClient := k8sClient.Client.CoreV1()

	return &PodService{
		podClient: coreClient,
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

// GetPod retrieves a specific Pod.
func (p *PodService) GetPod(
	ctx context.Context,
	ns string,
	podName string,
) (*corev1.Pod, error) {

	pod, err := p.podClient.Pods(ns).Get(
		ctx,
		podName,
		metaV1.GetOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("get pod %s/%s: %w", ns, podName, err)
	}

	return pod, nil
}

// GetPodLogs retrieves logs from a Pod container.
func (p *PodService) GetPodLogs(
	ctx context.Context,
	ns string,
	podName string,
	opts *corev1.PodLogOptions,
) ([]byte, error) {

	if opts == nil {
		tailLines := int64(100)
		opts = &corev1.PodLogOptions{
			TailLines: &tailLines,
		}
	}

	logs, err := p.podClient.
		Pods(ns).
		GetLogs(podName, opts).
		Do(ctx).
		Raw()

	if err != nil {
		return nil, fmt.Errorf("failed to get logs for pod %s/%s: %w", ns, podName, err)
	}

	return logs, nil
}

// CreatePod creates a new Pod.
func (p *PodService) CreatePod(
	ctx context.Context,
	pod *corev1.Pod,
) (*corev1.Pod, error) {

	if pod == nil {
		return nil, fmt.Errorf("pod cannot be nil")
	}

	createdPod, err := p.podClient.Pods(pod.Namespace).Create(
		ctx,
		pod,
		metaV1.CreateOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("create pod in namespace %s: %w", pod.Namespace, err)
	}

	return createdPod, nil
}

// // UpdatePod updates an existing Pod.
// func (p *PodService) UpdatePod(
// 	ctx context.Context,
// 	ns string,
// 	pod *corev1.Pod,
// ) (*corev1.Pod, error) {

// 	if pod == nil {
// 		return nil, fmt.Errorf("pod cannot be nil")
// 	}

// 	updatedPod, err := p.podClient.Pods(ns).Update(
// 		ctx,
// 		pod,
// 		metaV1.UpdateOptions{},
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("update pod in namespace %s: %w", ns, err)
// 	}

// 	return updatedPod, nil
// }

// DeletePod deletes a Pod.
func (p *PodService) DeletePod(
	ctx context.Context,
	ns string,
	podName string,
) error {

	err := p.podClient.Pods(ns).Delete(
		ctx,
		podName,
		metaV1.DeleteOptions{},
	)
	if err != nil {
		return fmt.Errorf("delete pod %s/%s: %w", ns, podName, err)
	}

	return nil
}
