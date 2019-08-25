package utils

import "encoding/base32"

func EncodeBase32(data []byte) string {
	return base32.StdEncoding.EncodeToString(data)
}

func DecodeBase32(data string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(data)
}
