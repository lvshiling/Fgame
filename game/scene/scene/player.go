package scene

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/aoi"
	"fgame/fgame/core/fsm"
	coretypes "fgame/fgame/core/types"
	activitytypes "fgame/fgame/game/activity/types"
	actvitytypes "fgame/fgame/game/activity/types"
	alliancetypes "fgame/fgame/game/alliance/types"
	cdcommon "fgame/fgame/game/cd/common"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/global"
	godsiegetypes "fgame/fgame/game/godsiege/types"
	itemtypes "fgame/fgame/game/item/types"
	pktypes "fgame/fgame/game/pk/types"
	"fgame/fgame/game/player/types"
	scenetypes "fgame/fgame/game/scene/types"
	teamtypes "fgame/fgame/game/team/types"
	tianshutypes "fgame/fgame/game/tianshu/types"
	worldbosstypes "fgame/fgame/game/worldboss/types"
	"fmt"

	"github.com/golang/protobuf/proto"
)

type GuaJiAction interface {
	OnEnter()
	OnExit()
	GuaJi(Player)
}

type GuaJiActionFactory interface {
	CreateAction() GuaJiAction
}

type GuaJiActionFactoryFunc func() GuaJiAction

func (f GuaJiActionFactoryFunc) CreateAction() GuaJiAction {
	return f()
}

type DummyGuaJiAction struct {
}

func (a *DummyGuaJiAction) OnEnter() {
	return
}

func (a *DummyGuaJiAction) GuaJi(p Player) {
	return
}

func (a *DummyGuaJiAction) OnExit() {
	return
}

func NewDummyGuaJiAction() *DummyGuaJiAction {
	return &DummyGuaJiAction{}
}

var (
	dummyActionInstance = NewDummyGuaJiAction()
)

var (
	guaJiStateMap        = make(map[scenetypes.GuaJiType]map[fsm.State]GuaJiActionFactory)
	defaultGuaJiStateMap = make(map[fsm.State]GuaJiActionFactory)
)

func RegisterDefaultGuaJiActionFactory(state fsm.State, action GuaJiActionFactory) {
	_, ok := defaultGuaJiStateMap[state]
	if ok {
		panic(fmt.Errorf("重复注册默认行为[%d]", state))
	}
	defaultGuaJiStateMap[state] = action
}

func RegisterGuaJiActionFactory(typ scenetypes.GuaJiType, state fsm.State, action GuaJiActionFactory) {
	stateMap, ok := guaJiStateMap[typ]
	if !ok {
		stateMap = make(map[fsm.State]GuaJiActionFactory)
		guaJiStateMap[typ] = stateMap
	}
	_, ok = stateMap[state]
	if ok {
		panic(fmt.Errorf("重复注册%d挂机类型,%d状态", typ, state))
	}
	stateMap[state] = action
}

func getDefaultGuaJiAction(state fsm.State) GuaJiAction {
	stateActionFactory, ok := defaultGuaJiStateMap[state]
	if !ok {
		return dummyActionInstance
	}
	action := stateActionFactory.CreateAction()
	return action
}

func GetGuaJiAction(typ scenetypes.GuaJiType, state fsm.State) GuaJiAction {
	stateMap, ok := guaJiStateMap[typ]
	if !ok {
		return getDefaultGuaJiAction(state)

	}
	stateActionFactory, ok := stateMap[state]
	if !ok {
		return getDefaultGuaJiAction(state)
	}
	action := stateActionFactory.CreateAction()
	return action
}

//挂机管理器
type PlayerGuaJiManager interface {
	IsGuaJi() bool
	StartGuaJi(scenetypes.GuaJiType) bool
	StopGuaJi()
	EnterGuaJi(scenetypes.GuaJiType)
	GetCurrentGuaJiType() scenetypes.GuaJiType
	GetLastGuaJiType() (scenetypes.GuaJiType, bool)
	ExitGuaJi()
	GuaJiIdle() bool
	GuaJiTrace() bool
	GuaJiAttack() bool
	GuaJiRun() bool
	GuaJiDead() bool
	GetCurrentGuaJiAction() GuaJiAction
	GetGuaJiDeadTimes() int32
}

//队伍管理器
type PlayerTeamManager interface {
	//获取队伍id
	GetTeamId() int64
	//获取队伍名字
	GetTeamName() string
	//获取队伍标识
	GetTeamPurpose() teamtypes.TeamPurposeType
	//设置队伍id
	SyncTeam(teamId int64, teamName string, purpose teamtypes.TeamPurposeType)
}

