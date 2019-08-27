package scene

import (
	"fgame/fgame/game/global"
	scenetemplate "fgame/fgame/game/scene/template"
	scenetypes "fgame/fgame/game/scene/types"
	"sync"
)

type SceneService interface {
	Start()
	Stop() (err error)
	GetSceneById(sceneId int64) (s Scene)
	GetSceneByMapId(mapId int32) (s Scene)
	GetFuBenSceneById(sceneId int64) (s Scene)
	GetActivityFuBenSceneById(sceneId int64) (s Scene)
	GetWorldSceneByMapId(mapId int32) (s Scene)
	GetBossSceneByMapId(mapId int32) (s Scene)
	GetActivitySceneByMapId(mapId int32) (s Scene)
	GetMarrySceneByMapId(mapId int32) (s Scene)
	GetTowerSceneByMapId(mapId int32) (s Scene)
	GetAllTowerScene() map[int64]Scene
}

//场景服务 主要管理数据
type sceneService struct {
	rwm sync.RWMutex
	//所有地图
	sceneMap map[int64]Scene
	//世界地图
	worldMap map[int64]Scene
	//副本
	fubenSceneMap map[int64]Scene
	//活动副本
	activityFuBenSceneMap map[int64]Scene
	//活动副本
	activitySceneMap map[int64]Scene
	//世界boss地图
	// worldBossMap map[int64]Scene
	// //跨服世界boss
	// crossWorldBossMap map[int64]Scene
	//结婚地图
	marrySceneMap map[int64]Scene
	//打宝塔地图
	towerMap map[int64]Scene
	//幻境boss
	// unrealBossMap map[int64]Scene
	// //外域boss
	// outlandBossMap map[int64]Scene
	// //藏经阁
	// cangJingGeMap map[int64]Scene
	// //珍惜
	// zhenXiMap map[int64]Scene
	// //定时
	// dingShiMap map[int64]Scene

	bossMapOfMap map[scenetypes.SceneType]map[int64]Scene
}

//场景服务初始化
func (ss *sceneService) init() (err error) {
	ss.sceneMap = make(map[int64]Scene)
	ss.fubenSceneMap = make(map[int64]Scene)
	ss.activitySceneMap = make(map[int64]Scene)
	ss.worldMap = make(map[int64]Scene)
	// ss.worldBossMap = make(map[int64]Scene)
	ss.marrySceneMap = make(map[int64]Scene)
	// ss.crossWorldBossMap = make(map[int64]Scene)
	ss.towerMap = make(map[int64]Scene)
	// ss.unrealBossMap = make(map[int64]Scene)
	// ss.outlandBossMap = make(map[int64]Scene)
	// ss.cangJingGeMap = make(map[int64]Scene)
	// ss.zhenXiMap = make(map[int64]Scene)
	ss.activityFuBenSceneMap = make(map[int64]Scene)
	// ss.dingShiMap = make(map[int64]Scene)
	ss.bossMapOfMap = make(map[scenetypes.SceneType]map[int64]Scene)
	err = ss.initWorld()
	if err != nil {
		return
	}
	err = ss.initBoss()
	if err != nil {
		return
	}
	// err = ss.initWorldBoss()
	// if err != nil {
	// 	return
	// }
	// err = ss.initCrossBoss()
	// if err != nil {
	// 	return
	// }
	// err = ss.initUnrealBoss()
	// if err != nil {
	// 	return
	// }
	// err = ss.initOutlandBoss()
	// if err != nil {
	// 	return
	// }
	// err = ss.initCangJingGe()
	// if err != nil {
	// 	return
	// }

	// err = ss.initZhenXi()
	// if err != nil {
	// 	return
	// }
	// err = ss.initDingShi()
	// if err != nil {
	// 	return
	// }
	return nil
}

func (ss *sceneService) initWorld() (err error) {
	//创建世界场景
	for _, to := range scenetemplate.GetSceneTemplateService().GetAllWorld() {
		if to.GetMap() == nil {
			continue
		}
		currentServerType := global.GetGame().GetServerType()
		if to.GetMapType().GameServerType() != currentServerType {
			continue
		}
		CreateScene(to, 0, nil)
	}
	return nil
}

func (ss *sceneService) initBoss() (err error) {
	//创建世界场景
	for _, to := range scenetemplate.GetSceneTemplateService().GetAllBoss() {
		if to.GetMap() == nil {
			continue
		}
		currentServerType := global.GetGame().GetServerType()
		if to.GetMapType().GameServerType() != currentServerType {
			continue
		}
		CreateScene(to, 0, nil)
	}
	return nil
}

