package main

import "fmt"

func quickSort(data []int) {
	// base condition
	if len(data) <= 1 {
		return
	}

	pivot := data[0]
	lessThanPivot := make([]int, 0)
	greaterThanPivot := make([]int, 0)

	// partitioning
	for _, value := range data[1:] {
		if value <= pivot {
			lessThanPivot = append(lessThanPivot, value)
		} else {
			greaterThanPivot = append(greaterThanPivot, value)
		}
	}

	// recursively quicksort left and right partition
	quickSort(lessThanPivot)
	quickSort(greaterThanPivot)

	copy(data, append(append(lessThanPivot, pivot), greaterThanPivot...))
}

func main() {
	data := []int{24, 9, 29, 19, 14, 27}
	quickSort(data)
	fmt.Println(data)
}

// note
// this quicksort implementation is use first element of array as a pivot
// because it use recursion we need to set the base condition as a condition of recursion to stop
// oh this case we set the recursion to end if length of a array is less than equal to 1

// after set the base condition, we set the first element of array as pivot
// and then we create 2 slices, lessThanPivot to collect data than less than pivot, and greaterThanPivot to collect data than greater than pivot

// next, we loop through the array, but we exclude the first element bc first element is our pivot, we use slice (data[1:]) that is mean we start from index 1 until end
// inside loop we check if the value is less than equal to pivot
// if true append the lessThanPivot with that value
// else append the graterThanPivot with that value

// after finishing the loop now we have 2 partition, lessThanPivot and greaterThanPivot
// we will quickSort each of partition again until hit the base case, after that it will append the pivot data into lessThanPivot and then append that into greaterThanPivot

// lets we jump to the case:
// ===========================================
// we have unsorted list [24,9,14,19,29,27]
// pivot will be the first element of list, so:
// pivot = 24
// next create the empty list to collect lessThanPivot and greaterThanPivot
// after doing loop to determine which one is less or greater than pivot we got:
// lessThanPivot = [19,9,14]
// greaterThanPivot = [29,27]

// first we quickSort the lessThanPivot (recursive 1)
// lessThanPivot = [19,9,14]
// on the process that will create another lessThanPivot and greaterThanPivot again and again until the base case is hit, for example:

// pivot = 19 				(recursive 2)
// after doing loop we got
// lessThanPivot = [9,14]
// greaterThanPivot = []

// we quickSort the lessThanPivot again, (recursive 3)
// lessThanPivot = [9,14]
// pivot = 9
// after doing loop we got
// lessThanPivot = []
// greaterThanPivot = [14]

// we quickSort the lessThanPivot again, (recursive 4)
// lessThanPivot = []
// this will hit out base case because the length is <= 1
// now we will jump to recursive 3 process that is quickSort the greaterThanPivot

// we quickSort the greaterThanPivot (recursive 3)
// greterThanPivot = [14]
// this will hit out base case because the length is <= 1
// now we will jump to recursive 3 process that is copy and append
// what we have now on recursive no 3 is
// pivot = 9
// lessThanPivot = []
// greaterThanPivot = [14]

// first we append pivot into lessThanPivot and we got [9]
// next we append greaterThanPivot into lessThanPivot and we got [9,14]

// after that we jump to recursive no 2
// what we have now on recursive no 2 is
// pivot = 19
// lessThanPivot = [9,14]
// greaterThanPivot = []

// first we append pivot into lessThanPivot and we got [9,14,19]
// next we append greaterThanPivot into lessThanPivot and we got [9,14,19]

// we have finished the lessThan pivot on recurstion no 1, after that will wil jump to recursive 1 to contiue quick sorting greaterThanPivot
// now we will quick sort greater than pivot [29,27] on recursion 1
// REMEMBER: WE STILL NOT DOING PROCESS COPY AND APPEND FOR RECURSION NO 1 SO WE STILL NOT FINISHED THE RECURSION NO 1 YET

// same like before we need to jump determine the pivot and create 2 partition again
// pivot = 29				(recursion 2b)
// after doing loop we got
// lessThanPivot = [27]
// greaterThanPivot = []

// next we quicksort lessThanPivot (recursion 3b)
// we hit the base case so we jump to recursion 2b to quickSort the greater that pivot recursion

// recursion 4b
// bc of greater that pivot is empty list we hit base case again, so we jump to recurstion 2 and continue to copy and append

// what we have now on recursive no 2b is
// pivot = 29
// lessThanPivot = [27]
// greaterThanPivot = []

// first we append pivot into lessThanPivot and we got [27,29]
// next we append greaterThanPivot into lessThanPivot and we got [27,29]

// now we have finished quicsort greater than pivot on recursion no 1, next we will continue to copy and append on recursion no 1

// what we have now on recursive no 1 is
// pivot = 24
// lessThanPivot = [9,14,19]
// greaterThanPivot = [27,29]

// first we append pivot into lessThanPivot and we got [9,14,19,24]
// next we append greaterThanPivot into lessThanPivot and we got [9,14,19,24,27,29]

// we finished quick sort the data
