package scene

import (
	"fgame/fgame/core/heartbeat"
	tulongeventtypes "fgame/fgame/cross/tulong/event/types"
	crosstulongtypes "fgame/fgame/cross/tulong/types"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/scene/scene"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	tulongtemplate "fgame/fgame/game/tulong/template"
	tulongtypes "fgame/fgame/game/tulong/types"
	"fmt"
)

//屠龙战场场景
func CreateTuLongScene(mapId int32, endTime int64, sh scene.SceneDelegate) (s scene.Scene) {
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	if mapTemplate.GetMapType() != scenetypes.SceneTypeCrossTuLong {
		return nil
	}
	s = scene.CreateScene(mapTemplate, endTime, sh)
	return s
}

//采集龙蛋信息
type collectEggInfo struct {
	playerId           int64                          //采集龙蛋玩家
	allianceId         int64                          //仙盟id
	collectStarTime    int64                          //采集开始时间
	pos                int32                          //龙蛋pos
	bornBiaoShi        int32                          //龙蛋出生标识
	status             crosstulongtypes.EggStatusType //龙蛋状态
	npc                scene.NPC                      //场景npc
	lastCollectEndTime int64                          //上次采集结束时间
}

func newCollectEggInfo(pos int32, bornBiaoShi int32) *collectEggInfo {
	collectEggInfo := &collectEggInfo{
		playerId:           0,
		allianceId:         0,
		collectStarTime:    0,
		pos:                pos,
		bornBiaoShi:        bornBiaoShi,
		status:             crosstulongtypes.EggStatusTypeInit,
		npc:                nil,
		lastCollectEndTime: 0,
	}
	return collectEggInfo
}

func (ce *collectEggInfo) GetPlayerId() int64 {
	return ce.playerId
}

func (ce *collectEggInfo) GetEggNpc() scene.NPC {
	return ce.npc
}

//大龙蛋信息
type bigEggInfo struct {
	bornBiaoShi int32                          //出生标识
	status      crosstulongtypes.EggStatusType //龙蛋状态
	npc         scene.NPC                      //场景npc
}

func newBigEggInfo(bigEggBornBiaoShi int32) *bigEggInfo {
	bigEgg := &bigEggInfo{
		bornBiaoShi: bigEggBornBiaoShi,
		status:      crosstulongtypes.EggStatusTypeInit,
	}
	return bigEgg
}

func (be *bigEggInfo) GetStatus() crosstulongtypes.EggStatusType {
	return be.status
}

func (be *bigEggInfo) GetBornBiaoShi() int32 {
	return be.bornBiaoShi
}

//仙盟信息
type allianceInfo struct {
	//仙盟人数
	num int32
	//成功采集过龙蛋
	collected bool
	//击杀boss数量
	killNum int32
	//仙盟成员获得物品
	itemInfoMap map[int64]map[int32]int32
}

func newAllianceInfo() *allianceInfo {
	info := &allianceInfo{
		num:         0,
		collected:   false,
		killNum:     0,
		itemInfoMap: make(map[int64]map[int32]int32),
	}
	return info
}

func (a *allianceInfo) GetKillNum() int32 {
	return a.killNum
}

func (a *allianceInfo) GetItemMap(playerId int64) map[int32]int32 {
	return a.itemInfoMap[playerId]
}

type TuLongSceneData interface {
	scene.SceneDelegate
	//获取当前仙盟人数
	GetAllianceNum(allianceId int64) int32
	//获取大龙蛋
	GetBigEgg() *bigEggInfo
	//获取采集龙蛋
	GetCollectEgg() map[int64]*collectEggInfo
	//是否已成功采集过龙蛋
	HasedCollectEgg(allianceId int64) bool
	//龙蛋应该存在
	SmallEggShouldExist(npcId int64) (eggInfo *collectEggInfo, flag bool)
	//能否采集
	IfCanCollectEgg(npcId int64) (*collectEggInfo, bool)
	//盟主采集龙蛋
	CollectEgg(playerId int64, allianceId int64, npcId int64) bool
	//玩家是否正在采集龙蛋
	HasingCollectEgg(playerId int64) (npcId int64, has bool)
	//采集被打断
	CollectEggInterrupt(npcId int64)
	//采集完成
	FinishCollectEgg(npcId int64)
	//击杀boss
	KillBoss(pl scene.Player)
	//获取allianceMap
	GetAllianceMap() map[int64]*allianceInfo
	//小龙蛋刷新判断
	SmallEggRebornCheck()
	// 是否排队
	IfLineup() bool
}

