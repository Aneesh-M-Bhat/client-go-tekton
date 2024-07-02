package main

import (
	"context"
	"fmt"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func listCustomRuns(client *clientSetType) {
	res, err := client.TektonV1beta1().CustomRuns("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range res.Items {
		fmt.Printf(" * %s \n", d.GetName())
	}
	prompt()
}

func createCustomRun(client *clientSetType, res *v1beta1.CustomRun) {
	createdRes, err := client.TektonV1beta1().CustomRuns("default").Create(context.TODO(), res, v1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(createdRes)
	prompt()
}

func deleteCustomRun(client *clientSetType, resName string) {
	err := client.TektonV1beta1().CustomRuns("default").Delete(context.TODO(), resName, v1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted CustomRun %s\n", resName)
	prompt()
}
