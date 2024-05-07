// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func main(){
    var array = []int{88,22,11,9,8,72,2}
    fmt.Println(selectionsort(array))
}

func selectionsort(array[]int)[]int{
    if len(array) <= 1{
        return array
    }
    for i:=0;i<=len(array)-1;i++{
        min:=i
        for j:=i+1;j<=len(array)-1;j++{
            if array[min]> array[j]{
                min = j
            }
            
        }
        array[min],array[i] = array[i],array[min]
    }
    return array
}