//屠龙战场数据
type tuLongSceneData struct {
	*scene.SceneDelegateBase
	s scene.Scene
	//开始时间
	starTime int64
	//采集龙蛋
	collectEggMap map[int64]*collectEggInfo
	//上次小龙蛋刷新时间
	lastCollectEggRefreshTime int64
	//大龙蛋
	bigEgg *bigEggInfo
	//仙盟map
	allianceMap map[int64]*allianceInfo
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
}

func CreateTuLongSceneData(bigEggBornBiaoShi int32) TuLongSceneData {
	csd := &tuLongSceneData{
		collectEggMap: make(map[int64]*collectEggInfo),
		bigEgg:        newBigEggInfo(bigEggBornBiaoShi),
		allianceMap:   make(map[int64]*allianceInfo),
	}
	csd.SceneDelegateBase = scene.NewSceneDelegateBase()
	return csd
}

func (sd *tuLongSceneData) GetScene() (s scene.Scene) {
	return sd.s
}

//场景开始
func (sd *tuLongSceneData) OnSceneStart(s scene.Scene) {
	sd.s = s
	now := global.GetGame().GetTimeService().Now()
	activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeCoressTuLong)
	activityTimeTemplate, _ := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
	startTime, _ := activityTimeTemplate.GetBeginTime(now)
	sd.starTime = startTime
	sd.lastCollectEggRefreshTime = startTime

	sd.refreshLongDan()
	//心跳任务
	sd.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	sd.heartbeatRunner.AddTask(CreateTuLongTask(sd, sd.starTime))
}

//刷怪
func (sd *tuLongSceneData) OnSceneRefreshGroup(s scene.Scene, currentGroup int32) {

}

//场景心跳
func (sd *tuLongSceneData) OnSceneTick(s scene.Scene) {
	sd.Heartbeat()
}

//怪物死亡
func (sd *tuLongSceneData) OnSceneBiologyAllDead(s scene.Scene) {

}

//生物进入
func (sd *tuLongSceneData) OnSceneBiologyEnter(s scene.Scene, npc scene.NPC) {
	//大Boss刷新
	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	if npcType == scenetypes.BiologyScriptTypeCrossBigBoss {
		//发送事件
		gameevent.Emit(tulongeventtypes.EventTypeTuLongBigEggStatusRefresh, sd, nil)
	}
}

func (sd *tuLongSceneData) OnSceneBiologyExit(s scene.Scene, npc scene.NPC) {

}

//怪物死亡
func (sd *tuLongSceneData) OnSceneBiologyDead(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
	//大Boss死亡
	npcType := npc.GetBiologyTemplate().GetBiologyScriptType()
	switch npcType {
	case scenetypes.BiologyScriptTypeCrossBigBoss:
		{

			sd.bigEgg.status = crosstulongtypes.EggStatusTypeDead
			//发送事件
			gameevent.Emit(tulongeventtypes.EventTypeTuLongBigEggStatusRefresh, sd, nil)
			break
		}
	case scenetypes.BiologyScriptTypeCrossSmallBoss:
		{
			collectEggInfo := sd.collectEggMap[npc.GetId()]
			if collectEggInfo != nil {
				collectEggInfo.status = crosstulongtypes.EggStatusTypeDead
			}
		}
	}
}

//生物重生
func (sd *tuLongSceneData) OnSceneBiologyReborn(s scene.Scene, npc scene.NPC) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
}

//玩家复活
func (sd *tuLongSceneData) OnScenePlayerReborn(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}

}

