package mldsa

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"slices"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

var _ cryptotypes.PubKey = (*PubKey)(nil)

func TestPKSuite(t *testing.T) {
	suite.Run(t, new(PKSuite))
}

type CommonSuite struct {
	suite.Suite
	pk *PubKey
	sk cryptotypes.PrivKey
}

func (suite *CommonSuite) SetupSuite() {
	sk, err := GenPrivKey()
	suite.Require().NoError(err)
	suite.sk = sk
	suite.pk = sk.PubKey().(*PubKey)

}

type PKSuite struct{ CommonSuite }

func (suite *PKSuite) TestString() {
	require := suite.Require()

	pk2 := &PubKey{Key: slices.Clone(suite.pk.Key)}
	require.True(suite.pk.Equals(pk2))
	require.Equal(suite.pk.Bytes(), pk2.Bytes())

	pk3 := &PubKey{Key: []byte("different-key")}
	require.False(suite.pk.Equals(pk3))
}

func (suite *PKSuite) TestType() {
	suite.Require().Equal("mldsa", suite.pk.Type())
}

func (suite *PKSuite) TestBytes() {

	require := suite.Require()

	pk2 := &PubKey{Key: slices.Clone(suite.pk.Key)}
	require.True(suite.pk.Equals(pk2))
	require.Equal(suite.pk.Bytes(), pk2.Bytes())

	pk3 := &PubKey{Key: []byte("different-key")}
	require.False(suite.pk.Equals(pk3))
}

func (suite *PKSuite) TestEquals() {
	require := suite.Require()

	skOther, err := GenPrivKey()
	require.NoError(err)
	pkOther := skOther.PubKey()
	pkOther2 := skOther.PubKey()

	require.False(suite.pk.Equals(pkOther))
	require.True(pkOther.Equals(pkOther2))
	require.True(pkOther2.Equals(pkOther))
	require.True(pkOther.Equals(pkOther), "Equals must be reflexive")
}
