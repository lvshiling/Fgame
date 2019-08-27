/*此类自动生成,请勿修改*/
package template

/*幸运抽奖配置*/
type LuckyDrewTemplateVO struct {

	//id
	Id int `json:"id"`

	//抽奖类型
	Type int32 `json:"type"`

	//抽奖id
	ChouJiangId int32 `json:"choujiang_id"`

	//掉落包id
	DropId int32 `json:"drop_id"`

	//随机到掉落包的概率
	Rate int32 `json:"rate"`

	//后台规则次数条件
	RewCount int32 `json:"rew_count"`

	//倍数
	RewTimes1 int32 `json:"rew_times1"`

	//概率
	Percent1 int32 `json:"percent1"`

	//倍数
	RewTimes2 int32 `json:"rew_times2"`

	//概率
	Percent2 int32 `json:"percent2"`

	//一定次数必定获取的掉落包
	MustGet1 int32 `json:"must_get1"`

	//参与次数
	MustAmount1 int32 `json:"must_amount1"`

	//一定次数必定获取的掉落包
	MustGet2 int32 `json:"must_get2"`

	//参与次数
	MustAmount2 int32 `json:"must_amount2"`

	//一定次数必定获取的掉落包
	MustGet3 int32 `json:"must_get3"`

	//参与次数
	MustAmount3 int32 `json:"must_amount3"`

	//一定次数必定获取的掉落包
	MustGet4 int32 `json:"must_get4"`

	//参与次数
	MustAmount4 int32 `json:"must_amount4"`

	//一定次数必定获取的掉落包
	MustGet5 int32 `json:"must_get5"`

	//参与次数
	MustAmount5 int32 `json:"must_amount5"`

	//一定次数必定获取的掉落包
	MustGet6 int32 `json:"must_get6"`

	//参与次数
	MustAmount6 int32 `json:"must_amount6"`

	//一定次数必定获取的掉落包
	MustGet7 int32 `json:"must_get7"`

	//参与次数
	MustAmount7 int32 `json:"must_amount7"`

	//一定次数必定获取的掉落包
	MustGet8 int32 `json:"must_get8"`

	//参与次数
	MustAmount8 int32 `json:"must_amount8"`

	//抽奖等级类型
	Level int32 `json:"level"`

	//消耗id
	ItemId string `json:"item_id"`

	//消耗数量
	ItemCount string `json:"item_count"`

	//消耗次数
	CostTimes int32 `json:"cost_times"`

	//额外奖励id
	GiveItemId string `json:"give_item_id"`

	//额外奖励数量
	GiveItemCount string `json:"give_item_count"`

	//称号id
	TitleId int32 `json:"titleId"`
}
