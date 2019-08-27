package common

//灵童信息信息
type LingTongInfo struct {
	LingTongId   int32             `json:"lingTongId"`
	FashionId    int32             `json:"fashionId"`
	Level        int32             `json:"level"`
	LingTongList []*LingTongDetail `json:"lingTongDetail"`
}

type LingTongDetail struct {
	LingTongId   int32  `json:"lingTongId"`
	LingTongName string `json:"lingTongName"`
	FashionId    int32  `json:"fashionId"`
	Level        int32  `json:"level"`
	LevelPro     int32  `json:"levelPro"`
	PeiYangLevel int32  `json:"peiYangLevel"`
	PeiYangPro   int32  `json:"peiYangPro"`
	StarLevel    int32  `json:"starLevel"`
	StarPro      int32  `json:"starPro"`
}
