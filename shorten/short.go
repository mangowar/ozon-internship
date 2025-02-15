package shorten

import (
	"hash"
	"hash/fnv"
	"strings"
)

const (
	alphabet     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"
	alphabet_len = uint64(len(alphabet))
)

// ShortenLink выполняет сжатие URL
func ShortenLink(url string) string {
	var (
		nums     []uint64
		builder  strings.Builder
		new_hash hash.Hash64
		value    uint64
	)
	new_hash = fnv.New64()
	new_hash.Write([]byte(url))
	value = new_hash.Sum64()
	for i := 0; i < 10; i++ {
		nums = append(nums, value%alphabet_len)
		value /= alphabet_len
	}
	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	return builder.String()
}

func TransfornLink(baseURL, short string) string {
	res := baseURL + "/r?url=" + short
	return res
}
