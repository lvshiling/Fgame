package shenmo

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShenMoRankDataTask struct {
	s *shenMoService
}

const (
	dataTaskTime = time.Minute
)

func (t *ShenMoRankDataTask) Run() {
	err := t.s.syncRemoteRankList()
	if err != nil {
		sta := status.Convert(err)
		if sta.Code() == codes.Canceled {
			//重新获取
			log.WithFields(
				log.Fields{
					"err": err,
				}).Warn("shenmo:同步神魔战场,重新获取客户端")
			err = t.s.resetClient()
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("shenmo:同步神魔战场,失败")
				return
			}
			err = t.s.syncRemoteRankList()
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("shenmo:同步神魔战场,失败")
				return
			}
			return
		}
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("shenmo:同步神魔战场,失败")
	}

	t.s.syncLocalRankList()
}

//间隔时间
func (t *ShenMoRankDataTask) ElapseTime() time.Duration {
	return dataTaskTime
}

func CreateShenMoRankDataTask(s *shenMoService) *ShenMoRankDataTask {
	t := &ShenMoRankDataTask{
		s: s,
	}
	return t
}
