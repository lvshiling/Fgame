/*此类自动生成,请勿修改*/
package template

/*宝箱配置*/
type BoxTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//多次开启
	Times int32 `json:"times"`

	//宝箱类型
	Type int32 `json:"type"`

	//最小转数
	ZhuanshuMin int32 `json:"zhuanshu_min"`

	//最小等级
	LevelMin int32 `json:"level_min"`

	//掉落包id(开天男)
	DropId string `json:"drop_id"`

	//掉落包id(开天女)
	DropIdNv string `json:"drop_id_nv"`

	//开启宝箱消耗的物品id(开天)
	UseItemId string `json:"use_item_id"`

	//开启宝箱消耗的物品数量(开天)
	UseItemAmount string `json:"use_item_amount"`

	//掉落包id(奕剑男)
	DropId2 string `json:"drop_id2"`

	//掉落包id(奕剑女)
	DropId2Nv string `json:"drop_id2_nv"`

	//开启宝箱消耗的物品id(奕剑)
	UseItemId2 string `json:"use_item_id2"`

	//开启宝箱消耗的物品数量(奕剑)
	UseItemAmount2 string `json:"use_item_amount2"`

	//掉落包id(破月)
	DropId3 string `json:"drop_id3"`

	//掉落包id(破月女)
	DropId3Nv string `json:"drop_id3_nv"`

	//开启宝箱消耗的物品id(破月)
	UseItemId3 string `json:"use_item_id3"`

	//开启宝箱消耗的物品数量(破月)
	UseItemAmount3 string `json:"use_item_amount3"`

	//自由选择物品种类
	FixationItemNum int32 `json:"fixation_item_num"`

	//开启宝箱消耗银两
	UseSilver int32 `json:"use_silver"`

	//开启宝箱消耗元宝
	UseGold int32 `json:"use_gold"`

	//开启宝箱消耗绑元
	UseBindgold int32 `json:"use_bindgold"`
}
