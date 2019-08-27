package template

import (
	"fgame/fgame/core/template"
	"fgame/fgame/core/template/validator"
	"fgame/fgame/core/utils"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/common/common"
	inventorytypes "fgame/fgame/game/inventory/types"
	materialtypes "fgame/fgame/game/material/types"
	playertypes "fgame/fgame/game/player/types"
	propertytypes "fgame/fgame/game/property/types"
	questtypes "fgame/fgame/game/quest/types"
	scenetypes "fgame/fgame/game/scene/types"
	soultypes "fgame/fgame/game/soul/types"

	xianfutypes "fgame/fgame/game/xianfu/types"
	"fmt"
)

type QuestTemplate struct {
	*QuestTemplateVO
	//任务类型
	questType questtypes.QuestType
	//任务子类型
	questSubType questtypes.QuestSubType
	//完成需要收集的物品
	reqItemMap map[playertypes.RoleType]map[int32]int32
	//前置任务
	prevQuestIds []int32
	//后置任务
	nextQuestIds []int32
	//接取消耗的物品
	consumeItemMap map[playertypes.RoleType]map[int32]int32
	//角色奖励
	roleRewardItemMap map[playertypes.RoleType]map[int32]int32
	//要求数据
	questDemandMap map[playertypes.RoleType]map[int32]int32
	//任务品质
	questLevel questtypes.QuestLevelType
	//奖励属性
	rewData *propertytypes.RewData
	//校验使用的map
	checkMap map[int32]struct{}
	//跳跃点生物id
	portBiologyTemplate *BiologyTemplate
	//跳跃点
	portalTemplate       *PortalTemplate
	guideReplicaTemplate *GuideReplicaTemplate
}

//任务类型
func (qt *QuestTemplate) GetQuestType() questtypes.QuestType {
	return qt.questType
}

func (qt *QuestTemplate) GetQuestSubType() questtypes.QuestSubType {
	return qt.questSubType
}

func (qt *QuestTemplate) IsAutoFinishByUsedFree() bool {
	return qt.IsAutoFinishFree == 0
}

func (qt *QuestTemplate) GetQuestDemandMap(roleType playertypes.RoleType) map[int32]int32 {
	return qt.questDemandMap[roleType]
}

func (qt *QuestTemplate) GetReqItemMap(roleType playertypes.RoleType) map[int32]int32 {
	return qt.reqItemMap[roleType]
}

func (qt *QuestTemplate) GetQuestLevel() questtypes.QuestLevelType {
	return qt.questLevel
}

func (qt *QuestTemplate) GetRewData() *propertytypes.RewData {
	return qt.rewData
}

func (qt *QuestTemplate) GetFeiXieNum() int32 {
	return qt.IsFeiXie
}

//获取接取需要消耗的物品
func (qt *QuestTemplate) GetConsumeItemMap(roleType playertypes.RoleType) map[int32]int32 {
	return qt.consumeItemMap[roleType]
}

//获取奖励的物品
func (qt *QuestTemplate) GetRewardItemMap(role playertypes.RoleType) map[int32]int32 {
	rewards, ok := qt.roleRewardItemMap[role]
	if !ok {
		return nil
	}
	return rewards
}

//获取前置id列表
func (qt *QuestTemplate) GetPrevQuestIds() []int32 {
	return qt.prevQuestIds
}

//获取后置id列表
func (qt *QuestTemplate) GetNextQuestIds() []int32 {
	return qt.nextQuestIds
}

//是否自动接受
func (qt *QuestTemplate) AutoAccept() bool {
	return qt.AcceptCreature == 0
}

//是否自动完成
func (qt *QuestTemplate) AutoCommit() bool {
	return qt.IsAutoCommi == 0
}

//获取跳跃点
func (qt *QuestTemplate) GetPortalTemplate() *PortalTemplate {
	return qt.portalTemplate
}

func (qt *QuestTemplate) TemplateId() int {
	return qt.Id
}

