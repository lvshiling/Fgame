package npc

import (
	"fgame/fgame/game/common/common"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	npceventtypes "fgame/fgame/game/npc/event/types"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
	"math"
	"time"
)

const (
	recoverElapseTime = time.Second
)

type npcRecoverTask struct {
	n               scene.NPC
	lastRecoverTime int64
}

func (t *npcRecoverTask) Run() {
	now := global.GetGame().GetTimeService().Now()
	//死亡
	if t.n.IsDead() {
		t.lastRecoverTime = now
		return
	}
	//非初始化和非返回
	if t.n.CurrentState() != scene.NPCStateBack && t.n.CurrentState() != scene.NPCStateInit {
		t.lastRecoverTime = now
		return
	}

	//不能回复
	if !t.n.GetBiologyTemplate().IsRecover() {
		t.lastRecoverTime = now
		return
	}

	maxHp := t.n.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	//满血
	if t.n.GetHP() >= maxHp {
		t.lastRecoverTime = now
		return
	}
	elapseTime := now - t.lastRecoverTime
	if elapseTime < int64(t.n.GetBiologyTemplate().AutoRecoverTime) {
		return
	}

	recoverTimes := elapseTime / int64(t.n.GetBiologyTemplate().AutoRecoverTime)
	recover := int64(math.Ceil(float64(recoverTimes) * float64(t.n.GetBiologyTemplate().AutoRecoverNum) / float64(common.MAX_RATE) * float64(maxHp)))
	//回血
	t.n.AddHP(recover)
	//通知回血
	gameevent.Emit(npceventtypes.EventTypeNPCAutoRecover, t.n, recover)
	t.lastRecoverTime += int64(t.n.GetBiologyTemplate().AutoRecoverTime) * recoverTimes

}

func (t *npcRecoverTask) ElapseTime() time.Duration {
	return recoverElapseTime
}

func CreateNPCRecoverTask(n scene.NPC) *npcRecoverTask {
	t := &npcRecoverTask{
		n: n,
	}
	return t
}
