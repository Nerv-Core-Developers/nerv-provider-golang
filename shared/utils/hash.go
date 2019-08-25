package utils

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/blake2b"
)

func Hash512(data []byte) [64]byte {
	hash := blake2b.Sum512(data)
	return hash
}

func Hash256(data []byte) [32]byte {
	hash := blake2b.Sum256(data)
	return hash
}

func RandomUUIDGenBytes() []byte {
	id := uuid.NewV4().Bytes()
	return id
}

func RandomUUIDGenString() string {
	id := uuid.NewV4().String()
	return id
}
