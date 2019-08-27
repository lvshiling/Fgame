/*此类自动生成,请勿修改*/
package template

/*转生配置*/
type ZhuanShengTemplateVO struct {

	//id
	Id int `json:"id"`

	//转生所需玩家转数
	NeedZhuanshu int32 `json:"need_zhuanshu"`

	//转生所需玩家等级
	NeedLevel int32 `json:"need_level"`

	//转生所需玩家飞升等级
	NeedFeisheng int32 `json:"need_feisheng"`

	//转生所需装备件数
	NeedEquipCount int32 `json:"need_equip_count"`

	//转生所需玩家装备转数
	NeedEquipZhuanshu int32 `json:"need_equip_zhuanshu"`

	//转生所需玩家装备等级
	NeedEquipLevel int32 `json:"need_equip_level"`

	//转生所需玩家装备强化等级
	NeedEquipStreng int32 `json:"need_equip_streng"`

	//新增转生所需装备品质描述
	NeedEquipQuality int32 `json:"need_equip_quality"`

	//生命加成（固定值）
	Hp int32 `json:"hp"`

	//攻击加成（固定值）
	Attack int32 `json:"attack"`

	//防御加成（固定值）
	Defence int32 `json:"defence"`
}
