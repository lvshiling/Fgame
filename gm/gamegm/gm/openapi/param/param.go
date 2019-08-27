package param

import "context"

type contextKey string

const (
	apiDataKey = contextKey("ApiDataKey")
)

func WithApiData(ctx context.Context, ls map[string]string) context.Context {
	return context.WithValue(ctx, apiDataKey, ls)
}

func ApiDataInContext(ctx context.Context) map[string]string {
	us, ok := ctx.Value(apiDataKey).(map[string]string)
	if !ok {
		return nil
	}
	return us
}
