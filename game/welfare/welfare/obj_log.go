package welfare

//元宝拉霸日志列表对象
type GoldLaBaLogObject struct {
	id         int64
	serverId   int32
	groupId    int32
	playerName string
	costGold   int32
	rewGold    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newGoldLaBaLogObject() *GoldLaBaLogObject {
	o := &GoldLaBaLogObject{}
	return o
}

func (o *GoldLaBaLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *GoldLaBaLogObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *GoldLaBaLogObject) GetPlayerName() string {
	return o.playerName
}

func (o *GoldLaBaLogObject) GetRewGold() int32 {
	return o.rewGold
}

func (o *GoldLaBaLogObject) GetGroupId() int32 {
	return o.groupId
}

func (o *GoldLaBaLogObject) GetCostGold() int32 {
	return o.costGold
}

func (o *GoldLaBaLogObject) GetDBId() int64 {
	return o.id
}

func (o *GoldLaBaLogObject) SetModified() {
	return
}

//抽奖日志列表对象
type DrewLogObject struct {
	id         int64
	serverId   int32
	groupId    int32
	playerName string
	itemId     int32
	itemNum    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newDrewLogObject() *DrewLogObject {
	o := &DrewLogObject{}
	return o
}

func (o *DrewLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *DrewLogObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *DrewLogObject) GetPlayerName() string {
	return o.playerName
}

func (o *DrewLogObject) GetItemId() int32 {
	return o.itemId
}

func (o *DrewLogObject) GetGroupId() int32 {
	return o.groupId
}

func (o *DrewLogObject) GetItemNum() int32 {
	return o.itemNum
}

func (o *DrewLogObject) GetDBId() int64 {
	return o.id
}

func (o *DrewLogObject) SetModified() {
	return
}

//疯狂宝箱日志列表对象
type CrazyBoxLogObject struct {
	id         int64
	serverId   int32
	groupId    int32
	playerName string
	itemId     int32
	itemNum    int32
	updateTime int64
	createTime int64
	deleteTime int64
}

func newCrazyBoxLogObject() *CrazyBoxLogObject {
	o := &CrazyBoxLogObject{}
	return o
}

func (o *CrazyBoxLogObject) GetCreateTime() int64 {
	return o.createTime
}

func (o *CrazyBoxLogObject) GetUpdateTime() int64 {
	return o.updateTime
}

func (o *CrazyBoxLogObject) GetPlayerName() string {
	return o.playerName
}

func (o *CrazyBoxLogObject) GetItemId() int32 {
	return o.itemId
}

func (o *CrazyBoxLogObject) GetGroupId() int32 {
	return o.groupId
}

func (o *CrazyBoxLogObject) GetItemNum() int32 {
	return o.itemNum
}

func (o *CrazyBoxLogObject) GetDBId() int64 {
	return o.id
}

func (o *CrazyBoxLogObject) SetModified() {
	return
}
