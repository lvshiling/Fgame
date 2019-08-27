package player

import (
	"fgame/fgame/game/common/common"
	funcopentypes "fgame/fgame/game/funcopen/types"
	"fgame/fgame/game/global"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/game/zhenfa/dao"
	zhenfatemplate "fgame/fgame/game/zhenfa/template"
	zhenfatypes "fgame/fgame/game/zhenfa/types"
	"fgame/fgame/pkg/idutil"
	"fgame/fgame/pkg/mathutils"
)

//玩家阵法管理器
type PlayerZhenFaDataManager struct {
	p player.Player
	//玩家阵法
	playerZhenFaMap map[zhenfatypes.ZhenFaType]*PlayerZhenFaObject
	//玩家阵旗
	playerZhenQiMap map[zhenfatypes.ZhenFaType]map[zhenfatypes.ZhenQiType]*PlayerZhenQiObject
	//玩家阵旗仙火
	playerZhenQiXianHuoMap map[zhenfatypes.ZhenFaType]*PlayerZhenQiXianHuoObject
	// 玩家阵法战力
	playerZhenFaPowerObject *PlayerZhenFaPowerObject
}

func (m *PlayerZhenFaDataManager) Player() player.Player {
	return m.p
}

//加载
func (m *PlayerZhenFaDataManager) Load() (err error) {
	err = m.loadZhenFa()
	if err != nil {
		return
	}
	err = m.loadZhenQi()
	if err != nil {
		return
	}
	err = m.loadZhenQiXianHuo()
	if err != nil {
		return
	}
	err = m.loadZhenFaPower()
	if err != nil {
		return
	}
	return nil
}

func (m *PlayerZhenFaDataManager) loadZhenFa() (err error) {
	m.playerZhenFaMap = make(map[zhenfatypes.ZhenFaType]*PlayerZhenFaObject)
	zhenFaList, err := dao.GetZhenFaDao().GetZhenFaList(m.p.GetId())
	if err != nil {
		return
	}

	for _, zhenFaInfo := range zhenFaList {
		obj := NewPlayerZhenFaObject(m.p)
		obj.FromEntity(zhenFaInfo)
		m.playerZhenFaMap[obj.GetZhenFaType()] = obj
	}
	return
}

func (m *PlayerZhenFaDataManager) loadZhenQi() (err error) {
	m.playerZhenQiMap = make(map[zhenfatypes.ZhenFaType]map[zhenfatypes.ZhenQiType]*PlayerZhenQiObject)
	zhenQiList, err := dao.GetZhenFaDao().GetZhenQiList(m.p.GetId())
	if err != nil {
		return
	}

	for _, zhenQiInfo := range zhenQiList {
		obj := NewPlayerZhenQiObject(m.p)
		obj.FromEntity(zhenQiInfo)
		m.addZhenQi(obj)
	}
	return
}

func (m *PlayerZhenFaDataManager) loadZhenQiXianHuo() (err error) {
	m.playerZhenQiXianHuoMap = make(map[zhenfatypes.ZhenFaType]*PlayerZhenQiXianHuoObject)
	zhenQiXianHuoList, err := dao.GetZhenFaDao().GetZhenQiXianHuoList(m.p.GetId())
	if err != nil {
		return
	}

	for _, zhenQiXianHuo := range zhenQiXianHuoList {
		obj := NewPlayerZhenQiXianHuoObject(m.p)
		obj.FromEntity(zhenQiXianHuo)
		m.playerZhenQiXianHuoMap[obj.GetZhenFaType()] = obj
	}
	return
}

func (m *PlayerZhenFaDataManager) loadZhenFaPower() (err error) {
	powerEntity, err := dao.GetZhenFaDao().GetZhenFaPowerEntity(m.p.GetId())
	if err != nil {
		return
	}

	now := global.GetGame().GetTimeService().Now()
	obj := NewPlayerZhenFaPowerObject(m.p)

	if powerEntity == nil {
		id, _ := idutil.GetId()
		obj.id = id
		obj.createTime = now
		obj.SetModified()
	} else {
		obj.FromEntity(powerEntity)
	}
	m.playerZhenFaPowerObject = obj

	return
}

func (m *PlayerZhenFaDataManager) addZhenQi(obj *PlayerZhenQiObject) {
	zhenQiMap, ok := m.playerZhenQiMap[obj.GetZhenFaType()]
	if !ok {
		zhenQiMap = make(map[zhenfatypes.ZhenQiType]*PlayerZhenQiObject)
		m.playerZhenQiMap[obj.GetZhenFaType()] = zhenQiMap
	}
	zhenQiMap[obj.GetZhenQiPos()] = obj
}

