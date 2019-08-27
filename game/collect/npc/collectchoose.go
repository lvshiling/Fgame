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
	scene.RegisterNPC(types.BiologyScriptTypePearl, scene.NPCFactoryFunc(CreateCollectChooseNPC))
}

type CollectChooseObject struct {
	pl        scene.Player
	startTime int64
}

func CreateCollectChooseObject() *CollectChooseObject {
	o := &CollectChooseObject{}
	return o
}

func (co *CollectChooseObject) GetPlayer() scene.Player {
	return co.pl
}

func (co *CollectChooseObject) GetStartTime() int64 {
	return co.startTime
}

type CollectChooseNPC struct {
	*npc.NPCBase
	collectChooseObject *CollectChooseObject
}

func (cn *CollectChooseNPC) GetCollect() *CollectChooseObject {
	return cn.collectChooseObject
}

func (cn *CollectChooseNPC) IfCanCollect(playerId int64) (flag, isMax bool) {
	if cn.collectChooseObject.pl != nil {
		return
	}
	flag = true
	return
}

// func (cn *CollectChooseNPC) GetCollectNPC() scene.NPC {
// 	return cn.NPCBase
// }

func (cn *CollectChooseNPC) StartCollect(pl scene.Player) (now int64, flag bool) {
	flag, _ = cn.IfCanCollect(pl.GetId())
	if !flag {
		return
	}
	now = global.GetGame().GetTimeService().Now()
	cn.collectChooseObject.pl = pl
	cn.collectChooseObject.startTime = now
	pl.Collect(cn)
	flag = true
	return
}

func (cn *CollectChooseNPC) CollectInterrupt(pl scene.Player) {
	cn.collectChooseObject.pl = nil
	cn.collectChooseObject.startTime = 0
}

func (cn *CollectChooseNPC) CollectFinish(finishType collecttypes.CollectChooseFinishType) {
	if cn.NPCBase.IsDead() {
		return
	}

	flag := cn.NPCBase.Recycle(cn.collectChooseObject.pl.GetId())
	if !flag {
		panic(fmt.Errorf("collectchoose:  CostHP should be ok"))
	}
	// n := cn.GetCollectNPC()
	eventData := collecteventtypes.CreateCollectChooseFinishEventData(cn, finishType)
	pl := cn.collectChooseObject.pl
	cn.collectChooseObject.pl.ClearCollect()
	cn.collectChooseObject.pl = nil
	cn.collectChooseObject.startTime = 0
	gameevent.Emit(collecteventtypes.EventTypeCollectChooseFinish, pl, eventData)
}

// func (cn *CollectChooseNPC) ExitScene(active bool) {
// 	cn.collectChooseObject.pl.ClearCollect()
// 	cn.collectChooseObject.pl = nil
// 	cn.collectChooseObject.startTime = 0
// 	cn.NPCBase.ExitScene(active)
// }

//心跳
func (cn *CollectChooseNPC) Heartbeat() {

	if cn.collectChooseObject.pl != nil && cn.collectChooseObject.startTime != 0 {
		now := global.GetGame().GetTimeService().Now()
		biologyTemplate := cn.GetBiologyTemplate()
		if biologyTemplate == nil {
			return
		}

		collectTime := biologyTemplate.CaiJiTime
		escape := now - cn.collectChooseObject.startTime
		if escape >= int64(collectTime) {
			cn.CollectFinish(collecttypes.CollectChooseFinishTypeLow)
		}
	}
	cn.NPCBase.Heartbeat()
}

func CreateCollectChooseNPC(ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, deadTime int64) scene.NPC {
	n := &CollectChooseNPC{}
	b := npc.NewNPCBase(n, ownerType, ownerId, ownerAllianceId, id, idInScene, biologyTemplate, pos, angle)
	collectObj := CreateCollectChooseObject()
	n.NPCBase = b
	n.collectChooseObject = collectObj
	n.Calculate()
	if deadTime != 0 {
		n.DeadInTime(deadTime)
	}
	return n
}
