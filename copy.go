package flightORM

import (
	"github.com/pingcap/errors"
	"github.com/vmihailenco/msgpack/v5"
)

type Codec interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
}

var codec Codec

func init() {
	codec = &MsgPack{}
}

type MsgPack struct{}

func (m *MsgPack) Encode(i interface{}) ([]byte, error) {
	return msgpack.Marshal(i)
}

func (m *MsgPack) Decode(bytes []byte, i interface{}) error {
	return msgpack.Unmarshal(bytes, i)
}

func Copy(dest, src interface{}) error {
	bytes, err := codec.Encode(src)
	if err != nil {
		return errors.Trace(err)
	}
	return codec.Decode(bytes, dest)
}
