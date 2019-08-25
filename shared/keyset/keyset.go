package keyset

import (
	"bytes"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/ed25519"
)

type KeySet struct {
	PrivateKey []byte
	PublicKey  []byte
}

func (key *KeySet) GenerateKey(seed []byte) (*KeySet, error) {
	var err error
	key.PublicKey, key.PrivateKey, err = ed25519.GenerateKey(bytes.NewBuffer(seed))
	if err != nil {
		return key, err
	}
	return key, nil
}

func (key *KeySet) Sign(dataHash []byte) ([]byte, error) {
	result := ed25519.Sign(key.PrivateKey, dataHash)
	return result, nil
}

func (key *KeySet) ImportByString(privateKey string) (*KeySet, error) {
	newKey := ed25519.PrivateKey{}
	newKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return key, err
	}
	key.PublicKey = newKey.Public().(ed25519.PublicKey)
	key.PrivateKey = newKey
	return key, nil
}

func (key *KeySet) ImportByBytes(privateKey []byte) (*KeySet, error) {
	newKey := ed25519.PrivateKey{}
	newKey = privateKey
	if len(newKey) < ed25519.PrivateKeySize {
		return key, errors.New("invalid key length")
	}
	key.PublicKey = newKey.Public().(ed25519.PublicKey)
	key.PrivateKey = newKey
	return key, nil
}

func VerifySig(data, signature, pbkey []byte) (bool, error) {
	isValid := ed25519.Verify(pbkey, data, signature)
	return isValid, nil
}
