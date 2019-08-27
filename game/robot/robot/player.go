package robot

import (
	"context"
	"fgame/fgame/common/message"
	"fgame/fgame/core/fsm"
	"fgame/fgame/core/heartbeat"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/battle/battle"
	battlecommon "fgame/fgame/game/battle/common"
	"fgame/fgame/game/buff/buff"
	buffcommon "fgame/fgame/game/buff/common"
	cdcommon "fgame/fgame/game/cd/common"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	gameevent "fgame/fgame/game/event"
	exceptioneventtypes "fgame/fgame/game/exception/event/types"
	playercommon "fgame/fgame/game/player/common"
	"fgame/fgame/game/player/types"
	robottypes "fgame/fgame/game/robot/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skillcommon "fgame/fgame/game/skill/common"
	"fgame/fgame/game/skill/skill"
	skilltemplate "fgame/fgame/game/skill/template"
	teamcommon "fgame/fgame/game/team/common"
	teamtypes "fgame/fgame/game/team/types"
	"fmt"
	"runtime/debug"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
)

type robotPlayer struct {
	*fsm.SubjectBase
	//消息队列
	msgQueue *MessageQueue
	po       playercommon.PlayerCommonObject
	//展示数据
	*battle.PlayerShowManager
	//pk管理器
	*battle.PlayerPKManager
	//竞技场数据
	*battle.PlayerArenaManager
	//竞技场pvp数据
	*battle.PlayerArenapvpManager
	//buff管理器
	*buff.BuffDataManager
	//技能管理器
	*skill.SkillManager
	//战斗管理器
	*battle.PlayerBattleManager
	//属性管理器
	*battle.PlayerBattlePropertyManager
	//系统属性
	*battle.SystemPropertyManager
	//场景管理器
	*battle.PlayerSceneManager
	//血池管理器
	*battle.PlayerXueChiManager
	//复活管理器
	*battle.PlayerReliveManager
	//队伍管理器
	*battle.PlayerTeamManager
	*battle.PlayerCollectManager
	// 打宝塔管理器
	*battle.PlayerTowerManager
	//天书管理器
	*battle.PlayerTianShuManager
	//幸运管理器
	*battle.PlayerLuckyManager
	//幻境BOSS管理器
	*battle.PlayerUnrealBossManager
	//外域BOSS管理器
	*battle.PlayerOutlandBossManager
	//仙盟管理器
	*battle.PlayerAllianceManager
	*battle.PlayerJieYiManager
	*battle.PlayerZhenYingManager
	*battle.PlayerBossReliveManager
	//神兽攻城管理器
	*battle.PlayerGodSiegeManager
	//挂机管理器
	*battle.PlayerGuaJiManager
	//状态数据管理器
	*battle.StateDataManager
	//灵童管理器
	*battle.PlayerLingTongShowManager
	//金银密窟
	*battle.PlayerDenseWatManager
	//神魔战场
	*battle.PlayerShenMoManager
	*battle.PlayerActivityManager
	*battle.TeShuSkillManager
	//移动
	*battle.MoveAction
	//cd组管理器
	cdGroupManager *cdcommon.CDGroupManager
	*RobotStateManager
	//心跳处理器
	heartbeatRunner heartbeat.HeartbeatTaskRunner
	//机器人类型
	robotType robottypes.RobotType

	attackTarget *scene.Enemy

	//竞技场
	arenaTeamId      int64
	arenaTeamName    string
	arenaTeamPurpose teamtypes.TeamPurposeType
	//复活次数
	canReliveTime int32
	questBeginId  int32
	questEndId    int32

	//是否展示服务器id
	showServerId bool
}

//----------------------------------scene object接口 ---------------------------------

//进入场景
func (p *robotPlayer) EnteringScene() bool {
	return true
}

//进入游戏
func (p *robotPlayer) EnterGame() bool {
	return true
}

//退出场景
func (p *robotPlayer) LeaveScene() bool {
	return true
}

//------------------------------玩家接口 --------------------------------

