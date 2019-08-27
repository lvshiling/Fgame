/*此类自动生成,请勿修改*/
package template

/*城战配置*/
type WarTemplateVO struct {

	//id
	Id int `json:"id"`

	//最大获得腰牌的上限
	PlateMax int32 `json:"plate_max"`

	//每次击杀获得腰牌的几率(万分比)
	PalteOdd int32 `json:"palte_odd"`

	//原地复活的最高次数
	RebornSitu int32 `json:"reborn_situ"`

	//采集旗子所需要的时间(毫秒)
	OccupyFlagTime int32 `json:"occupy_flag_time"`

	//击杀同一个玩家获得腰牌的间隔时间(毫秒)
	PlateTime int32 `json:"plate_time"`

	//掉落包
	DropId int32 `json:"drop_id"`

	//皇宫时间
	HuanggongTime int32 `json:"huanggong_time"`

	//进入后多少毫秒获得奖励
	FristTiem int64 `json:"frist_tiem"`

	//多少毫秒发一次奖励
	RewTiem int64 `json:"rew_tiem"`

	//单次发放经验值
	RewExp int64 `json:"rew_exp"`

	//单次发放的经验点
	RewExpPoint int32 `json:"rew_exp_point"`

	//奖励银两
	RewSilver int64 `json:"rew_silver"`

	//地图id
	MapId int32 `json:"map_id"`

	//单次发放经验值
	RewExpSafeArea int64 `json:"rew_exp_anquanqu"`

	//单次发放的经验点
	RewExpPointSafeArea int32 `json:"rew_exp_point_anquanqu"`

	//奖励银两
	RewSilverSafeArea int64 `json:"rew_silver_anquanqu"`

	//限制区域1
	FirstXianzhi int32 `json:"first_xianzhi"`

	//限制区域2
	SecondXianzhi int32 `json:"second_xianzhi"`

	//限制区域3
	ThirdXianzhi int32 `json:"third_xianzhi"`

	//限制区域4
	FourXianzhi int32 `json:"four_xianzhi"`

	//区域1固定点
	LahuiPos1 string `json:"lahui_pos1"`

	//区域2固定点
	LahuiPos2 string `json:"lahui_pos2"`

	//区域3固定点
	LahuiPos3 string `json:"lahui_pos3"`

	//击杀获得积分
	KillJiFen int32 `json:"kill_jifen"`

	//进入后多少毫秒获得奖励
	FristJiFenTime int64 `json:"frist_jifen_time"`

	//多少毫秒发一次奖励
	RewJiFenTime int64 `json:"rew_jifen_time"`

	//单次发放经验值
	RewJiFen int32 `json:"rew_jifen"`

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

	//防护罩无敌buff
	ProtectBuffId int32 `json:"wudi_buff_id"`

	//防护罩限制区域
	ProtectQuYuPos int32 `json:"zhaozi_quyu_pos"`

	//防护罩驱逐固定点
	ProtectLaHuiPos string `json:"zhaozi_lahui_pos"`
}
