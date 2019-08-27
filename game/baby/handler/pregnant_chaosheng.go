package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/baby/pbutil"
	playerbaby "fgame/fgame/game/baby/player"
	babytemplate "fgame/fgame/game/baby/template"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	playervip "fgame/fgame/game/vip/player"
	viptemplate "fgame/fgame/game/vip/template"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_BABY_BORN_CHAOSHENG_TYPE), dispatch.HandlerFunc(handleBabyBornChaoSheng))
}

//处理超生
func handleBabyBornChaoSheng(s session.Session, msg interface{}) (err error) {
	log.Debug("baby:处理超生消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = babyBornChaoSheng(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"error":    err,
			}).Error("baby:处理超生消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("baby:处理超生消息完成")
	return nil

}

//超生界面逻辑
func babyBornChaoSheng(pl player.Player) (err error) {

	vipManager := pl.GetPlayerDataManager(playertypes.PlayerVipDataManagerType).(*playervip.PlayerVipDataManager)
	babyManager := pl.GetPlayerDataManager(playertypes.PlayerBabyDataManagerType).(*playerbaby.PlayerBabyDataManager)
	pregnantInfo := babyManager.GetPregnantInfo()

	level, star := vipManager.GetVipLevel()
	vipTemplate := viptemplate.GetVipTemplateService().GetVipTemplate(level, star)
	if vipTemplate.BabyChaoSheng <= pregnantInfo.GetChaoShengNum() {
		log.WithFields(
			log.Fields{
				"playerId":     pl.GetId(),
				"VIP":          level,
				"canChaoSheng": vipTemplate.BabyChaoSheng,
				"hadChaoSheng": pregnantInfo.GetChaoShengNum(),
			}).Warn("baby:处理加速出生消息, 超生数量已满")
		playerlogic.SendSystemMessage(pl, lang.BabyFullChaoSheng)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	needGold := int64(babyConstantTemplate.GetChaoShengGold(int(pregnantInfo.GetChaoShengNum())))
	if !propertyManager.HasEnoughGold(needGold, false) {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"needGold": needGold,
			}).Warn("baby:处理加速出生消息, 元宝不足")
		playerlogic.SendSystemMessage(pl, lang.PlayerGoldNoEnough)
		return
	}

	goldCostReason := commonlog.GoldLogReasonBabyChaoShengUse
	flag := propertyManager.CostGold(needGold, false, goldCostReason, goldCostReason.String())
	if !flag {
		panic(fmt.Errorf("baby: 消耗元宝应该成功"))
	}

	//超生
	propertylogic.SnapChangedProperty(pl)
	babyManager.ChaoSheng()

	scMsg := pbutil.BuildSCBabyBornChaoSheng()
	pl.SendMsg(scMsg)
	return
}
