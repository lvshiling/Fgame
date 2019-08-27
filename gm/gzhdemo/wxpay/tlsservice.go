package wxpay

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/http"
)

//ca证书的位置，需要绝对路径
type TlsService struct {
	_caPath    string
	_tlsConfig *tls.Config
}

//采用单例模式初始化ca
func (m *TlsService) SetTLSConfig() error {
	if m._tlsConfig != nil {
		return nil
	}
	wechatCertPath := m._caPath + "/apiclient_cert.pem"
	wechatKeyPath := m._caPath + "/apiclient_key.pem"
	wechatCAPath := m._caPath + "/rootca.pem"

	// load cert
	cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	if err != nil {
		return err
	}
	// load root ca
	caData, err := ioutil.ReadFile(wechatCAPath)
	if err != nil {
		return err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)
	m._tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}
	return nil
}

//携带ca证书的安全请求
func (m *TlsService) SecurePost(url string, xmlContent []byte) (*http.Response, error) {
	if m._tlsConfig == nil {
		setErr := m.SetTLSConfig()
		if setErr != nil {
			return nil, setErr
		}
		if m._tlsConfig == nil {
			return nil, errors.New("配置密钥不存在")
		}
	}
	tlsConfig := m._tlsConfig
	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	return client.Post(
		url,
		"application/xml",
		bytes.NewBuffer(xmlContent))
}

func NewTlsService(p_path string) *TlsService {
	rst := &TlsService{
		_caPath: p_path,
	}
	return rst
}
