/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/10 18:04
*/

package entity

type UserInfoAuthResponse struct {
	UserInfoAuth struct {
		Common `json:",inline"`
	} `json:"alipay_user_info_auth_response"`
	AlipayCertSn string `json:"alipay_cert_sn"`
	Sign         string `json:"sign"`
}

// UserInfoShareResponse
// https://opendocs.alipay.com/apis/api_2/alipay.user.info.share
type UserInfoShareResponse struct {
	UserInfoShare `json:"alipay_user_info_share_response"`
	AlipayCertSn  string `json:"alipay_cert_sn"`
	Sign          string `json:"sign"`
}

type UserInfoShare struct {
	Common             `json:",inline"`
	UserId             string `json:"user_id"`
	Avatar             string `json:"avatar"`
	Province           string `json:"province"`
	City               string `json:"city"`
	NickName           string `json:"nick_name"`
	IsStudentCertified string `json:"is_student_certified"`
	UserType           string `json:"user_type"`
	UserStatus         string `json:"user_status"`
	IsCertified        string `json:"is_certified"`
	Gender             string `json:"gender"`
}

// UserCertifyOpenInitResponse
// https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.initialize
type UserCertifyOpenInitializeResponse struct {
	UserCertifyOpenInitialize `json:"alipay_user_certify_open_initialize_response"`
	AlipayCertSn              string `json:"alipay_cert_sn"`
	Sign                      string `json:"sign"`
}

type UserCertifyOpenInitialize struct {
	Common    `json:",inline"`
	CertifyId string `json:"certify_id"`
}

// UserCertifyOpenCertifyResponse
// https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.certify
type UserCertifyOpenCertifyResponse struct {
	UserCertifyOpenCertify `json:"alipay_user_certify_open_certify_response"`
	AlipayCertSn           string `json:"alipay_cert_sn"`
	Sign                   string `json:"sign"`
}

type UserCertifyOpenCertify struct {
	Common `json:",inline"`
}

// UserCertifyOpenQueryResponse
// https://opendocs.alipay.com/apis/api_2/alipay.user.certify.open.query
type UserCertifyOpenQueryResponse struct {
	UserCertifyOpenQuery `json:"alipay_user_certify_open_query_response"`
	AlipayCertSn         string `json:"alipay_cert_sn"`
	Sign                 string `json:"sign"`
}

type UserCertifyOpenQuery struct {
	Common       `json:",inline"`
	Passed       string `json:"passed"`
	IdentityInfo string `json:"identity_info"`
	MaterialInfo string `json:"material_info"`
}
