package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/songbuting/pbutil"
	playersongbuting "fgame/fgame/game/songbuting/player"
	songbutingtemplate "fgame/fgame/game/songbuting/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SONGBUTING_RECEIVE_TYPE), dispatch.HandlerFunc(handleSongBuTingReceive))
}

//处理送不停信息
func handleSongBuTingReceive(s session.Session, msg interface{}) (err error) {
	log.Debug("songbuting:处理获取送不停消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = songBuTingReceive(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("songbuting:处理获取送不停消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("songbuting:处理获取送不停消息完成")
	return nil
}

//获取送不停信息的逻辑
func songBuTingReceive(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerSongBuTingDataManagerType).(*playersongbuting.PlayerSongBuTingDataManager)
	songBuTingObj := manager.GetSongBuTingObj()
	songBuTingTempalte := songbutingtemplate.GetSongBuTingTemplateService().GetSongBuTingTemplate()
	if !songBuTingObj.GetIsReceive() {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("songbuting:单笔充值980元宝,才能享受享尊贵特权")
		needGoldStr := fmt.Sprintf("%d", songBuTingTempalte.NeedGold)
		playerlogic.SendSystemMessage(pl, lang.SongBuTingNoReceive, needGoldStr)
		return
	}

	if songBuTingObj.GetTimes() != 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("songbuting:今日福利已领完,请明日再来")
		playerlogic.SendSystemMessage(pl, lang.SongBuTingReceiveNumLimit)
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	rewItemMap := songBuTingTempalte.GetRewItemMap()
	rewData := songBuTingTempalte.GetRewData()
	if len(rewItemMap) != 0 {
		flag := inventoryManager.HasEnoughSlots(rewItemMap)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
			}).Warn("songbuting:背包已满")
			playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
			return
		}

		inventoryLogReason := commonlog.InventoryLogReasonSongBuTingRew
		reasonText := inventoryLogReason.String()
		flag = inventoryManager.BatchAdd(rewItemMap, inventoryLogReason, reasonText)
		if !flag {
			panic(fmt.Errorf("songbuting: songBuTingReceive BatchAdd should be ok"))
		}
		inventorylogic.SnapInventoryChanged(pl)
	}

	//添加奖励属性
	if rewData != nil {
		goldLog := commonlog.GoldLogReasonSongBuTingRew
		silverLog := commonlog.SilverLogReasonSongBuTingRew
		levelReason := commonlog.LevelLogReasonSongBuTingRew
		goldReasonText := goldLog.String()
		silverReasonText := silverLog.String()
		reasonText := levelReason.String()
		flag := propertyManager.AddRewData(rewData, goldLog, goldReasonText, silverLog, silverReasonText, levelReason, reasonText)
		if !flag {
			panic(fmt.Errorf("songbuting: songBuTingReceive AddRewData should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}

	manager.Receive()
	scSongBuTingChanged := pbutil.BuildSCSongBuTingChanged(songBuTingObj)
	pl.SendMsg(scSongBuTingChanged)
	return
}
