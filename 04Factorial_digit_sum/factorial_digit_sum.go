/*Write a function that calculates the sum of the digits of the factorial of a number.
n! means n x (n − 1) … x 3 x 2 x 1. For example, 10! = 10 x 9 x … x 3 x 2 x 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100!*/

//It is the main package
package main

//imports
import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

//Main function
func main() {
	//Get the number from the arguments
	num := getNumber()
	//claclutae factorial and add the numbers together
	sum := getSumofBigInt(getFactorialofnum(num))

	//Check if the given number is higher than 0
	if num > 0 {
		//Print the sum
		fmt.Println("The sum of factorial", num, " is :", sum)
	} else {
		fmt.Println("The number has to be higher than 0")
	}
	fmt.Println("\nYou can specify which number to use for the factorial by adding a number after the name of the runnable in terminal.")
	fmt.Println("For example: ./factorial_digit_sum.exe 101")

}

//Function used to get the command line argument or return 100
func getNumber() int64 {
	num := int64(100)
	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		//Parse the incoming command
		n, _ := strconv.ParseInt(os.Args[1], 10, 0)
		num = n
	}
	return num
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
