
package main
import "fmt"
type Node struct {
	value int
	next  *Node
}

type Linkedlist struct {
	head *Node
}

func (self *Linkedlist) AddNode(value int) {
	newnode := &Node{value: value}
	if self.head == nil {
		self.head = newnode
	} else {
		current := self.head
		for current.next != nil {
			current = current.next
		}
		current.next = newnode
	}
}

func (self *Linkedlist) Print() {
	current := self.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next // Advance to the next node
	}
	fmt.Println("nil")
}

func main() {
	Ld := &Linkedlist{}
	Ld.AddNode(10)
	Ld.AddNode(20)
	Ld.AddNode(30)
	Ld.AddNode(40)

	Ld.Print()
}
