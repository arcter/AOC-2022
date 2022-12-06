package main

import (
	"regexp"
	"strconv"
)

func in_range(RL int, RR int, P int) bool {
	return RL <= P && P <= RR
}

func parse_numline(line string) (int, int, int, int) {
	reg := regexp.MustCompile("[\\,\\-]+")
	var numbers = reg.Split(line, -1)
	var R1L, _ = strconv.Atoi(numbers[0])
	var R1R, _ = strconv.Atoi(numbers[1])
	var R2L, _ = strconv.Atoi(numbers[2])
	var R2R, _ = strconv.Atoi(numbers[3])
	return R1L, R1R, R2L, R2R

}

func task4_1(file_content []string) int {
	var res = 0
	for _, line := range file_content {
		var R1L, R1R, R2L, R2R = parse_numline(line)
		if (R1L <= R2L && R2R <= R1R) || (R1L >= R2L && R2R >= R1R) {
			res++
		}
	}
	return res
}

func task4_2(file_content []string) int {
	var res = 0
	for _, line := range file_content {
		var R1L, R1R, R2L, R2R = parse_numline(line)
		if in_range(R1L, R1R, R2L) || in_range(R1L, R1R, R2R) || in_range(R2L, R2R, R1L) || in_range(R2L, R2R, R1R) {
			res++
		}
	}
	return res
}

func task4(file_content []string) {
	println("Task4 results:")
	println(task4_1(file_content))
	println(task4_2(file_content))

}