func (qt *QuestTemplate) Patch() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(qt.FileName(), qt.TemplateId(), err)
			return
		}
	}()

	qt.questType = questtypes.QuestType(qt.QuestType)
	if !qt.questType.Valid() {
		err = fmt.Errorf("[%d] invalid", qt.QuestType)
		return template.NewTemplateFieldError("questType", err)
	}

	if qt.questType == questtypes.QuestTypeTuMo {
		qt.questLevel = questtypes.QuestLevelType(qt.QuestLevel)
		if !qt.questLevel.Valid() {
			err = fmt.Errorf("[%d] invalid", qt.questLevel)
			return template.NewTemplateFieldError("questLevel", err)
		}
	}

	qt.questSubType = questtypes.QuestSubType(qt.QuestSubType)
	if !qt.questSubType.Valid() {
		err = fmt.Errorf("[%d] invalid", qt.QuestSubType)
		return template.NewTemplateFieldError("questSubType", err)
	}
	qt.questDemandMap = make(map[playertypes.RoleType]map[int32]int32)
	//开天
	demandArr, err := coreutils.SplitAsIntArray(qt.QuestDemand)
	if err != nil {
		return template.NewTemplateFieldError("questDemand", err)
	}
	demandCountArr, err := coreutils.SplitAsIntArray(qt.QuestDemandCount)
	if err != nil {
		return template.NewTemplateFieldError("questDemandCount", err)
	}
	if len(demandArr) != len(demandCountArr) {
		return template.NewTemplateFieldError("questDemand or questDemandCount", err)
	}
	for i := 0; i < len(demandArr); i++ {
		demandId := demandArr[i]
		demandCount := demandCountArr[i]

		kaiTianDemandMap, exist := qt.questDemandMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianDemandMap = make(map[int32]int32)
			qt.questDemandMap[playertypes.RoleTypeKaiTian] = kaiTianDemandMap
		}
		kaiTianDemandMap[demandId] = demandCount
	}

	//奕剑
	demandArr2, err := coreutils.SplitAsIntArray(qt.QuestDemand2)
	if err != nil {
		return template.NewTemplateFieldError("questDemand2", err)
	}
	demandCountArr2, err := coreutils.SplitAsIntArray(qt.QuestDemandCount2)
	if err != nil {
		return template.NewTemplateFieldError("questDemandCount2", err)
	}
	if len(demandArr2) != len(demandCountArr2) {
		return template.NewTemplateFieldError("questDemand2 or questDemandCount2", err)
	}
	for i := 0; i < len(demandArr2); i++ {
		demandId := demandArr2[i]
		demandCount := demandCountArr2[i]

		yiJianDemandMap, exist := qt.questDemandMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianDemandMap = make(map[int32]int32)
			qt.questDemandMap[playertypes.RoleTypeYiJian] = yiJianDemandMap
		}
		yiJianDemandMap[demandId] = demandCount
	}

	//破月
	demandArr3, err := coreutils.SplitAsIntArray(qt.QuestDemand3)
	if err != nil {
		return template.NewTemplateFieldError("questDemand3", err)
	}
	demandCountArr3, err := coreutils.SplitAsIntArray(qt.QuestDemandCount3)
	if err != nil {
		return template.NewTemplateFieldError("questDemandCount2", err)
	}
	if len(demandArr3) != len(demandCountArr3) {
		return template.NewTemplateFieldError("questDemand3 or questDemandCount3", err)
	}
	for i := 0; i < len(demandArr3); i++ {
		demandId := demandArr3[i]
		demandCount := demandCountArr3[i]

		poYueDemandMap, exist := qt.questDemandMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueDemandMap = make(map[int32]int32)
			qt.questDemandMap[playertypes.RoleTypePoYue] = poYueDemandMap
		}
		poYueDemandMap[demandId] = demandCount
	}

	qt.reqItemMap = make(map[playertypes.RoleType]map[int32]int32)
	//开天
	reqItemArr, err := coreutils.SplitAsIntArray(qt.ReqItemId)
	if err != nil {
		return template.NewTemplateFieldError("reqItemId", err)
	}
	reqItemCountArr, err := coreutils.SplitAsIntArray(qt.ReqItemCount)
	if err != nil {
		return template.NewTemplateFieldError("reqItemCount", err)
	}
	if len(reqItemArr) != len(reqItemCountArr) {
		return template.NewTemplateFieldError("reqItemId or reqItemCount", err)
	}
	for i := 0; i < len(reqItemArr); i++ {
		reqItem := reqItemArr[i]
		reqItemCount := reqItemCountArr[i]
		tempItemTemplate := template.GetTemplateService().Get(int(reqItem), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", qt.ReqItemId)
			return template.NewTemplateFieldError("reqItemId", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(reqItemCount), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("reqItemCount", err)
		}

		kaiTianReqItemMap, exist := qt.reqItemMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianReqItemMap = make(map[int32]int32)
			qt.reqItemMap[playertypes.RoleTypeKaiTian] = kaiTianReqItemMap
		}
		kaiTianReqItemMap[reqItem] = reqItemCount
	}

	//奕剑
	reqItemArr2, err := coreutils.SplitAsIntArray(qt.ReqItemId2)
	if err != nil {
		return template.NewTemplateFieldError("reqItemId2", err)
	}
	reqItemCountArr2, err := coreutils.SplitAsIntArray(qt.ReqItemCount2)
	if err != nil {
		return template.NewTemplateFieldError("reqItemCount2", err)
	}
	if len(reqItemArr2) != len(reqItemCountArr2) {
		return template.NewTemplateFieldError("reqItemId2 or reqItemCount2", err)
	}
	for i := 0; i < len(reqItemArr2); i++ {
		reqItem2 := reqItemArr[i]
		reqItemCount2 := reqItemCountArr2[i]
		tempItemTemplate := template.GetTemplateService().Get(int(reqItem2), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", qt.ReqItemId2)
			return template.NewTemplateFieldError("reqItemId2", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(reqItemCount2), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("reqItemCount2", err)
		}

		yiJianReqItemMap, exist := qt.reqItemMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianReqItemMap = make(map[int32]int32)
			qt.reqItemMap[playertypes.RoleTypeYiJian] = yiJianReqItemMap
		}
		yiJianReqItemMap[reqItem2] = reqItemCount2
	}

	//破月
	reqItemArr3, err := coreutils.SplitAsIntArray(qt.ReqItemId3)
	if err != nil {
		return template.NewTemplateFieldError("reqItemId3", err)
	}
	reqItemCountArr3, err := coreutils.SplitAsIntArray(qt.ReqItemCount3)
	if err != nil {
		return template.NewTemplateFieldError("reqItemCount3", err)
	}
	if len(reqItemArr3) != len(reqItemCountArr3) {
		return template.NewTemplateFieldError("reqItemId3 or reqItemCount3", err)
	}
	for i := 0; i < len(reqItemArr3); i++ {
		reqItem3 := reqItemArr[i]
		reqItemCount3 := reqItemCountArr[i]
		tempItemTemplate := template.GetTemplateService().Get(int(reqItem3), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%d] invalid", qt.ReqItemId3)
			return template.NewTemplateFieldError("reqItemId3", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(reqItemCount3), float64(1), true)
		if err != nil {
			return template.NewTemplateFieldError("reqItemCount3", err)
		}

		yiJianReqItemMap, exist := qt.reqItemMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianReqItemMap = make(map[int32]int32)
			qt.reqItemMap[playertypes.RoleTypeYiJian] = yiJianReqItemMap
		}
		yiJianReqItemMap[reqItem3] = reqItemCount3
	}

	qt.prevQuestIds, err = coreutils.SplitAsIntArray(qt.PrevQuest)
	if err != nil {
		return template.NewTemplateFieldError("prevQuest", err)
	}
	//前置任务
	for _, prevQuestId := range qt.prevQuestIds {
		tempQuestTemplate := template.GetTemplateService().Get(int(prevQuestId), (*QuestTemplate)(nil))
		if tempQuestTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.PrevQuest)
			return template.NewTemplateFieldError("prevQuest", err)
		}
	}

	qt.nextQuestIds, err = coreutils.SplitAsIntArray(qt.NextQuest)
	if err != nil {
		return template.NewTemplateFieldError("NextQuest", err)
	}
	//后置任务
	for _, nextQuestId := range qt.nextQuestIds {
		tempQuestTemplate := template.GetTemplateService().Get(int(nextQuestId), (*QuestTemplate)(nil))
		if tempQuestTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.NextQuest)
			return template.NewTemplateFieldError("prevQuest", err)
		}
	}

	//后置id
	followQuestIds, err := coreutils.SplitAsIntArray(qt.FollowId)
	if err != nil {
		return template.NewTemplateFieldError("followId", err)
	}
	for _, followQuestId := range followQuestIds {
		tempQuestTemplate := template.GetTemplateService().Get(int(followQuestId), (*QuestTemplate)(nil))
		if tempQuestTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.FollowId)
			return template.NewTemplateFieldError("followId", err)
		}
		if coreutils.ContainInt32(qt.nextQuestIds, followQuestId) {
			err = fmt.Errorf("[%s] invalid", qt.FollowId)
			return template.NewTemplateFieldError("followId", err)
		}
		qt.nextQuestIds = append(qt.nextQuestIds, followQuestId)
	}

	//消耗材料
	qt.consumeItemMap = make(map[playertypes.RoleType]map[int32]int32)
	//开天 接取消耗的物品
	consumeItemArr, err := coreutils.SplitAsIntArray(qt.ConsumeItem)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.ConsumeItem)
		return template.NewTemplateFieldError("consumeItem", err)
	}
	consumeItemCountArr, err := coreutils.SplitAsIntArray(qt.ConsumeItemCount)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.ConsumeItemCount)
		return template.NewTemplateFieldError("consumeItemCount", err)
	}
	if len(consumeItemArr) != len(consumeItemCountArr) {
		err = fmt.Errorf("[%s] [%s] invalid", qt.ConsumeItem, qt.ConsumeItemCount)
		return template.NewTemplateFieldError("consumeItem or consumeItemCount", err)
	}
	for i := 0; i < len(consumeItemArr); i++ {
		consumeItem := consumeItemArr[i]

		consumeItemCount := consumeItemCountArr[i]
		tempItemTemplate := template.GetTemplateService().Get(int(consumeItem), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.ConsumeItem)
			return template.NewTemplateFieldError("consumeItem", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(consumeItemCount), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", qt.ConsumeItemCount)
			return template.NewTemplateFieldError("consumeItemCount", err)
		}

		kaiTianConsumeItemMap, exist := qt.consumeItemMap[playertypes.RoleTypeKaiTian]
		if !exist {
			kaiTianConsumeItemMap = make(map[int32]int32)
			qt.consumeItemMap[playertypes.RoleTypeKaiTian] = kaiTianConsumeItemMap
		}
		kaiTianConsumeItemMap[consumeItem] = consumeItemCount
	}

	//奕剑 接取消耗的物品
	consumeItemArr2, err := coreutils.SplitAsIntArray(qt.ConsumeItem2)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.ConsumeItem2)
		return template.NewTemplateFieldError("consumeItem2", err)
	}
	consumeItemCountArr2, err := coreutils.SplitAsIntArray(qt.ConsumeItemCount2)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.ConsumeItemCount2)
		return template.NewTemplateFieldError("consumeItemCount2", err)
	}
	if len(consumeItemArr2) != len(consumeItemCountArr2) {
		err = fmt.Errorf("[%s] [%s] invalid", qt.ConsumeItem2, qt.ConsumeItemCount2)
		return template.NewTemplateFieldError("consumeItem2 or consumeItemCount2", err)
	}
	for i := 0; i < len(consumeItemArr2); i++ {
		consumeItem2 := consumeItemArr2[i]

		consumeItemCount2 := consumeItemCountArr2[i]
		tempItemTemplate := template.GetTemplateService().Get(int(consumeItem2), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.ConsumeItem2)
			return template.NewTemplateFieldError("consumeItem2", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(consumeItemCount2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", qt.ConsumeItemCount2)
			return template.NewTemplateFieldError("consumeItemCount2", err)
		}

		yiJianConsumeItemMap, exist := qt.consumeItemMap[playertypes.RoleTypeYiJian]
		if !exist {
			yiJianConsumeItemMap = make(map[int32]int32)
			qt.consumeItemMap[playertypes.RoleTypeYiJian] = yiJianConsumeItemMap
		}
		yiJianConsumeItemMap[consumeItem2] = consumeItemCount2
	}

	//破月 接取消耗的物品
	consumeItemArr3, err := coreutils.SplitAsIntArray(qt.ConsumeItem3)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.ConsumeItem3)
		return template.NewTemplateFieldError("consumeItem3", err)
	}
	consumeItemCountArr3, err := coreutils.SplitAsIntArray(qt.ConsumeItemCount3)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.ConsumeItemCount3)
		return template.NewTemplateFieldError("consumeItemCount3", err)
	}
	if len(consumeItemArr3) != len(consumeItemCountArr3) {
		err = fmt.Errorf("[%s] [%s] invalid", qt.ConsumeItem3, qt.ConsumeItemCount3)
		return template.NewTemplateFieldError("consumeItem3 or consumeItemCount3", err)
	}
	for i := 0; i < len(consumeItemArr3); i++ {
		consumeItem3 := consumeItemArr3[i]

		consumeItemCount3 := consumeItemCountArr3[i]
		tempItemTemplate := template.GetTemplateService().Get(int(consumeItem3), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.ConsumeItem3)
			return template.NewTemplateFieldError("consumeItem3", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(consumeItemCount3), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", qt.ConsumeItemCount3)
			return template.NewTemplateFieldError("consumeItemCount3", err)
		}

		poYueConsumeItemMap, exist := qt.consumeItemMap[playertypes.RoleTypePoYue]
		if !exist {
			poYueConsumeItemMap = make(map[int32]int32)
			qt.consumeItemMap[playertypes.RoleTypePoYue] = poYueConsumeItemMap
		}
		poYueConsumeItemMap[consumeItem3] = consumeItemCount3
	}

	//角色奖励
	qt.roleRewardItemMap = make(map[playertypes.RoleType]map[int32]int32)
	//开天奖励的物品
	rewItemId1Arr, err := coreutils.SplitAsIntArray(qt.RewItemId1)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.RewItemId1)
		return template.NewTemplateFieldError("rewItemId1", err)
	}
	rewItemCount1Arr, err := coreutils.SplitAsIntArray(qt.RewItemCount1)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.RewItemCount1)
		return template.NewTemplateFieldError("rewItemCount1", err)
	}
	if len(rewItemId1Arr) != len(rewItemCount1Arr) {
		err = fmt.Errorf("[%s] [%s] invalid", qt.RewItemId1, qt.RewItemCount1)
		return template.NewTemplateFieldError("rewItemId1 or rewItemCount1", err)
	}
	rewItemId1Map := make(map[int32]int32)
	for i := 0; i < len(rewItemId1Arr); i++ {
		rewItemId1 := rewItemId1Arr[i]

		rewItemCount1 := rewItemCount1Arr[i]
		tempItemTemplate := template.GetTemplateService().Get(int(rewItemId1), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.RewItemId1)
			return template.NewTemplateFieldError("rewItemId1", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(rewItemCount1), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", qt.RewItemCount1)
			return template.NewTemplateFieldError("RewItemCount1", err)
		}
		rewItemId1Map[rewItemId1] = rewItemCount1
	}
	qt.roleRewardItemMap[playertypes.RoleTypeKaiTian] = rewItemId1Map

	rewItemId2Arr, err := coreutils.SplitAsIntArray(qt.RewItemId2)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.RewItemId2)
		return template.NewTemplateFieldError("rewItemId2", err)
	}
	rewItemCount2Arr, err := coreutils.SplitAsIntArray(qt.RewItemCount2)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.RewItemCount2)
		return template.NewTemplateFieldError("rewItemCount2", err)
	}
	if len(rewItemId2Arr) != len(rewItemCount2Arr) {
		err = fmt.Errorf("[%s] [%s] invalid", qt.RewItemId2, qt.RewItemCount2)
		return template.NewTemplateFieldError("rewItemId2 or rewItemCount2", err)
	}
	rewItemId2Map := make(map[int32]int32)
	for i := 0; i < len(rewItemId2Arr); i++ {
		rewItemId2 := rewItemId2Arr[i]

		rewItemCount2 := rewItemCount2Arr[i]
		tempItemTemplate := template.GetTemplateService().Get(int(rewItemId2), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.RewItemId2)
			return template.NewTemplateFieldError("rewItemId2", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(rewItemCount2), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", qt.RewItemCount2)
			return template.NewTemplateFieldError("RewItemCount2", err)
		}
		rewItemId2Map[rewItemId2] = rewItemCount2
	}
	qt.roleRewardItemMap[playertypes.RoleTypeYiJian] = rewItemId2Map

	rewItemId3Arr, err := coreutils.SplitAsIntArray(qt.RewItemId3)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.RewItemId3)
		return template.NewTemplateFieldError("rewItemId3", err)
	}
	rewItemCount3Arr, err := coreutils.SplitAsIntArray(qt.RewItemCount3)
	if err != nil {
		err = fmt.Errorf("[%s] invalid", qt.RewItemCount3)
		return template.NewTemplateFieldError("rewItemCount3", err)
	}
	if len(rewItemId3Arr) != len(rewItemCount3Arr) {
		err = fmt.Errorf("[%s] [%s] invalid", qt.RewItemId3, qt.RewItemCount3)
		return template.NewTemplateFieldError("rewItemId3 or rewItemCount3", err)
	}
	rewItemId3Map := make(map[int32]int32)
	for i := 0; i < len(rewItemId3Arr); i++ {
		rewItemId3 := rewItemId3Arr[i]

		rewItemCount3 := rewItemCount3Arr[i]
		tempItemTemplate := template.GetTemplateService().Get(int(rewItemId3), (*ItemTemplate)(nil))
		if tempItemTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.RewItemId3)
			return template.NewTemplateFieldError("rewItemId3", err)
		}
		//验证数量至少1
		err = validator.MinValidate(float64(rewItemCount3), float64(1), true)
		if err != nil {
			err = fmt.Errorf("[%s] invalid", qt.RewItemCount3)
			return template.NewTemplateFieldError("RewItemCount3", err)
		}
		rewItemId3Map[rewItemId3] = rewItemCount3
	}
	qt.roleRewardItemMap[playertypes.RoleTypePoYue] = rewItemId3Map

	if qt.RewGold > 0 || qt.RewExpPoint > 0 || qt.RewXp > 0 || qt.RewSilver > 0 || qt.RewBindGold > 0 {
		qt.rewData = propertytypes.CreateRewData(qt.RewXp, qt.RewExpPoint, qt.RewSilver, qt.RewGold, qt.RewBindGold)
	}

	if qt.Tiaoyuedian != 0 {
		tempBiologyTemplate := template.GetTemplateService().Get(int(qt.Tiaoyuedian), (*BiologyTemplate)(nil))
		if tempBiologyTemplate == nil {
			err = fmt.Errorf("[%d] invalid", qt.Tiaoyuedian)
			return template.NewTemplateFieldError("tiaoyuedian", err)
		}

		qt.portBiologyTemplate = tempBiologyTemplate.(*BiologyTemplate)
	}
	if qt.FubenId != 0 {
		tempGuideReplicaTemplate := template.GetTemplateService().Get(int(qt.FubenId), (*GuideReplicaTemplate)(nil))
		if tempGuideReplicaTemplate == nil {
			err = fmt.Errorf("[%s] invalid", qt.FubenId)
			return template.NewTemplateFieldError("FubenId", err)
		}
		qt.guideReplicaTemplate = tempGuideReplicaTemplate.(*GuideReplicaTemplate)
	}
	return nil
}

