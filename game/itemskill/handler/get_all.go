package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/core/session"
	"fgame/fgame/game/itemskill/pbutil"
	playeritemskill "fgame/fgame/game/itemskill/player"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ITEM_SKILL_ALL_GET_TYPE), dispatch.HandlerFunc(handleItemSkillAllGet))
}

//处理获取所有技能信息
func handleItemSkillAllGet(s session.Session, msg interface{}) (err error) {
	log.Debug("itemskill:处理获取获取所有技能消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = itemSkillAllGet(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("itemskill:处理获取获取所有技能消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("itemskill:处理获取获取所有技能消息完成")
	return nil

}

//获取获取所有技能界面信息的逻辑
func itemSkillAllGet(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerItemSkillDataManagerType).(*playeritemskill.PlayerItemSkillDataManager)
	itemSkills := manager.GetItemSkillAllMap()
	scItemSkillAllGet := pbutil.BuildSCItemSkillAllGet(itemSkills)
	pl.SendMsg(scItemSkillAllGet)
	return
}
