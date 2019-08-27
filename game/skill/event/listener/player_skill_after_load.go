package listener

import (
	"fgame/fgame/core/event"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/skill/pbutil"
	playerskill "fgame/fgame/game/skill/player"
)

//玩家登录成功后下发技能
func playerRoleSkillAfterLogin(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)

	manager := p.GetPlayerDataManager(types.PlayerSkillDataManagerType).(*playerskill.PlayerSkillDataManager)
	roleSkillMap := manager.GetRoleSkillMap()
	skillcdTime := pbutil.BuildSCSkillCdTime(nil)
	p.SendMsg(skillcdTime)

	skills := p.GetAllSkills()
	scSkillGet := pbutil.BuildSCSkillGet(skills)
	p.SendMsg(scSkillGet)

	scSkillTianFuGet := pbutil.BuildSCSkillTianFuGet(roleSkillMap)
	p.SendMsg(scSkillTianFuGet)
	return
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerAfterLoadFinish, event.EventListenerFunc(playerRoleSkillAfterLogin))
}
