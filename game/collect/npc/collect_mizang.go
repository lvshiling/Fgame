package npc

import (
	coretypes "fgame/fgame/core/types"
	collecteventtypes "fgame/fgame/game/collect/event/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/npc/npc"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
)

type miZangCollectPlayer struct {
	pl        scene.Player
	startTime int64
	finish    bool
}

func (co *miZangCollectPlayer) GetPlayer() scene.Player {
	return co.pl
}

func (co *miZangCollectPlayer) GetStartTime() int64 {
	return co.startTime
}

func (co *miZangCollectPlayer) Finish() bool {
	if co.finish {
		return false
	}
	co.finish = true
	return true
}

func (co *miZangCollectPlayer) IsFinish() bool {
	return co.finish
}

func newMiZangCollectPlayer(pl scene.Player, startTime int64) *miZangCollectPlayer {
	mp := &miZangCollectPlayer{}
	mp.pl = pl
	mp.startTime = startTime
	return mp
}

type CollecMiZangObject struct {
	playerMap map[int64]*miZangCollectPlayer
}

func CreateCollecMiZangObject() *CollecMiZangObject {
	o := &CollecMiZangObject{}
	o.playerMap = make(map[int64]*miZangCollectPlayer)
	return o
}

func (co *CollecMiZangObject) GetMiZangPlayer(playerId int64) *miZangCollectPlayer {
	st, ok := co.playerMap[playerId]
	if !ok {
		return nil
	}
	return st
}

func (co *CollecMiZangObject) GetStartTime(playerId int64) int64 {
	st, ok := co.playerMap[playerId]
	if !ok {
		return 0
	}
	return st.GetStartTime()
}

func (co *CollecMiZangObject) Finish(playerId int64) bool {
	st, ok := co.playerMap[playerId]
	if !ok {
		return false
	}
	return st.Finish()
}

func (co *CollecMiZangObject) StartCollect(p scene.Player, now int64) bool {
	playerId := p.GetId()
	_, ok := co.playerMap[playerId]
	if ok {
		return false
	}
	mp := newMiZangCollectPlayer(p, now)
	co.playerMap[playerId] = mp
	return true
}

func (co *CollecMiZangObject) CollectInterrupt(p scene.Player) {
	delete(co.playerMap, p.GetId())
}

type CollecMiZangNPC struct {
	*npc.NPCBase
	collecMiZangObject *CollecMiZangObject
	parentId           int32
	miZangTemplate     *gametemplate.BossMiZangTemplate
}

func (cn *CollecMiZangNPC) GetMiZangTemplate() *gametemplate.BossMiZangTemplate {
	return cn.miZangTemplate
}

//心跳
func (cn *CollecMiZangNPC) Heartbeat() {
	now := global.GetGame().GetTimeService().Now()
	biologyTemplate := cn.GetBiologyTemplate()

	collectTime := biologyTemplate.CaiJiTime
	for _, cp := range cn.collecMiZangObject.playerMap {
		escape := now - cp.GetStartTime()
		if escape >= int64(collectTime) {
			cn.collectFinish(cp.GetPlayer())
		}
	}

	cn.NPCBase.Heartbeat()
}

func (cn *CollecMiZangNPC) IfCanCollect(playerId int64) (flag, isMax bool) {
	p := cn.collecMiZangObject.GetMiZangPlayer(playerId)
	if p == nil {
		return true, false
	}
	return false, false
}

func (cn *CollecMiZangNPC) StartCollect(pl scene.Player) (now int64, flag bool) {
	flag, _ = cn.IfCanCollect(pl.GetId())
	if !flag {
		return
	}
	//TODO: 移动到外面
	now = global.GetGame().GetTimeService().Now()
	flag = cn.collecMiZangObject.StartCollect(pl, now)
	if !flag {
		return
	}
	pl.Collect(cn)
	flag = true
	return
}

func (cn *CollecMiZangNPC) CollectInterrupt(pl scene.Player) {
	playerId := pl.GetId()
	mp := cn.collecMiZangObject.GetMiZangPlayer(playerId)
	if mp == nil {
		return
	}

	cn.collecMiZangObject.CollectInterrupt(pl)
	pl.ClearCollect()
}

func (cn *CollecMiZangNPC) collectFinish(pl scene.Player) {
	playerId := pl.GetId()
	mp := cn.collecMiZangObject.GetMiZangPlayer(playerId)
	flag := mp.Finish()
	if !flag {
		return
	}

	gameevent.Emit(collecteventtypes.EventTypeCollectMiZangFinish, pl, cn)
}

func (cn *CollecMiZangNPC) ExitScene(active bool) {
	for _, mp := range cn.collecMiZangObject.playerMap {
		//清除所有
		cn.collecMiZangObject.CollectInterrupt(mp.GetPlayer())
		mp.GetPlayer().ClearCollect()
	}
	cn.NPCBase.ExitScene(active)
}

//
func (cn *CollecMiZangNPC) IfMiZangCanCollect(pl scene.Player) bool {
	playerId := pl.GetId()
	mp := cn.collecMiZangObject.GetMiZangPlayer(playerId)
	if mp == nil {
		return false
	}
	return mp.IsFinish()
}

//密藏采集完成
func (cn *CollecMiZangNPC) MiZangCollectFinish(pl scene.Player) (success bool) {
	if !cn.IfMiZangCanCollect(pl) {
		return false
	}
	playerId := pl.GetId()
	flag := cn.NPCBase.Recycle(playerId)
	if !flag {
		panic(fmt.Errorf("mizang collect:  CostHP should be ok"))
	}

	for _, mp := range cn.collecMiZangObject.playerMap {
		//清除所有
		cn.collecMiZangObject.CollectInterrupt(mp.GetPlayer())
		mp.GetPlayer().ClearCollect()
	}
	success = true
	return
}

func (cn *CollecMiZangNPC) MiZangCollectGiveUp(pl scene.Player) (success bool) {
	if !cn.IfMiZangCanCollect(pl) {
		return false
	}

	//清除所有
	cn.collecMiZangObject.CollectInterrupt(pl)
	pl.ClearCollect()
	success = true
	return
}

func (cn *CollecMiZangNPC) GetParentId() int32 {
	return cn.parentId
}

func CreateCollectMiZangNPC(ownerType scenetypes.OwnerType, ownerId int64, ownerAllianceId int64, id int64, idInScene int32, biologyTemplate *gametemplate.BiologyTemplate, pos coretypes.Position, angle float64, miZangTemplate *gametemplate.BossMiZangTemplate, parentId int32) scene.NPC {
	n := &CollecMiZangNPC{}
	b := npc.NewNPCBase(n, ownerType, ownerId, ownerAllianceId, id, idInScene, biologyTemplate, pos, angle)

	obj := CreateCollecMiZangObject()
	n.NPCBase = b
	n.collecMiZangObject = obj
	n.parentId = parentId
	n.miZangTemplate = miZangTemplate
	n.Calculate()
	return n
}
