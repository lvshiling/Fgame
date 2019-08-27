package chuangshi

import (
	"fgame/fgame/core/heartbeat"
	"fgame/fgame/core/runner"
	chuangshidao "fgame/fgame/cross/chuangshi/dao"
	chuangshiscene "fgame/fgame/cross/chuangshi/scene"
	activitytemplate "fgame/fgame/game/activity/template"
	activitytypes "fgame/fgame/game/activity/types"
	alliancetypes "fgame/fgame/game/alliance/types"
	chuangshidata "fgame/fgame/game/chuangshi/data"
	chuangshitemplate "fgame/fgame/game/chuangshi/template"
	chuangshitypes "fgame/fgame/game/chuangshi/types"
	"fgame/fgame/game/common/common"
	"fgame/fgame/game/global"
	"fgame/fgame/game/item/item"
	"fgame/fgame/game/scene/scene"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"math"
	"sort"
	"sync"

	log "github.com/Sirupsen/logrus"
)

// TODO xzk27  a)个人奖励为阵营内所有玩家均可以进行领取，当玩家有奖励未领取时，但却切换了阵营，则个人未领取的奖励直接清空
// 城防建筑在城池更换阵营时，随机一个城防建筑的等级会掉落A~B级掉级之外原先进度条多的经验全部置零处理，需要重新升回去
//	阵营更换卡
// 3)当城主所在的仙盟盟主发生改变时，由于城主仅能为仙盟盟主，因此该城城主对应替换为新的仙盟盟主(本服仙盟变更盟主同步过来)
// 当报名、投票期间，玩家更换了阵营，自动将玩家从报名名单以及神王候选人名单中移除

// 4)尊贵上线提醒：神王每次登陆时，将对阵营内所有玩家进行广告，将会在系统频道中刷出特殊边框的文本，文本内容为“我方阵营神王【玩家名称六字】上线了”
// 5)击杀提醒：当神王在创世之战地图被击杀时，则刷出系统公告“我方阵营神王【玩家名称六字】在【地图名称六字】被【玩家名称六字】击杀了，此等大辱岂可坐视不理   前往复仇”，点击前往复仇后进行进入对应地图判断
// 6)尊贵前往提醒：当神王每次前往创世之战对应地图时，则对本地图内所有玩家刷出系统功能，该公告需要与尊贵上线提醒一样的边框，文本内容为“我方神王【玩家名字六字】驾临了【地图名称六字】”

type ChuangShiService interface {
	runner.Task
	Start()
	Stop()
	// 城池阵营收益
	CityTask()
	// 神王竞选阶段
	ShenWangTask()

	//创世阵营
	GetChuangShiCampListList() []*Camp
	//获取创世城池信息
	GetChuangShiCityData(cityId int64) *CityData
	// 加入阵营
	JoinCamp(campType chuangshitypes.ChuangShiCampType, memList []*chuangshidata.MemberInfo) bool

	//神王竞选报名
	ChuangShiShenWangSignUp(campType chuangshitypes.ChuangShiCampType, playerId int64) (success bool, signList []*ChuangShiSignInfo)
	//神王投票
	ChuangShiShenWangVote(playerId int64, campType chuangshitypes.ChuangShiCampType, supportId int64) (success bool, voteList []*ChuangShiVoteInfo)
	//城主任命
	CityRenMing(playerId int64, cityId, beCommitId int64) (success bool)
	//城池工资分配
	CityPaySchedule(playerId int64, paramList []*chuangshidata.CityPayScheduleParam) (city *ChuangShiCityObject)
	//城池建设
	CityChengFangJianShe(playerId, cityId int64, jianSheType chuangshitypes.ChuangShiCityJianSheType, num int32) (success bool)
	//城池天气设置
	CityTianQiSet(playerId, cityId int64, level int32) bool
	//阵营工资分配
	CampPaySchedule(playerId int64, paramList []*chuangshidata.CamPayScheduleParam) (camp *Camp)
	//阵营工资领取
	CampPayReceive(playerId int64) (camp *Camp)

	//获取主城场景
	GetChuangShiMainScene(campType chuangshitypes.ChuangShiCampType) scene.Scene
	//获取中立场景
	GetChuangShiZhongLiScene() scene.Scene
	//获取附属场景
	GetChuangShiFuShuScene(cityId int64) scene.Scene
	//设置攻城目标（创建附属城场景）
	GongChengTargetFuShu(playerId, cityId int64) (flag bool)
}

