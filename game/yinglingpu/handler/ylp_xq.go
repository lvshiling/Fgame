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
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	ylplogic "fgame/fgame/game/yinglingpu/logic"
	ylppbutil "fgame/fgame/game/yinglingpu/pbutil"
	ylpplayer "fgame/fgame/game/yinglingpu/player"
	ylptmp "fgame/fgame/game/yinglingpu/template"
	ylptypes "fgame/fgame/game/yinglingpu/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_YLPU_SP_XIANGQIAN_TYPE), dispatch.HandlerFunc(handleYlpXq))
}

func handleYlpXq(s session.Session, msg interface{}) (err error) {
	log.Debug("yinglingpu:英灵普镶开始....")
	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csXiangQian := msg.(*uipb.CSYingLingPuSpXiangQian)
	tujianId := csXiangQian.GetTuJianId()
	suipianId := csXiangQian.GetSuiPianId()
	tujianTypeId := csXiangQian.GetTuJianType()
	tujianType := ylptypes.YingLingPuTuJianType(tujianTypeId)

	err = xiangQian(tpl, tujianId, tujianType, suipianId)
	if err != nil {
		log.WithFields(log.Fields{
			"playerId": tpl.GetId(),
			"error":    err,
		}).Error("yinglingpu:英灵普镶嵌失败")
		//需要返回
		return
	}
	log.Debug("yinglingpu:英灵普镶嵌结束....")
	return
}

func xiangQian(pl player.Player, tujianId int32, tujianType ylptypes.YingLingPuTuJianType, suipianId int32) (err error) {
	if !tujianType.Valid() {
		log.WithFields(log.Fields{
			"tujianId":   tujianId,
			"tujianType": tujianType,
			"suipianId":  suipianId,
			"playerId":   pl.GetId(),
		}).Warn("yinglingpu:镶嵌图鉴失败，图鉴类型错误")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}
	templateService := ylptmp.GetYingLingPuTemplateService()
	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	ylpManager := pl.GetPlayerDataManager(playertypes.PlayerYingLingPuManagerType).(*ylpplayer.PlayerYingLingPuManager)

	//碎片已经镶嵌
	exFlag := ylpManager.ExistsYlpSp(tujianId, tujianType, suipianId)
	if exFlag {
		log.WithFields(log.Fields{
			"tujianId":   tujianId,
			"tujianType": tujianType,
			"playerId":   pl.GetId(),
			"suipianId":  suipianId,
		}).Warn("yinglingpu:镶嵌图鉴失败，已经存在")
		playerlogic.SendSystemMessage(pl, lang.YingLingPuSpExists)
		return
	}

	suiPianTemplate := templateService.GetYingLingPuSuiPian(tujianId, suipianId, tujianType)
	//模板中木有
	if suiPianTemplate == nil {
		log.WithFields(log.Fields{
			"tujianId":   tujianId,
			"tujianType": tujianType,
			"playerId":   pl.GetId(),
			"suipianId":  suipianId,
		}).Warn("yinglingpu:镶嵌图鉴失败，未找到模板")
		playerlogic.SendSystemMessage(pl, lang.YingLingPuSpNotAllow)
		return
	}

	template := templateService.GetYingLingPuById(tujianId, tujianType)
	if template == nil {
		log.WithFields(log.Fields{
			"tujianId":   tujianId,
			"tujianType": tujianType,
			"playerId":   pl.GetId(),
			"suipianId":  suipianId,
		}).Warn("yinglingpu:镶嵌图鉴失败，英灵普未找到")
		playerlogic.SendSystemMessage(pl, lang.YingLingPuSpOther)
		return
	}

	flag := inventoryManager.HasEnoughItem(suiPianTemplate.UseItemId, suiPianTemplate.UseItemCount)
	if !flag {
		log.WithFields(log.Fields{
			"tujianId":   tujianId,
			"tujianType": tujianType,
			"playerId":   pl.GetId(),
			"suipianId":  suipianId,
		}).Warn("yinglingpu:镶嵌图鉴失败,物品不足")
		playerlogic.SendSystemMessage(pl, lang.InventoryItemNoEnough)
		return
	}
	reasonText := fmt.Sprintf(commonlog.InventoryLogReasonYingLingPuXq.String(), tujianType.String(), tujianId, suipianId)
	//背包里面的物品必须使用掉
	flag = inventoryManager.UseItem(suiPianTemplate.UseItemId, suiPianTemplate.UseItemCount, commonlog.InventoryLogReasonYingLingPuXq, reasonText)
	if !flag {
		panic(fmt.Sprintf("英灵普合成成功但是扣除物品失败"))
	}

	//镶嵌英灵普
	flag = ylpManager.AddYingLingPuSuiPian(tujianId, tujianType, suipianId)
	if !flag {
		panic(fmt.Errorf("yinglingpu:镶嵌图鉴类型[%s],id[%d],碎片[%d]应该成功", tujianType.String(), tujianId, suipianId))
	}
	inventorylogic.SnapInventoryChanged(pl)
	ylplogic.YingLingPuPropertyChanged(pl)
	successInfo := ylppbutil.BuildYlpXQ(tujianId, int32(tujianType), suipianId)
	pl.SendMsg(successInfo)
	return
}
