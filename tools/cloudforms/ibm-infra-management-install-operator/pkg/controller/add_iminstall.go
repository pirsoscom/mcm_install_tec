package controller

import (
	"github.ibm.com/IBMPrivateCloud/ibm-infra-management-install-operator/pkg/controller/iminstall"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, iminstall.Add)
}