type chuangShiService struct {
	rwm      sync.RWMutex
	hbRunner heartbeat.HeartbeatTaskRunner

	//阵营信息
	campMap map[chuangshitypes.ChuangShiCampType]*Camp
	//城池信息
	allCityMap map[int64]*CityData
	//阵营成员信息
	memberMap map[int64]*ChuangShiMemberObject
	// 投票记录
	voteRecordMap map[int64]*ChuangShiShenWangVoteRecordObject

	//主城
	mainSceneMap map[chuangshitypes.ChuangShiCampType]scene.Scene
	//中立
	zhongLiScene scene.Scene
	//附属城池
	fuShuSceneMap map[int64]scene.Scene
}

func (s *chuangShiService) init() (err error) {
	s.mainSceneMap = make(map[chuangshitypes.ChuangShiCampType]scene.Scene)

	//计算城市收益
	s.hbRunner = heartbeat.NewHeartbeatTaskRunner()
	s.hbRunner.AddTask(CreateCityTask(s))
	s.hbRunner.AddTask(CreateShenWangTask(s))

	//加载阵营
	err = s.loadCamps()
	if err != nil {
		return
	}

	//加载投票记录
	err = s.loadVoteRecord()
	if err != nil {
		return
	}

	return nil
}

func (s *chuangShiService) loadCamps() (err error) {
	s.campMap = make(map[chuangshitypes.ChuangShiCampType]*Camp)
	s.allCityMap = make(map[int64]*CityData)
	s.memberMap = make(map[int64]*ChuangShiMemberObject)

	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerId()
	eList, err := chuangshidao.GetChuangShiDao().GetCampList(platform, serverId)
	if err != nil {
		return
	}
	for _, e := range eList {
		campObj := NewChuangShiCampObject()
		err := campObj.FromEntity(e)
		if err != nil {
			return err
		}
		camp := createCamp(campObj)

		//成员
		memList, err := chuangshidao.GetChuangShiDao().GetChuangshiMemberList(platform, serverId, int32(campObj.campType))
		if err != nil {
			return err
		}
		for _, e := range memList {
			memObj := newChuangShiMemberObject(camp)
			memObj.FromEntity(e)

			s.addCampMember(memObj)
		}

		// 城池
		cList, err := chuangshidao.GetChuangShiDao().GetCityList(platform, serverId, int32(campObj.campType))
		if err != nil {
			return err
		}
		for _, c := range cList {
			cityObj := NewChuangShiCityObject(camp)
			cityObj.FromEntity(c)
			cityData := NewCityDatt()
			cityData.city = cityObj

			// 城池建设
			jiansheList, err := chuangshidao.GetChuangShiDao().GetCityJianSheList(platform, serverId, cityObj.id)
			if err != nil {
				return err
			}
			for _, e := range jiansheList {
				jianSheObj := NewChuangShiCityJianSheObject()
				err = jianSheObj.FromEntity(e)
				if err != nil {
					return err
				}
				cityData.jianSheList = append(cityData.jianSheList, jianSheObj)
			}

			flag := s.addCity(cityData)
			if !flag {
				return fmt.Errorf("chuangshi:加载重复的城市，阵营[%s],类型[%s],索引[%d]", cityObj.campType.String(), cityObj.typ.String(), cityObj.index)
			}
		}

		//报名
		signList, err := chuangshidao.GetChuangShiDao().GetShenWangSignUpList(platform, serverId, int32(campObj.campType))
		if err != nil {
			return err
		}
		for _, e := range signList {
			signObj := NewChuangShiShenWangSignUpObject()
			signObj.FromEntity(e)

			camp.addShenWangSign(signObj)
		}

		//投票
		voteList, err := chuangshidao.GetChuangShiDao().GetShenWangVoteList(platform, serverId, int32(campObj.campType))
		if err != nil {
			return err
		}
		for _, e := range voteList {
			voteObj := NewChuangShiShenWangVoteObject()
			voteObj.FromEntity(e)

			camp.addShenWangVote(voteObj)
		}

		s.addCamp(camp)
	}

	//初始化阵营
	s.initCamps()

	//初始化城市
	err = s.complementCities()
	if err != nil {
		return err
	}
	return
}

