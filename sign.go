/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/08 17:13
*/

package alipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
)

// DoRsaSign
// Default privKeyType PKCS8
// Default signType SHA256
func DoRsaSign(m map[string]interface{}, signType string, privateKeyType string, privateKey string) (sign string, err error) {
	var (
		block     *pem.Block
		h         hash.Hash
		key       *rsa.PrivateKey
		hashType  crypto.Hash
		encrypted []byte
	)
	pk := FormatPrivateKey(privateKey)

	if block, _ = pem.Decode([]byte(pk)); block == nil {
		err = errors.New("pem.Decode：privateKey decode error")
		return
	}

	switch privateKeyType {
	case PKCS1:
		if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			return
		}
	case PKCS8:
		pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return sign, err
		}
		pk8, ok := pkcs8Key.(*rsa.PrivateKey)
		if !ok {
			err = errors.New("parse PKCS8 key error")
			return sign, err
		}
		key = pk8
	default:
		if key, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
			return
		}
	}

	switch signType {
	case RSA:
		h = sha1.New()
		hashType = crypto.SHA1
	case RSA2:
		h = sha256.New()
		hashType = crypto.SHA256
	default:
		h = sha256.New()
		hashType = crypto.SHA256
	}
	if _, err = h.Write([]byte(EncodeMapParams(m))); err != nil {
		return
	}
	if encrypted, err = rsa.SignPKCS1v15(rand.Reader, key, hashType, h.Sum(nil)); err != nil {
		return
	}
	sign = base64.StdEncoding.EncodeToString(encrypted)
	return
}

type LoadPublicKeyFunc func() (publicKey *rsa.PublicKey, err error)

func loadPublicKeyFormBytes(input []byte) (publicKey *rsa.PublicKey, err error) {
	var (
		block  *pem.Block
		pubKey interface{}
		ok     bool
	)
	if block, _ = pem.Decode(input); block == nil {
		err = errors.New("支付宝公钥Decode错误")
		return
	}
	if pubKey, err = x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		err = fmt.Errorf("x509.ParsePKIXPublicKey：%w", err)
		return
	}
	if _, ok = pubKey.(*rsa.PublicKey); !ok {
		err = errors.New("public key 类型断言错误")
		return
	}
	return pubKey.(*rsa.PublicKey), nil
}

func loadPublicCertFormBytes(input []byte) (publicKey *rsa.PublicKey, err error) {
	var (
		block  *pem.Block
		pubKey *x509.Certificate
		ok     bool
	)
	if block, _ = pem.Decode(input); block == nil {
		err = errors.New("支付宝公钥Decode错误")
		return
	}
	if pubKey, err = x509.ParseCertificate(block.Bytes); err != nil {
		err = fmt.Errorf("x509.ParsePKIXPublicKey：%w", err)
		return
	}
	if _, ok = pubKey.PublicKey.(*rsa.PublicKey); !ok {
		err = errors.New("public key 类型断言错误")
		return
	}
	return pubKey.PublicKey.(*rsa.PublicKey), nil
}

func validateSign(data, sign, signType string, publicKey *rsa.PublicKey) (err error) {
	var (
		h        hash.Hash
		hashType crypto.Hash
	)
	signBytes, _ := base64.StdEncoding.DecodeString(sign)
	switch signType {
	case RSA:
		hashType = crypto.SHA1
	case RSA2:
		hashType = crypto.SHA256
	default:
		hashType = crypto.SHA256
	}
	h = hashType.New()
	h.Write([]byte(data))
	return rsa.VerifyPKCS1v15(publicKey, hashType, h.Sum(nil), signBytes)
}

func PublicKeyFormString(input string) LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		return loadPublicKeyFormBytes([]byte(FormatPublicKey(input)))
	}
}

func PublicCertFormBytes(input []byte) LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		return loadPublicCertFormBytes(input)
	}
}

func PublicCertFormPath(input string) LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		bytes, err := ioutil.ReadFile(input)
		if err != nil {
			err = fmt.Errorf("支付宝公钥文件读取失败: %w", err)
			return nil, err
		}
		return loadPublicCertFormBytes(bytes)
	}
}

func ValidateSignSync(data, sign string, loadPublicKey LoadPublicKeyFunc) (ok bool, err error) {
	publicKey, err := loadPublicKey()
	if err != nil {
		return false, err
	}
	if err := validateSign(data, sign, RSA2, publicKey); err != nil {
		return false, err
	}
	return true, nil
}

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
	case M:
		m := payload.(M)
		if m != nil {
			if v, ok := m["sign"]; ok {
				sign = v.(string)
				delete(m, "sign")
			}
			if v, ok := m["sign_type"]; ok {
				signType = v.(string)
				delete(m, "sign_type")
			}
			data = EncodeMapParams(m)
		}
	}
	if err := validateSign(data, sign, signType, publicKey); err != nil {
		return false, err
	}
	return true, nil
}
