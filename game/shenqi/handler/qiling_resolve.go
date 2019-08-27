package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	coreutils "fgame/fgame/core/utils"
	funcopentypes "fgame/fgame/game/funcopen/types"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"
	shenqilogic "fgame/fgame/game/shenqi/logic"
	"fgame/fgame/game/shenqi/pbutil"
	playershenqi "fgame/fgame/game/shenqi/player"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENQI_QILING_RESOLVE_TYPE), dispatch.HandlerFunc(handleShenQiQiLingResolve))
}

//处理分解神器器灵
func handleShenQiQiLingResolve(s session.Session, msg interface{}) (err error) {
	log.Debug("qilingresolve:处理神器器灵分解")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)
	csResolveEquip := msg.(*uipb.CSShenqiQilingResolve)
	itemIndexList := csResolveEquip.GetIndexList()

	err = shenQiQiLingResolve(tpl, itemIndexList)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"ItemIndex": itemIndexList,
				"error":     err,
			}).Error("qilingresolve:处理神器器灵分解,错误")

		return err
	}
	log.Debug("qilingresolve:处理神器器灵分解,完成")
	return nil
}

//分解
func shenQiQiLingResolve(pl player.Player, itemIndexList []int32) (err error) {
	if !pl.IsFuncOpen(funcopentypes.FuncOpenTypeShenQiResolve) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("qilingresolve:处理分解神器器灵,功能未开启")
		playerlogic.SendSystemMessage(pl, lang.CommonFuncNoOpen)
		return
	}

	if len(itemIndexList) < 1 {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("qilingresolve:处理分解神器器灵,没有装备")
		playerlogic.SendSystemMessage(pl, lang.ShenQiResolveNotQiLing)
		return
	}

	if coreutils.IfRepeatElementInt32(itemIndexList) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"indexList": itemIndexList,
			}).Warn("qilingresolve:处理分解神器器灵,索引重复")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	// 计算分解物品
	lingQiVal, flag := shenqilogic.CountResolveQiLing(pl, itemIndexList)
	if !flag {
		return
	}

	inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
	//消耗装备
	useReason := commonlog.InventoryLogReasonResolveCost
	flag, err = inventoryManager.BatchRemoveIndex(inventorytypes.BagTypeQiLing, itemIndexList, useReason, useReason.String())
	if err != nil {
		return
	}
	if !flag {
		panic(fmt.Errorf("qilingresolve:消耗物品应该成功"))
	}
	shenQiManager := pl.GetPlayerDataManager(playertypes.PlayerShenQiDataManagerType).(*playershenqi.PlayerShenQiDataManager)
	flag = shenQiManager.AddLingQiNum(lingQiVal)
	if !flag {
		panic(fmt.Errorf("qilingresolve:增加灵气值应该成功"))
	}

	inventorylogic.SnapInventoryChanged(pl)
	//发送事件
	scMsg := pbutil.BuildSCShenQiQilingResolve(shenQiManager.GetShenQiOjb().LingQiNum)
	pl.SendMsg(scMsg)
	return
}
