/*此类自动生成,请勿修改*/
package template

/*上古之灵基础配置*/
type ShangguzhilingBaseTemplateVO struct {

	//id
	Id int `json:"id"`

	//类型
	Type int32 `json:"type"`

	//名字
	Name string `json:"name"`

	//所需的上古之灵ID
	NeedSgzlId int32 `json:"need_sgzl_id"`

	//所需的上古之灵等级
	NeedSgzlLevel int32 `json:"need_sgzl_level"`

	//上古之灵升级起始ID
	SgzlLevelBeginId int32 `json:"sgzl_level_begin_id"`

	//进阶起始Id
	JinjieBeginId int32 `json:"jinjie_begin_id"`

	//hp
	Hp int32 `json:"hp"`

	//攻击力
	Attack int32 `json:"attack"`

	//防御力
	Defence int32 `json:"defence"`

	//上古之灵升级使用的物品ID
	SgzlLevelUseItemId string `json:"sgzl_level_use_item_id"`
}
