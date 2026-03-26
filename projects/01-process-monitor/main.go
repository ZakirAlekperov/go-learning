package main

import "fmt"

func main() {
	processName := []string{"ProcessA", "ProcessB", "ProcessC", "ProcessD", "ProcessE", "ProcessF", "ProcessG", "ProcessH", "ProcessI", "ProcessJ"}

	pids := []int{1234, 5678, 9101, 1121, 3141, 5161, 7181, 9202, 1222, 3242}

	status := []string{"Running", "Sleeping", "Stopped", "Zombie", "Running", "Sleeping", "Stopped", "Zombie", "Running", "Sleeping"}

	mbytes := []int{100, 200, 150, 300, 250, 400, 350, 450, 500, 550}

	pidsZombie := []int{}

	runningProcesses := 0
	sleepingProcesses := 0
	stoppedProcesses := 0
	zombieProcesses := 0

	for i, name := range status {
		if name == "Running" {
			runningProcesses++
		} else if name == "Sleeping" {
			sleepingProcesses++
		} else if name == "Stopped" {
			stoppedProcesses++
		} else if name == "Zombie" {
			zombieProcesses++
			pidsZombie = append(pidsZombie, pids[i])
		}
	}
	fmt.Printf("sleepingProcesses: %d\n", sleepingProcesses)
	fmt.Printf("runningProcesses: %d\n", runningProcesses)
	fmt.Printf("stoppedProcesses: %d\n", stoppedProcesses)
	fmt.Printf("zombieProcesses: %d\n", zombieProcesses)
	if zombieProcesses > 0 {
		fmt.Println("Warning: There are zombie processes running!", pidsZombie)
	}

	procesMemory := 0
	processIndex := 0
	for i, size := range mbytes {
		if size >= procesMemory {
			processIndex = i
			procesMemory = size
		}
	}

	fmt.Printf("maxMemoryProcess: %s, Memory: %dMB\n", processName[processIndex], procesMemory)

}
