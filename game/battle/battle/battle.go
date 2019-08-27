package battle

import (
	"fgame/fgame/core/heartbeat"
	coretypes "fgame/fgame/core/types"
	battlecommon "fgame/fgame/game/battle/common"
	battleeventtypes "fgame/fgame/game/battle/event/types"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
)

//战斗数据管理
type BattleManager struct {
	bo scene.BattleObject
	//阵营
	factionType scenetypes.FactionType
	maxAttackId int64
	//添加伤害计算
	damageMap map[int64]int64
	//敌人
	enemies             map[int64]*scene.Enemy
	attackTarget        scene.BattleObject
	defaultAttackTarget scene.BattleObject
	foreverAttackTarget scene.BattleObject
}

func (m *BattleManager) GetFactionType() scenetypes.FactionType {
	return m.factionType
}

func (m *BattleManager) SetFactionType(faction scenetypes.FactionType) {
	m.factionType = faction
}

func (m *BattleManager) AddDamage(attackId int64, damage int64) {
	m.damageMap[attackId] = m.damageMap[attackId] + damage
	m.checkMaxDamage(attackId)
}

func (m *BattleManager) checkMaxDamage(attackId int64) {
	if attackId == m.maxAttackId {
		return
	}
	if m.damageMap[attackId] > m.damageMap[m.maxAttackId] {
		originMaxAttackId := m.maxAttackId
		m.maxAttackId = attackId
		eventData := battleeventtypes.CreateBattleObjectMaxDamageChangedEventData(originMaxAttackId, attackId)
		gameevent.Emit(battleeventtypes.EventTypeBattleObjectMaxDamageChanged, m.bo, eventData)
	}
}

func (m *BattleManager) refreshMaxDamage() {
	maxDamage := int64(0)
	maxAttackId := int64(0)
	for attackId, damage := range m.damageMap {
		if damage >= maxDamage {
			maxDamage = damage
			maxAttackId = attackId
		}
	}
	originMaxAttackId := m.maxAttackId
	m.maxAttackId = maxAttackId
	eventData := battleeventtypes.CreateBattleObjectMaxDamageChangedEventData(originMaxAttackId, maxAttackId)
	gameevent.Emit(battleeventtypes.EventTypeBattleObjectMaxDamageChanged, m.bo, eventData)

}

func (m *BattleManager) ClearAllDamages() {
	for attackId, _ := range m.damageMap {
		delete(m.damageMap, attackId)
	}
	m.refreshMaxDamage()
}
func (m *BattleManager) ClearDamage(attackId int64) {
	delete(m.damageMap, attackId)
	m.refreshMaxDamage()
}

func (m *BattleManager) GetAllDamages() map[int64]int64 {
	return m.damageMap
}

//添加仇恨值
func (m *BattleManager) AddHate(target scene.BattleObject, val int) {
	if val <= 0 {
		panic(fmt.Errorf("battle: 添加仇恨值[%d]不大于0", val))
	}
	//判断是否在视野范围内
	if !m.bo.IsNeighbor(target.GetId()) {
		return
	}

	enemy, ok := m.enemies[target.GetId()]
	if !ok {
		enemy = scene.CreateEnemy(target, val)
		m.enemies[target.GetId()] = enemy
	}

	enemy.AddHate(val)
	gameevent.Emit(battleeventtypes.EventTypeBattleObjectHateChanged, m.bo, enemy)
	return
}

func (n *BattleManager) GetHate(target scene.BattleObject) int {
	enemy, ok := n.enemies[target.GetId()]
	if !ok {
		return 0
	}
	return enemy.GetHate()
}

//添加敌人
func (m *BattleManager) AddEnemy(e *scene.Enemy) bool {
	_, ok := m.enemies[e.BattleObject.GetId()]
	if ok {
		return false
	}
	m.enemies[e.BattleObject.GetId()] = e

	return true
}

