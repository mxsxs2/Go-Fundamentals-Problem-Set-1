/*Write a function that calculates the sum of the digits of the factorial of a number.
n! means n x (n − 1) … x 3 x 2 x 1. For example, 10! = 10 x 9 x … x 3 x 2 x 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100!

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
	"math/big"
	"os"
	"strconv"
)

//Main function
func main() {
	//Get the number from the arguments
	num, err := getNumber()
	//Calculate the factorial
	fact := getFactorialofnum(num)
	//claclutae factorial and add the numbers together
	sum := getSumofBigInt(fact)

	//Check if there was any error with the number intake
	if err != nil {
		//Write out the error. In this case the error is formatted
		fmt.Println(err)
	}

	//If the printvalue flag is turned on
	if isFlagOn("-printvalue") {
		//Print the factorial out
		fmt.Printf("%d!==", fact)
	}

	//Print the sum
	fmt.Println("The sum of factorial", num, " is :", sum)

	fmt.Println("\nYou can specify which number to use for the factorial by adding a number after the name of the runnable in terminal.")
	fmt.Println("For example: ./factorial_digit_sum.exe 101")
	fmt.Println("You can use -printvalue flag to print the factorial number as well")
	fmt.Println("For example: ./factorial_digit_sum.exe 101 -printvalue")

}

//Function used to get the command line argument or return 100
func getNumber() (int64, error) {
	num := int64(100)
	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		//Parse the incoming command and check if it can be converted to an int
		if n, err := strconv.ParseInt(os.Args[1], 10, 0); err == nil && n > 0 {
			//Return the number and no error
			return n, nil
		} else {
			//Return the initially set number and the error
			return num, errors.New("The command line input number could not be parsed or it is 0. Default 100 is used")
		}

	}
	//Return the initially set number and no error if there was no command line input
	return num, nil
}

//Function used to find a flag in command line arguments
func isFlagOn(flag string) bool {
	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		//Loop the arguments
		for _, arg := range os.Args {
			//If the flag was found then return true
			if arg == flag {
				return true
			}
		}
	}
	//Return false if there was no result
	return false
}

//Function used to get a factorial of a number
func getFactorialofnum(num int64) *big.Int {
	//Factorial holder
	factorial := new(big.Int)
	//Do factorial
	return factorial.MulRange(1, num)
}

//Function used to return the sum of a big int
func getSumofBigInt(factorialNumber *big.Int) int {
	//Sum of it
	sum := 0
	//Holder for the number at position
	n := new(big.Int)

	//Do factorial and loop it's numbers
	for factorialNumber.BitLen() > 0 {
		//Calculation: factorial/(position on power of 10)%10
		//And save to n
		factorialNumber.DivMod(factorialNumber, new(big.Int).SetUint64(uint64(10)), n)
		//Add n to the sum
		sum += int(n.Uint64())
	}

	return sum
}
