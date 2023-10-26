package registry

import (
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	// Error constant strings
	invalidAddrErr  = "invalid address provided for heuristic. expected %s, got %s"
	couldNotCastErr = "could not cast transit data value to %s type"
	noNestedArgsErr = "no nested args found in session params"
	zeroAddressErr  = "provided address cannot be the zero address"

	// Event declaration strings
	OutputProposedEvent   = "OutputProposed(bytes32,uint256,uint256,uint256)"
	WithdrawalProvenEvent = "WithdrawalProven(bytes32,address,address)"
	MessagePassed         = "MessagePassed(uint256,address,address,uint256,uint256,bytes,bytes32)"
)

var (
	MessagePassedSig    = crypto.Keccak256Hash([]byte(MessagePassed))
	OutputProposedSig   = crypto.Keccak256Hash([]byte(OutputProposedEvent))
	WithdrawalProvenSig = crypto.Keccak256Hash([]byte(WithdrawalProvenEvent))
)
