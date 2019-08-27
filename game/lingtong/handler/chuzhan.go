package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	playerlingtong "fgame/fgame/game/lingtong/player"

	"fgame/fgame/game/lingtong/pbutil"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_CHUZHAN_TYPE), dispatch.HandlerFunc(handleLingTongChuZhan))

}

//处理灵童出战信息
func handleLingTongChuZhan(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理获取灵童出战消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongChuZhan := msg.(*uipb.CSLingTongChuZhan)
	lingTongId := csLingTongChuZhan.GetLingTongId()

	err = lingTongChuZhan(tpl, lingTongId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"lingTongId": lingTongId,
				"error":      err,
			}).Error("lingtong:处理获取灵童出战消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Debug("lingtong:处理获取灵童出战消息完成")
	return nil

}

//获取灵童出战界面信息逻辑
func lingTongChuZhan(pl player.Player, lingTongId int32) (err error) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:模板为空")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(types.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	_, flag := manager.GetLingTongInfo(lingTongId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":   pl.GetId(),
			"lingTongId": lingTongId,
		}).Warn("lingtong:未激活该灵童")
		playerlogic.SendSystemMessage(pl, lang.LingTongNoActive)
		return
	}

	lingTongObj := manager.GetLingTong()
	if lingTongObj.GetLingTongId() == lingTongId {
		return
	}
	manager.LingTongChuZhan(lingTongId)
	scLingTongChuZhan := pbutil.BuildSCLingTongChuZhan(lingTongId)
	pl.SendMsg(scLingTongChuZhan)
	return
}
