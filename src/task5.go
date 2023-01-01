package main

import (
	"strings"
)

func task5_1(file_content []string) string {
	var res = ""
	var stacks [9]string
	for _, line := range file_content {
		if strings.Contains(line, "[") {
			var oi = 0
			for i := 2; i < len(line); i += 4 {
				if line[i] != ' ' {
					stacks[oi] = stacks[oi] + string(line[i])

				}
				oi++
			}
		}
	}
	for _, stack := range stacks {
		res += string(stack[len(stack)-1])
	}
	return res
}

func task5(file_content []string) {
	println("Task5 results:")
	println(task5_1(file_content))
	//println(task5_2(file_content))

}
