package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/click/pbutil"
	"fgame/fgame/game/click/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/quest/quest"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CLICK_TYPE), dispatch.HandlerFunc(handleClick))
}

//处理点击按钮事件信息
func handleClick(s session.Session, msg interface{}) (err error) {
	log.Debug("click:处理点击消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csClick := msg.(*uipb.CSClick)
	clickType := csClick.GetClickType()
	clickSubType := csClick.GetClickSubType()

	err = clickButton(tpl, clickType, clickSubType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"clickType":    clickType,
				"clickSubType": clickSubType,
				"error":        err,
			}).Error("dan:处理点击消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("dan:处理点击消息完成")
	return nil
}

//处理点击按钮信息逻辑
func clickButton(pl player.Player, clickType int32, clickSubType int32) (err error) {
	typ := types.ClickType(clickType)
	if !typ.Valid() {
		log.WithFields(log.Fields{
			"playerId":     pl.GetId(),
			"clickType":    clickType,
			"clickSubType": clickSubType,
		}).Warn("click:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	subType := types.GetClickSubType(typ, clickSubType)
	if !subType.Valid() {
		log.WithFields(log.Fields{
			"playerId":     pl.GetId(),
			"clickType":    clickType,
			"clickSubType": clickSubType,
		}).Warn("click:参数错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}


	quest.ClickHandle(pl, typ, subType)
	scClick := pbuitl.BuildSCClick()
	pl.SendMsg(scClick)
	return
}
