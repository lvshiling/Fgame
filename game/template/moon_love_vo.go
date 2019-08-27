/*此类自动生成,请勿修改*/
package template

/*月下情缘配置*/
type MoonloveTemplateVO struct {

	//id
	Id int `json:"id"`

	//最小等级
	MinLev int32 `json:"min_lev"`

	//最大等级
	MaxLev int32 `json:"max_lev"`

	//进入后多少毫秒获得奖励
	FristTiem int64 `json:"frist_tiem"`

	//多少毫秒发一次奖励
	RewTiem int64 `json:"rew_tiem"`

	//单次发放经验值
	RewExp int64 `json:"rew_exp"`

	//单次发放的经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//奖励银两
	RewSilver int64 `json:"rew_silver"`

	//奖励元宝
	RewGold int32 `json:"rew_gold"`

	//奖励绑定元宝
	RewBindGold int32 `json:"rew_bind_gold"`

	//奖励物品Id
	RewItemId string `json:"rew_item_id"`

	//奖励物品数量
	RewItemCount string `json:"rew_item_count"`

	//双人奖励倍数（物品不生效，向上取整）
	DoubleMan int32 `json:"double_man"`

	//采集次数限制
	CollectLimitCount int32 `json:"caiji_limit_count"`
}
