/*此类自动生成,请勿修改*/
package template

/*推送常量模板配置*/
type NoticeConstantTemplateVO struct {

	//id
	Id int `json:"id"`

	//好友信息弹窗CD
	FriendTime int64 `json:"friend_time"`

	//仇人信息弹窗CD
	ChouRenTime int64 `json:"chouren_time"`

	//被狗咬损失银两
	ChouRenSunShiSilver int32 `json:"chouren_sunshi_silver"`

	//购买防狗所需元宝
	BaoHuFei int32 `json:"baohufei"`

	//防狗持续时间
	BaoHuTime int32 `json:"baohu_time"`

	//虚假好友间隔
	TianjiaJiarenTime int32 `json:"tianjia_jiaren_time"`

	//同意赠送文本
	SongliTongyi string `json:"songli_tongyi"`

	//拒绝赠送文本
	SongliJujue string `json:"songli_jujue"`

	//索取成功文本
	ShouliTongyi string `json:"shouli_tongyi"`

	//索取失败文本
	ShouliJujue string `json:"shouli_jujue"`

	//被砸蛋收花的次数
	ShouliLimitCount int32 `json:"shouli_limit_count"`
}
