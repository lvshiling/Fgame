/*此类自动生成,请勿修改*/
package template

/*时装配置*/
type FashionTemplateVO struct {

	//id
	Id int `json:"id"`

	//时装名称-开天男
	Name11 string `json:"name_1_1"`

	//时装名称-开天女
	Name12 string `json:"name_1_2"`

	//时装名称-奕剑男
	Name21 string `json:"name_2_1"`

	//时装名称-奕剑女
	Name22 string `json:"name_2_2"`

	//时装名称-破月男
	Name31 string `json:"name_3_1"`

	//时装名称-破月女
	Name32 string `json:"name_3_2"`

	//时装类型
	Type int32 `json:"type"`

	//排序
	Pos int32 `json:"pos"`

	//需要物品id(开天男)
	NeedItemId int32 `json:"need_item_id"`

	//需要物品id(开天女)
	NeedItemIdNv int32 `json:"need_item_id_nv"`

	//需要物品id(奕剑男)
	NeedItemId2 int32 `json:"need_item_id2"`

	//需要物品id(奕剑女)
	NeedItemId2Nv int32 `json:"need_item_id2_nv"`

	//需要物品id(破月男)
	NeedItemId3 int32 `json:"need_item_id3"`

	//需要物品id(破月女)
	NeedItemId3Nv int32 `json:"need_item_id3_nv"`

	//需要物品数量
	NeedItemCount int32 `json:"need_item_count"`

	//有效性
	Time int64 `json:"time"`

	//属性
	Attr int32 `json:"attr"`

	//时装升星起始ID
	FashionUpgradeBeginId int32 `json:"fashion_upgrade_begin_id"`

	//获取途径
	Access string `json:"access"`

	//开天男
	ModelId1_1 int32 `json:"model_id_1_1"`

	//开天女
	ModelId1_2 int32 `json:"model_id_1_2"`

	//奕剑男
	ModelId2_1 int32 `json:"model_id_2_1"`

	//奕剑女
	ModelId2_2 int32 `json:"model_id_2_2"`

	//破月男
	ModelId3_1 int32 `json:"model_id_3_1"`

	//破月女
	ModelId3_2 int32 `json:"model_id_3_2"`
}
