/*此类自动生成,请勿修改*/
package template

/*装备宝库掉落配置*/
type EquipBaoKuTemplateVO struct {

	//id
	Id int `json:"id"`

	//后续id
	NextId int32 `json:"next_id"`

	//宝库类型
	Type int32 `json:"type"`

	//最小转数
	ZhuanshuMin int32 `json:"zhuanshu_min"`

	//最小等级
	LevelMin int32 `json:"level_min"`

	//掉落包id
	DropId int32 `json:"drop_id"`

	//一定次数必定获取的掉落包
	MustGet1 string `json:"must_get1"`

	//参与次数
	MustAmount1 string `json:"must_amount1"`

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

	//探寻一次获得的积分
	GiftJiFen int32 `json:"get_jifen"`

	//探寻一次获得的幸运值
	GiftXingYunZhi int32 `json:"get_xingyunzhi"`

	//幸运宝箱需要的幸运值
	NeedXingYunZhi int32 `json:"need_xingyunzhi"`

	//幸运宝箱掉落id
	ScriptXingYun int32 `json:"script_xingyun"`
}
