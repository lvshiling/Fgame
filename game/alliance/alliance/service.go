package alliance

import (
	"fgame/fgame/common/lang"
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/runner"
	coreutils "fgame/fgame/core/utils"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	alliancebossscene "fgame/fgame/game/alliance/boss_scene"
	"fgame/fgame/game/alliance/dao"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancescene "fgame/fgame/game/alliance/scene"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/center/center"
	chargetemplate "fgame/fgame/game/charge/template"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	goldequiptypes "fgame/fgame/game/goldequip/types"
	inventorytypes "fgame/fgame/game/inventory/types"
	"fgame/fgame/game/item/item"
	itemtypes "fgame/fgame/game/item/types"
	"fgame/fgame/game/merge/merge"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/game/scene/scene"
	biaochenpc "fgame/fgame/game/transportation/npc/biaoche"
	"fgame/fgame/game/transportation/transpotation"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"
)

type AllianceService interface {
	//定时器
	runner.Task
	//启动定时任务
	Start()
	//获取仙盟成员
	GetAllianceMember(memberId int64) *AllianceMemberObject
	//获取仙盟列表
	GetAllianceList() []*Alliance
	// //创建仙盟
	// CreateAlliance(campType chuangshitypes.ChuangShiCampType, createId int64, name string, role playertypes.RoleType, sex playertypes.SexType, plName string, createrVip int32, versionTyp alliancetypes.AllianceVersionType, typ alliancetypes.AllianceNewType) (a *Alliance, err error)
	//创建仙盟
	CreateAlliance(createId int64, name string, role playertypes.RoleType, sex playertypes.SexType, plName string, createrVip int32, createLingYuId int32, lev int32, versionTyp alliancetypes.AllianceVersionType, typ alliancetypes.AllianceNewType) (a *Alliance, err error)
	//解散仙盟
	DismissAlliance(memberId int64) (al *Alliance, tempMemberList []*AllianceMemberObject, err error)
	// //加入仙盟申请
	// ApplyJoinAlliance(allianceId int64, joinId int64, campType chuangshitypes.ChuangShiCampType, name string, role playertypes.RoleType, sex playertypes.SexType, level int32, force int64) (al *Alliance, joinApplyObj *AllianceJoinApplyObject, err error)
	//加入仙盟申请
	ApplyJoinAlliance(allianceId int64, joinId int64, name string, role playertypes.RoleType, sex playertypes.SexType, level int32, force int64) (al *Alliance, joinApplyObj *AllianceJoinApplyObject, err error)
	//同意加入
	AgreeAllianceJoinApply(agreeId int64, joinId int64, agree bool) (al *Alliance, joinApplyObj *AllianceJoinApplyObject, err error)
	//获取加入列表
	GetJoinApplyList(getMemberId int64) (joinApplyObj []*AllianceJoinApplyObject, err error)
	//同步成员数据
	SyncMemberInfo(memberId int64, name string, sex playertypes.SexType, level int32, force int64, zhuansheng int32, lingyuId int32, vip int32) (err error)
	//离线
	OfflineMember(memberId int64) (err error)
	//在线
	OnlineMember(memberId int64) (err error)
	//退出仙盟
	ExitAlliance(memberId int64) (al *Alliance, err error)
	//踢人
	Kick(memberId, kickMemberId int64) (member *AllianceMemberObject, kickMember *AllianceMemberObject, err error)
	//任命
	Commit(memberId, commitMemberId int64, position alliancetypes.AlliancePosition) (member *AllianceMemberObject, commitMember *AllianceMemberObject, err error)
	//转让
	Transfer(memberId, transferMemberId int64) (member *AllianceMemberObject, transferMember *AllianceMemberObject, err error)
	//弹劾
	Impeach(impeachId int64) (impeachMember *AllianceMemberObject, mengZhuMember *AllianceMemberObject, err error)
	//邀请
	Invitation(memberId, invitationId int64, name string, role playertypes.RoleType, sex playertypes.SexType, level int32, force int64) (al *Alliance, err error)
	//创建战场数据
	CreateAllianceSceneData(chengWaiId int32, endTime int64, defendAllianceId int64) alliancescene.AllianceSceneData
	//获取九霄城战数据
	GetAllianceSceneData() alliancescene.AllianceSceneData
	//获取联盟霸主
	GetAllianceHegemon() *AllianceHegemonObject
	//赢的霸主
	AllianceWin(winAllianceId int64)
	//设置城战守方
	SetHegemonDefence(allianceId int64)
	//捐献
	Donate(memberId int64, typ alliancetypes.AllianceJuanXianType) (mem *AllianceMemberObject, err error)
	//捐献虎符
	DonateHuFu(memberId int64) (mem *AllianceMemberObject, err error)
	//获取仙盟
	GetAlliance(allianceId int64) (al *Alliance)
	//修改公告
	ChangeAllianceNotice(memberId int64, content string) (al *Alliance, err error)
	//后台GM修改公告
	GMChangeAllianceNotice(allianceId int64, content string) (al *Alliance, err error)
	//设置自动处理仙盟申请
	ChangeAutoAgreeJoinApply(memberId int64, isAuto int32) error
	//仙盟日志列表
	GetAllianceLogList(memberId int64) (logList []*AllianceLogObject, err error)
	//更新战力
	UpdateForce(memberId int64, force int64) (err error)
	//更新等级
	UpdateLevel(memberId int64, level int32) (err error)
	//GM初始化霸主
	GMAllianceHegemonReset() (err error)
	//添加仙盟镖车
	AddTransportation(memberId int64) (member *AllianceMemberObject, biaoChe *biaochenpc.BiaocheNPC, err error)
	//获取仙盟镖车次数
	GetAllianceTransportTimes(memberId int64) int32
	//更新领域
	UpdateLingYu(memberId int64, lingyuId int32) (err error)
	//同意仙盟邀请
	AgreeAllianceInvitation(memberId int64, beInvitationId int64, agree bool) (al *Alliance, err error)
	//更新转生
	UpdateZhuanSheng(memberId int64, zhuansheng int32) (err error)
	//更新斗神
	UpdateDouShenList() (err error)
	//定时检查开启城战
	CheckAllianceScene() (err error)
	//仙盟仓库存入
	SaveInDepot(allianceId int64, itemData *droptemplate.DropItemData, propertyData inventorytypes.ItemPropertyData) (al *Alliance, err error)
	//是否有足够的位置
	HasEnoughDepotSlot(allianceId int64, itemData *droptemplate.DropItemData, propertyData inventorytypes.ItemPropertyData) (flag bool)
	//仙盟仓库取出
	TakeOutDepot(allianceId int64, depotIndex, itemId, num int32) error
	//仙盟仓库整理
	MergeDepot(pl player.Player, indexList []int32) error
	//仙盟自动销毁
	AutoRemoveDepot(memId int64, isAuto int32, zhuansheng int32, quality itemtypes.ItemQualityType) (al *Alliance, err error)
	//获取入盟时间
	GetAllianceMemberHasedJionTime(memberId int64) int64
	//召唤仙盟boss
	AllianceSummonBoss(pl player.Player) (err error)
	//仙盟boss结束
	AllianceBossEnd(allianceId int64)
	//仙盟成员进入仙盟boss
	AllianceBossEnter(pl player.Player) (err error)
	//是否在仙盟boss
	AllianceBossScene(allianceId int64) scene.Scene
	//仙盟boss信息
	AllianceBossInfo(allianceId int64) (status alliancetypes.AllianceBossStatus, level int32, exp int32, times int64, err error)
	//仙盟boss增加经验
	AllianceBossAddExp(allianceId int64, exp int32)
	// 霸主雕像
	GetCurModelList() []scene.RobotPlayer //获取当前雕像
	AddModel(model scene.RobotPlayer)     //添加雕像
	RemoveModel(model scene.RobotPlayer)  //添加雕像
	//仙盟合并
	AllianceMerge(inviteAllianceId, allianceId int64) error
	//仙盟改名
	UpdateAllianceName(pl player.Player, newName string) error
	//合并申请
	IfAllianceCanMergeApply(allianceId int64) bool
	//合并邀请
	AllianceMergeInvite(allianceId int64, inviteAllianceId int64) (err error)
	//清除合并
	ClearAllianceMergeInvite(allianceId int64)
	//GM 仙盟boss
	GMAllianceBossReset(allianceId int64) (flag bool)
	//改变仙盟阵营
	AllianceChangeCamp(allianceId int64, campType chuangshitypes.ChuangShiCampType)
	GMDismissAlliance(allianceId int64) (al *Alliance, tempMemberList []*AllianceMemberObject, err error)
}

const (
	doushenRankHours = 1 //斗神殿列表更新间隔：2小时
)

type allianceMergeApply struct {
	inviteAllianceId int64
	inviteTime       int64
}

func (a *allianceMergeApply) IsCD(now int64, cd int64) bool {
	if now-a.inviteTime > cd {
		return false
	}
	return true
}

type allianceService struct {
	rwm                      sync.RWMutex
	allianceList             []*Alliance
	allianceNameMap          map[string]*Alliance
	allianceMemberMap        map[int64]*AllianceMemberObject
	allianceSceneData        alliancescene.AllianceSceneData
	allianceHegemonObject    *AllianceHegemonObject
	lastDoushenRankTime      int64
	allianceBossSceneDataMap map[int64]alliancebossscene.AllianceBossSceneData
	hegemonWinnerModelMap    map[int64]scene.RobotPlayer //获胜方主城雕像机器人
	runner                   heartbeat.HeartbeatTaskRunner
	//合盟申请
	allianceMergeApply map[int64]*allianceMergeApply
}

