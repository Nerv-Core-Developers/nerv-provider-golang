package noid

// NOID - Nerv object identifier

import (
	"encoding/base32"
	"errors"
)

//Identifier format:
// [IDRegistry (4 bytes), IDType (1 byte), ObjectVersion, ID]
type Identifier []byte

func (id Identifier) CloneBytes() []byte {
	bytes := make([]byte, len(id))
	copy(bytes, id)
	return bytes
}

func (id Identifier) IsValid(additionalVerifier func(Identifier, IdentifierFormat) error, format IdentifierFormat) error {
	if len(id) < 256 {
		return errors.New("identifier length is invalid")
	}

	if additionalVerifier != nil {
		err := additionalVerifier(id, format)
		if err != nil {
			return err
		}
	}

	return nil
}

func (id Identifier) IDtoStringBase32() string {
	return base32.StdEncoding.EncodeToString(id)
}

func (id *Identifier) IDFromStringBase32(data string) error {
	idBytes, err := base32.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	var idtemp Identifier

	idtemp = make([]byte, len(idBytes))
	copy(idtemp, idBytes)
	id = &idtemp
	return nil
}

type IdentifierFormat struct {
	Types                []string
	ObjectVersionLength  int
	ObjectVersionChecker func(IDtype string, ObjectVersion []byte) error
	MinIDlength          int
	MaxIDlength          int
	IDChecker            func(IDtype string, ObjectVersion []byte, ID []byte) error
}

type IDType struct {
	Byte   byte
	String string
}
