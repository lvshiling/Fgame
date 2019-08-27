package player

import (
	commonlog "fgame/fgame/common/log"
	commonlogtypes "fgame/fgame/common/log/types"
	"fgame/fgame/game/alliance/dao"
	allianceeventtypes "fgame/fgame/game/alliance/event/types"
	alliancetemplate "fgame/fgame/game/alliance/template"
	alliancetypes "fgame/fgame/game/alliance/types"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	yuxitemplate "fgame/fgame/game/yuxi/template"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/timeutils"
	"fmt"
)

//玩家仙盟管理器
type PlayerAllianceDataManager struct {
	p player.Player
	//玩家仙盟对象
	playerAllianceObject *PlayerAllianceObject
	//玩家仙术对象map
	playerAllianceSkillObjectMap map[alliancetypes.AllianceSkillType]*PlayerAllianceSkillObject
	//上一次被击杀时间
	lastKilledTimeMap map[int64]int64
	//上一次被击杀通知时间
	lastBeKillNoticeTime int64
	//盟主id
	curMengZhuId int64
	//玩家仙盟职位
	curPos alliancetypes.AlliancePosition
	//上次一键加入时间
	lastBatchJoinTime int64
}

func (m *PlayerAllianceDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerAllianceDataManager) Load() (err error) {
	//加载玩家仙盟信息
	playerAliianceEntity, err := dao.GetAllianceDao().GetPlayerAlliance(m.p.GetId())
	if err != nil {
		return
	}
	if playerAliianceEntity == nil {
		m.initPlayerAllianceObject()
	} else {
		m.playerAllianceObject = newPlayerAllianceObject(m.p)
		m.playerAllianceObject.FromEntity(playerAliianceEntity)
	}

	//加载玩家仙术信息
	allianceSkillEntityArr, err := dao.GetAllianceDao().GetPlayerAllianceSkillList(m.p.GetId())
	for _, skillEntity := range allianceSkillEntityArr {
		skillObj := newPlayerAllianceSkillObject(m.p)
		skillObj.FromEntity(skillEntity)

		m.playerAllianceSkillObjectMap[skillObj.skillType] = skillObj
	}

	m.lastKilledTimeMap = make(map[int64]int64)

	return nil
}

//第一次初始化
func (m *PlayerAllianceDataManager) initPlayerAllianceObject() {
	o := newPlayerAllianceObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.donateMap = make(map[alliancetypes.AllianceJuanXianType]int32)
	o.sceneRewardMap = make(map[int32]int32)
	o.createTime = now
	o.totalWinTime = 0
	o.SetModified()
	m.playerAllianceObject = o
}

//初始化仙术
func (m *PlayerAllianceDataManager) initPlayerAllianceSkillObject(typ alliancetypes.AllianceSkillType) *PlayerAllianceSkillObject {
	if !typ.Valid() {
		panic("仙盟技能初始化，技能类型错误")
	}

	o := newPlayerAllianceSkillObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.level = 0
	o.skillType = typ
	o.createTime = now
	o.SetModified()

	m.playerAllianceSkillObjectMap[typ] = o

	return o
}

//加载后
func (m *PlayerAllianceDataManager) AfterLoad() (err error) {

	//刷新捐献次数
	err = m.refreshTimes()
	if err != nil {
		return err
	}

	return nil
}

//刷新次数
func (m *PlayerAllianceDataManager) refreshTimes() error {
	now := global.GetGame().GetTimeService().Now()
	lastJuanXuanTime := m.playerAllianceObject.lastJuanXuanTime

	flag, err := timeutils.IsSameDay(lastJuanXuanTime, now)
	if err != nil {
		return err
	}
	if !flag {
		//TODO 优化
		m.playerAllianceObject.donateMap = make(map[alliancetypes.AllianceJuanXianType]int32)
		m.playerAllianceObject.lastJuanXuanTime = now
		m.playerAllianceObject.updateTime = now
		m.playerAllianceObject.SetModified()
	}
	return nil
}

//心跳
func (m *PlayerAllianceDataManager) Heartbeat() {
}

//获取个人数据
func (m *PlayerAllianceDataManager) GetPlayerAllianceObject() *PlayerAllianceObject {
	m.refreshTimes()
	return m.playerAllianceObject
}

//获取仙盟id
func (m *PlayerAllianceDataManager) GetAllianceId() int64 {
	return m.playerAllianceObject.allianceId
}

