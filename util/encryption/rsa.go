package encryption

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"

	"github.com/lukmanlukmin/golib/log"

	"github.com/lestrrat-go/jwx/jwk"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ParsePubKeyFromString(pubKey string) *rsa.PublicKey {
	pubKeyByte, err := hex.DecodeString(pubKey)
	if err != nil {
		fatal(err)
	}
	pubAsli, err := x509.ParsePKIXPublicKey(pubKeyByte)
	if err != nil {
		fatal(err)
	}
	return pubAsli.(*rsa.PublicKey)
}

func GenerateRSAKeyString() (string, string) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	fatal(err)
	pub := &priv.PublicKey

	privKey := x509.MarshalPKCS1PrivateKey(priv)
	pubKey, err := x509.MarshalPKIXPublicKey(pub)
	fatal(err)

	return hex.EncodeToString(privKey), hex.EncodeToString(pubKey)
}

func DecodeHexRSAKeyString(privateKey string, pubKeystring string) ([]byte, []byte) {
	privByte, err := hex.DecodeString(privateKey)
	fatal(err)
	pubByte, err := hex.DecodeString(pubKeystring)
	fatal(err)
	return pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privByte,
		}), pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: pubByte,
		})
}

func PubKeyToJWKKey(pub *rsa.PublicKey) jwk.Key {
	key, err := jwk.New(pub)
	if err != nil {
		fatal(err)
	}
	if err := key.Set(jwk.KeyUsageKey, "enc"); err != nil {
		fatal(fmt.Errorf("failed to set key usage: %w", err))
	}
	if err := key.Set(jwk.AlgorithmKey, "RS256"); err != nil {
		fatal(fmt.Errorf("failed to set key algorithm: %w", err))
	}
	return key
}