//加载后
func (m *PlayerZhenFaDataManager) AfterLoad() (err error) {
	return
}

//心跳
func (m *PlayerZhenFaDataManager) Heartbeat() {

}

func (m *PlayerZhenFaDataManager) initZhenFa(zhenFaType zhenfatypes.ZhenFaType) (obj *PlayerZhenFaObject) {
	obj = m.GetZhenFaByType(zhenFaType)
	if obj != nil {
		return
	}
	obj = NewPlayerZhenFaObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	obj.typ = zhenFaType
	obj.level = 0
	obj.levelNum = 0
	obj.levelPro = 0
	obj.createTime = now
	obj.SetModified()
	m.playerZhenFaMap[obj.GetZhenFaType()] = obj
	m.initZhenQiAll(zhenFaType)
	m.initZhenFaXianHuo(zhenFaType)
	return
}

func (m *PlayerZhenFaDataManager) initZhenQiAll(zhenFaType zhenfatypes.ZhenFaType) {
	for zhenQiType := zhenfatypes.ZhenQiTypeMin; zhenQiType <= zhenfatypes.ZhenQiTypeMax; zhenQiType++ {
		obj := m.GetZhenQi(zhenFaType, zhenQiType)
		if obj != nil {
			continue
		}
		obj = NewPlayerZhenQiObject(m.p)
		now := global.GetGame().GetTimeService().Now()
		id, _ := idutil.GetId()
		obj.id = id
		obj.typ = zhenFaType
		obj.zhenQiPos = zhenQiType
		obj.number = 0
		obj.createTime = now
		obj.SetModified()
		m.addZhenQi(obj)
	}
}

func (m *PlayerZhenFaDataManager) initZhenFaXianHuo(zhenFaType zhenfatypes.ZhenFaType) (obj *PlayerZhenQiXianHuoObject) {
	obj = m.GetZhenQiXianHuo(zhenFaType)
	if obj != nil {
		return
	}
	obj = NewPlayerZhenQiXianHuoObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	obj.id = id
	obj.typ = zhenFaType
	obj.level = 0
	obj.luckyStar = 0
	obj.createTime = now
	obj.SetModified()
	m.playerZhenQiXianHuoMap[zhenFaType] = obj
	return
}

func (m *PlayerZhenFaDataManager) GetZhenFaMap() map[zhenfatypes.ZhenFaType]*PlayerZhenFaObject {
	return m.playerZhenFaMap
}

func (m *PlayerZhenFaDataManager) GetZhenQiMap() map[zhenfatypes.ZhenFaType]map[zhenfatypes.ZhenQiType]*PlayerZhenQiObject {
	return m.playerZhenQiMap
}

func (m *PlayerZhenFaDataManager) GetZhenQiXianHuoMap() map[zhenfatypes.ZhenFaType]*PlayerZhenQiXianHuoObject {
	return m.playerZhenQiXianHuoMap
}

func (m *PlayerZhenFaDataManager) GetZhenFaZhenQiMap(zhenFaType zhenfatypes.ZhenFaType) map[zhenfatypes.ZhenQiType]*PlayerZhenQiObject {
	playerZhenQiMap, ok := m.playerZhenQiMap[zhenFaType]
	if !ok {
		return nil
	}
	return playerZhenQiMap
}

func (m *PlayerZhenFaDataManager) GetZhenQiXianHuo(zhenFaType zhenfatypes.ZhenFaType) *PlayerZhenQiXianHuoObject {
	obj, ok := m.playerZhenQiXianHuoMap[zhenFaType]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerZhenFaDataManager) GetZhenQi(zhenFaType zhenfatypes.ZhenFaType, zhenQiType zhenfatypes.ZhenQiType) *PlayerZhenQiObject {
	playerZhenQiMap, ok := m.playerZhenQiMap[zhenFaType]
	if !ok {
		return nil
	}
	obj, ok := playerZhenQiMap[zhenQiType]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerZhenFaDataManager) GetZhenFaByType(zhenFaType zhenfatypes.ZhenFaType) *PlayerZhenFaObject {
	obj, ok := m.playerZhenFaMap[zhenFaType]
	if !ok {
		return nil
	}
	return obj
}

func (m *PlayerZhenFaDataManager) GetAllZhenFaLevel() (totalLevel int32) {
	for _, obj := range m.playerZhenFaMap {
		totalLevel += obj.GetLevel()
	}
	return
}

