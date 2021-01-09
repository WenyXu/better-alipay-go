/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 19:23
*/

package errors

import "fmt"

// FormatErrors format error array
func FormatErrors(errs ...error) error {
	return fmt.Errorf("err: %s", errs)
}
