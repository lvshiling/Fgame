package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/emperor/emperor"
	emperorlogic "fgame/fgame/game/emperor/logic"
	"fgame/fgame/game/emperor/pbutil"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_EMPEROR_OPEN_BOX_TYPE), dispatch.HandlerFunc(handleEmperorOpenBox))
}

//处理帝王信息
func handleEmperorOpenBox(s session.Session, msg interface{}) (err error) {
	log.Debug("emperor:处理帝王开宝箱信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = emperorOpenBox(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("emperor:处理帝王开宝箱信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("emperor:处理帝王开宝箱信息完成")
	return nil
}

//处理帝王开宝箱界面信息逻辑
func emperorOpenBox(pl player.Player) (err error) {
	emperorId, boxNum := emperor.GetEmperorService().GetEmperorIdAndBoxNum()
	if emperorId != pl.GetId() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("emperor:您不是帝王")
		playerlogic.SendSystemMessage(pl, lang.EmperorOpenBoxNoOwner)
		return
	}
	if boxNum <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("emperor:帝王宝箱无库存")
		playerlogic.SendSystemMessage(pl, lang.EmperorOpenBoxNoStorage)
		return
	}

	//帝王开宝箱
	emperorInfo, dropItemList, err := emperor.GetEmperorService().OpenBox(pl)
	if err != nil {
		return
	}

	if len(dropItemList) != 0 {
		//添加物品
		err = emperorlogic.OpenBoxReward(pl, dropItemList, true)
		if err != nil {
			return
		}
	}

	scEmperorOpenBox := pbuitl.BuildSCEmperorOpenBox(emperorInfo, dropItemList)
	pl.SendMsg(scEmperorOpenBox)
	return
}
