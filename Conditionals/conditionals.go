package main

import "fmt"

func main() {
	var name string = "Charles"
	var age int = 32

	if age > 18 {
		fmt.Println(name, "you are an old man")
	} else if age < 50 {
		fmt.Println("More or less young man")
	} else {
		fmt.Println("You really are an old man")
	}

    for i := 0; i < 10; i++ {
        fmt.Println("The number is: ", i)
    }

	dayOfWeek := 6

	switch dayOfWeek {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Thursday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Tuesday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Sartuday")
	default:
		fmt.Println("Wrong day")
	}
}
