package wxapi

import (
	"context"

	wechat "github.com/silenceper/wechat"
	"github.com/silenceper/wechat/oauth"
)

type WeChatService struct {
	Wc            *wechat.Config
	weChatService *wechat.Wechat
	oauthService  *oauth.Oauth
}

func (m *WeChatService) GetOauth() *oauth.Oauth {
	return m.oauthService
}

func (m *WeChatService) GetWeChat() *wechat.Wechat {
	return m.weChatService
}

func NewWeChatService(p_wc *wechat.Config) *WeChatService {
	rst := &WeChatService{
		Wc: p_wc,
	}
	rst.weChatService = wechat.NewWechat(p_wc)
	rst.oauthService = oauth.NewOauth(rst.weChatService.Context)
	return rst
}

type WeChatServiceKey string

const (
	wechatservicekey = WeChatServiceKey("wechatservicekey")
)

func WeChatServiceInContext(ctx context.Context) *WeChatService {
	c, ok := ctx.Value(wechatservicekey).(*WeChatService)
	if !ok {
		return nil
	}
	return c
}
func WithWeChatService(ctx context.Context, ss *WeChatService) context.Context {
	return context.WithValue(ctx, wechatservicekey, ss)
}
