package mldsa

import (
	"cosmossdk.io/depinject"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

type MldsaInterfaceInjector struct{}

// ✅ This is the new required marker method
func (MldsaInterfaceInjector) IsOnePerModuleType() {}

func (MldsaInterfaceInjector) Inject(reg codectypes.InterfaceRegistry) error {
	RegisterInterfaces(reg)
	return nil
}

// ✅ Implementing the interface explicitly
var _ depinject.OnePerModuleType = MldsaInterfaceInjector{}
