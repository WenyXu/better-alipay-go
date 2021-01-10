/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/11 1:54
*/

package entity

// ZhimaCreditScoreGetResponse
// https://opendocs.alipay.com/apis/api_8/zhima.credit.score.get
type ZhimaCreditScoreGetResponse struct {
	ZhimaCreditScoreGet `json:"zhima_credit_score_get_response"`
	AlipayCertSn        string `json:"alipay_cert_sn"`
	Sign                string `json:"sign"`
}

type ZhimaCreditScoreGet struct {
	Common  `json:",inline"`
	BizNo   string `json:"biz_no"`
	ZmScore string `json:"zm_score"`
}
