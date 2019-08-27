package grpcutil

import (
	"container/list"
	"container/ring"
	"errors"
	"sync"

	"google.golang.org/grpc"
)

var (
	ErrGrpcPoolExhausted = errors.New("grpc_pool: pool exhausted")
)

type Pool interface {
	Get() (*grpc.ClientConn, error)
	Put(*grpc.ClientConn) error
	Close() error
}

type pool struct {
	mu          sync.Mutex
	cond        *sync.Cond
	idle        *list.List
	active      int
	maxActive   int
	wait        bool
	addr        string
	dialOptions []grpc.DialOption
	ring        *ring.Ring
}

func (p *pool) Get() (conn *grpc.ClientConn, err error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	tval := p.ring.Value
	if tval == nil {
		conn, err = p.dialNew()
		if err != nil {
			return
		}
		p.ring.Value = conn
		tval = conn
	}
	p.ring = p.ring.Next()
	return tval.(*grpc.ClientConn), nil

	// p.mu.Lock()

	// //取出空闲的
	// if p.idle.Len() != 0 {
	// 	e := p.idle.Front()
	// 	ev := p.idle.Remove(e)

	// 	p.mu.Unlock()
	// 	return ev.(*grpc.ClientConn), nil
	// }

	// //还没达到最大激活量
	// if p.active < p.maxActive {
	// 	//新建连接
	// 	conn, err = p.dialNew()
	// 	if err != nil {
	// 		p.mu.Unlock()
	// 		return
	// 	}
	// 	p.active += 1
	// 	p.mu.Unlock()
	// 	return
	// }

	// //不等待释放
	// if !p.wait {
	// 	p.mu.Unlock()
	// 	return nil, ErrGrpcPoolExhausted
	// }

	// for p.idle.Len() == 0 {
	// 	p.cond.Wait()
	// }
	// e := p.idle.Front()
	// ev := p.idle.Remove(e)
	// p.mu.Unlock()
	// return ev.(*grpc.ClientConn), nil
}

func (p *pool) Put(conn *grpc.ClientConn) error {
	// p.mu.Lock()
	// defer p.mu.Unlock()
	// p.idle.PushBack(conn)
	// p.cond.Signal()
	return nil
}

func (p *pool) Close() error {
	for e := p.idle.Front(); e != nil; {
		ev := p.idle.Remove(e)
		ev.(*grpc.ClientConn).Close()
	}
	return nil
}

func (p *pool) dialNew() (conn *grpc.ClientConn, err error) {
	return grpc.Dial(p.addr, p.dialOptions...)
}

func NewPool(addr string, maxActive int, dialOptions ...grpc.DialOption) *pool {
	p := &pool{}
	p.addr = addr
	p.maxActive = maxActive
	p.wait = true
	p.dialOptions = dialOptions
	p.idle = list.New()
	p.cond = sync.NewCond(&p.mu)
	p.ring = ring.New(maxActive)
	return p
}