func (p *robotPlayer) Post(msg message.Message) {
	//TODO 可能造成死锁  做个超时处理
	p.msgQueue.Post(msg)
}

//tick
func (p *robotPlayer) Tick() {
	p.msgQueue.Tick()
}

//玩家定时心跳
func (p *robotPlayer) Heartbeat() {
	defer func() {
		if terr := recover(); terr != nil {
			debug.PrintStack()
			exceptionContent := string(debug.Stack())
			log.WithFields(
				log.Fields{
					"error": terr,
					"stack": string(debug.Stack()),
				}).Error("robot:Heartbeat,错误")
			gameevent.Emit(exceptioneventtypes.ExceptionEventTypeException, nil, exceptionContent)
		}
	}()

	p.MoveAction.Heartbeat()
	p.BuffDataManager.Heartbeat()
	p.SkillManager.Heartbeat()
	p.PlayerBattlePropertyManager.Heartbeat()
	//战斗
	p.PlayerBattleManager.Heartbeat()
	p.RobotStateManager.Heartbeat()

	p.heartbeatRunner.Heartbeat()
}

//复活
func (p *robotPlayer) Reborn(pos coretypes.Position) {
	flag := p.RobotStateManager.Idle()
	if !flag {
		return
	}
	p.PlayerBattlePropertyManager.Reborn(pos)

}

//重载
//返回是否死亡
func (p *robotPlayer) CostHP(hp int64, attackId int64) bool {
	if hp <= 0 {
		return false
	}
	dead := p.PlayerBattlePropertyManager.CostHP(hp, attackId)
	if dead {
		flag := p.RobotStateManager.RobotDead()
		if !flag {
			return false
			// panic(fmt.Errorf("robot：状态改变应该成功[%]"))
		}
		return true
	}
	return dead
}

// func (p *robotPlayer) Dead(attackId int64) bool {

// 	//状态改变
// 	flag := p.RobotStateManager.RobotDead()
// 	if !flag {
// 		return false
// 		// panic(fmt.Errorf("robot：状态改变应该成功[%]"))
// 	}
// 	return p.PlayerBattlePropertyManager.Dead(attackId)
// }

//重载
func (p *robotPlayer) UseSkill(skillId int32) bool {

	skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByType(skillId)
	if skillTemplate == nil {
		return false
	}
	//判断cd
	if p.SkillManager.IsSkillInCd(skillTemplate.TypeId) {
		return false
	}

	//技能动作
	if skillTemplate.IsPositive() {
		flag := p.RobotStateManager.Attack()
		if !flag {
			return false
		}
		p.SetSkillActionTime(int64(skillTemplate.ActionTime))

	}

	//使用技能
	flag := p.SkillManager.UseSkill(skillTemplate.TypeId)
	if !flag {
		panic(fmt.Errorf("robot:使用技能应该成功"))
	}

	return true
}

func (p *robotPlayer) SendMsg(msg proto.Message) error {
	return nil
}

func (p *robotPlayer) Close(err error) {
	//移除机器人
	return
}

func (p *robotPlayer) GetCDGroupManager() *cdcommon.CDGroupManager {
	return p.cdGroupManager
}

func (p *robotPlayer) GetContext() context.Context {
	return nil
}

func (p *robotPlayer) GetId() int64 {
	if p.po == nil {
		return 0
	}
	return p.po.GetId()
}

//快捷操作
func (p *robotPlayer) GetRole() types.RoleType {
	return p.po.GetRole()
}

func (p *robotPlayer) GetSex() types.SexType {
	return p.po.GetSex()
}

func (p *robotPlayer) GetName() string {
	// if merge.GetMergeService().GetMergeTime() > 0 {
	// 	return fmt.Sprintf("s%d.%s", p.po.GetServerId(), p.po.GetName())
	// }
	if p.showServerId {
		return fmt.Sprintf("s%d.%s", p.po.GetServerId(), p.po.GetName())
	}
	return p.po.GetName()
}

