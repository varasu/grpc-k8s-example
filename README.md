# grpc-k8s-example

The goal of this example is to showcase basic grpc client/server and interaction with k8s api (`client-go`).

## Steps to run grpc-k8s-example

You will need [go](https://go.dev/doc/install), [kind](https://kind.sigs.k8s.io/docs/user/quick-start/) and [docker](https://docs.docker.com/engine/install/).


Clone `grpc-k8s-example`.

```bash
git clone https://github.com/varasu/grpc-k8s-example.git
```

Enter `grpc-k8s-example` directory and build gclient binary.

```bash
cd gclient
go build -o ./gclient gservice/main.go
```
With `gclient` you could interact with grpc server - `gservice`.

## k8s
Build `gservice` as docker image.

```bash
docker build -t gservice-image .
```

Now we could create and start our k8s cluster with kind. We are going to use the default cluster configuration for the purpose of this example.

```bash
kind create cluster --name grpc-example
```

Do you see your cluster?
```bash
kind get clusters
```

Perfect. 

Now we are ready to load the image we previously build.

```bash
kind --name grpc-example load docker-image gservice-image
```

We need to create namespace and cluster role binding - `privileges` namespace could be able to do all the admin work. 

```bash
kubectl --context kind-grpc-example create namespace privileges
kubectl --context kind-grpc-example apply -f res/ClusterRoleBinding.yaml
```

We are ready for some deploying!

```bash
kubectl --context kind-grpc-example apply -f res/gservice-deployment.yaml
```

Before you could play and interact with the server, we should expose it. For the purpose of this example we will use plain old port forwarding mechanism.

```bash
kubectl port-forward --context kind-grpc-example --namespace privileges gservice-example-8b78f4cd-qj69v 8080:8080
```

You should replace the name of the pod. 

## Client
After we have everything set up and running, we could finally are ready to try to list all of the pods in `kube-system` namespace or just all.

```bash
gclient pods
gclient pods -n kube-system
```

## Usage
```
Execute commands to gservice grpc server.

Usage:
  gclient [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  pods        Fetch pods names from the gservice server.

Flags:
  -h, --help         help for gclient
  -t, --toggle       Help message for toggle
      --url string   gservice url (default "localhost:8080")

Use "gclient [command] --help" for more information about a command.
```

## Testing
E2E testing with `grpcurl` and `bash` on running instance.
