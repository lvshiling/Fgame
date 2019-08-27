package template

import (
	accounttypes "fgame/fgame/account/login/types"
	"fgame/fgame/core/template"
	scenetypes "fgame/fgame/game/scene/types"
	gametemplate "fgame/fgame/game/template"
	"math/rand"
	"sort"
	"sync"
)

type SceneTemplateService interface {
	GetMap(id int32) *gametemplate.MapTemplate
	GetAllWorld() map[int32]*gametemplate.MapTemplate
	GetAllBoss() map[int32]*gametemplate.MapTemplate

	GetCdGroup(id int32) *gametemplate.CdGroupTemplate
	GetQuestNPC(npcId int32) *gametemplate.SceneTemplate
	GetPortal(portalId int32) *gametemplate.PortalTemplate
	RandomWorldMap() *gametemplate.MapTemplate
	GetYaZhiTemplate(level int32) *gametemplate.YaZhiTemplate
	GetNPC(npcId int32) *gametemplate.SceneTemplate
	GetPortalSceneTemplate(portalId int32) *gametemplate.SceneTemplate
	//获取场景的传送阵
	GetPortalTemplateMapByMapId(mapId int32) map[int32]*gametemplate.PortalTemplate
	//获取场景机器人
	GetRobotQuestTemplate(sdkList []accounttypes.SDKType, mapId int32) gametemplate.RobotQuestTemplateInterface
	//set_type是世界boss
	GetWorldBossTemplate() []*gametemplate.BiologyTemplate
	GetDingShiSceneTemplate(mapId int32, biologyId int32) *gametemplate.SceneTemplate
	GetSceneTemplate(mapId int32, biologyId int32) *gametemplate.SceneTemplate
}

type sceneTemplateService struct {
	worldMap  map[int32]*gametemplate.MapTemplate
	worldList []*gametemplate.MapTemplate
	bossMap   map[int32]*gametemplate.MapTemplate

	mapTemplateMap map[int32]*gametemplate.MapTemplate
	cdGroupMap     map[int32]*gametemplate.CdGroupTemplate
	//记录所有任务npc
	questNPCMap map[int32]*gametemplate.SceneTemplate
	//传送阵
	portalMap map[int32]*gametemplate.PortalTemplate
	//传送阵场景对象
	portalSceneTemplate map[int32]*gametemplate.SceneTemplate
	//记录所有类型
	npcTemplateMap map[int32]*gametemplate.SceneTemplate
	//场景的传送阵
	portalTemplateMapOfMap map[int32]map[int32]*gametemplate.PortalTemplate
	//压制模板
	yaZhiTemplateList []*gametemplate.YaZhiTemplate
	//任务机器人配置
	baseRobotQuestTemplateMap map[int32]*gametemplate.RobotQuestTemplate
	//渠道任务机器人
	sdkRobotQuestTemplateMapOfMap map[accounttypes.SDKType]map[int32]*gametemplate.RobotQuestQudaoTemplate
	//世界Boss列表
	wordBossList []*gametemplate.BiologyTemplate
	//定时boss
	dingShiBossMap map[int32]map[int32]*gametemplate.SceneTemplate
}

func (s *sceneTemplateService) init() (err error) {
	s.mapTemplateMap = make(map[int32]*gametemplate.MapTemplate)
	s.worldMap = make(map[int32]*gametemplate.MapTemplate)
	s.worldList = make([]*gametemplate.MapTemplate, 0, 8)
	s.bossMap = make(map[int32]*gametemplate.MapTemplate)
	s.dingShiBossMap = make(map[int32]map[int32]*gametemplate.SceneTemplate)

	templateMapTemplateMap := template.GetTemplateService().GetAll((*gametemplate.MapTemplate)(nil))

	for _, tempMapTemplate := range templateMapTemplateMap {
		mapTemplate := tempMapTemplate.(*gametemplate.MapTemplate)
		s.mapTemplateMap[int32(mapTemplate.TemplateId())] = mapTemplate
		if mapTemplate.IsWorld() {
			s.worldMap[int32(mapTemplate.TemplateId())] = mapTemplate
			s.worldList = append(s.worldList, mapTemplate)
		}
		if mapTemplate.IsBoss() {
			s.bossMap[int32(mapTemplate.TemplateId())] = mapTemplate
		}

	}
	//初始化cd组
	s.initCdGroup()
	//初始化npc
	s.initNPC()
	//初始化压制
	s.initYaZhi()
	//初始化任务机器人配置
	s.initQuestRobot()
	//初始化世界BOSS配置
	s.initWorldBoss()
	//初始化定时
	s.initDingShi()
	return
}

