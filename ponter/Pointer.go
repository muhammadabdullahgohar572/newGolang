package main

import "fmt"

type Employee struct{
	EmployeeId int
}
func main() {

    

	// var a int = 10
	// var b *int = &a
	// fmt.Print(a,"\n")
	// fmt.Print(b)

     var a Employee;
   fmt.Print(a)
   fmt.Print("\n")
   var b *Employee;
   fmt.Print(b)
   fmt.Print("\n")

    new :=new(Employee)
   fmt.Print("value of new",new)
   fmt.Print("\n")
   fmt.Print("value of new *",*new)
  


}