package logic

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	coreutils "fgame/fgame/core/utils"
	activitytypes "fgame/fgame/game/activity/types"
	emaillogic "fgame/fgame/game/email/logic"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/longgong/pbutil"
	longgongtemplate "fgame/fgame/game/longgong/template"
	longgongtypes "fgame/fgame/game/longgong/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	welfarelogic "fgame/fgame/game/welfare/logic"

	"fmt"

	log "github.com/Sirupsen/logrus"
)

func PlayerEnterLongGongScene(pl player.Player, activityTemplate *gametemplate.ActivityTemplate, args ...string) (flag bool, err error) {
	ns := pl.GetScene()
	mapType := ns.MapTemplate().GetMapType()
	if mapType == scenetypes.SceneTypeLongGong {
		return
	}
	s := getLongGongScene(activityTemplate)
	if s == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("longgong:龙宫探宝场景不存在")
		playerlogic.SendSystemMessage(pl, lang.ActivityNotAtTime)
		return
	}

	bornPos := s.MapTemplate().GetBornPos()
	if !scenelogic.PlayerEnterScene(pl, s, bornPos) {
		return
	}
	flag = true
	return
}

// 龙宫探宝结束
func onFinishLongGong(s scene.Scene) {
	//通知活动结束
	plM := s.GetAllPlayers()
	for _, pl := range plM {
		scMsg := pbutil.BuildSCLonggongResult()
		pl.SendMsg(scMsg)
	}
	return
}

//龙宫排行奖励
func LongGongRankRewards(sd LongGongSceneData) {
	s := sd.GetScene()
	if s == nil {
		return
	}
	//邮件奖励
	rankList := s.GetAllRankList(longgongtypes.LongGongSceneRankTypeDamage)
	longGongService := longgongtemplate.GetLongGongTemplateService()

	for idx, rank := range rankList {
		num := int32(idx + 1)
		rankRewTemp := longGongService.GetLongGongRankTemplateByRankNum(num)
		if rankRewTemp == nil {
			continue
		}

		primevalTitle := lang.GetLangService().ReadLang(lang.LongGongEmailTitle)
		title := coreutils.FormatNoticeStr(primevalTitle)
		econtent := fmt.Sprintf(lang.GetLangService().ReadLang(lang.LongGongEmailContent), num)
		rewItemMap := rankRewTemp.GetRewEmailItemMap()
		expireType := inventorytypes.NewItemLimitTimeTypeNone
		expireTime := int64(0)
		newItemDataList := welfarelogic.ConvertToItemData(rewItemMap, expireType, expireTime)
		playerId := rank.GetPlayerId()
		endTime := s.GetEndTime()
		pl := player.GetOnlinePlayerManager().GetPlayerById(playerId)
		if pl == nil {
			emaillogic.AddOfflineEmailItemLevel(playerId, title, econtent, endTime, newItemDataList)
		} else {
			//异步发送
			ctx := scene.WithPlayer(context.Background(), pl)
			data := longgongtypes.NewLongGongRankEmailData(title, econtent, endTime, newItemDataList)
			pl.Post(message.NewScheduleMessage(sendLongGongRankRewards, ctx, data, nil))
		}
	}

	return
}

//龙宫伤害排行榜奖励邮件回调
func sendLongGongRankRewards(ctx context.Context, result interface{}, err error) error {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(player.Player)

	data := result.(*longgongtypes.LongGongRankEmailData)
	title := data.Title
	econtent := data.Econtent
	endTime := data.EndTime
	newItemDataList := data.RewItemDataList

	emaillogic.AddEmailItemLevel(tpl, title, econtent, endTime, newItemDataList)
	return nil
}

// 进入龙宫探宝
func onEnterLongGong(spl scene.Player, sd LongGongSceneData) {
	s := sd.GetScene()
	if s == nil {
		return
	}

	pl, ok := spl.(player.Player)
	if !ok {
		return
	}

	endTime := s.GetEndTime()
	pl.EnterActivity(activitytypes.ActivityTypeLongGong, endTime)
	rankMap := s.GetAllRanks()

	curHp := int64(0)
	heilongStatus := sd.GetHeiLongBossInfo().GetStatus()
	if heilongStatus == longgongtypes.HeiLongStatusTypeLive {
		bossNpc := sd.GetHeiLongBossInfo().GetNpc()
		if bossNpc != nil {
			curHp = bossNpc.GetHP()
		}
	}
	pCollectCount := sd.GetPlayerTreasureCollectCount(spl.GetId())
	tPearlCount := sd.GetPearlCollectCount()
	cn := sd.GetTreasureCollectPointNpc()

	scMsg := pbutil.BuildSCLonggongGet(spl, rankMap, curHp, heilongStatus, pCollectCount, tPearlCount, cn)
	pl.SendMsg(scMsg)
}
