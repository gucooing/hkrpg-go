package random

import (
	"encoding/hex"
	"math/rand"
)

func GetRandomByte(len int) []byte {
	ret := make([]byte, 0)
	for i := 0; i < len; i++ {
		r := uint8(rand.Intn(256))
		ret = append(ret, r)
	}
	return ret
}

func GetRandomByteHexStr(len int) string {
	return hex.EncodeToString(GetRandomByte(len))
}
