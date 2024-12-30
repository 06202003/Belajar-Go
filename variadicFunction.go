package main

import "fmt"

func sumAll(num ...int) (sum int) {
	sum = 0
	for _,number := range num{
		sum += number
	}
	return
}

func main(){
	fmt.Println(sumAll(1,2,3,4,5,6,7,8,9))
	numbers := []int{2,4,6,8,}
	total:= sumAll(numbers...)
	fmt.Println(total)
}