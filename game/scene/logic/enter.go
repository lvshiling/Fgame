package logic

import (
	"context"
	"fgame/fgame/common/lang"
	"fgame/fgame/common/message"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	"fgame/fgame/game/player"
	playerlogic "fgame/fgame/game/player/logic"
	"fgame/fgame/game/scene/pbutil"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"runtime/debug"

	log "github.com/Sirupsen/logrus"
)

//玩家进入地图位置
func PlayerEnterMapWithPos(pl scene.Player, mapId int32, pos coretypes.Position) (err error) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入场景,场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	ppl, ok := pl.(player.Player)
	if ok && !playerlogic.CheckCanEnterScene(ppl) {
		return
	}

	//TODO: 优化
	if !mapTemplate.IsWorld() && !mapTemplate.IsActivity() && !mapTemplate.IsActivitySub() && !mapTemplate.IsMarry() && !mapTemplate.IsBoss() && !mapTemplate.IsTower() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入场景,场景暂时未实现")
		playerlogic.SendSystemMessage(pl, lang.SceneNotWorldScene)
		return
	}
	var s scene.Scene
	if mapTemplate.IsWorld() {
		s = scene.GetSceneService().GetWorldSceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	} else if mapTemplate.IsActivity() || mapTemplate.IsActivitySub() {
		s = scene.GetSceneService().GetActivitySceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	} else if mapTemplate.IsMarry() {
		s = scene.GetSceneService().GetMarrySceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	} else if mapTemplate.IsTower() {
		s = scene.GetSceneService().GetTowerSceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	} else {
		s = scene.GetSceneService().GetBossSceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	}

	PlayerEnterScene(pl, s, pos)
	return
}

//玩家进入地图位置
func PlayerEnterMapWithPortal(pl scene.Player, mapId int32, pos coretypes.Position) (err error) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入场景,场景不存在")
		playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
		return
	}

	ppl, ok := pl.(player.Player)
	if ok && !playerlogic.CheckCanEnterScene(ppl) {
		return
	}

	//TODO: 优化
	if !mapTemplate.IsWorld() && !mapTemplate.IsActivity() && !mapTemplate.IsActivitySub() && !mapTemplate.IsMarry() && !mapTemplate.IsBoss() && !mapTemplate.IsTower() {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
				"mapId":    mapId,
			}).Warn("scene:处理进入场景,场景暂时未实现")
		playerlogic.SendSystemMessage(pl, lang.SceneNotWorldScene)
		return
	}
	var s scene.Scene
	if mapTemplate.IsWorld() {
		s = scene.GetSceneService().GetWorldSceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	} else if mapTemplate.IsActivity() || mapTemplate.IsActivitySub() {
		s = scene.GetSceneService().GetActivitySceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	} else if mapTemplate.IsMarry() {
		s = scene.GetSceneService().GetMarrySceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	} else if mapTemplate.IsTower() {
		s = scene.GetSceneService().GetTowerSceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	} else {
		s = scene.GetSceneService().GetBossSceneByMapId(mapId)
		if s == nil {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
					"mapId":    mapId,
				}).Warn("scene:处理进入场景,场景不存在")
			playerlogic.SendSystemMessage(pl, lang.SceneMapNoExist)
			return
		}
	}

	PlayerPortalEnterScene(pl, s, pos)
	return
}

//用户离开跨服 回到原场景
func PlayerEnterOriginScene(pl scene.Player) bool {
	originScene := pl.GetScene()
	if originScene != nil {
		return false
	}
	mapId := pl.GetMapId()
	pos := pl.GetPos()

	s := scene.GetSceneService().GetWorldSceneByMapId(mapId)
	if s == nil {
		return PlayerBackLastScene(pl)
	}

	return PlayerEnterScene(pl, s, pos)

}

//用户返回上一个场景
func AsyncPlayerBackLastScene(pl scene.Player) {
	ctx := scene.WithPlayer(context.Background(), pl)
	pl.Post(message.NewScheduleMessage(onPlayerBackLastScene, ctx, nil, nil))
}

