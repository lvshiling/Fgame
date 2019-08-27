package player

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/fsm"
	playereventtypes "fgame/fgame/cross/player/event/types"
	alliancecommon "fgame/fgame/game/alliance/common"
	"fgame/fgame/game/battle/battle"
	battlecommon "fgame/fgame/game/battle/common"
	"fgame/fgame/game/buff/buff"
	buffcommon "fgame/fgame/game/buff/common"
	cdcommon "fgame/fgame/game/cd/common"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	crosstypes "fgame/fgame/game/cross/types"
	densewatcommon "fgame/fgame/game/densewat/common"
	gameevent "fgame/fgame/game/event"
	jieyicommon "fgame/fgame/game/jieyi/common"
	pkcommon "fgame/fgame/game/pk/common"
	playercommon "fgame/fgame/game/player/common"
	"fgame/fgame/game/player/types"
	relivecommon "fgame/fgame/game/relive/common"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	gamesession "fgame/fgame/game/session"
	shenmocommon "fgame/fgame/game/shenmo/common"
	skillcommon "fgame/fgame/game/skill/common"
	"fgame/fgame/game/skill/skill"
	teamcommon "fgame/fgame/game/team/common"
	teamtypes "fgame/fgame/game/team/types"
	xuechicommon "fgame/fgame/game/xuechi/common"
	"fmt"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type Player struct {
	m sync.Mutex
	//外部对话
	s gamesession.Session
	//状态
	*fsm.SubjectBase
	//消息队列
	msgQueue *MessageQueue
	//玩家id
	playerId int64
	//玩家基本信息
	playerObject playercommon.PlayerCommonObject
	//技能
	*skill.SkillManager
	//buff管理器
	*buff.BuffDataManager
	//战斗管理器
	*battle.PlayerBattleManager
	//属性管理器
	*battle.PlayerBattlePropertyManager
	//场景管理器
	*battle.PlayerSceneManager
	//场景展示
	*battle.PlayerShowManager
	//pk管理器
	*battle.PlayerPKManager
	//系统属性
	*battle.SystemPropertyManager
	//竞技场数据
	*battle.PlayerArenaManager
	//pvp竞技场数据
	*battle.PlayerArenapvpManager
	//血池
	*battle.PlayerXueChiManager
	//复活管理器
	*battle.PlayerReliveManager
	*battle.PlayerCollectManager
	*battle.PlayerTeamManager
	*battle.PlayerTowerManager
	*battle.PlayerAllianceManager
	*battle.PlayerJieYiManager
	*battle.PlayerZhenYingManager
	*battle.PlayerBossReliveManager
	//天书管理器
	*battle.PlayerTianShuManager
	//神兽攻城管理器
	*battle.PlayerGodSiegeManager
	//幸运管理器
	*battle.PlayerLuckyManager
	*battle.StateDataManager
	//挂机管理器
	*battle.PlayerGuaJiManager

	//幻境BOSS管理器
	*battle.PlayerUnrealBossManager

	//外域BOSS管理器
	*battle.PlayerOutlandBossManager
	//灵童展示管理器
	*battle.PlayerLingTongShowManager
	//金银密窟管理器
	*battle.PlayerDenseWatManager
	//神魔战场管理器
	*battle.PlayerShenMoManager
	//移动
	*battle.MoveAction
	*battle.PlayerActivityManager
	*battle.TeShuSkillManager
	//cd组管理器
	cdGroupManager *cdcommon.CDGroupManager

	//竞技场队伍id
	arenaTeamId int64
	//竞技场队伍名字
	arenaTeamName string
	//队伍标识
	arenaTeamPurpose teamtypes.TeamPurposeType
	//跨服类型
	crossType crosstypes.CrossType
	logouting bool

	done chan struct{}

	showServerId bool
}

func (p *Player) Session() gamesession.Session {
	return p.s
}

func (p *Player) GetId() int64 {
	return p.playerId
}

func (p *Player) GetContext() context.Context {
	return p.s.Context()
}

