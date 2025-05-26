package mldsa

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &PubKey{})
	registry.RegisterImplementations((*cryptotypes.PrivKey)(nil), &PrivKey{})
}

func RegisterLegacyAmino(cdc *codec.LegacyAmino) {

	// cdc.RegisterConcrete(PubKey{}, "tendermint/PubKeymldsa", nil)
	cdc.RegisterConcrete(&PubKey{}, "/cosmos.crypto.mldsa.PubKey", nil)

	cdc.RegisterConcrete(PrivKey{}, "tendermint/PrivKeymldsa", nil)
}
