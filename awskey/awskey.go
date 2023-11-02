package awskey

import (
	"encoding/base32"
	"errors"
	"fmt"
)

type KeyType string

const (
	STSBearerToken       KeyType = "ABIA"
	ContextSpecific      KeyType = "ACCA"
	Group                KeyType = "AGPA"
	IAMUser              KeyType = "AIDA"
	EC2Profile           KeyType = "AIPA"
	AccessKey            KeyType = "AKIA"
	ManagedPolicy        KeyType = "ANPA"
	VersionManagedPolicy KeyType = "ANVA"
	PublicKey            KeyType = "APKA"
	Role                 KeyType = "AROA"
	Certificate          KeyType = "ASCA"
	STSKey               KeyType = "ASIA"
)

type Key struct {
	Type      KeyType
	AccountID string
}

func Decode(key string) (*Key, error) {
	var k Key
	if len(key) < 4 {
		return nil, errors.New("Invalid key too short")
	}

	k.Type = KeyType(key[:4])

	key = key[4:]
	b, err := base32.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("Invalid key %w", err)
	}

	var x uint64
	for i := 0; i < 6; i++ {
		x = x << 8
		x |= uint64(b[i])
	}

	x &= 0x7fffffffff80
	x >>= 7

	k.AccountID = fmt.Sprintf("%012d", x)

	return &k, nil

}