//玩家死亡
func (sd *tuLongSceneData) OnScenePlayerDead(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
}
func (sd *tuLongSceneData) OnScenePlayerBeforeEnter(s scene.Scene, p scene.Player) {

}

//玩家进入
func (sd *tuLongSceneData) OnScenePlayerEnter(s scene.Scene, p scene.Player) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
	p.SwitchPkState(pktypes.PkStateBangPai, pktypes.PkCommonCampDefault)
	sd.addAllianceNum(p.GetAllianceId())
	//发送事件
	gameevent.Emit(tulongeventtypes.EventTypeTuLongPlayerEnter, p, sd)

}

//玩家退出
func (sd *tuLongSceneData) OnScenePlayerExit(s scene.Scene, p scene.Player, active bool) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
	//清空采集龙蛋
	if p.GetMengZhuId() == p.GetId() {
		pos, hasCollect := sd.HasingCollectEgg(p.GetAllianceId())
		if hasCollect {
			sd.CollectEggInterrupt(pos)
		}
	}

	sd.subAllianceNum(p.GetAllianceId())
}

//场景完成
func (sd *tuLongSceneData) OnSceneFinish(s scene.Scene, success bool, useTime int64) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
	gameevent.Emit(tulongeventtypes.EventTypeTuLongSceneFinish, sd, nil)
}

//场景退出了
func (sd *tuLongSceneData) OnSceneStop(s scene.Scene) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
}

//场景获取物品
func (sd *tuLongSceneData) OnScenePlayerGetItem(s scene.Scene, pl scene.Player, itemData *droptemplate.DropItemData) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
	itemId := itemData.ItemId
	itemNum := itemData.Num
	allianceId := pl.GetAllianceId()
	playerId := pl.GetId()
	allianceInfo, exist := sd.allianceMap[allianceId]
	if !exist {
		return
	}
	itemMap, exist := allianceInfo.itemInfoMap[playerId]
	if !exist {
		itemMap = make(map[int32]int32)
		allianceInfo.itemInfoMap[playerId] = itemMap
	}
	itemMap[itemId] += itemNum
}

//玩家获得经验
func (sd *tuLongSceneData) OnScenePlayerGetExp(s scene.Scene, p scene.Player, num int64) {
	if sd.s != s {
		panic(fmt.Errorf("tulong:屠龙战场应该是同一个场景"))
	}
}

//心跳
func (sd *tuLongSceneData) Heartbeat() {
	sd.heartbeatRunner.Heartbeat()
}

//刷龙蛋
func (sd *tuLongSceneData) refreshLongDan() {
	tuLongConstTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongConstTemplate()
	if tuLongConstTemplate == nil {
		return
	}
	//刷大龙蛋
	bigEggBiologyTemplate := tuLongConstTemplate.GetBigEggBiologyTemplate()
	if bigEggBiologyTemplate != nil {
		tuLongPosTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongPosTemplate(tulongtypes.TuLongPosTypeBoss, sd.bigEgg.bornBiaoShi)
		n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, int64(0), 0, int32(0), bigEggBiologyTemplate, tuLongPosTemplate.GetPos(), 0, 0)
		if n != nil {
			//设置场景
			sd.s.AddSceneObject(n)
			sd.bigEgg.npc = n
		}
	}

	var biaoShiList []int32
	biaoShiList = append(biaoShiList, sd.bigEgg.bornBiaoShi)
	//刷小龙蛋
	smallEggBiologyTemplate := tuLongConstTemplate.GetSmallEggBiologyTemplate()
	tuLongLen := tulongtemplate.GetTuLongTemplateService().GetTuLongLen()
	for i := int32(1); i < tuLongLen; i++ {
		biaoShi, flag := tulongtemplate.GetTuLongTemplateService().GetTuLongPosBiaoShi(tulongtypes.TuLongPosTypeBoss, biaoShiList)
		if !flag {
			continue
		}
		//小龙蛋标识
		collectEggInfo := newCollectEggInfo(i, biaoShi)
		tuLongPosTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongPosTemplate(tulongtypes.TuLongPosTypeBoss, biaoShi)
		n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, int64(0), 0, int32(0), smallEggBiologyTemplate, tuLongPosTemplate.GetPos(), 0, 0)
		if n != nil {
			//设置场景
			sd.s.AddSceneObject(n)
			collectEggInfo.npc = n
		}
		sd.collectEggMap[n.GetId()] = collectEggInfo
		biaoShiList = append(biaoShiList, biaoShi)
	}

}

