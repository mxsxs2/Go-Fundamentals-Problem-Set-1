//Write a program that prints the numbers from 1 to 100, except for the following conditions. For multiples of three print “Fizz” instead of the number, and for the multiples of five print “Buzz”. 
//For numbers which are multiples of both three and five print “FizzBuzz”.

//It is the main package
package main
//imports
import(
	"fmt"
	"bytes"
	"strconv"
)
//Main function
func main() {
	
	//Loop from 1 to 100
	for i:=1; i<=100; i++{
		var output bytes.Buffer
		//Concat Fizz if it can be divided by 3
		if(i%3==0){
			output.WriteString("Fizz")
		}
		//Concat Fizz if it can be divided by 5
		if(i%5==0){
			output.WriteString("Buzz")
		}
		//Concat the number if it cannot be divided by any
		if(i%3!=0 && i%5!=0){
			//Convert int to string first
			output.WriteString(strconv.Itoa(i))
		}
		//Output the result
		fmt.Println(output.String()) 
	}
}