func (s *chuangShiService) initCamps() {
	now := global.GetGame().GetTimeService().Now()
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerId()

	for campType, _ := range chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCampTempAll() {
		camp := s.getCamp(campType)
		if camp != nil {
			continue
		}

		campObj := NewChuangShiCampObject()
		campObj.id, _ = idutil.GetId()
		campObj.platform = platform
		campObj.serverId = serverId
		campObj.campType = campType
		campObj.lastShouYiTime = now
		campObj.createTime = now
		campObj.targetCityMap = map[chuangshitypes.ChuangShiCampType]int64{}
		campObj.shenWangStatus = chuangshitypes.ShenWangStatusTypeEnd
		campObj.SetModified()

		camp = createCamp(campObj)
		s.addCamp(camp)
	}
}

func (s *chuangShiService) loadVoteRecord() (err error) {
	s.voteRecordMap = make(map[int64]*ChuangShiShenWangVoteRecordObject)
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerId()
	eList, err := chuangshidao.GetChuangShiDao().GetShenWangVoteRecordList(platform, serverId)
	if err != nil {
		return
	}

	for _, e := range eList {
		obj := NewChuangShiShenWangVoteRecordObject()
		obj.FromEntity(e)
		s.voteRecordMap[obj.playerId] = obj
	}

	return
}

func (s *chuangShiService) addCampMember(memberObj *ChuangShiMemberObject) (flag bool) {
	tmemObj := s.getCampMember(memberObj.playerId)
	if tmemObj != nil {
		return
	}

	camp := memberObj.GetCamp()
	if camp == nil {
		return
	}

	flag = camp.addMember(memberObj)
	if !flag {
		return
	}

	s.memberMap[memberObj.playerId] = memberObj
	return
}

func (s *chuangShiService) getCampMember(memberId int64) *ChuangShiMemberObject {
	mem, ok := s.memberMap[memberId]
	if !ok {
		return nil
	}

	return mem
}

func (s *chuangShiService) addCamp(camp *Camp) {
	s.campMap[camp.campObj.campType] = camp
}

func (s *chuangShiService) getCamp(campType chuangshitypes.ChuangShiCampType) *Camp {
	camp, ok := s.campMap[campType]
	if !ok {
		return nil
	}
	return camp
}

func (s *chuangShiService) getCity(cityId int64) *CityData {
	city, ok := s.allCityMap[cityId]
	if !ok {
		return nil
	}
	return city
}

func (s *chuangShiService) getCityByType(campType chuangshitypes.ChuangShiCampType, cityType chuangshitypes.ChuangShiCityType, index int32) *CityData {
	for _, cityData := range s.allCityMap {
		if cityData.city.originalCamp != campType {
			continue
		}

		if cityData.city.typ != cityType {
			continue
		}

		if cityData.city.index != index {
			continue
		}

		return cityData
	}
	return nil
}

func (s *chuangShiService) addCity(cityData *CityData) (flag bool) {
	cityObj := cityData.city
	if cityObj == nil {
		return
	}

	curCityData := s.getCity(cityObj.id)
	if curCityData != nil {
		return
	}

	camp := cityObj.GetCamp()
	if camp == nil {
		return
	}

	camp.addCity(cityData)
	s.allCityMap[cityObj.id] = cityData
	flag = true
	return
}

