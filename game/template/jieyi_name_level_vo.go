/*此类自动生成,请勿修改*/
package template

/*结义威名等级配置*/
type JieYiNameLevelTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级di
	NextId int32 `json:"next_id"`

	//等级
	Level int32 `json:"level"`

	//升级概率
	UpLevPercent int32 `json:"update_percent"`

	//所需银两
	UseMoney int32 `json:"use_money"`

	//所需声威值
	UseShengWei int32 `json:"use_shengwei"`

	//所需物品id
	UseItemId int32 `json:"use_item"`

	//所需物品数量
	UseItemCount int32 `json:"item_count"`

	//最小次数
	TimesMin int32 `json:"times_min"`

	//最大次数
	TimesMax int32 `json:"times_max"`

	//死亡掉落最小声威等级
	DeathMinLevel int32 `json:"shengwei_min"`

	//死亡掉落最大声威等级
	DeathMaxLevel int32 `json:"shengwei_max"`

	//死亡掉落声威等级概率
	DropPercent int32 `json:"shengwei_percent"`

	//掉在地上的声威值
	StarCount int32 `json:"star_count"`

	//该等级增加的生命
	Hp int64 `json:"hp"`

	//该等级增加的攻击
	Attack int64 `json:"attack"`

	//该等级增加的防御
	Defence int64 `json:"defence"`
}
