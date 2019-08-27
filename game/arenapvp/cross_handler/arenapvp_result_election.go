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
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.RegisterCross(codec.MessageType(crosspb.MessageType_IS_ARENAPVP_RESULT_ELECTION_TYPE), dispatch.HandlerFunc(handleArenapvpResultElection))
}

//处理跨服pvp结果
func handleArenapvpResultElection(s session.Session, msg interface{}) (err error) {
	log.Debug("pvp:处理跨服pvp结果")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	isMsg := msg.(*crosspb.ISArenapvpResultElection)
	win := isMsg.GetWin()
	ranking := isMsg.GetRanking()
	typeInt := isMsg.GetPvpType()

	pvpType := arenapvptypes.ArenapvpType(typeInt)
	err = arenapvpResultElection(tpl, win, ranking, pvpType)
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
func arenapvpResultElection(pl player.Player, win bool, ranking int32, pvpType arenapvptypes.ArenapvpType) (err error) {
	rankTemp := arenapvptemplate.GetArenapvpTemplateService().GetArenapvpRankTemplateByRankNum(ranking)
	addJiFen := int32(0)
	if rankTemp != nil {
		addJiFen = rankTemp.RewJifen
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)

		reasonGold := commonlog.GoldLogReasonArenapvpRew
		reasonSilver := commonlog.SilverLogReasonArenapvpRew
		reasonGoldText := fmt.Sprintf(reasonGold.String(), ranking, pvpType.String())
		reasonSilverText := fmt.Sprintf(reasonSilver.String(), ranking, pvpType.String())
		flag := propertyManager.AddMoney(int64(rankTemp.GetBindGold), int64(rankTemp.GetGold), reasonGold, reasonGoldText, int64(rankTemp.GetSilver), reasonSilver, reasonSilverText)
		if !flag {
			panic(fmt.Errorf("arenapvp:添加资源应该成功"))
		}

		itemMap := rankTemp.GetRewItemMap()
		if len(itemMap) > 0 {
			if inventoryManager.HasEnoughSlots(itemMap) {
				reason := commonlog.InventoryLogReasonArenapvpRew
				reasonText := fmt.Sprintf(reason.String(), ranking, pvpType.String())
				flag := inventoryManager.BatchAdd(itemMap, reason, reasonText)
				if !flag {
					panic(fmt.Errorf("arenapvp:添加物品应该成功"))
				}

			} else {
				title := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
				content := lang.GetLangService().ReadLang(lang.ArenapvpElectionWinMailContent)
				emaillogic.AddEmail(pl, title, content, itemMap)
			}
		}

		inventorylogic.SnapInventoryChanged(pl)
		propertylogic.SnapChangedProperty(pl)

		if !win {
			//失败发邮件
			title := lang.GetLangService().ReadLang(lang.ArenapvpElectionFailedMailTitle)
			content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.ArenapvpElectionFailedMailContent), addJiFen)
			emaillogic.AddEmail(pl, title, content, nil)
		}
	}

	arenapvpManager := pl.GetPlayerDataManager(playertypes.PlayerArenapvpDataManagerType).(*playerarenapvp.PlayerArenapvpDataManager)
	arenapvpManager.ArenapvpResult(win, pvpType, addJiFen)
	return
}
