package handler

import (
	"fgame/fgame/common/codec"
	uipb "fgame/fgame/common/codec/pb/ui"
	"fgame/fgame/common/dispatch"
	"fgame/fgame/common/lang"
	"fgame/fgame/core/session"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/processor"
	"fgame/fgame/game/rank/rank"
	ranktypes "fgame/fgame/game/rank/types"
	gamesession "fgame/fgame/game/session"
	welfarelogic "fgame/fgame/game/welfare/logic"
	"fgame/fgame/game/welfare/pbutil"
	playerwelfare "fgame/fgame/game/welfare/player"
	welfaretypes "fgame/fgame/game/welfare/types"
	"fgame/fgame/game/welfare/welfare"

	log "github.com/Sirupsen/logrus"
)

func init() {
	processor.Register(codec.MessageType(uipb.MessageType_CS_OPEN_ACTIVITY_RANK_CHARM_TYPE), dispatch.HandlerFunc(handleOpenActivityRankCharmList))
}

//处理魅力排行榜信息
func handleOpenActivityRankCharmList(s session.Session, msg interface{}) (err error) {
	log.Debug("welfare:处理获取魅力排行榜消息")

	gcs := gamesession.SessionInContext(s.Context())
	pl := gcs.Player()
	tpl := pl.(player.Player)

	csOpenActivityRankCharmList := msg.(*uipb.CSOpenActivityRankCharmList)
	page := csOpenActivityRankCharmList.GetPage()
	groupId := csOpenActivityRankCharmList.GetGroupId()

	err = rankCharmList(tpl, page, groupId)
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"page":     page,
				"error":    err,
			}).Error("welfare:处理获取魅力排行榜消息,错误")
		return
	}

	log.WithFields(
		log.Fields{
			"playerId": pl.GetId(),
		}).Debug("welfare:处理获取魅力排行榜消息完成")
	return nil

}

//获取魅力排行榜界面信息的逻辑
func rankCharmList(pl player.Player, page int32, groupId int32) (err error) {
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
	subType := welfaretypes.OpenActivityRankSubTypeCharm
	checkFlag := welfarelogic.CheckGroupId(pl, typ, subType, groupId)
	if !checkFlag {
		return
	}

	charmList, rankTime := rank.GetRankService().GetPropertyListByPage(ranktypes.RankTypeCharm, ranktypes.RankClassTypeLocalActivity, groupId, page)
	if page == 0 {
		nextPageList, _ := rank.GetRankService().GetPropertyListByPage(ranktypes.RankTypeCharm, ranktypes.RankClassTypeLocalActivity, groupId, page+1)
		if len(nextPageList) > 0 {
			charmList = append(charmList, nextPageList[0])
		}
	}

	welfareManager := pl.GetPlayerDataManager(playertypes.PlayerWelfareDataManagerType).(*playerwelfare.PlayerWelfareManager)
	curAddNum := welfareManager.GetActivityAddNumVal(groupId)
	startTime, endTime := welfare.GetWelfareService().CountOpenActivityTime(groupId)
	scCharmGet := pbutil.BuildSCOpenActivityRankCharmList(page, charmList, rankTime, startTime, endTime, curAddNum, groupId)
	pl.SendMsg(scCharmGet)
	return
}
