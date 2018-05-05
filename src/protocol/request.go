package protocol

import (
	"encoding/json"
)

type Request struct {
	Code     int32
	Sender   string
	Receiver string
	Data     string
}

func EncodeRequest(r *Request) ([]byte, error) {
	return json.Marshal(r)
}

func DecodeRequest(bytes []byte, r *Request) error {
	return json.Unmarshal(bytes, &r)
}
