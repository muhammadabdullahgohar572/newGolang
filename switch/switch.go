package main

import "fmt"

func main() {
	// a := 1

	// switch a {
	// case 2:
	// 	fmt.Print("This is value Of ",a)
	// case 3:
	// 	fmt.Print("This is value Of ",a)
	// case 4:
	// fmt.Print("This is value Of ",a)
	// case 1:
	// 	fmt.Print("This is value Of ",a)
	// 	break
	// 	fmt.Print("Value is None")
	// } 
	
	// a :=1;

	// switch a  {
	// case 1,2,3,4:
	// 	fmt.Print("This is value Of ",a,"It is a Positive Value ")
	// 	break;
	// 	case -5,-6,-7,-8:
	// 		fmt.Print("This is value Of ",a,"It is a Negative Value ")
    //     break;
	// 	case 5,7,11:
	// 		fmt.Print("This is value Of ",a,"It is a Prime Number ")
    //     break;
       
	// }
      

	var a interface{} = 12.4

	switch a.(type) {
	case string:
		fmt.Printf("The value is a string, its value is: %s\n", a.(string))
	case int:
		fmt.Printf("The value is an integer, its value is: %d\n", a.(int))
	case float64:
		fmt.Printf("The value is a float64, its value is: %f\n", a.(float64))
	default:
		fmt.Printf("The value is of unknown type\n")
	}
}