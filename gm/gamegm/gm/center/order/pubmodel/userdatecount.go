package pubmodel

type UserDateCount struct {
	DateTime   int64 `json:"dateTime"`
	LeiJiCount int   `json:"leijiCount"`
	DateCount  int   `json:"dateCount"`
}
