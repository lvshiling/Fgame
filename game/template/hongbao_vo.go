/*此类自动生成,请勿修改*/
package template

/*红包配置*/
type HongBaoTemplateVO struct {

	//id
	Id int `json:"id"`

	//红包类型
	Typ int32 `json:"type"`

	//领取红包需要的VIP最低等级与need_zhuanshu关系为或
	NeedVipLevel int32 `json:"need_vip_level"`

	//领取红包需要的转数与need_vip_level关系为或
	NeedZhuanShu int32 `json:"need_zhuanshu"`

	//奖励总量
	GeneralRew int32 `json:"general_rew"`

	//最佳手气最大值
	GoodProportionMax int32 `json:"good_proportion_max"`

	//最佳手气最小值
	GoodProportionMin int32 `json:"good_proportion_min"`

	//普通手气保底值
	ProportionMin int32 `json:"proportion_min"`

	//消耗的物品id
	ItemId int32 `json:"item_id"`

	//最小领取人数
	CountMin int32 `json:"hb_count_min"`

	//最大领取人数
	CountMax int32 `json:"hb_count_max"`
}
