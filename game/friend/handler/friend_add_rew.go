package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/friend/pbutil"
	playerfriend "fgame/fgame/game/friend/player"
	friendtemplate "fgame/fgame/game/friend/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_FRIEND_ADD_REW_TYPE), dispatch.HandlerFunc(handleFriendAddRew))
}

//处理领取添加好友奖励
func handleFriendAddRew(s session.Session, msg interface{}) (err error) {
	log.Debug("friend:处理领取添加好友奖励")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csMsg := msg.(*uipb.CSFriendAddRew)
	frNum := csMsg.GetFrNum()

	err = friendAddRew(tpl, frNum)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("friend:处理领取添加好友奖励,错误")
		return err
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("friend:处理领取添加好友奖励,完成")
	return nil
}

//处理领取添加好友奖励
func friendAddRew(pl player.Player, frNum int32) (err error) {

	addRewTemp := friendtemplate.GetFriendNoticeTemplateService().GetFriendAddTemplate(frNum)
	if addRewTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"frNum":    frNum,
			}).Warn("friend:处理领取添加好友奖励,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 是否领取
	manager := pl.GetPlayerDataManager(playertypes.PlayerFriendDataManagerType).(*playerfriend.PlayerFriendDataManager)
	if !manager.IsCanReceiveRew(frNum) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"frNum":    frNum,
			}).Warn("friend:处理领取添加好友奖励,已领取")
		playerlogic.SendSystemMessage(pl, lang.FriendAddRewHadDone)
		return
	}

	// 是否够好友数量
	curFriNum := int32(len(manager.GetFriends()))
	dummyNum := manager.GetDummyFriendNum()
	if curFriNum+dummyNum < frNum {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"dummyNum": dummyNum,
				"frNum":    frNum,
			}).Warn("friend:处理领取添加好友奖励,好友数量不足")
		playerlogic.SendSystemMessage(pl, lang.FriendNumNotEnough)
		return
	}

	// 背包是否足够
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	rewItemMap := addRewTemp.GetRewItemMap()
	if !inventoryManager.HasEnoughSlots(rewItemMap) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"frNum":    frNum,
			}).Warn("friend:处理领取添加好友奖励,好友数量不足")
		playerlogic.SendSystemMessage(pl, lang.InventorySlotNoEnough)
		return
	}

	// 添加奖励
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	addLevelReason := commonlog.LevelLogReasonFriendAddRew
	if addRewTemp.RewardExp > 0 {
		propertyManager.AddExp(int64(addRewTemp.RewardExp), addLevelReason, addLevelReason.String())
	}
	if addRewTemp.RewardExpPoint > 0 {
		propertyManager.AddExpPoint(int64(addRewTemp.RewardExpPoint), addLevelReason, addLevelReason.String())
	}

	addSilverReason := commonlog.SilverLogReasonAddFriendRew
	if addRewTemp.RewardSilver > 0 {
		propertyManager.AddSilver(int64(addRewTemp.RewardSilver), addSilverReason, addSilverReason.String())
	}

	addItemReason := commonlog.InventoryLogReasonAddFriendRew
	if len(rewItemMap) > 0 {
		flag := inventoryManager.BatchAdd(rewItemMap, addItemReason, addItemReason.String())
		if !flag {
			panic("添加好友奖励，批量添加物品应该成功")
		}
	}

	manager.AddRewRecord(frNum)

	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	scMsg := pbutil.BuildSCFriendAddRew(frNum)
	pl.SendMsg(scMsg)
	return
}
