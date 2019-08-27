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
	lingtongdevlogic "fgame/fgame/game/lingtongdev/logic"
	"fgame/fgame/game/lingtongdev/pbutil"
	playerlingtongdev "fgame/fgame/game/lingtongdev/player"
	tonglongdevtemplate "fgame/fgame/game/lingtongdev/template"
	"fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_LINGTONGDEV_TONGLING_TYPE), dispatch.HandlerFunc(handleLingTongDevTongLing))

}

//处理灵童养成类通灵信息
func handleLingTongDevTongLing(s session.Session, msg interface{}) (err error) {
	log.Debug("lingtongdev:处理灵童养成类通灵信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csLingTongDevTongLing := msg.(*uipb.CSLingTongDevTongLing)
	classType := csLingTongDevTongLing.GetClassType()

	err = lingTongDevTongLing(tpl, types.LingTongDevSysType(classType))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
				"error":     err,
			}).Error("lingtongdev:处理灵童养成类通灵信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Debug("lingtongdev:处理灵童养成类通灵信息完成")
	return nil

}

//灵童养成类通灵的逻辑
func lingTongDevTongLing(pl player.Player, classType types.LingTongDevSysType) (err error) {
	if !classType.Vaild() {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("LingTongDev:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	manager := pl.GetPlayerDataManager(playertypes.PlayerLingTongDevDataManagerType).(*playerlingtongdev.PlayerLingTongDevDataManager)
	lingTongInfo := manager.GetLingTongDevInfo(classType)
	if lingTongInfo == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("lingtongdev:请先激活灵童养成类系统")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevActiveSystem, classType.String())
		return
	}
	advancedId := lingTongInfo.GetAdvancedId()
	tongLingLevel := lingTongInfo.GetTongLingLevel()
	lingTongDevTemplate := tonglongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevNumber(classType, advancedId)
	if lingTongDevTemplate == nil {
		return
	}

	lingTongDevTongLingTemplate := tonglongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTongLingTemplate(classType, tongLingLevel+1)
	if lingTongDevTongLingTemplate == nil {
		log.WithFields(log.Fields{
			"playerId":  pl.GetId(),
			"classType": classType,
		}).Warn("lingtongdev:通灵等级满级")
		playerlogic.SendSystemMessage(pl, lang.LingTongDevTongLingReachedFull, classType.String())
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	needItems := lingTongDevTongLingTemplate.GetItemMap()
	if len(needItems) != 0 {
		flag := inventoryManager.HasEnoughItems(needItems)
		if !flag {
			log.WithFields(log.Fields{
				"playerId":  pl.GetId(),
				"classType": classType,
			}).Warn("LingTongDev:道具不足，无法通灵")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}

		//消耗物品
		reasonText := fmt.Sprintf(commonlog.InventoryLogReasonLingTongDevTongLing.String(), classType.String())
		flag = inventoryManager.BatchRemove(needItems, commonlog.InventoryLogReasonLingTongDevTongLing, reasonText)
		if !flag {
			panic(fmt.Errorf("lingtongdev:BatchRemove should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	//通灵判断
	pro, _, sucess := lingtongdevlogic.LingTongDevTongLing(lingTongInfo.GetTongLingNum(), lingTongInfo.GetTongLingPro(), lingTongDevTongLingTemplate)
	flag := manager.TongLing(classType, pro, sucess)
	if !flag {
		panic(fmt.Errorf("lingtongdev: lingTongDevTongLing should be ok"))
	}
	if sucess {
		lingtongdevlogic.LingTongDevPropertyChanged(pl, classType)
	}

	scLingTongDevTongLing := pbutil.BuildSCLingTongDevTongLing(int32(classType), lingTongInfo.GetTongLingLevel(), lingTongInfo.GetTongLingPro())
	pl.SendMsg(scLingTongDevTongLing)
	return
}
