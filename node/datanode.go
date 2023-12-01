package node

type DataNode[T Type] struct {
	PreNode *CalNode[T]
	BackNode []*CalNode[T]
	Data T
	Grad T
}