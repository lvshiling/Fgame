package arenapvp

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ArenapvpRankDataTask struct {
	s *arenapvpService
}

const (
	dataTaskTime = 10 * time.Second
)

func (t *ArenapvpRankDataTask) Run() {
	err := t.s.syncRemoteArenapvpData()
	if err != nil {
		sta := status.Convert(err)
		if sta.Code() == codes.Canceled {
			//重新获取
			log.WithFields(
				log.Fields{
					"err": err,
				}).Warn("arenapvp:同步pvp赛程,重新获取客户端")
			err = t.s.resetClient()
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("arenapvp:同步pvp赛程,失败")
				return
			}
			err = t.s.syncRemoteArenapvpData()
			if err != nil {
				log.WithFields(
					log.Fields{
						"err": err,
					}).Warn("arenapvp:同步pvp赛程,失败")
				return
			}
			return
		}
		log.WithFields(
			log.Fields{
				"err": err,
			}).Warn("arenapvp:同步pvp赛程,失败")
	}
}

//间隔时间
func (t *ArenapvpRankDataTask) ElapseTime() time.Duration {
	return dataTaskTime
}

func CreateArenapvpPlayerDataTask(s *arenapvpService) *ArenapvpRankDataTask {
	t := &ArenapvpRankDataTask{
		s: s,
	}
	return t
}
