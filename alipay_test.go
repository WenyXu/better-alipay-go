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

	"github.com/WenyXu/better-alipay-go/entity"

	"github.com/WenyXu/better-alipay-go/global"
	"github.com/WenyXu/better-alipay-go/m"
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
	err := s.Request("alipay.trade.create", m.NewMap(func(m m.M) {
		m.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}), &resp, options.SetAfterFunc(options.EmptyAfterFunc))

	_ = s.Request("alipay.trade.create", m.NewMap(func(m m.M) {
		m.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}), &resp)

	assert.Equal(t, err, nil)
	//fmt.Println(resp)
}

func TestService_Request_Trade_Page_Pay(t *testing.T) {
	// same as
	// alipay.trade.wap.pay
	// alipay.trade.page.pay
	// alipay.user.certify.open.certify
	data, err := s.MakeParam(global.AlipayTradePagePay, m.NewMap(func(param m.M) {
		param.Set("subject", "网站测试支付").
			Set("product_code", "FAST_INSTANT_TRADE_PAY").
			Set("out_trade_no", "123456789").
			Set("total_amount", "88.88")
	}))
	assert.Equal(t, nil, err)
	url := s.Options().Config.Url() + "?" + m.FormatURLParam(data)

	// page pay url
	fmt.Println(url)
}

func TestService_Request_App_Pay(t *testing.T) {
	// alipay.trade.app.pay
	data, err := s.MakeParam(global.AlipayTradeAppPay, m.NewMap(func(param m.M) {
		param.Set("subject", "app支付").
			Set("out_trade_no", "123456789").
			Set("total_amount", "88.88")
	}))
	assert.Equal(t, nil, err)
	// output
	fmt.Println(data)
}

func TestService_Request_AlipaySystemOauthToken(t *testing.T) {
	var resp entity.AlipaySystemOauthTokenResponse
	err := s.Request(global.AlipaySystemOauthToken, m.NewMap(func(m m.M) {
		m.Set("grant_type", "authorization_code").Set("code", "3a06216ac8f84b8c93507bb9774bWX11")
	}), &resp,
		options.SetMakeReqFunc(
			options.WithoutBizContentMakeReqFunc,
		),
	)
	assert.Equal(t, nil, err)
	fmt.Println(resp)

	var r2 map[string]interface{}
	err = s.Request("alipay.trade.create", m.NewMap(func(m m.M) {
		m.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}), &r2)
	assert.Equal(t, nil, err)
	fmt.Println(r2)
}
