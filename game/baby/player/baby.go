package player

import (
	commonlog "fgame/fgame/common/log"
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/game/baby/dao"
	babyeventtypes "fgame/fgame/game/baby/event/types"
	babytemplate "fgame/fgame/game/baby/template"
	babytypes "fgame/fgame/game/baby/types"
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	skilltemplate "fgame/fgame/game/skill/template"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//玩家宝宝管理器
type PlayerBabyDataManager struct {
	p                    player.Player
	playerPregnantObject *PlayerPregnantObject              //洞房信息
	playerBabyObjectList []*PlayerBabyObject                //宝宝对象
	toyBagMap            map[babytypes.ToySuitType]*BodyBag //宝宝玩具背包
	lastNoticeTime       int64                              //上次出生提醒
	hbRunner             heartbeat.HeartbeatTaskRunner
	//
	coupleBabyList  []*babytypes.CoupleBabyData //配偶宝宝信息
	babyPowerObject *PlayerBabyPowerObject
}

func (m *PlayerBabyDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerBabyDataManager) Load() (err error) {
	//加载玩家怀孕信息
	pregnantEntity, err := dao.GetBabyDao().GetPregnantEntity(m.p.GetId())
	if err != nil {
		return
	}
	if pregnantEntity == nil {
		m.initPlayerPregnantObject()
	} else {
		m.playerPregnantObject = NewPlayerPregnantObject(m.p)
		m.playerPregnantObject.FromEntity(pregnantEntity)
	}

	// 玩家宝宝信息
	babyList, err := dao.GetBabyDao().GetBabyEntityList(m.p.GetId())
	if err != nil {
		return
	}
	for _, entity := range babyList {
		obj := NewPlayerBabyObject(m.p)
		obj.FromEntity(entity)
		m.playerBabyObjectList = append(m.playerBabyObjectList, obj)
	}

	// 宝宝玩具数据
	toySlotList, err := dao.GetBabyDao().GetBabyToySlotEntityList(m.p.GetId())
	if err != nil {
		return
	}

	slotListMap := make(map[babytypes.ToySuitType][]*PlayerBabyToySlotObject)
	for _, slot := range toySlotList {
		pio := NewPlayerBabyToySlotObject(m.p)
		pio.FromEntity(slot)
		slotListMap[pio.suitType] = append(slotListMap[pio.suitType], pio)
	}

	m.toyBagMap = make(map[babytypes.ToySuitType]*BodyBag)
	for initType := babytypes.MinSuitType; initType <= babytypes.MaxSuitType; initType++ {
		slotList := slotListMap[initType]
		m.toyBagMap[initType] = createBodyBag(m.p, initType, slotList)
	}

	// 宝宝战力数据
	powerEntity, err := dao.GetBabyDao().GetBabyPowerEntity(m.p.GetId())
	if err != nil {
		return
	}
	if powerEntity == nil {
		m.initPlayerBabyPowerObject()
	} else {
		m.babyPowerObject = NewPlayerBabyPowerObject(m.p)
		m.babyPowerObject.FromEntity(powerEntity)
	}

	return nil
}

//初始化怀孕
func (m *PlayerBabyDataManager) initPlayerPregnantObject() {
	o := NewPlayerPregnantObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	o.id = id
	o.tonicPro = int32(0)
	o.pregnantTime = int64(0)
	o.createTime = now
	m.playerPregnantObject = o
	o.SetModified()
}

//初始化宝宝战力
func (m *PlayerBabyDataManager) initPlayerBabyPowerObject() {
	obj := NewPlayerBabyPowerObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	obj.createTime = now
	obj.SetModified()
	m.babyPowerObject = obj
}

//加载后
func (m *PlayerBabyDataManager) AfterLoad() (err error) {
	//m.hbRunner.AddTask(CreateCheckBabyBornTask(m.p))
	return nil
}

//心跳
func (m *PlayerBabyDataManager) Heartbeat() {
	//m.hbRunner.Heartbeat()
}

// 玩家洞房信息
func (m *PlayerBabyDataManager) GetPregnantInfo() *PlayerPregnantObject {
	return m.playerPregnantObject
}

// 是否怀孕
func (m *PlayerBabyDataManager) IsPregnant() bool {
	if m.playerPregnantObject.pregnantTime > 0 {
		return true
	}

	return false
}

// 怀孕
func (m *PlayerBabyDataManager) Pregnant() {
	now := global.GetGame().GetTimeService().Now()
	m.playerPregnantObject.pregnantTime = now
	m.playerPregnantObject.updateTime = now
	m.playerPregnantObject.SetModified()

	gameevent.Emit(babyeventtypes.EventTypeBabyPregnantChanged, m.p, nil)
}

/*
// 加速出生
func (m *PlayerBabyDataManager) AccelerateBorn() *PlayerBabyObject {
	if !m.IsPregnant() {
		return nil
	}

	return m.babyBorn()
}
*/

// 宝宝数量
func (m *PlayerBabyDataManager) GetBabyNum() int32 {
	return int32(len(m.playerBabyObjectList))
}

// 添加宝宝
func (m *PlayerBabyDataManager) AddBaby(quality, danbei int32, sex playertypes.SexType, talentList []*babytypes.TalentInfo) {
	pregnantTemplate := babytemplate.GetBabyTemplateService().GetBabyPregnantTemplateByQuality(quality)
	name := pregnantTemplate.GetBabyName(sex)
	m.initBaby(quality, danbei, sex, name, talentList)
}

// 宝宝超生
func (m *PlayerBabyDataManager) ChaoSheng() {
	now := global.GetGame().GetTimeService().Now()
	m.playerPregnantObject.chaoshengNum += 1
	m.playerPregnantObject.updateTime = now
	m.playerPregnantObject.SetModified()
}

// 是否超生
func (m *PlayerBabyDataManager) IsCanAddBaby() bool {
	nextBabyNum := m.GetBabyNum() + 1
	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	maxBabyLimit := babyConstantTemplate.BabyCount + m.playerPregnantObject.chaoshengNum
	if nextBabyNum > maxBabyLimit {
		return false
	}

	return true
}

// 吃补品
func (m *PlayerBabyDataManager) EatTonic(addPro int32) {
	maxPro := babytemplate.GetBabyTemplateService().GetBabyPregnantTonicMaxPro()
	remainPro := maxPro - m.playerPregnantObject.tonicPro
	if remainPro <= 0 {
		return
	}

	if addPro > remainPro {
		addPro = remainPro
	}

	now := global.GetGame().GetTimeService().Now()
	m.playerPregnantObject.tonicPro += addPro
	m.playerPregnantObject.updateTime = now
	m.playerPregnantObject.SetModified()

	gameevent.Emit(babyeventtypes.EventTypeBabyPregnantChanged, m.p, nil)
}

// 补品进度
func (m *PlayerBabyDataManager) IsFullTonic(addPro int32) bool {
	maxPro := babytemplate.GetBabyTemplateService().GetBabyPregnantTonicMaxPro()
	if m.playerPregnantObject.tonicPro+addPro > maxPro {
		return true
	}

	return false
}

// ------------宝宝操作---------------

// 玩家宝宝信息
func (m *PlayerBabyDataManager) GetBabyInfoList() []*PlayerBabyObject {
	return m.playerBabyObjectList
}

func (m *PlayerBabyDataManager) CheckBabyBorn() {
	if m.playerPregnantObject.pregnantTime <= 0 {
		return
	}

	// now := global.GetGame().GetTimeService().Now()
	// babyConstantTemp := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	// bornTime := babyConstantTemp.BornTime + m.playerPregnantObject.pregnantTime
	// leftTime := bornTime - now
	// if leftTime > 0 {
	// 	if m.lastNoticeTime == 0 {
	// 		m.lastNoticeTime = now
	// 	}

	// 	for _, noticeTime := range babyConstantTemp.GetBornNoticeTimeList() {
	// 		if m.lastNoticeTime <= int64(noticeTime) {
	// 			continue
	// 		}

	// 		if leftTime > int64(noticeTime) {
	// 			continue
	// 		}

	// 		evendData := babyeventtypes.CreatePlayerBabyBornMsgNoticeEventData(bornTime, int64(noticeTime))
	// 		gameevent.Emit(babyeventtypes.EventTypeBabyBornNotice, m.p, evendData)
	// 		m.lastNoticeTime = int64(noticeTime)
	// 		break
	// 	}

	// 	return
	// }
	qualityTemplate := babytemplate.GetBabyTemplateService().GetBabyQualityTemplate(m.playerPregnantObject.tonicPro)
	quality := qualityTemplate.RandomQuality()
	m.babyBorn(quality)
}

/*
//检查宝宝出生
func (m *PlayerBabyDataManager) CheckBabyBorn() {
	if m.playerPregnantObject.pregnantTime <= 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	babyConstantTemp := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	bornTime := babyConstantTemp.BornTime + m.playerPregnantObject.pregnantTime
	leftTime := bornTime - now
	if leftTime > 0 {
		if m.lastNoticeTime == 0 {
			m.lastNoticeTime = now
		}

		for _, noticeTime := range babyConstantTemp.GetBornNoticeTimeList() {
			if m.lastNoticeTime <= int64(noticeTime) {
				continue
			}

			if leftTime > int64(noticeTime) {
				continue
			}

			evendData := babyeventtypes.CreatePlayerBabyBornMsgNoticeEventData(bornTime, int64(noticeTime))
			gameevent.Emit(babyeventtypes.EventTypeBabyBornNotice, m.p, evendData)
			m.lastNoticeTime = int64(noticeTime)
			break
		}

		return
	}

	m.babyBorn()
}
*/

// 洞房结束,添加宝宝
func (m *PlayerBabyDataManager) AddBabyByDongFang() *PlayerBabyObject {
	babyTemp := babytemplate.GetBabyTemplateService().GetBabyTypeTemplate()
	if babyTemp == nil {
		return nil
	}

	babyObj := m.babyBorn(babyTemp.Quality)

	return babyObj
}

//宝宝出生
func (m *PlayerBabyDataManager) babyBorn(quality int32) *PlayerBabyObject {
	//qualityTemplate := babytemplate.GetBabyTemplateService().GetBabyQualityTemplate(m.playerPregnantObject.tonicPro)
	//quality := qualityTemplate.RandomQuality()
	pregnantTemplate := babytemplate.GetBabyTemplateService().GetBabyPregnantTemplateByQuality(quality)

	danbei := pregnantTemplate.GetRandonmBeiShu()
	sex := playertypes.RandomSex()
	name := pregnantTemplate.GetBabyName(sex)
	talentList := pregnantTemplate.GetInitRandonmTalentList()
	baby := m.initBaby(quality, danbei, sex, name, talentList)

	now := global.GetGame().GetTimeService().Now()
	m.playerPregnantObject.pregnantTime = 0
	m.playerPregnantObject.tonicPro = 0
	m.playerPregnantObject.updateTime = now
	m.playerPregnantObject.SetModified()

	gameevent.Emit(babyeventtypes.EventTypeBabyPregnantChanged, m.p, nil)
	gameevent.Emit(babyeventtypes.EventTypeBabyBorn, m.p, baby)
	return baby
}

func (m *PlayerBabyDataManager) initBaby(quality, danbei int32, sex playertypes.SexType, name string, talentList []*babytypes.TalentInfo) *PlayerBabyObject {
	now := global.GetGame().GetTimeService().Now()
	oldSkillList := m.GetEffectTalentSkillList()

	baby := NewPlayerBabyObject(m.p)
	id, _ := idutil.GetId()
	baby.id = id
	baby.sex = sex
	baby.name = name
	baby.quality = quality
	baby.skillList = talentList
	baby.learnLevel = 0
	baby.learnExp = 0
	baby.activateTimes = 0
	baby.lockTimes = 0
	baby.refreshTimes = 0
	baby.attrBeiShu = danbei
	baby.createTime = now
	baby.SetModified()
	m.playerBabyObjectList = append(m.playerBabyObjectList, baby)

	newSkillList := m.GetEffectTalentSkillList()
	eventData := babyeventtypes.CreatePlayerBabyTalentChangedEventData(baby.id, oldSkillList, newSkillList, baby.skillList)
	gameevent.Emit(babyeventtypes.EventTypeBabyTalentChanged, m.p, eventData)

	gameevent.Emit(babyeventtypes.EventTypeBabyAdd, m.p, baby)

	return baby
}

//宝宝转世
func (m *PlayerBabyDataManager) BabyZhuanShi(babyId int64) {
	now := global.GetGame().GetTimeService().Now()
	delIndex := 0
	for index, baby := range m.playerBabyObjectList {
		if baby.id != babyId {
			continue
		}
		delIndex = index
		baby.deleteTime = now
		baby.SetModified()
		break
	}

	m.playerBabyObjectList = append(m.playerBabyObjectList[:delIndex], m.playerBabyObjectList[delIndex+1:]...)

	m.playerPregnantObject.updateTime = now
	m.playerPregnantObject.SetModified()

	gameevent.Emit(babyeventtypes.EventTypeBabyZhuanShi, m.p, babyId)
}

// 宝宝改名
func (m *PlayerBabyDataManager) UpdateBabyName(babyId int64, newName string) {
	baby := m.getBabyObj(babyId)
	if baby == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	baby.name = newName
	baby.updateTime = now
	baby.SetModified()
}

// 天赋技能是否上限
func (m *PlayerBabyDataManager) IsFullTalent(babyId int64) bool {
	baby := m.getBabyObj(babyId)
	if baby == nil {
		return true
	}

	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	if len(baby.skillList) < int(babyConstantTemplate.LimitSkiiNum) {
		return false
	}

	return true
}

// 激活技能
func (m *PlayerBabyDataManager) ActivateSkill(babyId int64) (flag bool) {
	baby := m.getBabyObj(babyId)
	if baby == nil {
		return
	}

	pregnantTemp := babytemplate.GetBabyTemplateService().GetBabyPregnantTemplateByQuality(baby.quality)
	if pregnantTemp == nil {
		return
	}

	oldSkillList := m.GetEffectTalentSkillList()
	oldTalentList := baby.GetTalentSkillIdList()

	addTalent := pregnantTemp.GetRandonmTalent()
	now := global.GetGame().GetTimeService().Now()
	baby.skillList = append(baby.skillList, addTalent)
	baby.activateTimes += 1
	baby.updateTime = now
	baby.SetModified()

	newSkillList := m.GetEffectTalentSkillList()
	newTalentList := baby.GetTalentSkillIdList()

	eventData := babyeventtypes.CreatePlayerBabyTalentChangedEventData(babyId, oldSkillList, newSkillList, baby.skillList)
	gameevent.Emit(babyeventtypes.EventTypeBabyTalentChanged, m.p, eventData)

	reason := commonlog.BabyLogReasonTalentActivate
	reasonText := fmt.Sprintf(reason.String(), babyId)
	logEventData := babyeventtypes.CreatePlayerBabyTalentLogEventData(oldTalentList, newTalentList, []int32{addTalent.SkillId}, reason, reasonText)
	gameevent.Emit(babyeventtypes.EventTypeBabyTalentLog, m.p, logEventData)
	flag = true
	return
}

// 洗练技能
func (m *PlayerBabyDataManager) RefeshBabySkill(babyId int64) (flag bool) {
	baby := m.getBabyObj(babyId)
	if baby == nil {
		return
	}
	pregnantTemp := babytemplate.GetBabyTemplateService().GetBabyPregnantTemplateByQuality(baby.quality)
	if pregnantTemp == nil {
		return
	}

	oldSkillList := m.GetEffectTalentSkillList()
	oldTalentList := baby.GetTalentSkillIdList()

	var changedTalentList []int32
	for _, talentInfo := range baby.skillList {
		if talentInfo.Status == babytypes.SkillStatusTypeLock {
			continue
		}
		newTalentInfo := pregnantTemp.GetRandonmTalent()
		talentInfo.SkillId = newTalentInfo.SkillId
		talentInfo.Type = newTalentInfo.Type

		changedTalentList = append(changedTalentList, talentInfo.SkillId)
	}
	// 洗练消耗
	babyConstantTemplate := babytemplate.GetBabyTemplateService().GetBabyConstantTemplate()
	useItemMap := babyConstantTemplate.GetRefreshTalentUseItemMap(baby.refreshTimes)
	for _, num := range useItemMap {
		baby.refreshCostItemNum += num
	}
	now := global.GetGame().GetTimeService().Now()
	baby.refreshTimes += 1
	baby.updateTime = now
	baby.SetModified()

	newSkillList := m.GetEffectTalentSkillList()
	newTalentList := baby.GetTalentSkillIdList()

	eventData := babyeventtypes.CreatePlayerBabyTalentChangedEventData(babyId, oldSkillList, newSkillList, baby.skillList)
	gameevent.Emit(babyeventtypes.EventTypeBabyTalentChanged, m.p, eventData)

	reason := commonlog.BabyLogReasonTalentRefresh
	reasonText := fmt.Sprintf(reason.String(), babyId)
	logEventData := babyeventtypes.CreatePlayerBabyTalentLogEventData(oldTalentList, newTalentList, changedTalentList, reason, reasonText)
	gameevent.Emit(babyeventtypes.EventTypeBabyTalentLog, m.p, logEventData)
	flag = true
	return
}

// 获取技能信息
func (m *PlayerBabyDataManager) GetTalentInfo(babyId int64, skillIndex int32) *babytypes.TalentInfo {
	baby := m.getBabyObj(babyId)
	if baby == nil {
		return nil
	}

	for index, talent := range baby.skillList {
		if index != int(skillIndex) {
			continue
		}

		return talent
	}

	return nil
}

// 锁定技能
func (m *PlayerBabyDataManager) LockBabySkill(babyId int64, skillIndex int32, status babytypes.SkillStatusType) (flag bool) {
	baby := m.getBabyObj(babyId)
	if baby == nil {
		return
	}

	talent := m.GetTalentInfo(babyId, skillIndex)
	if talent == nil {
		return
	}

	if talent.Status == status {
		return
	}

	if status == babytypes.SkillStatusTypeLock {
		baby.lockTimes += 1
	}
	now := global.GetGame().GetTimeService().Now()
	talent.Status = status
	baby.updateTime = now
	baby.SetModified()

	flag = true
	return
}

// 宝宝读书
func (m *PlayerBabyDataManager) AddLearnExp(babyId int64, addLearnExp int32) (isUplevel bool) {
	baby := m.getBabyObj(babyId)
	if baby == nil {
		return
	}

	remainExp := baby.learnExp + addLearnExp
	oldLevel := baby.learnLevel
	newLevel := baby.learnLevel
	for {
		nextLearnTemp := babytemplate.GetBabyTemplateService().GetBabyLearnTemplate(newLevel + 1)
		if nextLearnTemp == nil {
			break
		}

		if remainExp < nextLearnTemp.Experience {
			break
		}

		remainExp -= nextLearnTemp.Experience
		newLevel += 1
	}

	if oldLevel != newLevel {
		baby.learnLevel = newLevel
		gameevent.Emit(babyeventtypes.EventTypeBabyLearnUplevel, m.p, baby)

		reason := commonlog.BabyLogReasonLearn
		reasonText := fmt.Sprintf(reason.String(), babyId)
		logEnventData := babyeventtypes.CreatePlayerBabyLevelLogEventData(oldLevel, newLevel, reason, reasonText)
		gameevent.Emit(babyeventtypes.EventTypeBabyLearnLog, m.p, logEnventData)
		isUplevel = true
	}

	now := global.GetGame().GetTimeService().Now()
	baby.learnExp = remainExp
	baby.updateTime = now
	baby.SetModified()

	return
}

// 宝宝信息
func (m *PlayerBabyDataManager) GetBabyInfo(babyId int64) *PlayerBabyObject {
	return m.getBabyObj(babyId)
}

// 宝宝玩具背包
func (m *PlayerBabyDataManager) GetBabyToyBag(suitType babytypes.ToySuitType) *BodyBag {
	return m.toyBagMap[suitType]
}

//玩具改变
func (m *PlayerBabyDataManager) GetChangedToySlotAndResetMap() map[babytypes.ToySuitType][]*PlayerBabyToySlotObject {
	changedMap := make(map[babytypes.ToySuitType][]*PlayerBabyToySlotObject)
	for suitType, bag := range m.toyBagMap {
		changedList := bag.GetChangedSlotAndReset()
		changedMap[suitType] = changedList
	}
	return changedMap
}

//获取所有玩具信息
func (m *PlayerBabyDataManager) GetAllToySlotMap() map[babytypes.ToySuitType][]*PlayerBabyToySlotObject {
	allSlotMap := make(map[babytypes.ToySuitType][]*PlayerBabyToySlotObject)
	for suitType, toyBag := range m.toyBagMap {
		allSlotMap[suitType] = toyBag.GetAll()
	}
	return allSlotMap
}

//获取所有套装数量
func (m *PlayerBabyDataManager) GetAllToyGroupNum() map[babytypes.ToySuitType]map[int32]int32 {
	allGroupMap := make(map[babytypes.ToySuitType]map[int32]int32)
	for suitType, tulongBag := range m.toyBagMap {
		allGroupMap[suitType] = tulongBag.CountSuitGroupNum()
	}

	return allGroupMap
}

//获取套装数量
func (m *PlayerBabyDataManager) GetToyGroupNumByType(suitType babytypes.ToySuitType) map[int32]int32 {
	toyBag := m.toyBagMap[suitType]
	return toyBag.CountSuitGroupNum()
}

//获取所有宝宝的有效天赋列表
func (m *PlayerBabyDataManager) GetEffectTalentSkillList() map[int32]int32 {
	loadSkillMap := make(map[int32]int32)
	for _, baby := range m.playerBabyObjectList {
		for _, talent := range baby.skillList {
			if talent.Type != babytypes.SkillTypeSkill {
				continue
			}

			// 计算所有可生效天赋技能
			skillTemp := skilltemplate.GetSkillTemplateService().GetSkillTemplate(talent.SkillId)
			if skillTemp == nil {
				continue
			}
			existSkillId, ok := loadSkillMap[skillTemp.TypeId]
			if !ok {
				loadSkillMap[skillTemp.TypeId] = talent.SkillId
				continue
			}

			//同组高级的生效
			existSkillTemp := skilltemplate.GetSkillTemplateService().GetSkillTemplate(existSkillId)
			if existSkillTemp.Lev >= skillTemp.Lev {
				continue
			}

			loadSkillMap[skillTemp.TypeId] = talent.SkillId
		}
	}

	return nil
}

func (m *PlayerBabyDataManager) getBabyObj(babyId int64) *PlayerBabyObject {
	for _, baby := range m.playerBabyObjectList {
		if baby.id != babyId {
			continue
		}

		return baby
	}

	return nil
}

//缓存使用
func (m *PlayerBabyDataManager) ToPregnantInfo() *babytypes.PregnantInfo {
	info := &babytypes.PregnantInfo{
		PregnantTime: m.playerPregnantObject.pregnantTime,
		TonicPro:     m.playerPregnantObject.tonicPro,
	}

	return info
}

func (m *PlayerBabyDataManager) GetPower() int64 {
	return m.babyPowerObject.power
}

func (m *PlayerBabyDataManager) SetPower(power int64) bool {
	if power < 0 {
		return false
	}

	if m.babyPowerObject.power == power {
		return false
	}

	now := global.GetGame().GetTimeService().Now()
	m.babyPowerObject.power = power
	m.babyPowerObject.updateTime = now
	m.babyPowerObject.SetModified()
	return true
}

// ---------配偶宝宝---------
//同步配偶所有宝宝
func (m *PlayerBabyDataManager) LoadAllCoupleBaby(babyDataList []*babytypes.CoupleBabyData) {
	m.coupleBabyList = babyDataList
}

//获取配偶宝宝列表
func (m *PlayerBabyDataManager) GetAllCoupleBabyList() []*babytypes.CoupleBabyData {
	return m.coupleBabyList
}

func CreatePlayerBabyDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerBabyDataManager{}
	m.p = p
	m.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	return m
}

func init() {
	player.RegisterPlayerDataManager(playertypes.PlayerBabyDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerBabyDataManager))
}
