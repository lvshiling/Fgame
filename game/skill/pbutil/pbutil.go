package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	skillcommon "fgame/fgame/game/skill/common"
	playerskill "fgame/fgame/game/skill/player"
)

func BuildSCSkillGet(skillMap map[int32]skillcommon.SkillObject) *uipb.SCSkillGet {
	scSkillGet := &uipb.SCSkillGet{}
	for _, obj := range skillMap {
		scSkillGet.SkillList = append(scSkillGet.SkillList, buildSkill(obj))
	}
	return scSkillGet
}

func BuildSCSkillAdd(skillId int32, level int32) *uipb.SCSkillAdd {
	skillAdd := &uipb.SCSkillAdd{}
	skillAdd.SkillId = &skillId
	skillAdd.Level = &level
	return skillAdd
}

func BuildSCSkillRemove(skillId int32, level int32) *uipb.SCSkillRemove {
	skillRemove := &uipb.SCSkillRemove{}
	skillRemove.SkillId = &skillId
	skillRemove.Level = &level
	return skillRemove
}

func BuildSCSkillUpgrade(oldSkillId int32) *uipb.SCSkillUpgrade {
	skillUpgrade := &uipb.SCSkillUpgrade{}
	skillUpgrade.OldSkillId = &oldSkillId
	return skillUpgrade
}

func BuildSCSkillUpgradeAll(oldSkillMap map[int32]int32) *uipb.SCSkillUpgradeAll {
	skillUpgradeAll := &uipb.SCSkillUpgradeAll{}
	for oldSkillId, _ := range oldSkillMap {
		skillUpgradeAll.OldSkillIdList = append(skillUpgradeAll.OldSkillIdList, int32(oldSkillId))
	}
	return skillUpgradeAll
}

func BuildSCSkillCdTime(skillCdMap map[int32]*playerskill.PlayerSkillCdObject) *uipb.SCSkillCdTime {
	skillCdTime := &uipb.SCSkillCdTime{}
	for _, skillCdObj := range skillCdMap {
		skillCdTime.SkillCdList = append(skillCdTime.SkillCdList, buildCdTime(skillCdObj))
	}
	return skillCdTime
}

func BuildSCSkillTianFuAwaken(skillId int32, tianFuId int32, level int32, sucess bool) *uipb.SCSkillTianFuAwaken {
	scSkillTianFuAwaken := &uipb.SCSkillTianFuAwaken{}
	scSkillTianFuAwaken.SkillId = &skillId
	scSkillTianFuAwaken.Sucess = &sucess
	scSkillTianFuAwaken.TianFuInfo = buildTianFu(tianFuId, level)
	return scSkillTianFuAwaken
}

func BuildSCSkillTianFuUpgrade(skillId int32, tianFuId int32, level int32, sucess bool) *uipb.SCSkillTianFuUpgrade {
	scSkillTianFuUpgrade := &uipb.SCSkillTianFuUpgrade{}
	scSkillTianFuUpgrade.SkillId = &skillId
	scSkillTianFuUpgrade.Sucess = &sucess
	scSkillTianFuUpgrade.TianFuInfo = buildTianFu(tianFuId, level)
	return scSkillTianFuUpgrade
}

func BuildSCSkillTianFuGet(roleSkillMap map[int32]*playerskill.PlayerSkillObject) *uipb.SCSkillTianFuGet {
	scSkillTianFuGet := &uipb.SCSkillTianFuGet{}
	for skillId, skillInfo := range roleSkillMap {
		if len(skillInfo.TianFuMap) == 0 {
			continue
		}
		scSkillTianFuGet.SkillTianFuList = append(scSkillTianFuGet.SkillTianFuList, buildSkillTianFu(skillId, skillInfo.TianFuMap))
	}
	return scSkillTianFuGet
}

func buildSkillTianFu(skillId int32, tianFuMap map[int32]int32) *uipb.SkillTianFuInfo {
	skillTianFuInfo := &uipb.SkillTianFuInfo{}
	skillTianFuInfo.SkillId = &skillId
	for tianFuId, level := range tianFuMap {
		skillTianFuInfo.TianFuList = append(skillTianFuInfo.TianFuList, buildTianFu(tianFuId, level))
	}
	return skillTianFuInfo

}

func buildTianFu(tianFuId int32, level int32) *uipb.TianFuInfo {
	tianFuInfo := &uipb.TianFuInfo{}
	tianFuInfo.TianFuId = &tianFuId
	tianFuInfo.Level = &level
	return tianFuInfo
}

func buildCdTime(skillCdObj *playerskill.PlayerSkillCdObject) *uipb.SkillCdInfo {
	skillCdInfo := &uipb.SkillCdInfo{}
	skillId := skillCdObj.SkillId
	lastTime := skillCdObj.LastTime
	skillCdInfo.SkillId = &skillId
	skillCdInfo.LastTime = &lastTime
	return skillCdInfo
}

func buildSkill(skill skillcommon.SkillObject) *uipb.SkillInfo {
	skillInfo := &uipb.SkillInfo{}
	skillId := skill.GetSkillId()
	level := skill.GetLevel()
	skillInfo.SkillId = &skillId
	skillInfo.Level = &level
	return skillInfo
}

func BuildSCSkillLearn(skillId int32, level int32) *uipb.SCSkillLearn {
	scSkillLearn := &uipb.SCSkillLearn{}
	scSkillLearn.SkillId = &skillId
	scSkillLearn.Level = &level
	return scSkillLearn
}
