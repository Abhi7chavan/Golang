package main
import ("fmt"
"bufio"
"strings"
"os"
"strconv"
)

func main(){ 
    reader := bufio.NewReader(os.Stdin)
    var t int;
    fmt.Scan(&t);
    for i:=0;i<t;i++{ 
        line,_:= reader.ReadString('\n')
        line = strings.TrimSpace(line)
        sizes := strings.Split(line," ")
        M,N := sizes[0],sizes[1]
        integerM,_:=strconv.Atoi(M)
        integerN,_:=strconv.Atoi(N)
        matrix:=ReadMatrix(integerM,integerN,reader)
        ans:=CalculateDigonalsum(integerM,integerN,matrix)
        
        fmt.Println("The Primary digonal sum",ans[0])
        fmt.Println("The Secondary digonal sum",ans[1])
    }
}

func ReadMatrix(M,N int,reader *bufio.Reader)[][]int{ 
    var matrix [][]int;
    for i:=0;i<M;i++{
    line,_:= reader.ReadString('\n')
    line = strings.TrimSpace(line);
    sizes:=strings.Split(line," ")
    var subarray []int;
    for _,value:=range(sizes){
        valueint,_:=strconv.Atoi(value);
        subarray=append(subarray,valueint)
    }
    matrix = append(matrix,subarray)
    }
    return matrix
}

func CalculateDigonalsum(M int ,N int,matrix [][]int)[]int{
    primarydigonal:= 0
    secondarydigonal:=0
    var digonalsum []int
    for i:=0;i<M;i++{ 
        primarydigonal+=matrix[i][i]
        secondarydigonal+=matrix[i][M-i-1]
    }
    digonalsum = append(digonalsum,primarydigonal)
    digonalsum =append(digonalsum,secondarydigonal)
    
    return digonalsum
    
}
