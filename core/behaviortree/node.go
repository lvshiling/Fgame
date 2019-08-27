package behaviortree

type Node interface {
}

type node struct {
	parent Node
}

type compositeNode struct {
	nodeList []Node
}
