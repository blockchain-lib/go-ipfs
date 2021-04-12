package proto

import (
	"bytes"
	"encoding/gob"
	"github.com/ipfs/go-cid"
)

const (
	V1            = "v1"
	MsgAddFile    = "FetchFile"
	MsgQuerySpace = "QuerySpace"
)

type (
	Message struct {
		Type string
		Data interface{}
	}
	FetchFile struct {
		Cid cid.Cid
	}

	QuerySpace struct {
		Space int64
	}
)

func init() {
	gob.Register(Message{})
	gob.Register(FetchFile{})
	gob.Register(QuerySpace{})
}

func (m Message) EncodeMessage() ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DecodeMessage(data []byte) (Message, error) {
	var buffer bytes.Buffer
	buffer.Write(data)
	dec := gob.NewDecoder(&buffer)
	var v Message
	err := dec.Decode(&v)
	if err != nil {
		return Message{}, err
	}
	return v, nil
}
