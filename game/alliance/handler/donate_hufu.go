package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	"fgame/fgame/game/alliance/alliance"
	alliancelogic "fgame/fgame/game/alliance/logic"
	"fgame/fgame/game/alliance/pbutil"
	chatlogic "fgame/fgame/game/chat/logic"
	chattypes "fgame/fgame/game/chat/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	noticelogic "fgame/fgame/game/notice/logic"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_ALLIANCE_DONATE_HUFU_TYPE), dispatch.HandlerFunc(handleAllianceDonateHuFu))
}

//处理仙盟捐献虎符
func handleAllianceDonateHuFu(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理仙盟捐献虎符")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = allianceDonateHuFu(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),

				"error": err,
			}).Error("alliance:处理仙盟捐献虎符,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("alliance:处理仙盟捐献虎符,完成")
	return nil

}

//仙盟捐献虎符
func allianceDonateHuFu(pl player.Player) (err error) {
	//获取模板
	huFuId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceItemHuFu)
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)

	//判断是物品足够
	if !inventoryManager.HasEnoughItem(huFuId, 1) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Warn("alliance:处理仙盟捐献虎符,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}

	//捐献
	memObj, err := alliance.GetAllianceService().DonateHuFu(pl.GetId())
	if err != nil {
		return
	}

	//消耗物品
	allianceId := memObj.GetAlliance().GetAllianceId()
	reason := commonlog.InventoryLogReasonAllianceDonateHuFu
	reasonText := fmt.Sprintf(reason.String(), allianceId)
	flag := inventoryManager.UseItem(huFuId, 1, reason, reasonText)
	if !flag {
		panic(fmt.Errorf("使用物品应该成功"))
	}

	//公告
	playerName := coreutils.FormatColor(chattypes.ColorTypePlayerName, coreutils.FormatNoticeStr(pl.GetName()))
	allianceName := coreutils.FormatColor(chattypes.ColorTypeModuleName, coreutils.FormatNoticeStr(memObj.GetAlliance().GetAllianceName()))
	hufuNum := coreutils.FormatColor(chattypes.ColorTypePower, fmt.Sprintf("%d", 1))
	content := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceDonateHuFuNotice), playerName, allianceName, hufuNum)
	chatlogic.SystemBroadcast(chattypes.MsgTypeText, []byte(content))
	noticelogic.NoticeNumBroadcast([]byte(content), 0, 1)

	// 广播
	alliancelogic.DonateHufu(pl)

	//同步背包
	inventorylogic.SnapInventoryChanged(pl)

	scAllianceDonateHuFu := pbutil.BuildSCAllianceDonateHuFu()
	pl.SendMsg(scAllianceDonateHuFu)
	return
}
