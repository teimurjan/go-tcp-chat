package protocol

import (
	"json"
)

type Address struct {
	ip   string
	port int32
	name string
}

type Protocol struct {
	code     int32
	receiver Address
	sender   Address
	dataSize int32
	data     []byte
}

// Encode implements marshling for Protocol struct
// returns sequence of bytes.
func Encode(p *Protocol) ([]byte, error) {
	return json.Marshal(p)
}

// Decode implements unmarshaling transforming
// sequence of bytes to Protocol structure.
func Decode(bytes []byte, p *Protocol) error {
	return json.Unmarshal(bytes, &p)
}
