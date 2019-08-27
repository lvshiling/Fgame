package robot

import (
	"context"
	"fgame/fgame/common/message"
	coretypes "fgame/fgame/core/types"
	"fgame/fgame/game/battle/battle"
	battlecommon "fgame/fgame/game/battle/common"
	dummytemplate "fgame/fgame/game/dummy/template"
	"fgame/fgame/game/fashion/fashion"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingtong/lingtong"
	lingtongtemplate "fgame/fgame/game/lingtong/template"
	marrytypes "fgame/fgame/game/marry/types"
	"fgame/fgame/game/player"
	playercommon "fgame/fgame/game/player/common"
	"fgame/fgame/game/player/types"
	playertypes "fgame/fgame/game/player/types"
	propertytypes "fgame/fgame/game/property/types"
	robottemplate "fgame/fgame/game/robot/template"
	robottypes "fgame/fgame/game/robot/types"
	"fgame/fgame/game/scene/scene"
	scenetypes "fgame/fgame/game/scene/types"
	skillcommon "fgame/fgame/game/skill/common"
	skilltemplate "fgame/fgame/game/skill/template"
	skilltypes "fgame/fgame/game/skill/types"
	"fgame/fgame/game/weapon/weapon"
	"fgame/fgame/pkg/idutil"
	"fmt"
	"sync"
)

type RobotService interface {
	CreateQuestRobot(beginQuestId int32, endQuestId int32, properties map[propertytypes.BattlePropertyType]int64, power int64, showServerId bool) scene.RobotPlayer
	// CreateClientTestRobot(copyPlayer scene.Player, showServerId bool)  scene.RobotPlayer
	CreateClientTestRobot(properties map[propertytypes.BattlePropertyType]int64, power int64, showServerId bool) scene.RobotPlayer
	CreateArenaRobot(serverId int32, properties map[propertytypes.BattlePropertyType]int64, reliveTime int32, power int64) scene.RobotPlayer
	CreateArenapvpRobot(platform int32, serverId int32, properties map[propertytypes.BattlePropertyType]int64, reliveTime int32, power int64) scene.RobotPlayer
	CreateTeamCopyRobot(serverId int32, properties map[propertytypes.BattlePropertyType]int64, reliveTime int32, power int64) scene.RobotPlayer
	CreateOneArenaRobot(info *playercommon.PlayerInfo, power int64, showServerId bool) scene.RobotPlayer
	CreateModelRobot(serverId int32, sexType playertypes.SexType, roleType playertypes.RoleType, fashionId int32, weaponId int32, wingId int32, showServerId bool) scene.RobotPlayer
	CreateMarryRecommentRobot(pl player.Player, num int32, showServerId bool) (pList []scene.RobotPlayer)
	GMClear()
	GetRobot(id int64) scene.RobotPlayer
	RemoveRobot(id int64)
	GetNumOfRobot() int32
}

type robotService struct {
	rwm       sync.RWMutex
	currentId int32
	robotMap  map[int64]scene.RobotPlayer
}

func (s *robotService) init() error {
	s.robotMap = make(map[int64]scene.RobotPlayer)
	return nil
}

func ConvertFromRobotPlayer(pl scene.RobotPlayer) *playercommon.PlayerInfo {
	info := &playercommon.PlayerInfo{
		PlayerId:    pl.GetId(),
		Name:        pl.GetName(),
		Role:        pl.GetRole(),
		Sex:         pl.GetSex(),
		Level:       pl.GetLevel(),
		Force:       pl.GetForce(),
		TeamId:      pl.GetTeamId(),
		OnlineState: playertypes.PlayerOnlineStateOnline,
		FashionId:   pl.GetFashionId(),
	}
	return info
}

func (s *robotService) GetNumOfRobot() int32 {
	return int32(len(s.robotMap))
}

func (s *robotService) CreateMarryRecommentRobot(pl player.Player, num int32, showServerId bool) (pList []scene.RobotPlayer) {
	sex := playertypes.SexTypeMan
	curSex := pl.GetSex()
	if curSex == sex {
		sex = playertypes.SexTypeWoman
	}

	for i := int32(0); i < num; i++ {
		id, _ := idutil.GetId()
		userId, _ := idutil.GetId()
		sexType := sex
		roleType := playertypes.RandomRole()
		name := dummytemplate.GetDummyTemplateService().GetRandomDummyName()
		serverId := global.GetGame().GetServerIndex()

		po := playercommon.NewPlayerCommonObject(id, userId, serverId, name, roleType, sexType, false)

		mountId := int32(0)
		mountAdvanceId := int32(0)
		mountHidden := true
		fourGodKey := int32(0)
		realm := int32(1)
		spouse := ""
		spouseId := int64(0)
		weddingStatus := int32(0)
		ringType := int32(0)
		ringLevel := int32(0)
		//展示
		randomFashionTemplate := fashion.GetFashionService().RandomFashionTemplate()
		fashionId := int32(randomFashionTemplate.TemplateId())
		weaponId := int32(weapon.GetWeaponService().RandomWeaponTemplate().TemplateId())
		weaponState := int32(0)

		titleId := int32(0)
		wingId := int32(0)
		lingYuId := int32(0)
		shenFaId := int32(0)
		faBaoId := int32(0)
		petId := int32(0)
		xianTiId := int32(0)
		baGua := int32(0)
		flyPetId := int32(0)
		developLevel := int32(0)
		shenYuKey := int32(0)
		showObj := battle.CreatePlayerShowObject(
			fashionId,
			weaponId,
			weaponState,
			titleId,
			wingId,
			mountId,
			mountAdvanceId,
			mountHidden,
			shenFaId,
			lingYuId,
			fourGodKey,
			realm,
			spouse,
			spouseId,
			weddingStatus,
			ringType,
			ringLevel,
			faBaoId,
			petId,
			xianTiId,
			baGua,
			flyPetId,
			developLevel,
			shenYuKey,
		)

		vip := pl.GetVip()
		level := pl.GetLevel()
		zhuansheng := pl.GetZhuanSheng()
		soulAwakenNum := pl.GetSoulAwakenNum()
		//TODO: zrc修改战力
		power := int64(0)
		playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, level, zhuansheng, soulAwakenNum, false)
		p := createRobotPlayer(robottypes.RobotTypeArena, po, showObj, nil, nil, nil, 0, playerBattleObject, showServerId, power)
		pList = append(pList, p)
	}

	return
}

