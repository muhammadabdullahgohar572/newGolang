package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		// fmt.Print("The value of ",i,"\n")
	}

	for j :=1; j <= 10; j++{
        
	     if j%2 ==0 {
			fmt.Print("This Number is Even ",j ,"\n");
		 }	else {
				fmt.Print("This Number is Odd ",j ,"\n");
			}
		 
	}

	
}
