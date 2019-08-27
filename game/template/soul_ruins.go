package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	itemtypes "fgame/fgame/game/item/types"
	propertytypes "fgame/fgame/game/property/types"
	scenetypes "fgame/fgame/game/scene/types"
	"fgame/fgame/game/soulruins/types"
	"fmt"
)

//帝魂配置
type SoulRuinsTemplate struct {
	*SoulRuinsTemplateVO
	typ               types.SoulRuinsType
	eventTyp          types.SoulRuinsEventType
	specialEventMap   map[types.SoulRuinsEventType]map[int32]int32 //特殊事件
	sweepEventDropMap map[types.SoulRuinsEventType]int32           //扫荡触发事件掉落包
	specialRateMap    map[types.SoulRuinsEventType]int32           //特殊事件概率
	groupIdMap        map[types.SoulRuinsEventType]int32           //事件刷怪组
	starTimeMap       map[types.SoulRuinsStarNumType]int32         //星数减少时限
	sweepNeedItemMap  map[int32]int32                              //扫荡需要物品
	sweepDropList     []int32                                      //扫荡固定掉落包
	rewData           *propertytypes.RewData                       //奖励属性
	mapTemplate       *MapTemplate                                 //副本地图
	robberSilver      int32
}

func (srt *SoulRuinsTemplate) TemplateId() int {
	return srt.Id
}

func (srt *SoulRuinsTemplate) GetType() types.SoulRuinsType {
	return srt.typ
}

func (srt *SoulRuinsTemplate) GetEventType() types.SoulRuinsEventType {
	return srt.eventTyp
}

func (srt *SoulRuinsTemplate) GetSpecialEventMap() map[types.SoulRuinsEventType]map[int32]int32 {
	return srt.specialEventMap
}

func (srt *SoulRuinsTemplate) GetSweepEventDropMap() map[types.SoulRuinsEventType]int32 {
	return srt.sweepEventDropMap
}

func (srt *SoulRuinsTemplate) GetSpecialRateMap() map[types.SoulRuinsEventType]int32 {
	return srt.specialRateMap
}

func (srt *SoulRuinsTemplate) GetGroupIdByEventType(eventType types.SoulRuinsEventType) int32 {
	return srt.groupIdMap[eventType]
}

func (srt *SoulRuinsTemplate) GetStarTimeMap() map[types.SoulRuinsStarNumType]int32 {
	return srt.starTimeMap
}

func (srt *SoulRuinsTemplate) GetSweepNeedItemMap() map[int32]int32 {
	return srt.sweepNeedItemMap
}

func (srt *SoulRuinsTemplate) GetSweepDropList() []int32 {
	return srt.sweepDropList
}

func (srt *SoulRuinsTemplate) GetRewData() *propertytypes.RewData {
	return srt.rewData
}

func (srt *SoulRuinsTemplate) GetRobberSilver() int32 {
	return srt.RobberSilver
}

func (srt *SoulRuinsTemplate) PatchAfterCheck() {

}

