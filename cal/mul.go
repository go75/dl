package cal

import "dl/node"

type MulNode[T node.Type] node.BaseCalNode[T]

func (n *MulNode[T])CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *MulNode[T]) Forward() {
	n.BackNode.Data = n.PreNodes[0].Data * n.PreNodes[1].Data
}

func (n *MulNode[T]) Backward() {
	// grad * b
	grad0 := n.BackNode.Grad * n.PreNodes[1].Data
	// grad * a
	grad1 := n.BackNode.Grad * n.PreNodes[0].Data

	n.PreNodes[0].Grad += grad0
	n.PreNodes[1].Grad += grad1
}