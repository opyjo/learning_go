package main

import "fmt"



func main() {
  originalArray := [3]int{1, 2, 3}
  newArray := modifyArray(originalArray)
    fmt.Println( originalArray)
    fmt.Println(newArray)

  originalSlice := []int{1, 2, 3}
  newSlice := modifySlice(originalSlice)
  fmt.Println(originalSlice)
  fmt.Println(newSlice)

}


func modifyArray(arr [3]int) [3]int {
  arr[0] = 88
  return arr
}

func modifySlice(theSlice []int) []int {
    theSlice[0] = 88
    return theSlice
}
