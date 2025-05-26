package mldsa

import (
	"crypto/sha256"

	"github.com/cloudflare/circl/sign/mldsa/mldsa44"
	cmtcrypto "github.com/cometbft/cometbft/crypto"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"golang.org/x/crypto/ripemd160"
)

var _ cryptotypes.PubKey = &PubKey{}

func (pubKey *PubKey) Address() cmtcrypto.Address {
	if len(pubKey.Key) != mldsa44.Scheme().PublicKeySize() {
		panic("length of pubkey is incorrect")
	}
	sha := sha256.Sum256(pubKey.Key)
	hasherRIPEMD160 := ripemd160.New()
	hasherRIPEMD160.Write(sha[:]) // does not error
	return cmtcrypto.Address(hasherRIPEMD160.Sum(nil))
}

func (pk *PubKey) Bytes() []byte {
	return pk.Key
}

func (pk *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	pub, err := mldsa44.Scheme().UnmarshalBinaryPublicKey(pk.Key)
	if err != nil {
		return false
	}
	return pub.Scheme().Verify(pub, msg, sig, nil)
}

func (pk *PubKey) Equals(other cryptotypes.PubKey) bool {
	Pk2, ok := other.(*PubKey)
	if !ok {
		return false
	}

	scheme := mldsa44.Scheme()
	pKey1, err1 := scheme.UnmarshalBinaryPublicKey(pk.Key)
	pKey2, err2 := scheme.UnmarshalBinaryPublicKey(Pk2.Key)

	if err1 != nil || err2 != nil {
		return false
	}

	return pKey1.Equal(pKey2)
}

func (pk *PubKey) Type() string {
	return "mldsa/PubKey"
}
