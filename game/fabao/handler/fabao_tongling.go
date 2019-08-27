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

	processor.Register(codec.MessageType(uipb.MessageType_CS_FABAO_TONGLING_TYPE), dispatch.HandlerFunc(handleFaBaoTongLing))

}

//处理法宝通灵信息
func handleFaBaoTongLing(s session.Session, msg interface{}) (err error) {
	log.Debug("fabao:处理法宝通灵信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = faBaoTongLing(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("fabao:处理法宝通灵信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("fabao:处理法宝通灵信息完成")
	return nil

}

//法宝通灵的逻辑
func faBaoTongLing(pl player.Player) (err error) {
	manager := pl.GetPlayerDataManager(types.PlayerFaBaoDataManagerType).(*playerfabao.PlayerFaBaoDataManager)
	faBaoInfo := manager.GetFaBaoInfo()
	advancedId := faBaoInfo.GetAdvancedId()
	tongLingLevel := faBaoInfo.GetTongLingLevel()
	fabaoTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoNumber(int32(advancedId))
	if fabaoTemplate == nil {
		return
	}

	tongLingTemplate := fabaotemplate.GetFaBaoTemplateService().GetFaBaoTongLingTemplate(tongLingLevel + 1)
	if tongLingTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
		}).Warn("fabao:通灵等级满级")
		playerlogic.SendSystemMessage(pl, lang.FaBaoTongLingReachedFull)
		return
	}

	useItem := tongLingTemplate.UseItem
	num := tongLingTemplate.ItemCount
	if useItem != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		curNum := inventoryManager.NumOfItems(useItem)
		if curNum < num {
			log.WithFields(log.Fields{
				"playerId": pl.GetId(),
				"num":      num,
			}).Warn("fabao:当前通灵数量不足")
			playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
			return
		}
		//消耗物品
		reasonText := commonlog.InventoryLogReasonFaBaoTongLing.String()
		flag := inventoryManager.UseItem(useItem, num, commonlog.InventoryLogReasonFaBaoTongLing, reasonText)
		if !flag {
			panic(fmt.Errorf("fabao:UseItem should be ok"))
		}
		//同步物品
		inventorylogic.SnapInventoryChanged(pl)
	}

	//通灵判断
	pro, _, sucess := fabaologic.FaBaoTongLing(faBaoInfo.GetTongLingNum(), faBaoInfo.GetTongLingPro(), tongLingTemplate)
	flag := manager.TongLing(pro, sucess)
	if !flag {
		panic(fmt.Errorf("fabao: faBaoTongLing should be ok"))
	}
	if sucess {
		fabaologic.FaBaoPropertyChanged(pl)
	}

	scFaBaoTongLing := pbutil.BuildSCFaBaoTongLing(faBaoInfo.GetTongLingLevel(), faBaoInfo.GetTongLingNum(), faBaoInfo.GetTongLingPro())
	pl.SendMsg(scFaBaoTongLing)
	return
}
