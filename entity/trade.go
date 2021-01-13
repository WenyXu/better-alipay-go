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
	Sign            string `json:"sign"`
}

// TradePayResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.pay
type TradePayResponse struct {
	TradePay `json:"alipay_trade_pay_response"`
	Sign     string `json:"sign"`
}

type TradePay struct {
	Common          `json:",inline"`
	TradeNo         string `json:"trade_no"`
	OutTradeNo      string `json:"out_trade_no"`
	BuyerLogonId    string `json:"buyer_logon_id"`
	SettleAmount    string `json:"settle_amount"`
	PayCurrency     string `json:"pay_currency"`
	PayAmount       string `json:"pay_amount"`
	SettleTransRate string `json:"settle_trans_rate"`
	TransPayRate    string `json:"trans_pay_rate"`
	TotalAmount     string `json:"total_amount"`
	TransCurrency   string `json:"trans_currency"`
	SettleCurrency  string `json:"settle_currency"`
	ReceiptAmount   string `json:"receipt_amount"`
	BuyerPayAmount  string `json:"buyer_pay_amount"`
	PointAmount     string `json:"point_amount"`
	InvoiceAmount   string `json:"invoice_amount"`
	GmtPayment      string `json:"gmt_payment"`
	FundBillList    []*struct {
		FundChannel string `json:"fund_channel"`
		BankCode    string `json:"bank_code"`
		Amount      string `json:"amount"`
		RealAmount  string `json:"real_amount"`
	} `json:"fund_bill_list"`
	CardBalance         string `json:"card_balance"`
	StoreName           string `json:"store_name"`
	BuyerUserId         string `json:"buyer_user_id"`
	DiscountGoodsDetail string `json:"discount_goods_detail"`
	VoucherDetailList   []*struct {
		Id                         string `json:"id"`
		Name                       string `json:"name"`
		Type                       string `json:"type"`
		Amount                     string `json:"amount"`
		MerchantContribute         string `json:"merchant_contribute"`
		OtherContribute            string `json:"other_contribute"`
		Memo                       string `json:"memo"`
		TemplateId                 string `json:"template_id"`
		PurchaseBuyerContribute    string `json:"purchase_buyer_contribute"`
		PurchaseMerchantContribute string `json:"purchase_merchant_contribute"`
		PurchaseAntContribute      string `json:"purchase_ant_contribute"`
	} `json:"voucher_detail_list"`
	AdvanceAmount    string `json:"advance_amount"`
	AuthTradePayMode string `json:"auth_trade_pay_mode"`
	ChargeAmount     string `json:"charge_amount"`
	ChargeFlags      string `json:"charge_flags"`
	SettlementId     string `json:"settlement_id"`
	BusinessParams   string `json:"business_params"`
	BuyerUserType    string `json:"buyer_user_type"`
	MdiscountAmount  string `json:"mdiscount_amount"`
	DiscountAmount   string `json:"discount_amount"`
	BuyerUserName    string `json:"buyer_user_name"`
}

// TradeQueryResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.query
type TradeQueryResponse struct {
	TradeQuery `json:"alipay_trade_query_response"`
	Sign       string `json:"sign"`
}

type TradeQuery struct {
	Common          `json:",inline"`
	TradeNo         string `json:"trade_no"`
	OutTradeNo      string `json:"out_trade_no"`
	BuyerLogonId    string `json:"buyer_logon_id"`
	TradeStatus     string `json:"trade_status"`
	TotalAmount     string `json:"total_amount"`
	TransCurrency   string `json:"trans_currency"`
	SettleCurrency  string `json:"settle_currency"`
	SettleAmount    string `json:"settle_amount"`
	PayCurrency     string `json:"pay_currency"`
	PayAmount       string `json:"pay_amount"`
	SettleTransRate string `json:"settle_trans_rate"`
	TransPayRate    string `json:"trans_pay_rate"`
	BuyerPayAmount  string `json:"buyer_pay_amount"`
	PointAmount     string `json:"point_amount"`
	InvoiceAmount   string `json:"invoice_amount"`
	SendPayDate     string `json:"send_pay_date"`
	ReceiptAmount   string `json:"receipt_amount"`
	StoreId         string `json:"store_id"`
	TerminalId      string `json:"terminal_id"`
	FundBillList    []*struct {
		FundChannel string `json:"fund_channel"`
		BankCode    string `json:"bank_code"`
		Amount      string `json:"amount"`
		RealAmount  string `json:"real_amount"`
	} `json:"fund_bill_list"`
	StoreName       string `json:"store_name"`
	BuyerUserId     string `json:"buyer_user_id"`
	ChargeAmount    string `json:"charge_amount"`
	ChargeFlags     string `json:"charge_flags"`
	SettlementId    string `json:"settlement_id"`
	TradeSettleInfo *struct {
		TradeSettleDetailList []*struct {
			OperationType     string `json:"operation_type"`
			OperationSerialNo string `json:"operation_serial_no"`
			OperationDt       string `json:"operation_dt"`
			TransOut          string `json:"trans_out"`
			TransIn           string `json:"trans_in"`
			Amount            string `json:"amount"`
		} `json:"trade_settle_detail_list"`
	} `json:"trade_settle_info"`
	AuthTradePayMode    string `json:"auth_trade_pay_mode"`
	BuyerUserType       string `json:"buyer_user_type"`
	MdiscountAmount     string `json:"mdiscount_amount"`
	DiscountAmount      string `json:"discount_amount"`
	BuyerUserName       string `json:"buyer_user_name"`
	Subject             string `json:"subject"`
	Body                string `json:"body"`
	AlipaySubMerchantId string `json:"alipay_sub_merchant_id"`
	ExtInfos            string `json:"ext_infos"`
}

// TradeCloseResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.close
type TradeCloseResponse struct {
	TradeClose `json:"alipay_trade_close_response"`
	Sign       string `json:"sign"`
}

type TradeClose struct {
	Common     `json:",inline"`
	TradeNo    string `json:"trade_no"`
	OutTradeNo string `json:"out_trade_no"`
}

// TradeCancelResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.cancel
type TradeCancelResponse struct {
	CancelResponse `json:"alipay_trade_cancel_response"`
	AlipayCertSn   string `json:"alipay_cert_sn"`
	Sign           string `json:"sign"`
}

type CancelResponse struct {
	Common             `json:",inline"`
	TradeNo            string `json:"trade_no"`
	OutTradeNo         string `json:"out_trade_no"`
	RetryFlag          string `json:"retry_flag"`
	Action             string `json:"action"`
	GmtRefundPay       string `json:"gmt_refund_pay"`
	RefundSettlementId string `json:"refund_settlement_id"`
}

// TradeFastpayRefundQueryResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.fastpay.refund.query
type TradeFastpayRefundQueryResponse struct {
	TradeFastpayRefundQuery `json:"alipay_trade_fastpay_refund_query_response"`
	AlipayCertSn            string `json:"alipay_cert_sn"`
	Sign                    string `json:"sign"`
}

type TradeFastpayRefundQuery struct {
	Common         `json:",inline"`
	TradeNo        string `json:"trade_no"`
	OutTradeNo     string `json:"out_trade_no"`
	OutRequestNo   string `json:"out_request_no"`
	RefundReason   string `json:"refund_reason"`
	TotalAmount    string `json:"total_amount"`
	RefundAmount   string `json:"refund_amount"`
	RefundRoyaltys []*struct {
		RefundAmount  string `json:"refund_amount"`
		RoyaltyType   string `json:"royalty_type"`
		ResultCode    string `json:"result_code"`
		TransOut      string `json:"trans_out"`
		TransOutEmail string `json:"trans_out_email"`
		TransIn       string `json:"trans_in"`
		TransInEmail  string `json:"trans_in_email"`
	} `json:"refund_royaltys"`
	GmtRefundPay                 string           `json:"gmt_refund_pay"`
	RefundDetailItemList         []*TradeFundBill `json:"refund_detail_item_list"`
	SendBackFee                  string           `json:"send_back_fee"`
	RefundSettlementId           string           `json:"refund_settlement_id"`
	PresentRefundBuyerAmount     string           `json:"present_refund_buyer_amount"`
	PresentRefundDiscountAmount  string           `json:"present_refund_discount_amount"`
	PresentRefundMdiscountAmount string           `json:"present_refund_mdiscount_amount"`
}

// TradeOrderSettleResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.order.settle
type TradeOrderSettleResponse struct {
	TradeOrderSettle `json:"alipay_trade_order_settle_response"`
	AlipayCertSn     string `json:"alipay_cert_sn"`
	Sign             string `json:"sign"`
}
type TradeOrderSettle struct {
	Common  `json:",inline"`
	TradeNo string `json:"trade_no"`
}

// TradePrecreateResponse
// https://opendocs.alipay.com/apis/api_1/alipay.trade.precreate
type TradePrecreateResponse struct {
	TradePrecreate `json:"alipay_trade_precreate_response"`
	NullResponse   *ErrorResponse `json:"null_response"`
	AlipayCertSn   string         `json:"alipay_cert_sn"`
	Sign           string         `json:"sign"`
}

type TradePrecreate struct {
	Common     `json:",inline"`
	OutTradeNo string `json:"out_trade_no"`
	QrCode     string `json:"qr_code"`
}
