/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/07 5:55
*/

/*
	Utility for key and cert file
*/
package cert

import (
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/WenyXu/better-alipay-go/global"
	"github.com/WenyXu/better-alipay-go/sign"
)

// LoadCertSN load root cert sn form path or bytes
func LoadCertSN(certPathOrData interface{}) (sn string, err error) {
	var certData []byte
	switch certPathOrData.(type) {
	case string:
		certData, err = ioutil.ReadFile(certPathOrData.(string))
	case []byte:
		certData = certPathOrData.([]byte)
	}
	if err != nil {
		return sn, err
	}

	if block, _ := pem.Decode(certData); block != nil {
		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return sn, err
		}
		name := cert.Issuer.String()
		serialNumber := cert.SerialNumber.String()
		h := md5.New()
		h.Write([]byte(name))
		h.Write([]byte(serialNumber))
		sn = hex.EncodeToString(h.Sum(nil))
	}
	if sn == "" {
		return "", errors.New("failed to load cert sn,check the cert path or data")
	}
	return sn, nil
}

// LoadRootCertSN load root cert sn form path or bytes
func LoadRootCertSN(rootCertPathOrData interface{}) (sn string, err error) {
	var certData []byte
	var certEnd = `-----END CERTIFICATE-----`
	switch rootCertPathOrData.(type) {
	case string:
		certData, err = ioutil.ReadFile(rootCertPathOrData.(string))
	case []byte:
		certData = rootCertPathOrData.([]byte)
	}

	pems := strings.Split(string(certData), certEnd)
	for _, c := range pems {
		if block, _ := pem.Decode([]byte(c + certEnd)); block != nil {
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				continue
			}
			if !allowedSignatureAlgorithm[cert.SignatureAlgorithm.String()] {
				continue
			}
			name := cert.Issuer.String()
			serialNumber := cert.SerialNumber.String()
			h := md5.New()
			h.Write([]byte(name))
			h.Write([]byte(serialNumber))
			if sn == "" {
				sn += hex.EncodeToString(h.Sum(nil))
			} else {
				sn += "_" + hex.EncodeToString(h.Sum(nil))
			}
		}
	}
	if sn == "" {
		return sn, errors.New("failed to get sn,please check your cert")
	}
	return sn, nil
}

// FormatPrivateKey  convert private key string to cert file text pattern
func FormatPrivateKey(privateKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

// FormatPublicKey convert public key string to cert file text pattern
func FormatPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

var allowedSignatureAlgorithm = map[string]bool{
	"MD2-RSA":       true,
	"MD5-RSA":       true,
	"SHA1-RSA":      true,
	"SHA256-RSA":    true,
	"SHA384-RSA":    true,
	"SHA512-RSA":    true,
	"SHA256-RSAPSS": true,
	"SHA384-RSAPSS": true,
	"SHA512-RSAPSS": true,
}

// LoadPrivateKeyFormString load private key form string
//
// by default key type is PKCS8, sign type is SHA256
func LoadPrivateKeyFormString(privateKeyType string, privateKey string) sign.LoadPrivateKeyFunc {
	return LoadPrivateKeyFormBytes(privateKeyType, []byte(FormatPrivateKey(privateKey)))
}

// LoadPrivateKeyFormBytes load private key form byte
func LoadPrivateKeyFormBytes(privateKeyType string, input []byte) sign.LoadPrivateKeyFunc {
	return func() (publicKey *rsa.PrivateKey, err error) {
		var (
			block *pem.Block
		)
		if block, _ = pem.Decode(input); block == nil {
			err = errors.New("pem.Decode：privateKey decode error")
			return
		}
		switch privateKeyType {
		case global.PKCS1:
			if publicKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
				return
			}
		case global.PKCS8:
			pkcs8Key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				return nil, err
			}
			pk8, ok := pkcs8Key.(*rsa.PrivateKey)
			if !ok {
				err = errors.New("parse PKCS8 key error")
				return nil, err
			}
			publicKey = pk8
		default:
			if publicKey, err = x509.ParsePKCS1PrivateKey(block.Bytes); err != nil {
				return
			}
		}
		return
	}
}

// LoadPublicKeyFormBytes load public key form byte
func LoadPublicKeyFormBytes(input []byte) (publicKey *rsa.PublicKey, err error) {
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

// LoadPublicCertFormBytes load public cert form byte
func LoadPublicCertFormBytes(input []byte) (publicKey *rsa.PublicKey, err error) {
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

// PublicKeyFormString load public key form string
func PublicKeyFormString(input string) sign.LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		return LoadPublicKeyFormBytes([]byte(FormatPublicKey(input)))
	}
}

// PublicCertFormBytes load public key form byte
func PublicCertFormBytes(input []byte) sign.LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		return LoadPublicCertFormBytes(input)
	}
}

// PublicCertFormPath load public key form path
func PublicCertFormPath(input string) sign.LoadPublicKeyFunc {
	return func() (publicKey *rsa.PublicKey, err error) {
		bytes, err := ioutil.ReadFile(input)
		if err != nil {
			err = fmt.Errorf("支付宝公钥文件读取失败: %w", err)
			return nil, err
		}
		return LoadPublicCertFormBytes(bytes)
	}
}
