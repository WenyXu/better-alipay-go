/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 17:30
*/

/*
	Utility for map[string]interface{}
*/
package m

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"

	errorsMsg "github.com/WenyXu/better-alipay-go/errors"
)

type M map[string]interface{}

type MapOptions struct {
	IgnoreEmptyString *bool
}

// NewMap return a new map, pass into a higher order func to setup map
func NewMap(f func(m M)) M {
	m := make(M)
	f(m)
	return m
}

func (o *MapOptions) SetIgnoreEmptyString(input bool) {
	o.IgnoreEmptyString = &input
}

// mergeOptions merge Options
func mergeOptions(opts ...MapOptions) MapOptions {
	merged := MapOptions{}
	for _, o := range opts {
		if o.IgnoreEmptyString != nil {
			merged.IgnoreEmptyString = o.IgnoreEmptyString
		}
	}
	return merged
}

// Set set map value by key
//
// if opts(MapOptions) SetIgnoreEmptyString(true),will ignore empty string("")
func (m M) Set(key string, value interface{}, opts ...MapOptions) M {
	mergedOption := mergeOptions(opts...)
	switch value.(type) {
	case func(M):
		_m := make(M)
		value.(func(M))(_m)
		m[key] = _m
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

// GetMapValueByKey get map value by key, marshal non-string type into string
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

// EncodeMapParams sort keys by asc, encode into url string
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

type MakeMapEndpoint func(M) error

// CombineMakeMapEndpointFunc combine MakeMapEndpoint func
func CombineMakeMapEndpointFunc(endpoint ...MakeMapEndpoint) MakeMapEndpoint {
	return func(target M) error {
		var errs []error
		for _, e := range endpoint {
			err := e(target)
			if err != nil {
				errs = append(errs, err)
			}
		}
		if len(errs) != 0 {
			return errorsMsg.FormatErrors(errs...)
		}
		return nil
	}

}

// FormatURLParam convert map into a url.Values
func FormatURLParam(m map[string]interface{}) (urlParam string) {
	v := url.Values{}
	for key, value := range m {
		v.Add(key, value.(string))
	}
	return v.Encode()
}

func MergeMap(source map[string]interface{}) MakeMapEndpoint {
	return func(target M) error {
		if source == nil {
			// allow nil input
			return nil
		}
		for k, v := range source {
			target[k] = v
		}
		return nil
	}
}
