package worker

import (
	"errors"
	"runtime/debug"
	"time"

	log "github.com/Sirupsen/logrus"
)

type Worker interface {
	Run() error
	Shutdown()
	Post(op Operation) error
}

func NewWorker(capacity int) Worker {
	w := &worker{}
	w.ops = make(chan Operation, capacity)
	w.done = make(chan struct{})
	return w
}

const (
	defaultCapacity = 10000
	defaultTimeout  = time.Duration(time.Second)
)

func NewDefaultWorker() Worker {
	return NewWorker(defaultCapacity)
}

type worker struct {
	ops  chan Operation
	done chan struct{}
}

var (
	PostOperationTimeoutError = errors.New("post operation time out")
)

func (w *worker) Post(op Operation) error {
	//TODO 判断worker是不是正在关闭
	select {
	case w.ops <- op:
		return nil
	case <-time.After(defaultTimeout):
		return PostOperationTimeoutError
	}
}

func (w *worker) Run() (err error) {
	go func() {
		log.Infoln("worker running")
		defer func() {
			close(w.done)
		}()
		for {
			op, ok := <-w.ops
			if !ok {
				break
			}
			w.run(op)
		}
		log.Infoln("worker done")
	}()
	return nil
}

var (
	OperationPanicError = errors.New("operation panic")
)

func (w *worker) run(op Operation) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			log.WithFields(
				log.Fields{
					"err":   r,
					"stack": string(debug.Stack()),
				}).Error("worker:操作错误")
			//异常回调
			terr, ok := r.(error)
			if ok {
				op.OnCallBack(nil, terr)
				return
			}
			op.OnCallBack(nil, OperationPanicError)
			return
		}
	}()
	result, err := op.Run()
	if err != nil {
		//错误回调
		op.OnCallBack(nil, err)
		return
	}
	//正确回调
	op.OnCallBack(result, nil)
}

func (w *worker) Shutdown() {
	close(w.ops)
	<-w.done
}
