/*此类自动生成,请勿修改*/
package template

/*上古之灵灵纹配置*/
type ShangguzhilingLingwenTemplateVO struct {

	//id
	Id int `json:"id"`

	//灵兽类型
	Type int32 `json:"type"`

	//灵纹类型
	SubType int32 `json:"sub_type"`

	//名字
	Name string `json:"name"`

	//需要上古之灵的等级
	NeedSgzlLevel int32 `json:"need_sgzl_level"`

	//灵纹升级起始Id
	LingwenLevelBeginId int32 `json:"lingwen_level_begin_id"`

	//上古之灵灵纹升级使用的物品ID
	LingwenLevelUseItemId string `json:"lingwen_level_use_item_id"`
}
