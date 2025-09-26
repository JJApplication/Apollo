package utils

// 二进制格式messagePack支持

import (
	"bytes"
	"github.com/vmihailenco/msgpack/v5"
)

func MarshalMsgPack(v interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	enc := msgpack.NewEncoder(&buffer)
	enc.SetCustomStructTag("json")
	err := enc.Encode(v)
	return buffer.Bytes(), err
}

func UnmarshalMsgPack(data []byte, v interface{}) error {
	dec := msgpack.NewDecoder(bytes.NewReader(data))
	dec.SetCustomStructTag("json")
	err := dec.Decode(v)
	return err
}
