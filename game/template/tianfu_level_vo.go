/*此类自动生成,请勿修改*/
package template

/*天赋等级配置*/
type TianFuLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//天赋等级
	Level int32 `json:"level"`

	//升级消耗的物品id
	UseItemId int32 `json:"use_item_id"`

	//升级消耗的物品数量
	UseItemCount int32 `json:"use_item_count"`

	//升级消耗的物品id
	UseItemIdYiJian int32 `json:"use_item_id_yijian"`

	//升级消耗的物品数量
	UseItemCountYiJian int32 `json:"use_item_count_yijian"`

	//升级消耗的物品id
	UseItemIdPoYue int32 `json:"use_item_id_poyue"`

	//升级消耗的物品数量
	UseItemCountPoYue int32 `json:"use_item_count_poyue"`

	//升阶成功率
	UpdateWfb int32 `json:"update_wfb"`

	//施法范围
	AreaType int32 `json:"area_type"`

	//技能作用半径
	AreaRadius int32 `json:"area_radius"`

	//施法角度
	AreaRange int32 `json:"area_range"`

	//作用对象
	SpecialTarget int32 `json:"special_target"`

	//特殊效果标识
	SpecialEffect int32 `json:"special_effect"`

	//特殊效果标识概率
	SpecialEffectRate int32 `json:"special_effect_rate"`

	//产生特殊效果的值
	SpecialEffectValue int32 `json:"special_effect_value"`

	//产生特殊效果的值
	SpecialEffectValue2 int32 `json:"special_effect_value2"`

	//产生特殊效果的值
	SpecialEffectValue3 int32 `json:"special_effect_value3"`

	//给予目标的buff
	AddStatus string `json:"add_status"`

	//增加该状态的概率
	AddStatusRate string `json:"add_status_rate"`

	//关联到动态buff等级表
	BuffDongTaiId string `json:"buff_dongtai_id"`
}
