package pbmodel

type RedeemCodeInfo struct {
	Id         int64  `json:"id"`
	RedeemCode string `json:"redeemCode"`
	RedeemId   int    `json:"redeemId"`
	UseNum     int    `json:"useNum"`
}
