/*此类自动生成,请勿修改*/
package template

/*创世之战常量模板配置*/
type ChuangShiYuGaoTemplateVO struct {

	//id
	Id int `json:"id"`

	//预告开始时间
	StartTime string `json:"yugao_star_time"`

	//预告持续时间
	DaoJiShiDay int32 `json:"daojishi_day"`

	//增加虚假人数间隔时间
	JiaRenTime int64 `json:"jiaren_time"`

	//增加的人数
	JiaRenCount int32 `json:"jiaren_count"`

	//报名活动的物品id
	BaoMingGet string `json:"baoming_get"`

	//报名活动的物品数量
	BaoMingGetCount string `json:"baoming_get_count"`
}
