package coin

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func (b *Block) Serialize() []byte{
	var result bytes.Buffer
	encoder :=gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err == nil {
		fmt.Println("encode fail")
	}
	return result.Bytes()
}