/*此类自动生成,请勿修改*/
package template

/*渠道模板配置*/
type QuDaoTemplateVO struct {

	//id
	Id int `json:"id"`

	//平台类型
	Type int32 `json:"type"`

	//平台名称
	Name string `json:"name"`

	//仙盟版本
	UnionVersion int32 `json:"union_type"`

	//比武大会显示平台名
	PlatName string `json:"plat_name"`

	//战力榜前几的玩家上下线有传音公告显示
	ChuanYinNum int32 `json:"chuanyin_num"`

	//是否需要添加好友才能聊天
	IsFriend int32 `json:"is_friend"`
}