func (s *allianceService) GetAllianceList() []*Alliance {
	return s.allianceList
}

func (s *allianceService) GetAlliance(allianceId int64) *Alliance {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	_, al := s.getAlliance(allianceId)
	return al
}

func (s *allianceService) Start() {
	//初始化斗神殿列表
	s.updateDouShenList()

	//初始城战场景
	s.checkAllianceScene()

	// 加载霸主雕像
	gameevent.Emit(allianceeventtypes.EventTypeAllianceLoadWinnerModel, nil, nil)

	// //初始化仙盟阵营处理
	// s.initCamp()

	return
}

//心跳
func (s *allianceService) Heartbeat() {
	s.runner.Heartbeat()
}

func (s *allianceService) UpdateDouShenList() (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	return s.updateDouShenList()
}

//更新斗神殿列表
func (s *allianceService) updateDouShenList() (err error) {

	now := global.GetGame().GetTimeService().Now()

	if !s.ifCanUpdate(now) {
		return
	}

	for _, al := range s.allianceList {
		err = al.updateDouShenForceList(now)
		if err != nil {
			return
		}
	}

	s.lastDoushenRankTime = now

	return
}

//能否更新
func (s *allianceService) ifCanUpdate(timestamp int64) (flag bool) {
	lastTime := s.lastDoushenRankTime
	diffTime := timestamp - lastTime
	if diffTime >= int64(common.MINUTE*doushenRankHours) {
		flag = true
	}
	return
}

func (s *allianceService) CheckAllianceScene() (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	return s.checkAllianceScene()
}

//定时城战活动
func (s *allianceService) checkAllianceScene() (err error) {
	if s.allianceSceneData != nil {
		return
	}
	acTemp := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeAlliance)
	now := global.GetGame().GetTimeService().Now()
	openTime := global.GetGame().GetServerTime()
	mergeTime := merge.GetMergeService().GetMergeTime()
	timeTemplate, err := acTemp.GetActivityTimeTemplate(now, openTime, mergeTime)
	if err != nil {
		return
	}
	if timeTemplate == nil {
		return
	}

	endTime, err := timeTemplate.GetEndTime(now)
	if err != nil {
		return
	}
	s.createSceneData(acTemp.Mapid, endTime, 0)
	return
}

// //初始化仙盟阵营
// func (s *allianceService) initCamp() {
// 	var campTypeList []chuangshitypes.ChuangShiCampType
// 	for campType, _ := range chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCampTempAll() {
// 		if campType == chuangshitypes.ChuangShiCampTypeNone {
// 			continue
// 		}
// 		campTypeList = append(campTypeList, campType)
// 	}

// 	campLen := len(campTypeList)
// 	if campLen == 0 {
// 		return
// 	}

// 	now := global.GetGame().GetTimeService().Now()
// 	for index, al := range s.allianceList {
// 		if al.allianceObject.campType != chuangshitypes.ChuangShiCampTypeNone {
// 			continue
// 		}

// 		campIndex := (index + 1) % campLen
// 		al.allianceObject.campType = campTypeList[campIndex]
// 		al.allianceObject.updateTime = now
// 		al.allianceObject.SetModified()
// 	}
// }

//创建仙盟
// func (s *allianceService) CreateAlliance(campType chuangshitypes.ChuangShiCampType, createId int64, name string, role playertypes.RoleType, sex playertypes.SexType, plName string, createrVip int32, versionTyp alliancetypes.AllianceVersionType, typ alliancetypes.AllianceNewType) (al *Alliance, err error) {
func (s *allianceService) CreateAlliance(createId int64, name string, role playertypes.RoleType, sex playertypes.SexType, plName string, createrVip int32, createLingYuId int32, lev int32, versionTyp alliancetypes.AllianceVersionType, typ alliancetypes.AllianceNewType) (al *Alliance, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	//验证用户是否在仙盟内
	memberObj := s.getAllianceMember(createId)

	if memberObj != nil {
		err = errorAllianceUserAlreadyInAlliance
		return
	}

	//验证名字合法
	al = s.getAllianceByName(name)

	if al != nil {
		err = errorAllianceNameExist
		return
	}

	now := global.GetGame().GetTimeService().Now()
	alObj := createAllianceObject()
	id, err := idutil.GetId()
	if err != nil {
		return
	}
	alObj.id = id
	alObj.serverId = global.GetGame().GetServerIndex()
	alObj.originServerId = global.GetGame().GetServerIndex()
	alObj.name = name
	alObj.notice = ""
	alObj.level = 1
	if versionTyp == alliancetypes.AllianceVersionTypeNew && typ == alliancetypes.AllianceNewTypeHigh {
		alObj.level = 2
	}
	alObj.jianShe = 0
	alObj.huFu = 0
	alObj.totalForce = 0
	alObj.createTime = now
	alObj.isAutoAgree = 1
	alObj.maxRemoveZhuanSheng = 0
	alObj.maxRemoveQuality = itemtypes.ItemQualityTypeBlue
	alObj.mengzhuId = createId
	// alObj.campType = campType
	alObj.createId = createId
	alObj.SetModified()

	al = createAlliance(alObj)
	memberObj = createAllianceMemberObject(al)
	memId, _ := idutil.GetId()
	memberObj.id = memId
	memberObj.memberId = createId
	memberObj.name = plName
	memberObj.position = alliancetypes.AlliancePosition(alliancetypes.AlliancePositionMengZhu)
	memberObj.joinTime = now
	memberObj.level = lev
	memberObj.role = role
	memberObj.sex = sex
	memberObj.lingyuId = createLingYuId
	memberObj.createTime = now
	memberObj.SetModified()

	//初始化仙盟boss
	allianceBossObj := s.initAllianceBoss(al.GetAllianceId())
	al.allianceBossObject = allianceBossObj

	flag := s.addAllianceMember(memberObj)
	if !flag {
		panic("alliance:添加成员应该成功")
	}
	flag = s.addAlliance(al)
	if !flag {
		panic("alliance:添加仙盟应该成功")
	}

	content := lang.GetLangService().ReadLang(lang.AllianceCreateLog)
	log := fmt.Sprintf(content, coreutils.FormatColor(alliancetypes.ColorTypeLogName, plName))
	al.addLog(log)

	depotNum := alliancetemplate.GetAllianceTemplateService().GetAllianceTemplate(versionTyp, alObj.GetLevel()).UnionStorage
	al.depotBag = createDepotBag(alObj.GetId(), nil, depotNum)

	//加载斗神殿
	al.updateDouShenForceList(now)

	return
}

//解散仙盟
func (s *allianceService) DismissAlliance(memberId int64) (al *Alliance, tempMemberList []*AllianceMemberObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.allianceSceneData != nil {
		err = errorAllianceDismissCanNotInSceneOpened
		return
	}
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	al = memObj.GetAlliance()

	if memObj.position != alliancetypes.AlliancePositionMengZhu {
		err = errorAlliancePrivilegeNotEnough
		return
	}

	biaoChe := transpotation.GetTransportService().GetTransportation(memberId)
	if biaoChe != nil {
		err = errorAllianceDismissCanNotInTransportation
		return
	}

	allianceId := al.GetAllianceId()
	_, ok := s.allianceBossSceneDataMap[allianceId]
	if ok {
		err = errorAllianceDismissCanNotInBoss
		return
	}

	mengZhuId := al.GetAllianceMengZhuId()

	//删除所有成员
	now := global.GetGame().GetTimeService().Now()
	tempMemberList = make([]*AllianceMemberObject, len(al.GetMemberList()))
	copy(tempMemberList, al.GetMemberList())
	for _, mem := range tempMemberList {
		now := global.GetGame().GetTimeService().Now()
		mem.deleteTime = now
		mem.SetModified()

		s.removeAllianceMember(mem, true)
	}

	alObj := al.GetAllianceObject()
	alObj.deleteTime = now
	alObj.SetModified()
	s.removeAlliance(al)

	gameevent.Emit(allianceeventtypes.EventTypeAllianceDismiss, nil, mengZhuId)
	return
}

