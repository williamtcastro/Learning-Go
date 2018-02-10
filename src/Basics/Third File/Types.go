package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	fmt.Println("In GO we have a bunch of a Types : ")
	fmt.Println("Bool, Int, Float32, Float64")
	fmt.Println("In go to declare a variable you need to write 'var name type' var(indicates that is indeed a variable) , name (the real name of the var), type (the type of the var that you are creating)")

	var num1 float64 = 5.6
	var num2 float64 = 9.5
	//var num1,num2 float64 = 5.6, 9.5
	//num1,num2 := 5.6, 9.5 this way the vars are theter with float64
	fmt.Println("The sum of 5.6 and 9.5 is", add(num1, num2))

	w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))

}
