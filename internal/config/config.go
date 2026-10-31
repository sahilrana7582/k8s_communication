package config

import (
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Port string

	K8s K8sConfig
}

type K8sConfig struct {
	KubeConfigPath string
	InCluster      bool
	Qps            int16
	Burst          int16
	RequestTimeout time.Duration
}

func Load() *Config {
	cfg := Config{
		Port: "8888",
		K8s:  K8sConfig{},
	}

	if home := os.Getenv("HOME"); home != "" {
		k8sConfigPath := filepath.Join(home, ".kube", "config")
		cfg.K8s.KubeConfigPath = k8sConfigPath
		cfg.K8s.Qps = 15
		cfg.K8s.Burst = 20
		cfg.K8s.RequestTimeout = 10 * time.Second
	}

	return &cfg
}