//申请加入仙盟
// func (s *allianceService) ApplyJoinAlliance(allianceId int64, joinId int64, campType chuangshitypes.ChuangShiCampType, name string, role playertypes.RoleType, sex playertypes.SexType, level int32, force int64) (al *Alliance, joinApplyObj *AllianceJoinApplyObject, err error) {
func (s *allianceService) ApplyJoinAlliance(allianceId int64, joinId int64, name string, role playertypes.RoleType, sex playertypes.SexType, level int32, force int64) (al *Alliance, joinApplyObj *AllianceJoinApplyObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	//仙盟不存在
	_, al = s.getAlliance(allianceId)
	if al == nil {
		err = errorAllianceNoExist
		return
	}

	// //阵营不同
	// if al.allianceObject.campType != campType {
	// 	err = errAllianceCampTypeNotSame
	// 	return
	// }

	//已经在仙盟内
	memObj := s.getAllianceMember(joinId)
	if memObj != nil {
		err = errorAllianceUserInAllianceToApply
		return
	}
	//判断是否已经满人
	if al.IsFull() {
		err = errorAllianceAlreadyFullApply
		return
	}

	if al.GetIsAutoAgree() == 1 {
		s.joinAlliance(al, level, name, force, joinId, role, sex, playertypes.PlayerOnlineStateOnline, 0, 0, 0, 0)
		return
	}

	//获取加入申请
	now := global.GetGame().GetTimeService().Now()
	_, joinApplyObj = al.getApply(joinId)
	if joinApplyObj != nil {
		if joinApplyObj.isApplyCD() {
			err = errorAllianceAlreadyApply
			return
		}
		joinApplyObj.updateTime = now
		joinApplyObj.SetModified()
		return
	}
	//添加申请
	joinApplyObj = createAllianceJoinApplyObject()
	applyId, _ := idutil.GetId()
	joinApplyObj.id = applyId
	joinApplyObj.allianceId = allianceId
	joinApplyObj.name = name
	joinApplyObj.sex = sex
	joinApplyObj.force = force
	joinApplyObj.level = level
	joinApplyObj.role = role
	joinApplyObj.joinId = joinId
	joinApplyObj.createTime = now
	joinApplyObj.SetModified()

	flag := al.addApply(joinApplyObj)
	if !flag {
		panic("alliance:添加申请应该成功")
	}
	return
}

//同意申请加入仙盟
func (s *allianceService) AgreeAllianceJoinApply(agreeId int64, joinId int64, agree bool) (al *Alliance, applyObj *AllianceJoinApplyObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	managerObj := s.getAllianceMember(agreeId)
	if managerObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	if managerObj.GetPosition() == alliancetypes.AlliancePositionMember {
		err = errorAlliancePrivilegeNotEnough
		return
	}

	// TODO xzk27 申请列表新增阵营类型

	al = managerObj.GetAlliance()
	now := global.GetGame().GetTimeService().Now()

	//申请id不存在
	_, applyObj = al.getApply(joinId)
	if applyObj == nil {
		err = errorAllianceApplyNoExist
		return
	}

	//已经在仙盟内
	memObj := s.getAllianceMember(joinId)
	if memObj != nil {
		err = errorAllianceUserAlreadyInAlliance
		goto AfterAgree
	}

	//判断是否已经满人
	if al.IsFull() {
		err = errorAllianceAlreadyFullApply
		goto AfterAgree
	}

	if agree {
		s.joinAlliance(al, applyObj.GetLevel(), applyObj.GetName(), applyObj.GetForce(), joinId, applyObj.GetRole(), applyObj.GetSex(), playertypes.PlayerOnlineStateOnline, 0, 0, 0, 0)
	}

AfterAgree:
	applyObj.deleteTime = now
	applyObj.SetModified()
	al.removeApply(applyObj)
	return
}

func (s *allianceService) joinAlliance(al *Alliance, level int32, applyName string, force, applyId int64, role playertypes.RoleType, sex playertypes.SexType, status playertypes.PlayerOnlineState, lastLogoutTime int64, lingyuId int32, zhuangSheng int32, vip int32) {
	now := global.GetGame().GetTimeService().Now()
	memberObj := createAllianceMemberObject(al)
	memId, _ := idutil.GetId()
	memberObj.id = memId
	memberObj.level = level
	memberObj.name = applyName
	memberObj.force = force
	memberObj.role = role
	memberObj.sex = sex
	memberObj.lingyuId = lingyuId
	memberObj.gongXian = 0
	memberObj.memberId = applyId
	memberObj.position = alliancetypes.AlliancePosition(alliancetypes.AlliancePositionMember)
	memberObj.joinTime = now
	memberObj.onlineStatus = status
	memberObj.zhuanSheng = zhuangSheng
	memberObj.vip = vip
	memberObj.lastLogoutTime = lastLogoutTime
	memberObj.createTime = now
	memberObj.SetModified()
	flag := s.addAllianceMember(memberObj)
	if !flag {
		panic("alliance:仙盟添加成员应该成功")
	}
	al.updateForce()

	content := lang.GetLangService().ReadLang(lang.AllianceJoinLog)
	log := fmt.Sprintf(content, coreutils.FormatColor(alliancetypes.ColorTypeLogName, applyName))
	al.addLog(log)
}

//获取仙盟成员
func (s *allianceService) GetAllianceMember(memberId int64) (memObj *AllianceMemberObject) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	memObj = s.getAllianceMember(memberId)
	return
}

//获取仙盟成员
func (s *allianceService) getAllianceMember(memberId int64) (memObj *AllianceMemberObject) {
	memObj, exist := s.allianceMemberMap[memberId]
	if !exist {
		return
	}
	return
}

func (s *allianceService) addAlliance(al *Alliance) (flag bool) {
	_, tal := s.getAlliance(al.GetAllianceId())
	if tal != nil {
		return
	}
	tal = s.getAllianceByName(al.GetAllianceObject().GetName())
	if tal != nil {
		return
	}
	flag = true
	s.allianceList = append(s.allianceList, al)
	s.allianceNameMap[al.GetAllianceObject().GetName()] = al
	//TODO 排序
	return
}

func (s *allianceService) removeAlliance(al *Alliance) {
	index, al := s.getAlliance(al.GetAllianceId())
	if al == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	al.allianceObject.deleteTime = now
	al.allianceObject.SetModified()
	s.allianceList = append(s.allianceList[:index], s.allianceList[index+1:]...)
	delete(s.allianceNameMap, al.GetAllianceObject().GetName())
}

func (s *allianceService) getAllianceByName(name string) (a *Alliance) {
	a, exist := s.allianceNameMap[name]
	if !exist {
		return
	}
	return
}

func (s *allianceService) getAlliance(allianceId int64) (index int32, al *Alliance) {
	for i, al := range s.allianceList {
		if al.GetAllianceId() == allianceId {
			return int32(i), al
		}
	}
	return -1, nil

}

func (s *allianceService) init() (err error) {
	//添加定时任务
	s.runner = heartbeat.NewHeartbeatTaskRunner()
	s.runner.AddTask(CreateAllianceSceneTask(s))
	s.runner.AddTask(CreateDouShenTask(s))

	ds := global.GetGame().GetDB()
	rs := global.GetGame().GetRedisService()
	//初始化dao
	err = dao.Init(ds, rs)
	if err != nil {
		return
	}

	s.allianceList = make([]*Alliance, 0, 16)
	s.allianceNameMap = make(map[string]*Alliance)
	s.allianceMemberMap = make(map[int64]*AllianceMemberObject)
	s.allianceBossSceneDataMap = make(map[int64]alliancebossscene.AllianceBossSceneData)
	s.hegemonWinnerModelMap = make(map[int64]scene.RobotPlayer)
	s.allianceMergeApply = make(map[int64]*allianceMergeApply)
	//TODO 优化加载部分帮派
	//加载所有帮派
	err = s.initAlliances()
	if err != nil {
		return
	}
	//初始化霸主
	err = s.initHegemon()
	if err != nil {
		return
	}

	return nil
}

func (s *allianceService) initAlliances() (err error) {
	allianceEntityList, err := dao.GetAllianceDao().GetAllAllianceList()
	if err != nil {
		return
	}
	//加载所有仙盟
	for _, allianceEntity := range allianceEntityList {
		allianceObject := createAllianceObject()
		allianceObject.FromEntity(allianceEntity)
		al := createAlliance(allianceObject)

		//加载成员
		memberEntityList, err := dao.GetAllianceDao().GetAllianceMemberList(allianceObject.GetId())
		if err != nil {
			return err
		}
		for _, memberEntity := range memberEntityList {
			memberObj := createAllianceMemberObject(al)
			memberObj.FromEntity(memberEntity)
			s.addAllianceMember(memberObj)
		}
		//加载申请请求
		joinApplyEntityList, err := dao.GetAllianceDao().GetAllianceJoinApplyList(allianceObject.GetId())
		if err != nil {
			return err
		}
		for _, joinApplyEntity := range joinApplyEntityList {
			joinApplyObj := createAllianceJoinApplyObject()
			joinApplyObj.FromEntity(joinApplyEntity)
			al.addApply(joinApplyObj)
		}

		//加载仙盟日志
		maxLogSize := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAlliancelogLimit)
		logList, err := dao.GetAllianceDao().GetPlayerAllianceLogList(allianceObject.GetId(), maxLogSize)
		if err != nil {
			return err
		}
		for _, log := range logList {
			o := createAllianceLogObject()
			o.FromEntity(log)
			al.appendLog(o)
		}

		// 加载邀请列表
		invitationEntityList, err := dao.GetAllianceDao().GetAllianceInvitationList(allianceObject.GetId())
		if err != nil {
			return err
		}
		for _, invitationEntity := range invitationEntityList {
			invitationObj := createAllianceInvitationObject()
			invitationObj.FromEntity(invitationEntity)
			al.addInvitation(invitationObj)
		}
		// 加载仓库
		depotItemEntityList, err := dao.GetAllianceDao().GetAllianceDepotItemList(allianceObject.GetId())
		if err != nil {
			return err
		}
		var depotItemList []*AllianceDepotItemObject
		for _, entity := range depotItemEntityList {
			depotObj := createAllianceDepotItemObject()
			depotObj.FromEntity(entity)
			depotItemList = append(depotItemList, depotObj)
		}

		// 判断创建仙盟版本
		versionType := checkAllianceVersion()
		depotNum := alliancetemplate.GetAllianceTemplateService().GetAllianceTemplate(versionType, al.GetAllianceLevel()).UnionStorage
		s.fixUpstarLevel(depotItemList)
		al.depotBag = createDepotBag(allianceObject.GetId(), depotItemList, depotNum)

		//加载仙盟boss
		err = s.initLoadAllianceBoss(al, allianceObject.GetId())
		if err != nil {
			return err
		}

		//添加仙盟到内存
		s.addAlliance(al)
	}
	return
}