//仙盟管理器
type PlayerAllianceManager interface {
	//获取仙盟id
	GetAllianceId() int64
	//获取队伍名字
	GetAllianceName() string
	//盟主
	GetMengZhuId() int64
	//当前职位
	GetMemPos() alliancetypes.AlliancePosition
	//同步仙盟
	SyncAlliance(allianceId int64, allianceName string, mengZhuId int64, pos alliancetypes.AlliancePosition)
}

//结义管理器
type PlayerJieYiManager interface {
	GetSceneJieYiName() string
	//获取结义名称
	GetJieYiName() string
	GetJieYiId() int64
	GetJieYiRank() int32
	SyncJieYi(jieYiId int64, jieYiName string, jieYiRank int32)
}

//pk管理器
type PlayerPkManager interface {
	SwitchPkState(pktypes.PkState, pktypes.PkCamp) bool
	GetPkState() pktypes.PkState
	GetPkValue() int32
	GetPkCamp() pktypes.PkCamp
	GetPkRedState() pktypes.PkRedState
	GetKillNum() int32
	GetLastKillTime() int64
	GetPkOnlineTime() int64
	GetPkLoginTime() int64
	Kill(white bool)
	ReducePkValue(value int32)
}

//战斗管理器
type PlayerBattleManager interface {
	BattleManager
	IsRobot() bool
	Battle() bool
	IsBattle() bool
	PvpBattle() bool
	IsPvpBattle() bool
	ClearPvpBattle()
	//是否可以退出卡死
	IfCanExitKaSi() bool
	//脱离卡死剩余时间
	ExitKaSiLeftTime() int64
	//退出卡死
	ExitKaSi()
	//可以获取物品
	IfCanGetDropItem(di DropItem) bool
	SyncSoulAwakenNum(soulAwakenNum int32)
	GetSoulAwakenNum() int32
	GetLevel() int32
	SyncLevel(level int32)
	GetZhuanSheng() int32
	SyncZhuanSheng(zhuanSheng int32)
	GetVip() int32
	SyncVip(vip int32)
	IsHuiYuanPlus() bool
	SyncHuiYuan(isHuiYuan bool)
}

//玩家场景管理器
type PlayerSceneManager interface {
	SceneObject
	RemoveLoadedPlayer(id int64)
	GetLoadedPlayers() map[int64]Player
	//获取进入场景邻居和清空
	GetEnterNeighborsAndClear() map[int64]aoi.AOI
	//获取退出场景邻居和清空
	GetLeaveNeighborsAndClear() map[int64]aoi.AOI
	SetEnterPos(pos coretypes.Position)
	//获取当前地图
	GetMapId() int32
	GetSceneId() int64
	GetPos() coretypes.Position
	//获取上一个地图
	GetLastMapId() int32
	GetLastSceneId() int64
	//获取上一个场景地点
	GetLastPos() coretypes.Position
}

//战斗属性管理器
type PlayerBattlePropertyManager interface {
	BattlePropertyManager
	//获取变更的战斗属性
	GetBattlePropertyChangedTypesAndReset() (battleChanged map[int32]int64)
}

//阵营管理器
type PlayerZhenYingManager interface {
	//获取阵营
	GetCamp() chuangshitypes.ChuangShiCampType
	//设置阵营
	SetCamp(chuangshitypes.ChuangShiCampType)
	SetGuanZhi(guanZhi chuangshitypes.ChuangShiGuanZhi)
	GetGuanZhi() chuangshitypes.ChuangShiGuanZhi
	//攻城原地复活次数
	SetChuangShiReliveTimes(times int32)
	GetChuangShiReliveTimes() int32
}

