package player

import (
	gameevent "fgame/fgame/game/event"
	"fgame/fgame/game/global"
	"fgame/fgame/game/itemskill/dao"
	itemskilleventtypes "fgame/fgame/game/itemskill/event/types"
	itemskilltemplate "fgame/fgame/game/itemskill/template"
	itemskilltypes "fgame/fgame/game/itemskill/types"
	"fgame/fgame/game/player"
	"fgame/fgame/game/player/types"
	"fgame/fgame/pkg/idutil"
)

//玩家物品技能管理器
type PlayerItemSkillDataManager struct {
	p player.Player
	//玩家物品技能map
	playerItemSkillObjectMap map[itemskilltypes.ItemSkillType]*PlayerItemSkillObject
}

func (pxfdm *PlayerItemSkillDataManager) Player() player.Player {
	return pxfdm.p
}

//加载
func (pxfdm *PlayerItemSkillDataManager) Load() (err error) {
	pxfdm.playerItemSkillObjectMap = make(map[itemskilltypes.ItemSkillType]*PlayerItemSkillObject)
	//加载玩家物品技能
	itemSkills, err := dao.GetItemSkillDao().GetItemSkillList(pxfdm.p.GetId())
	if err != nil {
		return
	}
	//物品技能信息
	for _, itemSkill := range itemSkills {
		pxfo := NewPlayerItemSkillObject(pxfdm.p)
		pxfo.FromEntity(itemSkill)
		pxfdm.addItemSkill(pxfo)
	}

	return nil
}

func (pxfdm *PlayerItemSkillDataManager) addItemSkill(obj *PlayerItemSkillObject) {
	typ := obj.Typ
	pxfdm.playerItemSkillObjectMap[typ] = obj
}

//加载后
func (pxfdm *PlayerItemSkillDataManager) AfterLoad() (err error) {
	return nil
}

//心跳
func (pxfdm *PlayerItemSkillDataManager) Heartbeat() {

}

func (pxfdm *PlayerItemSkillDataManager) GetItemSkillObjctByTyp(typ itemskilltypes.ItemSkillType) *PlayerItemSkillObject {
	if !typ.Valid() {
		return nil
	}
	itemSkill, exist := pxfdm.playerItemSkillObjectMap[typ]
	if !exist {
		return nil
	}
	return itemSkill
}

func (pxfdm *PlayerItemSkillDataManager) CreateItemSkillObjByRestitution(skillId int32) {
	skTemplate := itemskilltemplate.GetItemSkillTemplateService().GetTemplateBySkillId(skillId)
	if skTemplate == nil {
		return
	}

	itemSkill, exist := pxfdm.playerItemSkillObjectMap[skTemplate.GetType()]
	if exist {
		return
	}

	id, _ := idutil.GetId()
	now := global.GetGame().GetTimeService().Now()
	itemSkill = NewPlayerItemSkillObject(pxfdm.p)
	itemSkill.Id = id
	itemSkill.Typ = skTemplate.GetType()
	itemSkill.Level = skTemplate.Level
	itemSkill.CreateTime = now
	itemSkill.SetModified()

	pxfdm.addItemSkill(itemSkill)
	gameevent.Emit(itemskilleventtypes.EventTypeItemSkillActive, pxfdm.p, skTemplate.GetType())
	return
}

//获取玩家物品技能map
func (pxfdm *PlayerItemSkillDataManager) GetItemSkillAllMap() map[itemskilltypes.ItemSkillType]*PlayerItemSkillObject {
	return pxfdm.playerItemSkillObjectMap
}

//获取等级通过物品技能类型
func (pxfdm *PlayerItemSkillDataManager) GetItemSkillLevelByTyp(typ itemskilltypes.ItemSkillType) int32 {
	obj := pxfdm.GetItemSkillObjctByTyp(typ)
	if obj == nil {
		return 0
	}
	return obj.Level
}

//是否已激活
func (pxfdm *PlayerItemSkillDataManager) IfItemSkillExist(typ itemskilltypes.ItemSkillType) bool {
	obj := pxfdm.GetItemSkillObjctByTyp(typ)
	if obj == nil {
		return false
	}
	return true
}

//是否达到满级
func (pxfdm *PlayerItemSkillDataManager) ifFullLevel(typ itemskilltypes.ItemSkillType) bool {
	level := pxfdm.GetItemSkillLevelByTyp(typ)
	to := itemskilltemplate.GetItemSkillTemplateService().GetItemSkillTemplateByTypeAndLevel(typ, level)
	if to == nil {
		return true
	}
	if to.NextId == 0 {
		return true
	}

	return false
}

//能否升级
func (pxfdm *PlayerItemSkillDataManager) IfCanUpgrade(typ itemskilltypes.ItemSkillType) bool {
	flag := pxfdm.IfItemSkillExist(typ)
	if !flag {
		return false
	}
	flag = pxfdm.ifFullLevel(typ)
	if flag {
		return false
	}
	return true
}

//物品技能激活
func (pxfdm *PlayerItemSkillDataManager) ItemSkillActive(typ itemskilltypes.ItemSkillType, lev int32) (*PlayerItemSkillObject, bool) {
	if lev < 1 {
		return nil, false
	}
	flag := typ.Valid()
	if !flag {
		return nil, false
	}
	flag = pxfdm.IfItemSkillExist(typ)
	if flag {
		return nil, false
	}

	skTemplate := itemskilltemplate.GetItemSkillTemplateService().GetItemSkillTemplateByTypeAndLevel(typ, lev)
	if skTemplate == nil {
		return nil, false
	}
	id, _ := idutil.GetId()

	now := global.GetGame().GetTimeService().Now()
	pxfo := NewPlayerItemSkillObject(pxfdm.p)
	pxfo.Id = id
	pxfo.Typ = typ
	pxfo.Level = skTemplate.Level
	pxfo.CreateTime = now
	pxfo.SetModified()

	pxfdm.addItemSkill(pxfo)
	gameevent.Emit(itemskilleventtypes.EventTypeItemSkillActive, pxfdm.p, typ)
	return pxfo, true
}

//物品技能升级
func (pxfdm *PlayerItemSkillDataManager) Upgrade(typ itemskilltypes.ItemSkillType, addLev int32) (*PlayerItemSkillObject, bool) {
	if addLev < 1 {
		return nil, false
	}
	flag := pxfdm.IfCanUpgrade(typ)
	if !flag {
		return nil, false
	}

	obj := pxfdm.GetItemSkillObjctByTyp(typ)
	if obj == nil {
		return nil, false
	}

	skTemplate := itemskilltemplate.GetItemSkillTemplateService().GetItemSkillTemplateByTypeAndLevel(typ, obj.Level+addLev)
	if skTemplate == nil {
		return nil, false
	}

	oldLev := obj.Level
	now := global.GetGame().GetTimeService().Now()
	obj.UpdateTime = now
	obj.Level = skTemplate.Level
	obj.SetModified()
	data := itemskilleventtypes.CreateXianFuChallengeEventData(typ, oldLev)
	gameevent.Emit(itemskilleventtypes.EventTypeItemSkillUpgrade, pxfdm.p, data)
	return obj, true
}

func CreatePlayerItemSkillDataManager(p player.Player) player.PlayerDataManager {
	pxfdm := &PlayerItemSkillDataManager{}
	pxfdm.p = p
	return pxfdm
}

func init() {
	player.RegisterPlayerDataManager(types.PlayerItemSkillDataManagerType, player.PlayerDataManagerFactoryFunc(CreatePlayerItemSkillDataManager))
}
