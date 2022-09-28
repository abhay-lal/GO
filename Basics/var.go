package main

import "fmt" // package from go standard lib - used for formatting strings and printing messages to console , has bunch of methods
func main() {
	//strings
	var nameOne string = "Tony"
	var nameTwo = "Loki"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Ironman"
	nameThree = "Hulk"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Thor" // automatically treats := as string , shorthand version ..Can not be used outside function

	fmt.Println(nameFour)

	var ageOne int = 40
	var ageTwo = 18
	ageThree := 51

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits and memory
	// var num1 int8 = 25

	// var num2 int8 = -128

	// var num3 uint8 = 30

	var myscore float32 = 19.21
	var yourscore float64 = 340282038408.23
	whosescore := 4.3

	fmt.Println(myscore, yourscore, whosescore)

}
