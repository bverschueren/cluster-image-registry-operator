package swift

import (
	corev1 "k8s.io/api/core/v1"

	opapi "github.com/openshift/cluster-image-registry-operator/pkg/apis/dockerregistry/v1alpha1"
)

type driver struct {
	Config *opapi.OpenShiftDockerRegistryConfigStorageSwift
}

func NewDriver(c *opapi.OpenShiftDockerRegistryConfigStorageSwift) *driver {
	return &driver{
		Config: c,
	}
}

func (d *driver) GetName() string {
	return "swift"
}

func (d *driver) ConfigEnv() (envs []corev1.EnvVar, err error) {
	envs = append(envs,
		corev1.EnvVar{Name: "REGISTRY_STORAGE", Value: d.GetName()},
		corev1.EnvVar{Name: "REGISTRY_STORAGE_SWIFT_AUTHURL", Value: d.Config.AuthURL},
		corev1.EnvVar{Name: "REGISTRY_STORAGE_SWIFT_CONTAINER", Value: d.Config.Container},
		corev1.EnvVar{
			Name: "REGISTRY_STORAGE_SWIFT_USERNAME",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "image-registry-private-configuration",
					},
					Key: "REGISTRY_STORAGE_SWIFT_USERNAME",
				},
			},
		},
		corev1.EnvVar{
			Name: "REGISTRY_STORAGE_SWIFT_PASSWORD",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "image-registry-private-configuration",
					},
					Key: "REGISTRY_STORAGE_SWIFT_PASSWORD",
				},
			},
		},
	)
	return
}

func (d *driver) Volumes() ([]corev1.Volume, []corev1.VolumeMount, error) {
	return nil, nil, nil
}