func (m *PlayerZhenFaDataManager) Activate(zhenFaType zhenfatypes.ZhenFaType) (obj *PlayerZhenFaObject, flag bool) {
	if !m.p.IsFuncOpen(funcopentypes.FuncOpenTypeZhenFa) {
		return
	}
	obj = m.GetZhenFaByType(zhenFaType)
	if obj != nil {
		return
	}
	zhenFaJiHuoTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaJiHuoTemplate(zhenFaType)
	if zhenFaJiHuoTemplate == nil {
		return
	}

	if zhenFaJiHuoTemplate.NeedZhenFaType != 0 {
		needZhenFaType := zhenfatypes.ZhenFaType(zhenFaJiHuoTemplate.NeedZhenFaType)
		needObj := m.GetZhenFaByType(needZhenFaType)
		if needObj == nil || needObj.GetLevel() < zhenFaJiHuoTemplate.NeedZhenFaLevel {
			return
		}
	}
	obj = m.initZhenFa(zhenFaType)
	flag = true
	return
}

func (m *PlayerZhenFaDataManager) ZhenFaShengJi(zhenFaTpe zhenfatypes.ZhenFaType, sucess bool, pro int32) (flag bool) {
	obj := m.GetZhenFaByType(zhenFaTpe)
	if obj == nil {
		return
	}
	curLevel := obj.GetLevel()
	nextLevel := curLevel + 1
	zhenFaTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaTempalte(zhenFaTpe, nextLevel)
	if zhenFaTemplate == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		obj.level = zhenFaTemplate.Level
		obj.levelPro = 0
		obj.levelNum = 0
	} else {
		obj.levelNum++
		obj.levelPro += pro
	}
	obj.updateTime = now
	obj.SetModified()
	flag = true
	return
}

func (m *PlayerZhenFaDataManager) ZhenFaXianHuoShengJi(zhenFaType zhenfatypes.ZhenFaType, sucess bool, pro int32, protectFlag bool) (flag bool) {
	obj := m.GetZhenQiXianHuo(zhenFaType)
	if obj == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	luckStar := obj.GetLuckyStar()
	curLevel := obj.GetLevel()
	nextLevel := curLevel + 1
	zhenFaXianHuoTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaXianHuoTemplate(zhenFaType, nextLevel)
	if zhenFaXianHuoTemplate == nil {
		return
	}
	if sucess {
		if luckStar == 5 {
			nextLevel += 1
			nextZhenFaXianHuoTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaXianHuoTemplate(zhenFaType, nextLevel)
			if nextZhenFaXianHuoTemplate == nil {
				nextLevel -= 1
			}
			obj.luckyStar = 0
		}
		obj.levelPro = 0
		obj.levelNum = 0

		obj.level = nextLevel
		obj.updateTime = now
		obj.SetModified()
	} else {
		if luckStar < 5 {
			obj.luckyStar += 1
		}
		//降级
		if !protectFlag {
			if mathutils.RandomHit(common.MAX_RATE, int(zhenFaXianHuoTemplate.ReturnRate)) {
				if zhenFaXianHuoTemplate.GetReturnLevelTemplate() != nil {
					obj.level = zhenFaXianHuoTemplate.GetReturnLevelTemplate().Level
				}
			}
		}
		obj.levelNum += 1
		obj.levelPro += pro
		obj.updateTime = now
		obj.SetModified()
	}
	flag = true
	return
}

func (m *PlayerZhenFaDataManager) ZhenQiAdvanced(zhenFaType zhenfatypes.ZhenFaType, zhenQiType zhenfatypes.ZhenQiType, sucess bool, pro int32) (flag bool) {
	obj := m.GetZhenQi(zhenFaType, zhenQiType)
	if obj == nil {
		return
	}
	curNumber := obj.GetNumber()
	nextNumber := curNumber + 1
	zhenQiTemplate := zhenfatemplate.GetZhenFaTemplateService().GetZhenFaZhenQiTemplate(zhenFaType, zhenQiType, nextNumber)
	if zhenQiTemplate == nil {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		obj.number = nextNumber
		obj.numberPro = 0
		obj.numberNum = 0
	} else {
		obj.numberPro += pro
		obj.numberNum += 1
	}
	obj.updateTime = now
	obj.SetModified()
	flag = true
	return
}

func (m *PlayerZhenFaDataManager) GetPower() int64 {
	return m.playerZhenFaPowerObject.power
}

func (m *PlayerZhenFaDataManager) SetPower(power int64) {
	if power < 0 {
		return
	}
	if m.playerZhenFaPowerObject.power == power {
		return
	}
	now := global.GetGame().GetTimeService().Now()
	m.playerZhenFaPowerObject.power = power
	m.playerZhenFaPowerObject.updateTime = now
	m.playerZhenFaPowerObject.SetModified()
}

func CreatePlayerZhenFaDataManager(p player.Player) player.PlayerDataManager {
	m := &PlayerZhenFaDataManager{}
	m.p = p
	return m
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerZhenFaDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerZhenFaDataManager))
}
