package types

type SkillEventType string

const (
	//技能添加
	EventTypeSkillAdd SkillEventType = "skillAdd"
	//技能移除
	EventTypeSkillRemove SkillEventType = "skillRemove"
	//技能升级
	EventTypeSkillUpgrade SkillEventType = "SkillUpgrade"
	//使用技能
	EventTypeSkillUse SkillEventType = "SkillUse"
)

type SkillTianFuEventType string

const (
	//技能天机觉醒
	EventTypeSkillTianFuAwaken SkillTianFuEventType = "SkillTianFuAwaken"
	//技能天赋升级
	EventTypeSkillTianFuUpgrade SkillTianFuEventType = "SkillTianFuUpgrade"
)

type SkillTianFuUpgradeEventData struct {
	skillId  int32
	tianFuId int32
}

func CreateSkillTianFuUpgradeEventData(skillId int32, tianFuId int32) *SkillTianFuUpgradeEventData {
	data := &SkillTianFuUpgradeEventData{
		skillId:  skillId,
		tianFuId: tianFuId,
	}
	return data
}

func (d *SkillTianFuUpgradeEventData) GetSkillId() int32 {
	return d.skillId
}

func (d *SkillTianFuUpgradeEventData) GetTianFuId() int32 {
	return d.tianFuId
}
