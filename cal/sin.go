package cal

import (
	"dl/math"
	"dl/node"
)

type SinNode[T node.Type] node.BaseCalNode[T]

func (n *SinNode[T])CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *SinNode[T]) Forward() {
	n.BackNode.Data = math.Sin(n.PreNodes[0].Data)
}

func (n *SinNode[T]) Backward() {
	// grad * cos(a)
	grad := n.BackNode.Grad * math.Cos(n.PreNodes[0].Data)

	n.PreNodes[0].Grad += grad
}