package equipbaoku

import (
	"fgame/fgame/core/runner"
	"fgame/fgame/game/constant/constant"
	constanttypes "fgame/fgame/game/constant/types"
	droptemplate "fgame/fgame/game/drop/template"
	dummytemplate "fgame/fgame/game/dummy/template"
	equipbaokutypes "fgame/fgame/game/equipbaoku/types"
	"fgame/fgame/game/global"
	"fgame/fgame/pkg/mathutils"
	"sync"

	log "github.com/Sirupsen/logrus"
)

//装备宝库接口处理
type EquipBaoKuService interface {
	runner.Task
	//添加日志
	AddLog(plName string, itemId, num int32, typ equipbaokutypes.BaoKuType)
	//获取日志
	GetLogByTime(time int64, typ equipbaokutypes.BaoKuType) []*EquipBaoKuLogObject
	//清空日志
	GMClearLog()
}

type equipBaoKuService struct {
	rwm sync.RWMutex
	//装备宝库日志列表
	equipBaoKuLogList []*EquipBaoKuLogObject
	//材料宝库日志列表
	materialsBaoKuLogList []*EquipBaoKuLogObject
	//上次系统插入日志时间
	lastAddEquipDummyLogLime    int64
	lastAddMaterialDummyLogLime int64
}

//初始化
func (s *equipBaoKuService) init() (err error) {
	return
}

//TODO 修改为n秒加一次
//心跳
func (s *equipBaoKuService) Heartbeat() {

	err := s.addEquipDummyLog()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("equipBaoKu:系统生成装备假日志,错误")
		return
	}

	err = s.addMaterialDummyLog()
	if err != nil {
		log.WithFields(
			log.Fields{
				"error": err,
			}).Error("equipBaoKu:系统生成材料假日志,错误")
		return
	}
}

//生成系统假日志
func (s *equipBaoKuService) addEquipDummyLog() (err error) {
	now := global.GetGame().GetTimeService().Now()
	lastTime := s.lastAddEquipDummyLogLime
	diffTime := now - lastTime
	randTime := s.getRandomLogTime()
	if diffTime < randTime {
		return
	}

	name := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	dropId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEquipBaoKuDummyDrop)
	dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
	if dropData == nil {
		return
	}
	s.AddLog(name, dropData.ItemId, dropData.Num, equipbaokutypes.BaoKuTypeEquip)

	s.lastAddEquipDummyLogLime = now
	return
}

func (s *equipBaoKuService) addMaterialDummyLog() (err error) {
	now := global.GetGame().GetTimeService().Now()
	lastTime := s.lastAddMaterialDummyLogLime
	diffTime := now - lastTime
	randTime := s.getRandomLogTime()
	if diffTime < randTime {
		return
	}

	name := dummytemplate.GetDummyTemplateService().GetGameRandomDummyName()
	dropId := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeMaterialBaoKuDummyDrop)
	dropData := droptemplate.GetDropTemplateService().GetDropItemLevel(dropId)
	if dropData == nil {
		return
	}
	s.AddLog(name, dropData.ItemId, dropData.Num, equipbaokutypes.BaoKuTypeMaterials)

	s.lastAddMaterialDummyLogLime = now
	return
}

func (s *equipBaoKuService) GetList(typ equipbaokutypes.BaoKuType) []*EquipBaoKuLogObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()
	if typ == equipbaokutypes.BaoKuTypeEquip {
		return s.equipBaoKuLogList
	}
	return s.materialsBaoKuLogList
}

func (s *equipBaoKuService) AddLog(plName string, itemId, num int32, typ equipbaokutypes.BaoKuType) {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.appendLog(plName, itemId, num, typ)
}

func (s *equipBaoKuService) GetLogByTime(time int64, typ equipbaokutypes.BaoKuType) []*EquipBaoKuLogObject {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	if typ == equipbaokutypes.BaoKuTypeEquip {
		for index, log := range s.equipBaoKuLogList {
			if time < log.updateTime {
				return s.equipBaoKuLogList[index:]
			}
		}
	} else {
		for index, log := range s.materialsBaoKuLogList {
			if time < log.updateTime {
				return s.materialsBaoKuLogList[index:]
			}
		}
	}

	return nil
}

func (s *equipBaoKuService) GMClearLog() {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	s.equipBaoKuLogList = nil
	s.materialsBaoKuLogList = nil
}

func (s *equipBaoKuService) appendLog(playerName string, itemId, itemNum int32, typ equipbaokutypes.BaoKuType) {
	obj := s.createLogObj(playerName, itemId, itemNum, typ)
	if typ == equipbaokutypes.BaoKuTypeEquip {
		s.equipBaoKuLogList = append(s.equipBaoKuLogList, obj)
	} else {
		s.materialsBaoKuLogList = append(s.materialsBaoKuLogList, obj)
	}
}

func (s *equipBaoKuService) createLogObj(playerName string, itemId, itemNum int32, typ equipbaokutypes.BaoKuType) *EquipBaoKuLogObject {
	now := global.GetGame().GetTimeService().Now()
	maxLogLen := constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEquipBaoKuLogMaxNum)
	var obj *EquipBaoKuLogObject
	if typ == equipbaokutypes.BaoKuTypeEquip {
		// 装备宝库日志
		if len(s.equipBaoKuLogList) >= int(maxLogLen) {
			obj = s.equipBaoKuLogList[0]
			s.equipBaoKuLogList = s.equipBaoKuLogList[1:]
		} else {
			obj = NewEquipBaoKuLogObject()
			obj.createTime = now
		}
	} else {
		// 材料宝库日志
		if len(s.materialsBaoKuLogList) >= int(maxLogLen) {
			obj = s.materialsBaoKuLogList[0]
			s.materialsBaoKuLogList = s.materialsBaoKuLogList[1:]
		} else {
			obj = NewEquipBaoKuLogObject()
			obj.createTime = now
		}
	}

	obj.playerName = playerName
	obj.typ = typ
	obj.itemId = itemId
	obj.itemNum = itemNum
	obj.updateTime = now

	return obj
}

//系统假日志生成间隔
func (s *equipBaoKuService) getRandomLogTime() int64 {
	min := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEquipBaoKuLogAddTimeMin))
	max := int(constant.GetConstantService().GetConstant(constanttypes.ConstantTypeEquipBaoKuLogAddTimeMax))
	randTime := int64(mathutils.RandomRange(min, max))
	return randTime
}

var (
	once sync.Once
	cs   *equipBaoKuService
)

func Init() (err error) {
	once.Do(func() {
		cs = &equipBaoKuService{}
		err = cs.init()
	})
	return err
}

func GetEquipBaoKuService() EquipBaoKuService {
	return cs
}
