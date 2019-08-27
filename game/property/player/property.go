package player

import (
	commonlog "fgame/fgame/common/log"
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/core/storage"
	"fgame/fgame/core/template"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constantypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	goldequiptemplate "fgame/fgame/game/goldequip/template"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	propertycommon "fgame/fgame/game/property/common"
	"fgame/fgame/game/property/dao"
	gameentity "fgame/fgame/game/property/entity"
	propertyeventtypes "fgame/fgame/game/property/event/types"
	propertyplayerproperty "fgame/fgame/game/property/player/property"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	propertytypes "fgame/fgame/game/property/types"
	gametemplate "fgame/fgame/game/template"
	tianshutypes "fgame/fgame/game/tianshu/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
	"math"

	"github.com/pkg/errors"
)

//玩家属性数据
type PlayerPropertyObject struct {
	player   player.Player
	Id       int64
	PlayerId int64
	//等级
	Level      int32
	Exp        int64
	Silver     int64
	Gold       int64
	BindGold   int64
	Evil       int32
	ZhuanSheng int32
	//当前血量
	CurrentHP int64
	//当前体力
	CurrentTP     int64
	Power         int64
	Charm         int32
	GoldYuanLevel int32
	GoldYuanExp   int64
	UpdateTime    int64
	CreateTime    int64
	DeleteTime    int64
}

func NewPlayerPropertyObject(pl player.Player) *PlayerPropertyObject {
	pso := &PlayerPropertyObject{
		player: pl,
	}
	return pso
}

func convertPlayerPropertyObjectToEntity(ppo *PlayerPropertyObject) *gameentity.PlayerPropertyEntity {
	e := &gameentity.PlayerPropertyEntity{
		Id:            ppo.Id,
		PlayerId:      ppo.PlayerId,
		Level:         ppo.Level,
		Exp:           ppo.Exp,
		Silver:        ppo.Silver,
		Gold:          ppo.Gold,
		BindGold:      ppo.BindGold,
		Evil:          ppo.Evil,
		ZhuanSheng:    ppo.ZhuanSheng,
		CurrentHP:     ppo.CurrentHP,
		CurrentTP:     ppo.CurrentTP,
		Power:         ppo.Power,
		Charm:         ppo.Charm,
		GoldYuanLevel: ppo.GoldYuanLevel,
		GoldYuanExp:   ppo.GoldYuanExp,
		UpdateTime:    ppo.UpdateTime,
		CreateTime:    ppo.CreateTime,
		DeleteTime:    ppo.DeleteTime,
	}
	return e
}

func (ppo *PlayerPropertyObject) GetPlayerId() int64 {
	return ppo.PlayerId
}

func (ppo *PlayerPropertyObject) GetDBId() int64 {
	return ppo.Id
}

func (ppo *PlayerPropertyObject) ToEntity() (e storage.Entity, err error) {
	e = convertPlayerPropertyObjectToEntity(ppo)
	return
}

func (ppo *PlayerPropertyObject) FromEntity(e storage.Entity) (err error) {
	ppe, _ := e.(*gameentity.PlayerPropertyEntity)
	ppo.Id = ppe.Id
	ppo.Level = ppe.Level
	ppo.Exp = ppe.Exp
	ppo.Silver = ppe.Silver
	ppo.Gold = ppe.Gold
	ppo.BindGold = ppe.BindGold
	ppo.Evil = ppe.Evil
	ppo.PlayerId = ppe.PlayerId
	ppo.ZhuanSheng = ppe.ZhuanSheng
	ppo.CurrentHP = ppe.CurrentHP
	ppo.CurrentTP = ppe.CurrentTP
	ppo.Power = ppe.Power
	ppo.Charm = ppe.Charm
	ppo.GoldYuanLevel = ppe.GoldYuanLevel
	ppo.GoldYuanExp = ppe.GoldYuanExp
	ppo.UpdateTime = ppe.UpdateTime
	ppo.CreateTime = ppe.CreateTime
	ppo.DeleteTime = ppe.DeleteTime
	return
}

func (ppe *PlayerPropertyObject) SetModified() {
	e, err := ppe.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "Property"))
	}
	obj, _ := e.(types.PlayerDataEntity)
	if obj == nil {
		panic("never reach here")
	}
	ppe.player.AddChangedObject(obj)
	return
}

//玩家属性数据
type PlayerPowerRecordObject struct {
	player         player.Player
	Id             int64
	TodayInitPower int64
	HisMaxPower    int64
	UpdateTime     int64
	CreateTime     int64
	DeleteTime     int64
}

func NewPlayerPowerRecordObject(pl player.Player) *PlayerPowerRecordObject {
	pso := &PlayerPowerRecordObject{
		player: pl,
	}
	return pso
}

func convertPlayerPowerRecordObjectToEntity(ppo *PlayerPowerRecordObject) *gameentity.PlayerPowerRecordEntity {
	e := &gameentity.PlayerPowerRecordEntity{
		Id:             ppo.Id,
		PlayerId:       ppo.player.GetId(),
		HisMaxPower:    ppo.HisMaxPower,
		TodayInitPower: ppo.TodayInitPower,
		UpdateTime:     ppo.UpdateTime,
		CreateTime:     ppo.CreateTime,
		DeleteTime:     ppo.DeleteTime,
	}
	return e
}

func (ppo *PlayerPowerRecordObject) GetPlayerId() int64 {
	return ppo.player.GetId()
}

func (ppo *PlayerPowerRecordObject) GetDBId() int64 {
	return ppo.Id
}

func (ppo *PlayerPowerRecordObject) ToEntity() (e storage.Entity, err error) {
	e = convertPlayerPowerRecordObjectToEntity(ppo)
	return
}

func (ppo *PlayerPowerRecordObject) FromEntity(e storage.Entity) (err error) {
	ppe, _ := e.(*gameentity.PlayerPowerRecordEntity)
	ppo.Id = ppe.Id
	ppo.HisMaxPower = ppe.HisMaxPower
	ppo.TodayInitPower = ppe.TodayInitPower
	ppo.UpdateTime = ppe.UpdateTime
	ppo.CreateTime = ppe.CreateTime
	ppo.DeleteTime = ppe.DeleteTime
	return
}

func (ppe *PlayerPowerRecordObject) SetModified() {
	e, err := ppe.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "PowerRecord"))
	}
	obj, _ := e.(types.PlayerDataEntity)
	if obj == nil {
		panic("never reach here")
	}
	ppe.player.AddChangedObject(obj)
	return
}

