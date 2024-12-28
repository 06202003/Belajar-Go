package main

import "fmt"

func main() {
	type NRP int32

	var nrpDavid NRP = 2172003
	fmt.Println(nrpDavid)

	var onto int64 = 2172028
	var nrpOnto NRP = NRP(onto)
	fmt.Println(nrpOnto)
}