package helpers

import (
	"github.com/ethereum/go-ethereum/common"
)

func IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}
