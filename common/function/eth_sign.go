package function

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySig(from string, sigHex string, msg string) error {
	fromAddr := common.HexToAddress(from)

	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return err
	}
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return errors.New("sign format error")
	}
	sig[64] -= 27
	pubKey, err := crypto.SigToPub(signHash([]byte(msg)), sig)
	if err != nil {
		return err
	}

	if fromAddr != crypto.PubkeyToAddress(*pubKey) {
		return errors.New("address does not match")
	}
	return nil
}

func VerifySigEthHash(from string, sigHex string, msg string) error {
	fromAddr := common.HexToAddress(from)

	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return err
	}
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return errors.New("sign format error")
	}
	sig[64] -= 27
	pubKey, err := crypto.SigToPub(signHash(ethcommon.HexToHash(msg).Bytes()), sig)
	if err != nil {
		return err
	}

	address := crypto.PubkeyToAddress(*pubKey)
	fmt.Println(address)
	if fromAddr != address {
		return errors.New("address does not match")
	}
	return nil
}

// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L404
// signHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