//获取仙盟名字
func (m *PlayerAllianceDataManager) GetAllianceName() string {
	return m.playerAllianceObject.allianceName
}

//获取盟主id
func (m *PlayerAllianceDataManager) GetMengZhuId() int64 {
	return m.curMengZhuId
}

//获取玩家仙盟职位
func (m *PlayerAllianceDataManager) GetPlayerAlliancePos() alliancetypes.AlliancePosition {
	return m.curPos
}

//获取仙术数据
func (m *PlayerAllianceDataManager) getAllianceSkillObject(typ alliancetypes.AllianceSkillType) *PlayerAllianceSkillObject {
	skillObj, ok := m.playerAllianceSkillObjectMap[typ]
	if ok {
		return skillObj
	}

	return m.initPlayerAllianceSkillObject(typ)
}

//获取仙术等级
func (m *PlayerAllianceDataManager) GetAllianceSkillLevel(typ alliancetypes.AllianceSkillType) int32 {
	return m.getAllianceSkillObject(typ).level
}

//获取仙术数据列表
func (m *PlayerAllianceDataManager) GetPlayerAllianceSkillMap() map[alliancetypes.AllianceSkillType]*PlayerAllianceSkillObject {
	return m.playerAllianceSkillObjectMap
}

//获取激活仙术列表
func (m *PlayerAllianceDataManager) GetEffectiveAllianceSkillList() map[alliancetypes.AllianceSkillType]*PlayerAllianceSkillObject {
	effectiveSkill := make(map[alliancetypes.AllianceSkillType]*PlayerAllianceSkillObject)
	for skillType, skillObj := range m.playerAllianceSkillObjectMap {
		if m.IsOpenAllianceSkill(skillObj.level, skillType) {
			effectiveSkill[skillType] = skillObj
		}
	}

	return effectiveSkill
}

//是否仙术满级
func (m *PlayerAllianceDataManager) IsAllianceSkillFullLevel(typ alliancetypes.AllianceSkillType) bool {
	level := m.GetAllianceSkillLevel(typ)
	if level == 0 {
		return false
	}
	tem := alliancetemplate.GetAllianceTemplateService().GetAllianceSkillTemplateByType(level, typ)
	if tem.NextId != 0 {
		//仙术等级达到仙盟等级限制
		nexTem := tem.GetNextTemplate()
		allianceLevel := m.playerAllianceObject.allianceLevel
		isMax := nexTem.NeedUnionLevel > allianceLevel
		return isMax
	}

	return true
}

//仙术等级是否开启
func (m *PlayerAllianceDataManager) IsOpenAllianceSkill(level int32, typ alliancetypes.AllianceSkillType) bool {
	if level == 0 {
		level = 1
	}

	skillTemp := alliancetemplate.GetAllianceTemplateService().GetAllianceSkillTemplateByType(level, typ)
	if skillTemp == nil {
		return false
	}
	needLevel := skillTemp.OpenNeedLevel
	allianceLevel := m.playerAllianceObject.allianceLevel
	isOpen := needLevel <= allianceLevel

	return isOpen
}

//仙术升级
func (m *PlayerAllianceDataManager) UpgradeAllianceSkill(typ alliancetypes.AllianceSkillType) (flag bool) {
	flag = m.IsAllianceSkillFullLevel(typ)
	if flag {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj := m.getAllianceSkillObject(typ)
	obj.level += 1
	obj.updateTime = now
	obj.SetModified()
	flag = true

	eventData := allianceeventtypes.CreatePlayerAllianceSkillUpgradeEventData(typ, obj.level)
	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceSkillUpgrade, m.p, eventData)
	return
}

// GM设置技能等级
func (m *PlayerAllianceDataManager) GMSetAllianceSkill(typ alliancetypes.AllianceSkillType, level int32) {
	now := global.GetGame().GetTimeService().Now()
	obj := m.getAllianceSkillObject(typ)
	obj.level = level
	obj.updateTime = now
	obj.SetModified()

	eventData := allianceeventtypes.CreatePlayerAllianceSkillUpgradeEventData(typ, level)
	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceSkillUpgrade, m.p, eventData)
}

//贡献是否足够
func (m *PlayerAllianceDataManager) IsEnoughGongXian(needGongXian int64) bool {
	curGongXian := m.playerAllianceObject.currentGongXian
	isEnough := curGongXian >= needGongXian
	return isEnough
}

