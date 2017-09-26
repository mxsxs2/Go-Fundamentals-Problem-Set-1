/*Implement the square root function using Newton’s method. In this case, Newton’s method is to approximate sqrt(x) by picking a starting point z and then repeating:

z_next = z - ((z*z - x) / (2 * z))
To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, …). Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). How close are you to the math.Sqrt value?

Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:

z := float64(1)
z := 1.0*/

//It is the main package
package main

//imports
import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/*Function go calculate square root
This function is more accurate than the below one*/
func Sqrt(x float64) float64 {
	//Return 0 if the input is 0
	if x <= 0 {
		return 0
	}

	//Create variables for calculation
	z := x
	//Loop until it is close to 0.00000001
	for z_next := z + 0.1; math.Abs(z_next-z) > 0.00000001; {
		z = z_next
		//estimate the squerae root
		z_next = z - ((z*z - x) / (2 * z))
	}
	//Return the estimated number
	return z
}

//Main function
func main() {

	num := 0.0
	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		//Parse the incoming command
		n, _ := strconv.ParseFloat(os.Args[1], 64)
		num = n
	} else {
		//Set the default number
		num = 3.0

		fmt.Println("You can specify which number to use for the sqare root by adding a number after the name of the runnable in terminal.")
		fmt.Println("For example: ./nmfsr.exe 12")
		fmt.Println("Default number is 3\n")
	}

	if num <= 0 {
		fmt.Println("The number has to be higher than 0")
	} else {
		fmt.Println("Square root of", num, "in the new method:", Sqrt(num))
		fmt.Println("Square root of", num, "in GO's built in method. math.Sqrt(x)", math.Sqrt(num), "\n")
	}

}
