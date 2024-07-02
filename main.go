package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	client := getClient()

	for {
		fmt.Println("Select demo number (1-4) or 0 to exit:")
		fmt.Println("1. Task & TaskRun Demo")
		fmt.Println("2. Pipeline & PipelineRuns Demo")
		fmt.Println("3. CustomRuns")
		fmt.Println("4. ClusterTasks")

		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch option {
		case 0:
			fmt.Println("Exiting...")
			return
		case 1:
			taskWithRuns(client)
		case 2:
			pipelineWithRuns(client)
		case 3:
			customRun(client)
		case 4:
			clusterTask(client)
		default:
			fmt.Println("Invalid Option.")
		}
	}
}