//玩家展示管理器
type PlayerShowManager interface {
	//时装
	GetFashionId() int32
	SetFashionId(fashionId int32)
	//冰魂
	GetWeaponId() int32
	SetWeapon(weaponId int32, weaponState int32)
	//冰魂觉醒
	GetWeaponState() int32
	GetTitleId() int32
	SetTitleId(titleId int32)
	TitleHidden(hidden bool)
	GetWingId() int32
	SetWingId(wingId int32)
	GetMountId() int32
	GetMountAdvanceId() int32
	SetMountId(mountId int32, advanceId int32)
	MountHidden(hidden bool)
	IsMountHidden() bool
	MountSync(hidden bool)
	//获取四神钥匙
	GetFourGodKey() int32
	SetFourGodKey(fourGodKey int32)
	//身法
	GetShenFaId() int32
	SetShenFaId(shenFaId int32)
	//领域
	GetLingYuId() int32
	SetLingYuId(lingYuId int32)
	//天劫塔
	GetRealm() int32
	SetRealm(realm int32)
	//配偶
	GetSpouse() string
	SetSpouse(spouse string)
	SetSpouseId(id int64)
	GetSpouseId() int64
	//婚礼状态
	GetWeddingStatus() int32
	SetWeddingStatus(status int32)
	//设置模型
	SetModel(model int32)
	GetModel() int32
	//获取婚戒
	GetRingType() int32
	SetRingType(ringType int32)
	GetRingLevel() int32
	SetRingLevel(ringLevel int32)
	GetMarryDevelopLevel() int32
	SetMarryDevelopLevel(developLevel int32)
	//法宝id
	GetFaBaoId() int32
	SetFaBaoId(faBaoId int32)
	//宠物id
	GetPetId() int32
	SetPetId(petId int32)
	//仙体
	GetXianTiId() int32
	SetXianTiId(xianTiId int32)
	//八卦秘境
	GetBaGua() int32
	SetBaGua(level int32)
	//飞宠
	GetFlyPetId() int32
	SetFlyPetId(petId int32)
	//获取神域钥匙
	GetShenYuKey() int32
	SetShenYuKey(keyNum int32)
}

//玩家竞技管理器
type PlayerArenaManager interface {
	GetArenaReliveTime() int32
	SetArenaReliveTime(reliveTime int32)
	GetArenaWinTime() int32
	SetArenaWinTime(winTime int32)
	StartArenaBattle()
	StopArenaBattle()
	IsArenaBattle() bool
}

//玩家竞技pvp管理器
type PlayerArenapvpManager interface {
	GetArenapvpReliveTimes() int32
	SetArenapvpReliveTimes(reliveTimes int32)
	StartArenapvpBattle()
	StopArenapvpBattle()
	IsArenapvpBattle() bool
}

type PlayerXueChiManager interface {
	SetBloodLine(bloodLine int32)
	GetBloodLine() int32
	GetBlood() int64
	GetLastBloodTime() int64
	AddBlood(blood int64)
	RecoverHp(recover int64)
	SyncBlood(blood int64, bloodLine int32)
}

type PlayerReliveManager interface {
	RefreshReliveTime() bool
	GetCulReliveTime() int32
	GetLastReliveTime() int64
	Relive()
	SyncRelive(reliveTime int32, lastReliveTime int64)
}

type PlayerCollectManager interface {
	Collect(npc CollectNPC) bool
	HasCollect() (CollectNPC, bool)
	ClearCollect()
}

type PlayerTowerManager interface {
	CountTowerExp(exp int64)
	CountTowerItemMap(itemId, num int32)
	ResetCountTower()
	GetCountTowerExp() int64
	GetCountTowerItemMap() map[int32]int32
	StartDaBao()
	EndDaBao()
	IsOnDabao() bool
	IfNotDaBaoNotice() bool
}

type PlayerTianShuManager interface {
	GetTianShuRate(typ tianshutypes.TianShuType) int32
	AddTianShu(typ tianshutypes.TianShuType, rate int32)
	UplevelTianShu(typ tianshutypes.TianShuType, newRate int32)
}

type PlayerGodSiegeManager interface {
	GetGodSiegeLineUp() (isLineUp bool, godType godsiegetypes.GodSiegeType)
	GodSiegeCancleLineUp() (flag bool)
	GodSiegeLineUp(godType godsiegetypes.GodSiegeType) (flag bool)
	IsGodSiegeLineUp() bool
}

type PlayerLuckyManager interface {
	AddLucky(typ itemtypes.ItemType, subType itemtypes.ItemSubType, rate int32, expireTime int64)
	GetLuckyRate(typ itemtypes.ItemType, subType itemtypes.ItemSubType) int32
}

type PlayerUnrealBossManager interface {
	IsEnoughPilao(pilaoNum int32) bool
	IsPilaoNoticeCd() bool
	SynPilaoNum(pilaoNum int32)
	GetPilao() int32
}

type PlayerOutlandBossManager interface {
	IsZhuoQiLimit() bool
	IsZhuoQiNoticeCd() bool
	SynZhuoQiNum(num int32)
}

type PlayerLingTongManager interface {
	GetLingTong() LingTong
	UpdateLingTong(lingTong LingTong)
	HiddenLingTong(flag bool)
	IsLingTongHidden() bool
}

