package entity

//战力 MySql+redis
type PlayerForceData struct {
	ServerId   int32  `json:"serverId" gorm:"column:serverId"` //服务器id
	PlayerId   int64  `json:"playerId" gorm:"column:playerId"` //玩家id
	PlayerName string `json:"name" gorm:"column:name"`         //玩家名字
	GangName   string `json:"gangName" gorm:"column:gangName"` //帮派名字
	Force      int64  `json:"force" gorm:"column:power"`       //战斗力
	Level      int32  `json:"level" gorm:"column:level"`       //等级
	Role       int32  `json:"role" gorm:"column:role"`         //角色
	Sex        int32  `json:"sex" gorm:"column:sex"`           //性别
}

//帮派 MySql+redis
type PlayerGangData struct {
	ServerId int32  `json:"serverId" gorm:"column:serverId"`  //服务器id
	GangId   int64  `json:"gangId" gorm:"column:id"`          //帮派id
	GangName string `json:"gangName" gorm:"column:gangName" ` //帮派
	LeadName string `json:"LeadName" gorm:"column:leadName"`  //帮派大哥
	LeadId   int64  `json:"LeadId" gorm:"column:mengzhuId"`   //大哥id
	Power    int64  `json:"power" gorm:"column:totalForce"`   //帮派战斗力
	Role     int32  `json:"role" gorm:"column:role"`          //角色
	Sex      int32  `json:"sex" gorm:"column:sex"`            //性别
}

//坐骑 战翼 护体盾 暗器 法宝 仙体MySql+redis
type PlayerOrderData struct {
	ServerId   int32  `json:"serverId" gorm:"column:serverId"` //服务器id
	PlayerId   int64  `json:"playerId" gorm:"column:playerId"` //玩家id
	PlayerName string `json:"name" gorm:"column:name"`         //玩家名字
	Order      int32  `json:"order" gorm:"column:advancedId"`  //阶数
	Power      int64  `json:"power" gorm:"column:power"`       //战斗力
}

//兵魂MySql+redis
type PlayerWeaponData struct {
	ServerId   int32  `json:"serverId" gorm:"column:serverId"` //服务器id
	PlayerId   int64  `json:"playerId" gorm:"column:playerId"` //玩家id
	PlayerName string `json:"name" gorm:"column:name"`         //玩家名字
	Star       int32  `json:"star" gorm:"column:star"`         //兵魂总星数
	WearId     int32  `json:"wearId" gorm:"column:weaponWear"` //穿戴兵魂id
	Power      int64  `json:"power" gorm:"column:power"`       //战斗力
	Role       int32  `json:"role" gorm:"column:role"`         //角色
	Sex        int32  `json:"sex" gorm:"column:sex"`           //性别
}

//充值 消费 魅力 表白 次数 等级 灵童战力MySql+redis
type PlayerPropertyData struct {
	ServerId   int32  `json:"serverId" gorm:"column:serverId"` //服务器id
	PlayerId   int64  `json:"playerId" gorm:"column:playerId"` //玩家id
	PlayerName string `json:"name" gorm:"column:name"`         //玩家名字
	Num        int64  `json:"num" gorm:"column:num"`           //数值(排序)
	Power      int64  `json:"power" gorm:"column:power"`       //战斗力
}

type RankCommonData struct {
	Id    int64  `json:"id"`    //id
	Value int64  `json:"value"` //数值
	Name  string `json:"name"`  //名字
}

func NewRankCommonData(id int64, value int64, name string) *RankCommonData {
	data := &RankCommonData{
		Id:    id,
		Value: value,
		Name:  name,
	}
	return data
}

type RankCommonDataList []*RankCommonData

func (rcdl RankCommonDataList) Len() int {
	return len(rcdl)
}

func (rcdl RankCommonDataList) Less(i, j int) bool {
	return rcdl[i].Value < rcdl[j].Value
}

func (rcdl RankCommonDataList) Swap(i, j int) {
	rcdl[i], rcdl[j] = rcdl[j], rcdl[i]
}
