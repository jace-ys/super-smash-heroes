package response

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

type PbJSON *json.RawMessage

type protoMessage interface {
	proto.Message
}

func EncodePbToJSON(v protoMessage) *json.RawMessage {
	var buf bytes.Buffer
	m := &jsonpb.Marshaler{}
	m.Marshal(&buf, v)
	raw := json.RawMessage(buf.Bytes())
	return &raw
}
