package main

import ( "fmt" ; "math" ; "math/rand")

func main()  {
    fmt.Println("In Go the 'main()' always will run")
    maths()
}

func maths()  {
    fmt.Println("Let´s do some math using the 'math' lib,")
    fmt.Println("The square root of 4 is", math.Sqrt(4))
    fmt.Println("The sum betwen 3 and 2 is", (3+2))
    fmt.Println("A number from 1-100 is",rand.Intn(100))
}
