package scene

import (
	"fgame/fgame/core/fsm"
	coretypes "fgame/fgame/core/types"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/scene/types"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

//状态管理器
type NPCStateManager interface {
	//状态
	fsm.Subject
	//返回
	Back() bool
	//跟踪
	Trace() bool
	//待机
	Idle() bool
	//当前行为
	GetCurrentAction() NPCAction
}

//npc接口
type NPC interface {
	//战斗对象
	BattleObject
	//状态管理器
	NPCStateManager
	Heartbeat()
	PlayerCalled(playerId int64)
	AllianceCalled(allianceId int64)
	ChuangShiCampCalled(campType chuangshitypes.ChuangShiCampType)
	//回收
	Recycle(playerId int64) bool
	//获取主人类型
	GetOwnerType() scenetypes.OwnerType
	//获取主人
	GetOwnerId() int64
	//获取主人
	GetOwnerAllianceId() int64
	GetName() string
	//获取模型id  客户端使用
	GetTempId() int
	//场景id
	GetIdInScene() int32
	//获取生物模板
	GetBiologyTemplate() *gametemplate.BiologyTemplate
	//获取出生地点
	GetBornPosition() coretypes.Position

	//死亡时间
	GetDeadTime() int64
	ShouldRemove() bool
}

type CollectNPC interface {
	NPC
	CollectInterrupt(pl Player)
	// GetCollectNPC() NPC
	IfCanCollect(playerId int64) (flag, isMax bool)
	StartCollect(pl Player) (now int64, flag bool)
}

//npc工厂
type NPCFactory interface {
	CreateNPC(ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, deadTime int64) NPC
}

type NPCFactoryFunc func(ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, deadTime int64) NPC

func (nff NPCFactoryFunc) CreateNPC(ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, deadTime int64) NPC {
	return nff(ownerType, ownerId, ownerAllianceId, id, idInScene, biologyTemplate, pos, angle, deadTime)
}

var (
	npcFactoryMap map[types.BiologyScriptType]NPCFactory = make(map[types.BiologyScriptType]NPCFactory)
)

func RegisterNPC(bt types.BiologyScriptType, nf NPCFactory) {
	npcFactoryMap[bt] = nf
}

func CreateNPC(ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, deadTime int64) (n NPC) {
	nf, exist := npcFactoryMap[biologyTemplate.GetBiologyScriptType()]
	if !exist {
		return
	}
	n = nf.CreateNPC(ownerType, ownerId, ownerAllianceId, id, idInScene, biologyTemplate, pos, angle, deadTime)

	// n.UpdateBattleProperty(propertynpctypes.PropertyEffectorTypeMaskAll)
	return
}

type NPCState fsm.State

const (
	//待机状态
	NPCStateInit fsm.State = iota
	//追击状态
	NPCStateTrace
	//攻击
	NPCStateAttack
	//被打状态
	NPCStateAttacked
	//返回状态
	NPCStateBack
	//死亡
	NPCStateDead
)

const (
	EventNPCIdle     fsm.Event = "idle"
	EventNPCTrace              = "trace"
	EventNPCAttack             = "attack"
	EventNPCAttacked           = "attacked"
	EventNPCBack               = "back"
	EventNPCDead               = "dead"
)

var (
	npcStateMachine *fsm.StateMachine
)

var (
	transitions = []*fsm.Trasition{
		//初始化->追踪
		&fsm.Trasition{
			From:  NPCStateInit,
			To:    NPCStateTrace,
			Event: EventNPCTrace,
		},
		//初始化->攻击
		&fsm.Trasition{
			From:  NPCStateInit,
			To:    NPCStateAttack,
			Event: EventNPCAttack,
		},
		//初始化->被攻击
		&fsm.Trasition{
			From:  NPCStateInit,
			To:    NPCStateAttacked,
			Event: EventNPCAttacked,
		},
		//初始化->死亡
		&fsm.Trasition{
			From:  NPCStateInit,
			To:    NPCStateDead,
			Event: EventNPCDead,
		},
		//追踪->攻击
		&fsm.Trasition{
			From:  NPCStateTrace,
			To:    NPCStateAttack,
			Event: EventNPCAttack,
		},
		//追踪->被攻击
		&fsm.Trasition{
			From:  NPCStateTrace,
			To:    NPCStateAttacked,
			Event: EventNPCAttacked,
		},
		//追踪->返回
		&fsm.Trasition{
			From:  NPCStateTrace,
			To:    NPCStateBack,
			Event: EventNPCBack,
		},
		//追踪->返回
		&fsm.Trasition{
			From:  NPCStateTrace,
			To:    NPCStateDead,
			Event: EventNPCDead,
		},
		// 攻击->追踪
		&fsm.Trasition{
			From:  NPCStateAttack,
			To:    NPCStateTrace,
			Event: EventNPCTrace,
		},
		// 攻击->被攻击
		&fsm.Trasition{
			From:  NPCStateAttack,
			To:    NPCStateAttacked,
			Event: EventNPCAttacked,
		},
		// 攻击->返回
		&fsm.Trasition{
			From:  NPCStateAttack,
			To:    NPCStateBack,
			Event: EventNPCBack,
		},
		// 攻击->死亡
		&fsm.Trasition{
			From:  NPCStateAttack,
			To:    NPCStateDead,
			Event: EventNPCDead,
		},

		// 被攻击->追击
		&fsm.Trasition{
			From:  NPCStateAttacked,
			To:    NPCStateTrace,
			Event: EventNPCTrace,
		},
		// 被攻击->返回
		// &fsm.Trasition{
		// 	From:  NPCStateAttacked,
		// 	To:    NPCStateBack,
		// 	Event: EventNPCBack,
		// },
		// 被攻击 -> 死亡
		&fsm.Trasition{
			From:  NPCStateAttacked,
			To:    NPCStateDead,
			Event: EventNPCDead,
		},

		// 返回-> 死亡
		&fsm.Trasition{
			From:  NPCStateBack,
			To:    NPCStateDead,
			Event: EventNPCDead,
		},

		//返回->待机
		&fsm.Trasition{
			From:  NPCStateBack,
			To:    NPCStateInit,
			Event: EventNPCIdle,
		},
		// 死亡->待机
		&fsm.Trasition{
			From:  NPCStateDead,
			To:    NPCStateInit,
			Event: EventNPCIdle,
		},
	}
)
var (
	stateMachine *fsm.StateMachine
)

func init() {
	stateMachine = fsm.NewStateMachine(transitions)
}

func GetNPCStateMachine() *fsm.StateMachine {
	return stateMachine
}

type NPCAction interface {
	Tick(NPC)
}

type NPCActionHandler func(NPC)

func (h NPCActionHandler) Tick(n NPC) {
	h(n)
}

//默认行为
func dummyAction(n NPC) {
}

var (
	scriptStateMap  = make(map[scenetypes.BiologyScriptType]map[fsm.State]NPCAction)
	defaultStateMap = make(map[fsm.State]NPCAction)
)

func RegisterDefaultAction(state fsm.State, action NPCAction) {
	_, ok := defaultStateMap[state]
	if ok {
		panic(fmt.Errorf("repeate register default action  %d state script action", state))
	}
	defaultStateMap[state] = action
}

func RegisterAction(typ scenetypes.BiologyScriptType, state fsm.State, action NPCAction) {
	stateMap, ok := scriptStateMap[typ]
	if !ok {
		stateMap = make(map[fsm.State]NPCAction)
		scriptStateMap[typ] = stateMap
	}
	_, ok = stateMap[state]
	if ok {
		panic(fmt.Errorf("repeate register %d typ, %d state script action", typ, state))
	}
	stateMap[state] = action
}

func GetAction(typ scenetypes.BiologyScriptType, state fsm.State) NPCAction {
	stateMap, ok := scriptStateMap[typ]
	if !ok {
		return getDefaultAction(state)

	}
	stateAction, ok := stateMap[state]
	if !ok {
		return getDefaultAction(state)
	}
	return stateAction
}

func getDefaultAction(state fsm.State) NPCAction {
	stateAction, ok := defaultStateMap[state]
	if !ok {
		return NPCActionHandler(dummyAction)
	}
	return stateAction
}

type NPCActionState fsm.State

const (
	//待机状态
	NPCActionStateIdle fsm.State = iota
	//移动
	NPCActionStateMove
	//攻击
	NPCActionStateAttack
	//被攻击
	NPCActionStateAttacked
	//死亡
	NPCActionStateDead
)

const (
	EventNPCActionIdle     fsm.Event = "idle"
	EventNPCActionMove               = "trace"
	EventNPCActionAttack             = "attack"
	EventNPCActionAttacked           = "attacked"
	EventNPCActionDead               = "dead"
)

var (
	npcActionStateMachine *fsm.StateMachine
)

var (
	actionTransitions = []*fsm.Trasition{
		//初始化->移动
		&fsm.Trasition{
			From:  NPCActionStateIdle,
			To:    NPCActionStateMove,
			Event: EventNPCActionMove,
		},
		//初始化->攻击
		&fsm.Trasition{
			From:  NPCActionStateIdle,
			To:    NPCActionStateAttack,
			Event: EventNPCActionAttack,
		},
		//初始化->被攻击
		&fsm.Trasition{
			From:  NPCActionStateIdle,
			To:    NPCActionStateAttacked,
			Event: EventNPCActionAttacked,
		},
		//初始化->死亡
		&fsm.Trasition{
			From:  NPCActionStateIdle,
			To:    NPCActionStateDead,
			Event: EventNPCActionDead,
		},
		//移动->攻击
		&fsm.Trasition{
			From:  NPCActionStateMove,
			To:    NPCActionStateIdle,
			Event: EventNPCActionIdle,
		},
		//移动->攻击
		&fsm.Trasition{
			From:  NPCActionStateMove,
			To:    NPCActionStateAttack,
			Event: EventNPCActionAttack,
		},
		//移动->被攻击
		&fsm.Trasition{
			From:  NPCActionStateMove,
			To:    NPCActionStateAttacked,
			Event: EventNPCActionAttacked,
		},
		//移动->死亡
		&fsm.Trasition{
			From:  NPCActionStateMove,
			To:    NPCActionStateDead,
			Event: EventNPCActionDead,
		},
		// 攻击->空闲
		&fsm.Trasition{
			From:  NPCActionStateAttack,
			To:    NPCActionStateIdle,
			Event: EventNPCActionIdle,
		},
		// 攻击->被攻击
		&fsm.Trasition{
			From:  NPCActionStateAttack,
			To:    NPCActionStateAttacked,
			Event: EventNPCActionAttacked,
		},
		// 攻击->死亡
		&fsm.Trasition{
			From:  NPCActionStateAttack,
			To:    NPCActionStateDead,
			Event: EventNPCActionDead,
		},
		// 被攻击->空闲
		&fsm.Trasition{
			From:  NPCActionStateAttacked,
			To:    NPCActionStateIdle,
			Event: EventNPCActionIdle,
		},

		// 被攻击->死亡
		&fsm.Trasition{
			From:  NPCActionStateAttacked,
			To:    NPCActionStateDead,
			Event: EventNPCActionDead,
		},
		// 死亡->空闲
		&fsm.Trasition{
			From:  NPCActionStateDead,
			To:    NPCActionStateIdle,
			Event: EventNPCActionIdle,
		},
	}
)
