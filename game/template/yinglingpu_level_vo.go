/*此类自动生成,请勿修改*/
package template

/*英灵谱配置升级*/
type YinglingpuLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一部位id
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//升级需要消耗的物品
	UseItemId string `json:"use_item_id"`

	//升级需要消耗的物品数量
	UseItemCount string `json:"use_item_count"`

	//升级增加的生命
	Hp int32 `json:"hp"`

	//升级增加的攻击
	Attack int32 `json:"attack"`

	//升级增加的防御
	Defence int32 `json:"defence"`
}
