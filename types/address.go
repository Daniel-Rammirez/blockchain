package types

import (
	"encoding/hex"
	"fmt"
)

type Address [20]uint8

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}

	return b
}
func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		msg := fmt.Sprintf("Given bytes with length %d should be 32", len(b))
		panic(msg)
	}

	var value [20]uint8

	for i := range 20 {
		value[i] = b[i]
	}

	return Address(value)
}
