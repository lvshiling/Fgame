package types

type LingwenInfo struct {
	Level      int32 `json:"level"`
	Experience int64 `json:"experience"`
}

type LinglianInfo struct {
	PoolMark int32 `json:"poolMark"`
	// IsLock    bool  `json:"isLock"`
	// LockTimes int32 `json:"lockTimes"`
}
