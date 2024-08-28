package main

import "fmt"

func main() {
    show()
	fun(5, 10)
}

func show() {
	fmt.Print("Hi my name is Abdullah")
}

func fun(x int,y int)  {
 fmt.Print(x+y)
}