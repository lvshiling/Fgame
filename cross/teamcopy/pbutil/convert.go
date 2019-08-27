package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	playerpbutil "fgame/fgame/cross/player/pbutil"
	"fgame/fgame/cross/teamcopy/teamcopy"
	playertypes "fgame/fgame/game/player/types"
	teamtypes "fgame/fgame/game/team/types"
)

//组队副本玩家
func ConvertFromTeamPlayer(p *crosspb.TeamPlayer) *teamcopy.BattleTeamMember {
	playerId := p.GetPlayerId()
	force := p.GetForce()
	name := p.GetName()
	level := p.GetLevel()
	role := playertypes.RoleType(p.GetRole())
	sex := playertypes.SexType(p.GetSex())
	fashionId := p.GetFashionId()
	serverId := p.GetServerId()
	damage := int64(0)
	purpose := teamtypes.TeamPurposeType(p.GetPurpose())
	battleProperties := playerpbutil.ConvertFromBattleProperty(p.GetBattlePropertyData())
	skillList := playerpbutil.ConvertFromSkillDataList(p.GetSkillList())
	ap := teamcopy.CreateBattleTeamMemberObject(
		serverId,
		playerId,
		force,
		name,
		level,
		role,
		sex,
		fashionId,
		battleProperties,
		skillList,
		false,
		damage,
		purpose,
	)
	return ap
}

//组队副本玩家列表
func ConvertFromTeamPlayerList(pList []*crosspb.TeamPlayer) []*teamcopy.BattleTeamMember {
	apList := make([]*teamcopy.BattleTeamMember, 0, len(pList))
	for _, p := range pList {
		apList = append(apList, ConvertFromTeamPlayer(p))
	}
	return apList
}
