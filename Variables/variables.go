package main

import "fmt"
import "reflect"

func main() {
    var name string = "Charles"
    var age int = 32
    fmt.Println("Hello, my name is ", name, " I'm ", age);
        
    var version float32 = 1.0
    fmt.Println("My first software, version: ", version);
    
    var year = 2020
    fmt.Println("Type of variable is: ", reflect.TypeOf(year))
    
    city := "Vitória"
    fmt.Println("Type of variable is: ", reflect.TypeOf(city))

    fmt.Println("Which is Day of week today?")

    var dayOfWeek string
    fmt.Scan(&dayOfWeek)

    fmt.Println("Today is: ", dayOfWeek)
    fmt.Println("The pointer of variable is: ", &dayOfWeek)
}