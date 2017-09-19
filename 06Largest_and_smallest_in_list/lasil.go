//Write a function that returns the largest and smallest elements in a list.

//It is the main package
package main
//imports
import(
	"fmt"
)
//Main function
func main() {
	//Create a slice
	var numbersSlice []int
	//Size of the list
	var listSize int=0
	//Maximum size of the list.So we know when to stop taking in numbers
	var listMaxSize int
	//Temporary number container
	var temp int
	//Smallest number
	var smallest int
	//Largest number
	var largest int


	fmt.Println("Please enter the size of the list:")
	//Get the size of the list
	fmt.Scanf("%d",&listMaxSize)

	//Loop until we have all the numbers
	for listSize<listMaxSize{
		fmt.Println("Number",listSize+1)
		//Get the number
		fmt.Scanf("%d",&temp)
		//Add the number to the end
		numbersSlice=append(numbersSlice,temp)
		//Increase the list size
		listSize++
		//Initialize the smallest and largest if it is not done yet
		if(listSize==1){
			smallest=temp
			largest=temp
		}
	}

	//Loop the numbers
	for _, number:= range numbersSlice{
		//Set the smallest if this number is the current smallest number
		if(number<smallest){
			smallest=number
		}
		//Set the highest if this number is the current highest number
		if(number>largest){
			largest=number
		}
	}

	//Write the answer
	fmt.Println("The smallest in the list is:",smallest)
	fmt.Println("The largest in the list is:",largest)




	
}