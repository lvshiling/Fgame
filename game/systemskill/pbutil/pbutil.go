package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	sysskillcommon "fgame/fgame/game/systemskill/common"
	playersysskill "fgame/fgame/game/systemskill/player"
	sysskilltypes "fgame/fgame/game/systemskill/types"
)

func BuildSCSystemSkillGet(typ int32, sysSkill *playersysskill.SystemSkill) *uipb.SCSystemSkillGet {
	scSystemSkillGet := &uipb.SCSystemSkillGet{}
	scSystemSkillGet.Tag = &typ
	for _, sk := range sysSkill.GetSysSkillMap() {
		scSystemSkillGet.SystemSkillList = append(scSystemSkillGet.SystemSkillList, buildSystemSkill(sk))
	}
	return scSystemSkillGet
}

func BuildSCSystemSkillAllGet(sysSkillAllMap map[sysskilltypes.SystemSkillType]*playersysskill.SystemSkill) *uipb.SCSystemSkillAllGet {
	scSystemSkillAllGet := &uipb.SCSystemSkillAllGet{}
	for _, sysSkill := range sysSkillAllMap {
		for _, sk := range sysSkill.GetSysSkillMap() {
			scSystemSkillAllGet.SystemSkillList = append(scSystemSkillAllGet.SystemSkillList, buildSystemSkill(sk))
		}
	}
	return scSystemSkillAllGet
}

func BuildSCSystemSkillActive(sk *playersysskill.PlayerSystemSkillObject) *uipb.SCSystemSkillActive {
	scSystemSkillActive := &uipb.SCSystemSkillActive{}
	scSystemSkillActive.SystemSkill = buildSystemSkill(sk)
	return scSystemSkillActive
}

func BuildSCSystemSkillUpgrade(sk *playersysskill.PlayerSystemSkillObject) *uipb.SCSystemSkillUpgrade {
	scSystemSkillUpgrade := &uipb.SCSystemSkillUpgrade{}
	scSystemSkillUpgrade.SystemSkill = buildSystemSkill(sk)
	return scSystemSkillUpgrade
}

func buildSystemSkill(sk *playersysskill.PlayerSystemSkillObject) *uipb.SystemSkillInfo {
	sysSkillInfo := &uipb.SystemSkillInfo{}
	typ := int32(sk.Type)
	subType := int32(sk.SubType)
	level := sk.Level
	sysSkillInfo.Tag = &typ
	sysSkillInfo.SubType = &subType
	sysSkillInfo.Level = &level
	return sysSkillInfo
}

func BuildAllSystemSkillInfo(info *sysskillcommon.AllSystemSkillInfo) *uipb.AllSystemSkillInfo {
	allSystemSkillInfo := &uipb.AllSystemSkillInfo{}
	for _, systemSkill := range info.SystemSkillList {
		allSystemSkillInfo.SystemSkillList = append(allSystemSkillInfo.SystemSkillList, buildSystemSkillCacheInfo(systemSkill))
	}
	return allSystemSkillInfo
}

func buildSystemSkillCacheInfo(info *sysskillcommon.SystemSkillInfo) *uipb.SystemSkillCacheInfo {
	systemSkillCacheInfo := &uipb.SystemSkillCacheInfo{}
	typ := info.Type

	systemSkillCacheInfo.Tag = &typ
	for _, sysSkill := range info.SysSkillList {
		systemSkillCacheInfo.SkillSubList = append(systemSkillCacheInfo.SkillSubList, buildSystemSkillSubInfo(sysSkill.SubType, sysSkill.Level))
	}
	return systemSkillCacheInfo
}

func buildSystemSkillSubInfo(subType int32, level int32) *uipb.SystemSkillSubInfo {
	systemSkillSubInfo := &uipb.SystemSkillSubInfo{}
	systemSkillSubInfo.SubType = &subType
	systemSkillSubInfo.Level = &level
	return systemSkillSubInfo
}
