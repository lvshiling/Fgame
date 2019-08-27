package logic

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/template"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constantypes "fgame/fgame/game/constant/types"
	consttypes "fgame/fgame/game/constant/types"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/property/pbutil"
	playerproperty "fgame/fgame/game/property/player"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	xianzuncardlogic "fgame/fgame/game/xianzuncard/logic"
	"fmt"
	"math"
)

//推送战斗属性变化
func SnapChangedBattleProperty(pl scene.Player) {
	battleChanged := pl.GetBattlePropertyChangedTypesAndReset()
	if len(battleChanged) == 0 {
		return
	}
	force := pl.GetForce()
	scPlayerPropertyData := pbutil.BuildSCPlayerPropertyData(nil, battleChanged, force)
	pl.SendMsg(scPlayerPropertyData)
	return
}

//推送属性变化
func SnapChangedProperty(pl player.Player) {
	manager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	baseChanged := manager.GetChangedBasePropertiesAndReset()
	if len(baseChanged) == 0 {
		return
	}
	force := pl.GetForce()
	// basedChanged, _ := manager.GetChangedTypesAndReset()

	scPlayerPropertyData := pbutil.BuildSCPlayerPropertyData(baseChanged, nil, force)
	pl.SendMsg(scPlayerPropertyData)

	return
}

func CombineRes(rd *propertytypes.RewData, resMap map[itemtypes.ItemAutoUseResSubType]int32) *propertytypes.RewData {
	if len(resMap) <= 0 {
		panic(fmt.Errorf("property:resMap length should be more 0"))
	}
	rewExp := rd.GetRewExp()
	rewExpPoint := rd.GetRewExpPoint()
	rewSilver := rd.GetRewSilver()
	rewGold := rd.GetRewGold()
	rewBindGold := rd.GetRewBindGold()
	for typ, resNum := range resMap {
		if resNum <= 0 {
			panic(fmt.Errorf("property:resNum  should be more 0"))
		}
		switch typ {
		case itemtypes.ItemAutoUseResSubTypeSilver:
			{
				rewSilver += resNum
				break
			}
		case itemtypes.ItemAutoUseResSubTypeGold:
			{
				rewGold += resNum
				break
			}
		case itemtypes.ItemAutoUseResSubTypeBindGold:
			{
				rewBindGold += resNum
				break
			}
		}
	}
	newRd := propertytypes.CreateRewData(rewExp, rewExpPoint, rewSilver, rewGold, rewBindGold)
	return newRd
}

func GetRewDataMap(rd *propertytypes.RewData, level int32) (rewMap map[int32]int32) {

	rewMap = make(map[int32]int32)
	rewSilver := rd.RewSilver
	rewGold := rd.RewGold
	rewBindGold := rd.RewBindGold
	rewExpPoint := rd.RewExpPoint
	rewExp := rd.RewExp

	if rewSilver > 0 {
		rewMap[int32(consttypes.SilverItem)] = rewSilver
	}

	if rewGold > 0 {
		rewMap[int32(consttypes.GoldItem)] = rewGold
	}

	if rewBindGold > 0 {
		rewMap[int32(consttypes.BindGoldItem)] = rewBindGold
	}

	if rewExpPoint > 0 {
		tempLevelTemplate := template.GetTemplateService().Get(int(level), (*gametemplate.CharacterLevelTemplate)(nil))
		if tempLevelTemplate == nil {
			return
		}
		levelTemplate := tempLevelTemplate.(*gametemplate.CharacterLevelTemplate)
		exp := int64(math.Ceil(levelTemplate.GetExpRatio() * float64(rewExpPoint)))
		rewExp += int32(exp)
	}

	if rewExp > 0 {
		rewMap[int32(consttypes.ExpItem)] = rewExp
	}
	return
}

// 经验点转经验
func ExpPointConvertExp(expPoint int32, level int32) int64 {
	if expPoint > 0 {
		tempLevelTemplate := template.GetTemplateService().Get(int(level), (*gametemplate.CharacterLevelTemplate)(nil))
		if tempLevelTemplate == nil {
			return 0
		}
		levelTemplate := tempLevelTemplate.(*gametemplate.CharacterLevelTemplate)
		return int64(math.Ceil(levelTemplate.GetExpRatio() * float64(expPoint)))
	}
	return 0
}

