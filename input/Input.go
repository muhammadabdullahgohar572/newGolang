package main

import "fmt"

// "bufio"
// "fmt"
// "os"
// "strconv"
// "strings"


func main() {
    // fmt.Println("HI My Name is abdullah");
    // fmt.Println("Enter a your age");
	// input:=bufio.NewReader(os.Stdin);
    // age,_:=input.ReadString('\n');
	// // fmt.Printf("Your age is %v %T",age,age);
    // change,err :=strconv.ParseFloat(strings.TrimSpace(age),64)
	// if err != nil {
	//   fmt.Println("This is  not a number",err);
	// }else{
    //  fmt.Printf("This is a number %v ,%T",change,change);
	// }
     
	userdata :=map[any]Name{
		"abdullah":{
            name:"abdullah",
            emil:"abdullah@gmail.com",
            age:20,
        },
        "ali":{
            name:"ali",
            emil:"ali@gmail.com",
            age:20,
        },
        "ahmad":{
            name:"ahmad",
            emil:"ahmad@gmail.com",
            age:20,
        },
	}

	for key, value := range userdata {
		fmt.Printf("Key: %v\n", key)
		fmt.Printf("Name: %v\n", value.name)
		fmt.Printf("Email: %v\n", value.emil)
		fmt.Printf("Age: %v\n", value.age)
		fmt.Println() // To print a new line between entries
	}
}

type Name struct{
	name string
    emil string
	age int
}