func (m *BattleManager) RemoveEnemy(e *scene.Enemy) bool {
	_, ok := m.enemies[e.BattleObject.GetId()]
	if !ok {
		return false
	}
	delete(m.enemies, e.BattleObject.GetId())

	return true
}

func (m *BattleManager) GetEnemy(bo scene.BattleObject) *scene.Enemy {
	e, ok := m.enemies[bo.GetId()]
	if !ok {
		return nil
	}
	return e
}

//查找最仇恨的
func (m *BattleManager) FindHatestEnemy() (e *scene.Enemy) {
	maxHate := 0
	for _, te := range m.enemies {
		if te.GetHate() > maxHate {
			e = te
			maxHate = te.GetHate()
		}
	}
	return
}

func (m *BattleManager) GetEnemies() map[int64]*scene.Enemy {
	return m.enemies
}

func (m *BattleManager) SetAttackTarget(bo scene.BattleObject) {
	m.attackTarget = bo
}

func (m *BattleManager) GetAttackTarget() scene.BattleObject {
	return m.attackTarget
}

func (m *BattleManager) SetDefaultAttackTarget(bo scene.BattleObject) {
	m.defaultAttackTarget = bo
}

func (m *BattleManager) GetDefaultAttackTarget() scene.BattleObject {
	return m.defaultAttackTarget
}

func (m *BattleManager) SetForeverAttackTarget(bo scene.BattleObject) {
	m.foreverAttackTarget = bo
}

func (m *BattleManager) GetForeverAttackTarget() scene.BattleObject {
	return m.foreverAttackTarget
}

//死亡
func (m *BattleManager) AttackedMove(pos coretypes.Position, angle float64, moveSpeed float64, stopTime float64) {

}

//进入战斗
func (m *BattleManager) EnterBattle(bo scene.BattleObject) {

	//友军
	if !m.bo.IsEnemy(bo) {
		return
	}
	//敌人存在
	e := m.bo.GetEnemy(bo)
	if e != nil {
		return
	}

	//添加攻击目标
	e = scene.CreateDefaultEnemy(bo)
	m.bo.AddEnemy(e)
}

//退出战斗
func (m *BattleManager) ExitBattle(bo scene.BattleObject) {
	//有可能已经变成友军
	e := m.bo.GetEnemy(bo)
	if e == nil {
		return
	}
	m.bo.RemoveEnemy(e)
	//清空攻击目标
	if e.BattleObject == m.bo.GetAttackTarget() {
		m.bo.SetAttackTarget(nil)
	}
}

func (m *BattleManager) ResetEnemy() {

	for _, obj := range m.enemies {
		m.ExitBattle(obj.BattleObject)
	}
	//遍历周围生物 进入战斗
	for _, obj := range m.bo.GetNeighbors() {
		bo, ok := obj.(scene.BattleObject)
		if !ok {
			continue
		}
		if bo.IsDead() {
			continue
		}
		m.EnterBattle(bo)
	}
}

//战斗管理器
func CreateBattleManager(bo scene.BattleObject, factionType scenetypes.FactionType) *BattleManager {
	m := &BattleManager{
		bo: bo,
	}

	m.factionType = factionType
	m.damageMap = make(map[int64]int64)
	m.enemies = make(map[int64]*scene.Enemy)
	return m
}

//TODO 分开
//战斗管理器
type PlayerBattleManager struct {
	*BattleManager
	p scene.Player
	//是否是机器人
	robot bool
	//是否进入战斗
	battleTime int64
	//是否处于战斗状态
	isBattle bool
	//是否进入pvp
	pvpBattleTime int64
	//是否处于pvp
	isPvpBattle bool
	//上次脱离卡死时间
	lastExitKaSiTime int64
	//帝魂觉醒数量
	soulAwakenNum int32
	//等级
	level int32
	//转生
	zhuanSheng int32
	//vip
	vip int32
	//是否永久会员
	isHuiYuan bool

	hbRunner heartbeat.HeartbeatTaskRunner
}

func (m *PlayerBattleManager) IsRobot() bool {
	return m.robot
}

