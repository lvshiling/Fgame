package fsm

//状态
type State int32

//事件
type Event string

type Trasition struct {
	From  State
	To    State
	Event Event
}

type StateMachine struct {
	transitions []*Trasition
}

func (m *StateMachine) Trigger(s Subject, event Event) (flag bool) {

	t := m.findMatchTransition(s.CurrentState(), event)
	if t == nil {
		return false
	}
	s.OnExit(s.CurrentState())
	toState := t.To
	s.OnEnter(toState)
	return true
}

func (m *StateMachine) findMatchTransition(state State, event Event) *Trasition {
	for _, transition := range m.transitions {
		if transition.From == state && event == transition.Event {
			return transition
		}
	}
	return nil
}

func NewStateMachine(transitions []*Trasition) *StateMachine {
	sm := &StateMachine{}
	sm.transitions = transitions
	return sm
}
