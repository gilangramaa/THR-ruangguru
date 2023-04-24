package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	//yout code here!
	if position == "left" {
		return arr[1:]
	} else if position == "right" {
		return arr[:len(arr)-1]
	} else {
		return arr
	}
}

func main() {
	arr1 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(removeLeftRight(arr1, "left")) // Output: [2 3 4 5]

	arr2 := []interface{}{1, 2, 3, 4, 5}
	fmt.Println(removeLeftRight(arr2, "right")) // Output: [1 2 3 4]

	arr3 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeLeftRight(arr3, "left")) // Output: [Budi Joko Tono]

	arr4 := []interface{}{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeLeftRight(arr4, "right")) // Output: [Edo Budi Joko]
}