//获取大龙蛋
func (sd *tuLongSceneData) GetBigEgg() *bigEggInfo {
	return sd.bigEgg
}

//是否已成功采集过龙蛋
func (sd *tuLongSceneData) HasedCollectEgg(allianceId int64) bool {
	allianceInfo, exist := sd.allianceMap[allianceId]
	if !exist {
		return false
	}
	return allianceInfo.collected == true
}

func (sd *tuLongSceneData) SmallEggShouldExist(npcId int64) (eggInfo *collectEggInfo, flag bool) {
	eggInfo, flag = sd.collectEggMap[npcId]
	if !flag {
		return
	}
	flag = true
	return
}

//龙蛋是否正在采集
func (sd *tuLongSceneData) IfCanCollectEgg(npcId int64) (*collectEggInfo, bool) {
	collectEggInfo, exist := sd.SmallEggShouldExist(npcId)
	if !exist {
		return nil, false
	}
	if collectEggInfo.status != crosstulongtypes.EggStatusTypeInit {
		return collectEggInfo, false
	}
	if collectEggInfo.playerId != 0 {
		return collectEggInfo, false
	}
	return collectEggInfo, true
}

//玩家是否正在采集
func (sd *tuLongSceneData) HasingCollectEgg(playerId int64) (npcId int64, has bool) {
	npcId = 0
	has = false
	for curNpcId, collectEggInfo := range sd.collectEggMap {
		if collectEggInfo.playerId == playerId {
			has = true
			npcId = curNpcId
			return
		}
	}
	return
}

//获取采集龙蛋
func (sd *tuLongSceneData) GetCollectEgg() map[int64]*collectEggInfo {
	return sd.collectEggMap
}

//盟主采集龙蛋
func (sd *tuLongSceneData) CollectEgg(playerId int64, allianceId int64, npcId int64) bool {
	now := global.GetGame().GetTimeService().Now()
	collectEggInfo, flag := sd.IfCanCollectEgg(npcId)
	if !flag {
		return false
	}
	collectEggInfo.playerId = playerId
	collectEggInfo.allianceId = allianceId
	collectEggInfo.collectStarTime = now
	return true
}

//采集被打断
func (sd *tuLongSceneData) CollectEggInterrupt(npcId int64) {
	collectEggInfo, exist := sd.collectEggMap[npcId]
	if !exist {
		return
	}
	collectEggInfo.playerId = 0
	collectEggInfo.allianceId = 0
	collectEggInfo.collectStarTime = 0
}

//采集完成
func (sd *tuLongSceneData) FinishCollectEgg(npcId int64) {
	now := global.GetGame().GetTimeService().Now()
	collectEggInfo, exist := sd.collectEggMap[npcId]
	if !exist {
		return
	}
	if collectEggInfo.status != crosstulongtypes.EggStatusTypeInit {
		return
	}
	allianceId := collectEggInfo.allianceId
	allianceInfo, exist := sd.allianceMap[allianceId]
	if !exist {
		return
	}
	allianceInfo.collected = true
	collectEggInfo.status = crosstulongtypes.EggStatusTypeBoss
	collectEggInfo.lastCollectEndTime = now
	//移除小龙蛋
	sd.s.RemoveSceneObject(collectEggInfo.npc, false)
	collectEggInfo.npc = nil

	//刷小Boss
	tuLongTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongSmallBossTemplate(collectEggInfo.pos)
	if tuLongTemplate == nil {
		return
	}
	biologyTemplate := tuLongTemplate.GetBiologyTemplate()
	if biologyTemplate == nil {
		return
	}

	tuLongPosTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongPosTemplate(tulongtypes.TuLongPosTypeBoss, collectEggInfo.bornBiaoShi)
	n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, int64(0), 0, int32(0), biologyTemplate, tuLongPosTemplate.GetPos(), 0, 0)
	if n != nil {
		//设置场景
		sd.s.AddSceneObject(n)
		delete(sd.collectEggMap, npcId)
		collectEggInfo.npc = n
		sd.collectEggMap[n.GetId()] = collectEggInfo
	}
	//发送事件
	gameevent.Emit(tulongeventtypes.EventTypeTuLongCollectFinish, sd, npcId)
}

