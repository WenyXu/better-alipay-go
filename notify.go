/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/09 0:17
*/

package alipay

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

func ParseNotifyToStruct(reqOrValues interface{}, result interface{}) (err error) {
	m, err := ParseNotifyToMap(reqOrValues)
	if err != nil {
		return err
	}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, result)
}

func ParseNotifyToMap(reqOrValues interface{}) (m M, err error) {
	switch reqOrValues.(type) {
	case *http.Request:
		if err = reqOrValues.(*http.Request).ParseForm(); err != nil {
			return nil, err
		}
		return ParseNotifyByURLValues(reqOrValues.(*http.Request).Form)
	case url.Values:
		return ParseNotifyByURLValues(reqOrValues.(url.Values))
	default:
		return nil, errors.New("please input req or url values")
	}
}

func ParseNotifyByURLValues(value url.Values) (m M, err error) {
	m = make(M, len(value)+1)
	for k, v := range value {
		if len(v) == 1 {
			m.Set(k, v[0])
		}
	}
	return
}
