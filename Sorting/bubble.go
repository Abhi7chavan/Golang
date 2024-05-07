// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func main() {
  array := []int{8, 2, 11, 90, 21, 55, 29, 86}
  fmt.Println(bubblesort(array))
}

func bubblesort(array[]int) []int{
    if len(array) <= 1{
        return array
    }
    for i:= 0;i<len(array);i++ {
        for j:=1;j<len(array);j++{
            if array[j] < array[j-1]{
                array[j],array[j-1]=array[j-1],array[j]
                
            }
        }

        
    }
return array    
}
