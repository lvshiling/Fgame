package scene

import (
	"fgame/fgame/core/heartbeat"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	marryeventtypes "fgame/fgame/game/marry/event/types"
	marrytemplate "fgame/fgame/game/marry/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	"fmt"
	"sort"
)

type MarrySceneData interface {
	scene.SceneDelegate
	FinishClearData()
	//本期婚宴结束
	OnBanquetEnd()
	//获取婚礼清场时间
	GetClearTime() int64
	//重置婚礼清场时间
	ResetClearTime()
	//获取夫妻数据
	GetMarryData() *MarryData
	//获取婚期贺礼
	GetGift() *MarrySceneGift
	//当前婚期
	GetPeriod() (period int32)
	//当前婚宴状态
	GetStatus() MarrySceneStatusType
	//获取豪气榜
	GetHeroismList() []*MarryHeroism
	//获取夫妻名字
	GetBothName() (playerId int64, name string, spouseId int64, spouseName string)
	//赠送贺礼
	GiveGift(pl player.Player, period int32, itemId int32, num int32, silver int64, heriosm int64)
	//婚礼开始
	OnWeddingBegin(eventData *marryeventtypes.MarryWedStartEventData)
	//玩家修改名字
	OnWeddingPlayerNameChanged(pl player.Player)
}

type marrySceneData struct {
	*scene.SceneDelegateBase
	s    scene.Scene
	data *MarryData
	gift *MarrySceneGift
	npc  map[int64]scene.NPC
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
	//清场时间
	clearTime int64
}

func CreateMarrySceneData() MarrySceneData {
	marryData := CreateMarryData(-1, MarrySceneStatusTypeInit)
	scenegift := CreateMarrySceneGift(0, 0, 0, 0)
	msd := &marrySceneData{
		data:      marryData,
		gift:      scenegift,
		npc:       make(map[int64]scene.NPC),
		clearTime: 0,
	}
	msd.SceneDelegateBase = scene.NewSceneDelegateBase()

	return msd
}

func (sd *marrySceneData) GetScene() scene.Scene {
	return sd.s
}

//场景开始
func (sd *marrySceneData) OnSceneStart(s scene.Scene) {
	sd.s = s

	//心跳任务
	sd.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	sd.heartbeatRunner.AddTask(CreateMarrySceneTask(sd))
}

//刷怪
func (sd *marrySceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//心跳
func (sd *marrySceneData) Heartbeat() {
	sd.heartbeatRunner.Heartbeat()
}

//场景心跳
func (sd *marrySceneData) OnSceneTick(s scene.Scene) {
	sd.Heartbeat()
}

//生物进入
func (sd *marrySceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	switch npcType {
	case scenetypes.BiologyScriptTypeWedBanquet:
		{
			sd.data.Status = MarrySceneStatusBanquet
			sd.npc[npc.GetId()] = npc
			break
		}
	}

}

func (sd *marrySceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

func (sd *marrySceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *marrySceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("marry:结婚应该是同一个场景"))
	}

}

//怪物死亡
func (sd *marrySceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//玩家重生
func (sd *marrySceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {

}

//玩家死亡
func (sd *marrySceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("marry:结婚应该是同一个场景"))
	}
}
func (sd *marrySceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *marrySceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("marry:结婚应该是同一个场景"))
	}
	//发送事件
	gameevent.Emit(marryeventtypes.EventTypePlayerEnterMarryScene, p, sd)
}

//玩家退出
func (sd *marrySceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("marry:结婚应该是同一个场景"))
	}

}

//场景完成
func (sd *marrySceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("marry:结婚应该是同一个场景"))
	}

}

//场景退出了
func (sd *marrySceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("marry:结婚应该是同一个场景"))
	}
}

//场景获取物品
func (sd *marrySceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("marry:结婚应该是同一个场景"))
	}
}

//玩家获得经验
func (sd *marrySceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("marry:结婚应该是同一个场景"))
	}
}