//每日消费记录对象
type PlayerCycleCostRecordObject struct {
	player        player.Player
	id            int64
	costNum       int64
	preDayCostNum int64
	updateTime    int64
	createTime    int64
	deleteTime    int64
}

func NewPlayerCycleCostRecordObject(pl player.Player) *PlayerCycleCostRecordObject {
	o := &PlayerCycleCostRecordObject{
		player: pl,
	}
	return o
}

func convertNewPlayerCycleChargeRecordObjectToEntity(o *PlayerCycleCostRecordObject) (*gameentity.PlayerCycleCostRecordEntity, error) {
	e := &gameentity.PlayerCycleCostRecordEntity{
		Id:            o.id,
		PlayerId:      o.GetPlayerId(),
		CostNum:       o.costNum,
		PreDayCostNum: o.preDayCostNum,
		UpdateTime:    o.updateTime,
		CreateTime:    o.createTime,
		DeleteTime:    o.deleteTime,
	}
	return e, nil
}

func (o *PlayerCycleCostRecordObject) GetPlayerId() int64 {
	return o.player.GetId()
}

func (o *PlayerCycleCostRecordObject) GetDBId() int64 {
	return o.id
}

func (o *PlayerCycleCostRecordObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewPlayerCycleChargeRecordObjectToEntity(o)
	return e, err
}

func (o *PlayerCycleCostRecordObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*gameentity.PlayerCycleCostRecordEntity)

	o.id = pse.Id
	o.costNum = pse.CostNum
	o.preDayCostNum = pse.PreDayCostNum
	o.updateTime = pse.UpdateTime
	o.createTime = pse.CreateTime
	o.deleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerCycleCostRecordObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "CycleCostRecord"))
	}
	obj, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.player.AddChangedObject(obj)
	return
}

//所有魅力日志数据
type PlayerCharmAddLogObject struct {
	Id         int64
	p          player.Player
	Charm      int32
	SendId     int64
	UpdateTime int64
	CreateTime int64
	DeleteTime int64
}

func NewPlayerCharmAddLogObject(p player.Player) *PlayerCharmAddLogObject {
	o := &PlayerCharmAddLogObject{
		p: p,
	}
	return o
}

func convertNewCharmAddLogObjectToEntity(o *PlayerCharmAddLogObject) (*gameentity.PlayerCharmAddLogEntity, error) {
	e := &gameentity.PlayerCharmAddLogEntity{
		Id:         o.Id,
		PlayerId:   o.p.GetId(),
		Charm:      o.Charm,
		SendId:     o.SendId,
		UpdateTime: o.UpdateTime,
		CreateTime: o.CreateTime,
		DeleteTime: o.DeleteTime,
	}
	return e, nil
}

func (o *PlayerCharmAddLogObject) GetDBId() int64 {
	return o.Id
}

func (o *PlayerCharmAddLogObject) GetPlayerId() int64 {
	return o.p.GetId()
}

func (o *PlayerCharmAddLogObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertNewCharmAddLogObjectToEntity(o)
	return e, err
}

func (o *PlayerCharmAddLogObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*gameentity.PlayerCharmAddLogEntity)

	o.Id = pse.Id
	o.SendId = pse.SendId
	o.Charm = pse.Charm
	o.UpdateTime = pse.UpdateTime
	o.CreateTime = pse.CreateTime
	o.DeleteTime = pse.DeleteTime
	return nil
}

func (o *PlayerCharmAddLogObject) SetModified() {
	e, err := o.ToEntity()
	if err != nil {
		panic(errors.Wrap(err, "CharmAddLogObject"))
	}

	pe, ok := e.(types.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	o.p.AddChangedObject(pe)

	return
}

//玩家属性数据管理器
type PlayerPropertyDataManager struct {
	p                    player.Player
	playerPropertyObject *PlayerPropertyObject
	//基础属性
	baseProperty *propertycommon.BasePropertySegment
	//战斗属性组
	battlePropertyGroup *propertycommon.BattlePropertyGroup
	//每日消费记录
	playerCycleCostRecord *PlayerCycleCostRecordObject
	//战力变化记录
	playerPowerRecordObj *PlayerPowerRecordObject
	//玩家魅力增加日志数据
	charmAddLogList []*PlayerCharmAddLogObject
}

//是否属性变化过
func (ppdm *PlayerPropertyDataManager) IsChanged() bool {
	return ppdm.baseProperty.IsChanged() || ppdm.battlePropertyGroup.IsChanged()
}

//重置改变标记位
func (ppdm *PlayerPropertyDataManager) resetChanged() {
	ppdm.baseProperty.ResetChanged()
	ppdm.battlePropertyGroup.ResetChanged()
}

func (ppdm *PlayerPropertyDataManager) GetChangedTypesAndReset() (baseChanged map[int32]int64, battleChanged map[int32]int64) {
	baseChanged = ppdm.baseProperty.GetChangedTypes()
	battleChanged = ppdm.battlePropertyGroup.GetChangedTypes()
	ppdm.resetChanged()
	return
}

func (ppdm *PlayerPropertyDataManager) GetChangedBattlePropertiesAndReset() (battleChanged map[int32]int64) {
	battleChanged = ppdm.battlePropertyGroup.GetChangedTypes()
	ppdm.battlePropertyGroup.ResetChanged()
	return
}

func (ppdm *PlayerPropertyDataManager) GetChangedBasePropertiesAndReset() (baseChanged map[int32]int64) {
	baseChanged = ppdm.baseProperty.GetChangedTypes()
	ppdm.baseProperty.ResetChanged()
	return
}

//获得hp
func (ppdm *PlayerPropertyDataManager) GetHP() int64 {
	return ppdm.playerPropertyObject.CurrentHP
}

//是否有足够的钱
func (ppdm *PlayerPropertyDataManager) HasEnoughGold(gold int64, includeBind bool) bool {
	currentGold := ppdm.GetGold()
	if includeBind {
		currentGold += ppdm.GetBindGlod()
	}
	if currentGold < gold {
		return false
	}
	return true
}

//花费
func (ppdm *PlayerPropertyDataManager) Cost(bindGold int64, gold int64, goldReason commonlog.GoldLogReason, goldReasonText string, silver int64, silverReason commonlog.SilverLogReason, silverReasonText string) bool {
	if !ppdm.HasEnoughCost(bindGold, gold, silver) {
		return false
	}

	if silver != 0 {
		flag := ppdm.CostSilver(silver, silverReason, silverReasonText)
		if !flag {
			panic(fmt.Errorf("property:扣除银两[%d]应该成功", silver))
		}
	}
	if gold != 0 {
		flag := ppdm.CostGold(gold, false, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("property:扣除元宝[%d]应该成功", gold))
		}
	}
	if bindGold != 0 {
		flag := ppdm.CostGold(bindGold, true, goldReason, goldReasonText)
		if !flag {
			panic(fmt.Errorf("property:扣除绑定元宝[%d]应该成功", bindGold))
		}
	}
	return true
}

