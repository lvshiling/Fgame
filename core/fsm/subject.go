package fsm

//状态目标
type Subject interface {
	//当前状态
	CurrentState() State
	//进入状态前
	OnEnter(state State)
	//退出状态前
	OnExit(state State)
}

//状态目标基类
type SubjectBase struct {
	s State
}

func (sb *SubjectBase) CurrentState() State {
	return sb.s
}

func (sb *SubjectBase) OnExit(ss State) {
}

func (sb *SubjectBase) OnEnter(ss State) {
	sb.s = ss

}

func NewSubjectBase(s State) *SubjectBase {
	return &SubjectBase{
		s: s,
	}
}
