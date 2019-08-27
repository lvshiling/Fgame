/*此类自动生成,请勿修改*/
package template

/*救援引导副本配置*/
type GuideReplicaRescureTemplateVO struct {

	//id
	Id int `json:"id"`

	//小医仙生物id
	RescureBiologyId int32 `json:"biology_id"`

	//采集物生物id
	CollectBiologyId int32 `json:"biology_id2"`

	//奖励掉落
	DropId string `json:"drop_id"`

	//采集buffId
	BuffId int32 `json:"buff_id"`
}