//初始化cd组
func (s *sceneTemplateService) initCdGroup() {
	s.cdGroupMap = make(map[int32]*gametemplate.CdGroupTemplate)
	templateCdGroupTemplateMap := template.GetTemplateService().GetAll((*gametemplate.CdGroupTemplate)(nil))
	for _, templateCdGroupTemplate := range templateCdGroupTemplateMap {
		cdGroupTemplate := templateCdGroupTemplate.(*gametemplate.CdGroupTemplate)
		s.cdGroupMap[int32(cdGroupTemplate.TemplateId())] = cdGroupTemplate
	}
}

//初始化npc组
func (s *sceneTemplateService) initNPC() {
	s.npcTemplateMap = make(map[int32]*gametemplate.SceneTemplate)
	s.questNPCMap = make(map[int32]*gametemplate.SceneTemplate)
	s.portalSceneTemplate = make(map[int32]*gametemplate.SceneTemplate)
	s.portalTemplateMapOfMap = make(map[int32]map[int32]*gametemplate.PortalTemplate)

	templateSceneTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SceneTemplate)(nil))
	for _, tempSceneTemplate := range templateSceneTemplateMap {
		sceneTemplate := tempSceneTemplate.(*gametemplate.SceneTemplate)
		biologyTemplate := sceneTemplate.GetBiology()
		if biologyTemplate.GetBiologyType() == scenetypes.BiologyTypeNPC {
			s.questNPCMap[int32(biologyTemplate.TemplateId())] = sceneTemplate
		}
		if biologyTemplate.GetBiologyType() == scenetypes.BiologyTransmissionArray {
			s.portalSceneTemplate[biologyTemplate.PortalId] = sceneTemplate
		}
		s.npcTemplateMap[int32(biologyTemplate.TemplateId())] = sceneTemplate
	}
	//传送阵
	s.portalMap = make(map[int32]*gametemplate.PortalTemplate)
	templatePortalTemplateMap := template.GetTemplateService().GetAll((*gametemplate.PortalTemplate)(nil))
	for _, tempPortalTemplate := range templatePortalTemplateMap {
		portalTemplate := tempPortalTemplate.(*gametemplate.PortalTemplate)
		s.portalMap[int32(portalTemplate.TemplateId())] = portalTemplate
	}

	for _, sceneTemplate := range s.portalSceneTemplate {
		tempPortalTemplateMap, ok := s.portalTemplateMapOfMap[sceneTemplate.SceneID]
		if !ok {
			tempPortalTemplateMap = make(map[int32]*gametemplate.PortalTemplate)
			s.portalTemplateMapOfMap[sceneTemplate.SceneID] = tempPortalTemplateMap
		}
		tempPortalTemplateMap[sceneTemplate.GetBiology().PortalId] = s.portalMap[sceneTemplate.GetBiology().PortalId]
	}

}

func (s *sceneTemplateService) initYaZhi() {
	tempYaZhiTemplateMap := template.GetTemplateService().GetAll((*gametemplate.YaZhiTemplate)(nil))
	for _, tempYaZhiTemplate := range tempYaZhiTemplateMap {
		yaZhiTemplate := tempYaZhiTemplate.(*gametemplate.YaZhiTemplate)
		s.yaZhiTemplateList = append(s.yaZhiTemplateList, yaZhiTemplate)
	}
	sort.Sort(yazhiTemplateList(s.yaZhiTemplateList))
}

