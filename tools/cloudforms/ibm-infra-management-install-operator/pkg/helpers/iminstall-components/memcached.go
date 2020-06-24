package iminstalltools

import (
	"os"

	iminstallv1alpha1 "github.ibm.com/IBMPrivateCloud/ibm-infra-management-install-operator/pkg/apis/infra/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

func NewMemcachedDeployment(cr *iminstallv1alpha1.IMInstall) (*appsv1.Deployment, error) {
	memcachedImage := os.Getenv("MEMCACHED_IMAGE")
	if memcachedImage == "" {
		memcachedImage = cr.Spec.MemcachedImageName + ":" + cr.Spec.MemcachedImageTag
	}
	container := corev1.Container{
		Name:  "memcached",
		Image: memcachedImage,

		Ports: []corev1.ContainerPort{
			corev1.ContainerPort{
				ContainerPort: 11211,
				Protocol:      "TCP",
			},
		},

		LivenessProbe: &corev1.Probe{
			Handler: corev1.Handler{
				TCPSocket: &corev1.TCPSocketAction{
					Port: intstr.FromInt(11211),
				},
			},
		},
		ReadinessProbe: &corev1.Probe{
			Handler: corev1.Handler{
				TCPSocket: &corev1.TCPSocketAction{
					Port: intstr.FromInt(11211),
				},
			},
		},
		Env: []corev1.EnvVar{
			corev1.EnvVar{
				Name:  "MEMCACHED_MAX_MEMORY",
				Value: cr.Spec.MemcachedMaxMemory,
			},
			corev1.EnvVar{
				Name:  "MEMCACHED_MAX_CONNECTIONS",
				Value: cr.Spec.MemcachedMaxConnection,
			},
			corev1.EnvVar{
				Name:  "MEMCACHED_SLAB_PAGE_SIZE",
				Value: cr.Spec.MemcachedSlabPageSize,
			},
		},
	}

	err := addResourceReqs(cr.Spec.MemcachedMemoryLimit, cr.Spec.MemcachedMemoryRequest, cr.Spec.MemcachedCpuLimit, cr.Spec.MemcachedCpuRequest, &container)
	if err != nil {
		return nil, err
	}

	deploymentLabels := map[string]string{
		"app": cr.Spec.AppName,
	}
	podLabels := map[string]string{
		"name": "memcached",
		"app":  cr.Spec.AppName,
	}
	var repNum int32 = 1

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "memcached",
			Namespace: cr.ObjectMeta.Namespace,
			Labels:    deploymentLabels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &repNum,
			Selector: &metav1.LabelSelector{
				MatchLabels: podLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:        "memcached",
					Labels:      podLabels,
					Annotations: AnnotationsForPod(),
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{container},
				},
			},
		},
	}

	return deployment, nil
}

func NewMemcachedService(cr *iminstallv1alpha1.IMInstall) *corev1.Service {
	labels := map[string]string{
		"app": cr.Spec.AppName,
	}
	selector := map[string]string{
		"name": "memcached",
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "memcached",
			Namespace: cr.ObjectMeta.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Name: "memcached",
					Port: 11211,
				},
			},
			Selector: selector,
		},
	}
}