type PlayerDenseWatManager interface {
	GetDenseWatNum() int32
	SetDenseWatNum(num int32)
	GetDenseWatEndTime() int64
	SetDenseWatEndTime(endTime int64)
	SyncDenseWat(num int32, endTime int64)
}

type PlayerShenMoManager interface {
	GetShenMoGongXunNum() int32
	SetShenMoGongXunNum(num int32)
	GetShenMoKillNum() int32
	SetShenMoKillNum(killNum int32)
	GetShenMoEndTime() int64
	SetShenMoEndTime(endTime int64)
	SyncShenMo(gongXunNum int32, killNum int32, endTime int64)
}

//活动管理器
type PlayerActivityManager interface {
	EnterActivity(activityType actvitytypes.ActivityType, endTime int64)
	PlayerActivityPkManager
	PlayerActivityRankManager
	PlayerActivityCollectManager
	PlayerActvitiyTickRewManager
}

type PlayerActvitiyKillData struct {
	actvitiyType   activitytypes.ActivityType
	killedNum      int32 //杀次数
	lastKilledTime int64 //上次杀人时间
}

func (a *PlayerActvitiyKillData) GetKilledNum() int32 {
	return a.killedNum
}

func (a *PlayerActvitiyKillData) GetLastKilledTime() int64 {
	return a.lastKilledTime
}

func (a *PlayerActvitiyKillData) GetActivityType() activitytypes.ActivityType {
	return a.actvitiyType
}

func (a *PlayerActvitiyKillData) Reset() {
	a.killedNum = 0
	a.lastKilledTime = 0
}

func (a *PlayerActvitiyKillData) Kill(now int64) {
	a.killedNum += 1
	a.lastKilledTime = now
}

func CreatePlayerActvitiyKillData(actvitiyType activitytypes.ActivityType, killedNum int32, lastKilledTime int64) *PlayerActvitiyKillData {
	d := &PlayerActvitiyKillData{}
	d.actvitiyType = actvitiyType
	d.killedNum = killedNum
	d.lastKilledTime = lastKilledTime
	return d
}

type PlayerActivityPkManager interface {
	IfCanKilledInActivity(activityType activitytypes.ActivityType) (bool, int32)
	KilledInActivity(activityType activitytypes.ActivityType) bool
	GetPlayerActvitityKillMap() map[activitytypes.ActivityType]*PlayerActvitiyKillData
	SyncKillData(killData *PlayerActvitiyKillData)
}

//活动排行榜类型接口
type ActivityRankType interface {
	GetRankType() int32
}

type ActivityRankTypeFactory interface {
	CreateActivityRankType(rankType int32) ActivityRankType
}

type ActivityRankTypeFactoryFunc func(rankType int32) ActivityRankType

func (t ActivityRankTypeFactoryFunc) CreateActivityRankType(rankType int32) ActivityRankType {
	return t(rankType)
}

var (
	rankTypeMap = map[actvitytypes.ActivityType]ActivityRankTypeFactory{}
)

func RegistActivityRankTypeFactory(activityType actvitytypes.ActivityType, factory ActivityRankTypeFactory) {
	_, ok := rankTypeMap[activityType]
	if ok {
		panic(fmt.Errorf("活动排行榜类型已经注册，类型%d", activityType))
	}

	rankTypeMap[activityType] = factory
}

func GetActivityRankTypeFactory(activityType actvitytypes.ActivityType) ActivityRankTypeFactory {
	factory, ok := rankTypeMap[activityType]
	if !ok {
		return nil
	}

	return factory
}

type PlayerActvitiyRankData struct {
	actvitiyType activitytypes.ActivityType
	rankValueMap map[int32]int64
	endTime      int64
}

func (a *PlayerActvitiyRankData) GetRankMap() map[int32]int64 {
	return a.rankValueMap
}

func (a *PlayerActvitiyRankData) UpdateRankValue(rankType ActivityRankType, val int64) {
	a.rankValueMap[rankType.GetRankType()] = val
}

func (a *PlayerActvitiyRankData) GetRankValue(rankType ActivityRankType) int64 {
	return a.rankValueMap[rankType.GetRankType()]
}

func (a *PlayerActvitiyRankData) GetActivityType() activitytypes.ActivityType {
	return a.actvitiyType
}

func (a *PlayerActvitiyRankData) RefreshEndTime(endTime int64) bool {
	if a.endTime == endTime {
		return false
	}
	a.rankValueMap = make(map[int32]int64)
	a.endTime = endTime
	return true
}

func (a *PlayerActvitiyRankData) GetEndTime() int64 {
	return a.endTime
}

