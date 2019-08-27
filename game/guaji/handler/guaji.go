package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/guaji/pbutil"
	playerguaji "fgame/fgame/game/guaji/player"
	guajitypes "fgame/fgame/game/guaji/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GUA_JI_TYPE), dispatch.HandlerFunc(handleGuaJi))
}

//处理挂机
func handleGuaJi(s session.Session, msg interface{}) (err error) {
	log.Info("guaji:处理挂机")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csGuaJi := msg.(*uipb.CSGuaJi)
	guaJiDataList, flag := pbutil.ConvertFromGuaJiDataList(csGuaJi.GetGuaJiDataList())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"guaJiList": csGuaJi.GetGuaJiDataList(),
			}).Warn("guaji:处理挂机,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.GuaJiArgsInvalid)
		return
	}
	guaJiAdvanceMap, flag := pbutil.ConvertFromGuaJiAdvnaceSettingDataList(csGuaJi.GetGuaJiAdvanceSettingDataList())
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId":                pl.GetId(),
				"guaJiAdvanceSettingList": csGuaJi.GetGuaJiAdvanceSettingDataList(),
			}).Warn("guaji:处理挂机,参数错误")
		playerlogic.SendSystemMessage(tpl, lang.GuaJiArgsInvalid)
		return
	}
	err = guaJi(tpl, guaJiDataList, guaJiAdvanceMap)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"guaJiList": guaJiDataList,
				"error":     err,
			}).Error("guaji:处理挂机,错误")

		return err
	}
	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"guaJiList": guaJiDataList,
		}).Info("guaji:处理挂机,完成")
	return nil
}

//挂机
func guaJi(pl player.Player, guaJiDataList []*guajitypes.GuaJiData, guaJiAdvanceMap map[guajitypes.GuaJiAdvanceType]int32) (err error) {
	if len(guaJiDataList) == 0 {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("guaji:处理挂机,挂机列表为空")
		playerlogic.SendSystemMessage(pl, lang.GuaJiArgsInvalid)
		return
	}
	if !pl.IsGuaJiPlayer() {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"guaJiList": guaJiDataList,
			}).Warn("guaji:处理挂机,不是挂机玩家")
		playerlogic.SendSystemMessage(pl, lang.GuaJiNoGuaJiPlayer)
		return
	}

	guaJiManager := pl.GetPlayerDataManager(playertypes.PlayerGuaJiManagerType).(*playerguaji.PlayerGuaJiManager)
	originGuaJiDataList := guaJiManager.GetGuaJiTypeList()
	if len(originGuaJiDataList) > 0 {
		guaJiManager.StopGuaJiList()
	}
	log.WithFields(
		log.Fields{
			"playerId":        pl.GetId(),
			"guaJiList":       guaJiDataList,
			"guaJiAdvanceMap": guaJiAdvanceMap,
		}).Info("guaji:处理挂机,开始挂机")
	guaJiManager.StartGuaJiList(guaJiDataList, guaJiAdvanceMap)
	scGuaJi := pbutil.BuildSCGuaJi(guaJiDataList, guaJiAdvanceMap)
	pl.SendMsg(scGuaJi)
	return
}
