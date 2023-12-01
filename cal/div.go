package cal

import (
	"dl/math"
	"dl/node"
)

type DivNode[T node.Type] node.BaseCalNode[T]

func (n *DivNode[T])CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *DivNode[T]) Forward() {
	n.BackNode.Data = n.PreNodes[0].Data / n.PreNodes[1].Data
}

func (n *DivNode[T]) Backward() {
	// grad / b
	grad0 := n.BackNode.Grad / n.PreNodes[1].Data
	// -(grad * a / pow(b, 2))
	grad1 := n.BackNode.Grad * n.PreNodes[0].Data / math.Pow(n.PreNodes[1].Data, 2)

	n.PreNodes[0].Grad += grad0
	n.PreNodes[1].Grad -= grad1
}