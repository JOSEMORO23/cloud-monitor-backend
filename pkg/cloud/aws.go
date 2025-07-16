package cloud

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

// ListInstancesAWS imprime IDs y estados de instancias EC2
func ListInstancesAWS() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("error cargando configuración AWS: %w", err)
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{}
	result, err := client.DescribeInstances(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("error describiendo instancias: %w", err)
	}

	fmt.Println("✅ Instancias EC2 encontradas:")
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			fmt.Printf("- ID: %s, Estado: %s\n", *instance.InstanceId, instance.State.Name)
		}
	}

	return nil
}