package mldsa

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/suite"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

var _ cryptotypes.PrivKey = &PrivKey{}

func TestSKSuite(t *testing.T) {
	suite.Run(t, new(SKSuite))
}

type CommonSKSuite struct {
	suite.Suite
	sk cryptotypes.PrivKey
	pk cryptotypes.PubKey
}

func (suite *SKSuite) SetupTest() {
	sk, err := GenPrivKey()
	suite.Require().NoError(err)

	suite.sk = sk
	suite.pk = sk.PubKey().(*PubKey)
}

type SKSuite struct{ CommonSKSuite }

func (suite *SKSuite) TestPubKey() {
	pk := suite.sk.PubKey()
	suite.True(suite.sk.PubKey().Equals(pk))
}

func (suite *SKSuite) TestBytes() {
	require := suite.Require()

	originalSk, ok := suite.sk.(*PrivKey)
	require.True(ok)

	sk2 := &PrivKey{Key: slices.Clone(originalSk.Key)}
	require.True(suite.sk.Equals(sk2))
	require.Equal(suite.sk.Bytes(), sk2.Bytes())

	sk3 := &PrivKey{Key: []byte("other-privkey")}
	require.False(suite.sk.Equals(sk3))
}

func (suite *SKSuite) TestSign() {

	require := suite.Require()

	msg := []byte("sample message")
	sig, err := suite.sk.Sign(msg)
	require.NoError(err)

	valid := suite.pk.VerifySignature(msg, sig)
	require.True(valid)
}

func BenchmarkSigVerifyMLDSA(b *testing.B) {
	priv, _ := GenPrivKey() // assuming you have this
	msg := []byte("test message")
	sig, _ := priv.Sign(msg)
	pub := priv.PubKey().(*PubKey)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !pub.VerifySignature(msg, sig) {
			b.FailNow()
		}
	}
}
