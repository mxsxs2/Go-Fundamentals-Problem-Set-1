//Write a function to reverse a string in Go.

//It is the main package
package main
//imports
import(
	"fmt"
)

//Main function
func main() {
	//The input string
	var input string

	fmt.Println("Please input a word to reverse:")
	//Get the input
	fmt.Scanf("%s\n", &input)
	//Reverse and print out
	fmt.Println("Reversed:",reverse(input));
	
}
//Function used to reverse a string
func reverse(str string)string{
	//The new string holder
	var newString string;
	//Loop the string backwards
	for i:=len(str)-1; i>=0;i--{
		//Concatenate the character to the new string
		newString+=string(str[i])
	}
	//Return the new string
	return newString
}


