package main

import (
	"github.com/modmuss50/Advent-Of-Code-2018/utils"
	"strconv"
	"fmt"
)

func main(){

	lines := utils.ReadLinesFromFile("day1/input.txt")

	frequency := 0

	for _ , str := range lines {
		value, _ := strconv.Atoi(str)
		frequency += value
	}

	fmt.Println("Part 1:")
	fmt.Println(frequency)

	//Part 2 starts here
	frequency = 0
	var visited []int
	line := 0

	for true {
		value, _ := strconv.Atoi(lines[line])
		frequency += value
		if contains(visited, frequency){
			break
		}
		visited = append(visited, frequency)
		if line + 1 == len(lines) {
			line = 0
		} else {
			line ++
		}
	}

	fmt.Println("Part 2:")
	fmt.Println(frequency)

}


func contains(arr []int, val int) bool {
	for _, ent := range arr {
		if ent == val {
			return true
		}
	}
	return false
}