// func (ss *sceneService) initWorldBoss() (err error) {
// 	//创建世界场景
// 	for _, to := range scenetemplate.GetSceneTemplateService().GetAllWorldBoss() {
// 		if to.GetMap() == nil {
// 			continue
// 		}
// 		currentServerType := global.GetGame().GetServerType()
// 		if to.GetMapType().GameServerType() != currentServerType {
// 			continue
// 		}
// 		CreateScene(to, 0, nil)
// 	}
// 	return nil
// }

// func (ss *sceneService) initUnrealBoss() (err error) {
// 	//创建幻境BOSS场景
// 	for _, to := range scenetemplate.GetSceneTemplateService().GetAllUnrealBoss() {
// 		if to.GetMap() == nil {
// 			continue
// 		}
// 		currentServerType := global.GetGame().GetServerType()
// 		if to.GetMapType().GameServerType() != currentServerType {
// 			continue
// 		}
// 		CreateScene(to, 0, nil)
// 	}
// 	return nil
// }

// func (ss *sceneService) initOutlandBoss() (err error) {
// 	//创建外域BOSS场景
// 	for _, to := range scenetemplate.GetSceneTemplateService().GetAllOutlandBoss() {
// 		if to.GetMap() == nil {
// 			continue
// 		}
// 		currentServerType := global.GetGame().GetServerType()
// 		if to.GetMapType().GameServerType() != currentServerType {
// 			continue
// 		}
// 		CreateScene(to, 0, nil)
// 	}
// 	return nil
// }

// func (ss *sceneService) initCangJingGe() (err error) {
// 	//创建藏经阁场景
// 	for _, to := range scenetemplate.GetSceneTemplateService().GetAllCangJingGe() {
// 		if to.GetMap() == nil {
// 			continue
// 		}
// 		currentServerType := global.GetGame().GetServerType()
// 		if to.GetMapType().GameServerType() != currentServerType {
// 			continue
// 		}
// 		CreateScene(to, 0, nil)
// 	}
// 	return nil
// }

// func (ss *sceneService) initZhenXi() (err error) {
// 	//创建藏经阁场景
// 	for _, to := range scenetemplate.GetSceneTemplateService().GetAllZhenXi() {
// 		if to.GetMap() == nil {
// 			continue
// 		}
// 		currentServerType := global.GetGame().GetServerType()
// 		if to.GetMapType().GameServerType() != currentServerType {
// 			continue
// 		}
// 		CreateScene(to, 0, nil)
// 	}
// 	return nil
// }

// func (ss *sceneService) initDingShi() (err error) {
// 	//创建藏经阁场景
// 	for _, to := range scenetemplate.GetSceneTemplateService().GetAllDingShi() {
// 		if to.GetMap() == nil {
// 			continue
// 		}
// 		currentServerType := global.GetGame().GetServerType()
// 		if to.GetMapType().GameServerType() != currentServerType {
// 			continue
// 		}
// 		CreateScene(to, 0, nil)
// 	}
// 	return nil
// }

// func (ss *sceneService) initCrossBoss() (err error) {
// 	//创建世界场景
// 	for _, to := range scenetemplate.GetSceneTemplateService().GetAllCrossWorldBoss() {
// 		if to.GetMap() == nil {
// 			continue
// 		}
// 		currentServerType := global.GetGame().GetServerType()
// 		if to.GetMapType().GameServerType() != currentServerType {
// 			continue
// 		}
// 		CreateScene(to, 0, nil)
// 	}
// 	return nil
// }

func (ss *sceneService) OnSceneStart(s Scene) {
	ss.rwm.Lock()
	defer ss.rwm.Unlock()
	ss.addScene(s)
}

func (ss *sceneService) OnSceneStop(s Scene) {
	ss.rwm.Lock()
	defer ss.rwm.Unlock()
	ss.removeScene(s.Id())
}

//开始
func (ss *sceneService) Start() {
	return
}

//停止
func (ss *sceneService) Stop() (err error) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	for _, s := range ss.sceneMap {
		s.AsyncStop()
	}
	return nil
}

func (ss *sceneService) GetSceneById(sceneId int64) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	s, ok := ss.sceneMap[sceneId]
	if !ok {
		return nil
	}
	return
}

func (ss *sceneService) GetWorldSceneByMapId(mapId int32) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	for _, ts := range ss.worldMap {
		if ts.MapId() == mapId {
			return ts
		}
	}
	return nil
}

