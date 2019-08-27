package handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/cross/player/pbutil"
	"fgame/fgame/cross/player/player"
	"fgame/fgame/cross/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(crosspb.MessageType_SI_PLAYER_TESHU_SKILL_RESET_TYPE), dispatch.HandlerFunc(handlePlayerTeshuSkillReset))
}

//玩家仙盟变化
func handlePlayerTeshuSkillReset(s session.Session, msg interface{}) (err error) {
	log.Debug("login:处理跨服系统特殊技能推送消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player().(*player.Player)
	siPlayerTeshuSkillReset := msg.(*crosspb.SIPlayerTeshuSkillReset)
	skillDataList := siPlayerTeshuSkillReset.GetSkillList()

	err = teShuSkillReset(pl, skillDataList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": gcs.Player().GetId(),
			}).Error("login:玩家跨服特殊技能,失败")
		return err
	}

	log.Debug("login:处理跨服特殊技能消息完成")
	return nil
}

//玩家显示变化
func teShuSkillReset(pl *player.Player, skillPbList []*crosspb.TeShuSkillData) (err error) {
	skillList := pbutil.ConvertFromTeShuSkillDataList(skillPbList)
	pl.ResetTeShuSkills(skillList)
	return nil
}
