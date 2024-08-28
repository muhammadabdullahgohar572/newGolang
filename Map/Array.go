package main

import "fmt"

func main() {

	var arr []int = []int{1, 2, 3, 4};

	  arr2 := append(arr,5,6,7)
	  arr3 :=append(arr,arr2...);
	  fmt.Println(arr3); // Output: [1, 2, 3, 4, 5, 6, 7]	
	

	
}