package main

import (
	"fmt"
	"sort"
	"strconv"
)

func task1(file_content []string) {
	var max3, max, current int = 0, 0, 0
	var carried = []int{}
	for _, line := range file_content {
		if line == "" {
			carried = append(carried, current)
			current = 0
		} else {
			var intVar int
			intVar, _ = strconv.Atoi(line)
			current = current + intVar
		}
	}
	sort.Ints(carried)
	max = carried[len(carried)-1]
	max3 = carried[len(carried)-1] + carried[len(carried)-2] + carried[len(carried)-3]
	fmt.Println(max)
	fmt.Println(max3)
}
