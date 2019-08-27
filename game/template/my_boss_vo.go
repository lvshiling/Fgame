/*此类自动生成,请勿修改*/
package template

/*个人boss配置*/
type MyBossTemplateVO struct {

	//id
	Id int `json:"id"`

	//boss类型
	Type int32 `json:"type"`

	//vip等级限制
	NeedVipLevel int32 `json:"need_vip_level"`

	//生物id
	BiologyId int32 `json:"biology_id"`

	//地图id
	MapId int32 `json:"map_id"`

	//免费次数
	FreeTimes int32 `json:"free_times"`

	//总次数
	TimesCount int32 `json:"times_count"`

	//挑战所需物品
	UseItem int32 `json:"use_item"`

	//挑战所需物品数量
	UseCount int32 `json:"use_count"`
}
