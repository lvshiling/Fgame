package pbmodel

type RedeemInfo struct {
	Id                 int    `json:"id"`
	GiftBagName        string `json:"giftBagName"`
	GiftBagDesc        string `json:"giftBagDesc"`
	GiftBagContent     string `json:"giftBagContent"`
	RedeemNum          int    `json:"redeemNum"`
	RedeemUseNum       int    `json:"redeemUseNum"`
	RedeemPlayerUseNum int    `json:"redeemPlayerUseNum"`
	RedeemServerUseNum int    `json:"redeemServerUseNum"`
	SdkTypes           []int  `json:"sdkTypes"`
	SendType           int    `json:"sendType"`
	StartTime          int64  `json:"startTime"`
	EndTime            int64  `json:"endTime"`
	MinPlayerLevel     int    `json:"minPlayerLevel"`
	MinVipLevel        int    `json:"minVipLevel"`
	CreateFlag         int    `json:"createFlag"`
	CreateTime         int64  `json:"createTime"`
}
