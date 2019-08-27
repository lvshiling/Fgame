package pbutil

import (
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/cross/arena/arena"
	playerpbutil "fgame/fgame/cross/player/pbutil"
	playertypes "fgame/fgame/game/player/types"
)

//竞技场玩家
func ConvertFromArenaPlayer(p *crosspb.ArenaPlayer) *arena.MatchTeamMember {
	playerId := p.GetPlayerId()
	force := p.GetForce()
	name := p.GetName()
	level := p.GetLevel()
	role := playertypes.RoleType(p.GetRole())
	sex := playertypes.SexType(p.GetSex())
	fashionId := p.GetFashionId()
	serverId := p.GetServerId()
	battleProperties := playerpbutil.ConvertFromBattleProperty(p.GetBattlePropertyData())
	skillList := playerpbutil.ConvertFromSkillDataList(p.GetSkillList())
	winCount := p.GetWinCount()
	ap := arena.CreateMatchTeamMemberObject(
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
		winCount,
	)
	return ap
}

//竞技场玩家列表
func ConvertFromArenaPlayerList(pList []*crosspb.ArenaPlayer) []*arena.MatchTeamMember {
	apList := make([]*arena.MatchTeamMember, 0, len(pList))
	for _, p := range pList {
		apList = append(apList, ConvertFromArenaPlayer(p))
	}
	return apList
}
