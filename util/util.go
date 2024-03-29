package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func ArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func StringToArray(str string, delim string) []int {
	tmp := strings.Split(str, delim)
	var res []int
	for i := 0; i < len(tmp); i++ {
		num, err := strconv.Atoi(tmp[i])
		if err == nil {
			res = append(res, num)
		}
	}

	return res
}

func Sha256(bz []byte) []byte {
	h := sha256.New()
	h.Write(bz)
	return h.Sum(nil)
}

func Byte2Hex(data []byte) string {
	str := hex.EncodeToString(data)
	return str
}
