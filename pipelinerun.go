package main

import (
	"context"
	"fmt"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func listPipelineRuns(client *clientSetType) {
	res, err := client.TektonV1beta1().PipelineRuns("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range res.Items {
		fmt.Printf(" * %s \n", d.GetName())
	}
	prompt()
}

func createPipelineRun(client *clientSetType, res *v1beta1.PipelineRun) {
	createdRes, err := client.TektonV1beta1().PipelineRuns("default").Create(context.TODO(), res, v1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(createdRes)
	prompt()
}

func deletePipelineRun(client *clientSetType, resName string) {
	err := client.TektonV1beta1().PipelineRuns("default").Delete(context.TODO(), resName, v1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted PipelineRun %s\n", resName)
	prompt()
}
