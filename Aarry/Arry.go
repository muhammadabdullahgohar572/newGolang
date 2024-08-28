package main

import "fmt"

func main() {

	// Declare a variable with initial value of 10.

	// var a = [...]int{1, 2, 3, 5, 46, 8, 4, 7}
	// Declare a slice that references the array.

	// Print the slice.
	// fmt.Println(a);

	var b =[...]int{1,2,3,4,56,7,8,5,8,13,6}
	// Print the slice.
	arr := b;
	b[2]=20;
	fmt.Println(arr);
	
	var arr2 [3][3]int = [3][3]int{{1,2,3},{1,2,3},{1,2,3}}
	// Print the slice.
	fmt.Println(arr2);

	var age int = 25;
	var pi float64 = 3.14159;
	var name string = "John Doe";
	var isGoFun bool = true
	// var myMap map[string]int = map[string]int{"one":1, "two":2, "three":3}
	fmt.Printf("%d\n%f\n%s\n%t\n", age, pi, name, isGoFun)
// %d is the format specifier for an integer.
// %f is for a floating-point number.
// %s is for a string.
// %t is for a boolean.
}