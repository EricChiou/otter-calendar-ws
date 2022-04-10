package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"hash"
	"strings"
)

type algType string
type alg struct {
	HS256 algType
	HS384 algType
	HS512 algType
}

var Alg = alg{
	HS256: "HS256",
	HS384: "HS384",
	HS512: "HS512",
}

func Generate(payload interface{}, key string, alg algType) (string, error) {
	// header
	jwtHeader := base64.StdEncoding.EncodeToString([]byte(`{"typ":"JWT","alg":"` + alg + `"}`))

	// payload
	bytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	jwtPayload := base64.StdEncoding.EncodeToString(bytes)

	// signature
	jwtSignature := encryptSignature(jwtHeader+"."+jwtPayload, key, alg)

	return jwtHeader + "." + jwtPayload + "." + jwtSignature, nil
}

func Verify(payload interface{}, jwt, key string, alg algType) error {
	jwts := strings.Split(jwt, ".")
	if len(jwts) != 3 {
		return errors.New("invalid jwt")
	}

	// signature
	jwtSignature := encryptSignature(jwts[0]+"."+jwts[1], key, alg)
	if jwts[2] != jwtSignature {
		return errors.New("invalid jwt")
	}

	bytes, err := base64.StdEncoding.DecodeString(jwts[1])
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, payload)
}

func encryptSignature(str, key string, alg algType) string {
	var h hash.Hash
	switch alg {
	case "HS256":
		h = hmac.New(sha256.New, []byte(key))
	case "HS384":
		h = hmac.New(sha512.New384, []byte(key))
	case "HS512":
		h = hmac.New(sha512.New, []byte(key))
	default:
		h = hmac.New(sha256.New, []byte(key))
	}
	h.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