func (srt *SoulRuinsTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(srt.FileName(), srt.TemplateId(), err)
			return
		}
	}()

	//验证 type
	srt.typ = types.SoulRuinsType(srt.Type)
	if !srt.typ.Valid() {
		err = fmt.Errorf("[%d] invalid", srt.Type)
		err = template.NewTemplateFieldError("Type", err)
		return
	}

	//first_event
	if srt.FirstEvent != 0 {
		srt.eventTyp = types.SoulRuinsEventType(srt.FirstEvent)
		if !srt.eventTyp.Valid() {
			err = fmt.Errorf("[%d] invalid", srt.FirstEvent)
			err = template.NewTemplateFieldError("FirstEvent", err)
			return
		}
	}

	srt.starTimeMap = make(map[types.SoulRuinsStarNumType]int32)
	//time_1
	err = validator.MinValidate(float64(srt.Time1), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.Time1)
		return template.NewTemplateFieldError("Time1", err)
	}
	//time_2
	if srt.Time2 < srt.Time1 {
		err = fmt.Errorf("[%d] invalid", srt.Time2)
		return template.NewTemplateFieldError("Time2", err)
	}
	//time_3
	if srt.Time3 < srt.Time2 {
		err = fmt.Errorf("[%d] invalid", srt.Time3)
		return template.NewTemplateFieldError("Time3", err)
	}
	srt.starTimeMap[types.SoulRuinsStarNumTypeThree] = srt.Time1
	srt.starTimeMap[types.SoulRuinsStarNumTypeTwo] = srt.Time2
	srt.starTimeMap[types.SoulRuinsStarNumTypeOne] = srt.Time3

	srt.sweepNeedItemMap = make(map[int32]int32)
	//sweep_item_id
	if srt.SweepItemId != 0 {
		to := template.GetTemplateService().Get(int(srt.SweepItemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", srt.SweepItemId)
			err = template.NewTemplateFieldError("SweepItemId", err)
			return
		}
		itemType := to.(*ItemTemplate).GetItemType()
		if itemType != itemtypes.ItemTypeDefault {
			err = fmt.Errorf("[%d] invalid", srt.SweepItemId)
			err = template.NewTemplateFieldError("SweepItemId", err)
			return
		}

		//验证 soul_item_count
		err = validator.MinValidate(float64(srt.SweepItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", srt.SweepItemCount)
			return template.NewTemplateFieldError("SweepItemCount", err)
		}

		srt.sweepNeedItemMap[srt.SweepItemId] = srt.SweepItemCount

	}

	srt.specialEventMap = make(map[types.SoulRuinsEventType]map[int32]int32)
	srt.specialRateMap = make(map[types.SoulRuinsEventType]int32)
	srt.groupIdMap = make(map[types.SoulRuinsEventType]int32)
	srt.sweepEventDropMap = make(map[types.SoulRuinsEventType]int32)

	//赋值specialRateMap
	srt.specialRateMap[types.SoulRuinsEventTypeBoss] = srt.SpecialBossRate
	srt.specialRateMap[types.SoulRuinsEventTypeSoul] = srt.SoulRate
	srt.specialRateMap[types.SoulRuinsEventTypeRobber] = srt.RobberRate
	//赋值groupIdMap
	srt.groupIdMap[types.SoulRuinsEventTypeBoss] = srt.SpecialBossGroup
	srt.groupIdMap[types.SoulRuinsEventTypeSoul] = srt.SoulGroup
	srt.groupIdMap[types.SoulRuinsEventTypeRobber] = srt.RobberGroup
	//赋值sweepEventDropMap
	srt.sweepEventDropMap[types.SoulRuinsEventTypeBoss] = srt.SpecialBossDrop
	srt.sweepEventDropMap[types.SoulRuinsEventTypeSoul] = srt.SoulDrop
	srt.sweepEventDropMap[types.SoulRuinsEventTypeRobber] = srt.RobberDrop
	//赋值specialEventMap
	bossEventMap, ok := srt.specialEventMap[types.SoulRuinsEventTypeBoss]
	if !ok {
		bossEventMap = make(map[int32]int32)
		srt.specialEventMap[types.SoulRuinsEventTypeBoss] = bossEventMap
	}
	bossEventMap[srt.SpecialBossId] = srt.SpecialBossBuff

	soulEventMap, ok := srt.specialEventMap[types.SoulRuinsEventTypeSoul]
	if !ok {
		soulEventMap = make(map[int32]int32)
		srt.specialEventMap[types.SoulRuinsEventTypeSoul] = soulEventMap
	}
	soulEventMap[srt.SoulItemId] = srt.SoulItemCount

	robberEventMap, ok := srt.specialEventMap[types.SoulRuinsEventTypeRobber]
	if !ok {
		robberEventMap = make(map[int32]int32)
		srt.specialEventMap[types.SoulRuinsEventTypeRobber] = robberEventMap
	}
	robberEventMap[srt.RobberId] = srt.RobberCount

	//sweep_drop
	srt.sweepDropList = make([]int32, 0, 8)
	if srt.SweepDrop1 != "" {
		dropIdArr, err := utils.SplitAsIntArray(srt.SweepDrop1)
		if err != nil {
			return err
		}
		for _, dropId := range dropIdArr {
			srt.sweepDropList = append(srt.sweepDropList, dropId)
		}
	}

	//srt.sweepDropList = append(srt.sweepDropList, srt.SweepDrop2)
	//srt.sweepDropList = append(srt.sweepDropList, srt.SweepDrop3)
	//srt.sweepDropList = append(srt.sweepDropList, srt.SweepDrop4)

	//验证 rew_yinliang
	err = validator.MinValidate(float64(srt.RewYinliang), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.RewYinliang)
		return template.NewTemplateFieldError("RewYinliang", err)
	}

	//验证 rew_exp
	err = validator.MinValidate(float64(srt.RewExp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.RewExp)
		return template.NewTemplateFieldError("RewExp", err)
	}

	//验证 rew_uplev
	err = validator.MinValidate(float64(srt.RewUplev), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.RewUplev)
		return template.NewTemplateFieldError("RewUplev", err)
	}

	//验证 rew_gold
	err = validator.MinValidate(float64(srt.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.RewGold)
		return template.NewTemplateFieldError("RewGold", err)
	}

	if srt.RewYinliang > 0 || srt.RewGold > 0 || srt.RewUplev > 0 || srt.RewExp > 0 {
		srt.rewData = propertytypes.CreateRewData(srt.RewExp, srt.RewUplev, srt.RewYinliang, srt.RewGold, 0)
	}

	//验证 map_id
	to := template.GetTemplateService().Get(int(srt.MapId), (*MapTemplate)(nil))
	if to == nil {
		err = fmt.Errorf("[%d] invalid", srt.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}
	srt.mapTemplate = to.(*MapTemplate)

	//验证 robber_silver
	err = validator.MinValidate(float64(srt.RobberSilver), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.RobberSilver)
		return template.NewTemplateFieldError("RobberSilver", err)
	}

	return nil
}