func (s *allianceService) fixUpstarLevel(items []*AllianceDepotItemObject) {
	for _, itemObj := range items {
		if itemObj.IsEmpty() {
			continue
		}

		goldequipData, ok := itemObj.propertyData.(*goldequiptypes.GoldEquipPropertyData)
		if !ok {
			continue
		}
		itemTemp := item.GetItemService().GetItem(int(itemObj.itemId))
		if itemTemp.GetGoldEquipTemplate() == nil {
			log.Info("itemid:", itemObj.itemId)
			continue
		}
		maxLeve := itemTemp.GetGoldEquipTemplate().GetMaxUpstarLevel()
		goldequipData.FixUpstarLevel(maxLeve)
		itemObj.SetModified()
	}
}

// 加载仙盟boss
func (s *allianceService) initLoadAllianceBoss(al *Alliance, allianceId int64) (err error) {

	bossEntity, err := dao.GetAllianceDao().GetAllianceBoss(allianceId)
	if err != nil {
		return err
	}
	if bossEntity == nil {
		s.initAllianceBoss(allianceId)
	} else {
		al.allianceBossObject = createAllianceBossObject()
		al.allianceBossObject.FromEntity(bossEntity)
	}
	return
}

func (s *allianceService) initHegemon() (err error) {
	hegemonEntity, err := dao.GetAllianceDao().GetAllianceHegemon()
	if err != nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	hegemonObj := createAllianceHegemonObject()
	s.allianceHegemonObject = hegemonObj
	if hegemonEntity == nil {
		hegemonObj.id, err = idutil.GetId()
		if err != nil {
			return
		}
		hegemonObj.serverId = global.GetGame().GetServerIndex()
		hegemonObj.createTime = now
		hegemonObj.SetModified()

	} else {
		hegemonObj.FromEntity(hegemonEntity)
		//可能联盟解散了
		currentAllianceId := hegemonObj.GetAllianceId()
		if currentAllianceId != 0 {
			_, al := s.getAlliance(currentAllianceId)
			if al != nil {
				return
			}
			hegemonObj.allianceId = 0
			hegemonObj.winNum = 0
			hegemonObj.updateTime = now
			hegemonObj.SetModified()
		}

		defenceAllianceId := hegemonObj.GetDefenceAllianceId()
		if defenceAllianceId != 0 {
			_, al := s.getAlliance(defenceAllianceId)
			if al != nil {
				return
			}
			hegemonObj.defenceAllianceId = 0
			hegemonObj.updateTime = now
			hegemonObj.SetModified()
		}
	}
	return
}

func (s *allianceService) addAllianceMember(memberObj *AllianceMemberObject) (flag bool) {
	tmemObj := s.getAllianceMember(memberObj.GetMemberId())
	if tmemObj != nil {
		return
	}
	al := memberObj.GetAlliance()
	if al == nil {
		return
	}

	flag = al.addMember(memberObj)
	if !flag {
		return
	}

	s.allianceMemberMap[memberObj.GetMemberId()] = memberObj
	return
}

func (s *allianceService) removeAllianceMember(memberObj *AllianceMemberObject, isClearPlayerData bool) {
	al := memberObj.GetAlliance()
	if al != nil {
		al.removeMember(memberObj, isClearPlayerData)
	}
	delete(s.allianceMemberMap, memberObj.GetMemberId())
}

func (s *allianceService) GetJoinApplyList(memberId int64) (applyObjList []*AllianceJoinApplyObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	managerObj := s.getAllianceMember(memberId)
	if managerObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	if managerObj.GetPosition() == alliancetypes.AlliancePositionMember {
		err = errorAlliancePrivilegeNotEnough
		return
	}
	al := managerObj.GetAlliance()
	applyObjList = al.GetApplyList()
	//TODO 刷新申请列表
	return
}

func (s *allianceService) SyncMemberInfo(memberId int64, name string, sex playertypes.SexType, level int32, force int64, zhuansheng int32, lingyuId int32, vip int32) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	memObj.onlineStatus = playertypes.PlayerOnlineStateOnline
	memObj.level = level
	memObj.name = name
	memObj.sex = sex
	memObj.zhuanSheng = zhuansheng
	memObj.lingyuId = lingyuId
	memObj.vip = vip
	if force != memObj.force {
		memObj.force = force
		al := memObj.GetAlliance()
		al.updateForce()
	}
	memObj.SetModified()
	return
}

func (s *allianceService) OfflineMember(memberId int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	now := global.GetGame().GetTimeService().Now()
	memObj.onlineStatus = playertypes.PlayerOnlineStateOffline
	memObj.lastLogoutTime = now
	memObj.SetModified()
	return
}

func (s *allianceService) OnlineMember(memberId int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	memObj.SetModified()
	return
}

func (s *allianceService) ExitAlliance(memberId int64) (al *Alliance, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.allianceSceneData != nil {
		err = errorAllianceDismissCanNotInSceneOpened
		return
	}
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	if memObj.position == alliancetypes.AlliancePositionMengZhu {
		err = errorAllianceMengZhuNoExit
		return
	}
	now := global.GetGame().GetTimeService().Now()
	memObj.deleteTime = now
	memObj.SetModified()
	s.removeAllianceMember(memObj, true)
	al = memObj.GetAlliance()
	al.updateForce()

	content := lang.GetLangService().ReadLang(lang.AllianceExitLog)
	log := fmt.Sprintf(content, coreutils.FormatColor(alliancetypes.ColorTypeLogName, memObj.GetName()))
	al.addLog(log)

	return
}

//踢人
func (s *allianceService) Kick(memberId int64, kickMemberId int64) (memObj *AllianceMemberObject, kickMemObj *AllianceMemberObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.allianceSceneData != nil {
		err = errorAllianceDismissCanNotInSceneOpened
		return
	}
	memObj = s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	kickMemObj = s.getAllianceMember(kickMemberId)
	if kickMemObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	if memberId == kickMemberId {
		err = errorAllianceUserCanNotKickSelf
		return
	}

	al := memObj.GetAlliance()
	if kickMemObj.GetAlliance() != al {
		err = errorAllianceNotInSameAlliance
		return
	}

	if memObj.GetPosition() != alliancetypes.AlliancePositionMengZhu {
		if kickMemObj.GetPosition() != alliancetypes.AlliancePositionMember {
			err = errorAlliancePrivilegeNotEnough
			return
		}
	}

	now := global.GetGame().GetTimeService().Now()
	kickMemObj.deleteTime = now
	kickMemObj.SetModified()
	s.removeAllianceMember(kickMemObj, true)

	al.updateForce()

	positionName := coreutils.FormatColor(alliancetypes.ColorTypePosition, memObj.GetPosition().String())
	kickName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, memObj.GetName())
	beKickName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, kickMemObj.GetName())
	content := lang.GetLangService().ReadLang(lang.AllianceKickLog)
	log := fmt.Sprintf(content, positionName, kickName, beKickName)
	al.addLog(log)

	return
}

//任命
func (s *allianceService) Commit(memberId, commitMemberId int64, position alliancetypes.AlliancePosition) (memObj *AllianceMemberObject, commitMemObj *AllianceMemberObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj = s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	commitMemObj = s.getAllianceMember(commitMemberId)
	if commitMemObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	if memberId == commitMemberId {
		err = errorAllianceUserCanNotCommitSelf
		return
	}
	if memObj.GetPosition() != alliancetypes.AlliancePositionMengZhu {
		err = errorAllianceUserNotMengZhu
		return
	}
	al := memObj.GetAlliance()
	if commitMemObj.GetAlliance() != al {
		err = errorAllianceNotInSameAlliance
		return
	}

	//职位是否有空缺
	if position != alliancetypes.AlliancePositionMember {
		// 判断创建仙盟版本
		versionType := checkAllianceVersion()
		alTemp := alliancetemplate.GetAllianceTemplateService().GetAllianceTemplate(versionType, al.allianceObject.GetLevel())
		positionLimitNum := alTemp.GetAlliancePositionNum(position)
		positionNum := al.GetNumOfManagers(position)

		if positionNum >= positionLimitNum {
			err = errorAlliancePositionAlreadyFull
			return
		}
	}

	//普通成员
	if commitMemObj.GetPosition() == position && position == alliancetypes.AlliancePositionMember {
		err = errorAllianceAlreadyMember
		return
	}

	content := ""
	positionName := ""
	if position == alliancetypes.AlliancePositionMember {
		content = lang.GetLangService().ReadLang(lang.AllianceReleaseLog)
		positionName = commitMemObj.position.String()
	} else {
		content = lang.GetLangService().ReadLang(lang.AllianceCommitLog)
		positionName = position.String()
	}

	positionName = coreutils.FormatColor(alliancetypes.ColorTypePosition, positionName)
	commitName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, memObj.GetName())
	beCommitName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, commitMemObj.GetName())
	log := fmt.Sprintf(content, commitName, beCommitName, positionName)
	al.addLog(log)
	now := global.GetGame().GetTimeService().Now()
	commitMemObj.position = position
	commitMemObj.updateTime = now
	commitMemObj.SetModified()

	al.managerChanged(commitMemObj)

	return
}

