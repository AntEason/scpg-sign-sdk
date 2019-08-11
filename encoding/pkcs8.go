package encoding

import (
	"crypto/x509"
	"encoding/base64"
	"errors"
)

func ParsePKCS8PrivateKey(pkey string) (key interface{}, err error) {
	//var block *pem.Block
	//block, _ = pem.Decode(data)
	block,_:=base64.StdEncoding.DecodeString(pkey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	key, err =x509.ParsePKCS8PrivateKey(block)
	if err != nil {
		return nil, err
	}
	return key, err
}