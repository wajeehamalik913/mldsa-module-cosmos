package ante

import (
	"fmt"

	ed25519 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	secp256r1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/wajeehamalik913/mldsa-module/crypto/keys/mldsa"
)

// MldsaSigVerificationGasConsumer is used to consume gas for mldsa.PubKey signatures.
func MldsaSigVerificationGasConsumer(
	meter storetypes.GasMeter,
	sig signingtypes.SignatureV2,
	params types.Params,
) error {
	switch pk := sig.PubKey.(type) {
	case *mldsa.PubKey:
		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: mldsa")
		return nil

	case *secp256k1.PubKey:
		meter.ConsumeGas(params.SigVerifyCostSecp256k1, "ante verify: secp256k1")
		return nil

	case *ed25519.PubKey:
		meter.ConsumeGas(params.SigVerifyCostED25519, "ante verify: ed25519")
		return nil

	case *secp256r1.PubKey:
		meter.ConsumeGas(params.SigVerifyCostSecp256r1(), "ante verify: secp256r1")
		return nil

	case multisig.PubKey:
		multisignature, ok := sig.Data.(*signingtypes.MultiSignatureData)
		if !ok {
			return fmt.Errorf("expected *MultiSignatureData, got %T", sig.Data)
		}
		return ante.ConsumeMultisignatureVerificationGas(meter, multisignature, pk, params, sig.Sequence)

	default:
		return errorsmod.Wrapf(sdkerrors.ErrInvalidPubKey, "unrecognized public key type: %T", pk)
	}
}
