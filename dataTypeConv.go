package main

import "fmt"

func main() {
	var (
		num1 int8    = 10
		num2 float32 = 20.2
		name = "Yehezkiel David Setiawan"
	)

	num3 := int64(num1)
	fmt.Println(num3)
	num4 := float64(num2)
	num5 := float64(num3)
	fmt.Println(num4+num5)

	bitName := name[0]
	fmt.Println(bitName)
	StrName := string(bitName)
	fmt.Println(StrName)
}