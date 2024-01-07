package unit_tests

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
)

func TestPodTemplateRendersContainerImage(t *testing.T) {

	// Path to the helm chart we will test
	helmChartPath := "../../traefik-infrastruktur"

	// Setup the args. For this test, we will set the following input values:
	// - image=nginx:1.15.8
	options := &helm.Options{
		SetValues: map[string]string{"ServiceAccountName": "account"},
	}

	// Run RenderTemplate to render the template and capture the output.
	output := helm.RenderTemplate(t, options, helmChartPath, "serviceaccount", []string{"templates/00-account.yml"})

	// Now we use kubernetes/client-go library to render the template output into the Pod struct. This will
	// ensure the Pod resource is rendered correctly.
	var serviceaccount corev1.ServiceAccount
	helm.UnmarshalK8SYaml(t, output, &serviceaccount)

	// Finally, we verify the pod spec is set to the expected container image value
	expectedContainerImage := "account"
	podContainers := serviceaccount.ObjectMeta.Name
	if podContainers != expectedContainerImage {
		t.Fatalf("Rendered container image (%s) is not expected (%s)", podContainers, expectedContainerImage)
	}
}
