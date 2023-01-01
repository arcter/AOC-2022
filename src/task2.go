package main

func lookup(c byte) int {
	if c < 'X' {
		return int(c-'A') + 1
	}
	return int(c-'X') + 1
}

func match(Line string) int {
	task2_1 := 0
	var oponent, player = lookup(Line[0]), lookup(Line[2])
	//1 rock 2 paper 3 scisor
	//1 2
	//2 3
	//3 1
	if (player-oponent == 1) || (oponent-player == 2) {
		task2_1 = player + 6
	} else if player == oponent {
		task2_1 = player + 3
	} else {
		task2_1 = player
	}
	return task2_1
}
func match2(Line string) int {
	task2_2 := 0
	var oponent, outcome = lookup(Line[0]), lookup(Line[2])
	if outcome == 2 {
		task2_2 = oponent + 3
	}
	if outcome == 1 {
		//1 3
		//2 1
		//3 2
		switch oponent {
		case 1:
			task2_2 = 3
		case 2:
			task2_2 = 1
		case 3:
			task2_2 = 2

		}
	}
	if outcome == 3 {
		task2_2 = oponent%3 + 1 + 6
	}
	return task2_2
}

func task2(file_content []string) {
	var task2_1 int = 0
	var task2_2 int = 0
	for _, line := range file_content {
		task2_1 += match(line)
		task2_2 += match2(line)
	}
	println("Task2 results:")
	println(task2_1)
	println(task2_2)
}
