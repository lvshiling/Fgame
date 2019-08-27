package scene

import (
	"fgame/fgame/core/heartbeat"
	activitytypes "fgame/fgame/game/activity/types"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	lianyueventtypes "fgame/fgame/game/lianyu/event/types"
	lianyutypes "fgame/fgame/game/lianyu/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
	"sort"
)

//无间炼狱战场场景
func CreateLianYuScene(mapId int32, endTime int64, sd LianYuSceneData) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	activityType, ok := mapTemplate.GetMapType().ToActivityType()
	if !ok {
		return nil
	}
	if activityType != sd.GetAcitvityType() {
		return nil
	}
	s = scene.CreateScene(mapTemplate, endTime, sd)
	return s
}

type LianYuSceneData interface {
	scene.SceneDelegate

	//玩家人数
	GetScenePlayerNum() int32
	//获取boss
	GetBoss() *lianYuBoss
	//获取排行榜
	GetRank() []*LianYuRank
	//获取物品
	GetItemMap() map[int64]map[int32]int32
	//获取当前杀气
	GetShaQiNum(playerId int64) int32
	//获取活动类型
	GetAcitvityType() activitytypes.ActivityType
}

type LianYuRank struct {
	serviceId int32
	playerId  int64
	name      string
	shaqi     int32
}

func newLianYuRank(serverId int32, playerId int64, name string, shaqi int32) *LianYuRank {
	d := &LianYuRank{
		serviceId: serverId,
		playerId:  playerId,
		name:      name,
		shaqi:     shaqi,
	}
	return d
}

func (lr *LianYuRank) GetServerId() int32 {
	return lr.serviceId
}

func (lr *LianYuRank) GetName() string {
	return lr.name
}

func (lr *LianYuRank) GetShaQi() int32 {
	return lr.shaqi
}

//杀气记录排序
type LianYuRanktList []*LianYuRank

func (lrl LianYuRanktList) Len() int {
	return len(lrl)
}

func (lrl LianYuRanktList) Less(i, j int) bool {
	return lrl[i].shaqi < lrl[j].shaqi
}

func (lrl LianYuRanktList) Swap(i, j int) {
	lrl[i], lrl[j] = lrl[j], lrl[i]
}

type lianYuBoss struct {
	npc scene.NPC
	//boss状态
	bossStatus lianyutypes.LianYuBossStatusType
}

func newLianYuBoss() *lianYuBoss {
	d := &lianYuBoss{
		bossStatus: lianyutypes.LianYuBossStatusTypeInit,
	}
	return d
}

func (l *lianYuBoss) GetBossStatus() lianyutypes.LianYuBossStatusType {
	return l.bossStatus
}

func (l *lianYuBoss) GetNpc() scene.NPC {
	return l.npc
}

//无间炼狱战场数据
type lianYuSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//无间炼狱活动开始时间
	starTime int64
	//排行榜
	lianYuRankList []*LianYuRank
	//玩家获得的物品
	itemInfoMap map[int64]map[int32]int32
	//玩家杀气
	lianYuShaQiMap map[int64]*LianYuRank
	//当前人数
	num int32
	//boss
	boss *lianYuBoss
	//活动类型
	activityType activitytypes.ActivityType
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func CreateLianYuSceneData(activityType activitytypes.ActivityType) LianYuSceneData {
	csd := &lianYuSceneData{
		num:            0,
		lianYuRankList: make([]*LianYuRank, 0, lianyutypes.ShaQiRankSize),
		itemInfoMap:    make(map[int64]map[int32]int32),
		lianYuShaQiMap: make(map[int64]*LianYuRank),
		activityType:   activityType,
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

func (sd *lianYuSceneData) GetScene() (s scene.Scene) {
	return sd.s
}

//场景开始
func (sd *lianYuSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	sd.boss = newLianYuBoss()
	sd.starTime = s.GetStartTime()

	//心跳任务
	sd.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	sd.heartbeatRunner.AddTask(CreateLianYuTask(sd, sd.starTime))
}

//刷怪
func (sd *lianYuSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *lianYuSceneData) OnSceneTick(s scene.Scene) {
	sd.Heartbeat()
}

//怪物死亡
func (sd *lianYuSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//生物进入
func (sd *lianYuSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
	//Boss刷新
	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	if npcType == scenetypes.BiologyScriptTypeCrossLianYuBoss {
		sd.boss.npc = npc
		sd.boss.bossStatus = lianyutypes.LianYuBossStatusTypeLive
		//发送事件
		gameevent.Emit(lianyueventtypes.EventTypeLianYuBossStatusRefresh, sd, nil)
	}
}

func (sd *lianYuSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *lianYuSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
	//Boss死亡
	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	if npcType == scenetypes.BiologyScriptTypeCrossLianYuBoss {
		sd.boss.bossStatus = lianyutypes.LianYuBossStatusTypeDead
		sd.boss.npc = nil
		//发送事件
		gameevent.Emit(lianyueventtypes.EventTypeLianYuBossStatusRefresh, sd, nil)
	}
}

//生物重生
func (sd *lianYuSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
}

//玩家复活
func (sd *lianYuSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}

}

//玩家死亡
func (sd *lianYuSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
}

func (sd *lianYuSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *lianYuSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
	sd.num++
	//发送事件
	gameevent.Emit(lianyueventtypes.EventTypeLianYuPlayerEnter, sd, p)

}

//玩家退出
func (sd *lianYuSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
	sd.num--
	gameevent.Emit(lianyueventtypes.EventTypeLianYuPlayerExit, sd, nil)
}

//场景完成
func (sd *lianYuSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
	gameevent.Emit(lianyueventtypes.EventTypeLianYuSceneFinish, sd, nil)
}

//场景退出了
func (sd *lianYuSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
}

//场景获取物品
func (sd *lianYuSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
	itemId := itemData.ItemId
	itemNum := itemData.Num
	playerId := pl.GetId()
	itemMap, exist := sd.itemInfoMap[playerId]
	if !exist {
		itemMap = make(map[int32]int32)
		sd.itemInfoMap[playerId] = itemMap
	}
	itemMap[itemId] += itemNum

	sd.refreshRank(pl, itemId, itemNum)
}

//玩家获得经验
func (sd *lianYuSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("lianyu:无间炼狱战场应该是同一个场景"))
	}
}

