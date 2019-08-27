package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/systemskill/pbutil"
	playersysskill "fgame/fgame/game/systemskill/player"
	systemskilltypes "fgame/fgame/game/systemskill/types"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SYSTEM_SKILL_GET_TYPE), dispatch.HandlerFunc(handleSystemSkillGet))
}

//处理系统技能信息
func handleSystemSkillGet(s session.Session, msg interface{}) (err error) {
	log.Debug("systemskill:处理获取系统技能消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSystemSkillGet := msg.(*uipb.CSSystemSkillGet)
	typ := csSystemSkillGet.GetTag()
	err = systemSkillGet(tpl, systemskilltypes.SystemSkillType(typ))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"typ":      typ,
				"error":    err,
			}).Error("systemskill:处理获取系统技能消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Debug("systemskill:处理获取系统技能消息完成")
	return nil

}

//获取系统技能界面信息的逻辑
func systemSkillGet(pl player.Player, typ systemskilltypes.SystemSkillType) (err error) {
	if !typ.Valid() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"typ":      typ,
		}).Warn("systemskill:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	manager := pl.GetPlayerDataManager(types.PlayerSystemSkillDataManagerType).(*playersysskill.PlayerSystemSkillDataManager)
	systemSkills := manager.GetSystemSkillMap(typ)
	scSystemSkillGet := pbutil.BuildSCSystemSkillGet(int32(typ), systemSkills)
	pl.SendMsg(scSystemSkillGet)
	return
}
