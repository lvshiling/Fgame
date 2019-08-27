/*此类自动生成,请勿修改*/
package template

/*外域boss配置*/
type OutlandBossTemplateVO struct {

	//id
	Id int `json:"id"`

	//boss类型
	Type int32 `json:"type"`

	//生物id
	BiologyId int32 `json:"biology_id"`

	//地图id
	MapId int32 `json:"map_id"`

	//假掉落记录出现概率
	Rate int32 `json:"rate"`

	//战力推荐
	RecForce int64 `json:"rec_force"`
}
