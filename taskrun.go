package main

import (
	"context"
	"fmt"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func listTaskRuns(client *clientSetType) {
	res, err := client.TektonV1beta1().TaskRuns("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range res.Items {
		fmt.Printf(" * %s \n", d.GetName())
	}
	prompt()
}

func createTaskRun(client *clientSetType, res *v1beta1.TaskRun) {
	createdRes, err := client.TektonV1beta1().TaskRuns("default").Create(context.TODO(), res, v1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(createdRes)
	prompt()
}

func deleteTaskRun(client *clientSetType, resName string) {
	err := client.TektonV1beta1().TaskRuns("default").Delete(context.TODO(), resName, v1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted TaskRun %s\n", resName)
	prompt()
}