func (qt *QuestTemplate) beforeCheckNextData() (beforeDataList []int32) {
	for questId, _ := range qt.checkMap {
		beforeDataList = append(beforeDataList, questId)
	}
	return
}

func (qt *QuestTemplate) resetBeforeCheckMap(beforeDataList []int32) {
	for questId, _ := range qt.checkMap {
		flag := utils.ContainInt32(beforeDataList, questId)
		if !flag {
			delete(qt.checkMap, questId)
		}
	}
}

func (qt *QuestTemplate) checkNextQuest(curQuestId int32, nextQuestIds []int32) (err error) {
	for _, questId := range nextQuestIds {
		_, exist := qt.checkMap[questId]
		if exist {
			return fmt.Errorf("任务id:%d next_quest 包含任务:%d,会造成循环", curQuestId, questId)
		}
		to := template.GetTemplateService().Get(int(questId), (*QuestTemplate)(nil))
		questTemplate := to.(*QuestTemplate)
		if questTemplate == nil {
			continue
		}
		curNextQuestIds := questTemplate.GetNextQuestIds()
		if len(curNextQuestIds) == 0 {
			continue
		}
		beforeDataList := qt.beforeCheckNextData()
		qt.checkMap[questId] = struct{}{}
		err = qt.checkNextQuest(questId, curNextQuestIds)
		if err != nil {
			return
		}
		qt.resetBeforeCheckMap(beforeDataList)
	}
	return
}

