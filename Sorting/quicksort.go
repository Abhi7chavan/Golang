package main

import (
	"fmt"
)

func findDistinct(array []int) []int {
	distinct := []int{}
	seen := make(map[int]bool)

	for _, item := range array {
		if !seen[item] {
			distinct = append(distinct, item)
			seen[item] = true
		}
	}

	return distinct
}

func binarysearch(left int, right int, array []int, key int) bool {
    for left <= right {
        mid := left + (right-left)/2

        if array[mid] == key {
            return true
        } else if array[mid] < key {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return false
}



func linearsearch(array[] int,key int) bool{
	for _ ,item := range array{
	if item == key{
	return true;
	}
}
return false;
}

func quicksort(array[]int)[]int{
	if len(array) <= 1{
	return array
	}
	
	var pivot=array[0]
	
	var left, right []int
	
	for _ , element := range array[1:]{
	if  element<= pivot{
	left= append(left,element)
	}else{
	right =append(right,element)
	}
	
	}
	
	left = quicksort(left)
	right = quicksort(right)
	
	
	return append(append(left,pivot),right...)	
}

func main() {
    items := []int{95,78,46,58,45,86,99,99,251,320}
    fmt.Println(linearsearch(items, 86))

    left := 0
    right := len(items) - 1
    fmt.Println(binarysearch(left, right, items, 9))
    fmt.Println(findDistinct(items))
    fmt.Println(quicksort(items))
}