func (ss *sceneService) GetBossSceneByMapId(mapId int32) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	mapTemplate := scenetemplate.GetSceneTemplateService().GetMap(mapId)
	if mapTemplate == nil {
		return nil
	}
	bossMap, ok := ss.bossMapOfMap[mapTemplate.GetMapType()]
	if !ok {
		return nil
	}
	for _, ts := range bossMap {
		if ts.MapId() == mapId {
			return ts
		}
	}

	return nil
}

// func (ss *sceneService) GetCrossWorldSceneByMapId(mapId int32) (s Scene) {
// 	ss.rwm.RLock()
// 	defer ss.rwm.RUnlock()
// 	for _, ts := range ss.crossWorldBossMap {
// 		if ts.MapId() == mapId {
// 			return ts
// 		}
// 	}
// 	return nil
// }

// func (ss *sceneService) GetWorldBossSceneByMapId(mapId int32) (s Scene) {
// 	ss.rwm.RLock()
// 	defer ss.rwm.RUnlock()
// 	for _, ts := range ss.worldBossMap {
// 		if ts.MapId() == mapId {
// 			return ts
// 		}
// 	}
// 	return nil
// }

// func (ss *sceneService) GetUnrealBossSceneByMapId(mapId int32) (s Scene) {
// 	ss.rwm.RLock()
// 	defer ss.rwm.RUnlock()
// 	for _, ts := range ss.unrealBossMap {
// 		if ts.MapId() == mapId {
// 			return ts
// 		}
// 	}
// 	return nil
// }

// func (ss *sceneService) GetOutlandBossSceneByMapId(mapId int32) (s Scene) {
// 	ss.rwm.RLock()
// 	defer ss.rwm.RUnlock()
// 	for _, ts := range ss.outlandBossMap {
// 		if ts.MapId() == mapId {
// 			return ts
// 		}
// 	}
// 	return nil
// }

// func (ss *sceneService) GetCangJingGeSceneByMapId(mapId int32) (s Scene) {
// 	ss.rwm.RLock()
// 	defer ss.rwm.RUnlock()
// 	for _, ts := range ss.cangJingGeMap {
// 		if ts.MapId() == mapId {
// 			return ts
// 		}
// 	}
// 	return nil
// }

// func (ss *sceneService) GetZhenXiSceneByMapId(mapId int32) (s Scene) {
// 	ss.rwm.RLock()
// 	defer ss.rwm.RUnlock()
// 	for _, ts := range ss.zhenXiMap {
// 		if ts.MapId() == mapId {
// 			return ts
// 		}
// 	}
// 	return nil
// }

// func (ss *sceneService) GetDingShiSceneByMapId(mapId int32) (s Scene) {
// 	ss.rwm.RLock()
// 	defer ss.rwm.RUnlock()
// 	for _, ts := range ss.dingShiMap {
// 		if ts.MapId() == mapId {
// 			return ts
// 		}
// 	}
// 	return nil
// }

func (ss *sceneService) GetActivitySceneByMapId(mapId int32) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	for _, ts := range ss.activitySceneMap {
		if ts.MapId() == mapId {
			return ts
		}
	}
	return nil
}

func (ss *sceneService) GetMarrySceneByMapId(mapId int32) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	for _, ts := range ss.marrySceneMap {
		if ts.MapId() == mapId {
			return ts
		}
	}
	return nil
}

func (ss *sceneService) GetFuBenSceneById(sceneId int64) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	s, ok := ss.fubenSceneMap[sceneId]
	if !ok {
		return nil
	}
	return s
}

func (ss *sceneService) GetTowerSceneByMapId(mapId int32) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	for _, ts := range ss.towerMap {
		if ts.MapId() == mapId {
			return ts
		}
	}
	return nil
}

func (ss *sceneService) GetAllTowerScene() map[int64]Scene {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	return ss.towerMap
}

func (ss *sceneService) GetSceneByMapId(mapId int32) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	for _, ts := range ss.sceneMap {
		if ts.MapId() == mapId {
			return ts
		}
	}
	return nil
}

func (ss *sceneService) GetActivityFuBenSceneById(sceneId int64) (s Scene) {
	ss.rwm.RLock()
	defer ss.rwm.RUnlock()
	s, ok := ss.activityFuBenSceneMap[sceneId]
	if !ok {
		return nil
	}
	return nil
}

