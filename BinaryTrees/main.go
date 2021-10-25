package main

import "fmt"

func main() {
	var bt Tree = Tree{root: nil}
	bt.Add(12)
	bt.Add(5)
	bt.Add(2)
	bt.Add(9)
	bt.Add(18)
	bt.Add(14)
	bt.Add(19)
	bt.Add(13)
	bt.Add(17)
	bt.Add(16)
	bt.Add(15)
	bt.root = Delete(bt.root, 18)
	fmt.Println(bt.root)
}
