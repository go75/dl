package cal

import (
	"dl/math"
	"dl/node"
)

type CosNode[T node.Type] node.BaseCalNode[T]

func (n *CosNode[T])CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *CosNode[T]) Forward() {
	n.BackNode.Data = math.Cos(n.PreNodes[0].Data)
}

func (n *CosNode[T]) Backward() {
	// -(grad * sin(a))
	grad := n.BackNode.Grad * math.Sin(n.PreNodes[0].Data)

	n.PreNodes[0].Grad -= grad
}