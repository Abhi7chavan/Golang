// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"



type Friends struct{
    parent map[string]string
    rank map[string]int
}

func Friendbook() *Friends{
    return &Friends{parent:make(map[string]string),
    rank:make(map(string)int),
    }
}

func (fd *Friendbook) AddFriends(x string,y string){
    
}

func (fd *Friendbook) Find_parent(x string) string{
    if _,exits := fd.parent[x]; !exits{
        fd.parent[x] = x
        fd.rank[x] = 0
    }else{
        fd.parent[x] = fd.Find_parent(fd.parent[x])
    }
    return fd.parent[x]    
}

func (fd *Friendbook) Remove_Friend(x string,y string){
    if fd.AreFriends(x,y){
        fd.parent(y) = y
    }else{
        fmt.Println("%s and %s are not friends",x,y)
    }
}

func (fd *Friendbook) AreFriends(x string,y string) bool{
    return fd.find_parent(x) == fd.find_parent(y)
}
func main() {
  
}
