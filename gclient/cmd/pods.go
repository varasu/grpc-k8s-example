package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/varasu/grpc-k8s-example/gclient/internal/client"
)

var namespace string

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "Fetch pods names from the gservice server.",
	Long:  `This is the main command for grpc k8s example`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("call pods %s %s\n", namespace, url)

		gserviceClient := client.NewGService(url)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		pods, err := gserviceClient.ListPods(ctx, namespace)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Namespace: %s\n\n", namespace)
		if len(pods) == 0 {
			fmt.Println("No pods are running in this namespace.")
			return
		}
		for i, pod := range pods {
			fmt.Printf("%d: %s\n", i, pod.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(podsCmd)
	podsCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "k8s namespace")
}