// 初始化城市
func (s *chuangShiService) complementCities() (err error) {
	now := global.GetGame().GetTimeService().Now()
	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerId()
	for campType, campMap := range chuangshitemplate.GetChuangShiTemplateService().GetAllChuangShiCityTemplate() {
		for typ, cityMap := range campMap {
			for index, _ := range cityMap {
				cityData := s.getCityByType(campType, typ, index)
				if cityData != nil {
					continue
				}
				camp := s.getCamp(campType)
				if camp == nil {
					return fmt.Errorf("chuangshi:初始化城市，阵营不存在，campType:%d", campType)
				}
				cityObj := NewChuangShiCityObject(camp)
				cityObj.id, _ = idutil.GetId()
				cityObj.platform = platform
				cityObj.serverId = serverId
				cityObj.campType = campType
				cityObj.originalCamp = campType
				cityObj.typ = typ
				cityObj.index = index
				cityObj.createTime = now
				cityObj.SetModified()

				cityData = NewCityDatt()
				cityData.city = cityObj

				for jianSheType, _ := range chuangshitemplate.GetChuangShiTemplateService().GetAllChuangShiChengFangTemp() {
					jianSheObj := NewChuangShiCityJianSheObject()
					jianSheObj.id, _ = idutil.GetId()
					jianSheObj.platform = platform
					jianSheObj.serverId = serverId
					jianSheObj.cityId = cityObj.id
					jianSheObj.jianSheType = jianSheType
					jianSheObj.createTime = now
					jianSheObj.SetModified()

					cityData.jianSheList = append(cityData.jianSheList, jianSheObj)
				}

				flag := s.addCity(cityData)
				if !flag {
					return fmt.Errorf("chuangshi:初始化重复的城市，阵营[%s],类型[%s],索引[%d]", cityObj.campType.String(), cityObj.typ.String(), cityObj.index)
				}
			}
		}
	}
	return nil
}

func (s *chuangShiService) ChuangShiShenWangSignUp(campType chuangshitypes.ChuangShiCampType, playerId int64) (success bool, signList []*ChuangShiSignInfo) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	constantTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangshiConstantTemp()
	if !constantTemp.IfShenWangSign(now) {
		return
	}

	camp := s.getCamp(campType)
	if camp == nil {
		return
	}

	sign := camp.getShenWangSign(playerId)
	if sign != nil {
		return
	}

	id, _ := idutil.GetId()
	obj := NewChuangShiShenWangSignUpObject()
	obj.id = id
	obj.playerId = playerId
	obj.createTime = now
	obj.SetModified()
	camp.addShenWangSign(obj)

	signList = camp.GetShenWangSignList()
	success = true
	return
}

func (s *chuangShiService) ChuangShiShenWangVote(playerId int64, campType chuangshitypes.ChuangShiCampType, supportId int64) (success bool, voteList []*ChuangShiVoteInfo) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	mem := s.getCampMember(playerId)
	if mem == nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	constantTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangshiConstantTemp()
	if !constantTemp.IfShenWangVote(now) {
		return
	}

	camp := s.getCamp(campType)
	if camp == nil {
		return
	}

	voteObj := camp.getShenWangVote(supportId)
	if voteObj == nil {
		return
	}

	voteRecord, ok := s.voteRecordMap[playerId]
	if !ok {
		platform := global.GetGame().GetPlatform()
		serverId := global.GetGame().GetServerIndex()
		voteRecord = NewChuangShiShenWangVoteRecordObject()
		voteRecord.platform = platform
		voteRecord.serverId = serverId
		voteRecord.campType = campType
		voteRecord.playerId = playerId
		voteRecord.playerName = mem.playerName
		voteRecord.playerServerId = mem.playerServerId
		voteRecord.createTime = now
		voteRecord.SetModified()
		s.voteRecordMap[playerId] = voteRecord
	}

	if !voteRecord.IfCanVote(now) {
		return
	}

	voteRecord.houXuanPlatform = voteObj.Member.playerPlatform
	voteRecord.houXuanGameServerId = voteObj.Member.playerServerId
	voteRecord.houXuanPlayerId = voteObj.Member.playerId
	voteRecord.houXuanPlayerName = voteObj.Member.playerName
	voteRecord.lastVoteTime = now
	voteRecord.updateTime = now
	voteRecord.SetModified()

	voteObj.Vote.ticketNum += 1
	voteObj.Vote.updateTime = now
	voteObj.Vote.SetModified()
	voteList = camp.GetShenWangVoteList()
	success = true
	return
}

func (s *chuangShiService) GetChuangShiCampListList() (campList []*Camp) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	for _, camp := range s.campMap {
		campList = append(campList, camp)
	}

	return campList
}

func (s *chuangShiService) GetChuangShiCityData(cityId int64) (cityData *CityData) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	cityData, ok := s.allCityMap[cityId]
	if !ok {
		return nil
	}

	return cityData
}