func GetItemMapWithExpPoint(expPoint int32, level int32, curItemMap map[int32]int32) (itemMap map[int32]int32) {
	if expPoint <= 0 {
		return curItemMap
	}
	itemMap = make(map[int32]int32)
	for itemId, num := range curItemMap {
		itemMap[itemId] += num
	}
	tempLevelTemplate := template.GetTemplateService().Get(int(level), (*gametemplate.CharacterLevelTemplate)(nil))
	if tempLevelTemplate == nil {
		return curItemMap
	}
	levelTemplate := tempLevelTemplate.(*gametemplate.CharacterLevelTemplate)
	exp := int32(math.Ceil(levelTemplate.GetExpRatio() * float64(expPoint)))
	if exp >= 0 {
		itemMap[consttypes.ExpItem] += exp
	}
	return
}

//重新计算战力
func CulculateForce(propertyMap map[propertytypes.BattlePropertyType]int64) int64 {
	force := float64(0)
	forceTemplate := constant.GetConstantService().GetForceTemplate()
	for k, v := range forceTemplate.GetAllForceProperty() {
		val := propertyMap[k]
		tempForce := float64(val) / common.MAX_RATE * float64(v)
		force += tempForce
	}
	power := int64(math.Floor(force)) + propertyMap[propertytypes.BattlePropertyTypeForce]
	return power
}

//重新计算战力
func CulculateAllForce(propertyMap map[propertytypes.BattlePropertyType]int64) int64 {
	force := float64(0)
	forceTemplate := constant.GetConstantService().GetForceTemplate()
	for k, v := range forceTemplate.GetAllForceProperty() {
		val := propertyMap[k]
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
	power := int64(math.Floor(force)) + propertyMap[propertytypes.BattlePropertyTypeForce]
	return power
}

// 击杀怪加经验
func AddExpKillMonster(pl player.Player, monsterId int32, expBase int64, expPoint int64) (err error) {

	//判断经验加成
	luckyRate := pl.GetLuckyRate(itemtypes.ItemTypeResouceCard, itemtypes.ItemResourceCardSubTypeExp)
	expAddPercent := luckyRate
	if expAddPercent > 0 {
		expPoint = int64(math.Ceil(float64(expPoint) * (1 + float64(expAddPercent)/float64(common.MAX_RATE))))
		expBase = int64(math.Ceil(float64(expBase) * (1 + float64(expAddPercent)/float64(common.MAX_RATE))))
	}

	//仙尊卡加成
	teQuanPercent := xianzuncardlogic.MonsterExpExtralPercent(pl, monsterId)
	if teQuanPercent > 0 {
		expPoint = int64(math.Ceil(float64(expPoint) * (1 + float64(teQuanPercent)/float64(common.MAX_RATE))))
		expBase = int64(math.Ceil(float64(expBase) * (1 + float64(teQuanPercent)/float64(common.MAX_RATE))))
	}

	//判断
	//判断经验衰减
	// levelEscape := int32(math.Abs(float64(int32(bt.Level) - pl.GetLevel())))
	// yaZhiTemplate := scenetemplate.GetSceneTemplateService().GetYaZhiTemplate(levelEscape)
	// if yaZhiTemplate != nil {
	// 	expPoint = int64(math.Ceil(float64(expPoint) * (1 - float64(yaZhiTemplate.ExpPercent)/float64(common.MAX_RATE))))
	// 	expBase = int64(math.Ceil(float64(expBase) * (1 - float64(yaZhiTemplate.ExpPercent)/float64(common.MAX_RATE))))
	// }

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	reason := commonlog.LevelLogReasonMonsterKilled
	reasonText := fmt.Sprintf(reason.String(), monsterId)
	expBase = int64(math.Ceil(float64(expBase) * pl.GetWallowState().Rate()))
	if expBase != 0 {
		propertyManager.AddExp(expBase, reason, reasonText)
	}

	reasonExtra := commonlog.LevelLogReasonMonsterKilledExtra
	reasonExtraText := fmt.Sprintf(reasonExtra.String(), monsterId)
	expPointValue := int64(math.Ceil(float64(expPoint) * pl.GetWallowState().Rate()))
	if expPointValue != 0 {
		propertyManager.AddExpPoint(expPointValue, reasonExtra, reasonExtraText)
	}

	//推送
	SnapChangedProperty(pl)
	return
}

//资源和物品整合
func CombineRewDataAndItemData(rd *propertytypes.RewData, level int32, itemMap map[int32]int32) (rewMap map[int32]int32) {
	rewMap = GetRewDataMap(rd, level)
	for itemId, itemNum := range itemMap {
		rewMap[itemId] = rewMap[itemId] + itemNum
	}
	return
}
