/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/11 1:53
*/

package entity

//
type FundTransToAccountTransferResponse struct {
	FundTransToAccountTransfer `json:"alipay_fund_trans_toaccount_transfer_response"`
	AlipayCertSn               string `json:"alipay_cert_sn"`
	Sign                       string `json:"sign"`
}

type FundTransToAccountTransfer struct {
	Common   `json:",inline"`
	OutBizNo string `json:"out_biz_no"`
	OrderId  string `json:"order_id"`
	PayDate  string `json:"pay_date"`
}

// FundTransUniTransferResponse
// https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.uni.transfer
type FundTransUniTransferResponse struct {
	FundTransUniTransfer `json:"alipay_fund_trans_uni_transfer_response"`
	AlipayCertSn         string `json:"alipay_cert_sn"`
	Sign                 string `json:"sign"`
}

type FundTransUniTransfer struct {
	Common         `json:",inline"`
	OutBizNo       string `json:"out_biz_no"`
	OrderId        string `json:"order_id"`
	PayFundOrderId string `json:"pay_fund_order_id"`
	Status         string `json:"status"`
	TransDate      string `json:"trans_date"`
}

// FundTransCommonQueryResponse
// https://opendocs.alipay.com/apis/api_28/alipay.fund.trans.common.query
type FundTransCommonQueryResponse struct {
	FundTransCommonQuery `json:"alipay_fund_trans_common_query_response"`
	AlipayCertSn         string `json:"alipay_cert_sn"`
	Sign                 string `json:"sign"`
}

type FundTransCommonQuery struct {
	Common           `json:",inline"`
	OrderId          string `json:"order_id"`
	PayFundOrderId   string `json:"pay_fund_order_id"`
	OutBizNo         string `json:"out_biz_no"`
	TransAmount      string `json:"trans_amount"`
	Status           string `json:"status"`
	PayDate          string `json:"pay_date"`
	ArrivalTimeEnd   string `json:"arrival_time_end"`
	OrderFee         string `json:"order_fee"`
	ErrorCode        string `json:"error_code"`
	FailReason       string `json:"fail_reason"`
	DeductBillInfo   string `json:"deduct_bill_info"`
	TransferBillInfo string `json:"transfer_bill_info"`
}

// FundAccountQueryResponse
// https://opendocs.alipay.com/apis/api_28/alipay.fund.account.query
type FundAccountQueryResponse struct {
	FundAccountQuery `json:"alipay_fund_account_query_response"`
	AlipayCertSn     string `json:"alipay_cert_sn"`
	Sign             string `json:"sign"`
}

type FundAccountQuery struct {
	Common          `json:",inline"`
	AvailableAmount string `json:"available_amount"`
	ExtCardInfo     *struct {
		CardNo       string `json:"card_no"`
		BankAccName  string `json:"bank_acc_name"`
		CardBranch   string `json:"card_branch"`
		CardBank     string `json:"card_bank"`
		CardLocation string `json:"card_location"`
		CardDeposit  string `json:"card_deposit"`
		Status       string `json:"status"`
	} `json:"ext_card_info"`
}