//转让
func (s *allianceService) Transfer(memberId, transferMemberId int64) (memObj *AllianceMemberObject, transferMemObj *AllianceMemberObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if memberId == transferMemberId {
		return
	}

	memObj = s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	transferMemObj = s.getAllianceMember(transferMemberId)
	if transferMemObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	al := memObj.GetAlliance()
	if transferMemObj.GetAlliance() != al {
		err = errorAllianceNotInSameAlliance
		return
	}

	if memObj.GetPosition() != alliancetypes.AlliancePositionMengZhu {
		err = errorAlliancePrivilegeNotEnough
		return
	}

	biaoChe := transpotation.GetTransportService().GetTransportation(memberId)
	if biaoChe != nil {
		err = errorAllianceDismissCanNotInTransportation
		return
	}

	now := global.GetGame().GetTimeService().Now()
	al.GetAllianceObject().mengzhuId = transferMemberId
	al.GetAllianceObject().updateTime = now
	al.GetAllianceObject().SetModified()
	memObj.position = alliancetypes.AlliancePositionMember
	memObj.updateTime = now
	transferMemObj.position = alliancetypes.AlliancePositionMengZhu
	transferMemObj.updateTime = now
	memObj.SetModified()
	transferMemObj.SetModified()

	//TODO 优化
	al.managerChanged(transferMemObj)
	al.managerChanged(memObj)

	content := lang.GetLangService().ReadLang(lang.AllianceTransferLog)
	log := fmt.Sprintf(content, coreutils.FormatColor(alliancetypes.ColorTypeLogName, memObj.GetName()), coreutils.FormatColor(alliancetypes.ColorTypeLogName, transferMemObj.GetName()))
	al.addLog(log)

	//盟主变更
	gameevent.Emit(allianceeventtypes.EventTypeAllianceMengzhuChanged, al, memberId)

	return
}

//弹劾
func (s *allianceService) Impeach(impeachId int64) (impeachMemObj *AllianceMemberObject, mengZhuMemObj *AllianceMemberObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	impeachMemObj = s.getAllianceMember(impeachId)
	if impeachMemObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	al := impeachMemObj.GetAlliance()
	mengZhuId := al.GetAllianceObject().mengzhuId
	if mengZhuId == impeachId {
		err = errorAllianceUserCanNotImpeachSelf
		return
	}
	mengZhuMemObj = s.getAllianceMember(mengZhuId)
	if mengZhuMemObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	if mengZhuMemObj.GetPosition() != alliancetypes.AlliancePositionMengZhu {
		err = errorAllianceUserNotMengZhu
		return
	}

	//是否离线一周
	if !mengZhuMemObj.IsOfflineOneWeek() {
		err = errorAllianceImpeachConditionNotEnough
		return
	}

	//TODO 优化
	now := global.GetGame().GetTimeService().Now()
	impeachMemObj.position = alliancetypes.AlliancePositionMengZhu
	impeachMemObj.updateTime = now
	mengZhuMemObj.position = alliancetypes.AlliancePositionMember
	mengZhuMemObj.updateTime = now
	impeachMemObj.SetModified()
	mengZhuMemObj.SetModified()
	al.GetAllianceObject().mengzhuId = impeachId
	al.GetAllianceObject().SetModified()

	al.managerChanged(impeachMemObj)
	al.managerChanged(mengZhuMemObj)

	content := lang.GetLangService().ReadLang(lang.AllianceImpeachLog)
	log := fmt.Sprintf(content, coreutils.FormatColor(alliancetypes.ColorTypeLogName, impeachMemObj.GetName()))
	al.addLog(log)

	//盟主变更
	gameevent.Emit(allianceeventtypes.EventTypeAllianceMengzhuChanged, al, mengZhuId)
	return
}

//邀请
func (s *allianceService) Invitation(memberId, beInvitationId int64, name string, role playertypes.RoleType, sex playertypes.SexType, level int32, force int64) (al *Alliance, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	//不是仙盟成员
	member := s.getAllianceMember(memberId)
	if member == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	//已经在仙盟内
	invitationMemObj := s.getAllianceMember(beInvitationId)
	if invitationMemObj != nil {
		err = errorAllianceUserAlreadyInAlliance
		return
	}

	_, al = s.getAlliance(member.GetAllianceId())
	//判断是否已经满人
	if al.IsFull() {
		err = errorAllianceAlreadyFullInvitation
		return
	}

	//已经邀请过了
	_, invitationObj := al.getInvitation(beInvitationId)
	if invitationObj != nil {
		return
	}

	invitationObj = createAllianceInvitationObject()
	id, _ := idutil.GetId()
	invitationObj.id = id
	invitationObj.allianceId = al.GetAllianceId()
	invitationObj.invitationId = beInvitationId
	invitationObj.name = name
	invitationObj.sex = sex
	invitationObj.force = force
	invitationObj.level = level
	invitationObj.role = role
	now := global.GetGame().GetTimeService().Now()
	invitationObj.createTime = now
	invitationObj.SetModified()

	flag := al.addInvitation(invitationObj)
	if !flag {
		panic("alliance:邀请应该成功")
	}

	return
}

func (s *allianceService) getHighestAlliance() *Alliance {
	var maxAlliance *Alliance
	for _, al := range s.allianceList {
		if maxAlliance == nil {
			maxAlliance = al
			continue
		}
		if maxAlliance.GetAllianceObject().GetTotalForce() < al.GetAllianceObject().GetTotalForce() {
			maxAlliance = al
			continue
		}
	}
	return maxAlliance
}

func (s *allianceService) CreateAllianceSceneData(chengWaiId int32, endTime int64, defendAllianceId int64) (d alliancescene.AllianceSceneData) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if s.allianceSceneData != nil {
		d = s.allianceSceneData
		return
	}

	s.createSceneData(chengWaiId, endTime, defendAllianceId)
	return s.allianceSceneData
}

func (s *allianceService) createSceneData(chengWaiId int32, endTime int64, defendAllianceId int64) {

	currentHegemonAllianceId := defendAllianceId
	if currentHegemonAllianceId == 0 {
		currentHegemonAllianceId = s.allianceHegemonObject.defenceAllianceId
		if currentHegemonAllianceId == 0 {
			al := s.getHighestAlliance()
			if al == nil {
				return
			}
			currentHegemonAllianceId = al.GetAllianceId()
		}
	}
	_, al := s.getAlliance(currentHegemonAllianceId)
	if al == nil {
		al := s.getHighestAlliance()
		if al == nil {
			return
		}
		currentHegemonAllianceId = al.GetAllianceId()
	}
	_, al = s.getAlliance(currentHegemonAllianceId)
	if al == nil {
		return
	}

	allianceId := al.GetAllianceId()
	allianceName := al.GetAllianceName()
	allianceHuFu := al.GetAllianceObject().GetHuFu()
	allianceSceneData := alliancescene.CreateAllianceSceneData(chengWaiId, allianceId, allianceName, allianceHuFu, endTime)
	s.allianceSceneData = allianceSceneData
}

func (s *allianceService) GetAllianceSceneData() alliancescene.AllianceSceneData {
	return s.allianceSceneData
}

func (s *allianceService) GetAllianceHegemon() *AllianceHegemonObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	return s.allianceHegemonObject
}

func (s *allianceService) AllianceWin(winAllianceId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	_, al := s.getAlliance(winAllianceId)
	if al == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if s.allianceHegemonObject.allianceId == winAllianceId {
		s.allianceHegemonObject.winNum += 1
		s.allianceHegemonObject.updateTime = now
	} else {
		s.allianceHegemonObject.winNum = 1
		s.allianceHegemonObject.allianceId = winAllianceId
		s.allianceHegemonObject.updateTime = now
	}
	s.allianceHegemonObject.SetModified()

	s.allianceSceneData = nil
}

func (s *allianceService) SetHegemonDefence(allianceId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	_, al := s.getAlliance(allianceId)
	if al == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	s.allianceHegemonObject.defenceAllianceId = allianceId
	s.allianceHegemonObject.updateTime = now
	s.allianceHegemonObject.SetModified()
}

func (s *allianceService) Donate(memberId int64, typ alliancetypes.AllianceJuanXianType) (member *AllianceMemberObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	member = s.getAllianceMember(memberId)
	if member == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	unionDonateTemplate := alliancetemplate.GetAllianceTemplateService().GetUnionDonateTemplate(typ)
	if unionDonateTemplate == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	member.gongXian += int64(unionDonateTemplate.DonateContribution)
	member.updateTime = now
	member.SetModified()
	member.GetAlliance().AddJianShe(int64(unionDonateTemplate.DonateBuild))

	//日志
	donateTemplate := alliancetemplate.GetAllianceTemplateService().GetUnionDonateTemplate(typ)
	var logCode lang.LangCode
	if typ == alliancetypes.AllianceJuanXianTypeLingPai {
		logCode = lang.AllianceDonateItemLog
	} else {
		logCode = lang.AllianceDonateResouceLog
	}
	content := lang.GetLangService().ReadLang(logCode)
	donateName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, member.GetName())
	numStr := coreutils.FormatColor(alliancetypes.ColorTypeLogDonateNum, fmt.Sprint(donateTemplate.GetJuanXianNum()))
	log := fmt.Sprintf(content, donateName, numStr, typ.String())
	member.GetAlliance().addLog(log)

	return
}

