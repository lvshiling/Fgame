/*此类自动生成,请勿修改*/
package model

import logserverlog "fgame/fgame/logserver/log"

func init() {
	logserverlog.RegisterLogMsg((*PlayerJieYiDaoJu)(nil))
}

/*结义道具改变*/
type PlayerJieYiDaoJu struct {
	PlayerLogMsg `bson:",inline"`

	//之前道具类型
	BeforeDaoJu string `json:"beforeDaoJu"`

	//当前道具类型
	CurDaoJu string `json:"curDaoJu"`

	//进阶原因编号
	Reason int32 `json:"reason"`

	//进阶原因
	ReasonText string `json:"reasonText"`
}

func (c *PlayerJieYiDaoJu) LogName() string {
	return "player_jieyi_daoju"
}
