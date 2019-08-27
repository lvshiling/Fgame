package cross_handler

import (
	"fgame/fgame/common/codec"
	crosspb "fgame/fgame/common/codec/pb/cross"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	playerarenapvp "fgame/fgame/game/arenapvp/player"
	arenapvptemplate "fgame/fgame/game/arenapvp/template"
	arenapvptypes "fgame/fgame/game/arenapvp/types"
	emaillogic "fgame/fgame/game/email/logic"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_RESULT_BATTLE_TYPE), dispatch.HandlerFunc(handleArenapvpResultBattle))
}

//处理跨服pvp结果
func handleArenapvpResultBattle(s session.Session, msg interface{}) (err error) {
	log.Debug("pvp:处理跨服pvp结果")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isMsg := msg.(*crosspb.ISArenapvpResultBattle)
	win := isMsg.GetWin()
	typeInt := isMsg.GetPvpType()

	pvpType := arenapvptypes.ArenapvpType(typeInt)
	err = arenapvpResultBattle(tpl, win, pvpType)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"err":      err,
			}).Error("pvp:处理跨服pvp结果,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("pvp:处理跨服pvp结果,完成")
	return nil

}

//pvp结果
func arenapvpResultBattle(pl player.Player, win bool, pvpType arenapvptypes.ArenapvpType) (err error) {
	taoTaiTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpTaoTaiTemp(pvpType)
	if taoTaiTemp == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"pvpType":  pvpType,
			}).Warn("pvp:处理跨服pvp结果,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

	reasonGold := commonlog.GoldLogReasonArenapvpRew
	reasonSilver := commonlog.SilverLogReasonArenapvpRew
	reasonLevel := commonlog.LevelLogReasonArenapvpRew
	reasonGoldText := fmt.Sprintf(reasonGold.String(), win, pvpType.String())
	reasonSilverText := fmt.Sprintf(reasonSilver.String(), win, pvpType.String())
	reasonLevelText := fmt.Sprintf(reasonLevel.String(), win, pvpType.String())
	rd := taoTaiTemp.GetRewData(win)
	flag := propertyManager.AddRewData(rd, reasonGold, reasonGoldText, reasonSilver, reasonSilverText, reasonLevel, reasonLevelText)
	if !flag {
		panic(fmt.Errorf("arenapvp:添加资源应该成功"))
	}

	itemMap := taoTaiTemp.GetRewItemMap(win)
	if len(itemMap) > 0 {
		if inventoryManager.HasEnoughSlots(itemMap) {
			reason := commonlog.InventoryLogReasonArenapvpRew
			reasonText := fmt.Sprintf(reason.String(), win, pvpType.String())
			flag := inventoryManager.BatchAdd(itemMap, reason, reasonText)
			if !flag {
				panic(fmt.Errorf("arenapvp:添加物品应该成功"))
			}

		} else {
			title := lang.GetLangService().ReadLang(lang.ArenapvpBattleFinishMailTitle)
			content := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
			emaillogic.AddEmail(pl, title, content, itemMap)
		}
	}
	propertylogic.SnapChangedProperty(pl)
	inventorylogic.SnapInventoryChanged(pl)

	addJiFen := taoTaiTemp.GetJiFen(win)
	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpManager.ArenapvpResult(win, pvpType, addJiFen)
	return
}
