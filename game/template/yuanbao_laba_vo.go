/*此类自动生成,请勿修改*/
package template

/*元宝拉霸配置*/
type GoldLaBaTemplateVo struct {

	//id
	Id int `json:"id"`

	//下一级id
	NextId int `json:"next_id"`

	//关联运营活动id
	GroupId int32 `json:"type"`

	//次数
	Times int32 `json:"times"`

	//所需玩家充值金额
	InvestmentRecharge int32 `json:"investment_recharge"`

	//所需玩家元宝
	Investment int32 `json:"investment"`

	//返还元宝下限1
	ReturnMin1 int32 `json:"return_min1"`

	//返还元宝上限1
	ReturnMax1 int32 `json:"return_max1"`

	//概率1
	Percent1 int32 `json:"percent1"`

	//返还元宝下限2
	ReturnMin2 int32 `json:"return_min2"`

	//返还元宝上限2
	ReturnMax2 int32 `json:"return_max2"`

	//概率2
	Percent2 int32 `json:"percent2"`
}
