package utils

import "math/big"

// WeiToEther converts Weis to Ether
func WeiToEther(weis *big.Int) float64 {
	return float64(weis.Int64()) / 1000000000000000000
}