func (m *PlayerBattleManager) Heartbeat() {
	if m.isBattle {
		now := global.GetGame().GetTimeService().Now()
		outOfBattle := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeExitBattle)
		if now-m.battleTime > int64(outOfBattle) {
			m.ClearBattle()
		}
	}

	if m.isPvpBattle {
		now := global.GetGame().GetTimeService().Now()
		outOfBattle := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeExitBattle)
		if now-m.pvpBattleTime > int64(outOfBattle) {
			m.ClearPvpBattle()
		}
	}
	m.hbRunner.Heartbeat()
	//获取周围物体变更
	gameevent.Emit(battleeventtypes.EventTypePlayerSyncNeighbor, m.p, nil)

}

func (m *PlayerBattleManager) Battle() bool {
	m.battleTime = global.GetGame().GetTimeService().Now()
	if m.isBattle {
		return false
	}
	m.isBattle = true
	gameevent.Emit(battleeventtypes.EventTypePlayerEnterBattle, m.p, nil)
	return true
}

func (m *PlayerBattleManager) IsBattle() bool {
	return m.isBattle
}

func (m *PlayerBattleManager) ClearBattle() {
	if !m.isBattle {
		return
	}
	m.isBattle = false
	gameevent.Emit(battleeventtypes.EventTypePlayerExitBattle, m.p, nil)
}

func (m *PlayerBattleManager) PvpBattle() bool {
	m.pvpBattleTime = global.GetGame().GetTimeService().Now()
	if m.isPvpBattle {
		return false
	}
	m.isPvpBattle = true
	gameevent.Emit(battleeventtypes.EventTypePlayerEnterPVP, m.p, nil)
	return true
}

func (m *PlayerBattleManager) IsPvpBattle() bool {
	return m.isPvpBattle
}

func (m *PlayerBattleManager) ClearPvpBattle() {
	if !m.isPvpBattle {
		return
	}
	m.isPvpBattle = false
	gameevent.Emit(battleeventtypes.EventTypePlayerExitPVP, m.p, nil)
}

//是否是敌人
func (m *PlayerBattleManager) IsEnemy(bo scene.BattleObject) bool {
	//灵童忽略
	_, ok := bo.(scene.LingTong)
	if ok {
		return false
	}

	targetP, ok := bo.(scene.Player)
	if ok {
		if m.p.GetFactionType() == scenetypes.FactionTypeModel {
			return false
		}

		if m.p == targetP {
			return false
		}
		if m.p.GetPkState() == pktypes.PkStateAll {

			return true
		}
		switch m.p.GetPkState() {
		case pktypes.PkStateAll:
			return true
		case pktypes.PkStateBangPai:

			return m.p.GetAllianceId() != targetP.GetAllianceId()
		case pktypes.PkStateGroup:

			return m.p.GetTeamId() != targetP.GetTeamId()
		case pktypes.PkStateCamp:

			return m.p.GetPkCamp().Camp() != targetP.GetPkCamp().Camp()
		}
		return false
	}
	npc, ok := bo.(scene.NPC)
	if !ok {
		return false
	}
	switch npc.GetOwnerType() {
	//TODO 修改
	case scenetypes.OwnerTypeNone:
		return m.p.GetFactionType().IsEnemy(bo.GetFactionType())
	case scenetypes.OwnerTypePlayer:
		if m.p.GetId() == npc.GetOwnerId() {
			return false
		}
		if m.p.GetAllianceId() == 0 {
			return true
		}
		if m.p.GetAllianceId() == npc.GetOwnerAllianceId() && m.p.GetPkState() == pktypes.PkStateBangPai {
			return false
		}
		return true
	case scenetypes.OwnerTypeAlliance:
		if m.p.GetAllianceId() == npc.GetOwnerId() {
			return false
		}
		return true

	case scenetypes.OwnerTypeCamp:
		{
			ownerCamp := chuangshitypes.ChuangShiCampType(npc.GetOwnerId())
			if m.p.GetCamp() == ownerCamp {
				return false
			}
			return true
		}
	}
	return false
}

