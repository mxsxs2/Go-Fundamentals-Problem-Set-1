/*Write a function that returns the largest and smallest elements in a list.
It uses a simple solution to check largest and smallest in a list, stores the numbers in a list then loops the list.
The list content can be float and int as well as the int will be casted to floats

Could be even simpler if i dont store the numbers just check which is the smallest and largest while the user enters them.

Could implement a merge(or any other sort) sort for the numbers and just get the first and last number in the list. I believe that is a complete overkill here.
*/

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
	//Maximum size of the list.So we know when to stop taking in numbers
	var listMaxSize int

	fmt.Println("Please enter the size of the list:")
	//Get the size of the list
	readInt(&listMaxSize)

	//Get the list type
	types := getListType()
	//If the user want number list
	if types == 1 {
		//Get the list
		list := getNumberList(listMaxSize)
		//Get the smallest and largest
		smallest, largest := getLargestAndSmallestNumbers(list)
		//Write the answer
		fmt.Println("The smallest in the list is:", smallest)
		fmt.Println("The largest in the list is:", largest)
		fmt.Println("Your list was:", list)
	} else {
		//Get the list
		list := getStringList(listMaxSize)
		//Get the smallest and largest
		smallest, largest := getLargestAndSmallestString(list)
		//Write the answer
		fmt.Println("The smallest in the list is:", smallest)
		fmt.Println("The largest in the list is:", largest)
		fmt.Println("Your list was:", list)
	}

}

//Function used to get a list type
func getListType() int {
	types := 0
	//Loop until we get an answer
	for types != 1 && types != 2 {
		fmt.Println("Please enter the number of the type of list which you would like to use:")
		fmt.Println("1. Numbers")
		fmt.Println("2. Words/Sentences (lexicographically sort)")
		//Get the type
		readInt(&types)
	}
	return types
}

//Function used to read in a integer with input validation
func readInt(pointer *int) {
	//Read in the number
	if _, err := fmt.Scanf("%d", pointer); err != nil {
		//If there was an error then rread again
		fmt.Println("Invalid input. Please enter a number again:")
		//Re read the number
		readInt(pointer)
	}
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

//Function used to get the number list
func getNumberList(listMaxSize int) []float64 {
	//Temporary number holder
	temp := 0.0
	//Create a slice
	var numbersSlice []float64
	//Size of the list
	listSize := 0
	//Loop until we have all the numbers
	for listSize < listMaxSize {
		fmt.Println("Enter Number:", listSize+1)
		//Get the number
		readFloat(&temp)
		//Add the number to the end
		numbersSlice = append(numbersSlice, temp)
		//Increase the list size
		listSize++
	}
	//Return the numbers
	return numbersSlice
}

//Function used to get the string list
func getStringList(listMaxSize int) []string {
	//Create a slice
	var stringsSlice []string
	//Size of the list
	listSize := 0
	//Loop until we have all the strings
	for listSize < listMaxSize {
		fmt.Println("Enter String Number:", listSize+1)
		//Get the string
		temp := getInputString()
		//Add the number to the end
		stringsSlice = append(stringsSlice, temp)
		//Increase the list size
		listSize++
	}
	//Return the strings
	return stringsSlice
}

//Function used to get the largest and the smallest  number in a list
//This is a simple solution wich supports int and float numbers as well.
func getLargestAndSmallestNumbers(list []float64) (float64, float64) {
	//Smallest number
	var smallest float64
	//Largest number
	var largest float64
	//Loop the numbers
	for i, number := range list {
		//Initialize the largest and the smallest if there was none
		if i == 0 {
			smallest = number
			largest = number
		} else {
			//Set the smallest if this number is the current smallest number
			if number < smallest {
				smallest = number
			}
			//Set the highest if this number is the current highest number
			if number > largest {
				largest = number
			}
		}
	}
	//Return the smallest and the largest
	return smallest, largest
}

//Function used to get the largest and the smallest string in a list
func getLargestAndSmallestString(list []string) (string, string) {
	//Smallest string
	var smallest string
	//Largest string
	var largest string
	//Loop the strings
	for i, str := range list {
		//Initialize the largest and the smallest if there was none
		if i == 0 {
			smallest = str
			largest = str
		} else {
			//Set the smallest if this string is the current smallest string
			if strings.Compare(str, smallest) < 0 {
				smallest = str
			}
			//Set the highest if this string is the current highest string
			if strings.Compare(str, largest) > 0 {
				largest = str
			}
		}
	}
	//Return the smallest and the largest
	return smallest, largest
}
