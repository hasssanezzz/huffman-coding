package huffman

import "fmt"

type Node struct {
	Index      int
	Freq       int
	Byte       byte
	Left       *Node
	Right      *Node
	isInternal bool
}

func (n *Node) PrintSymbol() {
	fmt.Printf("%c -> %d\n", rune(n.Byte), n.Freq)
}
