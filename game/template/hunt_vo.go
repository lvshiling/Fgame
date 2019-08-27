/*此类自动生成,请勿修改*/
package template

/*寻宝配置*/
type HuntTemplateVO struct {

	//id
	Id int `json:"id"`

	//寻宝类型
	Type int32 `json:"type"`

	//掉落包id
	DropId int32 `json:"drop_id"`

	//一定次数必定获取的掉落包
	MustGet string `json:"must_get1"`

	//寻宝后台次数
	MustAmount string `json:"must_amount1"`

	//银两消耗
	SilverUse int32 `json:"silver_use"`

	//元宝消耗
	GoldUse int32 `json:"gold_use"`

	//绑元消耗
	BindGoldUse int32 `json:"bindgold_use"`

	//消耗物品id
	UseItemId int32 `json:"use_item_id"`

	//消耗物品数量
	UseItemCount int32 `json:"use_item_count"`
}
