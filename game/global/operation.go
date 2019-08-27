package global

import (
	"fgame/fgame/core/worker"
	"fgame/fgame/game/common/common"

	"sync"

	log "github.com/Sirupsen/logrus"
)

type OpeartionService interface {
	Stop()
	PostOperation(o worker.Operation)
}

const (
	defaultOperationCapacity = 10000
)

//TODO 分开全局和非全局
type opeartionService struct {
	m            sync.Mutex
	workerGroup  map[int64]worker.Worker
	globalWorker worker.Worker
	size         int32
}

func (os *opeartionService) Stop() {
	log.Info("global:关闭操作服务,开始")
	beforeTime := GetGame().GetTimeService().Now()
	for _, w := range os.workerGroup {
		w.Shutdown()
	}
	os.globalWorker.Shutdown()
	now := GetGame().GetTimeService().Now()
	elapse := (now - beforeTime) / int64(common.SECOND)
	log.WithFields(
		log.Fields{
			"costTime": elapse,
		}).Info("global:关闭操作服务")
	return
}

func (os *opeartionService) getWorkerId(id int64) int64 {
	workerId := id % int64(os.size)
	return workerId
}

func (os *opeartionService) PostOperation(o worker.Operation) {
	//判断是不是绑定
	no, ok := o.(worker.BindOperation)
	if ok {
		os.post(no.BindUUID(), o)
	} else {
		os.globalWorker.Post(o)
	}
	return
}

//TODO 优化
func (os *opeartionService) post(bindId int64, o worker.Operation) {
	os.m.Lock()
	defer os.m.Unlock()
	workerId := os.getWorkerId(bindId)
	w, exist := os.workerGroup[workerId]
	if !exist {
		log.WithFields(
			log.Fields{
				"workerId": workerId,
			}).Infoln("创建worker")
		w = worker.NewWorker(defaultOperationCapacity)
		w.Run()
		os.workerGroup[workerId] = w
	}
	w.Post(o)
}

func NewOperationService(operationPoolSize int32) OpeartionService {
	os := &opeartionService{}
	os.size = operationPoolSize
	os.workerGroup = make(map[int64]worker.Worker)
	os.globalWorker = worker.NewWorker(defaultOperationCapacity)
	os.globalWorker.Run()

	return os
}
