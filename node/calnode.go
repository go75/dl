package node

type BaseCalNode[T Type] struct {
	PreNodes []*DataNode[T]
	BackNode *DataNode[T]
}

type CalNode[T Type] interface {
	CalNode() *BaseCalNode[T]
	Forward()
	Backward()
}