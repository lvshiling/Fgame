package coupon

//兑换码
type RedeemUseNumEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	RedeemId   int32 `gorm:"column:redeemId"`
	UseNum     int32 `gorm:"column:useNum"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
	UpdateTime int64 `gorm:"column:updateTime"`
}

func (e *RedeemUseNumEntity) TableName() string {
	return "t_redeem_use_num"
}

//兑换码
type RedeemCodeEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	RedeemCode string `gorm:"column:redeemCode"`
	RedeemId   int32  `gorm:"column:redeemId"`
	UseNum     int32  `gorm:"column:useNum"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
	UpdateTime int64  `gorm:"column:updateTime"`
}

func (e *RedeemCodeEntity) TableName() string {
	return "t_redeem_code"
}

//兑换码
type RedeemCompleteCodeEntity struct {
	Id                 int64  `gorm:"primary_key;column:id"`
	RedeemCode         string `gorm:"column:redeemCode"`
	RedeemId           int32  `gorm:"column:redeemId"`
	UseNum             int32  `gorm:"column:useNum"`
	GiftBagName        string `gorm:"column:giftBagName"`
	GiftBagDesc        string `gorm:"column:giftBagDesc"`
	GiftBagContent     string `gorm:"column:giftBagContent"`
	ReedmUseNum        int32  `gorm:"column:redeemUseNum"`
	RedeemPlayerUseNum int32  `gorm:"column:redeemPlayerUseNum"`
	RedeemServerUseNum int32  `gorm:"column:redeemServerUseNum"`
	SdkTypes           string `gorm:"column:sdkTypes"`
	SendType           int32  `gorm:"column:sendType"`
	StartTime          int64  `gorm:"column:startTime"`
	EndTime            int64  `gorm:"column:endTime"`
	MinPlayerLevel     int32  `gorm:"column:minPlayerLevel"`
	MinVipLevel        int32  `gorm:"column:minVipLevel"`
	CreateTime         int64  `gorm:"column:createTime"`
	DeleteTime         int64  `gorm:"column:deleteTime"`
	UpdateTime         int64  `gorm:"column:updateTime"`
}

//兑换记录
type RedeemRecordEntity struct {
	Id             int64  `gorm:"primary_key;column:id"`
	RedeemCode     string `gorm:"column:redeemCode"`
	RedeemId       int32  `gorm:"column:redeemId"`
	PlatformId     int32  `gorm:"column:platformId"`
	ServerId       int32  `gorm:"column:serverId"`
	SdkType        int32  `gorm:"column:sdkType"`
	PlatformUserId string `gorm:"column:platformUserId"`
	UserId         int64  `gorm:"column:userId"`
	PlayerId       int64  `gorm:"column:playerId"`
	PlayerLevel    int32  `gorm:"column:playerLevel"`
	PlayerVipLevel int32  `gorm:"column:playerVipLevel"`
	PlayerName     string `gorm:"column:playerName"`
	CreateTime     int64  `gorm:"column:createTime"`
	DeleteTime     int64  `gorm:"column:deleteTime"`
	UpdateTime     int64  `gorm:"column:updateTime"`
}

func (e *RedeemRecordEntity) TableName() string {
	return "t_redeem_record"
}
