package ocm

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// GetFirstAPIServerEndpoint returns the first IP:port of the kubernetes Service from EndpointSlices
func GetFirstAPIServerEndpoint() (string, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load in-cluster config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", fmt.Errorf("failed to create clientset: %v", err)
	}

	endpointSlices, err := clientset.DiscoveryV1().EndpointSlices("default").
		List(context.TODO(), metav1.ListOptions{
			LabelSelector: "kubernetes.io/service-name=kubernetes",
		})
	if err != nil {
		return "", fmt.Errorf("failed to get endpoint slices: %v", err)
	}

	for _, slice := range endpointSlices.Items {
		for _, ep := range slice.Endpoints {
			for _, addr := range ep.Addresses {
				for _, port := range slice.Ports {
					if port.Name != nil && *port.Name == "https" {
						return fmt.Sprintf("%s:%d", addr, *port.Port), nil
					}
				}
			}
		}
	}

	return "", fmt.Errorf("no endpoint found")
}
