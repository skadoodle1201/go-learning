package main

import (
	"errors"
	"fmt"
)

func main() {

	numberA, numberB := 10, 0
	printFunc("Divide")
	fmt.Printf("NumberA : %d NumberB : %d \n", numberA, numberB)
	questiont, remainder, err := divideThese(numberA, numberB)
	// We Can use switch like this as well as the normal way on a variable
	switch {
	case err != nil:
		{
			fmt.Printf(err.Error())
			fmt.Println("After Error")
		}
	default:
		{
			fmt.Printf("NumberA: %d, NumberB: %d , questiont: %d , remainder: %d", numberA, numberB, questiont, remainder)
		}
	}

	var myArr [3]int16
	var myArr2 [3]int16 = [3]int16{1, 2, 3}
	myArr3 := [...]int32{1, 2, 4, 5, 6}
	fmt.Println(myArr[0])
	myArr[2] = 13
	myArr[1] = 4

	//Prints Everything Between from 1 to (3 - 1)
	fmt.Println(myArr[1:3])
	fmt.Println(myArr2, myArr3)

	//Slice is like vector in GO
	var intSlice []int32 = []int32{4, 5, 6}

	//If You know a specific length to use and
	//save time for creating duplicate slice and do append in background use "make"
	//creates a slice with cap as 4 and len 0 if given length initilize with default
	var intSlice2 []int32 = make([]int32, 4, 4)

	fmt.Println(intSlice, intSlice2)
	intSlice = append(intSlice, 8)
	fmt.Println("After Append", intSlice)
	//To Take length and Cap
	fmt.Println("length: ", len(intSlice2), " cap :: ", cap(intSlice2))

	var myMap = map[string]uint8{"adam": 13, "sam": 20}
	fmt.Println(myMap["adam"])

	//Map will always give a default value even if value not present so use the second value as bool to check
	age, status := myMap["sara"]
	fmt.Println(age, status)
	//delete a key from a map
	delete(myMap, "sara")

	// order changes as soon as you loop over a map
	for name := range myMap {
		fmt.Printf("Name : %s \t Age : %d \n", name, myMap[name])
	}
	for name, age := range myMap {
		fmt.Printf("Name : %s \t Age : %d \n", name, age)
	}
}

func printFunc(printThis string) {
	fmt.Printf(printThis)
}

func divideThese(numberA, numberB int) (questiont int, remainder int, err error) {
	// By Default This Will Be "nil"
	if numberB == 0 {
		err = errors.New("Cannot Divide By 0")
		return 0, 0, err
	}
	questiont = numberA / numberB
	remainder = numberA % numberB

	return questiont, remainder, err
}
