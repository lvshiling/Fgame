package npc

import (
	coretypes "fgame/fgame/core/types"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	collecttypes "fgame/fgame/game/collect/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/npc/npc"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/game/scene/types"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

func init() {
	scene.RegisterNPC(types.BiologyScriptTypeXianTaoQianNianCollect, scene.NPCFactoryFunc(CreateCollectPointNPC))
	scene.RegisterNPC(types.BiologyScriptTypeXianTaoBaiNianCollect, scene.NPCFactoryFunc(CreateCollectPointNPC))
	scene.RegisterNPC(types.BiologyScriptTypeLongGongTreasure, scene.NPCFactoryFunc(CreateCollectPointNPC))
}

type collectPlayer struct {
	pl        scene.Player
	startTime int64
}

func (co *collectPlayer) GetPlayer() scene.Player {
	return co.pl
}

func (co *collectPlayer) GetStartTime() int64 {
	return co.startTime
}

type CollectPointObject struct {
	playerMap       map[int64]collectPlayer
	totalCount      int32
	useCount        int32
	lastRecoverTime int64
}

func CreateCollectPointObject(tCount int32, now int64) *CollectPointObject {
	o := &CollectPointObject{
		totalCount:      tCount,
		useCount:        0,
		lastRecoverTime: now,
	}
	o.playerMap = make(map[int64]collectPlayer)
	return o
}

func (co *CollectPointObject) GetPlayer(playerId int64) scene.Player {
	st, ok := co.playerMap[playerId]
	if !ok {
		return nil
	}
	return st.GetPlayer()
}

func (co *CollectPointObject) GetStartTime(playerId int64) int64 {
	st, ok := co.playerMap[playerId]
	if !ok {
		return 0
	}
	return st.GetStartTime()
}

func (co *CollectPointObject) GetTotalCount() int32 {
	return co.totalCount
}

func (co *CollectPointObject) GetUseCount() int32 {
	return co.useCount
}

func (co *CollectPointObject) GetLastRecoverTime() int64 {
	return co.lastRecoverTime
}

func (co *CollectPointObject) GmSetUseCount(count int32) {
	if count > co.totalCount {
		count = co.totalCount
	}
	co.useCount = count
	return
}

func (co *CollectPointObject) GmSetLastRecoverTime(now int64) {
	co.lastRecoverTime = now
	return
}

type CollectPointNPC struct {
	*npc.NPCBase
	collectPointObject *CollectPointObject
}

func (cn *CollectPointNPC) GetCollect() *CollectPointObject {
	return cn.collectPointObject
}

func (cn *CollectPointNPC) IfCanCollect(playerId int64) (flag, isMax bool) {
	biologyTemplate := cn.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	if biologyTemplate.CaiJiLimitCount != 0 {
		totalCount := cn.collectPointObject.GetTotalCount()
		useCount := cn.collectPointObject.GetUseCount()
		if totalCount <= useCount {
			isMax = true
			return
		}
	}
	_, ok := cn.collectPointObject.playerMap[playerId]
	if ok {
		return
	}
	flag = true
	return
}

// func (cn *CollectPointNPC) GetCollectNPC() scene.NPC {
// 	return cn.NPCBase
// }

func (cn *CollectPointNPC) StartCollect(pl scene.Player) (now int64, flag bool) {
	flag, _ = cn.IfCanCollect(pl.GetId())
	if !flag {
		return
	}
	now = global.GetGame().GetTimeService().Now()
	st := collectPlayer{
		pl:        pl,
		startTime: now,
	}
	cn.collectPointObject.playerMap[pl.GetId()] = st
	pl.Collect(cn)
	flag = true
	return
}

func (cn *CollectPointNPC) CollectInterrupt(pl scene.Player) {
	playerId := pl.GetId()
	_, ok := cn.collectPointObject.playerMap[playerId]
	if ok {
		delete(cn.collectPointObject.playerMap, playerId)
	}
}

func (cn *CollectPointNPC) collectFinish(playerId int64) {
	if cn.NPCBase.IsDead() {
		return
	}
	cp, ok := cn.collectPointObject.playerMap[playerId]
	if !ok {
		return
	}

	biologyTemplate := cn.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}

	if biologyTemplate.CaiJiLimitCount != 0 {
		cn.collectPointObject.useCount += 1
		gameevent.Emit(collecteventtypes.EventTypeCollectPointChange, cn, nil)
	}

	if biologyTemplate.GetCollectPointFinishType() == collecttypes.CollectPointFinishTypeAway {
		totalCount := cn.collectPointObject.GetTotalCount()
		useCount := cn.collectPointObject.GetUseCount()
		if totalCount <= useCount {
			flag := cn.NPCBase.Recycle(playerId)
			if !flag {
				panic(fmt.Errorf("collect:  CostHP should be ok"))
			}
		}
	}
	// n := cn.GetCollectNPC()

	gameevent.Emit(collecteventtypes.EventTypeCollectFinish, cp.GetPlayer(), cn)
	cp.GetPlayer().ClearCollect()
	delete(cn.collectPointObject.playerMap, playerId)
}

//心跳
func (cn *CollectPointNPC) Heartbeat() {
	now := global.GetGame().GetTimeService().Now()
	biologyTemplate := cn.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}
	collectTime := biologyTemplate.CaiJiTime

	for _, cp := range cn.collectPointObject.playerMap {
		escape := now - cp.GetStartTime()
		if escape >= int64(collectTime) {
			cn.collectFinish(cp.GetPlayer().GetId())
		}
	}

	// recoverTime := biologyTemplate.CaiJiRecoverTime
	// if recoverTime != 0 && cn.collectPointObject.lastRecoverTime+recoverTime < now {
	// 	cn.collectPointObject.totalCount = biologyTemplate.CaiJiLimitCount
	// 	cn.collectPointObject.useCount = 0
	// 	cn.collectPointObject.lastRecoverTime = now
	// 	gameevent.Emit(collecteventtypes.EventTypeCollectPointChange, cn, nil)
	// }

	cn.NPCBase.Heartbeat()
}

func CreateCollectPointNPC(ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, deadTime int64) scene.NPC {
	n := &CollectPointNPC{}
	b := npc.NewNPCBase(n, ownerType, ownerId, ownerAllianceId, id, idInScene, biologyTemplate, pos, angle)
	now := global.GetGame().GetTimeService().Now()
	obj := CreateCollectPointObject(biologyTemplate.CaiJiLimitCount, now)
	n.NPCBase = b
	n.collectPointObject = obj
	n.Calculate()
	if deadTime != 0 {
		n.DeadInTime(deadTime)
	}
	return n
}
