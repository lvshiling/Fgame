package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/systemskill/pbutil"
	playersysskill "fgame/fgame/game/systemskill/player"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SYSTEM_SKILL_ALL_GET_TYPE), dispatch.HandlerFunc(handleSystemSkillAllGet))
}

//处理获取所有技能系统信息
func handleSystemSkillAllGet(s session.Session, msg interface{}) (err error) {
	log.Debug("systemskill:处理获取获取所有技能系统消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = systemSkillAllGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("systemskill:处理获取获取所有技能系统消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("systemskill:处理获取获取所有技能系统消息完成")
	return nil

}

//获取获取所有技能系统界面信息的逻辑
func systemSkillAllGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSystemSkillDataManagerType).(*playersysskill.PlayerSystemSkillDataManager)
	systemSkills := manager.GetSystemSkillAllMap()
	scSystemSkillAllGet := pbutil.BuildSCSystemSkillAllGet(systemSkills)
	pl.SendMsg(scSystemSkillAllGet)
	return
}
