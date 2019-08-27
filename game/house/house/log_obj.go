package house

//房子日志列表对象
type HouseLogObject struct {
	playerName  string
	houseType   int32
	houseLevel  int32
	houseIndex  int32
	operateType int32
	updateTime  int64
	createTime  int64
}

func NewHouseLogObject() *HouseLogObject {
	o := &HouseLogObject{}
	return o
}

func (o *HouseLogObject) GetPlayerName() string {
	return o.playerName
}

func (o *HouseLogObject) GetHouseType() int32 {
	return o.houseType
}

func (o *HouseLogObject) GetHouseLevel() int32 {
	return o.houseLevel
}

func (o *HouseLogObject) GetHouseIndex() int32 {
	return o.houseIndex
}

func (o *HouseLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *HouseLogObject) GetHouseOperateType() int32 {
	return o.operateType
}
