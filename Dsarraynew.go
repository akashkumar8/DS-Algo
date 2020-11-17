package main

import "fmt"


var integerArray [8]int

var stringArray [5]string

func main() {

	integerArray[0] = 10 
	integerArray[1] = 20
	integerArray[2] = 30
	integerArray[3] = 40
	integerArray[4] = 50
    integerArray[5] = 60
    integerArray[6] = 70
    integerArray[7] = 80


	fmt.Println("This is the integer Array: ", integerArray)

	stringArray[0] = "one"
	stringArray[1] = "two"
	stringArray[2] = "three"
	stringArray[3] = "four"
    stringArray[4] = "five"

	fmt.Println("This is the string array: ", stringArray)

}