package scene

import (
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancetemplate "fgame/fgame/game/alliance/template"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
	"sort"
)

const (
	rankSize = 3
)

type damageInfo struct {
	playerId   int64
	playerName string
	damage     int64
}

func newDamageInfo(playerId int64, playerName string, damage int64) *damageInfo {
	d := &damageInfo{
		playerId:   playerId,
		playerName: playerName,
		damage:     damage,
	}
	return d
}

func (d *damageInfo) GetPlayerId() int64 {
	return d.playerId
}

func (d *damageInfo) GetPlayerName() string {
	return d.playerName
}

func (d *damageInfo) GetDamage() int64 {
	return d.damage
}

type damageInfoList []*damageInfo

func (d damageInfoList) Len() int {
	return len(d)
}

func (d damageInfoList) Less(i, j int) bool {
	return d[i].damage < d[j].damage
}

func (d damageInfoList) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

//仙盟boss场景
func CreateAllianceBossScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeAllianceBoss {
		return nil
	}
	s = scene.CreateScene(mapTemplate, endTime, sh)
	return s
}

type AllianceBossSceneData interface {
	scene.SceneDelegate
	GetAllianceId() int64
	GetSummonTime() int64
	GetLevel() int32
	GetBossNpc() scene.NPC
	GetRankTop() []*damageInfo
	UpdateDamage(playerId int64, damage int64)
}

//仙盟boss数据
type allianceBossSceneData struct {
	*scene.SceneDelegateBase
	s          scene.Scene
	summonTime int64
	allianceId int64
	level      int32
	bossNpc    scene.NPC
	damageList []*damageInfo
}

func CreateAllianceBossSceneData(allianceId int64, level int32) AllianceBossSceneData {
	now := global.GetGame().GetTimeService().Now()
	csd := &allianceBossSceneData{
		summonTime: now,
		allianceId: allianceId,
		level:      level,
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

func (sd *allianceBossSceneData) GetScene() (s scene.Scene) {
	return sd.s
}

//场景开始
func (sd *allianceBossSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	allianceBossTemplate := alliancetemplate.GetAllianceTemplateService().GetAllianceBossTemplate(sd.level)
	if allianceBossTemplate == nil {
		panic(fmt.Errorf("allianceBoss:仙盟boss等级应该是有效的"))
	}
	biologyTemplate := allianceBossTemplate.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}

	n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, int64(0), 0, int32(0), biologyTemplate, allianceBossTemplate.GetPos(), 0, 0)
	if n != nil {
		//设置场景
		sd.s.AddSceneObject(n)
		sd.bossNpc = n
	}
}

//刷怪
func (sd *allianceBossSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *allianceBossSceneData) OnSceneTick(s scene.Scene) {
	sd.Heartbeat()
}

//怪物死亡
func (sd *allianceBossSceneData) OnSceneBiologyAllDead(s scene.Scene) {
	sd.s.Finish(true)
}

//生物进入
func (sd *allianceBossSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {

}

func (sd *allianceBossSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *allianceBossSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}

}

//生物重生
func (sd *allianceBossSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}
}

//玩家复活
func (sd *allianceBossSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}

}

//玩家死亡
func (sd *allianceBossSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}
}

func (sd *allianceBossSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *allianceBossSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}
	gameevent.Emit(allianceeventtypes.EventTypePlayerEnterAllianceBossScene, sd, p)
}

//玩家退出
func (sd *allianceBossSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}

}

//场景完成
func (sd *allianceBossSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}
	gameevent.Emit(allianceeventtypes.EventTypeAllianceBossSceneFinish, sd, success)
}

//场景退出了
func (sd *allianceBossSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}
}

//场景获取物品
func (sd *allianceBossSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}
}

//玩家获得经验
func (sd *allianceBossSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("allianceBoss:仙盟boss应该是同一个场景"))
	}
}

//心跳
func (sd *allianceBossSceneData) Heartbeat() {

}

func (sd *allianceBossSceneData) GetSummonTime() int64 {
	return sd.summonTime
}

func (sd *allianceBossSceneData) GetAllianceId() int64 {
	return sd.allianceId
}

func (sd *allianceBossSceneData) GetBossNpc() scene.NPC {
	return sd.bossNpc
}

func (sd *allianceBossSceneData) GetLevel() int32 {
	return sd.level
}

func (sd *allianceBossSceneData) GetRankTop() []*damageInfo {
	rankLen := len(sd.damageList)
	if rankLen == 0 {
		return nil
	}
	addLen := 0
	if rankSize >= rankLen {
		addLen = rankLen
	}
	return sd.damageList[0:addLen]
}

func (sd *allianceBossSceneData) addDamageInfo(pl scene.Player, damage int64) {
	if damage <= 0 {
		return
	}
	for _, damageInfo := range sd.damageList {
		if damageInfo.GetPlayerId() == pl.GetId() {
			damageInfo.damage += damage
			return
		}
	}
	damageInfo := newDamageInfo(pl.GetId(), pl.GetName(), damage)
	sd.damageList = append(sd.damageList, damageInfo)
}

func (sd *allianceBossSceneData) getOldRankTop() (oldList []*damageInfo) {
	oldDamageTopList := sd.GetRankTop()
	for _, oldDamageInfo := range oldDamageTopList {
		playerId := oldDamageInfo.GetPlayerId()
		playerName := oldDamageInfo.GetPlayerName()
		damage := oldDamageInfo.GetDamage()
		damageInfo := newDamageInfo(playerId, playerName, damage)
		oldList = append(oldList, damageInfo)
	}
	return
}

//更新伤害
func (sd *allianceBossSceneData) UpdateDamage(playerId int64, damage int64) {
	if damage <= 0 {
		return
	}
	pl := sd.GetScene().GetPlayer(playerId)
	if pl == nil {
		return
	}
	oldDamageTopList := sd.getOldRankTop()

	sd.addDamageInfo(pl, damage)
	sort.Sort(sort.Reverse(damageInfoList(sd.damageList)))
	newDamageTopList := sd.GetRankTop()

	pushFlag := false
	if len(oldDamageTopList) != len(newDamageTopList) {
		pushFlag = true
	} else {
		topLen := len(oldDamageTopList)
		for i := 0; i < topLen; i++ {
			if oldDamageTopList[i].GetPlayerId() != newDamageTopList[i].GetPlayerId() ||
				oldDamageTopList[i].GetDamage() != newDamageTopList[i].GetDamage() {
				pushFlag = true
				break
			}
		}
	}
	if pushFlag {
		gameevent.Emit(allianceeventtypes.EventTypeAllianceBossRankChanged, sd, nil)
	}
}