//是否有足够的消耗
func (ppdm *PlayerPropertyDataManager) HasEnoughCost(bindGold int64, gold int64, silver int64) bool {
	if !ppdm.HasEnoughSilver(int64(silver)) {
		return false
	}

	//判断有没有足够的元宝
	if !ppdm.HasEnoughGold(gold, false) {
		return false
	}

	if !ppdm.HasEnoughGold(gold+bindGold, true) {
		return false
	}
	return true
}

//增加奖励数据
func (ppdm *PlayerPropertyDataManager) AddRewData(rewData *propertytypes.RewData, goldReason commonlog.GoldLogReason, goldReasonText string, silverReason commonlog.SilverLogReason, silverReasonText string, levelReason commonlog.LevelLogReason, levelReasonText string) bool {
	rewExp := int64(rewData.GetRewExp())
	rewExpPoint := int64(rewData.GetRewExpPoint())
	rewSilver := int64(rewData.GetRewSilver())
	rewGold := int64(rewData.GetRewGold())
	rewBindGold := int64(rewData.GetRewBindGold())

	flag := ppdm.AddMoney(rewBindGold, rewGold, goldReason, goldReasonText, rewSilver, silverReason, silverReasonText)
	if !flag {
		panic(fmt.Errorf("property: AddRewData 调用 AddMoney 应该是ok的"))
	}

	if rewExpPoint > 0 {
		ppdm.AddExpPoint(rewExpPoint, levelReason, levelReasonText)
	}
	if rewExp > 0 {
		ppdm.AddExp(rewExp, levelReason, levelReasonText)
	}

	return true
}

//增加钱
func (ppdm *PlayerPropertyDataManager) AddMoney(bindGold int64, gold int64, goldReason commonlog.GoldLogReason, goldReasonText string, silver int64, silverReason commonlog.SilverLogReason, silverReasonText string) bool {
	if silver != 0 {
		ppdm.AddSilver(silver, silverReason, silverReasonText)
	}
	if gold != 0 {
		ppdm.AddGold(gold, false, goldReason, goldReasonText)
	}
	if bindGold != 0 {
		ppdm.AddGold(bindGold, true, goldReason, goldReasonText)
	}
	return true
}

const (
	MAX_GOLD = 500000
)

//获得元宝
func (ppdm *PlayerPropertyDataManager) AddGold(gold int64, bindGold bool, reason commonlogtypes.LogReason, reasonText string) {
	if gold <= 0 {
		return
		//panic(fmt.Errorf("property:添加元宝应该应该大于0"))
	}
	if gold >= MAX_GOLD {
		panic(fmt.Errorf("property:添加元宝应该应该小于%d", MAX_GOLD))
	}

	beforeGold := ppdm.GetGold()
	beforeBindGold := ppdm.GetBindGlod()
	if !bindGold {
		currentGold := ppdm.GetGold()
		currentGold += gold
		ppdm.setGold(currentGold)
		now := global.GetGame().GetTimeService().Now()
		ppdm.playerPropertyObject.UpdateTime = now
		ppdm.playerPropertyObject.SetModified()
		newGoldLogEventData := propertyeventtypes.CreatePlayerNewGoldChangedLogEventData(beforeGold, gold, reason, reasonText)
		gameevent.Emit(propertyeventtypes.EventTypePlayerNewGoldChangedLog, ppdm.p, newGoldLogEventData)
	} else {
		currentBindGold := ppdm.GetBindGlod()
		currentBindGold += gold
		ppdm.setBindGold(currentBindGold)
		now := global.GetGame().GetTimeService().Now()
		ppdm.playerPropertyObject.UpdateTime = now
		ppdm.playerPropertyObject.SetModified()
		newBindGoldLogEventData := propertyeventtypes.CreatePlayerNewBindGoldChangedLogEventData(beforeBindGold, gold, reason, reasonText)
		gameevent.Emit(propertyeventtypes.EventTypePlayerNewBindGoldChangedLog, ppdm.p, newBindGoldLogEventData)
	}

	data := propertyeventtypes.CreatePlayerGoldChangedLogEventData(beforeGold, beforeBindGold, gold, reason, reasonText)
	gameevent.Emit(propertyeventtypes.EventTypePlayerGoldChangedLog, ppdm.p, data)
}

