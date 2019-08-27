package coupon

type ExchangeFinishEventData struct {
	playerId   int64
	title      string
	content    string
	attachment string
}

func (d *ExchangeFinishEventData) GetTitle() string {
	return d.title
}
func (d *ExchangeFinishEventData) GetContent() string {
	return d.content
}

func (d *ExchangeFinishEventData) GetAttachment() string {
	return d.attachment
}

func CreateExchangeFinishEventData(title string, content string, attachment string) *ExchangeFinishEventData {
	d := &ExchangeFinishEventData{

		title:      title,
		content:    content,
		attachment: attachment,
	}
	return d
}

type ExchangeFailedEventData struct {
	code int32
	msg  string
}

func (d *ExchangeFailedEventData) GetCode() int32 {
	return d.code
}
func (d *ExchangeFailedEventData) GetMsg() string {
	return d.msg
}

func CreateExchangeFailedEventData(code int32, msg string) *ExchangeFailedEventData {
	d := &ExchangeFailedEventData{

		code: code,
		msg:  msg,
	}
	return d
}
