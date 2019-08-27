package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	goldequiplogic "fgame/fgame/game/goldequip/logic"
	"fgame/fgame/game/goldequip/pbutil"
	playergoldequip "fgame/fgame/game/goldequip/player"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	gamesession "fgame/fgame/game/session"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_GOLD_EQUIP_AUTO_FENJIE_TYPE), dispatch.HandlerFunc(handleGoldEquipAutoEat))
}

//处理金装自动分解
func handleGoldEquipAutoEat(s session.Session, msg interface{}) (err error) {
	log.Debug("alliance:处理金装自动分解")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csMsg := msg.(*uipb.CSGoldEquipAutoFenJie)
	isAuto := csMsg.GetIsAuto()
	qualityInt := csMsg.GetMaxQuality()

	quality := itemtypes.ItemQualityType(qualityInt)
	if !quality.Valid() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isAuto":   isAuto,
			}).Warn("alliance:处理金装自动分解,品质参数错误")
		playerlogic.SendSystemMessage(tpl, lang.CommonArgumentInvalid)
		return
	}
	zhuanShu := csMsg.GetZhuanShu()

	err = goldEquipAutoEat(tpl, isAuto, quality, zhuanShu)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"isAuto":   isAuto,
				"error":    err,
			}).Error("alliance:处理金装自动分解,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
			"isAuto":   isAuto,
		}).Debug("alliance:处理金装自动分解,完成")
	return nil
}

//金装自动分解
func goldEquipAutoEat(pl player.Player, isAuto int32, quality itemtypes.ItemQualityType, zhuanShu int32) (err error) {

	goldequipManager := pl.GetPlayerDataManager(playertypes.PlayerGoldEquipDataManagerType).(*playergoldequip.PlayerGoldEquipDataManager)
	goldequipManager.SetAutoFenJie(isAuto, quality, zhuanShu)

	// 自动分解
	goldequiplogic.AutoFenJieGoldEquip(pl)

	scMsg := pbutil.BuildSCGoldEquipAutoFenJie(isAuto, int32(quality), zhuanShu)
	pl.SendMsg(scMsg)
	return
}