func (p *Player) Tick() {
	p.msgQueue.Tick()
	return
}

func (p *Player) Heartbeat() {
	p.BuffDataManager.Heartbeat()
	//战斗
	p.PlayerBattleManager.Heartbeat()
	p.SkillManager.Heartbeat()
	p.MoveAction.Heartbeat()
	p.PlayerBattlePropertyManager.Heartbeat()
	p.PlayerXueChiManager.Heartbeat()
	p.PlayerReliveManager.Heartbeat()
	p.PlayerGuaJiManager.Heartbeat()
	return
}

func (p *Player) Post(msg message.Message) {
	p.msgQueue.Post(msg)
	return
}

func (p *Player) SendMsg(msg proto.Message) error {
	//发送消息
	p.s.Send(msg)
	return nil
}

func (p *Player) Close(err error) {
	p.Session().Close(true)
	return
}

func (p *Player) GetRole() types.RoleType {
	return p.playerObject.GetRole()
}

func (p *Player) GetName() string {
	if p.showServerId {
		return fmt.Sprintf("s%d.%s", p.playerObject.GetServerId(), p.playerObject.GetName())
	}
	return p.playerObject.GetName()
}

func (p *Player) GetOriginName() string {
	return p.playerObject.GetName()
}

func (p *Player) GetServerId() int32 {
	return p.playerObject.GetServerId()
}

func (p *Player) GetPlatform() int32 {
	return p.playerObject.GetPlatform()
}

func (p *Player) GetSex() types.SexType {
	return p.playerObject.GetSex()
}

func (p *Player) IsGuaJiPlayer() bool {
	return p.playerObject.IsGuaJi()
}

func (p *Player) GetTeamId() int64 {
	if p.arenaTeamId != 0 {
		return p.arenaTeamId
	}

	return p.PlayerTeamManager.GetTeamId()
}

func (p *Player) GetTeamName() string {
	if p.arenaTeamName != "" {
		return p.arenaTeamName
	}
	return p.PlayerTeamManager.GetTeamName()
}

func (p *Player) GetTeamPurpose() teamtypes.TeamPurposeType {
	if p.arenaTeamPurpose != teamtypes.TeamPurposeTypeNormal {
		return p.arenaTeamPurpose
	}
	return p.PlayerTeamManager.GetTeamPurpose()
}

//更新竞技场
func (p *Player) SetArenaTeam(teamId int64, teamName string, teamPurpose teamtypes.TeamPurposeType) {
	p.arenaTeamId = teamId
	p.arenaTeamName = teamName
	p.arenaTeamPurpose = teamPurpose
}

//获取跨服类型
func (p *Player) GetCrossType() crosstypes.CrossType {
	return p.crossType
}

func (p *Player) EnteringScene() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := GetPlayerStateMachine().Trigger(p, EventPlayerEnterScene)
	if !flag {
		return false
	}

	return true
}

//进入游戏
func (p *Player) EnterGame() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := GetPlayerStateMachine().Trigger(p, EventPlayerGaming)
	if !flag {
		return false
	}
	return true
}

//退出场景
func (p *Player) LeaveScene() bool {
	p.m.Lock()
	defer p.m.Unlock()
	flag := GetPlayerStateMachine().Trigger(p, EventPlayerLeaveScene)
	if !flag {
		return false
	}
	p.msgQueue.Pause()

	return true
}

func (p *Player) IsLogouting() bool {
	return p.logouting
}

func (p *Player) Logout() bool {
	p.m.Lock()
	defer p.m.Unlock()
	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
			"state":    p.CurrentState(),
		}).Info("player:玩家正在登出")
	switch p.CurrentState() {
	case PlayerStateAuth,
		PlayerStateLoaded,
		PlayerStateEnterScene,
		PlayerStateLeaveScene:
		flag := GetPlayerStateMachine().Trigger(p, EventPlayerLogout)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"state":    p.CurrentState(),
				}).Info("player:登出失败")
			return false
		}
		gameevent.Emit(playereventtypes.EventTypeCrossPlayerBeforeLogout, p, nil)
		return true
	case PlayerStateGaming:
		p.logouting = true
		gameevent.Emit(playereventtypes.EventTypeCrossPlayerExitSceneBeforeLogout, p, nil)
		return true
	case PlayerStateLogouting,
		PlayerStateLogouted:
		return true
	}
	return true
}

