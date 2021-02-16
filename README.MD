# Better Aliapy Go

[![Go Report Card](https://goreportcard.com/badge/github.com/WenyXu/better-alipay-go)](https://goreportcard.com/report/github.com/WenyXu/better-alipay-go)
[![GoDoc](https://pkg.go.dev/badge/github.com/WenyXu/better-alipay-go?status.svg)](https://pkg.go.dev/github.com/WenyXu/better-alipay-go?tab=doc)

A better Alipay SDK for Golang. With First-class function, without the tears 😢 [GoDoc](https://pkg.go.dev/github.com/WenyXu/better-alipay-go?tab=doc)
- Thread safe
- Hooks friendly
- Tracing friendly
- Dynamic configuration friendly
- Multiple configuration friendly

一个更好的支付宝（Alipay）SDK。函数优先，没有眼泪😢 [中文文档](http://better-alipay-go.terminal.im/)

- 多线程安全
- Hooks
- 链路跟踪
- 动态配置
- 多个配置


## Contents

- [Contents](#contents)
- [Installation](#installation)
- [Quick start](#quick-start)
- [Global Configuration](#global-configuration)
- [Dynamic Configuration](#dynamic-configuration)
  - [Basic](#basic-usage)
  - [Advanced](#advanced-usage)
- [Hooks](#hooks)
- [Tracing](#tracing)
- [Built-in Methods constant and Response structs](#built-in-methods-constant-and-response-structs)
- [Inspirited](#inspirited)


## Installation
To install Better-Alipay-Go package, you need to install Go and set your Go workspace first.

1. The first need [Go](https://golang.org/) installed, then you can use the below Go command to install it.

```sh
$ go get -u github.com/WenyXu/better-alipay-go
```

2. Import it in your code:

```go
import "github.com/WenyXu/better-alipay-go"
```

## Quick Start

```go
package main
import (
    alipay "github.com/WenyXu/better-alipay-go"
    "github.com/WenyXu/better-alipay-go/options"
    "github.com/WenyXu/better-alipay-go/global"
    "github.com/WenyXu/better-alipay-go/m"
    "os"
)
func main(){
	// init a default client with a app configuration
    s := alipay.Default(
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
	
    // of course, you can using a struct instead
    // var resp AlipayTradeCreateResponse
    resp := make(map[string]interface{})
    _ = s.Request("alipay.trade.create",m.NewMap(func(param m.M) {
          param.
              Set("subject", "网站测试支付").
              Set("buyer_id", "2088802095984694").
              Set("out_trade_no", "123456786543").
              Set("total_amount", "88.88"),
    }), &resp)

    // without biz-content request
    _ = s.Request(global.AlipaySystemOauthToken, m.NewMap(func(param m.M) {
        // set key value as public params 
    	param.
        	Set("grant_type", "authorization_code").
        	Set("code", "3a06216ac8f84b8c93507bb9774bWX11")
    }),
        &resp,
        options.SetMakeReqFunc(options.WithoutBizContentMakeReqFunc),
    )
	
}
    
```

## Global Configuration

You can use following functions to configure global configuration.

```go
package main

import (
	"github.com/WenyXu/better-alipay-go/options"
	"time"
)

func main() {
	// set Default Transport 
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	options.SetDefaultLocation(options.SetLocation(loc))
	
	// set Default Transport which implement http.RoundTripper interface
	transport := YourCustomTransport()
	options.SetDefaultTransport(transport)

	// set Default MakeRequestFunc which implement options.MakeRequestFunc func
	options.SetDefaultMakeReqFunc(yourMakeReqFunc)

	// set Default DecFunc which implement options.DecFunc func
	options.SetDefaultDecFunc(yourDecFunc)
	
	// set Default Logger which implement logger.Logger interface
	// built-in :
	// logger.NullLogger
	// logger.StdLogger
	options.SetDefaultLogger(yourLogger)
}

```

After above configuring, alipay.New / alipay.Default / options.newOptions will return new Options with your configured.

```go
// options.go 
func newOptions(opts ...Option) Options {
	opt := Options{
		Transport: DefaultTransport,
		Context:   context.Background(),
		MakeReq:   DefaultMakeReqFunc,
		Dec:       DefaultDecFunc,
		Logger:    DefaultLogger,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}
```

## Dynamic Configuration

If your application have multiple configurations, you can configure configuration dynamically.

When you call the Request func, it will make a copy form your current Options, then modify the Options. generally, it will be thread-safe.

### Basic Usage 

Configure per request
```go
package main
import (
    alipay "github.com/WenyXu/better-alipay-go"
    "github.com/WenyXu/better-alipay-go/options"
    "github.com/WenyXu/better-alipay-go/global"
    "github.com/WenyXu/better-alipay-go/m"
    "os"
)
func main(){
	// init a default client with a app configuration
    s := alipay.Default(
        options.AppId(os.Getenv("APP_ID")),
        options.PrivateKey(os.Getenv("PrivateKey")),
        options.AppCertPath("./cert_file/appCertPublicKey.crt"),
        options.RootCertPath("./cert_file/alipayRootCert.crt"),
        options.PublicCertPath("./cert_file/alipayCertPublicKey_RSA2.crt"),
        options.Production(false),
        options.PrivateKeyType(global.PKCS8),
        options.SignType(global.RSA2),
    )
	
    // of course, you can using a struct instead
    // var resp AlipayTradeCreateResponse
    resp := make(map[string]interface{})
	s.Request("alipay.trade.create",m.NewMap(func(param m.M) {
            param.
                Set("subject", "网站测试支付").
                Set("buyer_id", "2088802095984694").
                Set("out_trade_no", "123456786543").
                Set("total_amount", "88.88")
        }),
        &resp,
        // dynamic configuration
        options.AppAuthToken("your-app-auth-token"),
        options.AuthToken("your-auth-token"),
	)
}
```

### Advanced Usage
Configure with func which implement options.Option.
```go
    // implement options.Option
    // type Option func(*Options)
    func CustomOption(o *options.Options) {
        // modify your app config 
        // using 
        o.Config.AppId="whatever"
    }
    
    ...

    func main(){
        // use custom option in New or Default func 
		s := alipay.Default(
			...,
		    CustomOption
        )
        // or use custom option in Request of MakeParam func 
        s.Request("alipay.trade.create",m.NewMap(func(param m.M) {
              param.
              Set("subject", "网站测试支付").
              Set("buyer_id", "2088802095984694").
              Set("out_trade_no", "123456786543").
              Set("total_amount", "88.88")
          }),
          &resp,
          // your custom option
          CustomOption,
        )
    }   
    ...
    
    
```
Use a client without default app configuration and configure per request.
```go
package main
import (
    alipay "github.com/WenyXu/better-alipay-go"
    "github.com/WenyXu/better-alipay-go/options"
    "github.com/WenyXu/better-alipay-go/global"
    "github.com/WenyXu/better-alipay-go/m"
)

func main(){
	// init a empty client 
    s := alipay.Default()
    
    // of course, you can using a struct instead
    // var resp AlipayTradeCreateResponse
    resp := make(map[string]interface{})
	s.Request("alipay.trade.create",m.NewMap(func(param m.M) {
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

```

Load configuration form database or somewhere, and configure dynamically at per request.
```go
package main
import (
	alipay "github.com/WenyXu/better-alipay-go"
	"github.com/WenyXu/better-alipay-go/options"
	"github.com/WenyXu/better-alipay-go/m"
)

func MakeAppConfig(yourConfig ConfigEntity) options.Option {
	return func(o *options.Options) {
		// modify your app config 
		// using 
		o.Config.AppId=yourConfig.AppId
		...
	}
}
func main()  {
	// init a empty client 
	s := alipay.Default()

	// you config entity
	config:=ReadFormSomeWhere()
    
	// of course, you can using a struct instead
	// var resp AlipayTradeCreateResponse
	resp := make(map[string]interface{})
	s.Request("alipay.trade.create",m.NewMap(func(param m.M) {
              param.
                  Set("subject", "网站测试支付").
                  Set("buyer_id", "2088802095984694").
                  Set("out_trade_no", "123456786543").
                  Set("total_amount", "88.88")
	    }),
        &resp,
        // dynamic configuration pre Request func
        MakeAppConfig(config),
	)
    
}
```

## Hooks
In Options, we provide hooks which run before request started, and after response received, so you can do something like inject context, tracing's span etc. it is just similar to web middleware.
We provide a sample here, help you have a concept.

```go
// options.go
// DefaultBeforeFunc log the request body
func DefaultBeforeFunc(ctx context.Context, req *http.Request) context.Context {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        fmt.Printf("Read Request body with error: %s", err.Error())
        return ctx
    }
    req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
    fmt.Println(string(body))
    return ctx
}

// DefaultAfterFunc log the response body
func DefaultAfterFunc(ctx context.Context, resp *http.Response) context.Context {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Printf("Read Response body with error: %s", err.Error())
        return ctx
    }
    resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
    fmt.Println(string(body))
    return ctx
}

```

The Request fun will run before-hooks before the request started, and run after-hook after response received.

```go
//alipay.go
// Do Request
func (s service) Request(method string, param m.M, response interface{}, opts ...options.Option) (err error) {
    copyOpts := s.opts.Copy()
    
    ...
    
    // run before hooks before request started
    for _, f := range copyOpts.Before {
        ctx = f(ctx, req)
    }
    
    ...
    
    // do request
    resp, err := copyOpts.Transport.RoundTrip(req)
    
    
    ...
    
    // run after hooks after response received
    for _, f := range copyOpts.After {
        ctx = f(ctx, resp)
    }
    
    ...
}

```

## Tracing
You can inject context, when your call a Request func, and use hook to finish the tracing.

```go
package main

import (
	"context"
	alipay "github.com/WenyXu/better-alipay-go"
	"github.com/WenyXu/better-alipay-go/options"
	"github.com/WenyXu/better-alipay-go/m"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

func main() {
	s := alipay.Default()
	// get trace instance
	trace := yourTracingInstance()
	resp := make(map[string]interface{})
	s.Request("alipay.trade.create", m.NewMap(func(param m.M) {
		param.
			Set("subject", "网站测试支付").
			Set("buyer_id", "2088802095984694").
			Set("out_trade_no", "123456786543").
			Set("total_amount", "88.88")
	}),
	&resp,
	// inject your tracing context
	func(trace opentracing.Tracer) options.Option {
		return func(o *options.Options) {
			sp := opentracing.StartSpan("tarcing name")
			// injected
			o.Context = context.WithValue(o.Context,"span-key",sp)
		}
	}(trace),
	// handle span finish
	options.AppendAfterFunc(func(c context.Context, response *http.Response) context.Context {
            sp, ok := c.Get("span-key")
            if ok {
                // span finish
                sp.Finish()
            }
		    return c
	    }),
	)

}
```

## Built-in Methods constant and Response structs

Following methods:

```
        // https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.enterinfo.sync
	AlipayEcoMyCarParkingEnterInfoSync = "alipay.eco.mycar.parking.enterinfo.sync"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.exitinfo.sync
	AlipayEcoMyCarParkingExitInfoSync = "alipay.eco.mycar.parking.exitinfo.sync"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.vehicle.query
	AlipayEcoMyCarParkingVehicleQuery = "alipay.eco.mycar.parking.vehicle.query"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.order.sync
	AlipayEcoMyCarParkingOrderSync = "alipay.eco.mycar.parking.order.sync"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.order.update
	AlipayEcoMyCarParkingOrderUpdate = "alipay.eco.mycar.parking.order.update"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.config.set
	AlipayEcoMyCarParkingConfigSet = "alipay.eco.mycar.parking.config.set"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.update
	AlipayEcoMyCarParkingParkingLotInfoUpdate = "alipay.eco.mycar.parking.parkinglotinfo.update"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.create
	AlipayEcoMyCarParkingParkingLotInfoCreate = "alipay.eco.mycar.parking.parkinglotinfo.create"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.parkinglotinfo.query
	AlipayEcoMyCarParkingParkingLotInfoQuery = "alipay.eco.mycar.parking.parkinglotinfo.query"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.order.pay
	AlipayEcoMyCarParkingOrderPay = "alipay.eco.mycar.parking.order.pay"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.trade.order.query
	AlipayEcoMyCarTradeOrderQuery = "alipay.eco.mycar.trade.order.query"

	// https://opendocs.alipay.com/apis/api_19/alipay.eco.mycar.parking.agreement.query
	AlipayEcoMyCarParkingAgreement = "alipay.eco.mycar.parking.agreement.query"

	// https://opendocs.alipay.com/apis/api_9/alipay.user.info.auth
	AlipayUserInfoAuth = "alipay.user.info.auth"

	// https://opendocs.alipay.com/apis/api_9/alipay.system.oauth.token
	AlipaySystemOauthToken = "alipay.system.oauth.token"

	// https://opendocs.alipay.com/apis/api_9/alipay.open.auth.token.app
	AlipayOpenAuthTokenApp = "alipay.open.auth.token.app"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay
	AlipayTradePagePay = "alipay.trade.page.pay"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay
	AlipayTradeAppPay = "alipay.trade.app.pay"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay
	AlipayTradeWapPay = "alipay.trade.wap.pay"

	// https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.certify
	AlipayUserCertifyOpenCertify = "alipay.user.certify.open.certify"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.fastpay.refund.query
	AlipayTradeFastpayRefundQuery = "alipay.trade.fastpay.refund.query"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.order.settle
	AlipayTradeOrderSettle = "alipay.trade.order.settle"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.create
	AlipayTradeCreate = "alipay.trade.create"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.close
	AlipayTradeClose = "alipay.trade.close"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.cancel
	AlipayTradeCancel = "alipay.trade.cancel"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.refund
	AlipayTradeRefund = "alipay.trade.refund"

	// https://opendocs.alipay.com/apis/api_1/alipay.page.trade.refund
	AlipayTradePageRefund = "alipay.trade.page.refund"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.precreate
	AlipayTradePrecreate = "alipay.trade.precreate"

	// https://opendocs.alipay.com/apis/api_1/alipay.trade.query
	AlipayTradeQuery = "alipay.trade.query"

	// https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.toaccount.transfer
	AlipayFundTransToAccountTransfer = "alipay.fund.trans.toaccount.transfer"

	// https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.uni.transfer
	AlipayFundTransUniTransfer = "alipay.fund.trans.uni.transfer"

	// https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.common.query
	AlipayFundTransCommonQuery = "alipay.fund.trans.common.query"

	// https://opendocs.alipay.com/apis/api_28/alipay.fund.account.query
	AlipayFundAccountQuery = "alipay.fund.account.query"

	// https://opendocs.alipay.com/apis/api_2/alipay.user.info.share
	AlipayUserInfoShare = "alipay.user.info.share"

	// https://opendocs.alipay.com/apis/api_8/zhima.credit.score.get
	ZhimaCreditScoreGet = "zhima.credit.score.get"

	// https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.initialize
	AlipayUserCertifyOpenInitialize = "alipay.user.certify.open.initialize"

	// https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.query
	AlipayUserCertifyOpenQuery = "alipay.user.certify.open.query"
```

You can use Success() func of these structs to check if the response return successfully

```go
    ...
    var resp AlipayTradeCreate
    ...
    if resp.Success() {
    	....
    }   
```


## Inspirited
- [iGoogle-ink/gopay](https://github.com/iGoogle-ink/gopay)
- [smartwalle/alipay](https://github.com/smartwalle/alipay)