//消耗元宝
func (ppdm *PlayerPropertyDataManager) CostGold(gold int64, includeBind bool, reason commonlog.GoldLogReason, reasonText string) bool {
	if gold < 0 {
		panic(fmt.Errorf("property:消耗元宝应该不能小于0"))
	}
	if gold == 0 {
		return true
	}
	if !ppdm.HasEnoughGold(gold, includeBind) {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	remainGold := gold
	beforeBindGold := ppdm.GetBindGlod()
	if includeBind {
		bindGold := ppdm.GetBindGlod()
		changed := int64(0)
		if bindGold >= remainGold {
			bindGold -= remainGold
			changed = remainGold
			remainGold = 0
		} else {
			changed = bindGold
			remainGold -= bindGold
			bindGold = 0
		}
		ppdm.setBindGold(bindGold)
		ppdm.playerPropertyObject.UpdateTime = now
		ppdm.playerPropertyObject.SetModified()
		if changed > 0 {
			newBindGoldLogEventData := propertyeventtypes.CreatePlayerNewBindGoldChangedLogEventData(beforeBindGold, -changed, reason, reasonText)
			gameevent.Emit(propertyeventtypes.EventTypePlayerNewBindGoldChangedLog, ppdm.p, newBindGoldLogEventData)
		}
	}

	beforeGold := ppdm.GetGold()
	if remainGold > 0 {
		//zrc:特殊处理 预定婚礼不算元宝消费,等到举行婚礼才算,求婚的时候也不算，等到答应的时候才算
		if reason != commonlog.GoldLogReasonMarryWeddingGrade &&
			reason != commonlog.GoldLogReasonMarryProposal &&
			reason != commonlog.GoldLogReasonTrade {
			//消费元宝
			ppdm.UpdateCycleCostInfo(remainGold)
			gameevent.Emit(propertyeventtypes.EventTypePlayerGoldCost, ppdm.p, remainGold)
		}
		currentGold := ppdm.GetGold()
		currentGold -= remainGold
		if currentGold < 0 {
			panic(fmt.Errorf("property: 当前元宝不应该为负数"))
		}
		ppdm.setGold(currentGold)
		ppdm.playerPropertyObject.UpdateTime = now
		ppdm.playerPropertyObject.SetModified()

		newGoldLogEventData := propertyeventtypes.CreatePlayerNewGoldChangedLogEventData(beforeGold, -remainGold, reason, reasonText)
		gameevent.Emit(propertyeventtypes.EventTypePlayerNewGoldChangedLog, ppdm.p, newGoldLogEventData)
	}

	// 消耗元宝（含绑元）
	if reason != commonlog.GoldLogReasonMarryWeddingGrade &&
		reason != commonlog.GoldLogReasonMarryProposal &&
		reason != commonlog.GoldLogReasonTrade {
		gameevent.Emit(propertyeventtypes.EventTypePlayerGoldCostIncludeBind, ppdm.p, gold)
	}

	data := propertyeventtypes.CreatePlayerGoldChangedLogEventData(beforeGold, beforeBindGold, gold, reason, reasonText)
	gameevent.Emit(propertyeventtypes.EventTypePlayerGoldChangedLog, ppdm.p, data)
	return true
}

//是否有足够的钱
func (ppdm *PlayerPropertyDataManager) HasEnoughSilver(silver int64) bool {
	currentSilver := ppdm.GetSilver()
	if currentSilver < silver {
		return false
	}
	return true
}

//消耗银两
func (ppdm *PlayerPropertyDataManager) CostSilver(silverNum int64, reason commonlog.SilverLogReason, reasonText string) bool {
	if silverNum < 0 {
		panic(fmt.Errorf("property:消耗应该不能小于0"))
	}
	if silverNum == 0 {
		return true
	}
	currentSilver := ppdm.GetSilver()
	beforeSilver := currentSilver
	if currentSilver < int64(silverNum) {
		return false
	}
	currentSilver -= silverNum
	ppdm.setSilver(currentSilver)
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()

	data := propertyeventtypes.CreatePlayerSilverChangedLogEventData(beforeSilver, silverNum, reason, reasonText)
	gameevent.Emit(propertyeventtypes.EventTypePlayerSilverChangedLog, ppdm.p, data)
	return true
}

//获得银两
func (ppdm *PlayerPropertyDataManager) AddSilver(silverNum int64, reason commonlogtypes.LogReason, reasonText string) {
	if silverNum <= 0 {
		panic(fmt.Errorf("property:消耗应该不能小于0"))
	}

	tianshuRate := float64(ppdm.p.GetTianShuRate(tianshutypes.TianShuTypeSilver))
	extraSilver := int64(math.Ceil(float64(silverNum) * tianshuRate / float64(common.MAX_RATE)))

	currentSilver := ppdm.GetSilver()
	beforeSilver := currentSilver
	currentSilver += silverNum
	currentSilver += extraSilver

	ppdm.setSilver(currentSilver)
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()

	data := propertyeventtypes.CreatePlayerSilverChangedLogEventData(beforeSilver, silverNum, reason, reasonText)
	gameevent.Emit(propertyeventtypes.EventTypePlayerSilverChangedLog, ppdm.p, data)
}

//添加转生
func (ppdm *PlayerPropertyDataManager) SetZhuanSheng(zhuanShengNum int32, reason commonlog.ZhuanShengLogReason, reasonText string) {
	ppdm.setZhuanSheng(zhuanShengNum)
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()
	ppdm.p.SyncZhuanSheng(zhuanShengNum)
	gameevent.Emit(propertyeventtypes.EventTypePlayerZhuanShengChanged, ppdm.Player(), zhuanShengNum)
}

//添加魅力值
func (ppdm *PlayerPropertyDataManager) AddCharm(charm int32, reason commonlog.CharmLogReason, reasonText string) {
	if charm <= 0 {
		panic(fmt.Errorf("property:魅力值应该不能小于0"))
	}
	previousCharm := ppdm.GetCharm()
	currentCharm := ppdm.GetCharm()
	currentCharm += charm

	//超过最大值
	if currentCharm > propertytypes.MaxCharmLimit {
		currentCharm = propertytypes.MaxCharmLimit
	}

	eventData := currentCharm - previousCharm
	gameevent.Emit(propertyeventtypes.EventTypePlayerCharmAdd, ppdm.p, eventData)
	ppdm.setCharm(currentCharm)
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()
}

func (ppdm *PlayerPropertyDataManager) SetPower(power int64, mask uint64) {
	if ppdm.playerPropertyObject.Power == power {
		return
	}
	beforeForce := ppdm.playerPropertyObject.Power
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.Power = power
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()
	reasonText := playerpropertytypes.StringForMask(mask)
	eventData := propertyeventtypes.CreatePlayerForceChangedEventData(beforeForce, power, mask, reasonText)
	gameevent.Emit(propertyeventtypes.EventTypePlayerForceChanged, ppdm.p, eventData)

	// 更新战力记录
	ppdm.updatePowerRecord(beforeForce, power)
}

//战力记录
func (ppdm *PlayerPropertyDataManager) updatePowerRecord(beforePower, power int64) {
	now := global.GetGame().GetTimeService().Now()
	isSame, _ := timeutils.IsSameDay(now, ppdm.playerPowerRecordObj.UpdateTime)
	if !isSame {
		ppdm.playerPowerRecordObj.TodayInitPower = beforePower
	}

	hisMaxPower := ppdm.playerPowerRecordObj.HisMaxPower
	if power > hisMaxPower {
		ppdm.playerPowerRecordObj.HisMaxPower = power
	}

	ppdm.playerPowerRecordObj.UpdateTime = now
	ppdm.playerPowerRecordObj.SetModified()
}

func (ppdm *PlayerPropertyDataManager) GetTodayInitPower() int64 {
	return ppdm.playerPowerRecordObj.TodayInitPower
}

//TODO:zrc 优化退出存储
func (ppdm *PlayerPropertyDataManager) SetHp(hp int64) {
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.CurrentHP = hp
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()
}

//TODO:zrc 优化退出存储
func (ppdm *PlayerPropertyDataManager) SetTp(tp int64) {
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.CurrentTP = tp
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()
}

//设置罪恶值
func (ppdm *PlayerPropertyDataManager) SetEvil(evil int32) {
	ppdm.setEvil(evil)
	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()
}

//设置罪恶值在线时间
func (ppdm *PlayerPropertyDataManager) SetEvilOnlineTime(evilOnlineTime int64) {
	ppdm.setEvilOnlineTime(evilOnlineTime)
}

//添加经验点
func (ppdm *PlayerPropertyDataManager) AddExpPoint(expPoint int64, reason commonlog.LevelLogReason, reasonText string) {
	level := ppdm.GetLevel()
	tempLevelTemplate := template.GetTemplateService().Get(int(level), (*gametemplate.CharacterLevelTemplate)(nil))
	if tempLevelTemplate == nil {
		return
	}
	levelTemplate := tempLevelTemplate.(*gametemplate.CharacterLevelTemplate)
	exp := int64(math.Ceil(levelTemplate.GetExpRatio() * float64(expPoint)))
	ppdm.AddExp(exp, reason, reasonText)
}

//获取经验点
func (ppdm *PlayerPropertyDataManager) GetExpPoint(expPoint int64) int64 {
	level := ppdm.GetLevel()
	tempLevelTemplate := template.GetTemplateService().Get(int(level), (*gametemplate.CharacterLevelTemplate)(nil))
	if tempLevelTemplate == nil {
		return 0
	}
	levelTemplate := tempLevelTemplate.(*gametemplate.CharacterLevelTemplate)
	exp := int64(math.Ceil(levelTemplate.GetExpRatio() * float64(expPoint)))
	return exp
}

//添加经验
func (ppdm *PlayerPropertyDataManager) AddExp(exp int64, reason commonlogtypes.LogReason, reasonText string) {
	if exp <= 0 {
		panic(fmt.Errorf("property:添加经验至少为1,%d", exp))
	}

	tianshuRate := float64(ppdm.p.GetTianShuRate(tianshutypes.TianShuTypeExp))
	extraExp := int64(math.Ceil(float64(exp) * tianshuRate / float64(common.MAX_RATE)))

	currentExp := ppdm.GetExp()
	currentExp += exp
	currentExp += extraExp

	gameevent.Emit(propertyeventtypes.EventTypePlayerExpAdd, ppdm.p, exp)
	oldLevel := ppdm.GetLevel()
	changeExp := currentExp - ppdm.GetExp()
	levelUp := ppdm.checkExpAndLevel(currentExp)
	if levelUp {
		//检查等级
		ppdm.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeLevel.Mask())
	}

	data := propertyeventtypes.CreatePlayerExpChangedLogEventData(currentExp, ppdm.GetExp(), changeExp, oldLevel, ppdm.GetLevel(), reason, reasonText)
	gameevent.Emit(propertyeventtypes.EventTypePlayerExpChangedLog, ppdm.p, data)
}

