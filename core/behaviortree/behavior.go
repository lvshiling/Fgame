package behaviortree

//行为状态
type BehaviorStatus int32

const (
	BehaviorStatusIdle BehaviorStatus = iota
	BehaviorStatusRunning
)

type Behavior struct {
	node   Node
	task   Task
	status BehaviorStatus
}

func (b *Behavior) Update(input interface{}, output interface{}) BehaviorStatus {

	return b.status
}
