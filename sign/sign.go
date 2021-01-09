/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 17:13
*/

/*
	Utility for sign the param and validate sign & data
*/
package sign

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"hash"

	"github.com/WenyXu/better-alipay-go/global"
	m "github.com/WenyXu/better-alipay-go/m"
)

type LoadPublicKeyFunc func() (publicKey *rsa.PublicKey, err error)
type LoadPrivateKeyFunc func() (publicKey *rsa.PrivateKey, err error)

// Sign params
func Sign(param map[string]interface{}, signType string, loadPrivateFunc LoadPrivateKeyFunc) (sign string, err error) {
	var (
		h         hash.Hash
		hashType  crypto.Hash
		encrypted []byte
	)
	key, err := loadPrivateFunc()
	if err != nil {
		return sign, err
	}
	switch signType {
	case global.RSA:
		h = sha1.New()
		hashType = crypto.SHA1
	case global.RSA2:
		h = sha256.New()
		hashType = crypto.SHA256
	default:
		h = sha256.New()
		hashType = crypto.SHA256
	}
	if _, err = h.Write([]byte(m.EncodeMapParams(param))); err != nil {
		return
	}
	if encrypted, err = rsa.SignPKCS1v15(rand.Reader, key, hashType, h.Sum(nil)); err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(encrypted)
	return
}

func validateSign(data, sign, signType string, publicKey *rsa.PublicKey) (err error) {
	var (
		h        hash.Hash
		hashType crypto.Hash
	)
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	switch signType {
	case global.RSA:
		hashType = crypto.SHA1
	case global.RSA2:
		hashType = crypto.SHA256
	default:
		hashType = crypto.SHA256
	}
	h = hashType.New()
	h.Write([]byte(data))
	return rsa.VerifyPKCS1v15(publicKey, hashType, h.Sum(nil), signBytes)
}

// Validate Response Sign
func ValidateSignSync(data, sign string, loadPublicKey LoadPublicKeyFunc) (ok bool, err error) {
	publicKey, err := loadPublicKey()
	if err != nil {
		return false, err
	}
	if err := validateSign(data, sign, global.RSA2, publicKey); err != nil {
		return false, err
	}
	return true, nil
}

// Validate Notify Sign
func ValidateSign(payload interface{}, loadPublicKey LoadPublicKeyFunc) (ok bool, err error) {
	var (
		sign     string
		signType string
		data     string
	)
	publicKey, err := loadPublicKey()
	if err != nil {
		return false, err
	}
	switch payload.(type) {
	case m.M:
		param := payload.(m.M)
		if param != nil {
			if v, ok := param["sign"]; ok {
				sign = v.(string)
				delete(param, "sign")
			}
			if v, ok := param["sign_type"]; ok {
				signType = v.(string)
				delete(param, "sign_type")
			}
			data = m.EncodeMapParams(param)
		}
	}
	if err := validateSign(data, sign, signType, publicKey); err != nil {
		return false, err
	}
	return true, nil
}