//消耗贡献值
func (m *PlayerAllianceDataManager) UseGongXian(num int64) bool {
	if num <= 0 {
		panic("alliance: 消耗不能小于 0")
	}
	if m.playerAllianceObject.currentGongXian < num {
		return false
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.currentGongXian -= num
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()

	return true
}

//添加贡献值
func (m *PlayerAllianceDataManager) AddGongXian(num int64) {
	if num <= 0 {
		panic("alliance: 添加不能小于 0")
	}
	m.playerAllianceObject.currentGongXian += num
	m.playerAllianceObject.SetModified()
}

//同步仙盟
func (m *PlayerAllianceDataManager) SyncAlliance(allianceId, mengzhuId int64, name string, level int32, pos alliancetypes.AlliancePosition) (err error) {
	//同步allianceLevel前
	oldEffectiveSkillMap := m.GetEffectiveAllianceSkillList()

	m.curMengZhuId = mengzhuId
	m.curPos = pos
	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.allianceId = allianceId
	m.playerAllianceObject.allianceName = name
	m.playerAllianceObject.allianceLevel = level
	m.playerAllianceObject.SetModified()
	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceChanged, m.p, nil)
	//加载可激活仙术：同步allianceLevel后
	newEffectiveSkillMap := m.GetEffectiveAllianceSkillList()
	if len(oldEffectiveSkillMap) != len(newEffectiveSkillMap) {
		gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceSkillChanged, m.p, oldEffectiveSkillMap)
	}

	return
}

//同步仙盟盟主
func (m *PlayerAllianceDataManager) SyncAllianceMengzhu(mengzhuId int64) (err error) {
	m.curMengZhuId = mengzhuId

	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceMengZhuChanged, m.p, nil)
	return
}

//同步仙盟等级
func (m *PlayerAllianceDataManager) SyncAllianceLevel(allianceLevel int32) {
	//同步allianceLevel前
	oldEffectiveSkillMap := m.GetEffectiveAllianceSkillList()

	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.allianceLevel = allianceLevel
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()

	//加载可激活仙术：同步allianceLevel后
	newEffectiveSkillMap := m.GetEffectiveAllianceSkillList()
	if len(oldEffectiveSkillMap) != len(newEffectiveSkillMap) {
		gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceSkillChanged, m.p, oldEffectiveSkillMap)
	}
}

//同步仙盟职位
func (m *PlayerAllianceDataManager) SyncAlliancePos(pos alliancetypes.AlliancePosition) {
	m.curPos = pos
	gameevent.Emit(allianceeventtypes.EventTypePlayerAlliancePositionChanged, m.p, nil)
}

//同步仙盟职位
func (m *PlayerAllianceDataManager) SyncAllianceName(allianceName string) {
	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.allianceName = allianceName
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()

	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceChanged, m.p, nil)

}

//进入联盟
func (m *PlayerAllianceDataManager) JoinAlliance(allianceId, mengzhuId int64, name string, level int32, pos alliancetypes.AlliancePosition) {
	if allianceId == 0 {
		panic(fmt.Errorf("alliance:仙盟id不能为0"))
	}
	if m.playerAllianceObject.allianceId == allianceId {
		return
	}

	m.curMengZhuId = mengzhuId
	m.curPos = pos
	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.allianceId = allianceId
	m.playerAllianceObject.allianceName = name
	m.playerAllianceObject.allianceLevel = level
	m.playerAllianceObject.SetModified()
	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceJoin, m.p, nil)
}

//退出仙盟
func (m *PlayerAllianceDataManager) ExitAlliance(isClear bool) {
	if m.playerAllianceObject.allianceId == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.curMengZhuId = 0
	m.curPos = 0
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.allianceId = 0
	m.playerAllianceObject.allianceName = ""
	if isClear {
		m.playerAllianceObject.depotPoint = 0
	}

	m.playerAllianceObject.SetModified()
	gameevent.Emit(allianceeventtypes.EventTypePlayerAllianceExit, m.p, nil)
}

func (m *PlayerAllianceDataManager) IfCanDonate(typ alliancetypes.AllianceJuanXianType) bool {
	donateTemplate := alliancetemplate.GetAllianceTemplateService().GetUnionDonateTemplate(typ)
	if donateTemplate == nil {
		return false
	}
	if donateTemplate.DonateLimit == 0 {
		return true
	}
	useTimes := m.playerAllianceObject.donateMap[typ]
	return useTimes < donateTemplate.DonateLimit
}