func (s *chuangShiService) CityRenMing(playerId int64, cityId, beCommitId int64) (success bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	//
	mem := s.getCampMember(playerId)
	if mem == nil {
		return
	}

	if !mem.IfShenWang() {
		return
	}
	camp := mem.GetCamp()

	// 被任命者不是本阵营
	beCommitMem := camp.getMember(beCommitId)
	if beCommitMem == nil {
		return
	}

	//被任命者不是盟主
	if beCommitMem.alPos != alliancetypes.AlliancePositionMengZhu {
		return
	}

	//被任命者已经是城主
	orignalCity := camp.getCityByChengZhuId(beCommitId)
	if orignalCity != nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	commitCity := s.getCity(cityId)
	ownerMem := camp.getMember(commitCity.city.ownerId)
	if ownerMem != nil {
		//原城主移除官职
		ownerMem.pos = chuangshitypes.ChuangShiGuanZhiPingMing
		ownerMem.updateTime = now
		ownerMem.SetModified()
	}

	//任命
	commitCity.city.ownerId = beCommitId
	commitCity.city.updateTime = now
	commitCity.city.SetModified()

	beCommitMem.pos = chuangshitypes.ChuangShiGuanZhiChengZhu
	beCommitMem.updateTime = now
	beCommitMem.SetModified()

	success = true
	return
}

func (s *chuangShiService) CampPaySchedule(playerId int64, paramList []*chuangshidata.CamPayScheduleParam) (camp *Camp) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	//
	mem := s.getCampMember(playerId)
	if mem == nil {
		return
	}

	if !mem.IfShenWang() {
		return
	}

	camp = mem.GetCamp()
	remainJifen := camp.campObj.jifen
	remainDiamonds := camp.campObj.diamonds
	if remainJifen <= 0 && remainDiamonds <= 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	for _, param := range paramList {
		cityData := camp.getCity(param.CityId)
		if cityData == nil {
			continue
		}

		addJifen := int64(math.Floor(float64(int64(param.Ratio)*camp.campObj.jifen) / float64(common.MAX_RATE)))
		addDiamonds := int64(math.Floor(float64(int64(param.Ratio)*camp.campObj.diamonds) / float64(common.MAX_RATE)))

		cityData.city.jifen += addJifen
		cityData.city.diamonds += addDiamonds
		cityData.city.updateTime = now
		cityData.city.SetModified()

		remainJifen -= addJifen
		remainDiamonds -= addDiamonds
	}

	camp.campObj.jifen = remainJifen
	camp.campObj.diamonds = remainDiamonds
	camp.campObj.updateTime = now
	camp.campObj.SetModified()
	return
}

func (s *chuangShiService) CampPayReceive(playerId int64) (camp *Camp) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	//
	mem := s.getCampMember(playerId)
	if mem == nil {
		return
	}

	if !mem.IfShenWang() {
		return
	}

	camp = mem.GetCamp()
	addJifen := camp.campObj.payJifen
	addDiamonds := camp.campObj.payDiamonds

	now := global.GetGame().GetTimeService().Now()
	camp.campObj.jifen += addJifen
	camp.campObj.diamonds += addDiamonds
	camp.campObj.payJifen = 0
	camp.campObj.payDiamonds = 0
	camp.campObj.updateTime = now
	camp.campObj.SetModified()
	return
}

func (s *chuangShiService) CityPaySchedule(playerId int64, paramList []*chuangshidata.CityPayScheduleParam) (city *ChuangShiCityObject) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	//
	cityKing := s.getCampMember(playerId)
	if cityKing == nil {
		return
	}

	camp := cityKing.GetCamp()
	cityData := camp.getCityByChengZhuId(playerId)
	if cityData == nil {
		return
	}

	city = cityData.city
	remainJifen := city.jifen
	remainDiamonds := city.diamonds
	if remainJifen <= 0 && remainDiamonds <= 0 {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	for _, mem := range camp.GetMemberList() {
		if mem.allianceId != cityKing.allianceId {
			continue
		}

		ratio := chuangshidata.CityPayScheduleList(paramList).GetPayRatio(mem.alPos)
		if ratio == 0 {
			continue
		}

		addJifen := int64(math.Floor(float64(int64(ratio)*city.jifen) / float64(common.MAX_RATE)))
		addDiamonds := int64(math.Floor(float64(int64(ratio)*city.diamonds) / float64(common.MAX_RATE)))
		mem.scheduleJifen += addJifen
		mem.scheduleDiamonds += addDiamonds
		mem.updateTime = now
		mem.SetModified()

		remainJifen -= addJifen
		remainDiamonds -= addDiamonds
	}

	city.jifen = remainJifen
	city.diamonds = remainDiamonds
	city.updateTime = now
	city.SetModified()
	return
}