//添加经验和经验点
func (ppdm *PlayerPropertyDataManager) AddExpAndExpPoint(exp int64, expPoint int64, reason commonlogtypes.LogReason, reasonText string) {
	if exp <= 0 {
		panic(fmt.Errorf("property:添加经验至少为1,%d", exp))
	}

	exp += ppdm.GetExpPoint(expPoint)
	ppdm.AddExp(exp, reason, reasonText)
}

//扣经验
func (ppdm *PlayerPropertyDataManager) CostExp(exp int64, reason commonlog.LevelLogReason, reasonText string) (flag bool) {
	if exp <= 0 {
		panic(fmt.Errorf("property:扣除经验至少为1,%d", exp))
	}
	oldExp := ppdm.GetExp()
	newExp := ppdm.GetExp()
	newExp -= exp
	oldLevel := ppdm.GetLevel()
	newLevel := oldLevel
	for newExp < 0 {
		newLevel -= 1
		maxExp, flag := getMaxExp(newLevel)
		if !flag {
			return false
		}
		newExp += int64(maxExp)
	}
	ppdm.setExp(newExp)
	ppdm.setLevel(newLevel)

	if newLevel < oldLevel {
		//检查等级
		ppdm.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeLevel.Mask())
		ppdm.p.SyncLevel(newLevel)
		gameevent.Emit(propertyeventtypes.EventTypePlayerLevelChanged, ppdm.p, newLevel)

	}

	ppdm.playerPropertyObject.SetModified()
	//后台日志
	data := propertyeventtypes.CreatePlayerExpChangedLogEventData(oldExp, newExp, exp, oldLevel, newLevel, reason, reasonText)
	gameevent.Emit(propertyeventtypes.EventTypePlayerExpChangedLog, ppdm.p, data)
	return true
}

func (ppdm *PlayerPropertyDataManager) checkExpAndLevel(currentExp int64) (levelUp bool) {
	oldLevel := ppdm.GetLevel()
	newExp := currentExp
	newLevel := oldLevel
	for {
		exp, flag := getMaxExp(newLevel)
		if !flag {
			break
		}
		if newExp < int64(exp) {
			break
		}

		newLevel += 1
		newExp -= int64(exp)
	}
	if newLevel > oldLevel {
		now := global.GetGame().GetTimeService().Now()
		ppdm.playerPropertyObject.UpdateTime = now
		levelUp = true
		ppdm.setLevel(newLevel)
		ppdm.setExp(newExp)
		ppdm.playerPropertyObject.SetModified()
		ppdm.p.SyncLevel(newLevel)
		gameevent.Emit(propertyeventtypes.EventTypePlayerLevelChanged, ppdm.p, newLevel)

	} else {
		now := global.GetGame().GetTimeService().Now()
		ppdm.playerPropertyObject.UpdateTime = now
		ppdm.setExp(newExp)
		ppdm.playerPropertyObject.SetModified()
	}
	return
}

func getMaxExp(level int32) (exp int64, flag bool) {
	tempTemplateObject := template.GetTemplateService().Get(int(level), (*gametemplate.CharacterLevelTemplate)(nil))
	if tempTemplateObject == nil {
		return 0, false
	}
	tempLevelTemplate := tempTemplateObject.(*gametemplate.CharacterLevelTemplate)
	if tempLevelTemplate.GetNextLevelTemplate() == nil {
		return 0, false
	}
	exp = tempLevelTemplate.GetNextLevelTemplate().Experience
	flag = true
	return
}

