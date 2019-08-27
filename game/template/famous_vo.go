/*此类自动生成,请勿修改*/
package template

/*名人培养配置*/
type FamousTemplateVo struct {

	//id
	Id int `json:"id"`

	//活动id
	GroupId int32 `json:"group"`

	//名人名称
	Name string `json:"name"`

	//消耗物品id
	UseItemId string `json:"use_item_id"`

	//消耗物品数量
	UseItemCount string `json:"use_item_count"`

	//增加的好感值
	ItemIncreaseFavorable string `json:"item_increase_favorable"`

	//喂养次数限制
	UseLimit string `json:"use_limit"`

	//喂养奖励掉落包
	DropId string `json:"drop_id"`
}
