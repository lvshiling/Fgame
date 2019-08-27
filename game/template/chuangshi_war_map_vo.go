/*此类自动生成,请勿修改*/
package template

/*创世城战配置*/
type ChuangShiWarMapTemplateVO struct {

	//id
	Id int `json:"id"`

	//地图id
	MapId int32 `json:"map_id"`

	//采集旗子所需要的时间(毫秒)
	OccupyFlagTime int32 `json:"occupy_flag_time"`

	//限制区域1
	FirstXianzhi int32 `json:"first_xianzhi"`

	//限制区域2
	SecondXianzhi int32 `json:"second_xianzhi"`

	//区域1固定点
	LahuiPos1 string `json:"lahui_pos1"`

	//区域2固定点
	LahuiPos2 string `json:"lahui_pos2"`

	//保护罩地图x坐标
	ProtectPosX float64 `json:"pos_x1"`

	//保护罩地图y坐标
	ProtectPosY float64 `json:"pos_y1"`

	//保护罩地图z坐标
	ProtectPosZ float64 `json:"pos_z1"`

	//保护罩id
	ProtectId int32 `json:"zhaozi_id"`

	//保护罩生成时间
	ProtectRebornTime int64 `json:"zhaozi_reborn_time"`

	//玉玺采集物地图x坐标
	YuXiPosX float64 `json:"pos_x2"`

	//玉玺采集物地图y坐标
	YuXiPosY float64 `json:"pos_y2"`

	//玉玺采集物地图z坐标
	YuXiPosZ float64 `json:"pos_z2"`

	//玉玺采集物id
	YuxiId int32 `json:"yuxi_id"`

	//防护罩限制区域
	ProtectQuYuPos int32 `json:"zhaozi_quyu_pos"`

	//防护罩驱逐固定点
	ProtectLaHuiPos string `json:"zhaozi_lahui_pos"`
}