func (p *robotPlayer) GetOriginName() string {
	return p.po.GetName()
}

func (p *robotPlayer) GetServerId() int32 {
	return p.po.GetServerId()
}

func (p *robotPlayer) GetPlatform() int32 {
	return p.po.GetPlatform()
}

func (p *robotPlayer) GetAllianceId() int64 {
	return 0
}

func (p *robotPlayer) GetAllianceName() string {
	return ""
}

func (p *robotPlayer) GetMengZhuId() int64 {
	return 0
}

func (p *robotPlayer) GetTeamId() int64 {
	return p.arenaTeamId
}

func (p *robotPlayer) GetTeamName() string {
	return p.arenaTeamName
}

func (p *robotPlayer) GetTeamPurpose() teamtypes.TeamPurposeType {
	return p.arenaTeamPurpose
}

func (p *robotPlayer) SetArenaTeam(teamId int64, teamName string, teamPurpose teamtypes.TeamPurposeType) {
	p.arenaTeamId = teamId
	p.arenaTeamName = teamName
	p.arenaTeamPurpose = teamPurpose

}

//重载
func (p *robotPlayer) GetScene() scene.Scene {
	if p.PlayerSceneManager == nil {
		return nil
	}
	return p.PlayerSceneManager.GetScene()
}

//获取机器人类型
func (p *robotPlayer) GetRobotType() robottypes.RobotType {
	return p.robotType
}

func (p *robotPlayer) GetCanReliveTime() int32 {
	return p.canReliveTime
}

func (p *robotPlayer) GetQuestBeginId() int32 {
	return p.questBeginId
}

func (p *robotPlayer) GetQuestEndId() int32 {
	return p.questEndId
}

func (p *robotPlayer) IsGuaJiPlayer() bool {
	return false
}

const (
	queueCapacity = 10000
	maxTime       = time.Microsecond * 10
)

func createRobotPlayer(robotType robottypes.RobotType, po playercommon.PlayerCommonObject,
	showObj *battle.PlayerShowObject,
	buffs []buffcommon.BuffObject,
	skillList []skillcommon.SkillObject,
	battleProperties map[int32]int64,
	canReliveTime int32,
	playerBattleObject *battlecommon.PlayerBattleObject,
	showServerId bool,
	power int64,
) scene.RobotPlayer {
	return createRobotPlayerWithQuest(robotType, po, showObj, buffs, skillList, battleProperties, canReliveTime, playerBattleObject, showServerId, 0, 0, power)
}

func createRobotPlayerWithQuest(robotType robottypes.RobotType, po playercommon.PlayerCommonObject,
	showObj *battle.PlayerShowObject,
	buffs []buffcommon.BuffObject,
	skillList []skillcommon.SkillObject,
	battleProperties map[int32]int64,
	canReliveTime int32,
	playerBattleObject *battlecommon.PlayerBattleObject,
	showServerId bool,
	questBeginId int32,
	questEndId int32,
	power int64,
) scene.RobotPlayer {
	return createRobotPlayerBase(robotType, po, showObj, buffs, skillList, battleProperties, canReliveTime, playerBattleObject, showServerId, questBeginId, questEndId, power, scenetypes.FactionTypePlayer)
}

