/*此类自动生成,请勿修改*/
package template

/*宝宝解锁天赋模板配置*/
type BabyUnlockTalentTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一级
	NextId int `json:"next_id"`

	//次数
	Times int32 `json:"times"`

	//消耗元宝
	UseGold int32 `json:"use_gold"`

	//锁定消耗元宝
	SuodingGold int32 `json:"suoding_gold"`
}
