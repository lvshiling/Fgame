package wxpay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

type IWxPay interface {
}

type WxPayConfig struct {
	Key      string `json:"key"`       //商务号中的设置的密钥
	MchAppid string `json:"mch_appid"` //申请商户号的appid或商户号绑定的appid
	MchID    string `json:"mchid"`
	CertPath string `json:"certpath"` //证书绝对路径
}

type WxPayService struct {
	_config     *WxPayConfig
	_tlsService *TlsService
}

var transfers = "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"

//付款订单
type WithdrawOrder struct {
	XMLName        xml.Name `xml:"xml"`
	MchAppid       string   `xml:"mch_appid"`
	Mchid          string   `xml:"mchid"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	Openid         string   `xml:"openid"`
	CheckName      string   `xml:"check_name"`
	Amount         int      `xml:"amount"`
	Desc           string   `xml:"desc"`
	SpbillCreateIp string   `xml:"spbill_create_ip"`
}

//付款订单结果
type WithdrawResult struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	ResultCode     string `xml:"result_code"`
	ErrorCode      string `xml:"err_code"`
	ErrCodeDes     string `xml:"err_code_des"`
	PaymentNo      string `xml:"payment_no"`
	PartnerTradeNo string `xml:"partner_trade_no"`
	PaymentTime    string `xml:"payment_time"`
	Mchid          string `xml:"mchid"`
}

//付款，成功返回自定义订单号，微信订单号，true，失败返回错误信息，false
func (m *WxPayService) WithdrawMoney(openid, amount, partnerTradeNo, desc, clientIp string) (*WithdrawResult, string, error) {
	order := WithdrawOrder{}
	order.MchAppid = m._config.MchAppid
	order.Mchid = m._config.MchID
	order.Openid = openid
	order.Amount, _ = strconv.Atoi(amount)
	order.Desc = desc
	order.PartnerTradeNo = partnerTradeNo
	order.DeviceInfo = "WEB"
	order.CheckName = "NO_CHECK" //NO_CHECK：不校验真实姓名 FORCE_CHECK：强校验真实姓名
	order.SpbillCreateIp = clientIp
	order.NonceStr = GetRandomString(32)
	order.Sign = m.md5WithdrawOrder(order)
	xmlBody, _ := xml.MarshalIndent(order, " ", " ")
	resp, err := m._tlsService.SecurePost(transfers, xmlBody)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	var res WithdrawResult
	xmlerr := xml.Unmarshal(bodyByte, &res)
	if xmlerr != nil {
		return nil, "返回xml解析错误", xmlerr
	}
	if res.ReturnCode == "SUCCESS" && res.ResultCode == "SUCCESS" {
		return &res, "", nil
	}
	if res.ReturnCode != "SUCCESS" {
		return nil, res.ReturnMsg, errors.New("returnMsg错误")
	}
	return nil, res.ErrCodeDes, GetWxPayWithDrawError(res.ErrorCode)
}

//md5签名
func (m *WxPayService) md5WithdrawOrder(order WithdrawOrder) string {
	o := url.Values{}
	o.Add("mch_appid", order.MchAppid)
	o.Add("mchid", order.Mchid)
	o.Add("device_info", order.DeviceInfo)
	o.Add("partner_trade_no", order.PartnerTradeNo)
	o.Add("check_name", order.CheckName)
	o.Add("amount", strconv.Itoa(order.Amount))
	o.Add("spbill_create_ip", order.SpbillCreateIp)
	o.Add("desc", order.Desc)
	o.Add("nonce_str", order.NonceStr)
	o.Add("openid", order.Openid)
	r, _ := url.QueryUnescape(o.Encode())
	return Md5([]byte(r + "&key=" + m._config.Key))
}

func Md5(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	nt := hex.EncodeToString(cipherStr)
	return strings.ToUpper(nt)
}

func GetRandomString(p_len int) string {
	rst := uuid.NewV4().String()
	rst = strings.Replace(rst, "-", "", -1)
	return rst[:p_len]
}

func NewWxPayService(p_config *WxPayConfig) *WxPayService {
	if p_config == nil {
		return nil
	}
	rst := &WxPayService{}
	rst._config = p_config
	rst._tlsService = NewTlsService(rst._config.CertPath)
	return rst
}

var (
	instanceWxPayService *WxPayService
)

func InitInstance(p_config *WxPayConfig) {
	instanceWxPayService = NewWxPayService(p_config)
}

func GetWxPayServiceInstance() *WxPayService {
	return instanceWxPayService
}
