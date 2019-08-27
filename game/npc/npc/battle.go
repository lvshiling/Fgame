package npc

import (
	"fgame/fgame/core/aoi"
	coretypes "fgame/fgame/core/types"
	coreutils "fgame/fgame/core/utils"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	gameevent "fgame/fgame/game/event"
	npceventtypes "fgame/fgame/game/npc/event/types"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
)

//重写进入aoi
func (n *NPCBase) OnEnterAOI(other aoi.AOI) {
	n.SceneObjectBase.OnEnterAOI(other)
	bo, ok := other.(scene.BattleObject)
	if !ok {
		return
	}
	n.OnMove(bo, bo.GetPosition(), bo.GetAngle())
}

//重写退出aoi
func (n *NPCBase) OnLeaveAOI(other aoi.AOI, complete bool) {
	defer n.SceneObjectBase.OnLeaveAOI(other, complete)
	bo, ok := other.(scene.BattleObject)
	if !ok {
		return
	}

	n.exitBattle(bo)
	//退出场景清除伤害
	if complete {
		n.n.ClearDamage(other.GetId())
	}
}

//别的对象移动
func (n *NPCBase) OnMove(bo scene.BattleObject, pos coretypes.Position, angle float64) {
	n.enterBattle(bo)
}

//重生
func (n *NPCBase) OnReborn(other scene.BattleObject) {
	bo := other.(scene.BattleObject)
	n.enterBattle(bo)
}

//死亡
func (n *NPCBase) OnDead(other scene.BattleObject) {
	n.exitBattle(other)
}

//进入战斗
func (n *NPCBase) enterBattle(bo scene.BattleObject) {
	//友军
	if !n.IsEnemy(bo) {
		return
	}
	//敌人存在
	e := n.GetEnemy(bo)
	if e != nil {
		return
	}

	biologyTemplate := n.GetBiologyTemplate()

	//警戒距离小于0或者被动怪
	if biologyTemplate.GetAlertRadiusSquare() <= 0 || !biologyTemplate.IsPositive() {
		return
	}

	if coreutils.DistanceSquare(bo.GetPosition(), n.GetPosition()) <= biologyTemplate.GetAlertRadiusSquare() {
		//添加攻击目标
		e := scene.CreateDefaultEnemy(bo)
		n.AddEnemy(e)
	}
}

//退出战斗
func (n *NPCBase) exitBattle(bo scene.BattleObject) {
	//有可能已经变成友军
	e := n.GetEnemy(bo)
	if e == nil {
		return
	}
	n.RemoveEnemy(e)
	//清空攻击目标
	if e.BattleObject == n.GetAttackTarget() {
		n.SetAttackTarget(nil)
	}
}

func (n *NPCBase) clear() {
	n.ClearAllDamages()
	n.clearEnemy()
	n.resetEnemy()
}

func (n *NPCBase) clearEnemy() {
	enemies := n.GetEnemies()
	for _, obj := range enemies {
		n.exitBattle(obj.BattleObject)
	}
}

func (n *NPCBase) resetEnemy() {
	//遍历周围生物 进入战斗
	for _, obj := range n.GetNeighbors() {
		bo, ok := obj.(scene.BattleObject)
		if !ok {
			continue
		}
		if bo.IsDead() {
			continue
		}
		n.enterBattle(bo)
	}
}

//是否是敌人
func (n *NPCBase) IsEnemy(bo scene.BattleObject) bool {
	//灵童忽略
	switch bo.(type) {
	case scene.LingTong:
		return false
		break
	}

	if bo.GetFactionType() == scenetypes.FactionTypeModel {
		return false
	}

	if n.GetOwnerType() == scenetypes.OwnerTypeNone {
		return n.GetFactionType().IsEnemy(bo.GetFactionType())
	}

	switch tbo := bo.(type) {
	case scene.Player:
		if n.GetOwnerType() == scenetypes.OwnerTypeAlliance {
			if tbo.GetAllianceId() == n.GetOwnerId() {
				return false
			}
			return true
		}
		if n.GetOwnerType() == scenetypes.OwnerTypePlayer {
			if tbo.GetId() == n.GetOwnerId() {
				return false
			}
			if tbo.GetAllianceId() == 0 {
				return true
			}
			if tbo.GetAllianceId() == n.GetOwnerAllianceId() && tbo.GetPkState() == pktypes.PkStateBangPai {
				return false
			}
			return true
		}
		if n.GetOwnerType() == scenetypes.OwnerTypeCamp {
			ownerCamp := chuangshitypes.ChuangShiCampType(n.GetOwnerId())
			if tbo.GetCamp() == ownerCamp {
				return false
			}
			return true
		}
	case scene.NPC:
		if n.GetOwnerType() != tbo.GetOwnerType() {
			return false
		}
		return n.GetOwnerId() != tbo.GetOwnerId()
	}

	return false
}

func (n *NPCBase) GetSceneObjectSetType() scenetypes.BiologySetType {
	return n.GetBiologyTemplate().GetBiologySetType()
}

func (n *NPCBase) PlayerCalled(playerId int64) {
	n.clearEnemy()
	n.ownerType = scenetypes.OwnerTypePlayer
	n.ownerId = playerId
	n.resetEnemy()
	//TODO 重置所有敌人
	gameevent.Emit(npceventtypes.EventTypeNPCCampChanged, n, nil)
}

func (n *NPCBase) AllianceCalled(allianceId int64) {
	n.clearEnemy()
	n.ownerType = scenetypes.OwnerTypeAlliance
	n.ownerId = allianceId
	n.resetEnemy()
	//TODO 重置所有敌人
	gameevent.Emit(npceventtypes.EventTypeNPCCampChanged, n, nil)
}

func (n *NPCBase) ChuangShiCampCalled(campType chuangshitypes.ChuangShiCampType) {
	n.clearEnemy()
	n.ownerType = scenetypes.OwnerTypeCamp
	n.ownerId = int64(campType)
	n.resetEnemy()
	//TODO 重置所有敌人
	gameevent.Emit(npceventtypes.EventTypeNPCCampChanged, n, nil)
}

func (n *NPCBase) ResetOwner() {

	n.clearEnemy()
	n.ownerType = scenetypes.OwnerTypeNone
	n.SetFactionType(n.biologyTemplate.GetFactionType())
	n.ownerId = 0
	n.ownerAllianceId = 0
	n.resetEnemy()
	gameevent.Emit(npceventtypes.EventTypeNPCCampChanged, n, nil)
}

//获取主人类型
func (n *NPCBase) GetOwnerType() scenetypes.OwnerType {
	return n.ownerType
}

//获取主人
func (n *NPCBase) GetOwnerId() int64 {
	return n.ownerId
}

//获取主人
func (n *NPCBase) GetOwnerAllianceId() int64 {
	return n.ownerAllianceId
}
