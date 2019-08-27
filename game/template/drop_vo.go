/*此类自动生成,请勿修改*/
package template

/*掉落配置*/
type DropTemplateVO struct {

	//id
	Id int `json:"id"`

	//掉落包id
	DropId int32 `json:"drop_id"`

	//掉落物品id
	ItemId int32 `json:"item_id"`

	//掉落概率,最大为20亿
	Rate int64 `json:"rate"`

	//掉落最小数量
	MinCount int32 `json:"min_count"`

	//掉落最大数量
	MaxCount int32 `json:"max_count"`

	//最小堆数
	MinStack int32 `json:"min_stack"`

	//最大堆数
	MaxStack int32 `json:"max_stack"`

	//金装最小强化等级
	GoldEquipMin int32 `json:"gold_equip_min"`

	//金装最大强化等级
	GoldEquipMax int32 `json:"gold_equip_max"`

	//金装升星等级池id
	StrengthenPoolId int32 `json:"strengthen_pool_id"`

	//关联金装附件属性表
	FristFujiaId int32 `json:"frist_fujia_id"`

	//是否绑定
	BindType int32 `json:"bind_type"`

	//存活时间
	ExistTime int32 `json:"exist_time"`

	//保护时间
	ProtectedTime int32 `json:"protected_time"`

	//物品失效时间
	FailTime int32 `json:"fail_time"`

	//是否支持再次掉落0-否1-是
	IsDropAgin int32 `json:"is_drop_agin"`

	//是否受到掉落压制
	IsDropSuppress int32 `json:"is_drop_suppress"`
}