func (sd *tuLongSceneData) GetAllianceNum(allianceId int64) (num int32) {
	allianceInfo, exist := sd.allianceMap[allianceId]
	if !exist {
		return
	}
	return allianceInfo.num
}

//仙盟人数加1
func (sd *tuLongSceneData) addAllianceNum(allianceId int64) {
	allianceInfo, exist := sd.allianceMap[allianceId]
	if !exist {
		allianceInfo = newAllianceInfo()
		allianceInfo.num++
		sd.allianceMap[allianceId] = allianceInfo
		return
	}
	allianceInfo.num++
}

//仙盟人数减1
func (sd *tuLongSceneData) subAllianceNum(allianceId int64) {
	allianceInfo, exist := sd.allianceMap[allianceId]
	if !exist {
		return
	}
	allianceInfo.num--
}

func (sd *tuLongSceneData) KillBoss(pl scene.Player) {
	allianceId := pl.GetAllianceId()
	allianceInfo, exist := sd.allianceMap[allianceId]
	if !exist {
		return
	}
	allianceInfo.killNum++
}

func (sd *tuLongSceneData) GetAllianceMap() map[int64]*allianceInfo {
	return sd.allianceMap
}

//小龙蛋刷新判断
func (sd *tuLongSceneData) SmallEggRebornCheck() {
	tuLongConstTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongConstTemplate()
	now := global.GetGame().GetTimeService().Now()
	diffTime := now - sd.lastCollectEggRefreshTime
	if diffTime < int64(tuLongConstTemplate.SmallEggShuaXin) {
		return
	}

	sd.smallEggReborn()

	intervalCount := diffTime / sd.lastCollectEggRefreshTime
	sd.lastCollectEggRefreshTime += int64(tuLongConstTemplate.SmallEggShuaXin) * intervalCount
}

//是否排队
func (sd *tuLongSceneData) IfLineup() bool {
	plNum := len(sd.GetScene().GetAllPlayers())
	limitNum := tulongtemplate.GetTuLongTemplateService().GetTuLongConstTemplate().PlayerLimitCount
	return plNum >= int(limitNum)
}

func (sd *tuLongSceneData) smallEggReborn() {
	tuLongConstTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongConstTemplate()
	for _, eggInfo := range sd.collectEggMap {
		if eggInfo.status != crosstulongtypes.EggStatusTypeDead {
			continue
		}

		biaoShi := eggInfo.bornBiaoShi
		smallEggBiologyTemplate := tuLongConstTemplate.GetSmallEggBiologyTemplate()
		tuLongPosTemplate := tulongtemplate.GetTuLongTemplateService().GetTuLongPosTemplate(tulongtypes.TuLongPosTypeBoss, biaoShi)
		n := scene.CreateNPC(scenetypes.OwnerTypeNone, 0, int64(0), 0, int32(0), smallEggBiologyTemplate, tuLongPosTemplate.GetPos(), 0, 0)
		if n != nil {
			//设置场景
			sd.s.AddSceneObject(n)
			delete(sd.collectEggMap, eggInfo.npc.GetId())

			eggInfo.playerId = 0
			eggInfo.allianceId = 0
			eggInfo.collectStarTime = 0
			eggInfo.lastCollectEndTime = 0
			eggInfo.status = crosstulongtypes.EggStatusTypeInit
			eggInfo.npc = n
			sd.collectEggMap[n.GetId()] = eggInfo
		}
	}
}