func CreatePlayerActvitiyRankData(actvitiyType activitytypes.ActivityType, rankValueMap map[int32]int64, endTime int64) *PlayerActvitiyRankData {
	d := &PlayerActvitiyRankData{}
	d.actvitiyType = actvitiyType
	d.rankValueMap = rankValueMap
	d.endTime = endTime
	return d
}

type PlayerActivityRankManager interface {
	UpdateActivityRankValue(actvitiyType activitytypes.ActivityType, rankType ActivityRankType, val int64)
	GetActivityRankValue(actvitiyType activitytypes.ActivityType, rankType ActivityRankType) int64
	GetActivityRankMap() map[actvitytypes.ActivityType]*PlayerActvitiyRankData
}

type PlayerActvitiyCollectData struct {
	activityType activitytypes.ActivityType
	countMap     map[int32]int32
	endTime      int64
}

func (a *PlayerActvitiyCollectData) GetCountMap() map[int32]int32 {
	return a.countMap
}

func (a *PlayerActvitiyCollectData) GetActivityType() activitytypes.ActivityType {
	return a.activityType
}

func (a *PlayerActvitiyCollectData) GetEndTime() int64 {
	return a.endTime
}

func (a *PlayerActvitiyCollectData) RefreshData(endTime int64) {
	if a.endTime != endTime {
		a.countMap = map[int32]int32{}
		a.endTime = endTime
	}
}

func (a *PlayerActvitiyCollectData) UpdateData(biologyId int32) {
	a.countMap[biologyId] += 1
}

func CreatePlayerActvitiyCollectData(activityType activitytypes.ActivityType, countMap map[int32]int32, endTime int64) *PlayerActvitiyCollectData {
	d := &PlayerActvitiyCollectData{}
	d.activityType = activityType
	d.countMap = countMap
	d.endTime = endTime
	return d
}

type PlayerActivityCollectManager interface {
	//获取总采集次数
	GetActivityTotalCollectCount(activityType activitytypes.ActivityType) int32
	//获取采集次数信息
	GetActivityCollectCountMap(activityType activitytypes.ActivityType) map[int32]int32
	//更新值
	UpdateActivityCollect(activityType activitytypes.ActivityType, biologyId int32)
}

type PlayerActvitiyTickRewData struct {
	enterTime          int64
	resMap             map[int32]int32
	lastRewTime        int64
	specialResMap      map[int32]int32
	lastRewSpecialTime int64
}

func (a *PlayerActvitiyTickRewData) GetResMap() map[int32]int32 {
	return a.resMap
}

func (a *PlayerActvitiyTickRewData) GetSpecialResMap() map[int32]int32 {
	return a.specialResMap
}

func (a *PlayerActvitiyTickRewData) GetEnterTime() int64 {
	return a.enterTime
}

func (a *PlayerActvitiyTickRewData) GetLastRewTime() int64 {
	return a.lastRewTime
}

func (a *PlayerActvitiyTickRewData) GetLastRewSpecialTime() int64 {
	return a.lastRewSpecialTime
}

func (a *PlayerActvitiyTickRewData) RefreshData() {
	now := global.GetGame().GetTimeService().Now()
	a.enterTime = now
	a.lastRewTime = 0
	a.lastRewSpecialTime = 0
	a.resMap = map[int32]int32{}
	a.specialResMap = map[int32]int32{}
}

func (a *PlayerActvitiyTickRewData) UpdateTickRewData(resMap map[int32]int32, specialResMap map[int32]int32) {
	now := global.GetGame().GetTimeService().Now()
	if len(resMap) > 0 {
		for itemId, num := range resMap {
			a.resMap[itemId] += num
		}

		a.lastRewTime = now
	}

	if len(specialResMap) > 0 {
		for itemId, num := range specialResMap {
			a.specialResMap[itemId] += num
		}

		a.lastRewSpecialTime = now
	}
}

func CreatePlayerActvitiyTickRewData() *PlayerActvitiyTickRewData {
	d := &PlayerActvitiyTickRewData{}
	d.enterTime = global.GetGame().GetTimeService().Now()
	d.resMap = make(map[int32]int32)
	d.specialResMap = make(map[int32]int32)
	return d
}

type PlayerActvitiyTickRewManager interface {
	//获取定时奖励
	GetActivityTickRewData() *PlayerActvitiyTickRewData
	//更新奖励
	AddActivityTickRew(resMap, specialResMap map[int32]int32)
}

