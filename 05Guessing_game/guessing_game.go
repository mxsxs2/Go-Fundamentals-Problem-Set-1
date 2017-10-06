/*Write a guessing game where the user must guess a secret number.
  After every guess the program tells the user whether their number was too large or too small.
  At the end the number of tries needed should be printed. It counts only as one try if they input the same number multiple times consecutively.

 Used this link as an example for returning errors and error handling in general
https://gobyexample.com/errors
How to use command line arguments
https://gobyexample.com/command-line-arguments

*/

//It is the main package
package main

//imports
import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Main function
func main() {
	//Num of tries
	tries := 0
	//Last try
	var lastTry int
	//Holder for the tried numbers
	var triedNumbers []int

	//Get the random number
	number, max, err := getNumber()
	//Let the user know if there was any parsing error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Please guess the number which is between 0 and", max)

	//Loop until the number is found
	for number != lastTry {
		//Read in the number from the user into the lastTry variable
		readInt(&lastTry)

		//Compare the numbers
		cmp := compareNumbers(lastTry, number)

		//Check if this number was already tried
		if findNumberInSlice(lastTry, triedNumbers) {
			fmt.Println("You already tried", lastTry)
		} else {
			//Increase the number of tries if the tried number is not the same as previous
			tries++
			//Add the number to the tried numbers slice
			triedNumbers = append(triedNumbers, lastTry)
		}

		//Check if the number is smaller than try
		if cmp == -1 {
			fmt.Println("The guessed number", lastTry, "is larger")
			//Check if number is bigger than try
		} else if cmp == 1 {
			fmt.Println("The guessed number", lastTry, "is smaller")
		}

	}

	fmt.Println("Congratualtions you guessed the number with", tries, "tries")
}

//Function used to get the command line argument and return a random number between 0 and the command line argument.
//If the argument is invlaid or less than 0 then it createsthe random between 0 and 100
func getNumber() (int, int, error) {
	//Seed random so it has less chance to give the same numbers
	rand.Seed(time.Now().Unix())

	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		//Parse the incoming command and check if it can be converted to an int
		if n, err := strconv.ParseInt(os.Args[1], 10, 0); err == nil && n > 0 {
			//Return a random number between 0 and n. Return no error
			//The numbers converted to int from int64
			return rand.Intn(int(n)), int(n), nil
		} else {
			//Return a random number between 0 and 100. Set number and the error
			return rand.Intn(100), 100, errors.New("The command line input number could not be parsed or it is 0.\n Default 100 is used")
		}

	}
	//Return a random number between 0 and 100. No error if there was no command line input
	return rand.Intn(100), 100, nil
}

//Function used to read in a integer with input validation
func readInt(pointer *int) {
	//Read in the number
	if _, err := fmt.Scanf("%d", pointer); err != nil {
		//If there was an error then rread again
		fmt.Println("Invalid input. Please enter a number again:")
		//Re read the number
		readInt(pointer)
	}
}

//Function used to find a number in a slice
func findNumberInSlice(number int, slice []int) bool {
	//Get the lenght of the slice for the loop
	sliceLenght := len(slice)
	//Get the middle of the slice. It is rounded up so it will work with even 1 element
	middleOfSlice := int(math.Ceil(float64(sliceLenght) / 2))
	//Do a loop which goes from the back and from the front as well
	//"Devide and conquer" as John Healy said always in Data Structures and Agorithms class
	for i := 0; middleOfSlice > i; i++ {
		//Check if the the number is the same as any of the two current ones
		if slice[i] == number || slice[sliceLenght-1-i] == number {
			return true
		}
	}
	return false
}

//Function to check if the guessed number is larger or smaller
//If the first number is higher it returns -1
//If the numbers are the same then return 0
//If the second number is higher then returns 1
//It works the same way as the string compare in C
func compareNumbers(num1 int, num2 int) int {
	if num1 > num2 {
		return -1
	} else if num1 == num2 {
		return 0
	} else {
		return 1
	}
}
