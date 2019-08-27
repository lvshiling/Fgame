package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_RANK_SHIELD_TYPE), dispatch.HandlerFunc(handleRankShieldList))
}

//处理神盾尖刺排行榜信息
func handleRankShieldList(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取神盾尖刺排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csRankShieldGet := msg.(*uipb.CSOpenActivityRankShieldList)
	page := csRankShieldGet.GetPage()
	groupId := csRankShieldGet.GetGroupId()

	err = rankShieldList(tpl, page, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"page":     page,
				"error":    err,
			}).Error("welfare:处理获取神盾尖刺排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理获取神盾尖刺排行榜消息完成")
	return nil

}

//获取神盾尖刺排行榜界面信息的逻辑
func rankShieldList(pl player.Player, page int32, groupId int32) (err error) {
	if page < 0 {
		log.WithFields(
			log.Fields{
				"playerid": pl.GetId(),
				"page":     page,
			}).Warn("welfare:参数无效")
		playerlogic.SendSystemMessage(pl, lang.CommonArgumentInvalid)
		return
	}

	//检验活动
	typ := welfaretypes.OpenActivityTypeRank
	subType := welfaretypes.OpenActivityRankSubTypeShield
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	shieldList, rankTime := rank.GetRankService().GetOrderListByPage(ranktypes.RankTypeShield, ranktypes.RankClassTypeLocalActivity, groupId, page)
	if page == 0 {
		nextPageList, _ := rank.GetRankService().GetOrderListByPage(ranktypes.RankTypeShield, ranktypes.RankClassTypeLocalActivity, groupId, page+1)
		if len(nextPageList) > 0 {
			shieldList = append(shieldList, nextPageList[0])
		}
	}

	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scRankShieldGet := pbutil.BuildSCOpenActivityRankShieldList(page, shieldList, rankTime, startTime, endTime)
	pl.SendMsg(scRankShieldGet)
	return
}
