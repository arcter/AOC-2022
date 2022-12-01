package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("task1-1.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var max3, max, current int = 0, 0, 0
	var carried = []int{}
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		if line == "" {
			carried = append(carried, current)
			current = 0
		} else {
			var intVar int
			intVar, err = strconv.Atoi(line)
			current = current + intVar
		}
	}
	sort.Ints(carried)
	max = carried[len(carried)-1]
	max3 = carried[len(carried)-1] + carried[len(carried)-2] + carried[len(carried)-3]
	fmt.Println(max)
	fmt.Println(max3)
	readFile.Close()
}
