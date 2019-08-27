package cross

import (
	"sync"

	"google.golang.org/grpc"
)

type CrossService interface {
	Start()
	GetCross() Cross
}

type crossService struct {
	conn  *grpc.ClientConn
	cross Cross
}

func (s *crossService) init() (err error) {
	//注册

	s.cross = newCross()
	return
}

func (s *crossService) Start() {
	go func() {
		s.cross.Start()
	}()
	return
}

func (s *crossService) GetCross() Cross {
	return s.cross
}

func newCrossService() CrossService {
	s := &crossService{}
	return s
}

func GetCrossService() CrossService {
	return s
}

var (
	once sync.Once
	s    *crossService
)

func Init() (err error) {
	once.Do(func() {
		s = &crossService{}
		err = s.init()
	})
	return nil
}
