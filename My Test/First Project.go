package main

import "fmt"

func main()  {
	
	fmt.Println(add(1,2))
}

func add(num1 int,num2 int) (int, int)  {
	
	x := 1
	num1 += x 
	num2 += x

	return num1, num2
}

//returning two variables with one function