func (srt *SoulRuinsTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(srt.FileName(), srt.TemplateId(), err)
			return
		}
	}()

	//map_id 类型
	mapType := srt.mapTemplate.GetMapType()
	if mapType != scenetypes.SceneTypeFuBenSoulRuins {
		err = fmt.Errorf("[%d] invalid", srt.MapId)
		err = template.NewTemplateFieldError("MapId", err)
		return
	}

	//验证事件
	for typ, eventMap := range srt.specialEventMap {
		//boss
		if typ == types.SoulRuinsEventTypeBoss {
			err = srt.checkBossEvent(eventMap)
		}
		//帝魂降临
		if typ == types.SoulRuinsEventTypeSoul {
			err = srt.checkSoulEvent(eventMap)
		}
		//马贼
		if typ == types.SoulRuinsEventTypeRobber {
			err = srt.checkRobberEvent(eventMap)
		}
		if err != nil {
			return
		}

	}

	//验证概率
	for _, rate := range srt.specialRateMap {
		err = validator.RangeValidate(float64(rate), float64(0), true, float64(common.MAX_RATE), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", rate)
			err = template.NewTemplateFieldError("Rate", err)
			return
		}
	}

	//验证 事件刷怪组
	for _, groupId := range srt.groupIdMap {
		group := srt.mapTemplate.GetSceneBiologyMapByGroup(groupId)
		if group == nil {
			err = fmt.Errorf("[%d] invalid", groupId)
			err = template.NewTemplateFieldError("Group", err)
			return
		}
	}

	//验证 chapter
	err = validator.MinValidate(float64(srt.Chapter), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.Chapter)
		return template.NewTemplateFieldError("Chapter", err)
	}

	//验证 level
	err = validator.MinValidate(float64(srt.Level), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.Level)
		return template.NewTemplateFieldError("Level", err)
	}

	//验证 next_id
	if srt.NextId != 0 {
		diff := srt.NextId - int32(srt.Id)
		to := template.GetTemplateService().Get(int(srt.NextId), (*SoulRuinsTemplate)(nil))
		if to == nil || diff != 1 {
			err = fmt.Errorf("[%d] invalid", srt.NextId)
			err = template.NewTemplateFieldError("NextId", err)
			return
		}

		nextTemplate := to.(*SoulRuinsTemplate)
		//验证 chapter
		if srt.Chapter != nextTemplate.Chapter {
			err = fmt.Errorf("[%d] invalid", nextTemplate.Chapter)
			err = template.NewTemplateFieldError("Chapter", err)
			return
		}

		//验证 typ
		if srt.typ != nextTemplate.typ {
			err = fmt.Errorf("[%d] invalid", nextTemplate.Type)
			err = template.NewTemplateFieldError("Type", err)
			return
		}

		//验证level
		diffLevel := nextTemplate.Level - srt.Level
		if diffLevel != 1 {
			err = fmt.Errorf("[%d] invalid", nextTemplate.Level)
			err = template.NewTemplateFieldError("Level", err)
			return
		}
	}

	//验证 rew_time
	err = validator.MinValidate(float64(srt.RewTime), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", srt.RewTime)
		return template.NewTemplateFieldError("RewTime", err)
	}

	//验证 front_id
	if srt.FrontId != 0 {
		to := template.GetTemplateService().Get(int(srt.FrontId), (*SoulRuinsTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", srt.FrontId)
			err = template.NewTemplateFieldError("FrontId", err)
			return
		}
	}

	return nil
}