func (m *PlayerAllianceDataManager) Donate(typ alliancetypes.AllianceJuanXianType) bool {
	flag := m.IfCanDonate(typ)
	if !flag {
		return false
	}
	donateTemplate := alliancetemplate.GetAllianceTemplateService().GetUnionDonateTemplate(typ)
	gongxianNum := int64(donateTemplate.DonateContribution)

	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.donateMap[typ] += 1
	m.playerAllianceObject.lastJuanXuanTime = now
	m.playerAllianceObject.updateTime = now
	m.AddGongXian(gongxianNum)
	return true
}

func (m *PlayerAllianceDataManager) RefreshAllianceScene(endTime int64) {
	if m.playerAllianceObject.lastAllianceSceneEndTime != endTime {
		m.playerAllianceObject.lastAllianceSceneEndTime = endTime
		m.playerAllianceObject.sceneRewardMap = make(map[int32]int32)
		m.playerAllianceObject.warPoint = 0
		m.playerAllianceObject.SetModified()
	}
}

func (m *PlayerAllianceDataManager) GetRewardList() (rewardList []int32) {
	for index, _ := range m.playerAllianceObject.sceneRewardMap {
		rewardList = append(rewardList, index)
	}
	return rewardList
}

func (m *PlayerAllianceDataManager) GetReward(door int32) bool {
	if !m.IfCanGetReward(door) {
		return false
	}
	m.playerAllianceObject.sceneRewardMap[door] = 1
	m.playerAllianceObject.SetModified()
	return true
}

func (m *PlayerAllianceDataManager) IfCanGetReward(door int32) bool {
	_, exist := m.playerAllianceObject.sceneRewardMap[door]
	if exist {
		return false
	}
	return true
}

//增加城战积分
func (m *PlayerAllianceDataManager) AddWarPoint(warPoint int32) int32 {
	if warPoint <= 0 {
		panic(fmt.Errorf("property:城战积分应该大于0"))
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.warPoint += warPoint
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()

	return m.playerAllianceObject.warPoint
}

//是否足够城战积分
func (m *PlayerAllianceDataManager) IsEnoughWarPoint(warPoint int32) bool {
	return m.playerAllianceObject.warPoint >= warPoint
}

func (m *PlayerAllianceDataManager) GetWarPoint() int32 {
	return m.playerAllianceObject.warPoint
}

//增加腰牌
func (m *PlayerAllianceDataManager) AddYaoPai(yaoPai int32, yaoPaiLogReason commonlogtypes.LogReason, yaoPaiLogReasonText string) {
	if yaoPai <= 0 {
		panic(fmt.Errorf("property:腰牌应该大于0"))
	}
	now := global.GetGame().GetTimeService().Now()
	m.refreshYaoPai()
	currentYaoPai := m.playerAllianceObject.yaoPai
	currentYaoPai += yaoPai
	m.playerAllianceObject.yaoPai = currentYaoPai
	m.playerAllianceObject.lastYaoPaiUpdateTime = now
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()

	gameevent.Emit(allianceeventtypes.EventTypePlayerYaoPaiChanged, m.p, nil)
	return
}

//花费腰牌
func (m *PlayerAllianceDataManager) CostYaoPai(yaoPai int32, yaoPaiLogReason commonlog.YaoPaiLogReason, yaoPaiLogReasonText string) bool {
	if yaoPai <= 0 {
		panic(fmt.Errorf("property:腰牌应该大于0"))
	}

	now := global.GetGame().GetTimeService().Now()
	m.refreshYaoPai()
	currentYaoPai := m.playerAllianceObject.yaoPai
	currentYaoPai -= yaoPai
	m.playerAllianceObject.yaoPai = currentYaoPai
	m.playerAllianceObject.lastYaoPaiUpdateTime = now
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()
	gameevent.Emit(allianceeventtypes.EventTypePlayerYaoPaiChanged, m.p, nil)
	return true
}

//是否要足够的腰牌
func (m *PlayerAllianceDataManager) HasEnoughYaoPai() bool {
	m.refreshYaoPai()
	convertOneNeedNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeConvertNeedNum)
	currentYaoPai := m.playerAllianceObject.yaoPai
	return currentYaoPai >= convertOneNeedNum
}

func (m *PlayerAllianceDataManager) refreshYaoPai() error {
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(m.playerAllianceObject.lastYaoPaiUpdateTime, now)
	if err != nil {
		return err
	}
	if !isSame {
		m.playerAllianceObject.yaoPai = 0
		m.playerAllianceObject.lastYaoPaiUpdateTime = now
		m.playerAllianceObject.updateTime = now
		m.playerAllianceObject.SetModified()

	}
	return nil
}

