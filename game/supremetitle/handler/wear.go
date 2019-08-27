package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	scenelogic "fgame/fgame/game/scene/logic"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/supremetitle/pbutil"
	playersupremetitle "fgame/fgame/game/supremetitle/player"
	supremetitletemplate "fgame/fgame/game/supremetitle/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SUPREME_TITLE_WEAR_TYPE), dispatch.HandlerFunc(handleSupremeTitleWear))
}

//处理至尊称号穿戴信息
func handleSupremeTitleWear(s session.Session, msg interface{}) (err error) {
	log.Debug("supremetitle:处理至尊称号穿戴信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csTitleWear := msg.(*uipb.CSSupremeTitleWear)
	titleId := csTitleWear.GetTitleId()

	err = supremeTitleWear(tpl, titleId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"titleId":  titleId,
				"error":    err,
			}).Error("supremetitle:处理至尊称号穿戴信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Debug("supremetitle:处理至尊称号穿戴信息完成")
	return nil
}

//处理至尊称号穿戴信息逻辑
func supremeTitleWear(pl player.Player, titleId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSupremeTitleDataManagerType).(*playersupremetitle.PlayerSupremeTitleDataManager)
	titleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(titleId)
	if titleTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("supremetitle:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag := manager.HasedWeared(titleId)
	if flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("supremetitle:该至尊称号已穿戴,无需再次穿戴")
		playerlogic.SendSystemMessage(pl, lang.SupremeTitleRepeatWear)
		return
	}

	flag = manager.IfTitleExist(titleId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"titleId":  titleId,
		}).Warn("supremetitle:还没有该至尊称号,请先获取")
		playerlogic.SendSystemMessage(pl, lang.SupremeTitleNotHas)
		return
	}

	//移除当前穿戴buff
	curTitleId := manager.GetTitleId()
	supremeTitleTemplate := supremetitletemplate.GetTitleDingZhiTemplateService().GetTitleDingZhiTempalte(curTitleId)
	if supremeTitleTemplate != nil {
		scenelogic.RemoveBuff(pl, supremeTitleTemplate.BuffId)
	}

	//增加buff
	buffId := titleTemplate.BuffId
	if buffId != 0 {
		scenelogic.AddBuffs(pl, buffId, pl.GetId(), 1, common.MAX_RATE)
	}

	flag = manager.TitleWear(titleId)
	if !flag {
		panic(fmt.Errorf("supremetitle: supremeTitleWear  should be ok"))
	}
	scTitleWear := pbutil.BuildSCSupremeTitleWear(titleId)
	pl.SendMsg(scTitleWear)
	return
}
