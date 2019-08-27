/*此类自动生成,请勿修改*/
package template

/*特戒配置*/
type RingTemplateVO struct {

	//物品id
	Id int `json:"id"`

	//特戒类型
	Type int32 `json:"type"`

	//进阶起始id
	AdvanceBeginId int32 `json:"jinjie_begin_id"`

	//强化起始id
	StrengthenBeginId int32 `json:"strengthen_begin_id"`

	//净灵起始id
	JingLingBeginId int32 `json:"jingling_begin_id"`

	//特戒融合等级
	Level int32 `json:"level"`

	//该类型增加的生命
	Hp int64 `json:"hp"`

	//该类型增加的攻击
	Attack int64 `json:"attack"`

	//该类型增加的防御
	Defence int64 `json:"defence"`

	//特戒技能
	SkillId int32 `json:"skill_id"`

	//融合合成id
	FuseSynthesisId int32 `json:"ronghe_synthesis_id"`
}
