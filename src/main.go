package main

import (
	"bufio"
	"log"
	"os"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	var task1_file, err = readLines("../task1-1.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	task1(task1_file)
	var task2_file, err2 = readLines("../task2-1.txt")
	if err2 != nil {
		log.Fatalf("readLines: %s", err2)
	}
	task2(task2_file)
}