func (s *allianceService) DonateHuFu(memberId int64) (member *AllianceMemberObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	member = s.getAllianceMember(memberId)
	if member == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	member.GetAlliance().increaseHuFu()

	return
}

func (s *allianceService) ChangeAllianceNotice(memberId int64, content string) (al *Alliance, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	if memObj.GetPosition() != alliancetypes.AlliancePositionMengZhu {
		err = errorAlliancePrivilegeNotEnough
		return
	}

	al = memObj.GetAlliance()
	al.allianceObject.notice = content
	al.allianceObject.SetModified()

	return
}

func (s *allianceService) GMChangeAllianceNotice(allianceId int64, content string) (al *Alliance, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	_, al = s.getAlliance(allianceId)
	if al == nil {
		err = errorAllianceNoExist
	}

	al.allianceObject.notice = content
	al.allianceObject.SetModified()

	return
}

func (s *allianceService) ChangeAutoAgreeJoinApply(memberId int64, isAuto int32) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	if memObj.GetPosition() != alliancetypes.AlliancePositionMengZhu {
		err = errorAlliancePrivilegeNotEnough
		return
	}

	now := global.GetGame().GetTimeService().Now()
	al := memObj.GetAlliance()
	al.allianceObject.isAutoAgree = isAuto
	al.allianceObject.updateTime = now
	al.allianceObject.SetModified()
	return
}

func (s *allianceService) GetAllianceLogList(memberId int64) (logList []*AllianceLogObject, err error) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	logList = memObj.GetAlliance().GetLogList()

	return logList, nil
}

func (s *allianceService) UpdateForce(memberId int64, force int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	now := global.GetGame().GetTimeService().Now()
	memObj.force = force
	memObj.updateTime = now
	memObj.SetModified()
	memObj.GetAlliance().updateForce()
	return
}

func (s *allianceService) UpdateLevel(memberId int64, level int32) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	now := global.GetGame().GetTimeService().Now()
	memObj.level = level
	memObj.updateTime = now
	memObj.SetModified()

	return
}
func (s *allianceService) UpdateZhuanSheng(memberId int64, zhuansheng int32) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	now := global.GetGame().GetTimeService().Now()
	memObj.zhuanSheng = zhuansheng
	memObj.updateTime = now
	memObj.SetModified()

	return
}

func (s *allianceService) UpdateLingYu(memberId int64, lingyuId int32) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}
	now := global.GetGame().GetTimeService().Now()
	memObj.lingyuId = lingyuId
	memObj.updateTime = now
	memObj.SetModified()

	//斗神殿成员
	al := memObj.GetAlliance()
	for _, member := range al.GetDouShenList() {
		if member.GetMemberId() == memberId {
			gameevent.Emit(allianceeventtypes.EventTypeAllianceDouShenMemberLingYuChanged, al, member)
			break
		}
	}

	return
}

func (s *allianceService) AddTransportation(memberId int64) (member *AllianceMemberObject, biaoChe *biaochenpc.BiaocheNPC, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	member = s.getAllianceMember(memberId)
	if member == nil {
		err = errorAllianceTransportationNotInAlliance
		return
	}
	al := member.GetAlliance()
	if al == nil {
		err = errorAllianceNoExist
		return
	}

	if !member.IsMengZhu() {
		err = errorAllianceTransportationPositionNoEnough
		return
	}

	//次数判断
	if !al.hasEnoughTransportTimes() {
		err = errorAllianceTransportationAcceptNumNoEnough
		return
	}

	//消耗次数
	al.updateTransportTimes()

	biaoChe, err = transpotation.GetTransportService().AddAllianceTransportation(memberId, member.GetAllianceId(), al.GetAllianceName())
	return
}

func (s *allianceService) GetAllianceTransportTimes(memberId int64) int32 {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	member := s.getAllianceMember(memberId)
	if member == nil {
		return 0
	}
	al := member.GetAlliance()
	if al == nil {
		return 0
	}

	if !member.IsMengZhu() {
		return 0
	}

	al.refreshTimes()

	return al.allianceObject.transportTimes
}

func (s *allianceService) GMAllianceHegemonReset() (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	hegemonObj := s.allianceHegemonObject
	now := global.GetGame().GetTimeService().Now()

	hegemonObj.allianceId = 0
	hegemonObj.defenceAllianceId = 0
	hegemonObj.winNum = 0
	hegemonObj.updateTime = now
	hegemonObj.SetModified()

	return
}

func (s *allianceService) AgreeAllianceInvitation(memberId int64, beInvitationId int64, agree bool) (al *Alliance, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()

	managerObj := s.getAllianceMember(memberId)
	if managerObj == nil {
		err = errorAllianceUserNotInAlliance
	}

	al = managerObj.GetAlliance()
	//邀请id不存在
	_, beInvitationObj := al.getInvitation(beInvitationId)
	if beInvitationObj == nil {
		err = errorAllianceInvitationNoExist
		return
	}

	//已经在仙盟内
	beInvitationMemberObj := s.getAllianceMember(beInvitationId)
	if beInvitationMemberObj != nil {
		err = errorAllianceUserAlreadyInAlliance
		goto AfterAgree
	}

	//判断是否已经满人
	if al.IsFull() {
		err = errorAllianceAlreadyFullApply
		goto AfterAgree
	}

	if agree {
		memberObj := createAllianceMemberObject(al)
		memId, _ := idutil.GetId()
		memberObj.id = memId
		memberObj.level = beInvitationObj.GetLevel()
		memberObj.name = beInvitationObj.GetName()
		memberObj.force = beInvitationObj.GetForce()
		memberObj.role = beInvitationObj.GetRole()
		memberObj.sex = beInvitationObj.GetSex()
		memberObj.gongXian = 0
		memberObj.memberId = beInvitationId
		memberObj.position = alliancetypes.AlliancePosition(alliancetypes.AlliancePositionMember)
		memberObj.joinTime = now
		memberObj.createTime = now
		memberObj.SetModified()
		flag := s.addAllianceMember(memberObj)
		if !flag {
			panic("alliance:仙盟添加成员应该成功")
		}

		//更新仙盟总战力
		al.updateForce()

		content := lang.GetLangService().ReadLang(lang.AllianceJoinLog)
		log := fmt.Sprintf(content, coreutils.FormatColor(alliancetypes.ColorTypeLogName, beInvitationObj.GetName()))
		al.addLog(log)
	}

AfterAgree:
	beInvitationObj.deleteTime = now
	beInvitationObj.SetModified()
	al.removeInvitation(beInvitationObj)
	return
}

func (s *allianceService) SaveInDepot(allianceId int64, itemData *droptemplate.DropItemData, propertyData inventorytypes.ItemPropertyData) (al *Alliance, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	_, al = s.getAlliance(allianceId)
	if al == nil {
		err = errorAllianceNoExist
		return
	}

	itemId := itemData.GetItemId()
	num := itemData.GetNum()
	level := itemData.GetLevel()
	bind := itemData.GetBindType()

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return nil, fmt.Errorf("alliance:模板不存在")
	}

	if al.allianceObject.isAutoRemoveDepot == 1 {
		if itemTemplate.GetQualityType() <= al.allianceObject.maxRemoveQuality && itemTemplate.NeedZhuanShu <= al.allianceObject.maxRemoveZhuanSheng {
			return
		}
	}

	// 空间
	if !al.hasEnoughDepotSlot(itemData) {
		err = errorAllianceDepotSlotNoEnough
		return
	}

	//保存
	flag := al.depotBag.addLevelItem(itemId, num, level, bind, propertyData)
	if !flag {
		panic(fmt.Errorf("Alliance:保存到仙盟仓库应该成功"))
	}
	changedItemList := al.depotBag.getChangedSlotAndReset()
	gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotChanged, al, changedItemList)
	return
}

func (s *allianceService) HasEnoughDepotSlot(allianceId int64, itemData *droptemplate.DropItemData, propertyData inventorytypes.ItemPropertyData) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	_, al := s.getAlliance(allianceId)
	if al == nil {
		return
	}

	itemId := itemData.GetItemId()

	itemTemplate := item.GetItemService().GetItem(int(itemId))
	if itemTemplate == nil {
		return
	}

	if al.allianceObject.isAutoRemoveDepot == 1 {
		if itemTemplate.GetQualityType() <= al.allianceObject.maxRemoveQuality && itemTemplate.NeedZhuanShu <= al.allianceObject.maxRemoveZhuanSheng {
			return true
		}
	}

	// 空间
	if !al.hasEnoughDepotSlot(itemData) {
		return
	}

	return true
}