func (s *chuangShiService) CityChengFangJianShe(playerId, cityId int64, jianSheType chuangshitypes.ChuangShiCityJianSheType, num int32) (success bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	mem := s.getCampMember(playerId)
	if mem == nil {
		return
	}

	city := s.getCity(cityId)
	if city == nil {
		return
	}

	jianSheObj := city.GetChengFangJianShe(jianSheType)
	jianSheTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiChengFangTemp(jianSheType)
	if jianSheTemp == nil {
		return
	}
	jianSheLevelTemp := jianSheTemp.GetJianSheLevelTemp(jianSheObj.jianSheLevel)
	if jianSheLevelTemp == nil {
		return
	}

	itemTemp := item.GetItemService().GetItem(int(jianSheTemp.LevelItemId))
	addExp := itemTemp.TypeFlag1 * num

	newLevel := jianSheObj.jianSheLevel
	newExp := jianSheObj.jianSheExp
	newExp += addExp
	nextLevelTemp := jianSheLevelTemp.GetNextTemp()
	for nextLevelTemp != nil {
		if newExp < nextLevelTemp.NeedExp {
			break
		}

		newExp -= nextLevelTemp.NeedExp
		newLevel += 1
		nextLevelTemp = nextLevelTemp.GetNextTemp()
	}

	now := global.GetGame().GetTimeService().Now()
	jianSheObj.jianSheLevel = newLevel
	jianSheObj.jianSheExp = newExp
	jianSheObj.updateTime = now
	jianSheObj.SetModified()
	return
}

func (s *chuangShiService) CityTianQiSet(playerId, cityId int64, level int32) (success bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	mem := s.getCampMember(playerId)
	if mem == nil {
		return
	}

	cityData := s.getCity(cityId)
	if cityData == nil {
		return
	}

	//不是神王也不是当前城主
	if !mem.IfShenWang() && cityData.city.ownerId != playerId {
		return
	}

	jianSheObj := cityData.GetChengFangJianShe(chuangshitypes.ChuangShiCityJianSheTypeTianQi)
	if jianSheObj.skillLevelSet == level {
		success = true
		return
	}

	if jianSheObj.jianSheLevel < level {
		return
	}

	//激活
	now := global.GetGame().GetTimeService().Now()
	if !jianSheObj.IfActivate(level) {
		jianSheObj.skillMap[level] = now
	}

	// 设置为当前技能
	jianSheObj.skillLevelSet = level
	jianSheObj.updateTime = now
	jianSheObj.SetModified()

	success = true
	return
}

func (s *chuangShiService) GetChuangShiMainScene(campType chuangshitypes.ChuangShiCampType) scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	sc, ok := s.mainSceneMap[campType]
	if !ok {
		return nil
	}

	return sc
}

func (s *chuangShiService) GetChuangShiZhongLiScene() scene.Scene {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	return s.zhongLiScene
}

func (s *chuangShiService) GetChuangShiFuShuScene(cityId int64) scene.Scene {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	sc, ok := s.fuShuSceneMap[cityId]
	if !ok {
		return nil
	}

	return sc
}

