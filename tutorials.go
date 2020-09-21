package main

import (
	"fmt"
)

// we need "fmt" module and main package in almost all programs,so just import it

// main() is the entry point of the go program ,it will execute automatically,we donot need to call it
// explicitely like java public static void main() method
func main() {

	// var then name of Variable then  Type of variable = value that u want to assign to it
	// for variable refer this https://www.tutorialspoint.com/go/go_data_types.htm
	// IDENTIFIER NAMING CONVENSION
	// https://medium.com/better-programming/naming-conventions-in-go-short-but-descriptive-1fa7c6d2f32a#:~:text=The%20convention%20in%20Go%20is,first%20character%20should%20be%20uppercase.
	/**
	The convention in Go is to use MixedCaps or mixedCaps (simply camelCase) rather
	than underscores to write multi-word names.
	If an identifier needs to be visible outside the package,
	 its first character should be uppercase.
	 If you don’t have the intention to use it in another package,
	 you can safely stick to mixedCaps.
	*/

	var name string = "Rajat"
	var years uint8 = 25

	// if u see above we are explicitely we are defining the type of variable before we are using it
	// but we can skip this thing also , go determine this thing that which type of variable it is
	// like python
	var favFood = "ladyfinger"

	// This is also a another way of defining variable ,here we do not need to write var keyword also,
	// BUT WE CAN ASSIGN ANOTHER TYPE OF VALUE TO VARIABLE ONCE WE DEFINE STRING STRING WE CAN NOT
	// ASSIGN INTERGER OR ANY OTHER ,OTHER THAN STRING,LET IT BE ANY WAY U ARE ARE USING TO DEFINE
	// VARIABLES,ONE VARIABLE CAN HAVE ONE TYPE OF VALUE ONLY ,IT LIKE STATICALLY TYPES LANGUAGE
	// LIKE PYTHON WE CAN NOT ASSIGN ANOTHER KIND OF VALUE TO VARIABLE
	myFood := "Panir" //WALRUS OPERATOR

	// if u want to see type of variable then we can go for below Printf ,like in python we have type()
	// https://golang.org/pkg/fmt/
	// https://gobyexample.com/string-formatting
	fmt.Printf("%T ,%T", favFood, myFood)
	fmt.Printf("%v", favFood)

	fmt.Println("Hello saurabh in the World GO language", name, years, favFood)

	// now we learn ,how we can get user input

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type the birth year ,u were born")
	// scanner.Scan()
	// input, error := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Printf("You typed: %d,%q", input, error)

	// fmt.Printf("You will be %d year old at the end of 2020 december 31st", 2020-input)

	// if else

	z := 10

	if z == 10 {
		fmt.Printf("if  %d", z)
	} else if z > 10 {
		fmt.Printf("else if  %d", z)
	} else {
		fmt.Printf(" else %d ", z)
	}

	// for loop

	// y := 0

	// for y < 5 {
	// 	fmt.Println("for y value", y)
	// 	y++
	// }

	for x := 0; x <= 5; x++ {
		fmt.Println("for x value", x)
	}

	// we have also break to break the loop and
	// continue ,to continue means to skip current iteration and continue fron next one

	// Day 2 =====>
	// started with array

	var arr [5]int

	arr[1] = 10
	arr[2] = 11
	arr[3] = 12

	arrWithKnownValues := [4]int{0, 1, 2, 3}
	// [how many array inside][how many elements in each array]
	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("array-->", arr, len(arr))
	fmt.Println("array-->", arrWithKnownValues, arrWithKnownValues[3], len(arrWithKnownValues))
	fmt.Println("2Darray-->", arr2D, len(arr2D))

	// Slices
	// i just saw video for it

	// use of Range for iterating array elements

	iterArray := [5]int{11, 12, 13, 14, 15}

	for index, element := range iterArray {

		fmt.Printf("index : %d element : %d\n", index, element)
	}

	// map => nothing but dictionary n python

	// map[string] => means key will be of type string
	//  map[string]int => int means value will be of type integer
	var mp map[string]int = map[string]int{
		"apple":  4,
		"banana": 5,
		"pear":   6,
	}

	mp["apple"] = 10

	// To know particular key is exists in map or not ,then we do this below way

	// so val -> represent the value of the key if key exists else key would be 0 as we have
	// defined value of type int while defining map bove
	// and ok => represent true if key exists else false
	val, ok := mp["apple"]

	fmt.Printf("val -> %d ok -> %v \n", val, ok)
	fmt.Println("map---->", mp)

	// functions

	addAns, diif := add(10, 20)
	fmt.Println("ans---->", addAns, diif)

	// another way to define function some what related to javascript arrow function
	testFunc := func(x int) int {

		return x
	}
	ans := testFunc(15)
	fmt.Println("another way to define function -->", ans)
}
func add(a int, b int) (int, int) {
	return a + b, a - b
}
