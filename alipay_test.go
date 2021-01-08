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
	s = DefaultClient(
		AppId(os.Getenv("APP_ID")),
		PrivateKey(os.Getenv("PrivateKey")),
		AppCertSnPath("./cert/appCertPublicKey.crt"),
		RootCertSnPath("./cert/alipayRootCert.crt"),
		PublicCertSnPath("./cert/alipayCertPublicKey_RSA2.crt"),
		Production(false),
		PrivateKeyType(PKCS8),
		SignType(RSA2),
	)
	fmt.Printf("%p\n", s.Options().logger)
	os.Exit(m.Run())
}
func TestService_Request(t *testing.T) {
	resp := make(map[string]interface{})
	err := s.Request("alipay.trade.create", NewMap(func(m M) {
		m.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}), &resp, SetAfterFunc(EmptyAfterFunc))

	_ = s.Request("alipay.trade.create", NewMap(func(m M) {
		m.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}), &resp)

	assert.Equal(t, err, nil)
	//fmt.Println(resp)

}
