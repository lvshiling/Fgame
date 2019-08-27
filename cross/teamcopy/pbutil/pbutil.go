package pbutil

import (
	uipb "fgame/fgame/common/codec/pb/ui"
	teamscene "fgame/fgame/cross/teamcopy/scene"
	propertytypes "fgame/fgame/game/property/types"
	"fgame/fgame/game/scene/scene"
)

func BuildSCTeamCopySceneInfo(sd teamscene.TeamCopySceneData) *uipb.SCTeamCopySceneInfo {
	teamCopySceneInfo := &uipb.SCTeamCopySceneInfo{}
	startTime := sd.GetStartTime()
	teamObj := sd.GetTeamObj()
	purpose := int32(teamObj.GetTeamPurpose())
	teamCopySceneInfo.CreateTime = &startTime
	teamCopySceneInfo.Purpose = &purpose
	s := sd.GetScene()

	for _, mem := range teamObj.GetMemberList() {
		teamCopySceneInfo.PlayerShowList = append(teamCopySceneInfo.PlayerShowList, buildMember(s, mem))
	}
	return teamCopySceneInfo
}

func BuildSCTeamCopyPlayerStatusChanged(playerId int64, status int32) *uipb.SCTeamCopyPlayerDataChanged {
	teamCopyPlayerDataChanged := &uipb.SCTeamCopyPlayerDataChanged{}
	teamCopyPlayerDataChanged.PlayerShow = buildStatusChanged(playerId, status)
	return teamCopyPlayerDataChanged
}

func buildStatusChanged(playerId int64, status int32) *uipb.TeamCopyPlayerShowData {
	teamCopyPlayerShowData := &uipb.TeamCopyPlayerShowData{}
	teamCopyPlayerShowData.PlayerId = &playerId
	teamCopyPlayerShowData.Status = &status
	memStatus := teamscene.MemberStatus(status)
	switch memStatus {
	case teamscene.MemberStatusGoAway,
		teamscene.MemberStatusFailed,
		teamscene.MemberStatusOffline:
		hp := int64(0)
		maxHp := int64(0)
		level := int32(0)
		teamCopyPlayerShowData.Hp = &hp
		teamCopyPlayerShowData.MaxHp = &maxHp
		teamCopyPlayerShowData.Level = &level
	}
	return teamCopyPlayerShowData
}

func BuildSCTeamCopyPlayerDamageChanged(playerId int64, damage int64) *uipb.SCTeamCopyPlayerDataChanged {
	teamCopyPlayerDataChanged := &uipb.SCTeamCopyPlayerDataChanged{}
	teamCopyPlayerDataChanged.PlayerShow = buildDamageChanged(playerId, damage)
	return teamCopyPlayerDataChanged
}

func BuildSCTeamCopyPlayerHpChanged(p scene.Player) *uipb.SCTeamCopyPlayerDataChanged {
	scTeamCopyPlayerDataChanged := &uipb.SCTeamCopyPlayerDataChanged{}
	scTeamCopyPlayerDataChanged.PlayerShow = buildHpChanged(p)
	return scTeamCopyPlayerDataChanged
}

func BuildSCTeamCopyPlayerMaxHPChanged(p scene.Player) *uipb.SCTeamCopyPlayerDataChanged {
	scTeamCopyPlayerDataChanged := &uipb.SCTeamCopyPlayerDataChanged{}
	scTeamCopyPlayerDataChanged.PlayerShow = buildMaxHpChanged(p)
	return scTeamCopyPlayerDataChanged
}

func BuildSCTeamCopyPlayerLevelChanged(p scene.Player) *uipb.SCTeamCopyPlayerDataChanged {
	scTeamCopyPlayerDataChanged := &uipb.SCTeamCopyPlayerDataChanged{}
	scTeamCopyPlayerDataChanged.PlayerShow = buildLevelChanged(p)
	return scTeamCopyPlayerDataChanged
}

func BuildSCTeamCopyPlayerEnterScene(p scene.Player) *uipb.SCTeamCopyPlayerDataChanged {
	scTeamCopyPlayerDataChanged := &uipb.SCTeamCopyPlayerDataChanged{}
	scTeamCopyPlayerDataChanged.PlayerShow = buildPlayerShow(p)
	return scTeamCopyPlayerDataChanged
}