func (qt *QuestTemplate) Check() (err error) {
	defer func() {
		if err != nil {
			err = template.NewTemplateError(qt.FileName(), qt.TemplateId(), err)
			return
		}
	}()

	if len(qt.nextQuestIds) != 0 {
		qt.checkMap = make(map[int32]struct{})
		qt.checkMap[int32(qt.Id)] = struct{}{}
		err = qt.checkNextQuest(int32(qt.Id), qt.nextQuestIds)
		if err != nil {
			return
		}
		qt.checkMap = nil
	}

	//验证接取等级至少1
	err = validator.MinValidate(float64(qt.MinLevel), float64(1), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.MinLevel)
		return template.NewTemplateFieldError("minLevel", err)
	}

	//验证接取最大等级
	err = validator.RangeValidate(float64(qt.MaxLevel), float64(qt.MinLevel), true, float64(common.MAX_LEVEL), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.MaxLevel)
		return template.NewTemplateFieldError("maxLevel", err)
	}

	err = validator.MinValidate(float64(qt.MinZhuanshu), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.MinZhuanshu)
		return template.NewTemplateFieldError("minZhuanshu", err)
	}

	//验证接取最大转数
	err = validator.RangeValidate(float64(qt.MaxZhuanshu), float64(qt.MinZhuanshu), true, float64(common.MAX_ZHUAN), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.MaxZhuanshu)
		return template.NewTemplateFieldError("maxZhuanshu", err)
	}
	//验证银两
	err = validator.MinValidate(float64(qt.RewSilver), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.RewSilver)
		return template.NewTemplateFieldError("rewSilver", err)
	}
	//验证绑元
	err = validator.MinValidate(float64(qt.RewBindGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.RewBindGold)
		return template.NewTemplateFieldError("RewBindGold", err)
	}
	//验证元宝
	err = validator.MinValidate(float64(qt.RewGold), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.RewGold)
		return template.NewTemplateFieldError("RewGold", err)
	}
	//验证经验点
	err = validator.MinValidate(float64(qt.RewExpPoint), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.RewExpPoint)
		return template.NewTemplateFieldError("rewExpPoint", err)
	}

	err = validator.MinValidate(float64(qt.RewXp), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.RewXp)
		return template.NewTemplateFieldError("rewXp", err)
	}

	//验证个别子任务的quest_demand
	err = qt.checkQuestDemand(qt.questSubType)
	if err != nil {
		return
	}

	err = validator.MinValidate(float64(qt.IsFeiXie), float64(0), true)
	if err != nil {
		err = fmt.Errorf("[%d] invalid", qt.IsFeiXie)
		return template.NewTemplateFieldError("IsFeiXie", err)
	}

	//验证活跃度任务
	switch qt.questType {
	case questtypes.QuestTypeLiveness:
		to := template.GetTemplateService().Get(qt.Id, (*HuoYueTemplate)(nil))
		if to == nil {
			err = fmt.Errorf("[%d] invalid", qt.Id)
			return template.NewTemplateFieldError("Id", err)
		}
		if !qt.AutoAccept() {
			err = fmt.Errorf("[%d] invalid", qt.AcceptCreature)
			return template.NewTemplateFieldError("AcceptCreature", err)
		}

		if !qt.AutoCommit() {
			err = fmt.Errorf("[%d] invalid", qt.IsAutoCommi)
			return template.NewTemplateFieldError("IsAutoCommi", err)
		}
		break
	case questtypes.QuestTypeTianJiPai:
		if qt.AutoCommit() {
			err = fmt.Errorf("[%d] invalid", qt.IsAutoCommi)
			return template.NewTemplateFieldError("IsAutoCommi", err)
		}
		break
	}

	if qt.portBiologyTemplate != nil {
		if qt.portBiologyTemplate.GetBiologyType() != scenetypes.BiologyTransmissionArray {
			err = fmt.Errorf("[%d] invalid", qt.Tiaoyuedian)
			return template.NewTemplateFieldError("tiaoyuedian", err)
		}
	}

	//验证日环任务
	if qt.questType == questtypes.QuestTypeDaily || qt.questType == questtypes.QuestTypeDailyAlliance {
		if qt.AutoAccept() {
			err = fmt.Errorf("[%d] invalid", qt.AcceptCreature)
			return template.NewTemplateFieldError("AcceptCreature", err)
		}
	}

	//验证开服目标任务
	if qt.questType == questtypes.QuestTypeKaiFuMuBiao {
		if qt.MinLevel != 1 {
			err = fmt.Errorf("[%d] invalid", qt.MinLevel)
			return template.NewTemplateFieldError("MinLevel", err)
		}

		if qt.MaxLevel != 999 {
			err = fmt.Errorf("[%d] invalid", qt.MaxLevel)
			return template.NewTemplateFieldError("MaxLevel", err)
		}

		if qt.MinZhuanshu != 0 {
			err = fmt.Errorf("[%d] invalid", qt.MinZhuanshu)
			return template.NewTemplateFieldError("MinZhuanshu", err)
		}

		if qt.MaxZhuanshu != 999 {
			err = fmt.Errorf("[%d] invalid", qt.MaxZhuanshu)
			return template.NewTemplateFieldError("MaxZhuanshu", err)
		}
	}

	return nil
}

