/*此类自动生成,请勿修改*/
package template

/*时装升星配置*/
type FashionUpstarTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//星级
	Level int32 `json:"level"`

	//名字
	Name string `json:"name"`

	//升星需要消耗的物品id(开天男)
	UpstarItemId int32 `json:"upstar_item_id"`

	//升星需要消耗的物品id(开天女)
	UpstarItemIdNv int32 `json:"upstar_item_id_nv"`

	//升星需要消耗的物品id(奕剑男)
	UpstarItemId2 int32 `json:"upstar_item_id2"`

	//升星需要消耗的物品id(奕剑女)
	UpstarItemId2Nv int32 `json:"upstar_item_id2_nv"`

	//升星需要消耗的物品id(破月男)
	UpstarItemId3 int32 `json:"upstar_item_id3"`

	//升星需要消耗的物品id(破月女)
	UpstarItemId3Nv int32 `json:"upstar_item_id3_nv"`

	//升星消耗需要的物品数量
	UpstarItemCount int32 `json:"upstar_item_count"`

	//升星成功几率（万分比）
	UpstarRate int32 `json:"upstar_rate"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//每次随机加的最小祝福
	AddMin int32 `json:"add_min"`

	//每次随机加的最大祝福
	AddMax int32 `json:"add_max"`

	//每次随机加的最大祝福
	ZhufuMax int32 `json:"zhufu_max"`

	//生命加成（固定值）
	Hp int32 `json:"hp"`

	//攻击加成（固定）
	Attack int32 `json:"attack"`

	//防御加成（固定值）
	Defence int32 `json:"defence"`

	//装备基础全属性万分比
	EquipPercent int32 `json:"equip_percent"`
}
