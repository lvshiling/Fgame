/*此类自动生成,请勿修改*/
package template

/*玉玺常量配置*/
type YuXiConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//占有获胜时间
	WinTime int64 `json:"win_time"`

	//玉玺采集物id
	YuxiBiologyId int32 `json:"yuxi_biology_id"`

	//获胜奖励物品id
	MailItemId string `json:"mail_item_id"`

	//获胜奖励物品数量
	MailItemCount string `json:"mail_item_count"`

	//获胜每日奖励物品id
	WinDayItemId string `json:"win_day_item_id"`

	//获胜每日奖励物品数量
	WinDayItemCount string `json:"win_day_item_count"`

	//玉玺buffId
	BuffId int32 `json:"buff_id"`

	//千里救援cd
	RescueCd int64 `json:"qianlijiuyuan_cd"`

	//地图x坐标
	PosX float64 `json:"pos_x"`

	//地图y坐标
	PosY float64 `json:"pos_y"`

	//地图z坐标
	PosZ float64 `json:"pos_z"`

	//地图id
	MapId int32 `json:"map_id"`

	//玉玺争夺战获胜仙盟盟主雕像 地图x坐标
	ModelPosX float64 `json:"pos_x1"`

	//玉玺争夺战获胜仙盟盟主雕像 地图y坐标
	ModelPosY float64 `json:"pos_y1"`

	//玉玺争夺战获胜仙盟盟主雕像 地图z坐标
	ModelPosZ float64 `json:"pos_z1"`

	//玉玺争夺战获胜仙盟盟主配偶雕像 地图x坐标
	ModelCouplePosX float64 `json:"pos_x2"`

	//玉玺争夺战获胜仙盟盟主配偶雕像 地图y坐标
	ModelCouplePosY float64 `json:"pos_y2"`

	//玉玺争夺战获胜仙盟盟主配偶雕像 地图z坐标
	ModelCouplePosZ float64 `json:"pos_z2"`

	//雕像地图id
	ModelMapId int32 `json:"map_id1"`
}