//校验带指定的字样的
func (qt *QuestTemplate) checkQuestDemand(subType questtypes.QuestSubType) (err error) {
	switch subType {
	case questtypes.QuestSubTypeSpecialXianFu,
		questtypes.QuestSubTypeEnterSpecialXianFu,
		questtypes.QuestSubTypeUpgradeSpecialXianFu:
		{
			err = qt.checkXianFuDemand()
			if err != nil {
				return
			}
			break
		}
	case questtypes.QuestSubTypeEquipmentStrengthenLevel,
		questtypes.QuestSubTypeEquipmentUpgradeStar,
		questtypes.QuestSubTypeEquipmentUpgradeLevel:
		{
			err = qt.checkEquipmentDemand()
			if err != nil {
				return
			}
			break
		}
	case questtypes.QuestSubTypeSoulActive,
		questtypes.QuestSubTypeSoulStrengthenLevel,
		questtypes.QuestSubTypeSoulSpecialEmbed,
		questtypes.QuestSubTypeSoulUpgradeLevel:
		{
			err = qt.checkSoulDemand()
			if err != nil {
				return
			}
			break
		}
	case questtypes.QuestSubTypeSpecifiedSoulRuins:
		{
			err = qt.checkSoulRuinsDemand()
			if err != nil {
				return
			}
			break
		}
	case questtypes.QuestSubTypechallengeSpecialMaterialFuBen:
		{
			err = qt.checkMaterialDemand()
			if err != nil {
				return
			}
			break
		}
	default:
		break
	}
	return nil
}

