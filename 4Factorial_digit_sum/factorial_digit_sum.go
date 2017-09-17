/*Write a function that calculates the sum of the digits of the factorial of a number. 
n! means n x (n − 1) … x 3 x 2 x 1. For example, 10! = 10 x 9 x … x 3 x 2 x 1 = 3628800, 
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27. 
Find the sum of the digits in the number 100!*/

//It is the main package
package main
//imports
import(
	"fmt"
	"math/big"
)
//Main function
func main() {
	//Factorial holder
	factorial :=big.NewInt(1)
	//Sum of it
	var sum int =0;
	//Loop the 100
	for i:= 100; i>0; i--{
		//Multiply n-1
		factorial.Mul(factorial,big.NewInt(int64(i)))
	}

	//Holder for the number at position
	n := new(big.Int)

	//Loop until the end of the number
	for factorial.BitLen() >0{
		//Calculation: factorial/(position on power of 10)%10
		//And save to n
		factorial.DivMod(factorial, new(big.Int).SetUint64(uint64(10)), n)
		//Add n to the sum
		sum+=int(n.Uint64())
	}
	//Print the sum
	fmt.Println(sum)
	
}