package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	fabaologic "fgame/fgame/game/fabao/logic"
	"fgame/fgame/game/fabao/pbutil"
	playerfabao "fgame/fgame/game/fabao/player"
	fabaotemplate "fgame/fgame/game/fabao/template"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_FABAO_UNREAL_TYPE), dispatch.HandlerFunc(handleFaBaoUnreal))

}

//处理法宝幻化信息
func handleFaBaoUnreal(s session.Session, msg interface{}) (err error) {
	log.Debug("fabao:处理法宝幻化信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csFaBaoUnreal := msg.(*uipb.CSFaBaoUnreal)
	faBaoId := csFaBaoUnreal.GetFaBaoId()
	err = faBaoUnreal(tpl, int(faBaoId))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"faBaoId":  faBaoId,
				"error":    err,
			}).Error("fabao:处理法宝幻化信息,错误")
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"faBaoId":  faBaoId,
		}).Debug("fabao:处理法宝幻化信息完成")
	return nil

}

//法宝幻化的逻辑
func faBaoUnreal(pl player.Player, faBaoId int) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBao(int(faBaoId))
	//校验参数
	if faBaoTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"faBaoId":  faBaoId,
		}).Warn("fabao:幻化advancedId无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	advancedId := manager.GetFaBaoAdvancedId()
	if advancedId <= 0 {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"faBaoId":  faBaoId,
		}).Warn("fabao:请先激活法宝系统")
		playerlogic.SendSystemMessage(pl, lang.FaBaoUnrealActiveSystem)
		return
	}

	//是否已幻化
	flag := manager.IsUnrealed(faBaoId)
	if !flag {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		//获取物品幻化条件3,消耗物品数
		paramItemsMap := faBaoTemplate.GetMagicParamIMap()
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.HasEnoughItems(paramItemsMap)
			if !flag {
				log.WithFields(log.Fields{
					"playerId": pl.GetId(),
					"faBaoId":  faBaoId,
				}).Warn("FaBao:还有幻化条件未达成，无法解锁幻化")
				playerlogic.SendSystemMessage(pl, lang.FaBaoUnrealCondNotReached)
				return
			}
		}

		flag = manager.IsCanUnreal(faBaoId)
		if !flag {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"faBaoId":  faBaoId,
			}).Warn("FaBao:还有幻化条件未达成，无法解锁幻化")
			playerlogic.SendSystemMessage(pl, lang.FaBaoUnrealCondNotReached)
			return
		}

		//使用幻化条件3的物品
		inventoryReason := commonlog.InventoryLogReasonFaBaoUnreal
		reasonText := fmt.Sprintf(inventoryReason.String(), faBaoId)
		if len(paramItemsMap) != 0 {
			flag := inventoryManager.BatchRemove(paramItemsMap, inventoryReason, reasonText)
			if !flag {
				panic(fmt.Errorf("fabao:use item should be ok"))
			}
			//同步物品
			inventorylogic.SnapInventoryChanged(pl)
		}
		manager.AddUnrealInfo(faBaoId)
		fabaologic.FaBaoPropertyChanged(pl)
	}

	flag = manager.Unreal(faBaoId)
	if !flag {
		panic(fmt.Errorf("fabao:幻化应该成功"))
	}
	scFaBaoUnreal := pbutil.BuildSCFaBaoUnreal(int32(faBaoId))
	pl.SendMsg(scFaBaoUnreal)
	return
}
