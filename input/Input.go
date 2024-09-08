package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
    fmt.Println("HI My Name is abdullah");
    fmt.Println("Enter a your age");
	input:=bufio.NewReader(os.Stdin);
    age,_:=input.ReadString('\n');
	// fmt.Printf("Your age is %v %T",age,age);
    change,err :=strconv.ParseFloat(strings.TrimSpace(age),64)
	if err != nil {
	  fmt.Println("This is  not a number",err);
	}else{
     fmt.Printf("This is a number %v ,%T",change,change);
	}

}