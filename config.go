/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 18:06
*/

package alipay

import (
	"time"
)

type Config struct {
	loc                *time.Location
	appId              string
	privateKey         string
	privateKeyType     string
	appCertSN          string
	aliPayPublicCertSN string
	aliPayRootCertSN   string
	signType           string
	returnUrl          string
	notifyUrl          string
	charset            string
	format             string
	version            string
	appAuthToken       string
	authToken          string
	production         bool
}

func (c Config) Url() string {
	if c.production {
		return ServerUrlProduction
	}
	return ServerUrlDevelopment
}