func (s *robotService) CreateOneArenaRobot(playInfo *playercommon.PlayerInfo, power int64, showServerId bool) scene.RobotPlayer {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	id, _ := idutil.GetId()
	userId, _ := idutil.GetId()
	sexType := playInfo.Sex
	roleType := playInfo.Role
	name := playInfo.Name
	serverId := playInfo.ServerId

	po := playercommon.NewPlayerCommonObject(id, userId, serverId, name, roleType, sexType, false)

	mountId := int32(0)
	mountAdvanceId := int32(0)
	mountHidden := true
	fourGodKey := int32(0)
	realm := playInfo.RealmLevel
	spouse := playInfo.MarryInfo.SpouseName
	spouseId := playInfo.MarryInfo.SpouseId
	if spouseId != 0 {
		marryStatus := marrytypes.MarryStatusType(playInfo.MarryInfo.Status)
		switch marryStatus {
		case marrytypes.MarryStatusTypeMarried:
			{
				if playInfo.Sex == types.SexTypeMan {
					spouse += "的丈夫"
				} else {
					spouse += "的妻子"
				}
			}
		case marrytypes.MarryStatusTypeProposal,
			marrytypes.MarryStatusTypeEngagement:
			{
				if playInfo.Sex == types.SexTypeMan {
					spouse += "的未婚夫"
				} else {
					spouse += "的未婚妻"
				}
			}
		}
	}
	weddingStatus := int32(0)
	ringType := int32(0)
	ringLevel := int32(0)
	skillList := make([]skillcommon.SkillObject, 0, 16)
	for _, ski := range playInfo.SkillList {
		tempSki := skillcommon.CreateSkillObject(ski.GetSkillId(), ski.GetLevel(), nil)
		skillList = append(skillList, tempSki)
	}

	//展示
	fashionId := playInfo.FashionId
	weaponId := playInfo.AllWeaponInfo.Wear
	weaponState := int32(0)

	titleId := int32(0)
	wingId := playInfo.WingInfo.GetWingId()
	lingYuId := playInfo.LingyuInfo.LingyuId
	shenFaId := playInfo.ShenfaInfo.ShenfaId
	fabaoId := int32(0)
	petId := int32(0)
	xiantiId := int32(0)
	baGua := int32(0)
	flyPetId := int32(0)
	developLevel := int32(0)
	shenYuKey := int32(0)
	showObj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		fourGodKey,
		realm,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		fabaoId,
		petId,
		xiantiId,
		baGua,
		flyPetId,
		developLevel,
		shenYuKey,
	)

	vip := int32(0)
	level := playInfo.Level
	zhuansheng := int32(0)
	soulAwakenNum := int32(0)

	playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, level, zhuansheng, soulAwakenNum, false)
	p := createRobotPlayer(robottypes.RobotTypeArena, po, showObj, nil, skillList, playInfo.BattleProperty, 0, playerBattleObject, showServerId, power)
	return p
}

func (s *robotService) CreateModelRobot(
	serverId int32,
	sexType playertypes.SexType,
	roleType playertypes.RoleType,
	fashionId int32,
	weaponId int32,
	wingId int32,
	showServerId bool,
) scene.RobotPlayer {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	id, _ := idutil.GetId()
	userId, _ := idutil.GetId()
	name := ""

	po := playercommon.NewPlayerCommonObject(id, userId, serverId, name, roleType, sexType, false)

	mountId := int32(0)
	mountAdvanceId := int32(0)
	mountHidden := true
	fourGodKey := int32(0)
	realm := int32(0)
	spouse := ""
	spouseId := int64(0)
	weddingStatus := int32(0)
	ringType := int32(0)
	ringLevel := int32(0)
	skillList := make([]skillcommon.SkillObject, 0, 16)

	//展示
	weaponState := int32(0)
	titleId := int32(0)
	lingYuId := int32(0)
	shenFaId := int32(0)
	fabaoId := int32(0)
	petId := int32(0)
	xiantiId := int32(0)
	baGua := int32(0)
	flyPetId := int32(0)
	developLevel := int32(0)
	shenYuKey := int32(0)
	showObj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		fourGodKey,
		realm,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		fabaoId,
		petId,
		xiantiId,
		baGua,
		flyPetId,
		developLevel,
		shenYuKey,
	)

	vip := int32(0)
	level := int32(1)
	zhuansheng := int32(0)
	soulAwakenNum := int32(0)
	playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, level, zhuansheng, soulAwakenNum, false)
	battleProperties := map[int32]int64{}

	p := createRobotPlayerBase(robottypes.RobotTypeModel, po, showObj, nil, skillList, battleProperties, 0, playerBattleObject, showServerId, 0, 0, 0, scenetypes.FactionTypeModel)
	return p
}

