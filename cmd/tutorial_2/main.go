package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//defaults to int32 or int64 based in the system
	var basicInt int = 12
	fmt.Println(basicInt)

	//tells the compiler that this line is 8 bits and no more
	var eightBitInt int8 = 12
	fmt.Println(eightBitInt)

	//can only store positives and provides more memory storage
	var onlyPositiveNums uint8 = 12
	fmt.Println(onlyPositiveNums)

	//can store digits but will only return the value as it's stored
	var floatNums float32 = 12345678.9
	fmt.Println(floatNums) //>>123456789.0000

	//now this will give me the correct number
	var correctReturnValue float64 = 12345678.9
	fmt.Println(correctReturnValue) // >>12345678.9000

	//you cant add different types to each other, but you can make them the same
	//and then do it
	var num1 float32 = 12.1
	var num2 int8 = 1
	var num3 float32 = num1 + float32(num2)
	fmt.Println(num3)

	//integer devision results in an integer, it's cant contain decimals, so the result is rounded down
	var intNum1 = 3
	var intNum2 = 2
	fmt.Println(intNum1 / intNum2) //>>1

	//string can be used with both "" and ``, the difference is tha you cant do multiline with the first one
	var basicStr string = "String"
	var singleQuote string = `string
    string`
	fmt.Println(basicStr + singleQuote)

	//the strings can be tricky to use, for example if i were do to something like this
	fmt.Println(len("test")) //>>4 bytes, not amount of chars
	//it would return the number of bytes, so i were to use something not basic it would return more
	//to count the amount of chars in a string i can use a utf8.RuneCountInString()
	utf8.RuneCountInString("test") //4 chars

	//single chars are to be called runes
	var newRune rune = 'a'
	fmt.Println(newRune)

	//and the easiest data type in go are booleans
	var myBool bool = true
	fmt.Println(myBool)

	//we can create a var without initializing it, then it will just make a default value itself
	var myNewInt int //>>defaults to zero
	fmt.Println(myNewInt)

	//we can make a var without type declaration
	var myNewSrt = `some kind of string`
	//or drop the var too
	mySecondNewString := `some king of string`
	fmt.Println(myNewSrt + mySecondNewString)

	//we can initialize multiple var at the same time
	one, two := 1, 2
	fmt.Println(one, two)

	//we have constanses too
	const newConst string = "new string"
	fmt.Println(newConst)
}