//扣除元神经验
func (ppdm *PlayerPropertyDataManager) CostGoldYuanExp(exp int64, reason commonlog.GoldYuanLevelLogReason, reasonText string) {
	if exp <= 0 {
		panic(fmt.Errorf("property:添加经验至少为1,%d", exp))
	}

	oldLevel := ppdm.GetGoldYuanLevel()
	newLevel := oldLevel
	newExp := ppdm.GetGoldYuanExp()
	newExp -= exp

	for newExp < 0 {
		newLevel -= 1
		maxExp := ppdm.getGoldYuanMaxExp(newLevel)
		if maxExp == 0 {
			newExp = 0
			break
		}
		newExp += maxExp
	}

	if newLevel < oldLevel {
		ppdm.setGoldYuanLevel(newLevel)
		ppdm.setGoldYuanExp(newExp)

		//检查等级
		ppdm.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeGoldYuan.Mask())
	} else {
		ppdm.setGoldYuanExp(newExp)
	}

	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()
}

//添加元神经验
func (ppdm *PlayerPropertyDataManager) AddGoldYuanExp(exp int64, reason commonlog.GoldYuanLevelLogReason, reasonText string) {
	if exp <= 0 {
		panic(fmt.Errorf("property:添加经验至少为1,%d", exp))
	}
	currentExp := ppdm.GetGoldYuanExp()
	currentExp += exp

	isUplevel := ppdm.checkGoldYuanLevelAndExp(currentExp)
	if isUplevel {
		//检查等级
		ppdm.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeGoldYuan.Mask())
	}
}

func (ppdm *PlayerPropertyDataManager) checkGoldYuanLevelAndExp(curExp int64) (isUplevel bool) {
	oldLevel := ppdm.GetGoldYuanLevel()
	newLevel := oldLevel
	newExp := curExp

	for {
		maxExp := ppdm.getGoldYuanMaxExp(newLevel)
		if maxExp == 0 {
			newExp = 0
			break
		}
		if newExp < maxExp {
			break
		}

		newExp -= maxExp
		newLevel += 1
	}

	if newLevel > oldLevel {
		isUplevel = true
		ppdm.setGoldYuanLevel(newLevel)
		ppdm.setGoldYuanExp(newExp)
	} else {
		ppdm.setGoldYuanExp(newExp)
	}

	now := global.GetGame().GetTimeService().Now()
	ppdm.playerPropertyObject.UpdateTime = now
	ppdm.playerPropertyObject.SetModified()

	return isUplevel
}

func (ppdm *PlayerPropertyDataManager) getGoldYuanMaxExp(level int32) (maxExp int64) {
	curGoldYuanTemp := goldequiptemplate.GetGoldEquipTemplateService().GetGoldYuanTemplate(level)
	if curGoldYuanTemp == nil {
		return
	}

	if curGoldYuanTemp.GetNextGoldYuanTemplate() == nil {
		return
	}

	maxExp = int64(curGoldYuanTemp.GetNextGoldYuanTemplate().Exp)
	return
}

//快捷操作
func (ppdm *PlayerPropertyDataManager) GetLevel() int32 {
	return ppdm.playerPropertyObject.Level
}

func (ppdm *PlayerPropertyDataManager) GetExp() int64 {
	return ppdm.playerPropertyObject.Exp
}

func (ppdm *PlayerPropertyDataManager) GetTP() int64 {
	return ppdm.playerPropertyObject.CurrentTP
}

func (ppdm *PlayerPropertyDataManager) GetSilver() int64 {
	return ppdm.playerPropertyObject.Silver
}

func (ppdm *PlayerPropertyDataManager) GetGold() int64 {
	return ppdm.playerPropertyObject.Gold
}

func (ppdm *PlayerPropertyDataManager) GetBindGlod() int64 {
	return ppdm.playerPropertyObject.BindGold
}

func (ppdm *PlayerPropertyDataManager) GetEvil() int32 {
	return ppdm.playerPropertyObject.Evil
}

func (ppdm *PlayerPropertyDataManager) GetZhuanSheng() int32 {
	return ppdm.playerPropertyObject.ZhuanSheng
}

func (ppdm *PlayerPropertyDataManager) GetCharm() int32 {
	return ppdm.playerPropertyObject.Charm
}

func (ppdm *PlayerPropertyDataManager) GetGoldYuanExp() int64 {
	return ppdm.playerPropertyObject.GoldYuanExp
}

func (ppdm *PlayerPropertyDataManager) GetGoldYuanLevel() int32 {
	return ppdm.playerPropertyObject.GoldYuanLevel
}

//获取战斗属性
func (ppdm *PlayerPropertyDataManager) GetBattleProperty(battlePropertyType propertytypes.BattlePropertyType) int64 {
	return ppdm.battlePropertyGroup.Get(battlePropertyType)
}

func (ppdm *PlayerPropertyDataManager) GetMaxHP() int64 {
	return ppdm.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeMaxHP)
}

func (ppdm *PlayerPropertyDataManager) Player() player.Player {
	return ppdm.p
}

// 获取魅力日志并删除数据
func (m *PlayerPropertyDataManager) GetOfflineAddCharm() int32 {
	now := global.GetGame().GetTimeService().Now()
	charm := int32(0)
	for _, obj := range m.charmAddLogList {
		charm += obj.Charm
		obj.UpdateTime = now
		obj.DeleteTime = now
		obj.SetModified()
	}
	return charm
}

