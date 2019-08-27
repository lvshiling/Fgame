/*此类自动生成,请勿修改*/
package template

/*房子常量模板配置*/
type HouseConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//房子损坏CD
	BrokenCd int64 `json:"broken_cd"`

	//每日装修次数
	UplevLimitCount int32 `json:"uplev_limit_count"`

	//虚拟日志时间间隔
	JiaRiZhiTime string `json:"jiarizhi_time"`

	//总日志数量
	RiZhiCount int32 `json:"rizhi_count"`

	//初始房子类型
	FirstFangZiType int32 `json:"first_fangzi_type"`

	//房子最大数量
	FangZiCountMax int32 `json:"fangzi_count_max"`
}
