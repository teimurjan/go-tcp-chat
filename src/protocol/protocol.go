package protocol

import (
	"encoding/json"
)

type Address struct {
	Ip   string
	Port int32
	Name string
}

type Protocol struct {
	Code     int32
	Receiver Address
	Sender   Address
	DataSize int32
	Data     []byte
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
