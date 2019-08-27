package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/alliance/alliance"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
	inventorylogic "fgame/fgame/game/inventory/logic"
	playerinventory "fgame/fgame/game/inventory/player"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	propertylogic "fgame/fgame/game/property/logic"
	playerproperty "fgame/fgame/game/property/player"
	gamesession "fgame/fgame/game/session"
	"fgame/fgame/game/shenmo/pbutil"
	playershenmo "fgame/fgame/game/shenmo/player"
	"fgame/fgame/game/shenmo/shenmo"
	shenmotemplate "fgame/fgame/game/shenmo/template"
	shenmotypes "fgame/fgame/game/shenmo/types"
	"fgame/fgame/pkg/timeutils"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_SHENMO_GET_REWARD_TYPE), dispatch.HandlerFunc(handleShenMoGetReward))
}

//处理获取上次领取奖励的排行榜时间戳
func handleShenMoGetReward(s session.Session, msg interface{}) (err error) {
	log.Debug("shenmo:处理获取上次领取奖励的排行榜时间戳")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = shenMoGetReward(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("shenmo:处理获取上次领取奖励的排行榜时间戳,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("shenmo:处理获取上次领取奖励的排行榜时间戳")
	return nil
}

//处理获取上次领取奖励的排行榜时间戳
func shenMoGetReward(pl player.Player) (err error) {
	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("shenmo:您当前还未加入仙盟,无法领取奖励")
		playerlogic.SendSystemMessage(pl, lang.ShenMoGetMyRankNoAlliance)
		return
	}
	rankType := shenmotypes.RankTimeTypeLast
	serverId := global.GetGame().GetServerIndex()
	pos, rankTime := shenmo.GetShenMoService().GetMyRank(rankType, serverId, allianceId)
	if pos == 0 {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("shenmo:您的仙盟排名未在榜单内,无法领取奖励")
		playerlogic.SendSystemMessage(pl, lang.ShenMoGetRewardNoInRank)
		return
	}

	//判断玩家是否领取过奖励
	manager := pl.GetPlayerDataManager(types.PlayerShenMoWarDataManagerType).(*playershenmo.PlayerShenMoDataManager)
	if manager.IsHasedReward(rankTime) {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("shenmo:您已经领过周排行奖励了")
		playerlogic.SendSystemMessage(pl, lang.ShenMoGetRewardHasedGet)
		return
	}

	shenMoRankTemplate := shenmotemplate.GetShenMoTemplateService().GetShenMoRankTemplateByRankNum(pos)
	if shenMoRankTemplate == nil {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("shenmo:您的仙盟排名未在榜单内,无法领取奖励")
		playerlogic.SendSystemMessage(pl, lang.ShenMoGetRewardNoInRank)
		return
	}
	hasedJoinTime := alliance.GetAllianceService().GetAllianceMemberHasedJionTime(pl.GetId())
	if hasedJoinTime < int64(3*timeutils.DAY) {
		log.WithFields(log.Fields{
			"playerid": pl.GetId(),
		}).Warn("shenmo:您的仙盟排名未在榜单内,无法领取奖励")
		playerlogic.SendSystemMessage(pl, lang.ShenMoGetRewardNoEnoughTime)
		return
	}

	propertyManager := pl.GetPlayerDataManager(types.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	silver := int64(shenMoRankTemplate.GetSilver)
	if silver != 0 {
		reason := commonlog.SilverLogReasonShenMoWarRankReward
		reasonText := fmt.Sprintf(reason.String(), pos)
		propertyManager.AddSilver(silver, reason, reasonText)
	}
	bindGold := int64(shenMoRankTemplate.GetBindGold)
	if bindGold != 0 {
		reason := commonlog.GoldLogReasonShenMoWarRankReward
		reasonText := fmt.Sprintf(reason.String(), pos)
		propertyManager.AddGold(bindGold, true, reason, reasonText)
	}
	propertylogic.SnapChangedProperty(pl)

	rewItemMap := shenMoRankTemplate.GetRewItemMap()
	if len(rewItemMap) != 0 {
		inventoryManager := pl.GetPlayerDataManager(types.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if inventoryManager.HasEnoughSlots(rewItemMap) {
			reason := commonlog.InventoryLogReasonShenMoWar
			reasonText := fmt.Sprintf(reason.String(), pos)
			if !inventoryManager.BatchAdd(rewItemMap, reason, reasonText) {
				panic(fmt.Errorf("shenmo:处理领取周排行奖励应该是ok的"))
			}
		} else {
			title := lang.GetLangService().ReadLang(lang.ShenMoGetRewardTitle)
			content := lang.GetLangService().ReadLang(lang.ShenMoGetRewardContent)
			emaillogic.AddEmail(pl, title, content, rewItemMap)
		}
	}
	inventorylogic.SnapInventoryChanged(pl)

	addGongXunNum := shenMoRankTemplate.GetGongXun
	flag := manager.RewardGet(rankTime, addGongXunNum)
	if !flag {
		panic(fmt.Errorf("shenmo:RewardGet应该是ok的"))
	}

	scShenMoGetReward := pbutil.BuildSCShenMoGetReward(rankTime)
	pl.SendMsg(scShenMoGetReward)
	return
}
