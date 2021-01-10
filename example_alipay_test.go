/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/10 18:16
*/

package alipay

import (
	"github.com/WenyXu/better-alipay-go/global"
	"github.com/WenyXu/better-alipay-go/m"
	"github.com/WenyXu/better-alipay-go/options"
	"os"
)

func ExampleDefault_basic() {
	// init a default client with a app configuration
	_ = Default(
		options.AppId(os.Getenv("APP_ID")),
		options.PrivateKey(os.Getenv("PrivateKey")),
		// Depended on you AppCert type
		// if you'd like load the cert form []byte
		// just use:
		// options.AppCertBytes()
		// if you already save cert sn at somewhere
		// just use:
		// options.AppCertSn()
		options.AppCertPath("./cert_file/appCertPublicKey.crt"),
		// similar to the AppCertPath
		// also provide:
		// options.RootCertBytes()
		// options.RootCertSn()
		options.RootCertPath("./cert_file/alipayRootCert.crt"),
		// similar to the AppCertPath
		// also provide:
		// options.PublicCertBytes()
		// options.PublicCertSn()
		options.PublicCertPath("./cert_file/alipayCertPublicKey_RSA2.crt"),
		options.Production(false),
		// or global.PKCS1
		options.PrivateKeyType(global.PKCS8),
		// or global.RSA
		options.SignType(global.RSA2),
	)
}

func ExampleDefault_advanced() {
	s := Default()

	// of course, you can using a struct instead
	// var resp AlipayTradeCreateResponse
	resp := make(map[string]interface{})
	_ = s.Request("alipay.trade.create", m.NewMap(func(param m.M) {
		param.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}),
		&resp,
		// dynamic configuration pre Request func
		options.AppId(os.Getenv("APP_ID")),
		options.PrivateKey(os.Getenv("PrivateKey")),
		options.AppCertPath("./cert_file/appCertPublicKey.crt"),
		options.RootCertPath("./cert_file/alipayRootCert.crt"),
		options.PublicCertPath("./cert_file/alipayCertPublicKey_RSA2.crt"),
		options.Production(false),
		options.PrivateKeyType(global.PKCS8),
		options.SignType(global.RSA2),
	)
}

func MakeAppConfig(yourConfig struct{ AppId string }) options.Option {
	return func(o *options.Options) {
		// modify your app config
		// using
		o.Config.AppId = yourConfig.AppId
	}
}

func ExampleDefault_advanced_2() {
	s := Default()

	// of course, you can using a struct instead
	// var resp AlipayTradeCreateResponse
	resp := make(map[string]interface{})
	_ = s.Request("alipay.trade.create", m.NewMap(func(param m.M) {
		param.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}),
		&resp,
		// dynamic configuration pre Request func
		MakeAppConfig(struct{ AppId string }{AppId: ""}),
	)
}
