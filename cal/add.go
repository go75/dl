package cal

import "dl/node"

type AddNode[T node.Type] node.BaseCalNode[T]

func (n *AddNode[T])CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *AddNode[T]) Forward() {
	n.BackNode.Data = n.CalNode().PreNodes[0].Data + n.PreNodes[1].Data
}

func (n *AddNode[T]) Backward() {
	// grad
	n.PreNodes[0].Grad += n.BackNode.Grad
	// grad
	n.PreNodes[1].Grad += n.BackNode.Grad
}
