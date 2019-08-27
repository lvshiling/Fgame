/*此类自动生成,请勿修改*/
package template

/*苍龙棋局掉落配置*/
type ChessTemplateVO struct {

	//id
	Id int `json:"id"`

	//棋局类型
	Type int32 `json:"type"`

	//棋局id
	ChessId int32 `json:"chess_id"`

	//掉落包id
	DropId int32 `json:"drop_id"`

	//随机到掉落包的概率
	Rate int32 `json:"rate"`

	//一定次数必定获取的掉落包
	MustGet1 int32 `json:"must_get1"`

	//破解次数
	MustAmount1 int32 `json:"must_amount1"`

	//一定次数必定获取的掉落包
	MustGet2 int32 `json:"must_get2"`

	//破解次数
	MustAmount2 int32 `json:"must_amount2"`

	//一定次数必定获取的掉落包
	MustGet3 int32 `json:"must_get3"`

	//破解次数
	MustAmount3 int32 `json:"must_amount3"`

	//一定次数必定获取的掉落包
	MustGet4 int32 `json:"must_get4"`

	//破解次数
	MustAmount4 int32 `json:"must_amount4"`

	//一定次数必定获取的掉落包
	MustGet5 int32 `json:"must_get5"`

	//破解次数
	MustAmount5 int32 `json:"must_amount5"`

	//一定次数必定获取的掉落包
	MustGet6 int32 `json:"must_get6"`

	//破解次数
	MustAmount6 int32 `json:"must_amount6"`

	//一定次数必定获取的掉落包
	MustGet7 int32 `json:"must_get7"`

	//破解次数
	MustAmount7 int32 `json:"must_amount7"`

	//一定次数必定获取的掉落包
	MustGet8 int32 `json:"must_get8"`

	//破解次数
	MustAmount8 int32 `json:"must_amount8"`

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

	//破解赠送id
	GiftItem int32 `json:"gift_item"`

	//破解赠送数量
	GiftItemCount int32 `json:"gift_item_count"`
}
