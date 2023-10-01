package utils

import (
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

func EtherToWei(val *uint256.Int) *uint256.Int {
	return new(uint256.Int).Mul(val, uint256.NewInt(params.Ether))
}

func WeiToEther(val *uint256.Int) *uint256.Int {
	if val == nil {
		return nil
	}
	return new(uint256.Int).Div(val, uint256.NewInt(1e18))
}