func (s *robotService) CreateArenaRobot(serverId int32, properties map[propertytypes.BattlePropertyType]int64, reliveTime int32, power int64) scene.RobotPlayer {

	s.rwm.Lock()
	defer s.rwm.Unlock()
	robotTemplate := robottemplate.GetRobotTemplateService().GetRobotTemplate(power)
	if robotTemplate == nil {
		panic(fmt.Errorf("robot:一定可以获得机器人模板"))
	}
	id, _ := idutil.GetId()
	userId, _ := idutil.GetId()
	sex := robotTemplate.RandomSex()
	name := dummytemplate.GetDummyTemplateService().GetRandomDummyNameBySex(sex)
	roleType := robotTemplate.RandomRole()

	po := playercommon.NewPlayerCommonObject(id, userId, serverId, name, roleType, sex, false)

	//随机时装
	fashionId := robotTemplate.RandomFashion()
	//随机冰魂
	weaponId := robotTemplate.RandomWeapon()
	weaponState := int32(0)
	titleId := robotTemplate.RandomTitle()
	wingId := robotTemplate.RandomWing()
	mountId := robotTemplate.RandomMount()
	mountHidden := true
	mountAdvanceId := int32(0)
	lingYuId := robotTemplate.RandomField()
	shenFaId := robotTemplate.RandomShenfa()
	faBaoId := robotTemplate.RandomFabao()
	xianTiId := robotTemplate.RandomXianti()
	fourGodKey := int32(0)
	realm := int32(0)
	spouse := ""
	spouseId := int64(0)
	weddingStatus := int32(0)
	ringType := int32(0)
	ringLevel := int32(0)
	petId := int32(0)
	baGua := int32(0)
	flyPetId := int32(0)
	jueXueId := robotTemplate.RandomJueXue()
	diHunId := robotTemplate.RandomSoul()
	developLevel := int32(0)
	shenYuKey := int32(0)
	showObj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		fourGodKey,
		realm,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		faBaoId,
		petId,
		xianTiId,
		baGua,
		flyPetId,
		developLevel,
		shenYuKey,
	)
	roleLevel := robotTemplate.RandomLevel()

	skillList := make([]skillcommon.SkillObject, 0, 16)
	skillTemplateMap := skilltemplate.GetSkillTemplateService().GetAllSkillTemplates()
	for _, skillTemplate := range skillTemplateMap {
		switch skillTemplate.GetSkillFirstType() {
		case skilltypes.SkillFirstTypeNormal:
			if skillTemplate.GetRoleType() != roleType {
				continue
			}
			skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
			break
		case skilltypes.SkillFirstTypeRole:
			{
				//随机职业技能等级
				if skillTemplate.GetRoleType() != roleType {
					continue
				}
				maxLevel := skilltemplate.GetSkillTemplateService().GetMaxLevel(skillTemplate.TypeId)
				if roleLevel >= maxLevel {
					roleLevel = maxLevel
				}
				tempSkillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(skillTemplate.TypeId, roleLevel)
				if tempSkillTemplate == nil {
					continue
				}
				skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, roleLevel, nil))
				break
			}
		}
	}
	if jueXueId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(jueXueId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}
	if diHunId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(diHunId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}

	battleProperties := make(map[int32]int64)
	for k, v := range properties {
		battleProperties[int32(k)] = v
	}

	vip := int32(0)

	zhuansheng := int32(0)
	soulAwakenNum := int32(0)
	playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, roleLevel, zhuansheng, soulAwakenNum, false)
	p := createRobotPlayer(robottypes.RobotTypeArena, po, showObj, nil, skillList, battleProperties, reliveTime, playerBattleObject, false, power)
	s.robotMap[p.GetId()] = p

	//更新灵童
	lingTongId := robotTemplate.RandomLingTong()
	if lingTongId == 0 {
		return p
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	lingTongUid, _ := idutil.GetId()
	lingTongName := lingTongTemplate.Name
	pos := coretypes.Position{}
	angle := float64(0.0)
	//随机时装
	lingTongFashionId := robotTemplate.RandomLingTongFashion()
	//随机冰魂
	lingTongWeaponId := robotTemplate.RandomLingTongWeapon()
	lingTongWeaponState := int32(0)

	lingTongWingId := robotTemplate.RandomLingTongWing()
	lingTongMountId := robotTemplate.RandomLingTongMount()
	lingTongMountHidden := true
	lingTongLingYuId := robotTemplate.RandomLingTongLingyu()
	lingTongShenFaId := robotTemplate.RandomLingTongShenfa()
	lingTongFaBaoId := robotTemplate.RandomLingTongFabao()
	lingTongXianTiId := robotTemplate.RandomLingTongXianti()

	lingTongShowObj := lingtong.CreateLingTongShowObject(
		lingTongFashionId,
		lingTongWeaponId,
		lingTongWeaponState,
		0,
		lingTongWingId,
		lingTongMountId,
		lingTongMountHidden,
		lingTongShenFaId,
		lingTongLingYuId,
		lingTongFaBaoId,
		lingTongXianTiId,
	)
	lingTongBattleProperties := make(map[int32]int64)
	for k, v := range lingTongTemplate.GetLingTongBattlePropertyMap() {
		lingTongBattleProperties[int32(k)] = v
	}
	lingTong := lingtong.CreateLingTong(p, lingTongUid, lingTongName, pos, angle, lingTongTemplate, lingTongShowObj, lingTongBattleProperties)
	p.UpdateLingTong(lingTong)
	return p
}

