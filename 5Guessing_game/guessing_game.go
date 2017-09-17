/*Write a guessing game where the user must guess a secret number. 
  After every guess the program tells the user whether their number was too large or too small. 
  At the end the number of tries needed should be printed. It counts only as one try if they input the same number multiple times consecutively.*/

//It is the main package
package main
//imports
import(
	"fmt"
	"math/rand"
	"time"
)
//Main function
func main() {
	//Seed random so it has less chance to give the same numbers
	rand.Seed(time.Now().Unix())
	//The number, max 100
	number:=rand.Intn(100)
	//Num of tries
	var tries int = 0
	//Last try
	var lastTry int = -1
	//Previous try
	var prevTry int = -1

	fmt.Println("Please guess the number which is between 0 and 100")

	//Loop until the number is found
	for number!=lastTry{
		//Set the previous try
		prevTry=lastTry
		//Read in the number
		fmt.Scanf("%d",&lastTry)
		//Check if the number is smaller than try
		if(lastTry>number){
			fmt.Println(lastTry,"is larger")
		//Check if number is bigger than try
		}else if(lastTry<number){
			fmt.Println(lastTry,"is smaller")
		}

		//Increase the number of tries if the tried number is not the same as previous
		//In my understanding i didn't need to store them in an array and count every number only once
		if(prevTry!=lastTry){
			tries++
		}
	}

	fmt.Println("Congratualtions you guessed the number with",tries,"tries")
}