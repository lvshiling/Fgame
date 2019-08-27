package npc

import (
	coretypes "fgame/fgame/core/types"
	collecteventtypes "fgame/fgame/game/collect/event/types"
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
	scene.RegisterNPC(types.BiologyScriptTypeGeneralCollect, scene.NPCFactoryFunc(CreateNPC))
}

type CollectObject struct {
	pl        scene.Player
	startTime int64
}

func CreateCollectObject() *CollectObject {
	o := &CollectObject{}
	return o
}

func (co *CollectObject) GetPlayer() scene.Player {
	return co.pl
}

func (co *CollectObject) GetStartTime() int64 {
	return co.startTime
}

type CollectNPC struct {
	*npc.NPCBase
	collectObject *CollectObject
}

func (cn *CollectNPC) GetCollect() *CollectObject {
	return cn.collectObject
}

func (cn *CollectNPC) IfCanCollect(playerId int64) (flag, isMax bool) {
	if cn.collectObject.pl != nil {
		return
	}
	flag = true
	return
}

// func (cn *CollectNPC) GetCollectNPC() scene.NPC {
// 	return cn.NPCBase
// }

func (cn *CollectNPC) StartCollect(pl scene.Player) (now int64, flag bool) {
	flag, _ = cn.IfCanCollect(pl.GetId())
	if !flag {
		return
	}
	now = global.GetGame().GetTimeService().Now()
	cn.collectObject.pl = pl
	cn.collectObject.startTime = now
	pl.Collect(cn)
	flag = true
	return
}

func (cn *CollectNPC) CollectInterrupt(pl scene.Player) {
	cn.collectObject.pl = nil
	cn.collectObject.startTime = 0
}

func (cn *CollectNPC) collectFinish() {
	if cn.NPCBase.IsDead() {
		return
	}
	// curHp := cn.NPCBase.GetHP()
	// flag := cn.NPCBase.CostHP(curHp, cn.collectObject.pl.GetId())
	flag := cn.NPCBase.Recycle(cn.collectObject.pl.GetId())
	if !flag {
		panic(fmt.Errorf("collect:  CostHP should be ok"))
	}
	// n := cn.GetCollectNPC()

	spl := cn.collectObject.pl
	cn.collectObject.pl.ClearCollect()
	cn.collectObject.pl = nil
	cn.collectObject.startTime = 0
	gameevent.Emit(collecteventtypes.EventTypeCollectFinish, spl, cn)
}

//心跳
func (cn *CollectNPC) Heartbeat() {

	if cn.collectObject.pl != nil && cn.collectObject.startTime != 0 {
		now := global.GetGame().GetTimeService().Now()
		biologyTemplate := cn.GetBiologyTemplate()
		if biologyTemplate == nil {
			return
		}
		//collectTime := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeCollectTime)
		collectTime := biologyTemplate.CaiJiTime
		escape := now - cn.collectObject.startTime
		if escape >= int64(collectTime) {
			cn.collectFinish()
		}
	}
	cn.NPCBase.Heartbeat()
}

// func (cn *CollectNPC) ExitScene(active bool) {
// 	cn.collectObject.pl.ClearCollect()
// 	cn.collectObject.pl = nil
// 	cn.collectObject.startTime = 0
// 	cn.NPCBase.ExitScene(active)
// }

func CreateNPC(ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, deadTime int64) scene.NPC {
	n := &CollectNPC{}
	b := npc.NewNPCBase(n, ownerType, ownerId, ownerAllianceId, id, idInScene, biologyTemplate, pos, angle)
	collectObj := CreateCollectObject()
	n.NPCBase = b
	n.collectObject = collectObj
	n.Calculate()
	if deadTime != 0 {
		n.DeadInTime(deadTime)
	}
	return n
}
