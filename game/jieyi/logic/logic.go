package logic

import (
	"context"
	"fgame/fgame/common/message"
	coretypes "fgame/fgame/core/types"
	commomlogic "fgame/fgame/game/common/logic"
	consttypes "fgame/fgame/game/constant/types"
	"fgame/fgame/game/jieyi/jieyi"
	"fgame/fgame/game/jieyi/pbutil"
	playerjieyi "fgame/fgame/game/jieyi/player"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	playerproperty "fgame/fgame/game/property/player"
	playerpropertytypes "fgame/fgame/game/property/player/types"
	"fgame/fgame/game/scene/scene"
	gametemplate "fgame/fgame/game/template"
	"fgame/fgame/pkg/mathutils"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

func JieYiPropertyChange(pl player.Player) {
	propertyManager := pl.GetPlayerDataManager(playertypes.PlayerPropertyDataManagerType).(*playerproperty.PlayerPropertyDataManager)
	propertyManager.UpdateBattleProperty(playerpropertytypes.PlayerPropertyEffectorTypeJieYi.Mask())
}

//结义威名升级判断
func JieYiNameUpLev(curTimesNum int32, temp *gametemplate.JieYiNameLevelTemplate) (success bool) {
	timesMin := temp.TimesMin
	timesMax := temp.TimesMax
	updateRate := temp.UpLevPercent
	_, _, success = commomlogic.GetStatusAndProgress(curTimesNum, 0, timesMin, timesMax, 0, 0, updateRate, 0)
	return
}

//结义信物升级判断
func JieYiTokenUpLev(curTimesNum int32, curBless int32, temp *gametemplate.JieYiTokenLevelTemplate) (pro int32, randBless int32, sucess bool) {
	timesMin := temp.TimesMin
	timesMax := temp.TimesMax
	updateRate := temp.Rate
	blessMax := temp.ZhufuMax
	addMin := temp.AddMin
	addMax := temp.AddMax + 1

	randBless = int32(mathutils.RandomRange(int(addMin), int(addMax)))
	pro, sucess = commomlogic.AdvancedStatusAndProgress(curTimesNum, curBless, timesMin, timesMax, randBless, updateRate, blessMax)
	return
}

//声威值掉落
func ShengWeiZhiDrop(pl player.Player, attackId int64, attackName string, mapId int32, pos coretypes.Position) (itemId int32, dropNum int32) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if !manager.IsCanDropShengWeiZhi() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取声威值掉落,掉落冷却中")
		return
	}

	// 声威id
	itemId = int32(consttypes.ShengWei)

	// 掉落声威值判断
	flag, dropNum, dropLev := manager.ShengWeiDrop()
	if !flag {
		return
	}

	if dropLev > 0 {
		JieYiPropertyChange(pl)
	}

	if dropNum > 0 || dropLev > 0 {
		memberList := jieyi.GetJieYiService().GetJieYiMemberList(manager.GetPlayerJieYiObj().GetJieYiId())
		for _, member := range memberList {
			memberPl := player.GetOnlinePlayerManager().GetPlayerById(member.GetPlayerId())
			if memberPl != nil {
				scMsg := pbutil.BuildSCJieYiShengWeiZhiDrop(dropNum, dropLev, attackName, mapId, pos, pl.GetId(), pl.GetName())
				memberPl.SendMsg(scMsg)
			}
		}
		//TODO:xubin 推送自己？
		obj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
		if obj != nil {
			scMsg := pbutil.BuildSCJieBrotherInfoOnChange(obj)
			pl.SendMsg(scMsg)
		}
	}

	return
}

//声威值掉落
func DropShengWeiZhi(pl player.Player, attackId int64, attackName string) (itemId int32, dropNum int32) {
	manager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	if !manager.IsCanDropShengWeiZhi() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("massacre:处理获取声威值掉落,掉落冷却中")
		return
	}

	// 声威id
	itemId = int32(consttypes.ShengWei)

	// 掉落声威值判断
	flag, dropNum, dropLev := manager.ShengWeiDrop()
	if !flag {
		return
	}

	if dropLev > 0 {
		JieYiPropertyChange(pl)
	}

	// if dropNum > 0 || dropLev > 0 {
	// 	memberList := jieyi.GetJieYiService().GetJieYiMemberList(manager.GetPlayerJieYiObj().GetJieYiId())
	// 	for _, member := range memberList {
	// 		memberPl := player.GetOnlinePlayerManager().GetPlayerById(member.GetPlayerId())
	// 		if memberPl != nil {
	// 			scMsg := pbutil.BuildSCJieYiShengWeiZhiDrop(dropNum, dropLev, attackName, mapId, pos, pl.GetId(), pl.GetName())
	// 			memberPl.SendMsg(scMsg)
	// 		}
	// 	}

	// 	obj := jieyi.GetJieYiService().GetJieYiMemberInfo(pl.GetId())
	// 	if obj != nil {
	// 		scMsg := pbutil.BuildSCJieBrotherInfoOnChange(obj)
	// 		pl.SendMsg(scMsg)
	// 	}
	// }

	return
}

func JieYiMemberChanged(jieYi *jieyi.JieYi) {
	jieYiObj := jieYi.GetJieYiObject()
	//结义信息
	scJieYiInfoOnChange := pbutil.BuildSCJieYiInfoOnChange(jieYiObj)
	BroadcastJieYi(jieYi, scJieYiInfoOnChange)

	for _, mem := range jieYi.GetJieYiMemberList() {
		memId := mem.GetPlayerId()
		memPl := player.GetOnlinePlayerManager().GetPlayerById(memId)
		if memPl == nil {
			continue
		}
		ctx := scene.WithPlayer(context.Background(), memPl)
		msg := message.NewScheduleMessage(onJieYiMemberChanged, ctx, mem, nil)
		memPl.Post(msg)
	}
}

func onJieYiMemberChanged(ctx context.Context, result interface{}, err error) error {
	p := scene.PlayerInContext(ctx)
	pl := p.(player.Player)
	memberObj := result.(*jieyi.JieYiMemberObject)
	jieYi := memberObj.GetJieYi()
	daoJu := memberObj.GetDaoJuType()
	token := memberObj.GetTokenType()
	rank := memberObj.GetRank()
	name := jieYi.GetJieYiName()
	jieYiId := jieYi.GetJieYiObject().GetId()
	// 刷新被邀请人数据
	jieYiManager := pl.GetPlayerDataManager(playertypes.PlayerJieYiDataManagerType).(*playerjieyi.PlayerJieYiDataManager)
	jieYiManager.SyncJieYi(daoJu, token, jieYiId, name, rank)
	//属性变化
	JieYiPropertyChange(pl)
	return nil
}

func BroadcastJieYi(jieYi *jieyi.JieYi, msg proto.Message) {
	memberList := jieYi.GetJieYiMemberList()
	for _, member := range memberList {
		memberPl := player.GetOnlinePlayerManager().GetPlayerById(member.GetPlayerId())
		if memberPl != nil {
			memberPl.SendMsg(msg)
		}
	}
}

func BroadcastJieYiExculdeSelf(jieYi *jieyi.JieYi, pl player.Player, msg proto.Message) {
	memberList := jieYi.GetJieYiMemberList()
	for _, member := range memberList {
		memberPl := player.GetOnlinePlayerManager().GetPlayerById(member.GetPlayerId())
		if memberPl != nil && memberPl != pl {
			memberPl.SendMsg(msg)
		}
	}
}