func (s *chuangShiService) GongChengTargetFuShu(playerId, targetCityId int64) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	operateMem := s.getCampMember(playerId)
	if operateMem == nil {
		return
	}

	// 不是神王
	if operateMem.IfShenWang() {
		return
	}

	// 城池不存在
	targetCityData := s.getCity(targetCityId)
	if targetCityData == nil {
		return
	}

	// 不是附属城池
	if !targetCityData.ifFuShu() {
		return
	}
	targetCampType := targetCityData.city.campType
	targetOrignalCampType := targetCityData.city.originalCamp
	targetCityType := targetCityData.city.typ
	targetIndex := targetCityData.city.index

	//敌方两个阵营各一个
	camp := operateMem.GetCamp()
	if !camp.ifCanTarget(targetCampType) {
		return
	}

	// 临近目标
	// TODO xzk27 封装到模板里面
	_ = chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCityTemp(targetOrignalCampType, targetCityType, targetIndex)

	// 创建攻城场景
	now := global.GetGame().GetTimeService().Now()
	_, ok := s.fuShuSceneMap[targetCityId]
	if !ok {
		activityTemplate := activitytemplate.GetActivityTemplateService().GetActiveByType(activitytypes.ActivityTypeChuangShiZhiZhan)
		activityTimeTemplate, err := activityTemplate.GetActivityTimeTemplate(now, 0, 0)
		if err != nil {
			return
		}

		if activityTimeTemplate == nil {
			log.WithFields(
				log.Fields{
					"playerId": playerId,
					"now":      now,
					"type":     activitytypes.ActivityTypeChuangShiZhiZhan,
				}).Warnln("chuangshi:处理玩家进入城池错误,活动时间模板不存在")
			return
		}

		endTime, _ := activityTimeTemplate.GetEndTime(now)
		cityTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCityTemp(targetCampType, targetCityType, targetIndex)
		sc := chuangshiscene.CreateFuShuSceneData(cityTemp.GetMapId(), targetCampType, endTime)
		if sc != nil {
			s.fuShuSceneMap[targetCityId] = sc
		}
	}

	//
	camp.campObj.targetCityMap[targetCampType] = targetCityId
	camp.campObj.updateTime = now
	camp.campObj.SetModified()
	return
}

func (s *chuangShiService) JoinCamp(campType chuangshitypes.ChuangShiCampType, memList []*chuangshidata.MemberInfo) (flag bool) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	camp := s.getCamp(campType)
	if camp == nil {
		return
	}

	for _, mem := range memList {
		campMem := s.getCampMember(mem.PlayerId)
		if campMem != nil {
			return
		}
	}

	platform := global.GetGame().GetPlatform()
	serverId := global.GetGame().GetServerIndex()
	now := global.GetGame().GetTimeService().Now()
	for _, mem := range memList {
		id, _ := idutil.GetId()
		memObj := newChuangShiMemberObject(camp)
		memObj.id = id
		memObj.platform = platform
		memObj.serverId = serverId
		memObj.campType = campType
		memObj.playerPlatform = mem.Platform
		memObj.playerServerId = mem.ServerId
		memObj.playerId = mem.PlayerId
		memObj.playerName = mem.PlayerName
		memObj.force = mem.Force
		memObj.pos = chuangshitypes.ChuangShiGuanZhiPingMing
		memObj.alPos = mem.AlPos
		memObj.allianceId = mem.AllianceId
		memObj.allianceName = mem.AllianceName
		memObj.createTime = now
		memObj.SetModified()
		s.addCampMember(memObj)
	}

	flag = true
	return
}

//计算收益
func (s *chuangShiService) CityTask() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	constantTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangshiConstantTemp()
	for _, camp := range s.campMap {
		rewCount, newLastRewTime := constantTemp.RewCout(camp.campObj.lastShouYiTime, now)
		if rewCount <= 0 {
			continue
		}

		addJifen := int32(0)
		addDiamonds := int32(0)
		for _, cityData := range camp.cityList {
			city := cityData.city
			cityTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangShiCityTemp(city.originalCamp, city.typ, city.index)
			if cityTemp == nil {
				log.WithFields(log.Fields{
					"campType": city.originalCamp,
					"cityType": city.typ,
					"index":    city.index,
				}).Warningln("chuangshi：城池定时产出错误")
				continue
			}
			addJifen += cityTemp.PlayerRewJifen * rewCount
			addDiamonds += cityTemp.PlayerRewZuanshi * rewCount
		}

		camp.campObj.payJifen += int64(addJifen)
		camp.campObj.payDiamonds += int64(addDiamonds)
		camp.campObj.lastShouYiTime = newLastRewTime
		camp.campObj.updateTime = now
		camp.campObj.SetModified()
	}
}

