package types

import (
	droptemplate "fgame/fgame/game/drop/template"
)

type SceneEventType string

const (
	//怪物击杀
	EventTypeMonsterKilled SceneEventType = "MonsterKilled"
	//怪物被击退
	EventTypeMonsterHurted = "MonsterHurted"
	//用户被用户击杀
	EventTypePlayerKilled = "PlayerKilled"
	//生物移动
	EventTypeBattleObjectMove = "BattleObjectMove"
	//生物死亡
	EventTypeBattleObjectDead = "BattleObjectDead"

	//生物重生
	EventTypeBattleObjectReborn = "BattleObjectReborn"
	//生物攻击
	EventTypeBattleObjectAttack = "BattleObjectAttack"
	//生物被攻击
	EventTypeBattleObjectAttacked = "BattleObjectAttacked"
	//生物被治疗
	EventTypeBattleObjectCure = "BattleObjectCure"
	//掉落
	EventTypeMonsterDrop = "MonsterDrop"
	//掉落
	EventTypeBattleObjectDrop = "BattleObjectDrop"
	//自定义掉落
	EventTypeBattleCustomDrop = "BattleCustomDrop"
	//跨服击杀怪物
	EventTypeCrossKillBiology = "CrossKillBiology"
	//掉落直接入背包
	EventTypeBattleObjectDropIntoInventory = "BattleObjectDropIntoInventory"
	//任务机器人检测
	EventTypeSceneRobotCheck = "SceneRobotCheck"
	//场景排行变化
	EventTypeSceneRankChanged = "SceneRankChanged"
	//场景完成
	EventTypeSceneFinish = "SceneFinish"

	//场景日志
	EventTypeSceneInfoLog = "SceneInfoLog"
)

type MonsterDropData struct {
	attackId int64
	num      int32
}

func (d *MonsterDropData) GetAttackId() int64 {
	return d.attackId
}

func (d *MonsterDropData) GetNum() int32 {
	return d.num
}

func CreateMonsterDropData(attackId int64, num int32) *MonsterDropData {
	d := &MonsterDropData{
		attackId: attackId,
		num:      num,
	}
	return d
}

type BattleObjectDropIntoInventoryData struct {
	owerId       int64
	itemDataList []*droptemplate.DropItemData
}

func (d *BattleObjectDropIntoInventoryData) GetOwerId() int64 {
	return d.owerId
}

func (d *BattleObjectDropIntoInventoryData) GetItemDataList() []*droptemplate.DropItemData {
	return d.itemDataList
}

func CreateBattleObjectDropIntoInventoryData(owerId int64, itemDataList []*droptemplate.DropItemData) *BattleObjectDropIntoInventoryData {
	d := &BattleObjectDropIntoInventoryData{
		owerId:       owerId,
		itemDataList: itemDataList,
	}
	return d
}
