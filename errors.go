/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 19:23
*/

package alipay

import "fmt"

func FormatErrors(errs ...error) error {
	return fmt.Errorf("err: %s", errs)
}
