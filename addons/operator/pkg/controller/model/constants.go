package model

import corev1 "k8s.io/api/core/v1"

const (
	MoleServiceName              = "service"
	MoleConfigName               = "configmap"
	MoleIngressName              = "ingress"
	MoleJobName                  = "job"
	MoleDeploymentName           = "deployment"
	MolePodName                  = "pod"
	MoleHealthEndpoint           = "/api/health"
	MoleConfigVolumeName         = "volume"
	MoleLogsVolumeName           = "log"
	MoleMountPath                = "/mount"
	LogPath                      = "/tmp/dtstack/"
	MoleServiceAccountName       = "dtstack"
	MoleCom                      = "dtstack.com"
	DefaultMemoryRequest         = "100Mi"
	DefaultMemoryLimit           = "1Gi"
	DefaultLogSidecarMemoryLimit = "500Mi"
	DefaultCpuLimit              = "500m"
	DefaultCpuRequest            = "0"
	EnvHostAlias                 = "HostAlias"
)

var SupportResource = map[corev1.ResourceName]struct{}{
	corev1.ResourceCPU:    {},
	corev1.ResourceMemory: {},
}

var VolumeConfigMapMode int32 = 0755
