package flowmap

import (
	"dl/node"
)

type FlowMap[T node.Type] struct {
	nodes []node.CalNode[T]
}

func (m FlowMap[T]) Build() {
		
}

func (m FlowMap[T])Forward() {
	n := len(m.nodes)
	for i := 0; i < n; i++ {
		m.nodes[i].Forward()
	}
}

func (m FlowMap[T])Backward() {
	for i := len(m.nodes) - 1; i >= 0; i-- {
		m.nodes[i].Backward()
	}
}