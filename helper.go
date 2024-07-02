package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	clientset "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	AddToScheme = v1beta1.AddToScheme
)

type clientSetType = clientset.Clientset

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

func getClient() *clientset.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", "./config")
	if err != nil {
		fmt.Printf("Error building kubeconfig: %s\n", err.Error())
	}

	client, err := clientset.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, creating sample clientset\n", err.Error())
	}

	AddToScheme(scheme.Scheme)
	return client
}

func taskWithRuns(client *clientSetType) {
	createTask(client, getExampleTask())
	createTaskRun(client, getExampleTaskRun())
	listTasks(client)
	listTaskRuns(client)
	deleteTaskRun(client, "hello-task-run")
	deleteTask(client, "hello")
}

func pipelineWithRuns(client *clientSetType) {
	createTask(client, getExampleTask())
	createTask(client, getExampleTask2())
	createPipeline(client, getExamplePipeline())
	createPipelineRun(client, getExamplePipelineRun())
	listPipelines(client)
	listPipelineRuns(client)
	deletePipelineRun(client, "hello-goodbye-pipeline-run")
	deletePipeline(client, "hello-goodbye")
	deleteTask(client, "hello")
	deleteTask(client, "goodbye")
}

func customRun(client *clientSetType) {
	createCustomRun(client, getExampleCustomRun())
	listCustomRuns(client)
	deleteCustomRun(client, "example-run")
}

func clusterTask(client *clientSetType) {
	createClusterTask(client, getExampleClusterTask())
	listClusterTasks(client)
	deleteClusterTask(client, "example-cluster-task")
}
