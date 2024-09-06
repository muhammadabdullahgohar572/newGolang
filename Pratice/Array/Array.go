package main

import "fmt"

type name struct{
  name  string
  age int 
  city string
}
func main() {

	// allname:=map[any]name{
    //     1:{"sai",23,"hyderabad"},
	// 	2:{"sachin",23,"pune"},
    //     3:{"abdullah",24,"karachi"},
    //     4:{"sai",23,"hyderabad"},
	// }
	
	// for index,value :=range allname{
	// 	fmt.Println(index,value)
	// }

	//  a :=map[string]any{
	// 	"name":"sai",
    //      "age":23,
    //      "city":"hyderabad",
	//  }
	//  fmt.Println(a)
	//  fmt.Print("\n");
    //  _,ok :=a["id"]
	//   fmt.Println(ok) 

	//   b :=map[string]any{
	// 	"name":"Abdullah",
	// 	"age" :20,
	// 	"city":"hyderabad",
	//   }
	//   new,_ :=b["name"]
	//   fmt.Println(new)


	add :=1;

	for add < 10{ 
		if add == 3{
         goto logo;
        }
	if add ==5{
       add++;
	   continue;
	}
		fmt.Println(add)
		add++

	}
	logo:
	fmt.Println("hello world")
}