//神王阶段
func (s *chuangShiService) ShenWangTask() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	now := global.GetGame().GetTimeService().Now()
	constantTemp := chuangshitemplate.GetChuangShiTemplateService().GetChuangshiConstantTemp()
	for _, camp := range s.campMap {
		var status chuangshitypes.ShenWangStatusType
		switch camp.campObj.shenWangStatus {
		case chuangshitypes.ShenWangStatusTypeSign:
			{
				if !constantTemp.IfShenWangVote(now) {
					continue
				}
				status = chuangshitypes.ShenWangStatusTypeVote

				sort.Sort(sort.Reverse(signList(camp.signUpList)))
				for index, sign := range camp.signUpList {
					if index >= voteMaxLen {
						break
					}

					id, _ := idutil.GetId()
					voteObj := NewChuangShiShenWangVoteObject()
					voteObj.id = id
					voteObj.platform = sign.Sign.platform
					voteObj.serverId = sign.Sign.serverId
					voteObj.playerId = sign.Sign.playerId
					voteObj.playerServerId = sign.Sign.playerServerId
					voteObj.ticketNum = 0
					voteObj.campType = sign.Sign.campType
					voteObj.updateTime = now
					voteObj.SetModified()

					vote := &ChuangShiVoteInfo{}
					vote.Member = sign.Member
					vote.Vote = voteObj
					camp.voteList = append(camp.voteList, vote)
				}
			}
		case chuangshitypes.ShenWangStatusTypeVote:
			{
				status = chuangshitypes.ShenWangStatusTypeEnd
				if len(camp.voteList) != 0 {
					// 神王变更
					sort.Sort(sort.Reverse(voteList(camp.voteList)))
					shenWangVote := camp.voteList[0]
					shenWang := camp.getMember(shenWangVote.Vote.playerId)
					shenWang.pos = chuangshitypes.ChuangShiGuanZhiShenWang
					shenWang.updateTime = now
					shenWang.SetModified()

					oldShenWang := camp.getMember(camp.campObj.kingId)
					if oldShenWang != nil {
						oldShenWang.pos = chuangshitypes.ChuangShiGuanZhiPingMing
						oldShenWang.updateTime = now
						oldShenWang.SetModified()
					}

					//TODO xzk27 与神王同仙盟的盟主自动为主城城主，无法被任命为附属城城主，当神王无仙盟时，则自己同时为主城城主

					camp.campObj.kingId = shenWang.playerId
				}
			}
		case chuangshitypes.ShenWangStatusTypeEnd:
			{
				if !constantTemp.IfShenWangSign(now) {
					continue
				}
				status = chuangshitypes.ShenWangStatusTypeSign
			}
		}

		camp.campObj.shenWangStatus = status
		camp.campObj.updateTime = now
		camp.campObj.SetModified()
	}
}

func (s *chuangShiService) Heartbeat() {
	s.hbRunner.Heartbeat()
}

func (s *chuangShiService) Start() {
	s.createCityScene()
	return
}

func (s *chuangShiService) createCityScene() {
	for campType, campMap := range chuangshitemplate.GetChuangShiTemplateService().GetAllChuangShiCityTemplate() {
		for _, cityMap := range campMap {
			for _, cityTemp := range cityMap {
				mapId := cityTemp.GetMapId()
				endTime := int64(0)

				switch cityTemp.GetCityType() {
				case chuangshitypes.ChuangShiCityTypeMain:
					{
						sc := chuangshiscene.CreateMainSceneData(mapId, endTime)
						s.mainSceneMap[campType] = sc
					}
				case chuangshitypes.ChuangShiCityTypeZhongli:
					{
						sc := chuangshiscene.CreateZhongLiSceneData(mapId, endTime)
						s.zhongLiScene = sc
					}
				default:
					continue
				}
			}
		}
	}
}

func (s *chuangShiService) Stop() {
	return
}

var (
	once sync.Once
	as   *chuangShiService
)

func Init() (err error) {
	once.Do(func() {
		as = &chuangShiService{}
		err = as.init()
	})
	return err
}

func GetChuangShiService() ChuangShiService {
	return as
}
