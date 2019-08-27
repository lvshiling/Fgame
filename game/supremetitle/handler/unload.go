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
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/supremetitle/pbutil"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_SUPREME_TITLE_UNLOAD_TYPE), dispatch.HandlerFunc(handleSupremeTitleUnload))

}

//处理至尊称号卸下信息
func handleSupremeTitleUnload(s session.Session, msg interface{}) (err error) {
	log.Debug("supremetitle:处理至尊称号卸下信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = supremeTitleUnload(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("supremetitle:处理至尊称号卸下信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("supremetitle:处理至尊称号卸下信息完成")
	return nil

}

//至尊称号卸下逻辑
func supremeTitleUnload(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	titleWear := manager.GetTitleId()
	if titleWear == 0 {
		playerlogic.SendSystemMessage(pl, lang.SupremeTitleNotHas)
		return
	}
	//移除buff
	titleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(titleWear)
	if titleTemplate != nil && titleTemplate.BuffId != 0 {
		scenelogic.RemoveBuff(pl, titleTemplate.BuffId)
	}

	manager.TitleNoWear()
	scTitleUnload := pbutil.BuildSCSupremeTitleUnload(titleWear)
	pl.SendMsg(scTitleUnload)
	return
}