func (ss *sceneService) addScene(s Scene) (flag bool) {
	_, exist := ss.sceneMap[s.Id()]
	if exist {
		return false
	}
	ss.sceneMap[s.Id()] = s
	switch s.MapTemplate().GetMapType().MapType() {
	case scenetypes.MapTypeWorld:
		ss.worldMap[s.Id()] = s
		break
	case scenetypes.MapTypeBoss:
		bossMap, ok := ss.bossMapOfMap[s.MapTemplate().GetMapType()]
		if !ok {
			bossMap = make(map[int64]Scene)
			ss.bossMapOfMap[s.MapTemplate().GetMapType()] = bossMap
		}
		bossMap[s.Id()] = s
	// case scenetypes.MapTypeWorldBoss:
	// 	ss.worldBossMap[s.Id()] = s
	// 	break
	// case scenetypes.MapTypeCrossWorldBoss:
	// 	ss.crossWorldBossMap[s.Id()] = s
	// 	break
	case scenetypes.MapTypeFuBen:
		ss.fubenSceneMap[s.Id()] = s
		break
	case scenetypes.MapTypeActivityFuBen:
		ss.activityFuBenSceneMap[s.Id()] = s
		break
	case scenetypes.MapTypeActivity,
		scenetypes.MapTypeActivitySub:
		ss.activitySceneMap[s.Id()] = s
		break
	case scenetypes.MapTypeMarry:
		ss.marrySceneMap[s.Id()] = s
		break
	case scenetypes.MapTypeTower:
		ss.towerMap[s.Id()] = s
		break
		// case scenetypes.MapTypeUnrealBoss:
		// 	ss.unrealBossMap[s.Id()] = s
		// 	break
		// case scenetypes.MapTypeOutlandBoss:
		// 	ss.outlandBossMap[s.Id()] = s
		// 	break
		// case scenetypes.MapTypeCangJingGe:
		// 	ss.cangJingGeMap[s.Id()] = s
		// 	break
		// case scenetypes.MapTypeZhenXiBoss:
		// 	ss.zhenXiMap[s.Id()] = s
		// 	break
		// case scenetypes.MapTypeDingShiBoss:
		// 	ss.dingShiMap[s.Id()] = s
		// break
	}

	return true
}

func (ss *sceneService) removeScene(id int64) {
	s, exist := ss.sceneMap[id]
	if !exist {
		panic("remove no exist scene")
	}
	delete(ss.sceneMap, id)

	switch s.MapTemplate().GetMapType().MapType() {
	case scenetypes.MapTypeWorld:
		delete(ss.worldMap, id)
		break
	case scenetypes.MapTypeBoss:
		bossMap, ok := ss.bossMapOfMap[s.MapTemplate().GetMapType()]
		if ok {
			delete(bossMap, id)
		}
	// case scenetypes.MapTypeWorldBoss:
	// 	delete(ss.worldBossMap, id)
	// 	break
	// case scenetypes.MapTypeCrossWorldBoss:
	// 	delete(ss.crossWorldBossMap, id)
	// 	break
	case scenetypes.MapTypeFuBen:
		delete(ss.fubenSceneMap, id)
		break
	case scenetypes.MapTypeActivityFuBen:
		delete(ss.activityFuBenSceneMap, id)
		break
	case scenetypes.MapTypeActivity,
		scenetypes.MapTypeActivitySub:
		delete(ss.activitySceneMap, id)
		break
	case scenetypes.MapTypeMarry:
		delete(ss.marrySceneMap, id)
		break
	case scenetypes.MapTypeTower:
		delete(ss.towerMap, id)
		break
		// case scenetypes.MapTypeUnrealBoss:
		// 	delete(ss.unrealBossMap, id)
		// 	break
		// case scenetypes.MapTypeOutlandBoss:
		// 	delete(ss.outlandBossMap, id)
		// 	break
		// case scenetypes.MapTypeCangJingGe:
		// 	delete(ss.cangJingGeMap, id)
		// 	break
		// case scenetypes.MapTypeZhenXiBoss:
		// 	delete(ss.zhenXiMap, id)
		// 	break
		// case scenetypes.MapTypeDingShiBoss:
		// 	delete(ss.dingShiMap, id)
		// 	break
	}
}

var (
	once sync.Once
	cs   *sceneService
	// dingShiOnce sync.Once
	// ws          *dingShiService
)

// func InitDingShi() (err error) {
// 	dingShiOnce.Do(func() {
// 		ws = &dingShiService{}
// 		err = ws.init()
// 		if err != nil {
// 			return
// 		}

// 		return
// 	})
// 	return err
// }

func Init() (err error) {
	once.Do(func() {

		cs = &sceneService{}
		err = cs.init()

		return
	})
	return err
}

func GetSceneService() SceneService {
	return cs
}
