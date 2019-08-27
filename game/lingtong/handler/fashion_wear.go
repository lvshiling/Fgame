package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_FASHION_WEAR_TYPE), dispatch.HandlerFunc(handleLingTongFashionWear))

}

//处理灵童时装穿戴信息
func handleLingTongFashionWear(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理获取灵童时装穿戴消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongFashionWear := msg.(*uipb.CSLingTongFashionWear)
	fashionId := csLingTongFashionWear.GetFashionId()

	err = lingTongFashionWear(tpl, fashionId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"fashionId": fashionId,
				"error":     err,
			}).Error("lingtong:处理获取灵童时装穿戴消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Debug("lingtong:处理获取灵童时装穿戴消息完成")
	return nil
}

//获取灵童时装穿戴界面信息逻辑
func lingTongFashionWear(pl player.Player, fashionId int32) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongObj := manager.GetLingTong()
	if lingTongObj == nil || !lingTongObj.IsActivateSys() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("lingtong:请先激活灵童时装穿戴系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongActiveSystem)
		return
	}

	fashionTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongFashionTemplate(fashionId)
	if fashionTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:模板为空")
		return
	}

	flag := lingtongtemplate.GetLingTongTemplateService().IsBornFashion(fashionId)
	if flag {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:出生时装")
		return
	}

	fashionInfo := manager.GetFashionInfoById(fashionId)
	if fashionInfo == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:未激活的灵童时装,无法使用")
		playerlogic.SendSystemMessage(pl, lang.LingTongFashionUseNoActivate)
		return
	}

	lingTongId := lingTongObj.GetLingTongId()
	lingTongFashionInfo := manager.GetLingTongFashionById(lingTongId)
	if lingTongFashionInfo == nil {
		panic(fmt.Errorf("lingtong: 灵童时装应该是存在的"))
	}

	curFashionId := lingTongFashionInfo.GetFashionId()
	if curFashionId == fashionId {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"fashionId": fashionId,
		}).Warn("lingtong:当前出战的灵童时装已经是此时装")
		return
	}

	flag = manager.FashionWear(fashionId)
	if !flag {
		panic(fmt.Errorf("lingtong: lingTongFashionWear should be ok"))
	}

	scLingTongFashionWear := pbutil.BuildSCLingTongFashionWear(fashionId)
	pl.SendMsg(scLingTongFashionWear)
	return
}