//校验指定仙府
func (qt *QuestTemplate) checkXianFuDemand() (err error) {
	for _, questDemand := range qt.questDemandMap {
		if len(questDemand) != 1 {
			return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置只能一个", err)
		}
		for demand, _ := range questDemand {
			demandValue := xianfutypes.XianfuType(demand)
			if !demandValue.Valid() {
				return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置不是指定仙府类型", err)
			}
		}
	}
	return nil
}

//校验指定装备
func (qt *QuestTemplate) checkEquipmentDemand() (err error) {
	for _, questDemand := range qt.questDemandMap {
		if len(questDemand) != 1 {
			return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置只能一个", err)
		}
		for demand, _ := range questDemand {
			demandValue := inventorytypes.BodyPositionType(demand)
			if !demandValue.Valid() {
				return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置不是指定装备位置", err)
			}
		}
	}
	return nil
}

//校验指定的帝魂
func (qt *QuestTemplate) checkSoulDemand() (err error) {
	for _, questDemand := range qt.questDemandMap {
		if len(questDemand) != 1 {
			return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置只能一个", err)
		}
		for demand, _ := range questDemand {
			demandValue := soultypes.SoulType(demand)
			if !demandValue.Valid() {
				return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置不是指定帝魂标签", err)
			}
		}
	}
	return nil
}

