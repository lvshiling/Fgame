/*此类自动生成,请勿修改*/
package template

/*宝宝怀孕模板配置*/
type BabyPregnantTemplateVO struct {

	//id
	Id int `json:"id"`

	//补品值
	BupinQujian string `json:"bupin_qujian"`

	//品质
	Quality int32 `json:"quality"`

	//女宝宝名称
	NameNv string `json:"name_nv"`

	//男宝宝名称
	NameNan string `json:"name_nan"`

	//属性二进制
	AttrType int32 `json:"attr_type"`

	//属性倍数区间
	DanBeiQujian string `json:"danbei_qujian"`

	//天赋数量
	TalentCount int32 `json:"talent_count"`

	//天赋池id
	TalentBeginId int32 `json:"talent_begin_id"`

	//成长值
	GrowthNum int32 `json:"growth_num"`

	//读书最高等级
	LevelMax int32 `json:"level_max"`

	//属性提供比率
	AttrShareRate int32 `json:"attr_share_rate"`

	//宝宝出生概率
	Rate int64 `json:"rate"`
}
