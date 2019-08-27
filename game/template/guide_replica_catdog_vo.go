/*此类自动生成,请勿修改*/
package template

/*猫狗引导副本配置*/
type GuideReplicaCatDogTemplateVO struct {

	//id
	Id int `json:"id"`

	//猫生物id
	CatBiologyId int32 `json:"biology_id"`

	//狗生物id
	DogBiologyId int32 `json:"biology_id2"`

	//猫奖励掉落
	CatDropId string `json:"drop_id"`

	//狗奖励掉落
	DogDropId string `json:"drop_id2"`

	//默认奖励掉落
	DefaultDropId string `json:"drop_id3"`
}