func onPlayerBackLastScene(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(scene.Player)
	PlayerBackLastScene(tpl)
	return nil
}

//用户返回上一个场景
func PlayerBackLastScene(pl scene.Player) bool {
	var pos coretypes.Position
	var s scene.Scene
	dispatchLevel := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBackToZhuChengLevel))
	if dispatchLevel > pl.GetLevel() {
		lastMapId := pl.GetLastMapId()
		s = scene.GetSceneService().GetWorldSceneByMapId(lastMapId)
		pos = pl.GetLastPos()
	} else {
		dispatchMapId := int32(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeBackToZhuChengMapId))
		s = scene.GetSceneService().GetWorldSceneByMapId(dispatchMapId)
		pos = s.MapTemplate().GetRebornPos()
	}

	if s == nil {
		//TODO 以防卡死 需要有一个默认地图
		// err = fmt.Errorf("上一个场景[%d]不存在", lastMapId)
		return false
	}

	originScene := pl.GetScene()
	if originScene != nil {
		if originScene.MapTemplate().IsMarry() {
			//返回出生地点
			pos = s.MapTemplate().GetBornPos()
		}
	}

	return PlayerEnterScene(pl, s, pos)

}

//用户进入副本
func PlayerEnterSingleFuBenScene(pl scene.Player, s scene.Scene) bool {
	bornPos := s.MapTemplate().GetBornPos()
	return PlayerEnterScene(pl, s, bornPos)
}

//用户进入场景
func PlayerEnterScene(pl scene.Player, s scene.Scene, pos coretypes.Position) bool {
	return playerCommonEnterScene(pl, s, pos, scenetypes.SceneEnterTypeCommon)
}

//用户进入场景(传送阵)
func PlayerPortalEnterScene(pl scene.Player, s scene.Scene, pos coretypes.Position) bool {
	return playerCommonEnterScene(pl, s, pos, scenetypes.SceneEnterTypePortal)
}

//用户进入场景(定位其他玩家传送)
func PlayerTrackEnterScene(pl scene.Player, s scene.Scene, pos coretypes.Position) bool {
	return playerCommonEnterScene(pl, s, pos, scenetypes.SceneEnterTypeTrac)
}

//用户进入场景()
func playerCommonEnterScene(pl scene.Player, s scene.Scene, pos coretypes.Position, enterType scenetypes.SceneEnterType) bool {
	//限制进入方式
	if !s.MapTemplate().IfCanEnter(enterType) {
		log.WithFields(
			log.Fields{
				"playerId":  pl.GetId(),
				"enterType": enterType,
			}).Warn("scene:进入场景失败，该地图不允许的进入方式")
		playerlogic.SendSystemMessage(pl, lang.SceneEnterTypeError)
		return false
	}

	enterFlag, pos := scene.CheckEnterScene(pl, s, pos, enterType)
	if !enterFlag {
		return false
	}

	originS := pl.GetScene()
	//退出场景
	if originS != nil {
		if s == originS {
			log.WithFields(
				log.Fields{
					"playerId": pl.GetId(),
				}).Warn("player:玩家角色重复进入场景")
			playerlogic.SendSystemMessage(pl, lang.SceneRepeatEnter)
			return false
		}
		//退出场景 失败
		PlayerExitScene(pl, true)
	}

	//进入场景
	flag := pl.EnteringScene()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene:玩家尝试进入场景失败")
		return false
	}

	ctx := scene.WithScene(context.Background(), s)
	//TODO 处理场景已经 关闭了
	//异步进入场景
	result := &playerEnterResult{
		pl:  pl,
		pos: pos,
	}
	s.Post(message.NewScheduleMessage(onPlayerEnterScene, ctx, result, nil))
	return true
}

