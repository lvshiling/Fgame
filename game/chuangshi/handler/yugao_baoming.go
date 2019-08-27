package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/chuangshi/chuangshi"
	"fgame/fgame/game/chuangshi/pbutil"
	playerchuangshi "fgame/fgame/game/chuangshi/player"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_CHUANGSHI_YUGAO_BAOMING_TYPE), dispatch.HandlerFunc(handleChuangShiBaoMing))

}

//处理玩家报名参加创世之战消息
func handleChuangShiBaoMing(s session.Session, msg interface{}) (err error) {
	log.Debug("chuangshi:处理玩家报名参加创世之战消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	err = chuangShiBaoMing(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("chuangshi:处理玩家报名参加创世之战消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("chuangshi:处理玩家报名参加创世之战消息,成功")
	return nil
}

func chuangShiBaoMing(pl player.Player) (err error) {
	playerChuangShiManager := pl.GetPlayerDataManager(playertypes.PlayerChuangShiDataManagerType).(*playerchuangshi.PlayerChuangShiDataManager)
	if playerChuangShiManager.IsJoin() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Warn("chuangshi:玩家已经报名")
		playerlogic.SendSystemMessage(pl, lang.ChuangShiAlreadyBaoMing)
		return
	}

	// 检查创世之战是否可以报名
	flag, err := chuangshi.GetChuangShiService().CheckYuGaoTime()
	if err != nil {
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Warn("chuangshi:不在传世之战报名时间段内")
		playerlogic.SendSystemMessage(pl, lang.ChuangShiOutOfTime)
		return
	}
	yuGaoTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiYuGaoTemplate()
	itemMap := yuGaoTemp.GetReceiveItemMap()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	if !inventoryManager.HasEnoughSlots(itemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("chuangshi:背包空间不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	// 报名成功，添加物品
	if len(itemMap) != 0 {
		itemGetReason := commonlog.InventoryLogReasonBaoMingChuangShiGet
		itemGetReasonText := itemGetReason.String()
		flag := inventoryManager.BatchAdd(itemMap, itemGetReason, itemGetReasonText)
		if !flag {
			panic(fmt.Errorf("chuangshi: 报名创世之战添加物品应该成功"))
		}
	}

	//物品变化，同步背包数据
	inventorylogic.SnapInventoryChanged(pl)

	flag = playerChuangShiManager.BaoMingSucess()
	if !flag {
		panic("chuangshi: 创世之战报名应该成功")
	}

	chuangshi.GetChuangShiService().AddPlayerNum()

	scMsg := pbutil.BuildSCBaoMingChuangShi(itemMap)
	pl.SendMsg(scMsg)
	return
}
