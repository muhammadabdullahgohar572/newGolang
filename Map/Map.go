package main

import "fmt"

func mainn() {
	ma := map[string]any{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	 
	// delete(ma, "two")
	// fmt.Println( len(ma))
	tes,ok :=ma["six"]
	_,seven :=ma["seven"]

	fmt.Print(ma,"\n")
	fmt.Print(tes,ok,"\n")
	fmt.Print(seven)

}