package main

import "fmt"

 type student struct{
	name string
	age int
    city string

 }
func mainn() {
	fmt.Print("HI ");

	// ma :=map[any]any{
	// 	"name":"sachin",
    //     "age":23,
    //     "city":"pune",
	// }
	// fmt.Print(ma)

	// show(10,20)

    // var a =[...]int{1,2,3,4,5,6,7,8,9}
    //  var B=[...]int{10,20,30,40,50,60,70,80}
	//  c :=append(a[:],B[:]...)
	//  fmt.Print(c[:5])

       neww :=map[any]student{
         1:{"sachin",23,"pune"},
		 2:{"sachin",23,"pune"},
         3:{"sachin",23,"pune"},
         4:{"sachin",23,"pune"},
         5:{"sachin",23,"pune"},
         6:{"sachin",23,"pune"},
	   }
	   fmt.Print(neww)
}


// func  show(a int  ,b int ) {
// 	fmt.Print(a+b)
// }