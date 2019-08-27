package tower

//打宝塔日志列表对象
type TowerLogObject struct {
	playerName  string
	biologyName string
	itemId      int32
	itemNum     int32
	createTime  int64
}

func NewTowerLogObject() *TowerLogObject {
	o := &TowerLogObject{}
	return o
}

func (o *TowerLogObject) GetPlayerName() string {
	return o.playerName
}

func (o *TowerLogObject) GetBiologyName() string {
	return o.biologyName
}

func (o *TowerLogObject) GetItemId() int32 {
	return o.itemId
}

func (o *TowerLogObject) GetItemNum() int32 {
	return o.itemNum
}

func (o *TowerLogObject) GetCreateTime() int64 {
	return o.createTime
}
