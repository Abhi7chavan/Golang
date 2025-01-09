// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func solve(a,b int) int {
    if (b==0){
        return 1
    }
    
    ans:= solve(a,b/2)
    ans*=ans 
    if (b % 2 == 1){
        ans*=a
    }
    
    return ans
}

func main() {
  var a ,b int
  fmt.Scan(&a,&b)
  result:= solve(a,b)
  fmt.Println(result)
  
}
