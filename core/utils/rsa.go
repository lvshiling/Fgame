package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
)

func CheckRsaSign(publicKey, signStr, originalData string) (err error) {
	sign, err := base64.StdEncoding.DecodeString(signStr)
	if err != nil {
		return
	}
	public, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return
	}
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return
	}
	hashBytes := sha1.Sum([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hashBytes[:], sign)
}

func RsaSign(originalData, prvKey string) (sign string, err error) {
	keyBytes, err := base64.StdEncoding.DecodeString(prvKey)
	if err != nil {
		return
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {

		return "", err
	}

	hashBytes := sha1.Sum([]byte(originalData))

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hashBytes[:])
	if err != nil {

		return "", err
	}
	sign = base64.StdEncoding.EncodeToString(signature)
	return

}
