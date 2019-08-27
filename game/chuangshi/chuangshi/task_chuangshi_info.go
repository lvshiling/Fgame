package chuangshi

// import (
// 	"time"

// 	log "github.com/Sirupsen/logrus"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// type ChuangShiWholeInfoTask struct {
// 	s *chuangShiService
// }

// const (
// 	dataTaskTime = 10 * time.Second
// )

// func (t *ChuangShiWholeInfoTask) Run() {
// 	err := t.s.syncRemoteChuangshiData()
// 	if err != nil {
// 		sta := status.Convert(err)
// 		if sta.Code() == codes.Canceled {
// 			//重新获取
// 			log.WithFields(
// 				log.Fields{
// 					"err": err,
// 				}).Warn("chuangshi:同步创世信息,重新获取客户端")
// 			err = t.s.resetClient()
// 			if err != nil {
// 				log.WithFields(
// 					log.Fields{
// 						"err": err,
// 					}).Warn("chuangshi:同步创世信息,失败")
// 				return
// 			}
// 			err = t.s.syncRemoteChuangshiData()
// 			if err != nil {
// 				log.WithFields(
// 					log.Fields{
// 						"err": err,
// 					}).Warn("chuangshi:同步创世信息,失败")
// 				return
// 			}
// 			return
// 		}
// 		log.WithFields(
// 			log.Fields{
// 				"err": err,
// 			}).Warn("chuangshi:同步创世信息,失败")
// 	}
// }

// //间隔时间
// func (t *ChuangShiWholeInfoTask) ElapseTime() time.Duration {
// 	return dataTaskTime
// }

// func CreateArenapvpPlayerDataTask(s *chuangShiService) *ChuangShiWholeInfoTask {
// 	t := &ChuangShiWholeInfoTask{
// 		s: s,
// 	}
// 	return t
// }
