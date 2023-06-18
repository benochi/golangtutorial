package main

import "fmt"

func main() {

	var nameOne string = "I'm a string"
	var nameTwo = "I'm also a string."
	var nameThree string
	nameThree = "Goat"
	fmt.Println("Strings begin.")
	fmt.Println(nameOne, " + ", nameTwo, nameThree)

	nameOne = "Billy"
	fmt.Println(nameOne, " + ", nameTwo, nameThree)

	nameFour := "Yoshi"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
	fmt.Println("Strings end.")

	fmt.Println("Int begin.")

	var ageOne int = 10
	var ageTwo = 20
	ageThree := 50
	fmt.Println(ageOne, ageTwo, ageThree)

	//bits and memory declaring specific bits ofr an int.
	var numOne int8 = 25 //range -128 - +127
	var numTwo int8 = -128
	var numThree uint8 = 25 //unsigned must be positive but can go up to 255.

	fmt.Println(numOne, numTwo, numThree)
	fmt.Println("Int end.")
	//float64 has more precision, recommended. also default of floats.
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 100.10
	scoreThree := 75.1
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
