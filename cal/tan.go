package cal

import (
	"dl/math"
	"dl/node"
)

type TanNode[T node.Type] node.BaseCalNode[T]

func (n *TanNode[T])CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *TanNode[T]) Forward() {
	n.BackNode.Data = math.Tan(n.PreNodes[0].Data)
}

func (n *TanNode[T]) Backward() {
	// grad / pow(cos(a), 2)
	grad := n.BackNode.Grad / math.Pow(math.Cos(n.PreNodes[0].Data), 2)

	n.PreNodes[0].Grad += grad
}