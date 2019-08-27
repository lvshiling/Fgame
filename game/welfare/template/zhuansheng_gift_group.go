package template

import (
	playertypes "fgame/fgame/game/player/types"
	gametemplate "fgame/fgame/game/template"
)

// 运营活动-转生礼包奖励分组模板
type ZhuanSengGiftGroupTemplate struct {
	zhuanShengMap        map[playertypes.RoleType]map[playertypes.SexType][]*gametemplate.ZhuanShengGiftTemplate
	sexZhuanShengMap     map[playertypes.SexType][]*gametemplate.ZhuanShengGiftTemplate
	roleZhuanShengMap    map[playertypes.RoleType][]*gametemplate.ZhuanShengGiftTemplate
	commonZhuanShengList []*gametemplate.ZhuanShengGiftTemplate
}

func CreateZhuanShengGroupTemplate() *ZhuanSengGiftGroupTemplate {
	gt := &ZhuanSengGiftGroupTemplate{}
	gt.zhuanShengMap = make(map[playertypes.RoleType]map[playertypes.SexType][]*gametemplate.ZhuanShengGiftTemplate)
	gt.sexZhuanShengMap = make(map[playertypes.SexType][]*gametemplate.ZhuanShengGiftTemplate)
	gt.roleZhuanShengMap = make(map[playertypes.RoleType][]*gametemplate.ZhuanShengGiftTemplate)
	return gt
}

func (gt *ZhuanSengGiftGroupTemplate) AddTemplate(temp *gametemplate.ZhuanShengGiftTemplate) {
	if temp.Profession != 0 && temp.Gender == 0 {
		gt.roleZhuanShengMap[temp.GetRole()] = append(gt.roleZhuanShengMap[temp.GetRole()], temp)
		return
	}

	if temp.Profession == 0 && temp.Gender != 0 {
		gt.sexZhuanShengMap[temp.GetSex()] = append(gt.sexZhuanShengMap[temp.GetSex()], temp)
		return
	}

	if temp.Profession == 0 && temp.Gender == 0 {
		gt.commonZhuanShengList = append(gt.commonZhuanShengList, temp)
		return
	}

	sexMap, ok := gt.zhuanShengMap[temp.GetRole()]
	if !ok {
		sexMap = make(map[playertypes.SexType][]*gametemplate.ZhuanShengGiftTemplate)
		gt.zhuanShengMap[temp.GetRole()] = sexMap
	}

	sexMap[temp.GetSex()] = append(sexMap[temp.GetSex()], temp)
}

func (group *ZhuanSengGiftGroupTemplate) GetDiscountZhuanShengTemplateByType(role playertypes.RoleType, sex playertypes.SexType, typ int32) *gametemplate.ZhuanShengGiftTemplate {
	tempList := group.GetDiscountZhuanShengTemplate(role, sex)
	for _, temp := range tempList {
		if temp.Type == typ {
			return temp
		}
	}

	return nil
}

func (group *ZhuanSengGiftGroupTemplate) GetDiscountZhuanShengTemplate(role playertypes.RoleType, sex playertypes.SexType) []*gametemplate.ZhuanShengGiftTemplate {
	// 全部通用
	totalTempList := group.commonZhuanShengList

	// 性别通用
	sexCommonList := group.roleZhuanShengMap[role]
	if len(sexCommonList) > 0 {
		totalTempList = append(totalTempList, sexCommonList...)
	}

	// 角色通用
	roleCommonList := group.sexZhuanShengMap[sex]
	if len(roleCommonList) > 0 {
		totalTempList = append(totalTempList, roleCommonList...)
	}

	// 全部限制
	sexMap, ok := group.zhuanShengMap[role]
	if !ok {
		return totalTempList
	}
	tempList, ok := sexMap[sex]
	if !ok {
		return totalTempList
	}

	totalTempList = append(totalTempList, tempList...)
	return totalTempList
}