func (sd *marrySceneData) OnWeddingBegin(eventData *marryeventtypes.MarryWedStartEventData) {
	period := eventData.GetPeriod()
	playerId := eventData.GetPlayerId()
	playerName := eventData.GetPlayerName()
	playerRole := eventData.GetPlayerRole()
	playerSex := eventData.GetPlayerSex()
	spouseName := eventData.GetSpouseName()
	spouseId := eventData.GetSpouseId()
	spouseRole := eventData.GetSpouseRole()
	spouseSex := eventData.GetSpouseSex()
	grade := eventData.GetGrade()
	hunCheGrade := eventData.GetHunCheGrade()
	sugarGrade := eventData.GetSugarGrade()
	sd.data.Period = period
	sd.data.PlayerId = playerId
	sd.data.PlayerName = playerName
	sd.data.PlayerRole = int32(playerRole)
	sd.data.PlayerSex = int32(playerSex)
	sd.data.SpouseId = spouseId
	sd.data.SpouseName = spouseName
	sd.data.SpouseRole = int32(spouseRole)
	sd.data.SpouseSex = int32(spouseSex)
	sd.data.Grade = marrytypes.MarryBanquetSubTypeWed(grade)
	sd.data.HunCheGrade = marrytypes.MarryBanquetSubTypeHunChe(hunCheGrade)
	sd.data.SugarGrade = marrytypes.MarryBanquetSubTypeSugar(sugarGrade)
	sd.gift.PlayerId = playerId
	sd.gift.SpouseId = spouseId

	now := global.GetGame().GetTimeService().Now()
	wedBeginTime, _ := marrytemplate.GetMarryTemplateService().GetMarryFisrtWedTime(now)
	qingChangTime := marrytemplate.GetMarryTemplateService().GetMarryQingChangTime()
	durationTime := marrytemplate.GetMarryTemplateService().GetMarryDurationTime()
	wedEndTime := wedBeginTime + int64(period-1)*int64(qingChangTime+durationTime) + durationTime
	sd.clearTime = wedEndTime + EndDelayTime

	//发送事件
	gameevent.Emit(marryeventtypes.EventTypeMarrySceneWedBegin, sd, nil)
}

func (sd *marrySceneData) OnBanquetEnd() {
	for npcId, npc := range sd.npc {
		sd.s.RemoveSceneObject(npc, false)
		delete(sd.npc, npcId)
	}

	//发送事件
	gameevent.Emit(marryeventtypes.EventTypeMarrySceneWedEnd, sd, nil)
}

func (sd *marrySceneData) FinishClearData() {
	sd.data.Period = -1
	sd.data.PlayerId = 0
	sd.data.PlayerName = ""
	sd.data.SpouseId = 0
	sd.data.SpouseName = ""
	sd.data.Grade = 0
	sd.data.HunCheGrade = 0
	sd.data.SugarGrade = 0
	sd.gift.PlayerId = 0
	sd.gift.SpouseId = 0
	sd.gift.ItemMap = make(map[int32]int32)
	sd.gift.Silver = 0
	sd.data.Status = MarrySceneStatusTypeInit
	sd.data.HeroismList = make([]*MarryHeroism, 0, marrytypes.HeroisListLen)
}

func (sd *marrySceneData) GetPeriod() (period int32) {
	return sd.data.Period
}

func (sd *marrySceneData) GetMarryData() *MarryData {
	return sd.data
}

func (sd *marrySceneData) GetBothName() (playerId int64, name string, spouseId int64, spouseName string) {
	return sd.data.PlayerId, sd.data.PlayerName, sd.data.SpouseId, sd.data.SpouseName
}

func (sd *marrySceneData) GetHeroismList() []*MarryHeroism {
	return sd.data.HeroismList
}

func (sd *marrySceneData) GetGift() *MarrySceneGift {
	return sd.gift
}

func (sd *marrySceneData) GetStatus() MarrySceneStatusType {
	return sd.data.Status
}

func (sd *marrySceneData) GetClearTime() int64 {
	return sd.clearTime
}

