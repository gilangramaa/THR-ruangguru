package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		new_arr := append([]int{newData}, data...)
		return new_arr
	} else if position == "down" {
		new_arr := append(data, newData)
		return new_arr
	} else {
		return data
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	newData := 6
	position := "up"

	result := AddElement(data, newData, position)
	fmt.Println(result) // [6 1 2 3 4 5]

	position = "down"

	result = AddElement(data, newData, position)
	fmt.Println(result)
}