func (s *robotService) CreateArenapvpRobot(platform int32, serverId int32, properties map[propertytypes.BattlePropertyType]int64, reliveTime int32, power int64) scene.RobotPlayer {

	s.rwm.Lock()
	defer s.rwm.Unlock()
	robotTemplate := robottemplate.GetRobotTemplateService().GetRobotTemplate(power)
	if robotTemplate == nil {
		panic(fmt.Errorf("robot:一定可以获得机器人模板"))
	}
	id, _ := idutil.GetId()
	userId, _ := idutil.GetId()
	sex := robotTemplate.RandomSex()
	name := dummytemplate.GetDummyTemplateService().GetRandomDummyNameBySex(sex)
	roleType := robotTemplate.RandomRole()

	po := playercommon.NewBasicPlayerCommonObject(id, userId, serverId, name, roleType, sex, false, platform)

	//随机时装
	fashionId := robotTemplate.RandomFashion()
	//随机冰魂
	weaponId := robotTemplate.RandomWeapon()
	weaponState := int32(0)
	titleId := robotTemplate.RandomTitle()
	wingId := robotTemplate.RandomWing()
	mountId := robotTemplate.RandomMount()
	mountHidden := true
	mountAdvanceId := int32(0)
	lingYuId := robotTemplate.RandomField()
	shenFaId := robotTemplate.RandomShenfa()
	faBaoId := robotTemplate.RandomFabao()
	xianTiId := robotTemplate.RandomXianti()
	fourGodKey := int32(0)
	realm := int32(0)
	spouse := ""
	spouseId := int64(0)
	weddingStatus := int32(0)
	ringType := int32(0)
	ringLevel := int32(0)
	petId := int32(0)
	baGua := int32(0)
	flyPetId := int32(0)
	jueXueId := robotTemplate.RandomJueXue()
	diHunId := robotTemplate.RandomSoul()
	developLevel := int32(0)
	shenYuKey := int32(0)
	showObj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		fourGodKey,
		realm,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		faBaoId,
		petId,
		xianTiId,
		baGua,
		flyPetId,
		developLevel,
		shenYuKey,
	)
	roleLevel := robotTemplate.RandomLevel()

	skillList := make([]skillcommon.SkillObject, 0, 16)
	skillTemplateMap := skilltemplate.GetSkillTemplateService().GetAllSkillTemplates()
	for _, skillTemplate := range skillTemplateMap {
		switch skillTemplate.GetSkillFirstType() {
		case skilltypes.SkillFirstTypeNormal:
			if skillTemplate.GetRoleType() != roleType {
				continue
			}
			skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
			break
		case skilltypes.SkillFirstTypeRole:
			{
				//随机职业技能等级
				if skillTemplate.GetRoleType() != roleType {
					continue
				}
				maxLevel := skilltemplate.GetSkillTemplateService().GetMaxLevel(skillTemplate.TypeId)
				if roleLevel >= maxLevel {
					roleLevel = maxLevel
				}
				tempSkillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(skillTemplate.TypeId, roleLevel)
				if tempSkillTemplate == nil {
					continue
				}
				skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, roleLevel, nil))
				break
			}
		}
	}
	if jueXueId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(jueXueId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}
	if diHunId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(diHunId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}

	battleProperties := make(map[int32]int64)
	for k, v := range properties {
		battleProperties[int32(k)] = v
	}

	vip := int32(0)

	zhuansheng := int32(0)
	soulAwakenNum := int32(0)
	playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, roleLevel, zhuansheng, soulAwakenNum, false)
	p := createRobotPlayer(robottypes.RobotTypeArenapvp, po, showObj, nil, skillList, battleProperties, reliveTime, playerBattleObject, true, power)
	s.robotMap[p.GetId()] = p

	//更新灵童
	lingTongId := robotTemplate.RandomLingTong()
	if lingTongId == 0 {
		return p
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	lingTongUid, _ := idutil.GetId()
	lingTongName := lingTongTemplate.Name
	pos := coretypes.Position{}
	angle := float64(0.0)
	//随机时装
	lingTongFashionId := robotTemplate.RandomLingTongFashion()
	//随机冰魂
	lingTongWeaponId := robotTemplate.RandomLingTongWeapon()
	lingTongWeaponState := int32(0)

	lingTongWingId := robotTemplate.RandomLingTongWing()
	lingTongMountId := robotTemplate.RandomLingTongMount()
	lingTongMountHidden := true
	lingTongLingYuId := robotTemplate.RandomLingTongLingyu()
	lingTongShenFaId := robotTemplate.RandomLingTongShenfa()
	lingTongFaBaoId := robotTemplate.RandomLingTongFabao()
	lingTongXianTiId := robotTemplate.RandomLingTongXianti()

	lingTongShowObj := lingtong.CreateLingTongShowObject(
		lingTongFashionId,
		lingTongWeaponId,
		lingTongWeaponState,
		0,
		lingTongWingId,
		lingTongMountId,
		lingTongMountHidden,
		lingTongShenFaId,
		lingTongLingYuId,
		lingTongFaBaoId,
		lingTongXianTiId,
	)
	lingTongBattleProperties := make(map[int32]int64)
	for k, v := range lingTongTemplate.GetLingTongBattlePropertyMap() {
		lingTongBattleProperties[int32(k)] = v
	}
	lingTong := lingtong.CreateLingTong(p, lingTongUid, lingTongName, pos, angle, lingTongTemplate, lingTongShowObj, lingTongBattleProperties)
	p.UpdateLingTong(lingTong)
	return p
}

