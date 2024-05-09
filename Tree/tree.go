package main

import ("fmt"
"strings")

type Node struct{
    value string
    children []*Node
}
type Tree struct{
    root *Node
}
func NewNode(value string) *Node{
    return &Node{value:value}
}
func (n *Node) AddChild(child *Node){
    n.children = append(n.children,child)
    fmt.Printf("\nChild Added %s in Node\n",child)
}

func (t *Tree) PrintTree(node *Node,level int){
    if node != nil{
        fmt.Println(strings.Repeat(" ",level)+node.value)
    }
    
    for _ , child := range node.children{
        t.PrintTree(child,level+1)
    }
}
func main(){
tree := Tree{root:NewNode("A")}
b:= NewNode("B")
c:= NewNode("C")
d:= NewNode("D")
e:= NewNode("E")

tree.root.AddChild(b)
tree.root.AddChild(c)

b.AddChild(d)
c.AddChild(e)

tree.PrintTree(tree.root,0)
}
