/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/07 5:55
*/

package alipay

import (
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"strings"
)

func loadCertSN(certPathOrData interface{}) (sn string, err error) {
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

func loadRootCertSN(rootCertPathOrData interface{}) (sn string, err error) {
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

// FormatPrivateKey
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

// FormatPublicKey
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
