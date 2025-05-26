package mldsa

import (
	"github.com/cloudflare/circl/sign/mldsa/mldsa44"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

var _ cryptotypes.PrivKey = &PrivKey{}

func GenPrivKey() (cryptotypes.PrivKey, error) {
	_, sk, err := mldsa44.GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	return &PrivKey{Key: sk.Bytes()}, nil
}

func (sk *PrivKey) Sign(msg []byte) ([]byte, error) {
	privKey, err := mldsa44.Scheme().UnmarshalBinaryPrivateKey(sk.Key)
	if err != nil {
		return nil, err
	}
	return mldsa44.Scheme().Sign(privKey, msg, nil), err
}

func (sk *PrivKey) PubKey() cryptotypes.PubKey {
	privKey, err := mldsa44.Scheme().UnmarshalBinaryPrivateKey(sk.Key)
	if err != nil {
		return nil
	}
	pubKey := privKey.Public().(*mldsa44.PublicKey)
	return &PubKey{Key: pubKey.Bytes()}
}

func (sk *PrivKey) Bytes() []byte {
	return sk.Key
}

func (sk *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	Sk2, ok := other.(*PrivKey)
	if !ok {
		return false
	}

	PK1, err := mldsa44.Scheme().UnmarshalBinaryPrivateKey(sk.Key)
	if err != nil {
		return false
	}

	PK2, err := mldsa44.Scheme().UnmarshalBinaryPrivateKey(Sk2.Key)
	if err != nil {
		return false
	}

	return PK1.Equal(PK2)
}

func (sk *PrivKey) Type() string {
	return "mldsa/PrivKey"
}
