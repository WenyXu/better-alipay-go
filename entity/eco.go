/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/10 17:35
*/

package entity

type EnterInfo struct {
	Common    `json:",inline"`
	CarStatus string `json:"car_status"`
	SerialNo  string `json:"serial_no"`
}

type EnterInfoResponse struct {
	EnterInfo `json:"alipay_eco_mycar_parking_enterinfo_sync_response"`
	Sign      string `json:"sign"`
}

type ExitInfoResponse struct {
	Common `json:"alipay_eco_mycar_parking_exitinfo_sync_response"`
	Sign   string `json:"sign"`
}

type VehicleQuery struct {
	Common `json:",inline"`
	CarNum string `json:"car_number"`
}

type VehicleQueryResponse struct {
	VehicleQuery `json:"alipay_eco_mycar_parking_vehicle_query_response"`
	Sign         string `json:"sign"`
}

type OrderSyncResponse struct {
	Common `json:"alipay_eco_mycar_parking_order_sync_response"`
	Sign   string `json:"sign"`
}

type OrderUpdateResponse struct {
	Common `json:"alipay_eco_mycar_parking_order_update_response"`
	Sign   string `json:"sign"`
}

type ConfigQuery struct {
	Common               `json:",inline"`
	MerchantName         string `json:"merchant_name"`
	MerchantServicePhone string `json:"merchant_service_phone"`
	AccountNo            string `json:"account_no"`
	InterfaceInfoList    struct {
		InterfaceName string `json:"interface_name"`
		InterfaceType string `json:"interface_type"`
		InterfaceUrl  string `json:"interface_url"`
	} `json:"interface_info_list"`
}

type ConfigQueryResponse struct {
	ConfigQuery `json:"alipay_eco_mycar_parking_config_query_response"`
	Sign        string `json:"sign"`
}

type ConfigSetResponse struct {
	Common `json:"alipay_eco_mycar_parking_config_set_response"`
	Sign   string `json:"sign"`
}

type ParkingInfoUpdateResponse struct {
	Common `json:"alipay_eco_mycar_parking_parkinglotinfo_update_response"`
	Sign   string `json:"sign"`
}

type ParkingInfoCreate struct {
	Common          `json:",inline"`
	AlipayParkingId string `json:"parking_id"`
}

type ParkingInfoCreateResponse struct {
	ParkingInfoCreate `json:"alipay_eco_mycar_parking_parkinglotinfo_create_response"`
	Sign              string `json:"sign"`
}

type ParkingOrderPay struct {
	Common             `json:",inline"`
	UserID             string `json:"user_id"`
	TradeNo            string `json:"trade_no"`
	OutTradeNo         string `json:"out_trade_no"`
	TotalFee           string `json:"total_fee"`
	GmtPayment         string `json:"gmt_payment"`
	FundBillList       string `json:"fund_bill_list"`
	AdvanceAmount      string `json:"advance_amount"`
	AlipayRepaymentURL string `json:"alipay_repayment_url"`
}

type ParkingOrderPayResponse struct {
	ParkingOrderPay `json:"alipay_eco_mycar_parking_order_pay_response"`
	Sign            string `json:"sign"`
}

type ParkingOrderQueryRequest struct {
	BizTradeNo    string `json:"biz_trade_no"`
	TradeNo       string `json:"trade_no"`
	OutBizTradeNo string `json:"out_biz_trade_no"`
}

type ParkingOrderQuery struct {
	BuyerID           string  `json:"buyer_id"`
	BuyerLogonID      string  `json:"buyer_logon_id"`
	BizTradeNo        string  `json:"biz_trade_no"`
	TradeNo           string  `json:"trade_no"`
	ShopID            int64   `json:"shop_id"`
	TradeType         int     `json:"trade_type"`
	Subject           string  `json:"subject"`
	Summary           string  `json:"summary"`
	TotalFee          float64 `json:"total_fee"`
	TradeStatus       int     `json:"trade_status"`
	GmtPayment        string  `json:"gmt_payment"`
	GmtPaymentSuccess string  `json:"gmt_payment_success"`
	GmtRefund         string  `json:"gmt_refund"`
	GmtRefundSuccess  string  `json:"gmt_refund_success"`
	SendBackFee       float64 `json:"send_back_fee"`
	GmtClosed         string  `json:"gmt_closed"`
	GmtCreated        string  `json:"gmt_created"`
	GmtUpdated        string  `json:"gmt_updated"`
	OutBizTradeNo     string  `json:"out_biz_trade_no"`
	Common            `json:",inline"`
}

type ParkingOrderQueryResponse struct {
	ParkingOrderQuery `json:"alipay_eco_mycar_trade_order_query_response"`
	Sign              string `json:"sign"`
}

type AgreementQuery struct {
	Common          `json:",inline"`
	AgreementStatus string `json:"agreement_status"`
	UserAdvanceInfo struct {
		ConsultResponse               string `json:"consult_result"`
		UserAlipayParkingAllowAdvance bool   `json:"user_alipay_parking_allow_advance"`
		UserWaitRepaymentOrderCount   int    `json:"user_wait_repayment_order_count"`
		UserWaitRepaymentAmount       string `json:"user_wait_repayment_amount"`
	} `json:"user_advance_info"`
	AdvanceStatus string `json:"advance_status"`
	ExpireTime    string `json:"expire_time"`
}

type AgreementQueryResponse struct {
	AgreementQuery `json:"alipay_eco_mycar_parking_agreement_query_response"`
	Sign           string `json:"sign"`
}

type ParkingInfoQuery struct {
	Common                   `json:",inline"`
	OutParkingID             string `json:"out_parking_id"`
	ParkingID                string `json:"parking_id"`
	ParkingAddress           string `json:"parking_address"`
	ParkingLotType           string `json:"parking_lot_type"`
	ParkingPoiid             string `json:"parking_poiid"`
	ParkingMobile            string `json:"parking_mobile"`
	PayType                  string `json:"pay_type"`
	ParkingName              string `json:"parking_name"`
	MchntID                  string `json:"mchnt_id"`
	ShopingmallID            string `json:"shopingmall_id"`
	ParkingFeeDescription    string `json:"parking_fee_description"`
	TimeOut                  string `json:"time_out"`
	AgentID                  string `json:"agent_id"`
	ParkingLongitude         string `json:"parking_longitude"`
	ParkingLatitude          string `json:"parking_latitude"`
	MapPoiName               string `json:"map_poi_name"`
	MapPoiAddress            string `json:"map_poi_address"`
	ProvinceID               string `json:"province_id"`
	CityID                   string `json:"city_id"`
	AddressID                string `json:"address_id"`
	ParkingFeeDescriptionImg string `json:"parking_fee_description_img"`
}

type ParkingInfoQueryResponse struct {
	ParkingInfoQuery `json:"alipay_eco_mycar_parking_parkinglotinfo_query_response"`
	Sign             string `json:"sign"`
}

type OpenAuthToken struct {
	Common          `json:",inline"`
	UserID          string `json:"user_id"`
	AuthAppID       string `json:"auth_app_id"`
	AppAuthToken    string `json:"app_auth_token"`
	AppRefreshToken string `json:"app_refresh_token"`
	ExpiresIn       string `json:"expires_in"`
	ReExpiresIn     string `json:"re_expires_in"`
}

type OpenAuthTokenResponse struct {
	OpenAuthToken `json:"alipay_open_auth_token_app_response"`
	Sign          string `json:"sign"`
}