func (s *robotService) CreateTeamCopyRobot(
	serverId int32,
	properties map[propertytypes.BattlePropertyType]int64,
	reliveTime int32,
	power int64) scene.RobotPlayer {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	robotTemplate := robottemplate.GetRobotTemplateService().GetRobotTemplate(power)
	if robotTemplate == nil {
		panic(fmt.Errorf("robot:一定可以获得机器人模板"))
	}
	id, _ := idutil.GetId()
	userId, _ := idutil.GetId()
	sex := robotTemplate.RandomSex()
	name := dummytemplate.GetDummyTemplateService().GetRandomDummyNameBySex(sex)
	roleType := robotTemplate.RandomRole()

	po := playercommon.NewPlayerCommonObject(id, userId, serverId, name, roleType, sex, false)

	//随机时装
	fashionId := robotTemplate.RandomFashion()
	//随机冰魂
	weaponId := robotTemplate.RandomWeapon()
	weaponState := int32(0)
	titleId := robotTemplate.RandomTitle()
	wingId := robotTemplate.RandomWing()
	mountId := robotTemplate.RandomMount()
	mountHidden := true
	mountAdvanceId := int32(0)
	lingYuId := robotTemplate.RandomField()
	shenFaId := robotTemplate.RandomShenfa()
	faBaoId := robotTemplate.RandomFabao()
	xianTiId := robotTemplate.RandomXianti()
	fourGodKey := int32(0)
	realm := int32(0)
	spouse := ""
	spouseId := int64(0)
	weddingStatus := int32(0)
	ringType := int32(0)
	ringLevel := int32(0)
	petId := int32(0)
	baGua := int32(0)
	flyPetId := int32(0)
	jueXueId := robotTemplate.RandomJueXue()
	diHunId := robotTemplate.RandomSoul()
	developLevel := int32(0)
	shenYuKey := int32(0)
	showObj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		fourGodKey,
		realm,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		faBaoId,
		petId,
		xianTiId,
		baGua,
		flyPetId,
		developLevel,
		shenYuKey,
	)
	roleLevel := robotTemplate.RandomLevel()

	skillList := make([]skillcommon.SkillObject, 0, 16)
	skillTemplateMap := skilltemplate.GetSkillTemplateService().GetAllSkillTemplates()
	for _, skillTemplate := range skillTemplateMap {
		switch skillTemplate.GetSkillFirstType() {
		case skilltypes.SkillFirstTypeNormal:
			if skillTemplate.GetRoleType() != roleType {
				continue
			}
			skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
			break
		case skilltypes.SkillFirstTypeRole:
			{
				//随机职业技能等级
				if skillTemplate.GetRoleType() != roleType {
					continue
				}
				maxLevel := skilltemplate.GetSkillTemplateService().GetMaxLevel(skillTemplate.TypeId)
				if roleLevel >= maxLevel {
					roleLevel = maxLevel
				}
				tempSkillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(skillTemplate.TypeId, roleLevel)
				if tempSkillTemplate == nil {
					continue
				}
				skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, roleLevel, nil))
				break
			}
		}
	}
	if jueXueId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(jueXueId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}
	if diHunId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(diHunId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}

	battleProperties := make(map[int32]int64)
	for k, v := range properties {
		battleProperties[int32(k)] = v
	}

	vip := int32(0)

	zhuansheng := int32(0)
	soulAwakenNum := int32(0)
	playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, roleLevel, zhuansheng, soulAwakenNum, false)
	p := createRobotPlayer(robottypes.RobotTypeTeamCopy, po, showObj, nil, skillList, battleProperties, reliveTime, playerBattleObject, false, power)
	s.robotMap[p.GetId()] = p

	//更新灵童
	lingTongId := robotTemplate.RandomLingTong()
	if lingTongId == 0 {
		return p
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	lingTongUid, _ := idutil.GetId()
	lingTongName := lingTongTemplate.Name
	pos := coretypes.Position{}
	angle := float64(0.0)
	//随机时装
	lingTongFashionId := robotTemplate.RandomLingTongFashion()
	//随机冰魂
	lingTongWeaponId := robotTemplate.RandomLingTongWeapon()
	lingTongWeaponState := int32(0)

	lingTongWingId := robotTemplate.RandomLingTongWing()
	lingTongMountId := robotTemplate.RandomLingTongMount()
	lingTongMountHidden := true
	lingTongLingYuId := robotTemplate.RandomLingTongLingyu()
	lingTongShenFaId := robotTemplate.RandomLingTongShenfa()
	lingTongFaBaoId := robotTemplate.RandomLingTongFabao()
	lingTongXianTiId := robotTemplate.RandomLingTongXianti()

	lingTongShowObj := lingtong.CreateLingTongShowObject(
		lingTongFashionId,
		lingTongWeaponId,
		lingTongWeaponState,
		0,
		lingTongWingId,
		lingTongMountId,
		lingTongMountHidden,
		lingTongShenFaId,
		lingTongLingYuId,
		lingTongFaBaoId,
		lingTongXianTiId,
	)
	lingTongBattleProperties := make(map[int32]int64)
	for k, v := range lingTongTemplate.GetLingTongBattlePropertyMap() {
		lingTongBattleProperties[int32(k)] = v
	}
	lingTong := lingtong.CreateLingTong(p, lingTongUid, lingTongName, pos, angle, lingTongTemplate, lingTongShowObj, lingTongBattleProperties)
	p.UpdateLingTong(lingTong)
	return p
}

// func (s *robotService) CreateClientTestRobot(copyPlayer scene.Player, showServerId bool)  scene.RobotPlayer {
// 	s.rwm.Lock()
// 	defer s.rwm.Unlock()
// 	id, _ := idutil.GetId()
// 	userId := int64(0)
// 	s.currentId += 1
// 	sex := copyPlayer.GetSex()
// 	name := dummytemplate.GetDummyTemplateService().GetRandomDummyNameBySex(sex)
// 	role := copyPlayer.GetRole()