func (p *Player) LogoutSave() bool {
	p.m.Lock()
	defer p.m.Unlock()
	return p.logoutSave()
}

func (p *Player) logoutSave() bool {
	if p.logouting {
		flag := GetPlayerStateMachine().Trigger(p, EventPlayerLogout)
		if !flag {
			log.WithFields(
				log.Fields{
					"playerId": p.GetId(),
					"state":    p.CurrentState(),
				}).Info("player:登出失败")
			return false
		}
	}
	defer func() {
		close(p.done)
	}()

	log.WithFields(
		log.Fields{
			"playerId": p.GetId(),
		}).Info("player:玩家登出,成功")
	gameevent.Emit(playereventtypes.EventTypeCrossPlayerLogout, p, nil)
	return true
}

func (p *Player) GetScene() scene.Scene {
	if p.PlayerSceneManager == nil {
		return nil
	}
	return p.PlayerSceneManager.GetScene()
}

func (p *Player) GetCDGroupManager() *cdcommon.CDGroupManager {
	return p.cdGroupManager
}

func (p *Player) Load(
	commonObj playercommon.PlayerCommonObject,
	showObj *battle.PlayerShowObject,
	pkObj pkcommon.PlayerPkObject,
	basePropertyObj map[int32]int64,
	battlePropertyObj map[int32]int64,
	skillList []skillcommon.SkillObject,
	buffDataList []buffcommon.BuffObject,
	teamObj teamcommon.PlayerTeamObject,
	allianceObj alliancecommon.PlayerAllianceObject,
	arenaObj *battle.PlayerArenaObject,
	arenapvpObj *battle.PlayerArenapvpObject,
	xueChiObj *xuechicommon.PlayerXueChiObject,
	reliveObj *relivecommon.PlayerReliveObject,
	battleObj *battlecommon.PlayerBattleObject,
	denseWatObj *densewatcommon.PlayerDenseWatObject,
	shenMoObj *shenmocommon.PlayerShenMoObject,
	crossType crosstypes.CrossType,
	power int64,
	lingTong scene.LingTong,
	activityKillDataList []*scene.PlayerActvitiyKillData,
	activityRankDataList []*scene.PlayerActvitiyRankData,
	jieYiObj jieyicommon.PlayerJieYiObject,
	// chuangShiObj chuangshidata.CommonPlayerChuangShiObject,
	bossReliveList []*scene.PlayerBossReliveData,
	teShuSkillList []*scene.TeshuSkillObject,
	showServerId bool) bool {
	flag := GetPlayerStateMachine().Trigger(p, EventPlayerLoaded)
	if !flag {
		return false
	}
	//TODO 使用基础属性
	p.cdGroupManager = cdcommon.NewCDGroupManager()
	//基础对象
	p.playerObject = commonObj
	hp := int64(0)
	tp := int64(0)

	p.PlayerBattlePropertyManager = battle.CreatePlayerBattlePropertyManager(p, hp, tp, power)
	p.PlayerShowManager = battle.CreatePlayerShowManagerWithObject(p, showObj)
	p.PlayerSceneManager = battle.CreatePlayerSceneManager(p)
	p.PlayerPKManager = battle.CreatePlayerPKManagerWithObject(p, pkObj)
	p.BuffDataManager = buff.CreateBuffDataManagerWithBuffList(p, buffDataList)
	p.SkillManager = skill.CreateSkillManager(p, p.cdGroupManager, skillList)
	p.PlayerBattleManager = battle.CreatePlayerBattleManagerWithObject(p, false, battleObj, scenetypes.FactionTypePlayer)
	p.SystemPropertyManager = battle.CreateSystemPropertyManagerWithData(p, battlePropertyObj)

	//竞技场pvp
	p.PlayerArenapvpManager = battle.CreatePlayerArenapvpManagerWithObject(p, arenapvpObj)
	//仙盟数据
	p.PlayerArenaManager = battle.CreatePlayerArenaManagerWithObject(p, arenaObj)
	p.PlayerXueChiManager = battle.CreatePlayerXueChiManager(p, xueChiObj.GetBlood(), xueChiObj.GetBloodLine())
	p.PlayerReliveManager = battle.CreatePlayerReliveManagerWithObj(p, reliveObj)
	p.PlayerCollectManager = battle.CreatePlayerCollectManager(p)
	p.PlayerTeamManager = battle.CreatePlayerTeamManagerWithObject(p, teamObj)

	//打宝塔管理器
	p.PlayerTowerManager = battle.CreatePlayerTowerManager(p)
	p.PlayerTianShuManager = battle.CreatePlayerTianShuManager(p)
	p.PlayerAllianceManager = battle.CreatePlayerAllianceManagerWithObject(p, allianceObj)
	// 阵营
	// camp := chuangShiObj.GetCampType()
	// guanZhi := chuangShiObj.GetPos()
	camp := chuangshitypes.RandomChuangShiCamp()
	guanZhi := chuangshitypes.RandomChuangShiGuanZhi()
	p.PlayerZhenYingManager = battle.CreatePlayerZhenYingManager(p, camp, guanZhi)

	p.PlayerBossReliveManager = battle.CreatePlayerBossReliveManager(p, bossReliveList)
	//结义
	p.PlayerJieYiManager = battle.CreatePlayerJieYiManagerWithObject(p, jieYiObj)
	//神兽攻城管理器
	p.PlayerGodSiegeManager = battle.CreatePlayerGodSiegeManager(p)
	//幸运管理器
	p.PlayerLuckyManager = battle.CreatePlayerLuckyManager(p)
	p.StateDataManager = battle.CreateStateDateManager(p)
	//挂机管理器
	p.PlayerGuaJiManager = battle.CreatePlayerGuaJiManager(p)

	//幻境BOSS管理器
	p.PlayerUnrealBossManager = battle.CreatePlayerUnrealBossManager(p)

	//外域BOSS管理器
	p.PlayerOutlandBossManager = battle.CreatePlayerOutlandBossManager(p)
	//灵童
	p.PlayerLingTongShowManager = battle.CreatePlayerLingTongShowManagerWithLingTong(p, lingTong)
	//金银密窟
	p.PlayerDenseWatManager = battle.CreatePlayerDenseWatManager(p, denseWatObj.GetNum(), denseWatObj.GetEndTime())
	//神魔战场
	p.PlayerShenMoManager = battle.CreatePlayerShenMoManager(p, shenMoObj.GetGongXunNum(), shenMoObj.GetKillNum(), shenMoObj.GetEndTime())
	//活动管理器
	tickRewData := scene.CreatePlayerActvitiyTickRewData()
	p.PlayerActivityManager = battle.CreatePlayerActivityManager(p, activityKillDataList, activityRankDataList, nil, tickRewData)
	//移动管理器
	p.MoveAction = battle.CreateMoveAction(p)
	p.TeShuSkillManager = battle.CreateTeShuSkillManager(p, teShuSkillList)
	//设置跨服类型
	p.crossType = crossType
	p.showServerId = showServerId
	gameevent.Emit(playereventtypes.EventTypeCrossPlayerAfterLoad, p, nil)

	//更新战斗属性
	p.Calculate()
	return true
}

const (
	queueCapacity = 1000
	maxTime       = time.Microsecond * 10
)

func NewPlayer(s gamesession.Session, playerId int64) *Player {
	p := &Player{}
	p.SubjectBase = fsm.NewSubjectBase(PlayerStateAuth)
	p.s = s
	p.msgQueue = NewMessageQueue(p, queueCapacity, maxTime)
	p.playerId = playerId
	p.done = make(chan struct{})
	return p
}
