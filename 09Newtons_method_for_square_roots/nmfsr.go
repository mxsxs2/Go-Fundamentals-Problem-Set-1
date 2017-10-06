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
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Main function
func main() {
	//Get the number from the command line
	num, err := getCommandlineNumber()

	//If there vas an error with the command line value
	if err != nil {
		//Write out the error
		fmt.Println(err)
		fmt.Println("Please enter a number:")
		//Read the number again
		readFloat(&num)
	}

	fmt.Println("You can specify which number to use for the sqare root by adding a number after the name of the runnable in terminal.")
	fmt.Println("For example: ./nmfsr.exe 12")
	fmt.Println("You can use -verbose flag to print every guess of the Sqrt function")
	fmt.Println("For example: ./nmfsr.exe 12 -verbose\n")

	if num <= 0 {
		fmt.Println("The number has to be higher than 0")
	} else {
		fmt.Println("Square root of", num, "in the new method:", Sqrt(num))
		fmt.Println("Square root of", num, "in GO's built in method. math.Sqrt(x)", math.Sqrt(num), "\n")
	}

}

//Function used to get the command line argument or return 100
func getCommandlineNumber() (float64, error) {
	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		//Parse the incoming command and check if it can be converted to a float
		if n, err := strconv.ParseFloat(os.Args[1], 64); err == nil && n != 0 {
			//Return the number and no error
			return n, nil
		}
	}
	//Return 0 and the error
	return 0, errors.New("The command line input number could not be parsed or it is 0")
}

//Function used to read in a float with input validation
func readFloat(pointer *float64) {
	//Read in the number
	if _, err := fmt.Scanf("%g", pointer); err != nil {
		//If there was an error then rread again
		fmt.Println("Invalid input. Please enter a number again:")
		//Re read the number
		readFloat(pointer)
	}
}

//Function used to find a flag in command line arguments
func isFlagOn(flag string) bool {
	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		//Loop the arguments
		for _, arg := range os.Args {
			//If the flag was found then return true
			if strings.Compare(arg, flag) == 0 {
				return true
			}
		}
	}
	//Return false if there was no result
	return false
}

/*Function go calculate square root
This function is more accurate than the below one*/
func Sqrt(x float64) float64 {
	//Return 0 if the input is 0
	if x <= 0 {
		return 0
	}

	//Create variables for calculation
	z := x
	//Loop until it is close to 0.000000001
	for zNext := z + 0.1; math.Abs(zNext-z) > 0.000000001; {
		//Write the current guess out if the verbose flag exist
		if isFlagOn("-verbose") {
			fmt.Println("Current guess:", z)
		}

		z = zNext
		//estimate the square root
		zNext = z - ((z*z - x) / (2 * z))
	}
	//Return the estimated number
	return z
}