// 	serverId := copyPlayer.GetServerId()
// 	po := playercommon.NewPlayerCommonObject(id, userId, serverId, name, role, sex, false)

// 	//随机时装
// 	randomFashionTemplate := fashion.GetFashionService().RandomFashionTemplate()
// 	fashionId := int32(randomFashionTemplate.TemplateId())

// 	weaponTemplate := weapon.GetWeaponService().RandomWeaponTemplate()
// 	weaponId := int32(weaponTemplate.TemplateId())
// 	// weaponId := int32(0)
// 	weaponState := int32(0)
// 	if weaponTemplate.IsAwaken > 0 {
// 		weaponState = int32(1)
// 	}

// 	titleTemplate := title.GetTitleService().RandomTitleTemplate()
// 	titleId := int32(titleTemplate.TemplateId())

// 	wingTemplate := wing.GetWingService().RandomWingTemplate()
// 	wingId := int32(wingTemplate.TemplateId())
// 	// wingId := int32(0)
// 	// mountTemplate := mount.GetMountService().RandomMountTemplate()
// 	// mountId := int32(mountTemplate.TemplateId())
// 	mountId := int32(0)
// 	mountHidden := true
// 	lingYuTemplate := lingyutemplate.GetLingyuTemplateService().RandomLingYuTemplate()
// 	lingYuId := int32(lingYuTemplate.TemplateId())
// 	// lingYuId := int32(0)
// 	shenFaTemplate := shenfatemplate.GetShenfaTemplateService().RandomShenFaTemplate()
// 	shenFaId := int32(shenFaTemplate.TemplateId())
// 	// shenFaId := int32(0)
// 	fourGodKey := int32(0)
// 	realm := int32(0)
// 	spouse := ""
// 	spouseId := int64(0)
// 	weddingStatus := int32(0)
// 	ringType := int32(0)
// 	ringLevel := int32(0)
// 	petId := int32(0)
// 	faBaoId := int32(0)
// 	xianTiId := int32(0)
// 	baGua := int32(0)
// 	flyPetId := int32(0)
// 	showObj := battle.CreatePlayerShowObject(
// 		fashionId,
// 		weaponId,
// 		weaponState,
// 		titleId,
// 		wingId,
// 		mountId,
// 		mountHidden,
// 		shenFaId,
// 		lingYuId,
// 		fourGodKey,
// 		realm,
// 		spouse,
// 		spouseId,
// 		weddingStatus,
// 		ringType,
// 		ringLevel,
// 		faBaoId,
// 		petId,
// 		xianTiId,
// 		baGua,
// 		flyPetId,
// 	)
// 	skillList := make([]skillcommon.SkillObject, 0, 16)
// 	allSkills := skilltemplate.GetSkillTemplateService().GetAllSkillTemplates()
// 	for _, skillTemplate := range allSkills {
// 		if skillTemplate.GetSkillFirstType() == skilltypes.SkillFirstTypeMonsterPassive {
// 			continue
// 		}
// 		if skillTemplate.IsLimitRole() && skillTemplate.GetRoleType() != role {
// 			continue
// 		}

// 		sk := skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev)
// 		skillList = append(skillList, sk)
// 	}

// 	battleProperties := make(map[int32]int64)
// 	for pt := propertytypes.MinBattlePropertyType; pt <= propertytypes.MaxBattlePropertyType; pt++ {
// 		val := copyPlayer.GetBattleProperty(pt)
// 		battleProperties[int32(pt)] = val
// 	}

// 	vip := copyPlayer.GetVip()
// 	level := copyPlayer.GetLevel()
// 	zhuansheng := copyPlayer.GetZhuanSheng()
// 	soulAwakenNum := copyPlayer.GetSoulAwakenNum()
// 	power := int64(0)
// 	playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, level, zhuansheng, soulAwakenNum)
// 	p := createRobotPlayer(robottypes.RobotTypeTest, po, showObj, nil, skillList, battleProperties, 0, playerBattleObject, showServerId, power)
// 	s.robotMap[p.GetId()] = p
// 	return p
// }

func (s *robotService) GMClear() {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	for _, p := range s.robotMap {
		sc := p.GetScene()
		if sc != nil {
			ctx := scene.WithScene(context.Background(), sc)
			//TODO 造成死锁
			sc.Post(message.NewScheduleMessage(onRobotExit, ctx, p, nil))
		}
		delete(s.robotMap, p.GetId())
	}
}

func (s *robotService) GetRobot(id int64) scene.RobotPlayer {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	robotPlayer, ok := s.robotMap[id]
	if !ok {
		return nil
	}
	return robotPlayer
}

func onRobotExit(ctx context.Context, result interface{}, err error) error {
	s := scene.SceneInContext(ctx)
	p := result.(scene.Player)
	s.RemoveSceneObject(p, true)
	return nil
}

func (s *robotService) RemoveRobot(id int64) {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	delete(s.robotMap, id)
}

