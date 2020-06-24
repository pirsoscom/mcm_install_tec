package iminstalltools

import (
	iminstallv1alpha1 "github.ibm.com/IBMPrivateCloud/ibm-infra-management-install-operator/pkg/apis/infra/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func AppSecret(cr *iminstallv1alpha1.IMInstall) *corev1.Secret {

	labels := map[string]string{
		"app": cr.Spec.AppName,
	}
	secret := map[string]string{
		"encryption-key": generateEncryptionKey(),
	}
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "app-secrets",
			Namespace: cr.ObjectMeta.Namespace,
			Labels:    labels,
		},
		StringData: secret,
	}
}
