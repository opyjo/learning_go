package main

import "fmt"

func main() {
	// myArray := [3]int{10, 20, 30}
	// mySlice := []int{10, 20, 30}

	// fmt.Println("Array:", myArray)
	// fmt.Println("Slice:", mySlice)
	// fmt.Println("Slice length:", len(mySlice))
	// fmt.Println("Slice capacity:", cap(mySlice))

	// mySlice = append(mySlice, 40)
	// fmt.Println("Slice after append:", mySlice)

	// partOfSlice := mySlice[1:3]
	// fmt.Println("Part of slice:", partOfSlice)


	

	// for i, value := range mySlice {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, value)
	// }

	arr:= []int{1, 2, 3}
	s:= arr[1:3]
	fmt.Println(s)
	fmt.Println(s[0])
	s[0] = 20
	fmt.Println(arr)
	fmt.Println(s)
}
