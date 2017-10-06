//Write a function to reverse a string in Go.

//It is the main package
package main

//imports
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

//Main function
func main() {
	//The input string
	input, err := getCommandlineWord()
	//If there was an error with the input
	if err != nil {
		fmt.Println("Please input a word to reverse:")
		//Get input from user
		input = getInputString()
	}

	//Reverse and print out
	fmt.Println("Reversed:", reverse(input))

	fmt.Println("\nYou can specify a word to reverse by adding a number after the name of the runnable in terminal.")
	fmt.Println("For example: ./reverse_string.exe krisztian")

}

//Function used to get the command line word
func getCommandlineWord() (string, error) {
	//Check if there is any argument supplied
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}
	//Return 0 and the error
	return "", errors.New("No command line argument")
}

//Function used to get the input string
func getInputString() string {
	//Create the reader
	reader := bufio.NewReader(os.Stdin)
	//Get the input string
	input, err := reader.ReadString('\n')
	//Trim the input so no trailing space or \n is added
	input = strings.TrimSpace(input)
	//Check if there is anything entered
	if err != nil || len(input) < 1 {
		fmt.Println("You should write something! Please try again:")
		//Get the string again
		getInputString()
	}
	return input
}

//Function used to reverse a string
func reverse(str string) string {
	//The new string holder
	var newString string
	//Loop the string backwards
	for i := len(str) - 1; i >= 0; i-- {
		//Concatenate the character to the new string
		newString += string(str[i])
	}
	//Return the new string
	return newString
}
