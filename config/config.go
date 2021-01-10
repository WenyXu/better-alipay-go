/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 18:06
*/

/*
	Module for configure the sdk configurations
*/
package config

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/WenyXu/better-alipay-go/cert"
	"github.com/WenyXu/better-alipay-go/global"
	m "github.com/WenyXu/better-alipay-go/m"
	"github.com/WenyXu/better-alipay-go/sign"
)

type Config struct {
	Loc                *time.Location
	AppId              string
	PrivateKey         string
	PrivateKeyType     string
	AppCertSN          string
	AliPayPublicCertSN string
	AliPayRootCertSN   string
	SignType           string
	ReturnUrl          string
	NotifyUrl          string
	Charset            string
	Format             string
	Version            string
	AppAuthToken       string
	AuthToken          string
	Production         bool
}

// Url get alipay gateway endpoint
func (c Config) Url() string {
	if c.Production {
		return global.ServerUrlProduction
	}
	return global.ServerUrlDevelopment
}

// SetPublicParam Set Config's Public Param to map[string]interface{}
/*
	target.
	Set("app_id", c.AppId).
	Set("format", c.Format).
	Set("charset", c.Charset).
	Set("sign_type", c.SignType).
	Set("timestamp", time.Now().In(c.Loc).Format(global.TimeLayout))
*/
func SetPublicParam(c Config) m.MakeMapEndpoint {
	return func(target m.M) error {
		target.
			Set("app_id", c.AppId).
			Set("format", c.Format).
			Set("charset", c.Charset).
			Set("sign_type", c.SignType).
			Set("timestamp", time.Now().In(c.Loc).Format(global.TimeLayout))
		return nil
	}
}

// SetOptionalParam Set Config's Optional Param to map[string]interface{}
/*
	target.
	Set("app_cert_sn", c.AppCertSN, opt).
	Set("alipay_root_cert_sn", c.AliPayRootCertSN, opt).
	Set("return_url", c.ReturnUrl, opt).
	Set("notify_url", c.NotifyUrl, opt).
	Set("app_auth_token", c.AppAuthToken, opt).
	Set("auth_token", c.AuthToken, opt)
*/
func SetOptionalParam(c Config) m.MakeMapEndpoint {
	return func(target m.M) error {
		opt := m.MapOptions{}
		opt.SetIgnoreEmptyString(true)
		// add public params
		target.
			Set("app_cert_sn", c.AppCertSN, opt).
			Set("alipay_root_cert_sn", c.AliPayRootCertSN, opt).
			Set("return_url", c.ReturnUrl, opt).
			Set("notify_url", c.NotifyUrl, opt).
			Set("app_auth_token", c.AppAuthToken, opt).
			Set("auth_token", c.AuthToken, opt)
		return nil
	}
}

// SetMethod set method
func SetMethod(method string) m.MakeMapEndpoint {
	return func(target m.M) error {
		target.Set("method", method)
		return nil
	}
}

// SetBizContent set bizContent
func SetBizContent(source m.M) m.MakeMapEndpoint {
	return func(target m.M) error {
		bytes, err := json.Marshal(source)
		if err != nil {
			return fmt.Errorf("json.Marshal：%w", err)
		}
		target["biz_content"] = string(bytes)
		return nil
	}
}

// SignParam sign current params, also set the sign result into map[string]interface{}
func SignParam(c Config) m.MakeMapEndpoint {
	return func(target m.M) error {
		s, err := sign.Sign(target, c.SignType, cert.LoadPrivateKeyFormString(c.PrivateKeyType, c.PrivateKey))
		if err != nil {
			return err
		}
		target.Set("sign", s)
		return nil
	}
}
