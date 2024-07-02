package main

import (
	"context"
	"fmt"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func listTasks(client *clientSetType) {
	res, err := client.TektonV1beta1().Tasks("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range res.Items {
		fmt.Printf(" * %s \n", d.GetName())
	}
	prompt()
}

func createTask(client *clientSetType, res *v1beta1.Task) {
	createdRes, err := client.TektonV1beta1().Tasks("default").Create(context.TODO(), res, v1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(createdRes)
	prompt()
}

func deleteTask(client *clientSetType, resName string) {
	err := client.TektonV1beta1().Tasks("default").Delete(context.TODO(), resName, v1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted Task %s\n", resName)
	prompt()
}
