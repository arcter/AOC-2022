package main

import "strings"

func char_to_int(charF byte) int {
	if charF > 'Z' {
		return int(charF) - 'a' + 1
	}
	return int(charF) - 'A' + 27

}

func duplicat_find(Line string) int {
	var res = 0
	var len_of_half = len(Line) / 2
	var front, back = uniqe(Line[0:len_of_half]), uniqe(Line[len_of_half:])
	for _, charF := range front {
		for _, charB := range back {
			if charF == charB {
				return char_to_int(byte(charF))
			}
		}
	}
	return res
}

func task3_1(file_content []string) int {
	var task3_1 = 0
	for _, line := range file_content {
		task3_1 += duplicat_find(line)
	}
	return task3_1
}

func uniqe(line string) string {
	var bytes = []byte(line)
	inResult := make(map[byte]bool)
	var result []byte
	for _, str := range bytes {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return string(result)
}

func task3_2(file_content []string) int {
	var task3_2 = 0
	for i := 0; i < len(file_content); i += 3 {
		var line1 = uniqe(file_content[i])
		var line2 = uniqe(file_content[i+1])
		var line3 = uniqe(file_content[i+2])
		for _, char := range line1 {

			if strings.Contains(line2, string(char)) && strings.Contains(line3, string(char)) {
				task3_2 += char_to_int(byte(char))
			}
		}
	}
	return task3_2
}

func task3(file_content []string) {
	println("Task3 results:")
	println(task3_1(file_content))
	println(task3_2(file_content))

}
