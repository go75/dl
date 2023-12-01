package cal

import (
	"dl/math"
	"dl/node"
)

type PowNode[T node.Type] node.BaseCalNode[T]

func (n *PowNode[T]) CalNode() *node.BaseCalNode[T] {
	return (*node.BaseCalNode[T])(n)
}

func (n *PowNode[T]) Forward() {
	n.BackNode.Data = math.Pow(n.PreNodes[0].Data, n.PreNodes[1].Data)
}

func (n *PowNode[T]) Backward() {
	// grad * b * pow(a, b-1)
	grad0 := n.BackNode.Grad * n.PreNodes[1].Data * math.Pow(n.PreNodes[0].Data, n.PreNodes[1].Data - 1)
	// gard * pow(a, b) * ln(a)
	grad1 := n.BackNode.Grad * math.Pow(n.PreNodes[0].Data, n.PreNodes[1].Data) * math.Ln(n.PreNodes[0].Data)

	n.PreNodes[0].Grad += grad0
	n.PreNodes[1].Grad += grad1
}