/*此类自动生成,请勿修改*/
package template

/*神龙现世配置*/
type DragonTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一培养阶段
	NextId int32 `json:"next_id"`

	//神龙喂食所需的物品id
	DragonItemId string `json:"dragon_item_id"`

	//神龙喂食所需的物品数量
	DragonItemAmount string `json:"dragon_item_amount"`

	//龙神各阶段模型id
	DragonModelId int32 `json:"dragon_model_id"`

	//技能id
	DragonSkill int32 `json:"dragon_skill"`

	//坐骑id
	DragonMount int32 `json:"dragon_mount"`

	//材料产出地
	DragonChandi string `json:"dragon_chandi"`

	//坐骑幻化卡id
	ItemId int32 `json:"item_id"`
}
