package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
)

var publicKey *rsa.PublicKey
var privateKey *rsa.PrivateKey

func LoadPublicKey(filePath string) (*rsa.PublicKey, error) {
	byteBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(byteBuffer)
	if block == nil {
		return nil, errors.New("load public key error")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey = pub.(*rsa.PublicKey)
	return publicKey, nil
}

func LoadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	byteBuffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(byteBuffer))
	if block == nil {
		return nil, errors.New("load private key error")
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func Encrypt(s string) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, publicKey, []byte(s), nil)
	if err != nil {
		return "", err
	}

	encrypt := base64.StdEncoding.EncodeToString(ciphertext)
	return encrypt, nil
}

func EncryptByKey(s string, pk *rsa.PublicKey) (string, error) {
	ciphertext, err := rsa.EncryptOAEP(sha512.New(), rand.Reader, pk, []byte(s), nil)
	if err != nil {
		return "", err
	}

	encrypt := base64.StdEncoding.EncodeToString(ciphertext)
	return encrypt, nil
}

func Decrypt(s string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	decrypt, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, privateKey, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decrypt), nil
}

func DecryptByKey(s string, pk *rsa.PrivateKey) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	decrypt, err := rsa.DecryptOAEP(sha512.New(), rand.Reader, pk, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(decrypt), nil
}