func (s *allianceService) TakeOutDepot(allianceId int64, depotIndex, itemId, num int32) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	if num <= 0 || depotIndex < 0 {
		panic(fmt.Errorf("alliance：仙盟仓库取出物品错误，allianceId:%d,index:%d,num:%d", allianceId, depotIndex, num))
	}

	_, al := s.getAlliance(allianceId)
	if al == nil {
		err = errorAllianceNoExist
		return
	}

	it := al.depotBag.getByIndex(depotIndex)
	if it == nil || it.IsEmpty() {
		err = errorAllianceDepotItemNotExist
		return
	}
	if it.num < num {
		err = errorAllianceDepotItemNotEnough
		return
	}
	if it.itemId != itemId {
		err = errorAllianceDepotItemNotExist
		return
	}

	//移除
	al.depotBag.removeIndex(depotIndex, num)
	changedItemList := al.depotBag.getChangedSlotAndReset()
	gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotChanged, al, changedItemList)
	return
}

func (s *allianceService) AutoRemoveDepot(memId int64, isAuto int32, zhuansheng int32, quality itemtypes.ItemQualityType) (al *Alliance, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	memObj := s.getAllianceMember(memId)
	if memObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	al = memObj.GetAlliance()
	if al == nil {
		err = errorAllianceNoExist
		return
	}

	//移除
	if isAuto == 1 {
		depotItemList := make([]*AllianceDepotItemObject, len(al.depotBag.GetItemList()))
		copy(depotItemList, al.depotBag.GetItemList())
		for _, itemObj := range depotItemList {
			if itemObj.IsEmpty() {
				continue
			}
			itemId := itemObj.itemId
			itemNum := itemObj.num
			itemTemp := item.GetItemService().GetItem(int(itemId))
			if itemTemp.GetQualityType() > quality {
				continue
			}
			if itemTemp.NeedZhuanShu > zhuansheng {
				continue
			}
			al.depotBag.removeIndex(itemObj.index, itemObj.num)

			// 移除日志
			removeItemReason := commonlog.AllianceLogReasonDepotItemRemove
			depotLogEventData := allianceeventtypes.CreateAllianceDepotItemChangedLogEventData(itemId, itemNum, removeItemReason, removeItemReason.String())
			gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotItemChangedLog, al, depotLogEventData)
		}
		changedItemList := al.depotBag.getChangedSlotAndReset()
		gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotChanged, al, changedItemList)
	}

	now := global.GetGame().GetTimeService().Now()
	al.allianceObject.isAutoRemoveDepot = isAuto
	al.allianceObject.maxRemoveZhuanSheng = zhuansheng
	al.allianceObject.maxRemoveQuality = quality
	al.allianceObject.updateTime = now
	al.allianceObject.SetModified()

	// 设置日志
	settingReason := commonlog.AllianceLogReasonDepotAutoRemoveSetting
	settingReasonTetx := fmt.Sprintf(settingReason.String(), isAuto, zhuansheng, quality)
	eventData := allianceeventtypes.CreateAllianceDepotSettingChangedLogEventData(settingReason, settingReasonTetx)
	gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotSettingChangedLog, al, eventData)

	return
}

func (s *allianceService) MergeDepot(pl player.Player, indexList []int32) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	allianceId := pl.GetAllianceId()
	_, al := s.getAlliance(allianceId)
	if al == nil {
		err = errorAllianceNoExist
		return
	}

	_, member := al.getMember(pl.GetId())
	if !member.IsMengZhu() {
		err = errorAllianceUserNotMengZhu
		return
	}

	for _, index := range indexList {
		it := al.depotBag.getByIndex(index)
		itemId := it.itemId
		itemNum := it.num
		if it == nil || it.IsEmpty() {
			err = errorAllianceDepotItemNotExist
			return
		}

		al.depotBag.removeIndex(index, itemNum)

		// 仓库整理日志
		removeItemReason := commonlog.AllianceLogReasonDepotMergeItem
		removeItemReasonText := fmt.Sprintf(removeItemReason.String(), pl.GetId())
		depotLogEventData := allianceeventtypes.CreateAllianceDepotItemChangedLogEventData(itemId, itemNum, removeItemReason, removeItemReasonText)
		gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotItemChangedLog, al, depotLogEventData)
	}

	al.depotBag.merge()
	itemList := al.depotBag.GetItemList()
	gameevent.Emit(allianceeventtypes.EventTypeAllianceDepotMerge, al, itemList)
	return
}

func (s *allianceService) GetAllianceMemberHasedJionTime(memberId int64) (hasedJoinTime int64) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	memObj := s.getAllianceMember(memberId)
	if memObj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	return now - memObj.GetJoinTime()
}

func (s *allianceService) initAllianceBoss(allianceId int64) *AllianceBossObject {
	now := global.GetGame().GetTimeService().Now()
	obj := createAllianceBossObject()
	id, _ := idutil.GetId()

	obj.id = id
	obj.serverId = global.GetGame().GetServerIndex()
	obj.allianceId = allianceId
	obj.summonTime = 0
	obj.bossExp = 0
	obj.bossLevel = 1
	obj.isSummon = 0
	obj.createTime = now
	obj.SetModified()
	return obj
}

//召唤仙盟boss
func (s *allianceService) AllianceSummonBoss(pl player.Player) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		err = errorAllianceUserNotInAlliance
		return
	}
	_, allianceObj := s.getAlliance(allianceId)
	if allianceObj == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	if pl.GetId() != allianceObj.GetAllianceMengZhuId() && pl.GetId() != allianceObj.GetFuMengZhuId() {
		err = errorAllianceBossSummonNoMengZhu
		return
	}

	allianceBossObject := allianceObj.GetAllianceBossObject()
	allianceBossObject.IsCrossFive()
	if allianceBossObject.GetIsSummon() {
		err = ErrorAllianceBossSummonedBoss
		return
	}

	level := allianceBossObject.GetBossLevel()
	allianceBossTemplate := alliancetemplate.GetAllianceTemplateService().GetAllianceBossTemplate(level)
	if allianceBossTemplate == nil {
		err = ErrorAllianceBossSummonedBoss
		return
	}

	allianceBossSceneData := alliancebossscene.CreateAllianceBossSceneData(allianceId, level)
	sc := scene.CreateFuBenScene(allianceBossTemplate.MapId, allianceBossSceneData)
	if sc == nil {
		return
	}
	allianceBossObject.SetIsSummon(true)
	s.allianceBossSceneDataMap[pl.GetAllianceId()] = allianceBossSceneData
	eventData := allianceeventtypes.CreateAllianceBossSummonSucessEventData(pl, sc)
	gameevent.Emit(allianceeventtypes.EvnetTypeAllianceBossSummonSucess, allianceObj, eventData)
	return
}

func (s *allianceService) AllianceBossEnd(allianceId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	delete(s.allianceBossSceneDataMap, allianceId)
}

func (s *allianceService) AllianceBossEnter(pl player.Player) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	allianceId := pl.GetAllianceId()
	if allianceId == 0 {
		err = errorAllianceUserNotInAlliance
		return
	}

	_, al := s.getAlliance(allianceId)
	if al == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	allianceBossObject := al.GetAllianceBossObject()
	allianceBossObject.IsCrossFive()
	if !allianceBossObject.GetIsSummon() {
		err = errorAllianceBossEnterNoStart
		return
	}

	sceneData, ok := s.allianceBossSceneDataMap[allianceId]
	if !ok {
		err = errorAllianceBossTodayFinish
		return
	}
	gameevent.Emit(allianceeventtypes.EventTypeAllowPlayerEnterAllianceBoss, sceneData, pl)
	return
}

func (s *allianceService) AllianceBossScene(allianceId int64) scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	_, al := s.getAlliance(allianceId)
	if al == nil {
		return nil
	}

	sc, ok := s.allianceBossSceneDataMap[allianceId]
	if !ok {
		return nil
	}
	if sc == nil {
		return nil
	}
	return sc.GetScene()
}

func (s *allianceService) AllianceBossInfo(allianceId int64) (status alliancetypes.AllianceBossStatus,
	level int32, exp int32, times int64, err error) {

	s.rwm.Lock()
	defer s.rwm.Unlock()

	if allianceId == 0 {
		err = errorAllianceUserNotInAlliance
		return
	}

	_, al := s.getAlliance(allianceId)
	if al == nil {
		err = errorAllianceUserNotInAlliance
		return
	}

	allianceBossObject := al.GetAllianceBossObject()
	allianceBossObject.IsCrossFive()

	level = allianceBossObject.GetBossLevel()
	exp = allianceBossObject.GetBossExp()

	status = alliancetypes.AllianceBossStatusDead

	times = -1
	if !allianceBossObject.GetIsSummon() {
		status = alliancetypes.AllianceBossStatusInit
		times = 0
	} else {
		sceneData, ok := s.allianceBossSceneDataMap[allianceId]
		if ok && sceneData.GetBossNpc().GetHP() != 0 {
			status = alliancetypes.AllianceBossStatusSummon
			times = sceneData.GetSummonTime()
		}
	}
	return
}

//仙盟boss增加经验
func (s *allianceService) AllianceBossAddExp(allianceId int64, exp int32) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	_, al := s.getAlliance(allianceId)
	if al == nil {
		return
	}

	allianceBossObject := al.GetAllianceBossObject()
	allianceBossObject.IsCrossFive()

	allianceBossObject.AddExp(exp)
}