type PlayerBossReliveData struct {
	bossType   worldbosstypes.BossType
	reliveTime int32
}

func (d *PlayerBossReliveData) GetBossType() worldbosstypes.BossType {
	return d.bossType
}

func (d *PlayerBossReliveData) GetReliveTime() int32 {
	return d.reliveTime
}

func (d *PlayerBossReliveData) Relive() {
	d.reliveTime += 1
}
func (d *PlayerBossReliveData) Reset() {
	d.reliveTime = 0
}

func (d *PlayerBossReliveData) Sync(reliveTime int32) {
	d.reliveTime = reliveTime
}

func CreatePlayerBossReliveData(bossType worldbosstypes.BossType, reliveTime int32) *PlayerBossReliveData {
	d := &PlayerBossReliveData{}
	d.bossType = bossType
	d.reliveTime = reliveTime
	return d
}

type PlayerBossReliveManager interface {
	//获取总采集次数
	GetBossReliveTime(bossType worldbosstypes.BossType) int32
	PlayerBossRelive(bossType worldbosstypes.BossType)
	PlayerBossReset(bossType worldbosstypes.BossType)
	PlayerBossReliveSync(bossType worldbosstypes.BossType, reliveTime int32)
	GetPlayerBossReliveMap() map[worldbosstypes.BossType]*PlayerBossReliveData
}

//玩家接口
type Player interface {
	//cd组
	GetCDGroupManager() *cdcommon.CDGroupManager
	GetExtraSpeed() int64
	//玩家场景管理器
	PlayerSceneManager
	//buff
	BuffManager
	//技能
	SkillManager
	//特殊技能
	TeShuSkillManager
	//状态数据管理器
	StateDataManager
	//玩家pk管理器
	PlayerPkManager
	//玩家战斗
	PlayerBattleManager
	//玩家系统管理器
	SystemPropertyManager
	//玩家战斗属性
	PlayerBattlePropertyManager
	//玩家展示管理器
	PlayerShowManager
	//玩家技能管理器
	SkillActionManager
	//玩家竞技管理器
	PlayerArenaManager
	//玩家竞技pvp管理器
	PlayerArenapvpManager
	//玩家血池管理器
	PlayerXueChiManager
	//玩家复活管理器
	PlayerReliveManager
	//玩家采集管理器
	PlayerCollectManager
	//玩家队伍管理器
	PlayerTeamManager
	//玩家打宝塔管理器
	PlayerTowerManager
	//玩家天书管理器
	PlayerTianShuManager
	//玩家幸运管理器
	PlayerLuckyManager
	//玩家仙盟管理器
	PlayerAllianceManager
	//结义管理器
	PlayerJieYiManager
	//玩家神兽攻城管理器
	PlayerGodSiegeManager
	//玩家挂机管理器
	PlayerGuaJiManager
	//玩家幻境BOSS管理器
	PlayerUnrealBossManager
	//玩家外域BOSS管理器
	PlayerOutlandBossManager
	//玩家灵童管理器
	PlayerLingTongManager
	//玩家金银密窟管理器
	PlayerDenseWatManager
	//玩家神魔战场管理器
	PlayerShenMoManager
	//玩家活动管理器
	PlayerActivityManager
	//阵营
	PlayerZhenYingManager
	//boss复活数据
	PlayerBossReliveManager
	//移动管理器
	MoveManager
	GetContext() context.Context
	Tick()
	Heartbeat()
	Post(msg message.Message)
	//返回上一个场景
	BackLastScene()
	//重生

	SendMsg(msg proto.Message) error
	Close(err error)

	GetRole() types.RoleType
	GetName() string
	GetOriginName() string
	GetSex() types.SexType
	GetServerId() int32
	GetPlatform() int32
	GetForce() int64
	UpdateForce(force int64)

	//设置竞技场队伍
	SetArenaTeam(teamId int64, teamName string, purpose teamtypes.TeamPurposeType)

	//进入场景
	EnteringScene() bool
	//进入游戏
	EnterGame() bool
	//离开场景
	LeaveScene() bool
	GetDeadTime() int64
	IsGuaJiPlayer() bool
}

const (
	playerContextKey contextKey = "player"
)

func WithPlayer(parent context.Context, p Player) context.Context {
	return context.WithValue(parent, playerContextKey, p)
}

func PlayerInContext(ctx context.Context) Player {
	p := ctx.Value(playerContextKey)
	if p == nil {
		return nil
	}
	tp, ok := p.(Player)
	if !ok {
		return nil
	}
	return tp
}
