package iminstalltools

import (
	iminstallv1alpha1 "github.ibm.com/IBMPrivateCloud/ibm-infra-management-install-operator/pkg/apis/infra/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func HttpdServiceAccount(cr *iminstallv1alpha1.IMInstall) *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Spec.AppName + "-httpd",
			Namespace: cr.ObjectMeta.Namespace,
		},
	}
}

func OrchestratorServiceAccount(cr *iminstallv1alpha1.IMInstall) *corev1.ServiceAccount {
	return &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      orchestratorObjectName(cr),
			Namespace: cr.ObjectMeta.Namespace,
		},
	}
}

func OrchestratorRole(cr *iminstallv1alpha1.IMInstall) *rbacv1.Role {
	return &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      orchestratorObjectName(cr),
			Namespace: cr.ObjectMeta.Namespace,
		},
		Rules: []rbacv1.PolicyRule{
			rbacv1.PolicyRule{
				APIGroups: []string{""},
				Resources: []string{"pods", "pods/finalizers"},
				Verbs:     []string{"*"},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{"apps"},
				Resources: []string{"deployments", "deployments/scale"},
				Verbs:     []string{"*"},
			},
			rbacv1.PolicyRule{
				APIGroups: []string{"extensions"},
				Resources: []string{"deployments", "deployments/scale"},
				Verbs:     []string{"*"},
			},
		},
	}
}

func OrchestratorRoleBinding(cr *iminstallv1alpha1.IMInstall) *rbacv1.RoleBinding {
	return &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      orchestratorObjectName(cr),
			Namespace: cr.ObjectMeta.Namespace,
		},
		RoleRef: rbacv1.RoleRef{
			Kind:     "Role",
			Name:     orchestratorObjectName(cr),
			APIGroup: "rbac.authorization.k8s.io",
		},
		Subjects: []rbacv1.Subject{
			rbacv1.Subject{
				Kind: "ServiceAccount",
				Name: orchestratorObjectName(cr),
			},
		},
	}
}

func orchestratorObjectName(cr *iminstallv1alpha1.IMInstall) string {
	return cr.Spec.AppName + "-orchestrator"
}
