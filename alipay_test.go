/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 19:56
*/

package alipay

import (
	"fmt"
	"os"
	"testing"

	"github.com/WenyXu/better-alipay-go/global"
	_map "github.com/WenyXu/better-alipay-go/m"
	"github.com/WenyXu/better-alipay-go/options"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var (
	s Service
)

func TestMain(m *testing.M) {
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Println(err)
	}
	s = Default(
		options.AppId(os.Getenv("APP_ID")),
		options.PrivateKey(os.Getenv("PrivateKey")),
		options.AppCertPath("./cert_file/appCertPublicKey.crt"),
		options.RootCertPath("./cert_file/alipayRootCert.crt"),
		options.PublicCertPath("./cert_file/alipayCertPublicKey_RSA2.crt"),
		options.Production(false),
		options.PrivateKeyType(global.PKCS8),
		options.SignType(global.RSA2),
	)
	fmt.Printf("%p\n", s.Options().Logger)
	os.Exit(m.Run())
}

func TestService_Request(t *testing.T) {
	resp := make(map[string]interface{})
	err := s.Request("alipay.trade.create", _map.NewMap(func(m _map.M) {
		m.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}), &resp, options.SetAfterFunc(options.EmptyAfterFunc))

	_ = s.Request("alipay.trade.create", _map.NewMap(func(m _map.M) {
		m.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}), &resp)

	assert.Equal(t, err, nil)
	//fmt.Println(resp)

}
