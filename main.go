package main

import (
	"fmt"
)

const PORT = 999

func main() {
	strTypeTest := "we are test type of string elements"

	// for k, v = range strTypeTest{
	// 	fmt.Printf("Key: %v,  Value: %v, Type: %T", k, v, v)
	// }
	var a, s, d = 0, "", 0
	q, w, e := "asd", 0.5, 1

	fmt.Println(PORT)
	fmt.Println(a, s, d, q, w, e)

	for k := 0; k < len(strTypeTest); k++ {
		v := strTypeTest[k]
		s := string(v)
		fmt.Printf("Key: %v,  Value: %v,  Char Value: %v, Type: %T \n", k, v, s, v)
	}
	fmt.Println(strTypeTest)
}
