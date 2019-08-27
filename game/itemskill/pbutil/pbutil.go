package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	playeritemskill "fgame/fgame/game/itemskill/player"
	itemskilltypes "fgame/fgame/game/itemskill/types"
)

func BuildSCItemSkillAllGet(itemSkillAllMap map[itemskilltypes.ItemSkillType]*playeritemskill.PlayerItemSkillObject) *uipb.SCItemSkillAllGet {
	scItemSkillAllGet := &uipb.SCItemSkillAllGet{}
	for _, skill := range itemSkillAllMap {
		scItemSkillAllGet.ItemSkillList = append(scItemSkillAllGet.ItemSkillList, buildItemSkill(skill))
	}
	return scItemSkillAllGet
}

func BuildSCItemSkillActive(sk *playeritemskill.PlayerItemSkillObject) *uipb.SCItemSkillActive {
	scItemSkillActive := &uipb.SCItemSkillActive{}
	scItemSkillActive.ItemSkill = buildItemSkill(sk)
	return scItemSkillActive
}

func BuildSCItemSkillUpgrade(sk *playeritemskill.PlayerItemSkillObject) *uipb.SCItemSkillUpgrade {
	scItemSkillUpgrade := &uipb.SCItemSkillUpgrade{}
	scItemSkillUpgrade.ItemSkill = buildItemSkill(sk)
	return scItemSkillUpgrade
}

func buildItemSkill(sk *playeritemskill.PlayerItemSkillObject) *uipb.ItemSkillInfo {
	sysSkillInfo := &uipb.ItemSkillInfo{}
	typ := int32(sk.Typ)
	level := sk.Level
	sysSkillInfo.Typ = &typ
	sysSkillInfo.Level = &level
	return sysSkillInfo
}
