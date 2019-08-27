/*此类自动生成,请勿修改*/
package template

/*衣橱模板配置*/
type YiChuTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//套装类型
	Type int32 `json:"type"`

	//套装类型
	SubType int32 `json:"sub_type"`

	//件数
	Number int32 `json:"number"`

	//食丹上限
	ShiDanLimit int32 `json:"shidan_limit"`

	//生命上限
	Hp int64 `json:"hp"`

	//攻击值
	Attack int64 `json:"attack"`

	//防御值
	Defence int64 `json:"defence"`

	//生命万分比
	HpPercent int64 `json:"hp_percent"`

	//攻击万分比
	AttackPercent int64 `json:"attack_percent"`

	//防御万分比
	DefPercent int64 `json:"def_percent"`

	//技能id
	SkillId int32 `json:"skill_id"`

	//技能id2
	SkillId2 int32 `json:"skill_id2"`
}
