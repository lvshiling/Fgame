package scene

import (
	"fgame/fgame/core/fsm"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

//状态管理器
type LingTongStateManager interface {
	//状态
	fsm.Subject
	//跟踪
	Trace() bool
	//待机
	Idle() bool
	//当前行为
	GetCurrentAction() LingTongAction
}

type LingTongShowManager interface {
	//时装
	GetLingTongFashionId() int32
	SetLingTongFashionId(fashionId int32)
	//冰魂
	GetLingTongWeaponId() int32
	SetLingTongWeapon(weaponId int32, weaponState int32)
	//冰魂觉醒
	GetLingTongWeaponState() int32
	GetLingTongTitleId() int32
	SetLingTongTitleId(titleId int32)
	GetLingTongWingId() int32
	SetLingTongWingId(wingId int32)
	GetLingTongMountId() int32
	SetLingTongMountId(mountId int32)
	LingTongMountHidden(hidden bool)
	IsLingTongMountHidden() bool
	//身法
	GetLingTongShenFaId() int32
	SetLingTongShenFaId(shenFaId int32)
	//领域
	GetLingTongLingYuId() int32
	SetLingTongLingYuId(lingYuId int32)
	//法宝id
	GetLingTongFaBaoId() int32
	SetLingTongFaBaoId(faBaoId int32)
	//仙体
	GetLingTongXianTiId() int32
	SetLingTongXianTiId(xianTiId int32)
}

//灵童接口
type LingTong interface {
	//战斗对象
	BattleObject
	Heartbeat()
	//获取主人
	GetOwner() Player
	//获取名字
	GetName() string
	UpdateName(name string)
	//展示
	LingTongShowManager
	//状态
	LingTongStateManager
	GetLingTongTemplate() *gametemplate.LingTongTemplate
	UpdateLingTongTemplate(lingTong *gametemplate.LingTongTemplate, name string)
	GetLingTongId() int32
	GetLoadedPlayers() map[int64]Player
	RemoveLoadedPlayer(id int64)
}

type LingTongState fsm.State

const (
	//待机状态
	LingTongStateInit fsm.State = iota
	//追击状态
	LingTongStateTrace
	//攻击
	LingTongStateAttack
)

const (
	EventLingTongIdle   fsm.Event = "idle"
	EventLingTongTrace            = "trace"
	EventLingTongAttack           = "attack"
)

var (
	lingTongTransitions = []*fsm.Trasition{
		//初始化->追踪
		&fsm.Trasition{
			From:  LingTongStateInit,
			To:    LingTongStateTrace,
			Event: EventLingTongTrace,
		},
		//初始化->攻击
		&fsm.Trasition{
			From:  LingTongStateInit,
			To:    LingTongStateAttack,
			Event: EventLingTongAttack,
		},
		//追踪->攻击
		&fsm.Trasition{
			From:  LingTongStateTrace,
			To:    LingTongStateAttack,
			Event: EventLingTongAttack,
		},
		//追踪->初始化
		&fsm.Trasition{
			From:  LingTongStateTrace,
			To:    LingTongStateInit,
			Event: EventLingTongIdle,
		},
		// 攻击->追踪
		&fsm.Trasition{
			From:  LingTongStateAttack,
			To:    LingTongStateTrace,
			Event: EventLingTongTrace,
		},
		// 攻击->初始化
		&fsm.Trasition{
			From:  LingTongStateAttack,
			To:    LingTongStateInit,
			Event: EventLingTongIdle,
		},
	}
)
var (
	lingTongStateMachine *fsm.StateMachine
)

func init() {
	lingTongStateMachine = fsm.NewStateMachine(lingTongTransitions)
}

func GetLingTongStateMachine() *fsm.StateMachine {
	return lingTongStateMachine
}

type LingTongAction interface {
	OnEnter()
	Action(LingTong)
	OnExit()
}

type DummyLingTongAction struct {
}

func (a *DummyLingTongAction) OnEnter() {
	return
}

func (a *DummyLingTongAction) Action(p LingTong) {
	return
}

func (a *DummyLingTongAction) OnExit() {
	return
}

func NewDummyLingTongAction() *DummyLingTongAction {
	return &DummyLingTongAction{}
}

type LingTongActionHandler func(LingTong)

func (h LingTongActionHandler) Action(n LingTong) {
	h(n)
}

type LingTongActionFactory interface {
	CreateAction() LingTongAction
}

type LingTongActionFactoryFunc func() LingTongAction

func (f LingTongActionFactoryFunc) CreateAction() LingTongAction {
	return f()
}

var (
	dummyLingTongActionInstance = NewDummyLingTongAction()
)

var (
	lingTongStateMap = make(map[fsm.State]LingTongActionFactory)
)

func RegisterLingTongActionFactory(state fsm.State, actionFactory LingTongActionFactory) {
	_, ok := lingTongStateMap[state]
	if ok {
		panic(fmt.Errorf("重复注册灵童%d行为", state))
	}
	lingTongStateMap[state] = actionFactory
}

func GetLingTongAction(state fsm.State) LingTongAction {

	stateActionFactory, ok := lingTongStateMap[state]
	if !ok {
		return dummyLingTongActionInstance
	}
	return stateActionFactory.CreateAction()
}
