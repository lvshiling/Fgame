package listener

import (
	"fgame/fgame/core/event"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/center/center"
	crosslogic "fgame/fgame/game/cross/logic"
	crosstypes "fgame/fgame/game/cross/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playereventtypes "fgame/fgame/game/player/event/types"
	"fgame/fgame/game/player/pbutil"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/logic"
	scenelogic "fgame/fgame/game/scene/logic"
	playerscene "fgame/fgame/game/scene/player"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"

	log "github.com/Sirupsen/logrus"
)

func playerLoadFinish(target event.EventTarget, data event.EventData) (err error) {
	p := target.(player.Player)
	//刷新
	flag, err := p.AfterLoad()
	if err != nil {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
				"error":    err,
			}).Error("player:玩家加载角色数据后,失败")
		//发送异常消息
		//断开链接
		p.Close(nil)
		return
	}
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("player:玩家加载角色数据后,加载失败")
		//发送异常消息
		//断开链接
		p.Close(nil)
		return
	}

	flag = player.GetOnlinePlayerManager().PlayerEnterServer(p)
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": p.GetId(),
			}).Warn("player:玩家角色进入游戏服务器,失败")
		//TODO 断开连接
		p.Close(nil)
		return
	}
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家角色进入游戏服务器,成功")
	allianceFlag := center.GetCenterService().IsAllianceOpen()
	tradeFlag := center.GetCenterService().IsTradeOpen()
	info := pbutil.BuildSCPlayerInfo(p, allianceFlag, tradeFlag)
	p.SendMsg(info)
	gameevent.Emit(playereventtypes.EventTypePlayerAfterLoadFinish, p, nil)

	//p.UpdateBattleProperty(playerpropertytypes.PropertyEffectorTypeMaskAll)
	//判断是否在跨服中
	crossType := p.GetCrossType()
	if crossType != crosstypes.CrossTypeNone {
		//挂机玩家
		if p.IsGuaJiPlayer() {
			crosslogic.PlayerExitCross(p)
		} else {
			//进入跨服
			crosslogic.PlayerReenterCross(p, crossType)
		}
	}

	var mapId int32
	var sceneId int64
	var pos coretypes.Position
	if global.PRESSURE {
		//随机地图
		mapTemplate := scenetemplate.GetSceneTemplateService().RandomWorldMap()
		mapId = int32(mapTemplate.TemplateId())
		pos = mapTemplate.GetMap().RandomPosition()
	} else {
		psdm, _ := p.GetPlayerDataManager(playertypes.PlayerSceneDataManagerType).(*playerscene.PlayerSceneDataManager)
		mapId = psdm.GetPlayerScene().GetMapId()
		sceneId = psdm.GetPlayerScene().GetSceneId()
		pos = coretypes.Position{
			X: psdm.GetPlayerScene().GetPosX(),
			Y: psdm.GetPlayerScene().GetPosY(),
			Z: psdm.GetPlayerScene().GetPosZ(),
		}
	}
	var s scene.Scene
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)

	switch mapTemplate.GetMapType().MapType() {
	// case scenetypes.MapTypeWorldBoss:
	// 	s = scene.GetSceneService().GetWorldBossSceneByMapId(mapId)
	// 	break
	case scenetypes.MapTypeWorld:
		s = scene.GetSceneService().GetWorldSceneByMapId(mapId)
		break
	case scenetypes.MapTypeActivity,
		scenetypes.MapTypeActivitySub:
		s = scene.GetSceneService().GetActivitySceneByMapId(mapId)
		break
	case scenetypes.MapTypeMarry:
		s = scene.GetSceneService().GetMarrySceneByMapId(mapId)
		break
	case scenetypes.MapTypeFuBen:
		s = scene.GetSceneService().GetFuBenSceneById(sceneId)
		break
	case scenetypes.MapTypeTower:
		s = scene.GetSceneService().GetTowerSceneByMapId(mapId)
		break
	case scenetypes.MapTypeBoss:
		s = scene.GetSceneService().GetBossSceneByMapId(mapId)
		break
		// case scenetypes.MapTypeOutlandBoss:
		// 	s = scene.GetSceneService().GetOutlandBossSceneByMapId(mapId)
		// 	break
		// case scenetypes.MapTypeActivityFuBen:
		// 	s = scene.GetSceneService().GetActivityFuBenSceneById(sceneId)
		// 	break
		// case scenetypes.MapTypeCangJingGe:
		// 	s = scene.GetSceneService().GetCangJingGeSceneByMapId(mapId)
		// 	break
		// case scenetypes.MapTypeZhenXiBoss:
		// 	s = scene.GetSceneService().GetZhenXiSceneByMapId(mapId)
		// 	break
		// case scenetypes.MapTypeDingShiBoss:
		// 	s = scene.GetSceneService().GetDingShiSceneByMapId(mapId)
		// 	break
	}
	if !mapTemplate.GetMapType().IsReConnect() {
		s = nil
	}

	if s == nil {
		if mapTemplate.IsWorld() {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
				}).Error("player:玩家角色进入游戏服务器,世界场景不存在")
			//TODO 断线
			err = fmt.Errorf("世界场景[%d]不存在", mapId)
			return
		}
		//返回上一个场景
		scenelogic.PlayerBackLastScene(p)
		return
	}
	//挂机
	if p.IsGuaJiPlayer() {
		if !mapTemplate.IsWorld() {
			scenelogic.PlayerBackLastScene(p)
			return
		}
	}
	//不在掩码上,回到出生点
	if !s.MapTemplate().GetMap().IsWalkable(pos.X, pos.Z) {
		pos = s.MapTemplate().GetBornPos()
	}
	if !logic.PlayerEnterScene(p, s, pos) {
		scenelogic.PlayerBackLastScene(p)
	}
	return nil
}

func init() {
	gameevent.AddEventListener(playereventtypes.EventTypePlayerLoadFinish, event.EventListenerFunc(playerLoadFinish))
}
