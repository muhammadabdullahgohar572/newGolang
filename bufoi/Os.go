package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// "bufio"
// "fmt"
// "os"
// "strconv"
// "strings"



func main() {

	// fmt.Printf("HI my name is Abdullah .How Are  you ")
	// fmt.Printf("\n")
	// rander :=bufio.NewReader(os.Stdin);
    // fmt.Printf("Enter your age ");
       
	// age,_:=rander.ReadString('\n');
	// fmt.Printf("Your age is and %T %v",age)


// ---------------------------------------------------------------------------
    

     // fmt.Print("HI my name is Abdullah.Hoe are you");
		// fmt.Print("\n");
		// render :=bufio.NewReader(os.Stdin);
		// fmt.Printf("Enter your age ");
		// name,_:=render.ReadString('\n')
		// fmt.Printf("Your name is %s ,%T",name,name)

		//  fmt.Print("HI my name is Abdullah.Hoe are you");
		// fmt.Print("\n");
		// render :=bufio.NewReader(os.Stdin);
		// fmt.Printf("Enter your age ");
		// name,_:=render.ReadString('\n')
		// // fmt.Printf("Your name is %s ,%T",name,name)
       
		// change,err :=strconv.ParseFloat(strings.TrimSpace(name),64)
		// if err!=nil{
        //     fmt.Println("Error occurred while converting string to float")
        // }else{
        //     fmt.Printf("Converted age is%v",change+1,)
        //     fmt.Printf("%T",change)
		// 	}
          

          fmt.Print("HI MY NAME IS ABDULLAH");
         fmt.Print("\n");

		 render:=bufio.NewReader(os.Stdin);
		 fmt.Printf("Enter your age ");
		 age,_ :=render.ReadString('\n');
		 fmt.Printf("This is your Age and %v %T",age,age)

		 change,err :=strconv.ParseFloat(strings.TrimSpace(age),64)
		 if err !=nil {
			fmt.Print("Error occurred while converting string to float");
		 }else {
			fmt.Printf("Converted age is %v",change+1);
		 }
}