/*此类自动生成,请勿修改*/
package template

/*仙桃大会配置*/
type XianTaoTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int `json:"next_id"`

	//仙桃类型
	Typ int32 `json:"type"`

	//提交的物品数量区间
	XianTaoCount string `json:"xiantao_count"`

	//提交获得奖励经验值
	RewExp int32 `json:"rew_xp"`

	//提交获得奖励经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//提交获得奖励银两
	RewSilver int32 `json:"rew_silver"`

	//提交获得奖励绑元
	RewBindGold int32 `json:"rew_bind_gold"`

	//提交获得奖励元宝
	RewGold int32 `json:"rew_gold"`

	//提交获得奖励物品id
	RewItemId string `json:"rew_item_id"`

	//提交获得奖励物品数量
	RewItemCount string `json:"rew_item_count"`
}
