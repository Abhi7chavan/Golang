package main
import "fmt"


func mergesort(array[]int) []int{
    if len(array) <= 1{
        return array
    }
    mid := (len(array)/2)
    left:= mergesort(array[:mid])
    right:= mergesort(array[mid:])
    
    return merge(left,right)
}

func merge(left,right []int) []int{
    L,R := 0,0
    var result [] int
    for L < len(left) && R < len(right){
        if left[L] < right[R]{
            result = append(result,left[L])
            L++
        }else{
            result = append(result,right[R])
            R++
        }
    }
    result = append(result,left[L:]...)
    result= append(result,right[R:]...)
    return result
}
func main(){
    var array= []int{12,4,5,7,99,3}
    fmt.Println(mergesort(array))
    
}
