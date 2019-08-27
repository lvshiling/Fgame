package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/cache/dao"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/lingtong/pbutil"
	playerlingtong "fgame/fgame/game/lingtong/player"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONG_RENAME_TYPE), dispatch.HandlerFunc(handleLingTongRename))

}

//处理灵童改名信息
func handleLingTongRename(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtong:处理灵童改名信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongRename := msg.(*uipb.CSLingTongRename)
	lingTongId := csLingTongRename.GetLingTongId()
	lingTongName := csLingTongRename.GetLingTongName()

	err = lingTongRename(tpl, lingTongId, lingTongName)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"lingTongId":   lingTongId,
				"lingTongName": lingTongName,
				"error":        err,
			}).Error("lingtong:处理灵童改名信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":     pl.GetId(),
			"lingTongId":   lingTongId,
			"lingTongName": lingTongName,
		}).Debug("lingtong:处理灵童改名信息完成")
	return nil

}

//灵童改名逻辑
func lingTongRename(pl player.Player, lingTongId int32, lingTongName string) (err error) {
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	if lingTongTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":     pl.GetId(),
			"lingTongId":   lingTongId,
			"lingTongName": lingTongName,
		}).Warn("lingtong:模板为空")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	if len(lingTongName) == 0 {
		log.WithFields(log.Fields{
			"playerId":     pl.GetId(),
			"lingTongId":   lingTongId,
			"lingTongName": lingTongName,
		}).Warn("lingtong:名字不能为空")
		return
	}
	serverId := global.GetGame().GetServerIndex()
	playerCacheEntity, err := dao.GetCacheDao().GetPlayerCacheByName(lingTongName, serverId)
	if err != nil {
		return
	}
	if playerCacheEntity != nil {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"lingTongName": lingTongName,
			}).Warn("lingtong:处理改名请求,名字已经存在")
		playerlogic.SendSystemMessage(pl, lang.NameAlreadyExist)
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDataManagerType).(*playerlingtong.PlayerLingTongDataManager)
	lingTongInfo, flag := manager.GetLingTongInfo(lingTongId)
	if !flag {
		log.WithFields(log.Fields{
			"playerId":     pl.GetId(),
			"lingTongId":   lingTongId,
			"lingTongName": lingTongName,
		}).Warn("lingtong:未激活该灵童")
		playerlogic.SendSystemMessage(pl, lang.LingTongNoActive)
		return
	}

	if lingTongInfo.GetLingTongName() == lingTongName {
		log.WithFields(log.Fields{
			"playerId":     pl.GetId(),
			"lingTongId":   lingTongId,
			"lingTongName": lingTongName,
		}).Warn("lingtong:处理改名请求,名字已经存在")
		playerlogic.SendSystemMessage(pl, lang.NameAlreadyExist)
		return
	}

	useItem := lingTongTemplate.NameItemId
	num := lingTongTemplate.NameItemCount
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId":     pl.GetId(),
				"lingTongId":   lingTongId,
				"lingTongName": lingTongName,
			}).Warn("lingtong:物品不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongRename.String(), lingTongId)
		flag = inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonLingTongRename, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtong:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	flag = manager.Rename(lingTongId, lingTongName)
	if !flag {
		panic(fmt.Errorf("lingtong:lingTongRename should be ok"))
	}
	scLingTongRename := pbutil.BuildSCLingTongRename(lingTongInfo.GetLingTongId(), lingTongInfo.GetLingTongName())
	pl.SendMsg(scLingTongRename)
	return
}
