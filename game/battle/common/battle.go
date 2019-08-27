package common

type PlayerBattleObject struct {
	level         int32
	zhuanSheng    int32
	vip           int32
	soulAwakenNum int32
	isHuiYuan     bool
}

func (o *PlayerBattleObject) GetLevel() int32 {
	return o.level
}

func (o *PlayerBattleObject) GetVip() int32 {
	return o.vip
}

func (o *PlayerBattleObject) GetZhuanSheng() int32 {
	return o.zhuanSheng
}

func (o *PlayerBattleObject) GetIsHuiYuan() bool {
	return o.isHuiYuan
}

func (o *PlayerBattleObject) GetSoulAwakenNum() int32 {
	return o.soulAwakenNum
}

func CreatePlayerBattleObject(vip int32, level int32, zhuanSheng int32, soulAwakenNum int32, isHuiYuan bool) *PlayerBattleObject {
	o := &PlayerBattleObject{
		vip:           vip,
		level:         level,
		zhuanSheng:    zhuanSheng,
		soulAwakenNum: soulAwakenNum,
		isHuiYuan:     isHuiYuan,
	}
	return o
}
