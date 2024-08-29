package main

import "fmt"

func main() {
	

	 a :=map[string]any{
		"name":"sai",
         "age":23,
         "city":"hyderabad",
	 }
	 fmt.Println(a)
	 fmt.Print("\n");
     _,ok :=a["id"]
	  fmt.Println(ok) 
}