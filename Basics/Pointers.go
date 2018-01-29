package main

import "fmt"

func main()  {
    x := 15
    a := &x //memory addres , when & is used the var will point the memory addres that is being used
    fmt.Println(a) //reference the memory addres
    fmt.Println(*a) //print the value in the memory , when * is used will point the value that is contained in the memory
    *a = 5 //reference the value and change it to the value of 5
    fmt.Println(x) //print the new reference value of x
}