//GM 仙盟boss
func (s *allianceService) GMAllianceBossReset(allianceId int64) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()

	_, al := s.getAlliance(allianceId)
	if al == nil {
		return
	}

	al.allianceBossObject.IsCrossFive()
	al.allianceBossObject.bossLevel = 1
	al.allianceBossObject.bossExp = 0
	al.allianceBossObject.isSummon = 0
	al.allianceBossObject.updateTime = now
	al.allianceBossObject.SetModified()
	flag = true
	return
}

//仙盟改变阵营
func (s *allianceService) AllianceChangeCamp(allianceId int64, campType chuangshitypes.ChuangShiCampType) {
	// s.rwm.Lock()
	// defer s.rwm.Unlock()

	// _, al := s.getAlliance(allianceId)
	// if al == nil {
	// 	return
	// }

	// if al.allianceObject.campType == campType {
	// 	return
	// }

	// now := global.GetGame().GetTimeService().Now()
	// al.allianceObject.campType = campType
	// al.allianceObject.updateTime = now
	// al.allianceObject.SetModified()

	// gameevent.Emit(allianceeventtypes.EventTypeAllianceCampTypeChanged, al, nil)
	return
}

func (s *allianceService) GetCurModelList() (robotList []scene.RobotPlayer) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	for _, robot := range s.hegemonWinnerModelMap {
		robotList = append(robotList, robot)
	}
	return robotList
}

func (s *allianceService) AddModel(model scene.RobotPlayer) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.hegemonWinnerModelMap[model.GetId()] = model
}

func (s *allianceService) RemoveModel(model scene.RobotPlayer) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	delete(s.hegemonWinnerModelMap, model.GetId())
}

func (s *allianceService) AllianceMerge(inviteAllianceId, allianceId int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	//没有合并邀请
	mergeApply, ok := s.allianceMergeApply[inviteAllianceId]
	if !ok {
		err = errorAllianceMergeApplyExpired
		return
	}
	//合并仙盟已经换了
	if mergeApply.inviteAllianceId != allianceId {
		err = errorAllianceMergeApplyExpired
		return
	}
	hemengCd := int64(alliancetemplate.GetAllianceTemplateService().GetAllianceConstantTemp().HemengCd) * int64(common.SECOND)
	if !mergeApply.IsCD(now, hemengCd) {
		err = errorAllianceMergeApplyExpired
		return
	}

	//清除合并申请
	s.clearAllianceMergeInvite(inviteAllianceId)

	_, parentAl := s.getAlliance(inviteAllianceId)
	if parentAl == nil {
		return
	}
	_, childAl := s.getAlliance(allianceId)
	if childAl == nil {
		return
	}
	if parentAl.IfFull(childAl.NumOfMembers()) {
		err = errorAllianceAlreadyFullApply
		return
	}

	//解散仙盟
	tempMemberList := make([]*AllianceMemberObject, len(childAl.GetMemberList()))
	copy(tempMemberList, childAl.GetMemberList())
	for _, mem := range tempMemberList {

		mem.deleteTime = now
		mem.SetModified()
		s.removeAllianceMember(mem, false)

		s.joinAlliance(parentAl, mem.level, mem.name, mem.force, mem.memberId, mem.role, mem.sex, mem.onlineStatus, mem.lastLogoutTime, mem.lingyuId, mem.zhuanSheng, mem.vip)
	}
	s.removeAlliance(childAl)

	childName := coreutils.FormatColor(alliancetypes.ColorTypeLogName, coreutils.FormatNoticeStr(childAl.GetAllianceName()))
	log := fmt.Sprintf(lang.GetLangService().ReadLang(lang.AllianceInviteAllianceMergeLog), childName)
	parentAl.addLog(log)

	parentAl.allianceObject.lastMergeTime = now
	parentAl.allianceObject.updateTime = now
	parentAl.allianceObject.SetModified()

	//合并事件
	gameevent.Emit(allianceeventtypes.EventTypeAllianceMerge, parentAl, tempMemberList)
	gameevent.Emit(allianceeventtypes.EventTypeAllianceMergeLog, parentAl, allianceId)

	return
}

//修改仙盟名称
func (s *allianceService) UpdateAllianceName(pl player.Player, newName string) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	_, al := s.getAlliance(pl.GetAllianceId())
	if al == nil {
		err = errorAllianceNoExist
		return
	}

	_, mem := al.getMember(pl.GetId())
	if mem == nil {
		err = errorAllianceUserNotInAlliance
	}

	if !mem.IsMengZhu() && !mem.IsFuMengZhu() {
		err = errorAlliancePrivilegeNotEnough
		return
	}

	//验证名字合法
	tal := s.getAllianceByName(newName)

	if tal != nil {
		err = errorAllianceNameExist
		return
	}

	// 合盟后改名
	if !al.ifCanAllianceRename() {
		return
	}

	oldName := al.allianceObject.name
	now := global.GetGame().GetTimeService().Now()
	al.allianceObject.name = newName
	al.allianceObject.updateTime = now
	al.allianceObject.SetModified()

	// content := lang.GetLangService().ReadLang(lang.AllianceImpeachLog)
	// log := fmt.Sprintf(content, coreutils.FormatColor(alliancetypes.ColorTypeLogName, impeachMemObj.GetName()))
	// al.addLog(log)

	//仙盟改名
	gameevent.Emit(allianceeventtypes.EventTypeAllianceNameChanged, al, oldName)
	return
}

//是否可以合并
func (s *allianceService) IfAllianceCanMergeApply(allianceId int64) bool {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	return s.ifAllianceCanMergeApply(allianceId)
}

func (s *allianceService) ifAllianceCanMergeApply(allianceId int64) bool {
	mergeApply, ok := s.allianceMergeApply[allianceId]
	if !ok {
		return true
	}
	now := global.GetGame().GetTimeService().Now()
	hemengCd := int64(alliancetemplate.GetAllianceTemplateService().GetAllianceConstantTemp().HemengCd) * int64(common.SECOND)

	if mergeApply.IsCD(now, hemengCd) {

		return false
	}
	return true
}

//合并邀请
func (s *allianceService) AllianceMergeInvite(allianceId int64, inviteAllianceId int64) (err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	now := global.GetGame().GetTimeService().Now()
	hemengCd := int64(alliancetemplate.GetAllianceTemplateService().GetAllianceConstantTemp().HemengCd) * int64(common.SECOND)
	//被别人邀请中
	for _, mergeApply := range s.allianceMergeApply {
		if mergeApply.inviteAllianceId != inviteAllianceId {
			continue
		}
		if mergeApply.IsCD(now, hemengCd) {
			err = errorAllianceMergeApplying
			return err
		}
	}

	mergeApply, ok := s.allianceMergeApply[allianceId]
	if !ok {
		mergeApply = &allianceMergeApply{}
		mergeApply.inviteTime = now
		mergeApply.inviteAllianceId = inviteAllianceId
		s.allianceMergeApply[allianceId] = mergeApply
		return nil
	}

	if mergeApply.IsCD(now, hemengCd) {
		err = errorAllianceMergeApplyCd
		return err
	}

	//修改当前邀请的
	mergeApply.inviteTime = now
	mergeApply.inviteAllianceId = inviteAllianceId
	return nil

}

//清除合并
func (s *allianceService) ClearAllianceMergeInvite(allianceId int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	s.clearAllianceMergeInvite(allianceId)
}

//清除合并
func (s *allianceService) clearAllianceMergeInvite(allianceId int64) {
	delete(s.allianceMergeApply, allianceId)
}

func checkAllianceVersion() alliancetypes.AllianceVersionType {
	// 判断创建仙盟版本
	versionType := alliancetypes.AllianceVersionTypeOld
	sdkList := center.GetCenterService().GetSdkList()
	for _, sdk := range sdkList {
		temp := chargetemplate.GetChargeTemplateService().GetQuDaoTemplateByType(sdk)
		if temp == nil {
			continue
		}

		if temp.GetAllianceVersion() == alliancetypes.AllianceVersionTypeNew {
			versionType = alliancetypes.AllianceVersionTypeNew
			break
		}
	}
	return versionType
}

//解散仙盟
func (s *allianceService) GMDismissAlliance(allianceId int64) (al *Alliance, tempMemberList []*AllianceMemberObject, err error) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	if s.allianceSceneData != nil {
		err = errorAllianceDismissCanNotInSceneOpened
		return
	}
	_, al = s.getAlliance(allianceId)
	if al == nil {
		return
	}

	_, ok := s.allianceBossSceneDataMap[allianceId]
	if ok {
		err = errorAllianceDismissCanNotInBoss
		return
	}

	mengZhuId := al.GetAllianceMengZhuId()
	//删除所有成员
	now := global.GetGame().GetTimeService().Now()
	tempMemberList = make([]*AllianceMemberObject, len(al.GetMemberList()))
	copy(tempMemberList, al.GetMemberList())
	for _, mem := range tempMemberList {
		now := global.GetGame().GetTimeService().Now()
		mem.deleteTime = now
		mem.SetModified()

		s.removeAllianceMember(mem, true)
	}

	alObj := al.GetAllianceObject()
	alObj.deleteTime = now
	alObj.SetModified()
	s.removeAlliance(al)

	gameevent.Emit(allianceeventtypes.EventTypeAllianceDismiss, nil, mengZhuId)
	return
}

var (
	once sync.Once
	as   *allianceService
)

func Init() (err error) {
	once.Do(func() {
		as = &allianceService{}
		err = as.init()
	})
	return err
}

func GetAllianceService() AllianceService {
	return as
}
