/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/07 5:48
*/

/*
	Global constants
*/
package global

const (
	RSA2        = "RSA2"
	RSA         = "RSA"
	PKCS1       = "PKCS1"
	PKCS8       = "PKCS8"
	ContentType = "application/x-www-form-urlencoded;charset=utf-8"
	TimeLayout  = "2006-01-02 15:04:05"

	ServerUrlProduction  = "https://openapi.alipay.com/gateway.do"    // 正式环境请求地址
	ServerUrlDevelopment = "https://openapi.alipaydev.com/gateway.do" // 沙箱环境请求地址

	PublicAppAuthUrlProduction  = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"    // 正式环境授权登录地址
	PublicAppAuthUrlDevelopment = "https://openauth.alipaydev.com/oauth2/publicAppAuthorize.htm" // 沙箱环境授权登录地址

	AppToAppAuthUrlProduction  = "https://openauth.alipay.com/oauth2/appToAppAuth.htm"    // 正式环境第三方授权登录地址
	AppToAppAuthUrlDevelopment = "https://openauth.alipaydev.com/oauth2/appToAppAuth.htm" // 沙箱环境第三方授权登录地址

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

	// https://opendocs.alipay.com/apis/00tdbu
	AlipayEcoMyCarParkingOrderRefund = "alipay.eco.mycar.parking.order.refund"

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
)
