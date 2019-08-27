/*此类自动生成,请勿修改*/
package template

/*资源找回配置*/
type FoundTemplateVO struct {

	//id
	Id int `json:"id"`

	//任务名称
	Name string `json:"name"`

	//任务类型
	Type int32 `json:"type"`

	//最低限制
	LevelMin int32 `json:"level_min"`

	//最高限制
	LevelMax int32 `json:"level_max"`

	//单次完美找回消耗的绑元
	FoundUsing int32 `json:"found_using"`

	//单次完美找回获得的银两数量
	FoundSilver int32 `json:"found_silver"`

	//单次完美找回获得的元宝数量
	FoundGold int32 `json:"found_gold"`

	//单次完美找回获得的绑元数量
	FoundBindgold int32 `json:"found_bindgold"`

	//单次完美找回获得的经验值
	FoundExp int32 `json:"found_exp"`

	//单次完美找回获得的经验点
	FoundExpPoint int32 `json:"found_exp_point"`

	//单次完美找回获得的物品id
	FoundItemId string `json:"found_item_id"`

	//单次完美找回获得的物品数量
	FoundItemAmount string `json:"found_item_amount"`

	//单次普通找回消耗的银两数量
	FoundUsingSilver int32 `json:"found_using_silver"`

	//单次普通找回获得的银两数量
	FoundSilver2 int32 `json:"found_silver2"`

	//单次普通找回获得的元宝数量
	FoundGold2 int32 `json:"found_gold2"`

	//单次普通找回获得的绑元数量
	FoundBindgold2 int32 `json:"found_bindgold2"`

	//单次普通找回获得的经验值
	FoundExp2 int32 `json:"found_exp2"`

	//单次普通找回获得的经验点
	FoundExpPoint2 int32 `json:"found_exp_point2"`

	//单次普通找回获得的物品id
	FoundItemId2 string `json:"found_item_id2"`

	//单次普通找回获得的物品数量
	FoundItemAmount2 string `json:"found_item_amount2"`

	//功能开启id
	OpenId int32 `json:"open_id"`
}
