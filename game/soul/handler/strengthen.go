package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	soullogic "fgame/fgame/game/soul/logic"
	"fgame/fgame/game/soul/pbutil"
	playersoul "fgame/fgame/game/soul/player"
	"fgame/fgame/game/soul/soul"
	soultypes "fgame/fgame/game/soul/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {

	processor.Register(codec.MessageType(uipb.MessageType_CS_SOUL_STRENGTHEN_TYPE), dispatch.HandlerFunc(handleSoulStrengthen))

}

//处理强化信息
func handleSoulStrengthen(s session.Session, msg interface{}) (err error) {
	log.Debug("soul:处理强化信息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csSoulStrengthen := msg.(*uipb.CSSoulStrengthen)
	soulTag := csSoulStrengthen.GetSoulTag()

	err = soulStrengthen(tpl, soultypes.SoulType(soulTag))
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"soulTag":  soulTag,
				"error":    err,
			}).Error("soul:处理强化信息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("soul:处理强化信息完成")
	return nil

}

//强化逻辑
func soulStrengthen(pl player.Player, soulTag soultypes.SoulType) (err error) {
	soulManager := pl.GetPlayerDataManager(types.PlayerSoulDataManagerType).(*playersoul.PlayerSoulDataManager)
	flag := soulTag.Valid()
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	flag = soulManager.IfSoulTagExist(soulTag)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:未激活的帝魂,无法强化")
		playerlogic.SendSystemMessage(pl, lang.SoulNotActiveNotStrengthen)
		return
	}

	flag = soulManager.IfCanStrengthen(soulTag)
	if !flag {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:强化阶别已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.SoulStrengthenReackLimit)
		return
	}

	soulInfo := soulManager.GetSoulInfoByTag(soulTag)
	nextLevel := soulInfo.StrengthenLevel + 1
	strengthenTemplate := soul.GetSoulService().GetSoulStrengthenTemplateByLevel(soulTag, nextLevel)
	if strengthenTemplate == nil {
		log.WithFields(log.Fields{
			"playerId": pl.GetId(),
			"soulTag":  soulTag,
		}).Warn("soul:强化阶别已达最高阶")
		playerlogic.SendSystemMessage(pl, lang.SoulStrengthenReackLimit)
		return
	}

	//升级需要消耗的银两
	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	costSilver := int64(strengthenTemplate.UseSilver)
	//是否足够银两
	if costSilver != 0 {
		flag := propertyManager.HasEnoughSilver(int64(costSilver))
		if !flag {
			log.WithFields(log.Fields{
				"playerid": pl.GetId(),
			}).Warn("soul:银两不足,无法强化")
			playerlogic.SendSystemMessage(pl, lang.PlayerSilverNoEnough)
			return
		}

		//消耗银两
		reasonSliverText := commonlog.SilverLogReasonSoulStrengthen.String()
		flag = propertyManager.CostSilver(costSilver, commonlog.SilverLogReasonSoulStrengthen, reasonSliverText)
		if !flag {
			panic(fmt.Errorf("soul: soulStrengthen CostSilver should be ok"))
		}
		propertylogic.SnapChangedProperty(pl)
	}
	//强化判断
	pro, _, sucess := soullogic.SoulStrengthen(soulInfo.StrengthenNum, soulInfo.StrengthenPro, strengthenTemplate)
	soulManager.SoulStrengthen(soulTag, pro, sucess)
	//同步属性
	if sucess {
		soullogic.SoulPropertyChanged(pl)
	}
	scSoulStrengthen := pbutil.BuildSCSoulStrengthen(int32(soulTag), soulInfo.StrengthenLevel, soulInfo.StrengthenPro)
	pl.SendMsg(scSoulStrengthen)
	return
}