func (srt *SoulRuinsTemplate) checkBossEvent(bossEventMap map[int32]int32) (err error) {
	for bossId, buffId := range bossEventMap {
		//验证 special_boss_id
		to := template.GetTemplateService().Get(int(bossId), (*BiologyTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", bossId)
			err = template.NewTemplateFieldError("SpecialBossId", err)
			return
		}
		biologyType := to.(*BiologyTemplate).GetBiologyScriptType()
		if biologyType != scenetypes.BiologyScriptTypeSoulBoss {
			err = fmt.Errorf("[%d] invalid", srt.SpecialBossId)
			err = template.NewTemplateFieldError("SpecialBossId", err)
			return
		}

		//TODO 校验special_boss_buff
		err = validator.MinValidate(float64(buffId), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", buffId)
			return template.NewTemplateFieldError("special_boss_buff", err)
		}
	}
	return
}

func (srt *SoulRuinsTemplate) checkSoulEvent(soulEventMap map[int32]int32) (err error) {
	for itemId, num := range soulEventMap {
		//验证 soul_item_id
		to := template.GetTemplateService().Get(int(itemId), (*ItemTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", itemId)
			err = template.NewTemplateFieldError("SoulItemId", err)
			return
		}

		//验证 soul_item_count
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			return template.NewTemplateFieldError("SoulItemCount", err)
		}
	}
	return

}

func (srt *SoulRuinsTemplate) checkRobberEvent(robberEventMap map[int32]int32) (err error) {
	for robberId, num := range robberEventMap {
		//验证 robber_id
		to := template.GetTemplateService().Get(int(robberId), (*BiologyTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", robberId)
			err = template.NewTemplateFieldError("RobberId", err)
			return
		}
		biologyType := to.(*BiologyTemplate).GetBiologyScriptType()
		if biologyType != scenetypes.BiologyScriptTypeRobber {
			err = fmt.Errorf("[%d] invalid", robberId)
			err = template.NewTemplateFieldError("RobberId", err)
			return
		}

		//验证 robber_count
		err = validator.MinValidate(float64(num), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%d] invalid", num)
			return template.NewTemplateFieldError("RobberCount", err)
		}

	}
	return

}

func (srt *SoulRuinsTemplate) FileName() string {
	return "tb_soul_ruins.json"
}

func init() {
	template.Register((*SoulRuinsTemplate)(nil))
}
