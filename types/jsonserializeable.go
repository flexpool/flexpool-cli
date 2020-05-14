package types

import (
	"fmt"
	"math/big"
)

// JSONBigInt is a JSON Serializable big.Int wrapper
type JSONBigInt struct {
	big.Int
}

// MarshalJSON marshals big.Int
func (b JSONBigInt) MarshalJSON() ([]byte, error) {
	return []byte(b.String()), nil
}

// UnmarshalJSON unmarshals big.Int
func (b *JSONBigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	var z big.Int
	_, ok := z.SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Int = z
	return nil
}