//加载
func (ppdm *PlayerPropertyDataManager) Load() (err error) {
	//TODO 数据加载封装
	pse, err := dao.GetPropertyDao().GetPropertyEntity(ppdm.p.GetId())
	if err != nil {
		return
	}
	if pse == nil {
		ppdm.initPlayerPropertyObject()
	} else {
		ppdm.playerPropertyObject = NewPlayerPropertyObject(ppdm.p)
		err = ppdm.playerPropertyObject.FromEntity(pse)
		if err != nil {
			return
		}
	}

	//加载玩家档次首充信息
	cycleRecordEntity, err := dao.GetPropertyDao().GetPlayerCycleCostRecordEntity(ppdm.p.GetId())
	if err != nil {
		return
	}

	if cycleRecordEntity != nil {
		obj := NewPlayerCycleCostRecordObject(ppdm.p)
		obj.FromEntity(cycleRecordEntity)
		ppdm.playerCycleCostRecord = obj
	} else {
		ppdm.initCycleCostObject()
	}

	entity, err := dao.GetPropertyDao().GetPowerRecordEntity(ppdm.p.GetId())
	if err != nil {
		return
	}
	if entity != nil {
		ppdm.playerPowerRecordObj = NewPlayerPowerRecordObject(ppdm.p)
		ppdm.playerPowerRecordObj.FromEntity(entity)
	} else {
		ppdm.initPlayerPowerRecordObject()
	}

	// 玩家魅力增加日志数据
	charmAddLogEntityList, err := dao.GetPropertyDao().GetCharmAddLogList(ppdm.p.GetId())
	if err != nil {
		return
	}
	for _, charmAddLog := range charmAddLogEntityList {
		charmAddLogObject := NewPlayerCharmAddLogObject(ppdm.p)
		charmAddLogObject.FromEntity(charmAddLog)
		ppdm.charmAddLogList = append(ppdm.charmAddLogList, charmAddLogObject)
	}

	return nil
}

func (ppdm *PlayerPropertyDataManager) AfterLoad() (err error) {
	//刷新
	ppdm.refresh()

	//加载玩家魅力
	charm := ppdm.GetOfflineAddCharm()
	if charm != 0 {
		charmAddReason := commonlog.CharmLogReasonGiftReward
		ppdm.AddCharm(charm, charmAddReason, charmAddReason.String())
	}

	return nil
}

//第一次初始化
func (ppdm *PlayerPropertyDataManager) initPlayerPropertyObject() {
	ppo := NewPlayerPropertyObject(ppdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	ppo.Id = id
	//生成id
	ppo.PlayerId = ppdm.p.GetId()
	ppo.Level = initLevel
	ppo.Exp = 0
	ppo.Silver = 0
	ppo.Gold = 0
	ppo.BindGold = 0
	ppo.Evil = 0
	ppo.CurrentHP = 0
	ppo.CurrentTP = 0
	ppo.ZhuanSheng = 0
	ppo.Charm = 0
	ppo.CreateTime = now
	ppdm.playerPropertyObject = ppo
	ppo.SetModified()
}

//第一次初始化
func (ppdm *PlayerPropertyDataManager) initPlayerPowerRecordObject() {
	ppo := NewPlayerPowerRecordObject(ppdm.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	ppo.Id = id
	ppo.HisMaxPower = 0
	ppo.TodayInitPower = 0
	ppo.CreateTime = now
	ppo.SetModified()
	ppdm.playerPowerRecordObj = ppo
}

func (ppdm *PlayerPropertyDataManager) Heartbeat() {

}

func (ppdm *PlayerPropertyDataManager) GetModuleForce(typ playerpropertytypes.PropertyEffectorType) int64 {
	return ppdm.battlePropertyGroup.GetPropertySegment(typ).GetForce()
}

func (ppdm *PlayerPropertyDataManager) GetModule(typ playerpropertytypes.PropertyEffectorType) *propertycommon.SystemPropertySegment {
	return ppdm.battlePropertyGroup.GetPropertySegment(typ)
}

//更新战斗属性
func (ppdm *PlayerPropertyDataManager) UpdateBattleProperty(mask uint64) {
	//各个系统
	for _, effType := range propertyplayerproperty.GetAllPlayerPropertyEffectors() {
		if effType.Mask()&mask != 0 {
			pef := propertyplayerproperty.GetPlayerPropertyEffector(effType)
			if pef == nil {
				panic("never reach here")
			}
			//获取相对应属性
			p := ppdm.battlePropertyGroup.GetPropertySegment(effType)
			//重置属性
			p.Clear()
			//属性作用
			pef(ppdm.p, p)
		}
	}
	ppdm.battlePropertyGroup.UpdateProperty()
	ppdm.updateForce(mask)
	for _, effType := range propertyplayerproperty.GetAllPlayerPropertyEffectors() {
		p := ppdm.battlePropertyGroup.GetPropertySegment(effType)
		p.UpdateModuleProperty()

		gameevent.Emit(propertyeventtypes.EventTypePlayerPropertyEffectorChanged, ppdm.p, effType)
	}

	gameevent.Emit(propertyeventtypes.EventTypePlayerSystemPropertyChanged, ppdm.p, nil)

}

func (m *PlayerPropertyDataManager) updateForce(mask uint64) {
	force := float64(0)
	forceTemplate := constant.GetConstantService().GetForceTemplate()
	for k, v := range forceTemplate.GetAllForceProperty() {
		val := m.battlePropertyGroup.Get(k)
		if k == propertytypes.BattlePropertyTypeMoveSpeed {
			initSpeed := constant.GetConstantService().GetConstant(constantypes.ConstantTypeInitMoveSpeed)
			val -= int64(initSpeed)
		}
		if k == propertytypes.BattlePropertyTypeHit {
			hit := constant.GetConstantService().GetConstant(constantypes.ConstantTypeInitalHit)
			val -= int64(hit)
		}
		tempForce := float64(val) / common.MAX_RATE * float64(v)
		force += tempForce
	}
	power := int64(math.Floor(force)) + m.battlePropertyGroup.Get(propertytypes.BattlePropertyTypeForce)
	m.SetPower(power, mask)
}

func (ppdm *PlayerPropertyDataManager) GetForce() int64 {
	return ppdm.playerPropertyObject.Power
}

func (ppdm *PlayerPropertyDataManager) setLevel(level int32) {
	ppdm.playerPropertyObject.Level = level
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeLevel, int64(level))

}
func (ppdm *PlayerPropertyDataManager) setExp(exp int64) {
	ppdm.playerPropertyObject.Exp = exp
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeExp, exp)
}
func (ppdm *PlayerPropertyDataManager) setSilver(silver int64) {
	ppdm.playerPropertyObject.Silver = silver
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeSilver, silver)
}

func (ppdm *PlayerPropertyDataManager) setGold(gold int64) {
	ppdm.playerPropertyObject.Gold = gold
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeGold, int64(gold))
}

func (ppdm *PlayerPropertyDataManager) setBindGold(bindGold int64) {
	ppdm.playerPropertyObject.BindGold = bindGold
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeBindGold, int64(bindGold))
}

