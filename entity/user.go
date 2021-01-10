/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/10 18:04
*/

package entity

type UserInfoAuthResponse struct {
	UserInfoAuth struct {
		Common `json:",inline"`
	} `json:"alipay_user_info_auth_response,omitempty"`
	AlipayCertSn string `json:"alipay_cert_sn,omitempty"`
	Sign         string `json:"sign"`
}
