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

	num := int64(100)
	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		//Parse the incoming command
		n, _ := strconv.ParseInt(os.Args[1], 10, 0)
		num = n
	}
	//Check if the given number is higher than 0
	if num > 0 {

		//Factorial holder
		factorial := big.NewInt(num)
		//Sum of it
		var sum int = 0
		//Loop the 100
		for i := 100; i > 0; i-- {
			//Multiply n-1
			factorial.Mul(factorial, big.NewInt(int64(i)))
		}

		//Holder for the number at position
		n := new(big.Int)

		//Loop until the end of the number
		for factorial.BitLen() > 0 {
			//Calculation: factorial/(position on power of 10)%10
			//And save to n
			factorial.DivMod(factorial, new(big.Int).SetUint64(uint64(10)), n)
			//Add n to the sum
			sum += int(n.Uint64())
		}
		//Print the sum
		fmt.Println("The sum of factorial", num, " is :", sum)
	} else {
		fmt.Println("The number has to be higher than 0")
	}
	fmt.Println("\nYou can specify which number to use for the factorial by adding a number after the name of the runnable in terminal.")
	fmt.Println("For example: ./factorial_digit_sum.exe 101")

}
