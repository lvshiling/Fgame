package behaviortree

type State int32

type Task interface {
	Node() Node
}
