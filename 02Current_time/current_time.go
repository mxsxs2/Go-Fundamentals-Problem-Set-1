//Write a program that prints the current time and date to the console.
//It is the main package
package main

//import format and time
import (
	"fmt"
	"time"
)

//Main function
func main() {
	//Get the time
	ctime := time.Now()
	//Print the time
	fmt.Println("The time is", ctime.Format("15:04"))
	//Print the date
	fmt.Println("The current date is", ctime.Format(" 02/01/2006"))
}
