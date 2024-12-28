package main

import "fmt"

func main() {
	const fName = "Yehezkiel"
	const lName = "Setiawan"
	const (
		mName = "David"
		nrp = 2172003
	)

	// Error
	// fName = "Laurentius"
	// lName = "Yudha"
	
	fmt.Println(fName, mName, lName, "-", nrp)
}