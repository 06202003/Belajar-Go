package main

import "fmt"

func helloWorld(){
	fmt.Println("Hello World")
}

func paramTest(num1 int,num2 int){
	fmt.Println(num1+num2)
}

func returnTest(num int) (y,x int) {
	x = num%10
	y = (num/10)*10
	return
}

func main(){
	helloWorld()
	fmt.Println("==================================================")
	a:=1
	b:=2
	paramTest(a,b)
	fmt.Println("==================================================")
	var num1 int
	fmt.Print("Angka puluhan: ")
	fmt.Scan(&num1)
	fmt.Println(returnTest(num1))
}