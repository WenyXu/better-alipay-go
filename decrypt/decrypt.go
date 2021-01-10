/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/11 1:42
*/

package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)

// Decrypt decrypt data to struct
// https://opendocs.alipay.com/mini/introduce/aes
// https://opendocs.alipay.com/open/common/104567
func Decrypt(encrypted, secretKey string, result interface{}) (err error) {
	var (
		key, originData []byte
		iv              = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		block           cipher.Block
		blockMode       cipher.BlockMode
	)
	key, err = base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		return
	}
	secretData, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return err
	}
	if block, err = aes.NewCipher(key); err != nil {
		return fmt.Errorf("aes.NewCipher：%w", err)
	}
	if len(secretData)%len(key) != 0 {
		return errors.New("incorrect encrypted data")
	}
	blockMode = cipher.NewCBCDecrypter(block, iv)
	originData = make([]byte, len(secretData))
	blockMode.CryptBlocks(originData, secretData)
	if len(originData) > 0 {
		originData = PKCS5UnPadding(originData)
	}
	if err = json.Unmarshal(originData, &result); err != nil {
		return fmt.Errorf("json.Unmarshal(%s)：%w", string(originData), err)
	}
	return
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])   //找到Byte数组最后的填充byte
	return origData[:(length - unpadding)] //只截取返回有效数字内的byte数组
}