//是否有兑换次数
func (m *PlayerAllianceDataManager) HasEnoughConvetTiems() bool {
	m.refreshConvertTimes()
	maxConvertTimes := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeConvertLimit)
	currentConvetTimes := m.playerAllianceObject.convertTimes
	return maxConvertTimes > currentConvetTimes
}

//更新兑换次数
func (m *PlayerAllianceDataManager) UpdateConvertTimes(addTimes int32) {
	currentConvertTimes := m.playerAllianceObject.convertTimes
	currentConvertTimes += addTimes
	now := global.GetGame().GetTimeService().Now()

	m.playerAllianceObject.convertTimes = currentConvertTimes
	m.playerAllianceObject.lastConvertUpdateTime = now
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()

	return

}

//重置原地复活次数
func (m *PlayerAllianceDataManager) RestReliveTime() {
	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.reliveTime = 0
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()
	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneImmediate, m.p, m.playerAllianceObject.reliveTime)
}

//增加原地复活次数
func (m *PlayerAllianceDataManager) AddReliveTime() {
	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.reliveTime += 1
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.lastReliveTime = now
	m.playerAllianceObject.SetModified()
	gameevent.Emit(allianceeventtypes.EventTypeAllianceSceneImmediate, m.p, m.playerAllianceObject.reliveTime)
}

//GM重置兑换次数
func (m *PlayerAllianceDataManager) GMResetConvertTimes() {
	yesterday, _ := timeutils.BeginOfYesterday()
	m.playerAllianceObject.lastConvertUpdateTime = yesterday
	m.refreshConvertTimes()
}

//GM重置捐献次数
func (m *PlayerAllianceDataManager) GMResetDonateTimes() {
	yesterday, _ := timeutils.BeginOfYesterday()
	m.playerAllianceObject.lastJuanXuanTime = yesterday
	m.refreshTimes()
}

func (m *PlayerAllianceDataManager) refreshConvertTimes() error {
	now := global.GetGame().GetTimeService().Now()
	isSame, err := timeutils.IsSameDay(m.playerAllianceObject.lastConvertUpdateTime, now)
	if err != nil {
		return err
	}
	if !isSame {
		m.playerAllianceObject.convertTimes = 0
		m.playerAllianceObject.lastConvertUpdateTime = now
		m.playerAllianceObject.updateTime = now
		m.playerAllianceObject.SetModified()

	}
	return nil
}

//获取兑换次数
func (m *PlayerAllianceDataManager) GetCanConvertTimes() int32 {
	m.refreshConvertTimes()
	m.refreshYaoPai()

	maxConvertTimes := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeConvertLimit)
	curConvertTimes := m.playerAllianceObject.convertTimes
	curYaoPai := m.playerAllianceObject.yaoPai
	convertOneNeedNum := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeConvertNeedNum)

	canConvertTimes := maxConvertTimes - curConvertTimes
	enoughConvertTimes := curYaoPai / convertOneNeedNum
	if canConvertTimes > enoughConvertTimes {
		return enoughConvertTimes
	} else {
		return canConvertTimes
	}

}

func (m *PlayerAllianceDataManager) GetYaoPai() int32 {
	m.refreshYaoPai()
	currentYaoPai := m.playerAllianceObject.yaoPai
	return currentYaoPai
}

//跨天更新仙盟个人次数信息
func (m *PlayerAllianceDataManager) RefreshAlliancePlayerInfo() {
	m.refreshYaoPai()
	m.refreshConvertTimes()
	m.refreshTimes()
}

//获取上次击杀时间
func (m *PlayerAllianceDataManager) GetLastBeKilledTime(playerId int64) int64 {
	return m.lastKilledTimeMap[playerId]
}

//杀掉仙盟玩家
func (m *PlayerAllianceDataManager) KilledByAlliancePlayer(playerId int64) {
	now := global.GetGame().GetTimeService().Now()
	m.lastKilledTimeMap[playerId] = now
}

//清除列表
func (m *PlayerAllianceDataManager) ClearAllKillTime() {
	for playerId, _ := range m.lastKilledTimeMap {
		delete(m.lastKilledTimeMap, playerId)
	}
}

//获取仓库贡献
func (m *PlayerAllianceDataManager) AddDepotPoint(point int32) {
	if point < 0 {
		panic(fmt.Errorf("Alliance:仓库积分point应该大于0"))
	}

	if m.playerAllianceObject.allianceId == 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.depotPoint += point
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()
}