func (sd *marrySceneData) ResetClearTime() {
	sd.clearTime = 0
}

func (sd *marrySceneData) GiveGift(pl player.Player, period int32, itemId int32, num int32, silver int64, heriosm int64) {
	if sd.data.Period != period || sd.data.Status != MarrySceneStatusBanquet {
		return
	}

	if itemId != 0 {
		curNum, exist := sd.gift.ItemMap[itemId]
		if exist {
			num += curNum
		}
		sd.gift.ItemMap[itemId] = num
	}

	sd.gift.Silver += silver
	changeFlag := sd.heroismRankChange(heriosm, pl)
	if changeFlag {
		//发送事件
		gameevent.Emit(marryeventtypes.EventTypeMarryHeriosmChange, sd, nil)
	}

}

//TODO: 修改场景通用排行榜
func (sd *marrySceneData) heroismRankChange(heriosm int64, pl player.Player) (changeFlag bool) {
	oldMarryHeroisList := make([]MarryHeroism, 0, marrytypes.HeroisTopLen)
	newPlayerIdList := make([]int64, 0, marrytypes.HeroisTopLen)

	//更新
	isRanking := false
	for index, heroismData := range sd.data.HeroismList {
		if index < marrytypes.HeroisTopLen {
			d := MarryHeroism{
				PlayerId: heroismData.PlayerId,
				Name:     heroismData.Name,
				Heroism:  heroismData.Heroism,
			}
			oldMarryHeroisList = append(oldMarryHeroisList, d)
		}
		if heroismData.PlayerId == pl.GetId() {
			heroismData.Heroism += heriosm
			isRanking = true
		}
	}

	if !isRanking {
		if len(sd.data.HeroismList) >= marrytypes.HeroisListLen {
			lastHerois := sd.data.HeroismList[marrytypes.HeroisListLen-1].Heroism
			if lastHerois > heriosm {
				changeFlag = false
				return
			}
		}

		newRankData := &MarryHeroism{}
		newRankData.Heroism = heriosm
		newRankData.Name = pl.GetName()
		newRankData.PlayerId = pl.GetId()
		sd.data.HeroismList = append(sd.data.HeroismList, newRankData)
	}

	sort.Sort(sort.Reverse(MarryHeroismList(sd.data.HeroismList)))
	if len(sd.data.HeroismList) > marrytypes.HeroisListLen {
		sd.data.HeroismList = sd.data.HeroismList[:marrytypes.HeroisListLen]
	}

	//前三是否变化
	for index, heriosm := range sd.data.HeroismList {
		if index < marrytypes.HeroisTopLen {
			newPlayerIdList = append(newPlayerIdList, heriosm.PlayerId)
		}
	}

	if len(oldMarryHeroisList) != len(newPlayerIdList) {
		changeFlag = true
		return
	}
	for i := int32(0); i < int32(len(oldMarryHeroisList)); i++ {
		if oldMarryHeroisList[i].PlayerId != newPlayerIdList[i] {
			changeFlag = true
			return
		}
		if oldMarryHeroisList[i].Heroism != sd.data.HeroismList[i].Heroism {
			changeFlag = true
			return
		}
	}
	return
}

//玩家修改名字
func (sd *marrySceneData) OnWeddingPlayerNameChanged(pl player.Player) {
	changed := false
	if sd.data.PlayerId == pl.GetId() {
		changed = true
		sd.data.PlayerName = pl.GetName()
	}

	if sd.data.SpouseId == pl.GetId() {
		changed = true
		sd.data.SpouseName = pl.GetName()
	}

	for index, heroismData := range sd.data.HeroismList {
		if heroismData.PlayerId == pl.GetId() {
			heroismData.Name = pl.GetName()
			if index < marrytypes.HeroisTopLen {
				changed = true
			}
			break
		}
	}

	if changed {
		gameevent.Emit(marryeventtypes.EventTypeMarryScenePlayerNameChanged, sd, nil)
	}
}
