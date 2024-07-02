package main

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getExampleTask() *v1beta1.Task {
	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{
			Kind:       "Task",
			APIVersion: "tekton.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "hello",
		},
		Spec: v1beta1.TaskSpec{
			Steps: []v1beta1.Step{
				{
					Name:   "echo",
					Image:  "alpine",
					Script: "#!/bin/sh\n" + "echo \"Hello World\"",
				},
			},
		},
	}

	return task
}

func getExampleTaskRun() *v1beta1.TaskRun {
	taskrun := &v1beta1.TaskRun{
		TypeMeta: v1.TypeMeta{
			Kind:       "TaskRun",
			APIVersion: "tekton.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "hello-task-run",
		},
		Spec: v1beta1.TaskRunSpec{
			TaskRef: &v1beta1.TaskRef{
				Name: "hello",
			},
		},
	}

	return taskrun
}

func getExampleTask2() *v1beta1.Task {
	task := &v1beta1.Task{
		TypeMeta: v1.TypeMeta{
			Kind:       "Task",
			APIVersion: "tekton.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "goodbye",
		},
		Spec: v1beta1.TaskSpec{
			Params: v1beta1.ParamSpecs{
				{Name: "username", Type: "string"},
			},
			Steps: []v1beta1.Step{
				{
					Name:   "goodbye",
					Image:  "ubuntu",
					Script: "#!/bin/bash\n" + "echo \"Goodbye $(params.username)!\"",
				},
			},
		},
	}

	return task
}

func getExamplePipeline() *v1beta1.Pipeline {
	pipeline := &v1beta1.Pipeline{
		TypeMeta: v1.TypeMeta{
			Kind:       "Pipeline",
			APIVersion: "tekton.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "hello-goodbye",
		},
		Spec: v1beta1.PipelineSpec{
			Params: v1beta1.ParamSpecs{
				{
					Name: "username",
					Type: "string",
				},
			},
			Tasks: []v1beta1.PipelineTask{
				{
					Name: "hello",
					TaskRef: &v1beta1.TaskRef{
						Name: "hello",
					},
				},
				{
					Name:     "goodbye",
					RunAfter: []string{"hello"},
					TaskRef:  &v1beta1.TaskRef{Name: "goodbye"},
					Params: v1beta1.Params{
						{
							Name: "username",
							Value: v1beta1.ParamValue{
								Type:      v1beta1.ParamTypeString,
								StringVal: "$(params.username)",
							},
						},
					},
				},
			},
		},
	}

	return pipeline
}

func getExamplePipelineRun() *v1beta1.PipelineRun {
	pr := &v1beta1.PipelineRun{
		TypeMeta: v1.TypeMeta{
			Kind:       "PipelineRun",
			APIVersion: "tekton.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "hello-goodbye-pipeline-run",
		},
		Spec: v1beta1.PipelineRunSpec{
			PipelineRef: &v1beta1.PipelineRef{
				Name: "hello-goodbye",
			},
			Params: v1beta1.Params{
				{
					Name: "username",
					Value: v1beta1.ParamValue{
						Type:      v1beta1.ParamTypeString,
						StringVal: "Tekton",
					},
				},
			},
		},
	}

	return pr
}

func getExampleClusterTask() *v1beta1.ClusterTask {
	cr := &v1beta1.ClusterTask{
		TypeMeta: v1.TypeMeta{
			Kind:       "ClusterTask",
			APIVersion: "tekton.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "example-cluster-task",
		},
		Spec: v1beta1.TaskSpec{
			Steps: []v1beta1.Step{
				{
					Name:   "step1",
					Image:  "alpine",
					Script: "#!/bin/sh\necho \"This is step 1\"",
				},
				{
					Name:   "step2",
					Image:  "alpine",
					Script: "#!/bin/sh\necho \"This is step 2\"",
				},
			},
		},
	}

	return cr
}

func getExampleCustomRun() *v1beta1.CustomRun {
	cr := &v1beta1.CustomRun{
		TypeMeta: v1.TypeMeta{
			Kind:       "CustomRun",
			APIVersion: "tekton.dev/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name: "example-run",
		},
		Spec: v1beta1.CustomRunSpec{
			CustomRef: &v1beta1.TaskRef{
				APIVersion: "example.dev/v1beta1",
				Kind:       "Example",
				Name:       "an-existing-example-task",
			},
		},
	}

	return cr
}
