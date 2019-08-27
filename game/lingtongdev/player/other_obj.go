package player

import (
	"fgame/fgame/core/storage"
	"fgame/fgame/game/global"
	"fgame/fgame/game/lingtongdev/dao"
	"fgame/fgame/game/lingtongdev/entity"
	lingtongdevtemplate "fgame/fgame/game/lingtongdev/template"
	"fgame/fgame/game/lingtongdev/types"
	lingtongdevtypes "fgame/fgame/game/lingtongdev/types"
	"fgame/fgame/game/player"
	playertypes "fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
	"fmt"
)

//灵童养成非进阶对象
type PlayerLingTongOtherObject struct {
	player     player.Player
	id         int64
	playerId   int64
	classType  lingtongdevtypes.LingTongDevSysType
	typ        lingtongdevtypes.LingTongDevType
	seqId      int32
	level      int32
	upNum      int32
	upPro      int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func NewPlayerLingTongOtherObject(pl player.Player) *PlayerLingTongOtherObject {
	pwo := &PlayerLingTongOtherObject{
		player: pl,
	}
	return pwo
}

func convertLingTongOtherObjectToEntity(pwo *PlayerLingTongOtherObject) (*entity.PlayerLingTongOtherEntity, error) {

	e := &entity.PlayerLingTongOtherEntity{
		Id:         pwo.id,
		PlayerId:   pwo.playerId,
		ClassType:  int32(pwo.classType),
		Type:       int32(pwo.typ),
		SeqId:      pwo.seqId,
		Level:      pwo.level,
		UpNum:      pwo.upNum,
		UpPro:      pwo.upPro,
		UpdateTime: pwo.updateTime,
		CreateTime: pwo.createTime,
		DeleteTime: pwo.deleteTime,
	}
	return e, nil
}

func (pwo *PlayerLingTongOtherObject) GetPlayerId() int64 {
	return pwo.playerId
}

func (pwo *PlayerLingTongOtherObject) GetDBId() int64 {
	return pwo.id
}

func (pwo *PlayerLingTongOtherObject) GetClassType() lingtongdevtypes.LingTongDevSysType {
	return pwo.classType
}

func (pwo *PlayerLingTongOtherObject) GetType() lingtongdevtypes.LingTongDevType {
	return pwo.typ
}

func (pwo *PlayerLingTongOtherObject) GetSeqId() int32 {
	return pwo.seqId
}

func (pwo *PlayerLingTongOtherObject) GetLevel() int32 {
	return pwo.level
}

func (pwo *PlayerLingTongOtherObject) GetUpNum() int32 {
	return pwo.upNum
}

func (pwo *PlayerLingTongOtherObject) GetUpPro() int32 {
	return pwo.upPro
}

func (pwo *PlayerLingTongOtherObject) UpSatr(pro int32, sucess bool) (flag bool) {
	now := global.GetGame().GetTimeService().Now()
	if sucess {
		lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(pwo.classType, int(pwo.seqId))
		if lingTongDevTemplate == nil {
			return
		}
		lingTongDevUpstarTemplate := lingTongDevTemplate.GetLingTongDevUpstarByLevel(pwo.level + 1)
		if lingTongDevUpstarTemplate == nil {
			return
		}
		pwo.level += 1
		pwo.upNum = 0
		pwo.upPro = pro
	} else {
		pwo.upNum += 1
		pwo.upPro += pro
	}
	pwo.updateTime = now
	pwo.SetModified()
	return true
}

func (pwo *PlayerLingTongOtherObject) ToEntity() (e storage.Entity, err error) {
	e, err = convertLingTongOtherObjectToEntity(pwo)
	return e, err
}

func (pwo *PlayerLingTongOtherObject) FromEntity(e storage.Entity) error {
	pse, _ := e.(*entity.PlayerLingTongOtherEntity)

	pwo.id = pse.Id
	pwo.playerId = pse.PlayerId
	pwo.classType = lingtongdevtypes.LingTongDevSysType(pse.ClassType)
	pwo.typ = lingtongdevtypes.LingTongDevType(pse.Type)
	pwo.seqId = pse.SeqId
	pwo.level = pse.Level
	pwo.upNum = pse.UpNum
	pwo.upPro = pse.UpPro
	pwo.updateTime = pse.UpdateTime
	pwo.createTime = pse.CreateTime
	pwo.deleteTime = pse.DeleteTime
	return nil
}

func (pwo *PlayerLingTongOtherObject) SetModified() {
	e, err := pwo.ToEntity()
	if err != nil {
		panic(fmt.Errorf("lingtongother: err[%s]", err.Error()))
	}
	obj, ok := e.(playertypes.PlayerDataEntity)
	if !ok {
		panic("never reach here")
	}

	pwo.player.AddChangedObject(obj)
	return
}

type LingTongOtherContainer struct {
	p         player.Player
	classType types.LingTongDevSysType
	otherMap  map[types.LingTongDevType]map[int32]*PlayerLingTongOtherObject
}

func NewLingTongOtherContainer(p player.Player, classType types.LingTongDevSysType) *LingTongOtherContainer {
	d := &LingTongOtherContainer{
		p:         p,
		classType: classType,
		otherMap:  make(map[types.LingTongDevType]map[int32]*PlayerLingTongOtherObject),
	}
	return d
}

func (m *LingTongOtherContainer) GetOtherMap() map[types.LingTongDevType]map[int32]*PlayerLingTongOtherObject {
	return m.otherMap
}

func (m *LingTongOtherContainer) GetOtherObj(seqId int32) *PlayerLingTongOtherObject {
	return m.getOtherObj(seqId)
}

func (m *LingTongOtherContainer) getOtherObj(seqId int32) *PlayerLingTongOtherObject {
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(m.classType, int(seqId))
	if lingTongDevTemplate == nil {
		return nil
	}
	typ := lingTongDevTemplate.GetType()
	otherMap, ok := m.otherMap[typ]
	if !ok {
		return nil
	}
	obj, ok := otherMap[seqId]
	if !ok {
		return nil
	}
	return obj
}

//增加非进阶灵童养成
func (m *LingTongOtherContainer) newPlayerLingTongOtherObject(seqId int32) (pwo *PlayerLingTongOtherObject) {
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(m.classType, int(seqId))
	if lingTongDevTemplate == nil {
		return nil
	}
	typ := lingTongDevTemplate.GetType()
	otherMap, exist := m.otherMap[typ]
	if !exist {
		otherMap = make(map[int32]*PlayerLingTongOtherObject)
		m.otherMap[typ] = otherMap
	}
	pwo, exist = otherMap[seqId]
	if exist {
		return
	}
	pwo = NewPlayerLingTongOtherObject(m.p)
	now := global.GetGame().GetTimeService().Now()
	id, _ := idutil.GetId()
	pwo.id = id
	//生成id
	pwo.playerId = m.p.GetId()
	pwo.typ = typ
	pwo.seqId = seqId
	pwo.classType = m.classType
	pwo.level = 0
	pwo.upNum = 0
	pwo.upPro = 0
	pwo.createTime = now
	otherMap[seqId] = pwo
	pwo.SetModified()
	return
}

func (m *LingTongOtherContainer) IfCanUpStar(seqId int32) (*PlayerLingTongOtherObject, bool) {
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(m.classType, int(seqId))
	if lingTongDevTemplate == nil {
		return nil, false
	}

	obj, flag := m.IfLingTongSkinExist(seqId)
	if !flag {
		return nil, false
	}

	if lingTongDevTemplate.GetUpstarBeginId() == 0 {
		return nil, false
	}

	level := obj.level
	if level <= 0 {
		return obj, true
	}
	lingTongDevUpstarTemplate := lingTongDevTemplate.GetLingTongDevUpstarByLevel(level)
	if lingTongDevUpstarTemplate.GetNextId() != 0 {
		return obj, true
	}
	return nil, false
}

//是否已拥有该灵童养成皮肤
func (m *LingTongOtherContainer) IfLingTongSkinExist(seqId int32) (*PlayerLingTongOtherObject, bool) {
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(m.classType, int(seqId))
	if lingTongDevTemplate == nil {
		return nil, false
	}
	obj := m.getOtherObj(seqId)
	if obj == nil {
		return nil, false
	}
	return obj, true
}

func (m *PlayerLingTongDevDataManager) loadOther() (err error) {
	m.playerOtherMap = make(map[types.LingTongDevSysType]*LingTongOtherContainer)
	//加载玩家非进阶灵童养成信息
	otherList, err := dao.GetLingTongDevDao().GetLingTongOtherList(m.p.GetId())
	if err != nil {
		return
	}
	//非进阶灵童养成信息
	for _, otherObj := range otherList {
		pwo := NewPlayerLingTongOtherObject(m.p)
		pwo.FromEntity(otherObj)

		m.addOther(pwo)
	}
	return
}

func (m *PlayerLingTongDevDataManager) addOther(obj *PlayerLingTongOtherObject) {
	classType := obj.GetClassType()
	typ := obj.GetType()
	seqId := obj.GetSeqId()

	containerObj, ok := m.playerOtherMap[classType]
	if !ok {
		containerObj = NewLingTongOtherContainer(m.p, classType)
		m.playerOtherMap[classType] = containerObj
	}
	otherMap, ok := containerObj.otherMap[typ]
	if !ok {
		otherMap = make(map[int32]*PlayerLingTongOtherObject)
		containerObj.otherMap[typ] = otherMap
	}
	otherMap[seqId] = obj
}

//增加非进阶灵童养成
func (m *PlayerLingTongDevDataManager) initPlayerLingTongOtherObject(classType types.LingTongDevSysType, seqId int32) {
	containerObj := m.GetLingTongDevOtherMap(classType)
	if containerObj == nil {
		containerObj = NewLingTongOtherContainer(m.p, classType)
		m.playerOtherMap[classType] = containerObj
	}
	obj := containerObj.getOtherObj(seqId)
	if obj != nil {
		return
	} else {
		obj = containerObj.newPlayerLingTongOtherObject(seqId)
	}
	return
}

//获取玩家非进阶灵童养成对象
func (m *PlayerLingTongDevDataManager) GetLingTongDevOtherMap(classType types.LingTongDevSysType) *LingTongOtherContainer {
	containerObj, ok := m.playerOtherMap[classType]
	if !ok {
		return nil
	}
	return containerObj
}

//灵童养成皮肤升星
func (m *PlayerLingTongDevDataManager) Upstar(classType types.LingTongDevSysType, seqId int32, pro int32, sucess bool) (flag bool) {
	lingTongDevTemplate := lingtongdevtemplate.GetLingTongDevTemplateService().GetLingTongDevTemplate(classType, int(seqId))
	if lingTongDevTemplate == nil {
		return
	}
	containerObj := m.GetLingTongDevOtherMap(classType)
	if containerObj == nil {
		return
	}
	obj, flag := containerObj.IfCanUpStar(seqId)
	if !flag {
		return
	}
	return obj.UpSatr(pro, sucess)
}
