package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"
)

func encodeSha256(str string) []byte {
	alg := sha256.New()
	alg.Write([]byte(str))
	return alg.Sum(nil)
}

func decodeBase58(bt []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bt)
	if err != nil {
		panic(err.Error())
	}
	return string(encoded)
}

func GenerateShortURL(originalUrl string) string {
	urlHashBytes := encodeSha256(originalUrl)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := decodeBase58([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}