func (s *robotService) CreateQuestRobot(beginQuestId int32, endQuestId int32, properties map[propertytypes.BattlePropertyType]int64, power int64, showServerId bool) scene.RobotPlayer {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	robotTemplate := robottemplate.GetRobotTemplateService().GetRobotTemplate(power)
	if robotTemplate == nil {
		return nil
	}
	id, _ := idutil.GetId()
	userId := int64(0)
	sex := robotTemplate.RandomSex()
	name := dummytemplate.GetDummyTemplateService().GetRandomDummyNameBySex(sex)
	roleType := robotTemplate.RandomRole()

	serverId := global.GetGame().GetServerIndex()
	po := playercommon.NewPlayerCommonObject(id, userId, serverId, name, roleType, sex, false)

	//随机时装
	fashionId := robotTemplate.RandomFashion()
	//随机冰魂
	weaponId := robotTemplate.RandomWeapon()
	weaponState := int32(0)
	titleId := robotTemplate.RandomTitle()
	wingId := robotTemplate.RandomWing()
	mountId := robotTemplate.RandomMount()
	mountAdvanceId := int32(0)
	mountHidden := true
	lingYuId := robotTemplate.RandomField()
	shenFaId := robotTemplate.RandomShenfa()
	faBaoId := robotTemplate.RandomFabao()
	xianTiId := robotTemplate.RandomXianti()
	fourGodKey := int32(0)
	realm := int32(0)
	spouse := ""
	spouseId := int64(0)
	weddingStatus := int32(0)
	ringType := int32(0)
	ringLevel := int32(0)
	petId := int32(0)
	baGua := int32(0)
	flyPetId := int32(0)
	jueXueId := robotTemplate.RandomJueXue()
	diHunId := robotTemplate.RandomSoul()
	developLevel := int32(0)
	shenYuKey := int32(0)
	showObj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		fourGodKey,
		realm,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		faBaoId,
		petId,
		xianTiId,
		baGua,
		flyPetId,
		developLevel,
		shenYuKey,
	)
	roleLevel := robotTemplate.RandomLevel()

	skillList := make([]skillcommon.SkillObject, 0, 16)
	skillTemplateMap := skilltemplate.GetSkillTemplateService().GetAllSkillTemplates()
	for _, skillTemplate := range skillTemplateMap {
		switch skillTemplate.GetSkillFirstType() {
		case skilltypes.SkillFirstTypeNormal:
			if skillTemplate.GetRoleType() != roleType {
				continue
			}
			skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
			break
		case skilltypes.SkillFirstTypeRole:
			{
				//随机职业技能等级
				if skillTemplate.GetRoleType() != roleType {
					continue
				}
				maxLevel := skilltemplate.GetSkillTemplateService().GetMaxLevel(skillTemplate.TypeId)
				if roleLevel >= maxLevel {
					roleLevel = maxLevel
				}
				tempSkillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(skillTemplate.TypeId, roleLevel)
				if tempSkillTemplate == nil {
					continue
				}
				skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, roleLevel, nil))
				break
			}
		}
	}
	if jueXueId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(jueXueId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}
	if diHunId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(diHunId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}

	battleProperties := make(map[int32]int64)
	for k, v := range properties {
		battleProperties[int32(k)] = v
	}

	vip := int32(0)

	zhuansheng := int32(0)
	soulAwakenNum := int32(0)
	playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, roleLevel, zhuansheng, soulAwakenNum, false)
	p := createRobotPlayerWithQuest(robottypes.RobotTypeQuest, po, showObj, nil, skillList, battleProperties, 0, playerBattleObject, showServerId, beginQuestId, endQuestId, power)
	s.robotMap[p.GetId()] = p

	//更新灵童
	lingTongId := robotTemplate.RandomLingTong()
	if lingTongId == 0 {
		return p
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	lingTongUid, _ := idutil.GetId()
	lingTongName := lingTongTemplate.Name
	pos := coretypes.Position{}
	angle := float64(0.0)
	//随机时装
	lingTongFashionId := robotTemplate.RandomLingTongFashion()
	//随机冰魂
	lingTongWeaponId := robotTemplate.RandomLingTongWeapon()
	lingTongWeaponState := int32(0)

	lingTongWingId := robotTemplate.RandomLingTongWing()
	lingTongMountId := robotTemplate.RandomLingTongMount()
	lingTongMountHidden := true
	lingTongLingYuId := robotTemplate.RandomLingTongLingyu()
	lingTongShenFaId := robotTemplate.RandomLingTongShenfa()
	lingTongFaBaoId := robotTemplate.RandomLingTongFabao()
	lingTongXianTiId := robotTemplate.RandomLingTongXianti()

	lingTongShowObj := lingtong.CreateLingTongShowObject(
		lingTongFashionId,
		lingTongWeaponId,
		lingTongWeaponState,
		0,
		lingTongWingId,
		lingTongMountId,
		lingTongMountHidden,
		lingTongShenFaId,
		lingTongLingYuId,
		lingTongFaBaoId,
		lingTongXianTiId,
	)
	lingTongBattleProperties := make(map[int32]int64)
	for k, v := range lingTongTemplate.GetLingTongBattlePropertyMap() {
		lingTongBattleProperties[int32(k)] = v
	}
	lingTong := lingtong.CreateLingTong(p, lingTongUid, lingTongName, pos, angle, lingTongTemplate, lingTongShowObj, lingTongBattleProperties)
	p.UpdateLingTong(lingTong)
	return p

}

