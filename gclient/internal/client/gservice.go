package client

import (
	"context"
	"fmt"
	"log"

	api "github.com/varasu/grpc-k8s-example/gservice/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GService struct {
	client api.GServiceClient
}

func NewGService(URL string) GService {
	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("%v", err)
	}
	client := api.NewGServiceClient(conn)
	return GService{client: client}
}

func (c *GService) ListPods(ctx context.Context, namespace string) ([]*api.Pod, error) {
	resp, err := c.client.ListPods(ctx, &api.ListPodsRequest{Namespace: namespace})
	if err != nil {
		return nil, fmt.Errorf("ListPods err: %w", err)
	}
	return resp.Pods, nil
}
