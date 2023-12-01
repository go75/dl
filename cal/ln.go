package cal

import (
	"dl/math"
	"dl/node"
)

type LnNode[T node.Type] node.BaseCalNode[T]

func (n *LnNode[T])CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *LnNode[T]) Forward() {
	n.BackNode.Data = math.Ln(n.PreNodes[0].Data)
}

func (n *LnNode[T]) Backward() {
	// grad / a
	grad := n.BackNode.Grad / n.PreNodes[0].Data

	n.PreNodes[0].Grad += grad
}