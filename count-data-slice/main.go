package main

import "fmt"

func howManyElements(data []any) int {
	//your code here!

	return len(data)
}

func main() {
	input1 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(howManyElements(input1)) // output: 5

	input2 := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(howManyElements(input2)) // output: 10

	input3 := []interface{}{1, 1, 1, 5, 5, 5}
	fmt.Println(howManyElements(input3)) // output: 6

	input4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(input4)) // output: 4

	input5 := []interface{}{"Edo", "Budi", "Joko", "Tono", "Edo", "Budi", "Joko", "Tono"}
	fmt.Println(howManyElements(input5)) // output: 8

	input6 := []interface{}{true, false, true, false, true, false}
	fmt.Println(howManyElements(input6)) // output: 6
}
