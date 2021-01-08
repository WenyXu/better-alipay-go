/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 17:27
*/

package alipay

import (
	"net/url"
)

func FormatURLParam(m map[string]interface{}) (urlParam string) {
	v := url.Values{}
	for key, value := range m {
		v.Add(key, value.(string))
	}
	return v.Encode()
}
