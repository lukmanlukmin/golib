package encryption

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/xdg-go/pbkdf2"
)

func EncryptPassword(password string, salt string) []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), 10000, 512, sha512.New)
}

const hextable = "0123456789abcdef"

func Encode(dst, src []byte) int {
	j := 0
	for _, v := range src {
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return len(src) * 2
}
func EncodedLen(n int) int { return n * 2 }
func EncodeToString(src []byte) string {
	dst := make([]byte, EncodedLen(len(src)))
	Encode(dst, src)
	return string(dst)
}

func IsSamePassword(passwordHash []byte, passwordPlain string) bool {
	return (hex.EncodeToString(passwordHash) == passwordPlain)
}