type yazhiTemplateList []*gametemplate.YaZhiTemplate

func (l yazhiTemplateList) Len() int {
	return len(l)
}

func (l yazhiTemplateList) Less(i, j int) bool {
	return l[i].LevelMin < l[j].LevelMin
}

func (l yazhiTemplateList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

//初始化任务机器人
func (s *sceneTemplateService) initQuestRobot() {
	//初始化机器人配置
	s.baseRobotQuestTemplateMap = make(map[int32]*gametemplate.RobotQuestTemplate)
	templateRobotQuestTemplateMap := template.GetTemplateService().GetAll((*gametemplate.RobotQuestTemplate)(nil))
	for _, tempRobotQuestTemplate := range templateRobotQuestTemplateMap {
		robotQuestTemplate := tempRobotQuestTemplate.(*gametemplate.RobotQuestTemplate)
		s.baseRobotQuestTemplateMap[robotQuestTemplate.MapId] = robotQuestTemplate
	}
	//初始化渠道机器人配置
	s.sdkRobotQuestTemplateMapOfMap = make(map[accounttypes.SDKType]map[int32]*gametemplate.RobotQuestQudaoTemplate)
	templateRobotQuestQudaoTemplateMap := template.GetTemplateService().GetAll((*gametemplate.RobotQuestQudaoTemplate)(nil))
	for _, tempRobotQuestQudaoTemplate := range templateRobotQuestQudaoTemplateMap {
		robotQuestQudaoTemplate := tempRobotQuestQudaoTemplate.(*gametemplate.RobotQuestQudaoTemplate)
		sdkRobotQuestTemplateMap, ok := s.sdkRobotQuestTemplateMapOfMap[robotQuestQudaoTemplate.GetSDKType()]
		if !ok {
			sdkRobotQuestTemplateMap = make(map[int32]*gametemplate.RobotQuestQudaoTemplate)
			s.sdkRobotQuestTemplateMapOfMap[robotQuestQudaoTemplate.GetSDKType()] = sdkRobotQuestTemplateMap
		}
		sdkRobotQuestTemplateMap[robotQuestQudaoTemplate.MapId] = robotQuestQudaoTemplate
	}
}

//初始化世界BOSS
func (s *sceneTemplateService) initWorldBoss() {
	biologyTemplateMap := template.GetTemplateService().GetAll((*gametemplate.BiologyTemplate)(nil))
	for _, tempObj := range biologyTemplateMap {
		biologyTemp := tempObj.(*gametemplate.BiologyTemplate)
		if biologyTemp.GetBiologySetType() != scenetypes.BiologySetTypeWorldBoss {
			continue
		}
		s.wordBossList = append(s.wordBossList, biologyTemp)
	}
}

//初始化定时刷新
func (s *sceneTemplateService) initDingShi() {
	sceneTemplateMap := template.GetTemplateService().GetAll((*gametemplate.SceneTemplate)(nil))
	for _, tempObj := range sceneTemplateMap {
		sceneTemp := tempObj.(*gametemplate.SceneTemplate)
		if sceneTemp.GetBiology().GetRebornType() != scenetypes.BiologyRebornTypeTime {
			continue
		}
		dingShiBossMap, ok := s.dingShiBossMap[sceneTemp.SceneID]
		if !ok {
			dingShiBossMap = make(map[int32]*gametemplate.SceneTemplate)
			s.dingShiBossMap[sceneTemp.SceneID] = dingShiBossMap
		}
		dingShiBossMap[int32(sceneTemp.GetBiology().TemplateId())] = sceneTemp
	}
}

func (s *sceneTemplateService) GetAllWorld() map[int32]*gametemplate.MapTemplate {
	return s.worldMap
}

func (s *sceneTemplateService) GetAllBoss() map[int32]*gametemplate.MapTemplate {
	return s.bossMap
}

func (s *sceneTemplateService) GetMap(id int32) *gametemplate.MapTemplate {
	mapTemplate, ok := s.mapTemplateMap[id]
	if !ok {
		return nil
	}
	return mapTemplate
}

func (s *sceneTemplateService) GetCdGroup(id int32) *gametemplate.CdGroupTemplate {
	cdGroupTemplate, ok := s.cdGroupMap[id]
	if !ok {
		return nil
	}
	return cdGroupTemplate
}

func (s *sceneTemplateService) GetQuestNPC(id int32) *gametemplate.SceneTemplate {
	sceneTemplate, ok := s.questNPCMap[id]
	if !ok {
		return nil
	}
	return sceneTemplate
}

func (s *sceneTemplateService) GetNPC(id int32) *gametemplate.SceneTemplate {
	sceneTemplate, ok := s.npcTemplateMap[id]
	if !ok {
		return nil
	}
	return sceneTemplate
}

func (s *sceneTemplateService) GetPortal(portalId int32) *gametemplate.PortalTemplate {
	portalTemplate, ok := s.portalMap[portalId]
	if !ok {
		return nil
	}
	return portalTemplate
}

func (s *sceneTemplateService) GetPortalSceneTemplate(portalId int32) *gametemplate.SceneTemplate {
	portalTemplate, ok := s.portalSceneTemplate[portalId]
	if !ok {
		return nil
	}
	return portalTemplate
}

func (s *sceneTemplateService) RandomWorldMap() *gametemplate.MapTemplate {
	num := len(s.worldList)
	index := rand.Intn(num)
	return s.worldList[index]
}

func (s *sceneTemplateService) GetYaZhiTemplate(level int32) *gametemplate.YaZhiTemplate {
	var yazhiTemplate *gametemplate.YaZhiTemplate
	for _, tempYazhiTemplate := range s.yaZhiTemplateList {
		if tempYazhiTemplate.LevelMin > level {
			break
		}
		if tempYazhiTemplate.LevelMax < level {
			continue
		}
		yazhiTemplate = tempYazhiTemplate
		break
	}
	return yazhiTemplate
}

func (s *sceneTemplateService) GetPortalTemplateMapByMapId(mapId int32) map[int32]*gametemplate.PortalTemplate {
	tempMap, ok := s.portalTemplateMapOfMap[mapId]
	if !ok {
		return nil
	}
	return tempMap
}

func (s *sceneTemplateService) GetRobotQuestTemplate(sdkTypeList []accounttypes.SDKType, mapId int32) gametemplate.RobotQuestTemplateInterface {
	for _, sdkType := range sdkTypeList {
		sdkRobotQuestTemplateMap, ok := s.sdkRobotQuestTemplateMapOfMap[sdkType]
		if !ok {
			continue
		}
		robotQuestTemplate, ok := sdkRobotQuestTemplateMap[mapId]
		if !ok {
			return nil
		}
		return robotQuestTemplate
	}

	robotQuestTemplate, ok := s.baseRobotQuestTemplateMap[mapId]
	if !ok {
		return nil
	}
	return robotQuestTemplate
}

func (s *sceneTemplateService) GetWorldBossTemplate() []*gametemplate.BiologyTemplate {
	return s.wordBossList
}

func (s *sceneTemplateService) GetDingShiSceneTemplate(mapId int32, biologyId int32) *gametemplate.SceneTemplate {
	dingShiBossMap, ok := s.dingShiBossMap[mapId]
	if !ok {
		return nil
	}
	bossTemplate, ok := dingShiBossMap[biologyId]
	if !ok {
		return nil
	}
	return bossTemplate
}

func (s *sceneTemplateService) GetSceneTemplate(mapId int32, biologyId int32) *gametemplate.SceneTemplate {
	dingShiBossMap, ok := s.dingShiBossMap[mapId]
	if !ok {
		return nil
	}
	bossTemplate, ok := dingShiBossMap[biologyId]
	if !ok {
		return nil
	}
	return bossTemplate
}

var (
	once sync.Once
	s    *sceneTemplateService
)

func Init() (err error) {
	once.Do(func() {
		s = &sceneTemplateService{}
		err = s.init()
	})
	return err
}

func GetSceneTemplateService() SceneTemplateService {
	return s
}
