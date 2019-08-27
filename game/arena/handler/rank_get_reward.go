package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/session"
	"fgame/fgame/game/arena/arena"
	"fgame/fgame/game/arena/pbutil"
	playerarena "fgame/fgame/game/arena/player"
	arenatemplate "fgame/fgame/game/arena/template"
	arenatypes "fgame/fgame/game/arena/types"
	emaillogic "fgame/fgame/game/email/logic"
	"fgame/fgame/game/global"
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
	processor.Register(codec.MessageType(uipb.MessageType_CS_ARENA_GET_REWARD_TYPE), dispatch.HandlerFunc(handleArenaGetReward))
}

//处理获取上次领取奖励的排行榜时间戳
func handleArenaGetReward(s session.Session, msg interface{}) (err error) {
	log.Debug("arena:处理获取上次领取奖励的排行榜时间戳")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	err = arenaGetReward(tpl)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Error("arena:处理获取上次领取奖励的排行榜时间戳,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("arena:处理获取上次领取奖励的排行榜时间戳")
	return nil
}

//处理获取上次领取奖励的排行榜时间戳
func arenaGetReward(pl player.Player) (err error) {
	timeType := arenatypes.RankTimeTypeLast
	serverId := global.GetGame().GetServerIndex()
	pos, _, rankTime := arena.GetArenaService().GetMyRank(timeType, serverId, pl.GetId())
	if pos == 0 {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
			}).Warn("arena:您的连胜排名未在榜单内,无法领取奖励")
		playerlogic.SendSystemMessage(pl, lang.ArenaGetRewardNoInRank)
		return
	}

	arenaRankTemplate := arenatemplate.GetArenaTemplateService().GetArenaRankTemplateByRankNum(pos)
	if arenaRankTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"pos":      pos,
			}).Warn("arena:领取周榜奖励失败,模板不存在")
		playerlogic.SendSystemMessage(pl, lang.CommonTemplateNotExist)
		return
	}

	//判断玩家是否领取过奖励
	arenaManager := pl.GetPlayerDataManager(playertypes.PlayerArenaDataManagerType).(*playerarena.PlayerArenaDataManager)
	if arenaManager.IsHasedReward(rankTime) {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"rankTime": rankTime,
			}).Warn("arena:您已经领过周排行奖励了")
		playerlogic.SendSystemMessage(pl, lang.ArenaGetRewardHasedGet)
		return
	}

	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	silver := int64(arenaRankTemplate.GetSilver)
	if silver != 0 {
		reason := commonlog.SilverLogReasonArenaRankReward
		reasonText := fmt.Sprintf(reason.String(), pos)
		propertyManager.AddSilver(silver, reason, reasonText)
	}

	bindGold := int64(arenaRankTemplate.GetBindGold)
	if bindGold != 0 {
		reason := commonlog.GoldLogReasonArenaRankReward
		reasonText := fmt.Sprintf(reason.String(), pos)
		propertyManager.AddGold(bindGold, true, reason, reasonText)
	}

	gold := int64(arenaRankTemplate.GetGold)
	if gold != 0 {
		reason := commonlog.GoldLogReasonArenaRankReward
		reasonText := fmt.Sprintf(reason.String(), pos)
		propertyManager.AddGold(gold, false, reason, reasonText)
	}
	propertylogic.SnapChangedProperty(pl)

	rewItemMap := arenaRankTemplate.GetRewItemMap()
	if len(rewItemMap) != 0 {
		inventoryManager := pl.GetPlayerDataManager(playertypes.PlayerInventoryDataManagerType).(*playerinventory.PlayerInventoryDataManager)
		if inventoryManager.HasEnoughSlots(rewItemMap) {
			reason := commonlog.InventoryLogReasonArenaRankReward
			reasonText := fmt.Sprintf(reason.String(), pos)
			if !inventoryManager.BatchAdd(rewItemMap, reason, reasonText) {
				panic(fmt.Errorf("arena:处理领取周排行奖励应该是ok的"))
			}
		} else {
			title := lang.GetLangService().ReadLang(lang.ArenaMailTitle)
			content := lang.GetLangService().ReadLang(lang.EmailInventorySlotNoEnough)
			emaillogic.AddEmail(pl, title, content, rewItemMap)
		}
	}
	inventorylogic.SnapInventoryChanged(pl)

	flag := arenaManager.RankRewardGet(rankTime)
	if !flag {
		panic(fmt.Errorf("arena:RewardGet应该是ok的"))
	}

	scMsg := pbutil.BuildSCArenaGetReward(rankTime)
	pl.SendMsg(scMsg)
	return
}
