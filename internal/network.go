package internal

import (
	"fmt"
	"net"
	"github.com/m1gwings/treedrawer/tree"
)


func buildNetTree() *tree.Tree {
	rootTree := tree.NewTree(tree.NodeString("This machine"))

	ifaces, _ := net.Interfaces()
	for index, i := range ifaces {

		addr, _ := i.Addrs()
		v4, v4ipnet, _ := net.ParseCIDR(addr[1].String())

		rootTree.AddChild(tree.NodeString(i.Name))

		nodeNetName, _ := rootTree.Child(index)
		nodeNetName.AddChild(tree.NodeString(i.Flags.String()))

		nodeNetFlags, _ := nodeNetName.Child(0)
		nodeNetFlags.AddChild(tree.NodeString(v4.String()))

		nodeAddr, _ := nodeNetFlags.Child(0)
		nodeAddr.AddChild(tree.NodeString(v4ipnet.String()))

	}
	
	return rootTree

}

func DisplayNetTree() {
	tree_ := buildNetTree()
	fmt.Print(tree_)
}