//心跳
func (sd *lianYuSceneData) Heartbeat() {
	sd.heartbeatRunner.Heartbeat()
}

//玩家人数
func (sd *lianYuSceneData) GetScenePlayerNum() int32 {
	return sd.num
}

//获取boss
func (sd *lianYuSceneData) GetBoss() *lianYuBoss {
	return sd.boss
}

//获取排行榜
func (sd *lianYuSceneData) GetRank() []*LianYuRank {
	return sd.lianYuRankList
}

//获取物品
func (sd *lianYuSceneData) GetItemMap() map[int64]map[int32]int32 {
	return sd.itemInfoMap
}

//获取杀气数量
func (sd *lianYuSceneData) GetShaQiNum(playerId int64) int32 {
	data, ok := sd.lianYuShaQiMap[playerId]
	if !ok {
		return 0
	}
	return data.shaqi
}

//获取活动类型
func (sd *lianYuSceneData) GetAcitvityType() activitytypes.ActivityType {
	return sd.activityType
}

func (sd *lianYuSceneData) getShaQiRankIdList() (lianYuRanList []*LianYuRank) {
	for index, shaQiRank := range sd.lianYuRankList {
		if index > lianyutypes.ShaQiRankSize {
			continue
		}
		lianYuRank := newLianYuRank(shaQiRank.serviceId, shaQiRank.playerId, shaQiRank.name, shaQiRank.shaqi)
		lianYuRanList = append(lianYuRanList, lianYuRank)
	}
	return
}

func (sd *lianYuSceneData) refreshRank(pl scene.Player, itemId int32, itemNum int32) {
	rankItemId := constanttypes.ShaQiItem
	if sd.activityType == activitytypes.ActivityTypeLianYu {
		rankItemId = constanttypes.ShaLuXin
	}
	if itemId == int32(rankItemId) {
		oldShaQiRankList := sd.getShaQiRankIdList()
		lianYuShaQi, exist := sd.lianYuShaQiMap[pl.GetId()]
		if !exist {
			lianYuShaQi = newLianYuRank(pl.GetServerId(), pl.GetId(), pl.GetName(), itemNum)
			sd.lianYuShaQiMap[pl.GetId()] = lianYuShaQi
		} else {
			lianYuShaQi.shaqi += itemNum
		}
		gameevent.Emit(lianyueventtypes.EventTypeLianYuPlayerShaQiChanged, pl, lianYuShaQi.shaqi)

		isAppend := true
		for _, lianYuRank := range sd.lianYuRankList {
			if lianYuRank.playerId == pl.GetId() {
				isAppend = false
				break
			}
		}
		if isAppend {
			sd.lianYuRankList = append(sd.lianYuRankList, lianYuShaQi)
		}
		sort.Sort(sort.Reverse(LianYuRanktList(sd.lianYuRankList)))
		if len(sd.lianYuRankList) > lianyutypes.ShaQiRankSize {
			sd.lianYuRankList = sd.lianYuRankList[:lianyutypes.ShaQiRankSize]
		}

		isRefresh := false
		if len(oldShaQiRankList) != len(sd.lianYuRankList) {
			isRefresh = true
		}
		if !isRefresh {
			for i := int(0); i < len(oldShaQiRankList); i++ {
				if oldShaQiRankList[i].playerId != sd.lianYuRankList[i].playerId ||
					oldShaQiRankList[i].shaqi != sd.lianYuRankList[i].shaqi {
					isRefresh = true
					break
				}
			}
		}
		if isRefresh {
			gameevent.Emit(lianyueventtypes.EventTypeLianYuShaQiRankChanged, sd, nil)
		}
	}
}