//校验指定的帝陵副本
func (qt *QuestTemplate) checkSoulRuinsDemand() (err error) {
	for _, questDemand := range qt.questDemandMap {
		if len(questDemand) != 1 {
			return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置只能一个", err)
		}
	}
	return nil
}

//校验指定的材料副本
func (qt *QuestTemplate) checkMaterialDemand() (err error) {
	for _, questDemand := range qt.questDemandMap {
		if len(questDemand) != 1 {
			return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置只能一个", err)
		}
		for demand, _ := range questDemand {
			demandValue := materialtypes.MaterialType(demand)
			if !demandValue.Valid() {
				return template.NewTemplateFieldError("该任务的对应子类型的quest_demand 配置不是指定材料副本标签", err)
			}
		}
	}
	return
}

//校验指定的材料副本
func (qt *QuestTemplate) GetGuideReplicaTemplate() *GuideReplicaTemplate {
	return qt.guideReplicaTemplate
}

func (qt *QuestTemplate) PatchAfterCheck() {

	if qt.portBiologyTemplate != nil {
		qt.portalTemplate = qt.portBiologyTemplate.GetPortalTemplate()
	}
}
func (qt *QuestTemplate) FileName() string {
	return "tb_quest.json"
}

func init() {
	template.Register((*QuestTemplate)(nil))
}
