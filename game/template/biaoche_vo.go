/*此类自动生成,请勿修改*/
package template

/*镖车模板配置*/
type BiaocheTemplateVO struct {

	//id
	Id int `json:"id"`

	//镖车类型
	Type int32 `json:"type"`

	//镖车id
	BiaocheId int32 `json:"biaoche_id"`

	//镖车起点NPC
	BiologyId int32 `json:"biology_id"`

	//押镖所需银两
	BiaocheSilver int32 `json:"biaoche_silver"`

	//押镖所需元宝
	BiaocheGold int32 `json:"biaoche_gold"`

	//押镖成功奖励银两
	BiaocheAwardSilver int32 `json:"biaoche_award_silver"`

	//押镖成功奖励元宝
	BiaocheAwardGold int32 `json:"biaoche_award_gold"`

	//押镖成功获得物品id
	BiaocheAwardItemId string `json:"biaoche_award_item_id"`

	//押镖成功获得物品数量
	BiaocheAwardItemCount string `json:"biaoche_award_item_count"`

	//劫镖获得银两
	JiebiaoAwardSilver int32 `json:"jiebiao_award_silver"`

	//劫镖获得元宝
	JiebiaoAwardGold int32 `json:"jiebiao_award_gold"`

	//劫镖成功获得的物品id
	JiebiaoAwardItemId string `json:"jiebiao_award_item_id"`

	//劫镖成功获得的物品数量
	JiebiaoAwardItemCount string `json:"jiebiao_award_item_count"`

	//押镖失败获得银两
	BiaocheLoseSilver int32 `json:"biaoche_lose_silver"`

	//押镖失败获得元宝
	BiaocheLoseGold int32 `json:"biaoche_lose_gold"`

	//押镖失败获得的物品id
	BiaocheLoseItemId string `json:"biaoche_lose_item_id"`

	//押镖失败获得的物品数量
	BiaocheLoseItemCount string `json:"biaoche_lose_item_count"`
}
