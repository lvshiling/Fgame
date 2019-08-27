/*此类自动生成,请勿修改*/
package template

/*灵池争夺配置*/
type OneArenaTemplateVO struct {

	//id
	Id int `json:"id"`

	//下一个id
	NextId int32 `json:"next_id"`

	//灵池名字
	Name string `json:"name"`

	//灵池等级
	Level int32 `json:"level"`

	//灵池位置
	PosId int32 `json:"pos_id"`

	//战斗场景地图id
	MapId int32 `json:"map_id"`

	//灵池守护者关联生物表id
	BiologyId int32 `json:"biology_id"`

	//掉落id
	DropId string `json:"drop_id"`

	//奖励发送时间
	RefreshTime int32 `json:"refresh_time"`

	//战斗失败冷却时间
	CoolTime int32 `json:"cool_time"`

	//奖励预览
	RewItemId string `json:"rew_item_id"`

	//灵池守卫坐标x
	PosX float64 `json:"pos_x"`

	//灵池守卫坐标y
	PosY float64 `json:"pos_y"`

	//灵池守卫坐标z
	PosZ float64 `json:"pos_z"`
}
