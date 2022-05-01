package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func ToByte(anything interface{}) []byte {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	if err := enc.Encode(anything); err != nil {
		log.Panic(err)
	}

	return buf.Bytes()
}
