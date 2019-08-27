/*此类自动生成,请勿修改*/
package template

/*绝学配置*/
type JueXueTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//绝学类型
	Type int32 `json:"type"`

	//绝学等级
	Level int32 `json:"level"`

	//是否可顿悟
	IsInsight int32 `json:"is_insight"`

	//名称
	Name string `json:"name"`

	//获得战力
	Power int32 `json:"power"`

	//技能描述
	Des string `json:"des"`

	//激活所需物品ID
	NeedItemId int32 `json:"need_item_id"`

	//激活所需物品数量
	NeedItemNum int32 `json:"need_item_num"`

	//关联技能ID
	Skill int32 `json:"skill"`
}
