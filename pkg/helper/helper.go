package helper

import "github.com/ivanrad/base62"

func GenerateShortCode(url []byte) string {
	return base62.EncodeToString(url)
}
