/*此类自动生成,请勿修改*/
package template

/*竞技场pvp常量模板配置*/
type ArenapvpConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//是否优先匹配假人
	IsJiaren int32 `json:"is_jiaren"`

	//进入会场多少时间获得积分
	CunhuoFristTiem int32 `json:"cunhuo_frist_tiem"`

	//会场获得积分间隔
	CunhuoTime int32 `json:"cunhuo_time"`

	//会场获得积分数量
	CunhuoJifen int32 `json:"cunhuo_jifen"`

	//幸运奖开始时间
	XingYunFirstTime int64 `json:"xingyun_first_time"`

	//幸运奖循环时间
	XingYunTime int64 `json:"xingyun_time"`

	//幸运奖人数
	XingYunPlayerCount int32 `json:"xingyun_player_count"`

	//幸运奖假人概率
	XinyunRobotRate int32 `json:"xinyun_robot_rate"`

	//幸运奖物品id
	XingYunItemId int32 `json:"xingyun_item_id"`

	//死亡损失积分百分比
	JifenBekillPercent int32 `json:"jifen_bekill_percent"`

	//插入机器人条件
	ZhenshiPlayerRobot int32 `json:"zhenshi_player_robot"`

	//插入机器人间隔
	RobotAddTime int64 `json:"robot_add_time"`

	//机器人上限
	RobotMax int32 `json:"robot_max"`

	//机器人下限
	AttrRatioMin int64 `json:"attr_min"`

	//机器人上限
	AttrRatioMax int64 `json:"attr_max"`

	//门票价格
	RuchangUseBindgold int32 `json:"ruchang_use_bindgold"`
}
