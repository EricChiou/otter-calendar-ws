package sha3

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

type sha3Long int
type long struct {
	L224 sha3Long
	L256 sha3Long
	L384 sha3Long
	L512 sha3Long
}

var Long = long{
	L224: 224,
	L256: 256,
	L384: 384,
	L512: 512,
}

func Encrypt(s string, l sha3Long) string {
	switch l {
	case 224:
		byte28 := sha3.Sum224([]byte(s))
		return hex.EncodeToString(byte28[:])

	case 256:
		byte32 := sha3.Sum256([]byte(s))
		return hex.EncodeToString(byte32[:])

	case 384:
		byte48 := sha3.Sum384([]byte(s))
		return hex.EncodeToString(byte48[:])

	case 512:
		byte64 := sha3.Sum512([]byte(s))
		return hex.EncodeToString(byte64[:])

	default:
		byte48 := sha3.Sum256([]byte(s))
		return hex.EncodeToString(byte48[:])
	}
}
