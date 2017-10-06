//Write a function that merges two sorted lists into a new sorted list. [1,4,6],[2,3,5] → [1,2,3,4,5,6].
//Since the numbers were specified in the excersise sheet i believe i do not have to implement any input or even string merging
//I used float 64 to store the arrays as this way the function can be used for floats and ints as well. (If it is used for int after merge it has to be casted to int)

//It is the main package
package main

//imports
import (
	"fmt"
)

//Main function
func main() {

	fmt.Println("The merged and sorted array from [1.1,3.2,4,6] and [2,3.1,5] is:")
	//Merge and print array
	fmt.Println(merge([]float64{1.1, 3.2, 4, 6}, []float64{2, 3.1, 5}))
}

//Function which merges two sorted slices (technically this is the second part of a merge sort)
func merge(l1 []float64, l2 []float64) []float64 {
	//Initialize the new slice from the first slice
	var newlist []float64
	//If the size of the l1 or l2 are higher than 0
	for len(l1) > 0 || len(l2) > 0 {
		//If both are higher than 0
		if len(l1) > 0 && len(l2) > 0 {
			//If the first number of the first array is lower
			if l1[0] <= l2[0] {
				//Add to the new list
				newlist = append(newlist, l1[0])
				//Resize the slice
				l1 = l1[1:len(l1)]
			} else {
				//Add to the new list
				newlist = append(newlist, l2[0])
				//Resize the slice
				l2 = l2[1:len(l2)]
			}
			//If just the first array is higher then 0
		} else if len(l1) > 0 {
			newlist = append(newlist, l1[0])
			l1 = l1[1:len(l1)]
			//If just the second array is higher then 0
		} else if len(l2) > 0 {
			newlist = append(newlist, l2[0])
			l2 = l2[1:len(l2)]
		}
	}

	//Retrun the new sorted array
	return newlist
}
