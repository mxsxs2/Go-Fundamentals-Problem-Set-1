/*Implement the square root function using Newton’s method. In this case, Newton’s method is to approximate sqrt(x) by picking a starting point z and then repeating:

z_next = z - ((z*z - x) / (2 * z))
To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, …). Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). How close are you to the math.Sqrt value?

Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:

z := float64(1)
z := 1.0*/

//It is the main package
package main
//imports
import(
	"fmt"
	"math"
)

//Function go calculate square root
func Sqrt(x float64) float64 {
	//Initialize the estimated sqrt to the number itself
 	z := x
    
	//Function to estimate the squerae root
    calculate := func() float64 {
    	return z - ((z*z - x) / (2 * z))
    }
    //Loop until the estimation is very close to 0
    for z_next := calculate(); math.Abs(z_next - z) > 0.00000001
    {	
		//Set the new estimation to be the old one
    	z = z_next
		//Recalculate the estimation
		z_next=calculate()
    }
	//Return the estimated number
    return z
}
//Main function
func main() {

	fmt.Println("1 new:",Sqrt(1))
	fmt.Println("1 math.Sqrt(x)",math.Sqrt(1))

	fmt.Println("2 new:",Sqrt(2))
	fmt.Println("3 math.Sqrt(x)",math.Sqrt(2))

	fmt.Println("3 new:",Sqrt(3))
	fmt.Println("3 math.Sqrt(x)",math.Sqrt(3))
}


