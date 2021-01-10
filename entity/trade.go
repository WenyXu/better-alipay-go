/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/10 17:34
*/

package entity

type TradeCreate struct {
	Common     `json:",inline"`
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}

// TradeCreateResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.create
type TradeCreateResponse struct {
	TradeCreate `json:"alipay_trade_create_response"`
	Sign        string `json:"sign"`
}

type TradeFundBill struct {
	FundChannel string `json:"fund_channel"`
	BankCode    string `json:"bank_code"`
	Amount      string `json:"amount"`
	RealAmount  string `json:"real_amount"`
	FundType    string `json:"fund_type"`
}

type TradeRefund struct {
	Common                  `json:",inline"`
	TradeNo                 string           `json:"trade_no"`
	OutTradeNo              string           `json:"out_trade_no"`
	BuyerLogonId            string           `json:"buyer_logon_id"`
	FundChange              string           `json:"fund_change"`
	RefundFee               string           `json:"refund_fee"`
	RefundCurrency          string           `json:"refund_currency"`
	GmtRefundPay            string           `json:"gmt_refund_pay"`
	RefundDetailItemList    []*TradeFundBill `json:"refund_detail_item_list"`
	StoreName               string           `json:"store_name"`
	BuyerUserId             string           `json:"buyer_user_id"`
	RefundPresetPaytoolList []*struct {
		Amount         []string `json:"amount"`
		AssertTypeCode string   `json:"assert_type_code"`
	} `json:"refund_preset_paytool_list"`
	RefundSettlementId           string `json:"refund_settlement_id"`
	PresentRefundBuyerAmount     string `json:"present_refund_buyer_amount"`
	PresentRefundDiscountAmount  string `json:"present_refund_discount_amount"`
	PresentRefundMdiscountAmount string `json:"present_refund_mdiscount_amount"`
}

// TradeRefundResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.refund
type TradeRefundResponse struct {
	TradeRefund  `json:"alipay_trade_refund_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	SignData     string `json:"-"`
	Sign         string `json:"sign"`
}

type TradePageRefund struct {
	Common       `json:",inline"`
	TradeNo      string `json:"trade_no"`
	OutTradeNo   string `json:"out_trade_no"`
	OutRequestNo string `json:"out_request_no"`
	RefundAmount string `json:"refund_amount"`
}

// TradePageRefundResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.page.refund
type TradePageRefundResponse struct {
	TradePageRefund `json:"alipay_trade_page_refund_response"`
	AlipayCertSn    string `json:"alipay_cert_sn"`
	SignData        string `json:"-"`
	Sign            string `json:"sign"`
}
