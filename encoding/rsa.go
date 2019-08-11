package encoding

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
)

func Sign(data string ,priv *rsa.PrivateKey) (string, error) {

	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, priv,
		crypto.SHA1, hashed)
	if signature == nil {
		return "", err
	}
	sign := base64.StdEncoding.EncodeToString(signature)  //转换成base64返回
	return string(sign), nil
}