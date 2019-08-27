/*此类自动生成,请勿修改*/
package template

/*创世城模板配置*/
type ChuangShiCityTemplateVO struct {

	//id
	Id int `json:"id"`

	//阵营
	Camp int32 `json:"camp"`

	//城市类型
	Type int32 `json:"type"`

	//名字
	Name string `json:"name"`

	//索引
	SuoyinId int32 `json:"suoyin_id"`

	//相邻城池id
	NearCityId string `json:"near_city_id"`

	//阵营积分
	ZhenyingRewJifen int32 `json:"zhenying_rew_jifen"`

	//阵营奖励钻石
	ZhenyingRewZuanshi int32 `json:"zhenying_rew_zuanshi"`

	//个人积分
	PlayerRewJifen int32 `json:"player_rew_jifen"`

	//个人奖励钻石
	PlayerRewZuanshi int32 `json:"player_rew_zuanshi"`

	//城常量id
	CityConstantId int32 `json:"city_constant_id"`
}