func buildPlayerShow(p scene.Player) *uipb.TeamCopyPlayerShowData {
	teamCopyPlayerShowData := &uipb.TeamCopyPlayerShowData{}
	playerId := p.GetId()
	hp := p.GetHP()
	maxHP := p.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	level := p.GetLevel()
	role := int32(p.GetRole())
	sex := int32(p.GetSex())
	status := int32(teamscene.MemberStatusOnline)
	teamCopyPlayerShowData.PlayerId = &playerId
	teamCopyPlayerShowData.Hp = &hp
	teamCopyPlayerShowData.MaxHp = &maxHP
	teamCopyPlayerShowData.Status = &status
	teamCopyPlayerShowData.Level = &level
	teamCopyPlayerShowData.Role = &role
	teamCopyPlayerShowData.Sex = &sex
	return teamCopyPlayerShowData
}

func buildLevelChanged(p scene.Player) *uipb.TeamCopyPlayerShowData {
	teamCopyPlayerShowData := &uipb.TeamCopyPlayerShowData{}
	playerId := p.GetId()
	level := p.GetLevel()
	teamCopyPlayerShowData.PlayerId = &playerId
	teamCopyPlayerShowData.Level = &level
	return teamCopyPlayerShowData
}

func buildMaxHpChanged(p scene.Player) *uipb.TeamCopyPlayerShowData {
	teamCopyPlayerShowData := &uipb.TeamCopyPlayerShowData{}
	playerId := p.GetId()
	hp := p.GetHP()
	maxHP := p.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
	teamCopyPlayerShowData.PlayerId = &playerId
	teamCopyPlayerShowData.Hp = &hp
	teamCopyPlayerShowData.MaxHp = &maxHP
	return teamCopyPlayerShowData
}

func buildHpChanged(p scene.Player) *uipb.TeamCopyPlayerShowData {
	teamCopyPlayerShowData := &uipb.TeamCopyPlayerShowData{}
	playerId := p.GetId()
	hp := p.GetHP()
	teamCopyPlayerShowData.PlayerId = &playerId
	teamCopyPlayerShowData.Hp = &hp
	return teamCopyPlayerShowData
}

func buildDamageChanged(playerId int64, damage int64) *uipb.TeamCopyPlayerShowData {
	teamCopyPlayerShowData := &uipb.TeamCopyPlayerShowData{}
	teamCopyPlayerShowData.PlayerId = &playerId
	teamCopyPlayerShowData.Damage = &damage
	return teamCopyPlayerShowData
}

func buildMember(s scene.Scene, mem *teamscene.TeamMemberObject) *uipb.TeamCopyPlayerShowData {
	teamCopyPlayerShowData := &uipb.TeamCopyPlayerShowData{}
	playerId := mem.GetPlayerId()
	name := mem.GetName()
	damage := mem.GetDamage()
	status := mem.GetStatus()
	role := int32(mem.GetRole())
	sex := int32(mem.GetSex())

	teamCopyPlayerShowData.PlayerId = &playerId
	teamCopyPlayerShowData.Name = &name
	teamCopyPlayerShowData.Damage = &damage
	teamCopyPlayerShowData.Role = &role
	teamCopyPlayerShowData.Sex = &sex

	pl := s.GetPlayer(playerId)
	if pl != nil {
		level := pl.GetLevel()
		hp := pl.GetHP()
		maxHP := pl.GetBattleProperty(propertytypes.BattlePropertyTypeMaxHP)
		memStatus := int32(status)
		teamCopyPlayerShowData.Hp = &hp
		teamCopyPlayerShowData.MaxHp = &maxHP
		teamCopyPlayerShowData.Status = &memStatus
		teamCopyPlayerShowData.Level = &level
	} else {
		hp := int64(0)
		maxHp := int64(0)
		level := int32(0)
		if status == teamscene.MemberStatusOnline {
			status = teamscene.MemberStatusOffline
		}
		memStatus := int32(status)
		teamCopyPlayerShowData.Hp = &hp
		teamCopyPlayerShowData.MaxHp = &maxHp
		teamCopyPlayerShowData.Status = &memStatus
		teamCopyPlayerShowData.Level = &level
	}
	return teamCopyPlayerShowData
}
