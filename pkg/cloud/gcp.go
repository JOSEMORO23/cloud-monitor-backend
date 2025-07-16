package cloud

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	"google.golang.org/api/option"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

// ListInstancesGCP imprime nombres de instancias en un proyecto/Zona
func ListInstancesGCP(projectID, zone, credentialsFile string) error {
	ctx := context.Background()

	client, err := compute.NewInstancesRESTClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return fmt.Errorf("error creando cliente GCP: %w", err)
	}
	defer client.Close()

	req := &computepb.ListInstancesRequest{
		Project: projectID,
		Zone:    zone,
	}

	it := client.List(ctx, req)
	fmt.Println("✅ Instancias en GCP:")
	for {
		instance, err := it.Next()
		if err != nil {
			break
		}
		fmt.Printf("- Nombre: %s\n", instance.GetName())
	}

	return nil
}