//消耗仓库贡献
func (m *PlayerAllianceDataManager) CostDepotPoint(costPoint int32) {
	if costPoint < 0 {
		panic(fmt.Errorf("Alliance:仓库积分point应该大于0"))
	}

	if m.playerAllianceObject.allianceId == 0 {
		return
	}

	curPoint := m.playerAllianceObject.depotPoint
	if curPoint < costPoint {
		panic(fmt.Errorf("Alliance:当前仓库积分:%d不足,需要:%d", curPoint, costPoint))
	}

	curPoint -= costPoint
	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.depotPoint = curPoint
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()
	return
}

//是否积分足够
func (m *PlayerAllianceDataManager) HasEnoughPoint(needPoint int32) bool {
	if needPoint < 0 {
		panic(fmt.Errorf("Alliance:仓库积分point应该大于0"))
	}

	curPoint := m.playerAllianceObject.depotPoint
	return curPoint >= needPoint
}

func (m *PlayerAllianceDataManager) GetDepotPoint() int32 {
	currentYaoPai := m.playerAllianceObject.depotPoint
	return currentYaoPai
}

//获取上次被击杀仙盟推送时间
func (m *PlayerAllianceDataManager) GetLastBeKilledNoticeTime() int64 {
	return m.lastBeKillNoticeTime
}

//更新上次被击杀仙盟推送时间
func (m *PlayerAllianceDataManager) UpdateBeKillNoticeTime() {
	now := global.GetGame().GetTimeService().Now()
	m.lastBeKillNoticeTime = now
}

//是否召集CD
func (m *PlayerAllianceDataManager) IsMemberCallCD(callType alliancetypes.AllianceCallType) bool {
	cd := int64(0)
	lastTime := int64(0)
	switch callType {
	case alliancetypes.AllianceCallTypeCommon:
		{
			cd = int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceMemberResuceCD))
			lastTime = m.playerAllianceObject.lastMemberCallTime
		}
	case alliancetypes.AllianceCallTypeYuXi:
		{
			cd = yuxitemplate.GetYuXiTemplateService().GetYuXiConstTemplate().RescueCd
			lastTime = m.playerAllianceObject.lastYuXiMemberCallTime
		}
	default:
		return true
	}

	now := global.GetGame().GetTimeService().Now()
	elapse := now - lastTime
	if elapse < cd {
		return true
	}

	return false
}

//更新召集时间
func (m *PlayerAllianceDataManager) UpdateLastMemberCall(callType alliancetypes.AllianceCallType) (lastTime int64, flag bool) {
	now := global.GetGame().GetTimeService().Now()
	switch callType {
	case alliancetypes.AllianceCallTypeCommon:
		{
			m.playerAllianceObject.lastMemberCallTime = now
		}
	case alliancetypes.AllianceCallTypeYuXi:
		{
			m.playerAllianceObject.lastYuXiMemberCallTime = now
		}
	default:
		return
	}

	m.playerAllianceObject.SetModified()

	lastTime = now
	flag = true
	return
}

//是否一键加入CD
func (m *PlayerAllianceDataManager) IsOnBatchJoinCD() bool {
	now := global.GetGame().GetTimeService().Now()
	cd := int64(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeAllianceJoinBatchCD))
	if now-m.lastBatchJoinTime > cd {
		return false
	}
	return true
}

//增加获胜次数
func (m *PlayerAllianceDataManager) AddWinTime() {
	now := global.GetGame().GetTimeService().Now()
	m.playerAllianceObject.totalWinTime += 1
	m.playerAllianceObject.updateTime = now
	m.playerAllianceObject.SetModified()
	gameevent.Emit(allianceeventtypes.EventTypeAllianceWinChengZhan, m.p, m.playerAllianceObject.totalWinTime)
}

//获取城战获胜次数
func (m *PlayerAllianceDataManager) GetWinTime() int32 {
	return m.playerAllianceObject.totalWinTime
}

func (m *PlayerAllianceDataManager) UpdateLastBatchJoinTime() {
	now := global.GetGame().GetTimeService().Now()
	m.lastBatchJoinTime = now
}

func createPlayerAllianceDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerAllianceDataManager{}
	m.p = p
	m.playerAllianceSkillObjectMap = make(map[alliancetypes.AllianceSkillType]*PlayerAllianceSkillObject)
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerAllianceDataManagerType, player.PlayerDataManagerFactoryFunc(createPlayerAllianceDataManager))
}