func (ppdm *PlayerPropertyDataManager) setCharm(charm int32) {
	ppdm.playerPropertyObject.Charm = charm
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeCharm, int64(charm))
}

func (ppdm *PlayerPropertyDataManager) setGoldYuanLevel(level int32) {
	ppdm.playerPropertyObject.GoldYuanLevel = level
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeGoldYuanLevel, int64(level))
}

func (ppdm *PlayerPropertyDataManager) setGoldYuanExp(exp int64) {
	ppdm.playerPropertyObject.GoldYuanExp = exp
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeGoldYuanExp, int64(exp))
}

func (ppdm *PlayerPropertyDataManager) setEvil(evil int32) {
	ppdm.playerPropertyObject.Evil = evil
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeEvil, int64(evil))
}

func (ppdm *PlayerPropertyDataManager) setEvilOnlineTime(onlineTime int64) {
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeEvilOnlineTime, int64(onlineTime))
}

func (ppdm *PlayerPropertyDataManager) setCurrentHP(currentHP int64) {
	ppdm.playerPropertyObject.CurrentHP = currentHP
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeHP, currentHP)
}

func (ppdm *PlayerPropertyDataManager) setCurrentTP(currentTP int64) {
	ppdm.playerPropertyObject.CurrentTP = currentTP
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeTP, currentTP)
}

func (ppdm *PlayerPropertyDataManager) setZhuanSheng(zhuanSheng int32) {
	ppdm.playerPropertyObject.ZhuanSheng = zhuanSheng
	ppdm.baseProperty.Set(propertytypes.BasePropertyTypeZhuanSheng, int64(zhuanSheng))
}

const (
	initLevel      = 1
	initZhuanSheng = 0
)

//刷新数据
func (ppdm *PlayerPropertyDataManager) refresh() {
	//第一次初始化
	if ppdm.playerPropertyObject.UpdateTime == 0 {
		now := global.GetGame().GetTimeService().Now()
		role := ppdm.Player().GetRole()
		sex := ppdm.Player().GetSex()
		playerCreateTemplate := constant.GetConstantService().GetPlayerCreateTemplate(role, sex)
		ppdm.setLevel(initLevel)
		ppdm.setExp(0)
		ppdm.setSilver(int64(playerCreateTemplate.Money))
		ppdm.setGold(0)
		ppdm.setBindGold(0)
		ppdm.setEvil(0)
		ppdm.setCurrentHP(0)
		ppdm.setCurrentTP(0)
		ppdm.setZhuanSheng(initZhuanSheng)
		ppdm.setCharm(0)
		ppdm.setGoldYuanExp(0)
		ppdm.setGoldYuanLevel(0)

		ppdm.playerPropertyObject.UpdateTime = now
		ppdm.playerPropertyObject.SetModified()
	} else {
		ppdm.setLevel(ppdm.GetLevel())
		ppdm.setExp(ppdm.GetExp())
		ppdm.setSilver(ppdm.GetSilver())
		ppdm.setGold(ppdm.GetGold())
		ppdm.setBindGold(ppdm.GetBindGlod())
		ppdm.setEvil(ppdm.GetEvil())
		ppdm.setCurrentHP(ppdm.GetHP())
		ppdm.setCurrentTP(ppdm.GetTP())
		ppdm.setZhuanSheng(ppdm.GetZhuanSheng())
		ppdm.setCharm(ppdm.GetCharm())
		ppdm.setGoldYuanExp(ppdm.GetGoldYuanExp())
		ppdm.setGoldYuanLevel(ppdm.GetGoldYuanLevel())

		//s	ppdm.checkExpAndLevel(ppdm.GetExp())
	}
	return
}

func (ppdm *PlayerPropertyDataManager) ToBaseProperties() map[int32]int64 {
	properties := make(map[int32]int64)
	for typ := propertytypes.BasePropertyTypeHP; typ <= propertytypes.BasePropertyTypeGoldYuanExp; typ++ {
		val := ppdm.baseProperty.Get(typ)
		properties[int32(typ)] = val
	}
	return properties
}

func (ppdm *PlayerPropertyDataManager) ToBattleProperties() map[int32]int64 {
	properties := make(map[int32]int64)
	for typ := propertytypes.MinBattlePropertyType; typ <= propertytypes.MaxBattlePropertyType; typ++ {
		val := ppdm.battlePropertyGroup.Get(typ)
		properties[int32(typ)] = val
	}
	return properties
}

//更新每日消费信息
func (ppdm *PlayerPropertyDataManager) UpdateCycleCostInfo(costNum int64) {
	ppdm.refreshCycleCost()

	now := global.GetGame().GetTimeService().Now()
	ppdm.playerCycleCostRecord.costNum += costNum
	ppdm.playerCycleCostRecord.updateTime = now
	ppdm.playerCycleCostRecord.SetModified()
}

//获取今日消费数
func (ppdm *PlayerPropertyDataManager) GetTodayCostNum() int64 {
	ppdm.refreshCycleCost()
	return ppdm.playerCycleCostRecord.costNum
}

//刷新每日消费信息
func (ppdm *PlayerPropertyDataManager) refreshCycleCost() {
	now := global.GetGame().GetTimeService().Now()
	diff, _ := timeutils.DiffDay(now, ppdm.playerCycleCostRecord.updateTime)
	if diff != 0 {
		if diff == 1 {
			ppdm.playerCycleCostRecord.preDayCostNum = ppdm.playerCycleCostRecord.costNum
		}
		ppdm.playerCycleCostRecord.updateTime = now
		ppdm.playerCycleCostRecord.costNum = 0
		ppdm.playerCycleCostRecord.SetModified()
	}
}

//初始化每日消费对象
func (ppdm *PlayerPropertyDataManager) initCycleCostObject() {
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()

	obj := NewPlayerCycleCostRecordObject(ppdm.p)
	obj.id = id
	obj.costNum = 0
	obj.preDayCostNum = 0
	obj.createTime = now
	obj.updateTime = now
	obj.SetModified()
	ppdm.playerCycleCostRecord = obj
}

func CreatePlayerPropertyDataManager(p player.Player) player.PlayerDataManager {

	ppdm := &PlayerPropertyDataManager{}
	ppdm.p = p
	ppdm.battlePropertyGroup = propertycommon.NewBattlePropertyGroup()
	ppdm.baseProperty = propertycommon.NewBasePropertySegment()
	return ppdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerPropertyDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerPropertyDataManager))
}
