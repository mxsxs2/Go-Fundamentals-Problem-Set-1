//Write a function that tests whether a string is a palindrome.

//It is the main package
package main
//imports
import(
	"fmt"
	"strings"
)
//Main function
func main() {
	//The input string
	var input string

	fmt.Println("Please input a word to check if it is a palindrome:")
	//Get the input
	fmt.Scanf("%s\n", &input)
	//Check if it is palindrome
	if(isPalindrome(strings.ToLower(input))){
		fmt.Println("Your text is palindrome")
	}else{
		fmt.Println("Your text is not palindrome")
	}
}


//Function used to check if a text is palindrome
func isPalindrome(text string) bool{
	var front int = 0
	back:= len(text)-1
	//Loop until the two numbers meet
	for front<back{
		//If there is any difference then it is not a palindrome
		if(string([]rune(text)[front])!=string([]rune(text)[back])){
			return false
		}
		front++
		back--
	}
	//It is a palindrome
	return true
}