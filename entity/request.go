/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/07 5:19
*/

package entity

import "github.com/pkg/errors"

type Common struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

func (c Common) Success() bool {
	return c.Code == "10000"
}

func (c Common) ErrorWrap(err error) error {
	if err != nil || !c.Success() {
		if err != nil {
			return errors.Wrap(err, c.Msg)
		}
		return errors.New(c.Msg)
	}
	return nil
}
