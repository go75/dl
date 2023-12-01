package cal

import "dl/node"

type SubNode[T node.Type] node.BaseCalNode[T]

func (n *SubNode[T])CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *SubNode[T]) Forward() {
	n.BackNode.Data = n.PreNodes[0].Data - n.PreNodes[1].Data
}

func (n *SubNode[T]) Backward() {
	// gard
	n.PreNodes[0].Grad += n.BackNode.Grad
	// -grad
	n.PreNodes[1].Grad -= n.BackNode.Grad
}