//用户重新进入
func PlayerReenterScene(pl scene.Player) bool {
	s := pl.GetScene()
	if s == nil {
		return false
	}
	pos := pl.GetPosition()
	//退出场景 失败
	PlayerExitScene(pl, true)

	//进入场景
	flag := pl.EnteringScene()
	if !flag {
		log.WithFields(
			log.Fields{
				"playerId": pl.GetId(),
			}).Warn("scene:玩家尝试进入场景失败")
		return false
	}

	ctx := scene.WithScene(context.Background(), s)
	//TODO 处理场景已经 关闭了
	//异步进入场景
	result := &playerEnterResult{
		pl:  pl,
		pos: pos,
	}
	s.Post(message.NewScheduleMessage(onPlayerEnterScene, ctx, result, nil))
	return true
}

type playerEnterAndExitResult struct {
	s   scene.Scene
	pos coretypes.Position
}

//用户异步进入场景跳到另外一个场景
func AsyncPlayerEnterScene(pl scene.Player, s scene.Scene, pos coretypes.Position) (err error) {
	originS := pl.GetScene()
	if originS == nil {
		PlayerEnterScene(pl, s, pos)
		return nil
	} else {
		//进入场景

		result := &playerEnterAndExitResult{
			s:   s,
			pos: pos,
		}
		ctx := scene.WithPlayer(context.Background(), pl)
		pl.Post(message.NewScheduleMessage(onPlayerExitAndEnterScene, ctx, result, nil))
	}
	return nil
}

func onPlayerExitAndEnterScene(ctx context.Context, result interface{}, err error) (rerr error) {
	pl := scene.PlayerInContext(ctx)
	tpl := pl.(scene.Player)
	playerEnterAndExitResult := result.(*playerEnterAndExitResult)
	s := playerEnterAndExitResult.s
	pos := playerEnterAndExitResult.pos
	PlayerEnterScene(tpl, s, pos)
	return nil
}

const (
	defaultAngle = 0
)

type playerEnterResult struct {
	pl  scene.Player
	pos coretypes.Position
}

//回调
func onPlayerEnterScene(ctx context.Context, result interface{}, err error) (rerr error) {
	re := result.(*playerEnterResult)
	sc := scene.SceneInContext(ctx)
	p := re.pl
	var flag bool
	//防止卡死在场景里
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"playerId":         p.GetId(),
					"error":            r,
					"exceptionContent": exceptionContent,
				}).Error("player:玩家进入游戏,异常")
			//TODO 发送异常代码
			p.Close(nil)
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
			return
		}
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
				}).Error("player:玩家进入游戏失败")
			p.Close(nil)
		}
	}()

	pos := re.pos
	if !sc.MapTemplate().GetMap().IsWalkable(pos.X, pos.Z) {
		pos = sc.MapTemplate().GetBornPos()
	}
	pos.Y = sc.MapTemplate().GetMap().GetHeight(pos.X, pos.Z)
	p.SetEnterPos(pos)

	//进入游戏
	flag = p.EnterGame()
	//已经登出了
	if !flag {
		return
	}

	//设置场景
	sc.AddSceneObject(p)
	//发送进入场景数据
	enterSceneMsg := pbutil.BuildEnterScene(p)
	rerr = p.SendMsg(enterSceneMsg)
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"mapId":    sc.MapId(),
		}).Info("player:玩家进入场景")

	return
}

//npc进入
type npcEnterResult struct {
	n   scene.NPC
	pos coretypes.Position
}

//npc进入场景
func NPCEnterScene(n scene.NPC, s scene.Scene, pos coretypes.Position) {
	os := n.GetScene()
	if os == s {
		return
	}
	if os != nil {
		//移除
		os.RemoveSceneObject(n, true)
	}

	ctx := scene.WithScene(context.Background(), s)
	result := &npcEnterResult{
		n:   n,
		pos: pos,
	}
	//异步进入场景
	s.Post(message.NewScheduleMessage(onNPCEnterScene, ctx, result, nil))
}

//回调
func onNPCEnterScene(ctx context.Context, result interface{}, err error) (rerr error) {
	tResult := result.(*npcEnterResult)
	s := scene.SceneInContext(ctx)
	n := tResult.n
	pos := tResult.pos
	n.SetPosition(pos)
	s.AddSceneObject(n)
	return
}