func createRobotPlayerBase(robotType robottypes.RobotType, po playercommon.PlayerCommonObject,
	showObj *battle.PlayerShowObject,
	buffs []buffcommon.BuffObject,
	skillList []skillcommon.SkillObject,
	battleProperties map[int32]int64,
	canReliveTime int32,
	playerBattleObject *battlecommon.PlayerBattleObject,
	showServerId bool,
	questBeginId int32,
	questEndId int32,
	power int64,
	factionType scenetypes.FactionType,
) scene.RobotPlayer {
	p := &robotPlayer{
		po: po,
	}
	p.msgQueue = NewMessageQueue(p, queueCapacity, maxTime)
	p.robotType = robotType
	p.SubjectBase = fsm.NewSubjectBase(RobotPlayerStateIdle)
	p.heartbeatRunner = heartbeat.NewHeartbeatTaskRunner()
	p.cdGroupManager = cdcommon.NewCDGroupManager()
	p.BuffDataManager = buff.CreateBuffDataManagerWithBuffList(p, buffs)
	p.PlayerSceneManager = battle.CreatePlayerSceneManager(p)
	p.PlayerBattleManager = battle.CreatePlayerBattleManagerWithObject(p, true, playerBattleObject, factionType)
	p.PlayerBattlePropertyManager = battle.CreatePlayerBattlePropertyManager(p, 0, 0, power)
	p.PlayerPKManager = battle.CreatePlayerPKManager(p)
	p.SkillManager = skill.CreateSkillManager(p, p.cdGroupManager, skillList)
	p.SystemPropertyManager = battle.CreateSystemPropertyManagerWithData(p, battleProperties)
	p.PlayerShowManager = battle.CreatePlayerShowManagerWithObject(p, showObj)
	p.RobotStateManager = NewRobotStateManager(p)
	arenaObj := battle.CreatePlayerArenaObject(0, 0)
	p.PlayerArenaManager = battle.CreatePlayerArenaManagerWithObject(p, arenaObj)
	arenapvpObj := battle.CreatePlayerArenapvpObject(0)
	p.PlayerArenapvpManager = battle.CreatePlayerArenapvpManagerWithObject(p, arenapvpObj)
	p.PlayerXueChiManager = battle.CreatePlayerXueChiManager(p, 0, 0)
	p.PlayerReliveManager = battle.CreatePlayerReliveManager(p, 0, 0)
	p.PlayerCollectManager = battle.CreatePlayerCollectManager(p)
	p.PlayerTowerManager = battle.CreatePlayerTowerManager(p)
	p.PlayerAllianceManager = battle.CreatePlayerAllianceManager(p)
	camp := chuangshitypes.RandomChuangShiCamp()
	guanZhi := chuangshitypes.RandomChuangShiGuanZhi()
	p.PlayerZhenYingManager = battle.CreatePlayerZhenYingManager(p, camp, guanZhi)
	p.PlayerBossReliveManager = battle.CreatePlayerBossReliveManager(p, nil)
	p.PlayerJieYiManager = battle.CreatePlayerJieYiManager(p)
	teamObj := teamcommon.CreatePlayerTeamObject(0, "", teamtypes.TeamPurposeTypeNormal)
	p.PlayerTeamManager = battle.CreatePlayerTeamManagerWithObject(p, teamObj)
	p.PlayerTianShuManager = battle.CreatePlayerTianShuManager(p)
	p.PlayerLuckyManager = battle.CreatePlayerLuckyManager(p)
	p.PlayerGuaJiManager = battle.CreatePlayerGuaJiManager(p)
	p.StateDataManager = battle.CreateStateDateManager(p)
	p.PlayerUnrealBossManager = battle.CreatePlayerUnrealBossManager(p)
	p.PlayerOutlandBossManager = battle.CreatePlayerOutlandBossManager(p)
	p.PlayerDenseWatManager = battle.CreatePlayerDenseWatManager(p, 0, 0)
	p.PlayerShenMoManager = battle.CreatePlayerShenMoManager(p, 0, 0, 0)

	p.PlayerLingTongShowManager = battle.CreatePlayerLingTongShowManager(p)
	p.PlayerActivityManager = battle.CreatePlayerActivityManager(p, nil, nil, nil, nil)

	p.TeShuSkillManager = battle.CreateTeShuSkillManager(p, nil)
	p.MoveAction = battle.CreateMoveAction(p)
	p.showServerId = showServerId
	p.questBeginId = questBeginId
	p.questEndId = questEndId
	p.Calculate()
	p.heartbeatRunner.AddTask(CreateRobotTask(p))
	if robotType == robottypes.RobotTypeTest {
		p.heartbeatRunner.AddTask(CreateRandomRobotTask(p))
	}

	p.canReliveTime = canReliveTime

	return p
}
