package iminstalltools

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Important product values needed for annotations
const LicensingCloudpakId = "7f6eda41081c4e08a255be1f0b4aef2d"
const LicensingCloudpakName = "IBM Cloud Pak for Multicloud Management"
const LicensingCloudpakVersion = "2.0"
const LicensingProductChargedContainers = "All"
const LicensingProductCloudpakRatio = "1:1"
const LicensingProductMetric = "MANAGED_VIRTUAL_SERVER"

const LicensingProductName = "IBM Cloud Pak for Multicloud Management Infrastructure Management"
const LicensingProductID = "4747644c9bae4473aa336e128c3cc3e9"
const LicensingProductVersion = "2.0"

func AnnotationsForPod() map[string]string {
	return map[string]string{
		"productName":              LicensingProductName,
		"productID":                LicensingProductID,
		"productVersion":           LicensingProductVersion,
		"productMetric":            LicensingProductMetric,
		"cloudpakId":               LicensingCloudpakId,
		"cloudpakName":             LicensingCloudpakName,
		"cloudpakVersion":          LicensingCloudpakVersion,
		"productChargedContainers": LicensingProductChargedContainers,
	}
}

func addResourceReqs(memLimit, memReq, cpuLimit, cpuReq string, c *corev1.Container) error {
	if memLimit == "" && memReq == "" && cpuLimit == "" && cpuReq == "" {
		return nil
	}

	if memLimit != "" || cpuLimit != "" {
		c.Resources.Limits = make(map[corev1.ResourceName]resource.Quantity)
	}

	if memLimit != "" {
		limit, err := resource.ParseQuantity(memLimit)
		if err != nil {
			return err
		}
		c.Resources.Limits["memory"] = limit
	}

	if cpuLimit != "" {
		limit, err := resource.ParseQuantity(cpuLimit)
		if err != nil {
			return err
		}
		c.Resources.Limits["cpu"] = limit
	}

	if memReq != "" || cpuReq != "" {
		c.Resources.Requests = make(map[corev1.ResourceName]resource.Quantity)
	}

	if memReq != "" {
		req, err := resource.ParseQuantity(memReq)
		if err != nil {
			return err
		}
		c.Resources.Requests["memory"] = req
	}

	if cpuReq != "" {
		req, err := resource.ParseQuantity(cpuReq)
		if err != nil {
			return err
		}
		c.Resources.Requests["cpu"] = req
	}

	return nil
}
