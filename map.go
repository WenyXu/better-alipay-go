/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 17:30
*/

package alipay

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

type M map[string]interface{}

type MapOptions struct {
	IgnoreEmptyString *bool
}

func NewMap(f func(m M)) M {
	m := make(M)
	f(m)
	return m
}
func (o *MapOptions) SetIgnoreEmptyString(input bool) {
	o.IgnoreEmptyString = &input
}

func mergeOptions(opts ...MapOptions) MapOptions {
	merged := MapOptions{}
	for _, o := range opts {
		if o.IgnoreEmptyString != nil {
			merged.IgnoreEmptyString = o.IgnoreEmptyString
		}
	}
	return merged
}

func (m M) Set(key string, value interface{}, opts ...MapOptions) M {
	mergedOption := mergeOptions(opts...)
	switch value.(type) {
	case func(map[string]interface{}):
		value.(func(map[string]interface{}))(m)
	case string:
		if mergedOption.IgnoreEmptyString != nil {
			if *mergedOption.IgnoreEmptyString == true {
				if value != "" {
					m[key] = value
				}
			}
		} else {
			// default not ignore empty string
			m[key] = value
		}
	default:
		m[key] = value
	}
	return m
}

func GetMapValueByKey(m map[string]interface{}, key string) (string, error) {
	if m != nil {
		if value, ok := m[key]; ok {
			switch value.(type) {
			case string:
				return value.(string), nil
			default:
				bytes, err := json.Marshal(value)
				if err != nil {
					return "", fmt.Errorf(err.Error())
				}
				return string(bytes), nil
			}
		}
	}
	return "", fmt.Errorf("input map is nil")
}

func EncodeMapParams(m map[string]interface{}) string {
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v, err := GetMapValueByKey(m, k); err == nil {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}

type MakeMapEndpoint func(M) (M, error)

func CombineMakeMapEndpointFunc(endpoint ...MakeMapEndpoint) (map[string]interface{}, error) {
	//var values = url.Values{}
	//values.Add("method", method)
	target := make(map[string]interface{})
	var errs []error
	for _, e := range endpoint {
		_, err := e(target)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		return target, FormatErrors(errs...)
	}
	return target, nil
}

func SetPublicParam(c Config) MakeMapEndpoint {
	return func(target M) (M, error) {
		target.
			Set("app_id", c.appId).
			Set("format", c.format).
			Set("charset", c.charset).
			Set("sign_type", c.signType).
			Set("timestamp", time.Now().In(c.loc).Format(TimeLayout))
		return target, nil
	}
}

func SetOptionsParam(c Config) MakeMapEndpoint {
	return func(target M) (M, error) {
		opt := MapOptions{}
		opt.SetIgnoreEmptyString(true)
		// add public params
		target.
			Set("app_cert_sn", c.appCertSN, opt).
			Set("alipay_root_cert_sn", c.aliPayRootCertSN, opt).
			Set("return_url", c.returnUrl, opt).
			Set("notify_url", c.notifyUrl, opt).
			Set("app_auth_token", c.appAuthToken, opt).
			Set("auth_token", c.authToken, opt)
		return target, nil
	}
}

func SetMethod(method string) MakeMapEndpoint {
	return func(target M) (M, error) {
		target.Set("method", method)
		return target, nil
	}
}

func SetBizContent(source M) MakeMapEndpoint {
	return func(target M) (M, error) {
		bytes, err := json.Marshal(source)
		if err != nil {
			return target, fmt.Errorf("json.Marshal：%w", err)
		}
		target["biz_content"] = string(bytes)
		return target, nil
	}
}

func SignParam(c Config) MakeMapEndpoint {
	return func(target M) (M, error) {
		sign, err := DoRsaSign(target, c.signType, c.privateKeyType, c.privateKey)
		if err != nil {
			return target, err
		}
		target.Set("sign", sign)
		return target, nil
	}
}
