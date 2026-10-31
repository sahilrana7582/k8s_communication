package k8s

import (
	"fmt"

	"github.com/sahilrana27582/k8s_communication/internal/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient struct {
	Client *kubernetes.Clientset
}

func NewClient(cfg *config.Config) (*K8sClient, error) {
	kubeCfg, err := clientcmd.BuildConfigFromFlags("", cfg.K8s.KubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("Error for kubeconfig: %w", err)
	}

	kubeCfg.QPS = float32(cfg.K8s.Qps)
	kubeCfg.Burst = int(cfg.K8s.Burst)
	kubeCfg.Timeout = cfg.K8s.RequestTimeout

	clientSet, err := kubernetes.NewForConfig(kubeCfg)
	if err != nil {
		return nil, fmt.Errorf("Error while creating k8s client: %w", err)
	}

	return &K8sClient{Client: clientSet}, nil
}
