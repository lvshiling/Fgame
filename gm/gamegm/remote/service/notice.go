package service

import (
	"context"
	cnclient "fgame/fgame/center/client"
	"net/http"

	"github.com/codegangsta/negroni"
)

type INotice interface {
	RefreshNotice() (err error)
}

type notice struct {
	manager cnclient.NoticeManager
}

func (m *notice) RefreshNotice() (err error) {
	ctx := context.Background()
	_, err = m.manager.RefreshNotice(ctx)
	return
}

func NewNotice(p_host string, p_port int32) (INotice, error) {
	rst := &notice{}
	config := &cnclient.Config{
		Host: p_host,
		Port: p_port,
	}
	client, err := cnclient.NewClient(config)
	if err != nil {
		return nil, err
	}
	manager := cnclient.NewNoticeManager(client)
	rst.manager = manager
	return rst, nil
}

const (
	noticeKey = contextKey("Notice")
)

func WithNotice(ctx context.Context, ls INotice) context.Context {
	return context.WithValue(ctx, noticeKey, ls)
}

func NoticeInContext(ctx context.Context) INotice {
	us, ok := ctx.Value(noticeKey).(INotice)
	if !ok {
		return nil
	}
	return us
}

func SetupNoticeHandler(ls INotice) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, req *http.Request, hf http.HandlerFunc) {
		ctx := req.Context()
		nctx := WithNotice(ctx, ls)
		nreq := req.WithContext(nctx)
		hf(rw, nreq)
	})
}
