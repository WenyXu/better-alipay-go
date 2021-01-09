/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/09 0:17
*/

/*
	notify util
*/
package notify

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	_map "github.com/WenyXu/better-alipay-go/m"
)

// ParseNotifyToStruct parse notify form into struct, reqOrValues accept a *http.Request or url.Values value
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

// ParseNotifyToMap parse notify form into map, reqOrValues accept a *http.Request or url.Values value
func ParseNotifyToMap(reqOrValues interface{}) (m _map.M, err error) {
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

// ParseNotifyToMap parse notify form into map, pass into url.Values
func ParseNotifyByURLValues(value url.Values) (m _map.M, err error) {
	m = make(_map.M, len(value)+1)
	for k, v := range value {
		if len(v) == 1 {
			m.Set(k, v[0])
		}
	}
	return
}
