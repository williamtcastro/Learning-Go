package main

import ( "fmt" ; "math" ; "math/rand")

func main()  {
    fmt.Printf("In Go the 'main()' always will run")
    maths()
}

func maths()  {
    fmt.Printf("Let´s do some math using the 'math' lib,\n")
    fmt.Printf("The square root of 4 is", math.Sqrt(4),"\n")
    fmt.Printf("The sum betwen 3 and 2 is", (3+2), "\n")
    fmt.Printf("A number from 1-100 is",rand.Intn(100), "\n")
}
