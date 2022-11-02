package server

import (
	"context"
	"log"

	api "github.com/varasu/grpc-k8s-example/gservice/api/v1"

	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	port = ":50051"
)

type grpcServer struct {
	api.UnimplementedGServiceServer
	Clientset *kubernetes.Clientset
}

func (s *grpcServer) ListPods(ctx context.Context, req *api.ListPodsRequest) (*api.Pods, error) {
	namespace := string(req.Namespace)

	log.Printf("get pods for %s\n", namespace)
	pods, err := s.Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	data := []*api.Pod{}
	for _, pod := range pods.Items {
		log.Printf("%+v", pod)
		_pod := api.Pod{Name: pod.Name}
		data = append(data, &_pod)
	}

	return &api.Pods{Pods: data}, nil
}

func NewGRPCServer() *grpc.Server {
	var clientset *kubernetes.Clientset
	var err error
	if clientset, err = NewInClusterClientset(); err != nil {
		log.Fatal(err)
	}
	gsrv := grpc.NewServer()
	srv := grpcServer{
		Clientset: clientset,
	}

	api.RegisterGServiceServer(gsrv, &srv)
	return gsrv
}

type Pod struct {
	Name string
}
