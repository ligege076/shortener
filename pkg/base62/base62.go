package base62

import (
	"math"
	"strings"
)

var (
	baseStr    string
	baseStrLen uint64
)

// MustInit 使用 base62 这个包前必须先调用，完成初始化
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string!")
	}
	baseStr = bs
	baseStrLen = uint64(len(bs))
}

// Int2String 十进制数字转为 base62 字符串
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(baseStr[0])
	}

	bl := []byte{}
	for seq > 0 {
		mod := seq % baseStrLen
		div := seq / baseStrLen
		bl = append(bl, baseStr[mod])
		seq = div
	}

	// 最后把得到的数反转一下
	return string(reverse(bl))
}

// String2Int base62 字符串转为十进制数字
func String2Int(s string) (seq uint64) {
	bl := []byte(s)
	bl = reverse(bl)

	for idx, b := range bl {
		base := math.Pow(float64(baseStrLen), float64(idx))
		seq += uint64(strings.Index(baseStr, string(b))) * uint64(base)
	}

	return seq
}

// reverse 反转字节切片
func reverse(bl []byte) []byte {
	for i, j := 0, len(bl)-1; i < j; i, j = i+1, j-1 {
		bl[i], bl[j] = bl[j], bl[i]
	}
	return bl
}