func (m *PlayerBattleManager) GetSceneObjectSetType() scenetypes.BiologySetType {
	return scenetypes.BiologySetTypePlayer
}

//退出卡死cd
const (
	exitKaSiTime = int64(common.MINUTE * 10)
)

func (m *PlayerBattleManager) IfCanExitKaSi() bool {
	now := global.GetGame().GetTimeService().Now()
	return (now - m.lastExitKaSiTime) >= exitKaSiTime
}

func (m *PlayerBattleManager) ExitKaSiLeftTime() int64 {
	now := global.GetGame().GetTimeService().Now()
	elapse := now - m.lastExitKaSiTime
	return exitKaSiTime - elapse
}

func (m *PlayerBattleManager) ExitKaSi() {
	now := global.GetGame().GetTimeService().Now()
	m.lastExitKaSiTime = now
}

func (m *PlayerBattleManager) IfCanGetDropItem(di scene.DropItem) bool {
	ownerId := di.GetOwnerId()
	if ownerId == 0 {
		return true
	}
	switch di.GetOwnerType() {
	case scenetypes.DropOwnerTypeAlliance:
		return m.p.GetAllianceId() == ownerId

	case scenetypes.DropOwnerTypePlayer:
		return m.p.GetId() == ownerId

	case scenetypes.DropOwnerTypeTeam:
		return m.p.GetTeamId() == ownerId
	}
	return false
}

func (m *PlayerBattleManager) SyncSoulAwakenNum(soulAwakenNum int32) {
	m.soulAwakenNum = soulAwakenNum
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerSoulAwakenChanged, m.p, nil)
}

func (m *PlayerBattleManager) GetSoulAwakenNum() int32 {
	return m.soulAwakenNum
}

func (m *PlayerBattleManager) SyncLevel(level int32) {
	m.level = level
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerLevelChanged, m.p, nil)
}

func (m *PlayerBattleManager) GetLevel() int32 {
	return m.level
}

func (m *PlayerBattleManager) SyncZhuanSheng(zhuanSheng int32) {
	m.zhuanSheng = zhuanSheng
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerZhuanShengChanged, m.p, nil)
}

func (m *PlayerBattleManager) GetZhuanSheng() int32 {
	return m.zhuanSheng
}

func (m *PlayerBattleManager) SyncVip(vip int32) {
	m.vip = vip
	gameevent.Emit(battleeventtypes.EventTypeBattlePlayerVipChanged, m.p, nil)
}

func (m *PlayerBattleManager) GetVip() int32 {
	return m.vip
}

func (m *PlayerBattleManager) SyncHuiYuan(isHuiYuan bool) {
	m.isHuiYuan = isHuiYuan
}

func (m *PlayerBattleManager) IsHuiYuanPlus() bool {
	return m.isHuiYuan
}

//战斗管理器
func CreatePlayerBattleManager(
	p scene.Player,
	robot bool,

) *PlayerBattleManager {

	return CreatePlayerBattleManagerWithObject(p, robot, nil, scenetypes.FactionTypePlayer)
}

func CreatePlayerBattleManagerWithObject(
	p scene.Player,
	robot bool,
	obj *battlecommon.PlayerBattleObject,
	factionType scenetypes.FactionType,
) *PlayerBattleManager {
	m := &PlayerBattleManager{
		p: p,
	}
	m.BattleManager = CreateBattleManager(p, factionType)
	m.robot = robot
	if obj != nil {
		m.soulAwakenNum = obj.GetSoulAwakenNum()
		m.level = obj.GetLevel()
		m.vip = obj.GetVip()
		m.zhuanSheng = obj.GetZhuanSheng()
		m.isHuiYuan = obj.GetIsHuiYuan()
	}
	m.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	m.hbRunner.AddTask(CreateMountTask(p))
	m.hbRunner.AddTask(CreatePkTask(p))

	return m
}
