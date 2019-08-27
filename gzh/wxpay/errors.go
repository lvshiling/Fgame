package wxpay

import "errors"

var (
	WxPayWithDrawError_NO_AUTH               = errors.New("NO_AUTH:没有该接口权限")
	WxPayWithDrawError_AMOUNT_LIMIT          = errors.New("AMOUNT_LIMIT:金额超限")
	WxPayWithDrawError_PARAM_ERROR           = errors.New("PARAM_ERROR:参数错误")
	WxPayWithDrawError_OPENID_ERROR          = errors.New("OPENID_ERROR:Openid错误")
	WxPayWithDrawError_SEND_FAILED           = errors.New("SEND_FAILED:付款错误")
	WxPayWithDrawError_NOTENOUGH             = errors.New("NOTENOUGH:余额不足")
	WxPayWithDrawError_SYSTEMERROR           = errors.New("SYSTEMERROR:系统繁忙，请稍后再试")
	WxPayWithDrawError_NAME_MISMATCH         = errors.New("NAME_MISMATCH:姓名校验出错")
	WxPayWithDrawError_SIGN_ERROR            = errors.New("SIGN_ERROR:签名错误")
	WxPayWithDrawError_XML_ERROR             = errors.New("XML_ERROR:Post内容出错")
	WxPayWithDrawError_FATAL_ERROR           = errors.New("FATAL_ERROR:两次请求参数不一致")
	WxPayWithDrawError_FREQ_LIMIT            = errors.New("FREQ_LIMIT:没有该接口权限")
	WxPayWithDrawError_MONEY_LIMIT           = errors.New("MONEY_LIMIT:已经达到今日付款总额上限/已达到付款给此用户额度上限")
	WxPayWithDrawError_CA_ERROR              = errors.New("CA_ERROR:商户API证书校验出错")
	WxPayWithDrawError_V2_ACCOUNT_SIMPLE_BAN = errors.New("V2_ACCOUNT_SIMPLE_BAN:无法给非实名用户付款")
	WxPayWithDrawError_PARAM_IS_NOT_UTF8     = errors.New("PARAM_IS_NOT_UTF8:请求参数中包含非utf8编码字符")
	WxPayWithDrawError_SENDNUM_LIMIT         = errors.New("SENDNUM_LIMIT:该用户今日付款次数超过限制,如有需要请登录微信支付商户平台更改API安全配置")
	WxPayWithDrawError_Other                 = errors.New("Other:未知错误")
)

var (
	errorMap = make(map[string]error)
)

func init() {
	errorMap["NO_AUTH"] = WxPayWithDrawError_NO_AUTH
	errorMap["AMOUNT_LIMIT"] = WxPayWithDrawError_AMOUNT_LIMIT
	errorMap["PARAM_ERROR"] = WxPayWithDrawError_PARAM_ERROR
	errorMap["OPENID_ERROR"] = WxPayWithDrawError_OPENID_ERROR
	errorMap["SEND_FAILED"] = WxPayWithDrawError_SEND_FAILED
	errorMap["NOTENOUGH"] = WxPayWithDrawError_NOTENOUGH
	errorMap["SYSTEMERROR"] = WxPayWithDrawError_SYSTEMERROR
	errorMap["NAME_MISMATCH"] = WxPayWithDrawError_NAME_MISMATCH
	errorMap["SIGN_ERROR"] = WxPayWithDrawError_SIGN_ERROR
	errorMap["XML_ERROR"] = WxPayWithDrawError_XML_ERROR
	errorMap["FATAL_ERROR"] = WxPayWithDrawError_FATAL_ERROR
	errorMap["FREQ_LIMIT"] = WxPayWithDrawError_FREQ_LIMIT
	errorMap["MONEY_LIMIT"] = WxPayWithDrawError_MONEY_LIMIT
	errorMap["CA_ERROR"] = WxPayWithDrawError_CA_ERROR
	errorMap["V2_ACCOUNT_SIMPLE_BAN"] = WxPayWithDrawError_V2_ACCOUNT_SIMPLE_BAN
	errorMap["PARAM_IS_NOT_UTF8"] = WxPayWithDrawError_PARAM_IS_NOT_UTF8
	errorMap["SENDNUM_LIMIT"] = WxPayWithDrawError_SENDNUM_LIMIT
}

func GetWxPayWithDrawError(p_string string) error {
	if value, ok := errorMap[p_string]; ok {
		return value
	}
	return WxPayWithDrawError_Other
}
