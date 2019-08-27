package logic

import (
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	moonloveeventtypes "fgame/fgame/game/moonlove/event/types"
	"fgame/fgame/game/moonlove/pbutil"
	moonlovetypes "fgame/fgame/game/moonlove/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"fmt"
	"sort"
	"sync"
)

var (
	moonLoveSceneMutex sync.Mutex
)

const (
	MAX_RANKING_TOP = 3
	MAX_RANKING     = 10
)

func getMoonLoveScene(activityTemplate *gametemplate.ActivityTemplate) (s scene.Scene) {
	//加锁
	moonLoveSceneMutex.Lock()
	defer moonLoveSceneMutex.Unlock()

	s = scene.GetSceneService().GetActivitySceneByMapId(activityTemplate.Mapid)
	if s == nil {
		//TODO 修改
		now := global.GetGame().GetTimeService().Now()
		var activityTimeTemplate *gametemplate.ActivityTimeTemplate
		for _, tempActivityTimeTemplate := range activityTemplate.GetTimeList() {
			startTime, _ := tempActivityTimeTemplate.GetBeginTime(now)
			endTime, _ := tempActivityTimeTemplate.GetEndTime(now)
			if now >= startTime && now <= endTime {
				activityTimeTemplate = tempActivityTimeTemplate
				break
			}
		}

		if activityTimeTemplate == nil {
			return nil
		}
		endTime, _ := activityTimeTemplate.GetEndTime(now)
		sd := createMoonloveSceneData()
		s = createMoonLoveScene(activityTemplate.Mapid, endTime, sd)
		if s == nil {
			return nil
		}
	}
	return
}

func createMoonLoveScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeYueXiaQingYuan {
		return nil
	}
	s = scene.CreateActivityScene(mapId, endTime, sh)
	return s
}

type MoonloveSceneData interface {
	scene.SceneDelegate
	GetGenerousRank() []*moonlovetypes.RankData                                  //豪气榜
	GetCharmRank() []*moonlovetypes.RankData                                     //魅力榜
	IsCouple(playerId int64) bool                                                //是否双人赏月
	ReleaseCouple(playerId int64)                                                //解除双人赏月
	CombineCouple(playerId, otherPlayerId int64)                                 //双人赏月
	ChramRankChange(friCharm int32, friName string, friId int64) int32           //计算魅力排名
	GenerousRankChange(plGenerousNum int32, sendName string, sendId int64) int32 //计算豪气排名
	PlayerNameChanged(pl player.Player)                                          //玩家修改名字
	GetCoupleMap() map[int64]*moonlovetypes.MoonloveDoubleData
}

//月下情缘场景数据
type moonloveSceneData struct {
	*scene.SceneDelegateBase
	s             scene.Scene
	charmRank     []*moonlovetypes.RankData                   //魅力榜
	generousRank  []*moonlovetypes.RankData                   //豪气榜
	playerExpMap  map[int64]int64                             //玩家经验奖励
	coupleDataMap map[int64]*moonlovetypes.MoonloveDoubleData //玩家赏月状态数据
}

func (sd *moonloveSceneData) GetScene() scene.Scene {
	return sd.s
}

func (sd *moonloveSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
}

func (sd *moonloveSceneData) OnSceneTick(s scene.Scene) {

	for _, pl := range s.GetAllPlayers() {
		pl := pl.(player.Player)
		isDouble := sd.IsCouple(pl.GetId())

		onSceneTickRew(pl, isDouble)
	}
}

//场景完成
func (sd *moonloveSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}

	for playerId, spl := range s.GetAllPlayers() {
		expCount := sd.playerExpMap[playerId]
		pl := spl.(player.Player)
		onFinishMoonlove(pl, expCount)
	}

	//发放排行榜奖励
	onSendRankRewards(sd.charmRank, sd.generousRank)

}
func (sd *moonloveSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}
}

