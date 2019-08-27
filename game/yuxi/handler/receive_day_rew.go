package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	pbutil "fgame/fgame/game/yuxi/pbutil"
	playeryuxi "fgame/fgame/game/yuxi/player"
	yuxitemplate "fgame/fgame/game/yuxi/template"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_YUXI_RECEIVE_DAY_REW_TYPE), dispatch.HandlerFunc(handleYuXiDayRew))
}

//处理玉玺之战每日奖励
func handleYuXiDayRew(s session.Session, msg interface{}) (err error) {
	log.Debug("yuxi:处理玉玺之战每日奖励消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = yuXiDayRew(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("yuxi:处理玉玺之战每日奖励消息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("yuxi:处理玉玺之战每日奖励消息完成")
	return nil
}

//处理玉玺之战每日奖励信息逻辑
func yuXiDayRew(pl player.Player) (err error) {

	//获胜仙盟
	hegemonAlliance := alliance.GetAllianceService().GetAllianceHegemon()
	if pl.GetAllianceId() != hegemonAlliance.GetDefenceAllianceId() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("yuxi:处理玩家玉玺之战每日奖励，不是获胜联盟成员")
		playerlogic.SendSystemMessage(pl, lang.YuXiWinNotWinerMember)
		return
	}

	//是否领取
	yuXiManager := pl.GetPlayerDataManager(playertypes.PlayerYuXiDataManagerType).(*playeryuxi.PlayerYuXiDataManager)
	yuXiManager.RefreshTimes()

	yuXiInfo := yuXiManager.GetPlayerYuXiInfo()
	if yuXiInfo.IsReceive() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("yuxi:处理玩家玉玺之战每日奖励，已领取每日奖励")
		playerlogic.SendSystemMessage(pl, lang.YuXiWinHadReceiveDayRew)
		return
	}
	dayItemMap := yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate().GetDayItemMap()
	if len(dayItemMap) == 0 {
		log.WithFields(
			log.Fields{
				"playerId":   pl.GetId(),
				"dayItemMap": dayItemMap,
			}).Warn("yuxi:处理玩家玉玺之战每日奖励，没有可领取的奖励")
		playerlogic.SendSystemMessage(pl, lang.YuXiWinHadNotDayRew)
		return
	}
	//背包空间
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlots(dayItemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("yuxi:处理玩家玉玺之战每日奖励，背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	//增加物品
	itemGetReason := commonlog.InventoryLogReasonYuXiDayRew
	flag := inventoryManager.BatchAdd(dayItemMap, itemGetReason, itemGetReason.String())
	if !flag {
		panic("yuxi:yuxi rewards add item should be ok")
	}
	//同步资源
	inventorylogic.SnapInventoryChanged(pl)

	if flag := yuXiManager.ReceiveDayRew(); !flag {
		panic("yuxi:领取每日奖励应该成功")
	}

	scMsg := pbutil.BuildSCYuXiReceiveDayRew(dayItemMap)
	pl.SendMsg(scMsg)
	return
}
