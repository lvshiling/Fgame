package idutil

import (
	"github.com/zheng-ji/goSnowFlake"
)

var (
	iw *goSnowFlake.IdWorker
)

func GetId() (id int64, err error) {
	// Params: Given the workerId, 0 < workerId < 1024
	// if iw == nil {
	// 	tiw, err := goSnowFlake.NewIdWorker(1)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	iw = tiw
	// }
	// id, err = iw.NextId()
	// if err != nil {
	// 	return
	// }
	// return
	return s.NextId(), nil
}
