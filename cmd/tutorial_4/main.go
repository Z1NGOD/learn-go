package main

import "fmt"

func main() {
	//arrays have :
	//fixed lengh
	//same type
	//indexable
	//contiguous in memory
	var intArr [3]int32
	fmt.Println(intArr[0], intArr[1:3], &intArr[0]) //& -> pointer where it is stored in memory

	//we can make arrays like this too:
	var intArr2 [3]int32 = [3]int32{1, 2, 3} //declaration on the spot
	intArr3 := [3]int32{1, 2, 3}             // use of := operator
	intArr4 := [...]int32{1, 2, 3}           // [...] -> if we dont know how much we will have in this array
	intArr5 := []int32{}                     // just use something like this
	fmt.Println(intArr2, intArr3, intArr4, intArr5)

	//to make a slice (array), we use make func
	intArr6 := make([]int32, 3, 8) //3-> lengh of a slice, 8->capacity
	fmt.Println(intArr6)

	//we can create maps wih make func too
	myMap := make(map[string]uint8)
	myMap2 := map[string]uint8{"Adam": 32}
	fmt.Println(myMap, myMap2)
	//maps always return something, so if we were to access the value of some key in map and there
	//would be nothing, then we get 0,   myMap2["Adam"]
	//luckily for us maps can return second value -> bool
	var newInt, isOk = myMap2["Sergay"]
	if isOk {
		fmt.Println(newInt)
	} else {
		fmt.Println("Invalid key")
	}

	//we can iterate though map with loops
	for name := range myMap2 {
		fmt.Println(name)
	}
	//we can get key value too
	for name, key := range myMap2 {
		fmt.Printf("This is the name: %v, and key:%v", name, key)
	}
	// we can use loops like in other languages too
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
