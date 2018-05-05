package main

import (
	"encoding/json"
	"fmt"
)

/*type Address struct {
	ip   string
	port int32
	name string
}*/

type Protocol struct {
	code int32
	//receiver Address
	//sender   Address
	dataSize int32
	//data     []byte
}

// Encode implements marshling for Protocol struct
// returns sequence of bytes.
func Encode(p *Protocol) ([]byte, error) {
	result, err := json.Marshal(p)
	fmt.Println(result)
	return result, err
}

// Decode implements unmarshaling transforming
// sequence of bytes to Protocol structure.
func Decode(bytes []byte, p *Protocol) error {
	return json.Unmarshal(bytes, &p)
}

func main() {
	//a := Address{"192.17.0.5", 9999, "Alice"}
	p := &Protocol{0, 32} //, []byte("Hello, Alice. I will kill you")}
	var pr *Protocol
	encoded, e := Encode(p)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Printf("%T, %v", encoded, string(encoded))
	Decode(encoded, pr)
	fmt.Printf("%v", pr)
}