//生物进入
func (sd *moonloveSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *moonloveSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

func (sd *moonloveSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

func (sd *moonloveSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {

}

func (sd *moonloveSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//玩家重生
func (sd *moonloveSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

func (sd *moonloveSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {

}

//玩家退出场景
func (sd *moonloveSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("xianfu:仙府应该是同一个场景"))
	}

	state := sd.IsCouple(p.GetId())
	if state {
		sd.ReleaseCouple(p.GetId())
	}
}

func (sd *moonloveSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入场景
func (sd *moonloveSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("moonlove:月下情缘应该是同一个场景"))
	}
	_, isExist := sd.playerExpMap[p.GetId()]
	if !isExist {
		sd.playerExpMap[p.GetId()] = 0
	}

	pl := p.(player.Player)
	sceneStartTime := s.GetStartTime()
	exp := sd.playerExpMap[p.GetId()]

	onScenePlayerEnter(pl, s.GetEndTime())
	onPushSceneInfo(pl, sd.charmRank, sd.generousRank, exp, sceneStartTime)

	//判断玩家是否改过名字
	sd.PlayerNameChanged(pl)
}

//场景获取物品
func (sd *moonloveSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {

}

//玩家获得经验
func (sd *moonloveSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("moonlove:月下情缘应该是同一个场景"))
	}
	_, isExist := sd.playerExpMap[p.GetId()]
	if isExist {
		sd.playerExpMap[p.GetId()] += num
	} else {
		sd.playerExpMap[p.GetId()] = num
	}

	pl := p.(player.Player)
	exp := sd.playerExpMap[p.GetId()]
	onPushSceneExpInfo(pl, exp)
}

//刷怪
func (sd *moonloveSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

func (sd *moonloveSceneData) GetCharmRank() []*moonlovetypes.RankData {
	return sd.charmRank
}

func (sd *moonloveSceneData) GetGenerousRank() []*moonlovetypes.RankData {
	return sd.generousRank
}

func (sd *moonloveSceneData) IsCouple(playerId int64) bool {
	_, isCouple := sd.coupleDataMap[playerId]
	return isCouple
}

func (sd *moonloveSceneData) ReleaseCouple(playerId int64) {

	moonData, isExist := sd.coupleDataMap[playerId]
	if isExist {
		delete(sd.coupleDataMap, moonData.PlayerId)
		delete(sd.coupleDataMap, moonData.OhtherPayerId)

		ed := moonloveeventtypes.CreateMoonloveDoubleReleaseEventData(moonData.PlayerId, moonData.OhtherPayerId)
		gameevent.Emit(moonloveeventtypes.EventTypeMoonloveDoubleRelease, nil, ed)

	}

}

func (sd *moonloveSceneData) CombineCouple(playerId, otherPlayerId int64) {

	moonData := createMoonData(playerId, otherPlayerId)
	sd.coupleDataMap[playerId] = moonData
	sd.coupleDataMap[otherPlayerId] = moonData
	ed := moonloveeventtypes.CreateMoonloveDoubleCombineEventData(moonData.PlayerId, moonData.OhtherPayerId)
	gameevent.Emit(moonloveeventtypes.EventTypeMoonloveDoubleCombine, nil, ed)
}

func createMoonData(playerId, otherPlayerId int64) *moonlovetypes.MoonloveDoubleData {
	moonData := &moonlovetypes.MoonloveDoubleData{}
	moonData.PlayerId = playerId
	moonData.OhtherPayerId = otherPlayerId

	return moonData
}

func (sd *moonloveSceneData) getRankingIndex(rankTyp moonlovetypes.MoonloveRankType, playerId int64) (index int32, rankData *moonlovetypes.RankData) {
	var rankList []*moonlovetypes.RankData
	if rankTyp == moonlovetypes.MoonloveRankTypeGenerousCharm {
		rankList = sd.charmRank
	} else {
		rankList = sd.generousRank
	}
	for index, rankData := range rankList {
		if rankData.PlayerId == playerId {
			return int32(index), rankData
		}
	}
	return -1, nil
}

func (sd *moonloveSceneData) ChramRankChange(friCharm int32, friName string, friId int64) (ranking int32) {
	//更新
	beforeIndex, beforeRankData := sd.getRankingIndex(moonlovetypes.MoonloveRankTypeGenerousCharm, friId)
	if beforeRankData == nil {
		newRankData := &moonlovetypes.RankData{}
		newRankData.Number = friCharm
		newRankData.Name = friName
		newRankData.PlayerId = friId
		sd.charmRank = append(sd.charmRank, newRankData)
	} else {
		beforeRankData.Number = friCharm
	}

	sort.Sort(sort.Reverse(moonlovetypes.RankDataList(sd.charmRank)))
	sortRanks := make([]*moonlovetypes.RankData, len(sd.charmRank))
	copy(sortRanks, sd.charmRank)
	if len(sd.charmRank) > MAX_RANKING {
		sd.charmRank = sd.charmRank[:MAX_RANKING]
	}

	//不上榜或排名不变
	afterIndex, curRankData := sd.getRankingIndex(moonlovetypes.MoonloveRankTypeGenerousCharm, friId)
	ranking = afterIndex + 1
	if ranking <= MAX_RANKING {
		//广播前十
		for plId, _ := range sd.s.GetAllPlayers() {
			scenePlayer := sd.s.GetPlayer(plId)
			if scenePlayer == nil {
				continue
			}

			scPushRank := pbutil.BuildMoonlovePushCharmRank(sd.charmRank)
			scenePlayer.SendMsg(scPushRank)
		}
	}

	if curRankData == nil || beforeIndex == afterIndex {
		return
	}

	pushCharmRankChanged(sd.s, sortRanks, int(afterIndex))
	return
}

func (sd *moonloveSceneData) GenerousRankChange(generousNum int32, sendName string, sendId int64) (ranking int32) {
	//更新
	beforeIndex, beforeRankData := sd.getRankingIndex(moonlovetypes.MoonloveRankTypeGenerous, sendId)
	if beforeRankData == nil {
		newRankData := &moonlovetypes.RankData{}
		newRankData.Number = generousNum
		newRankData.Name = sendName
		newRankData.PlayerId = sendId
		sd.generousRank = append(sd.generousRank, newRankData)

	} else {
		beforeRankData.Number = generousNum
	}

	sort.Sort(sort.Reverse(moonlovetypes.RankDataList(sd.generousRank)))
	sortRanks := make([]*moonlovetypes.RankData, len(sd.generousRank))
	copy(sortRanks, sd.generousRank)
	if len(sd.generousRank) > MAX_RANKING {
		sd.generousRank = sd.generousRank[:MAX_RANKING]
	}

	//不上榜或排名不变
	afterIndex, curRankData := sd.getRankingIndex(moonlovetypes.MoonloveRankTypeGenerous, sendId)
	ranking = afterIndex + 1
	if ranking <= MAX_RANKING {
		//广播前十
		for plId, _ := range sd.s.GetAllPlayers() {
			scenePlayer := sd.s.GetPlayer(plId)
			if scenePlayer == nil {
				continue
			}
			scPushRank := pbutil.BuildMoonlovePushGenerousRank(sd.generousRank)
			scenePlayer.SendMsg(scPushRank)
		}
	}

	if curRankData == nil || beforeIndex == afterIndex {
		return
	}

	pushGenerousChanged(sd.s, sortRanks, int(afterIndex))
	return
}

func (sd *moonloveSceneData) PlayerNameChanged(pl player.Player) {
	changed := false
	for index, charmData := range sd.charmRank {
		if charmData.PlayerId == pl.GetId() {
			charmData.Name = pl.GetName()
			if index < MAX_RANKING_TOP {
				changed = true
			}
			break
		}
	}
	if changed {
		scPushRank := pbutil.BuildMoonlovePushCharmRank(sd.charmRank)
		sd.GetScene().BroadcastMsg(scPushRank)
	}
	changed = false
	for index, generousData := range sd.generousRank {
		if generousData.PlayerId == pl.GetId() {
			generousData.Name = pl.GetName()
			if index < MAX_RANKING_TOP {
				changed = true
			}
			break
		}
	}
	if changed {
		scPushRank := pbutil.BuildMoonlovePushGenerousRank(sd.generousRank)
		sd.GetScene().BroadcastMsg(scPushRank)
	}
}

func (sd *moonloveSceneData) GetCoupleMap() map[int64]*moonlovetypes.MoonloveDoubleData {
	return sd.coupleDataMap
}

func createMoonloveSceneData() MoonloveSceneData {
	sd := &moonloveSceneData{
		playerExpMap:  make(map[int64]int64),
		coupleDataMap: make(map[int64]*moonlovetypes.MoonloveDoubleData),
	}
	sd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return sd
}
