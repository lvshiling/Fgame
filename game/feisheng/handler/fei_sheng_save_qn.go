package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	feishenglogic "fgame/fgame/game/feisheng/logic"
	"fgame/fgame/game/feisheng/pbutil"
	playerfeisheng "fgame/fgame/game/feisheng/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FEI_SHENG_SAVE_QN_TYPE), dispatch.HandlerFunc(handleFeiShengSaveQn))
}

//处理飞升保存潜能
func handleFeiShengSaveQn(s session.Session, msg interface{}) (err error) {
	log.Debug("wing:处理飞升保存潜能消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSFeiShengSaveQn)
	qnInfo := csMsg.GetQnInfo()

	err = feiShengSaveQn(tpl, qnInfo.GetTi(), qnInfo.GetLi(), qnInfo.GetGu())
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("feisheng:处理飞升保存潜能消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("feisheng:处理飞升保存潜能消息完成")
	return nil
}

//飞升保存潜能界面逻辑
func feiShengSaveQn(pl player.Player, ti, li, gu int32) (err error) {

	feiManager := pl.GetPlayerDataManager(playertypes.PlayerFeiShengDataManagerType).(*playerfeisheng.PlayerFeiShengDataManager)
	feiShengInfo := feiManager.GetFeiShengInfo()
	totalQn := ti + li + gu
	if totalQn <= 0 {
		log.WithFields(
			log.Fields{
				"playerid":   pl.GetId(),
				"totalAddQn": totalQn,
				"leftQn":     feiShengInfo.GetLeftPotential(),
			}).Warn("feisheng:飞升保存潜能失败，还未进行加点")
		playerlogic.SendSystemMessage(pl, lang.FeiShengHadLeftQn)
		return
	}

	if feiShengInfo.GetLeftPotential() < totalQn {
		log.WithFields(
			log.Fields{
				"playerid":   pl.GetId(),
				"totalAddQn": totalQn,
				"leftQn":     feiShengInfo.GetLeftPotential(),
			}).Warn("feisheng:飞升保存潜能失败，潜能点不足")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	feiManager.SaveQn(ti, li, gu)
	feishenglogic.FeiShengPropertyChanged(pl)

	scMsg := pbutil.BuildSCFeiShengSaveQn(ti, li, gu)
	pl.SendMsg(scMsg)
	return
}
