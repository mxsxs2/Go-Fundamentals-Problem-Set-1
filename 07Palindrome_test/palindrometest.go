//Write a function that tests whether a string is a palindrome.

//It is the main package
package main

//imports
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Main function
func main() {

	fmt.Println("Please input a word or sentence to check if it is a palindrome:")
	//Get the input
	input := getInputString()
	//Check if it is palindrome
	if isPalindrome(strings.ToLower(input)) {
		fmt.Println("Your text is palindrome")
	} else {
		fmt.Println("Your text is not palindrome")
	}
}

//Function used to check if a text is palindrome
func isPalindrome(text string) bool {
	//Set the two directions of the loop
	front := 0
	back := len(text) - 1
	//Loop until the two numbers meet
	for front < back {
		//If there is any difference then it is not a palindrome. It is converted into runes for more simple number comparison
		if rune(text[front]) != rune(text[back]) {
			return false
		}
		front++
		back--
	}
	//It is a palindrome
	return true
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