func (s *robotService) CreateClientTestRobot(properties map[propertytypes.BattlePropertyType]int64, power int64, showServerId bool) scene.RobotPlayer {
	s.rwm.Lock()
	defer s.rwm.Unlock()
	robotTemplate := robottemplate.GetRobotTemplateService().GetRobotTemplate(power)
	if robotTemplate == nil {
		return nil
	}
	id, _ := idutil.GetId()
	userId := int64(0)
	sex := robotTemplate.RandomSex()
	name := dummytemplate.GetDummyTemplateService().GetRandomDummyNameBySex(sex)
	roleType := robotTemplate.RandomRole()

	serverId := global.GetGame().GetServerIndex()
	po := playercommon.NewPlayerCommonObject(id, userId, serverId, name, roleType, sex, false)

	//随机时装
	fashionId := robotTemplate.RandomFashion()
	//随机冰魂
	weaponId := robotTemplate.RandomWeapon()
	weaponState := int32(0)
	titleId := robotTemplate.RandomTitle()
	wingId := robotTemplate.RandomWing()
	mountId := robotTemplate.RandomMount()
	mountAdvanceId := int32(0)
	mountHidden := true
	lingYuId := robotTemplate.RandomField()
	shenFaId := robotTemplate.RandomShenfa()
	faBaoId := robotTemplate.RandomFabao()
	xianTiId := robotTemplate.RandomXianti()
	fourGodKey := int32(0)
	realm := int32(0)
	spouse := ""
	spouseId := int64(0)
	weddingStatus := int32(0)
	ringType := int32(0)
	ringLevel := int32(0)
	petId := int32(0)
	baGua := int32(0)
	flyPetId := int32(0)
	jueXueId := robotTemplate.RandomJueXue()
	diHunId := robotTemplate.RandomSoul()
	developLevel := int32(0)
	shenYuKey := int32(0)
	showObj := battle.CreatePlayerShowObject(
		fashionId,
		weaponId,
		weaponState,
		titleId,
		wingId,
		mountId,
		mountAdvanceId,
		mountHidden,
		shenFaId,
		lingYuId,
		fourGodKey,
		realm,
		spouse,
		spouseId,
		weddingStatus,
		ringType,
		ringLevel,
		faBaoId,
		petId,
		xianTiId,
		baGua,
		flyPetId,
		developLevel,
		shenYuKey,
	)
	roleLevel := robotTemplate.RandomLevel()

	skillList := make([]skillcommon.SkillObject, 0, 16)
	skillTemplateMap := skilltemplate.GetSkillTemplateService().GetAllSkillTemplates()
	for _, skillTemplate := range skillTemplateMap {
		switch skillTemplate.GetSkillFirstType() {
		case skilltypes.SkillFirstTypeNormal:
			if skillTemplate.GetRoleType() != roleType {
				continue
			}
			skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
			break
		case skilltypes.SkillFirstTypeRole:
			{
				//随机职业技能等级
				if skillTemplate.GetRoleType() != roleType {
					continue
				}
				maxLevel := skilltemplate.GetSkillTemplateService().GetMaxLevel(skillTemplate.TypeId)
				if roleLevel >= maxLevel {
					roleLevel = maxLevel
				}
				tempSkillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplateByTypeAndLevel(skillTemplate.TypeId, roleLevel)
				if tempSkillTemplate == nil {
					continue
				}
				skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, roleLevel, nil))
				break
			}
		}
	}
	if jueXueId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(jueXueId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}
	if diHunId != 0 {
		skillTemplate := skilltemplate.GetSkillTemplateService().GetSkillTemplate(diHunId)
		skillList = append(skillList, skillcommon.CreateSkillObject(skillTemplate.TypeId, skillTemplate.Lev, nil))
	}

	battleProperties := make(map[int32]int64)
	for k, v := range properties {
		battleProperties[int32(k)] = v
	}

	vip := int32(0)

	zhuansheng := int32(0)
	soulAwakenNum := int32(0)
	playerBattleObject := battlecommon.CreatePlayerBattleObject(vip, roleLevel, zhuansheng, soulAwakenNum, false)
	p := createRobotPlayer(robottypes.RobotTypeTest, po, showObj, nil, skillList, battleProperties, 0, playerBattleObject, showServerId, power)
	s.robotMap[p.GetId()] = p

	//更新灵童
	lingTongId := robotTemplate.RandomLingTong()
	if lingTongId == 0 {
		return p
	}
	lingTongTemplate := lingtongtemplate.GetLingTongTemplateService().GetLingTongTemplate(lingTongId)
	lingTongUid, _ := idutil.GetId()
	lingTongName := dummytemplate.GetDummyTemplateService().GetRandomDummyName()
	pos := coretypes.Position{}
	angle := float64(0.0)
	//随机时装
	lingTongFashionId := robotTemplate.RandomLingTongFashion()
	//随机冰魂
	lingTongWeaponId := robotTemplate.RandomLingTongWeapon()
	lingTongWeaponState := int32(0)

	lingTongWingId := robotTemplate.RandomLingTongWing()
	lingTongMountId := robotTemplate.RandomLingTongMount()
	lingTongMountHidden := true
	lingTongLingYuId := robotTemplate.RandomLingTongLingyu()
	lingTongShenFaId := robotTemplate.RandomLingTongShenfa()
	lingTongFaBaoId := robotTemplate.RandomLingTongFabao()
	lingTongXianTiId := robotTemplate.RandomLingTongXianti()

	lingTongShowObj := lingtong.CreateLingTongShowObject(
		lingTongFashionId,
		lingTongWeaponId,
		lingTongWeaponState,
		0,
		lingTongWingId,
		lingTongMountId,
		lingTongMountHidden,
		lingTongShenFaId,
		lingTongLingYuId,
		lingTongFaBaoId,
		lingTongXianTiId,
	)
	lingTongBattleProperties := make(map[int32]int64)
	for k, v := range lingTongTemplate.GetLingTongBattlePropertyMap() {
		lingTongBattleProperties[int32(k)] = v
	}
	lingTong := lingtong.CreateLingTong(p, lingTongUid, lingTongName, pos, angle, lingTongTemplate, lingTongShowObj, lingTongBattleProperties)
	p.UpdateLingTong(lingTong)
	return p
}

var (
	once sync.Once

	cs *robotService
)

func Init() (err error) {
	once.Do(func() {

		cs = &robotService{}
		err = cs.init()
	})
	return err
}

func GetRobotService() RobotService {
	if cs == nil {
		return nil
	}

	return cs
}
