package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func ProtobufToJson(message proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	return marshaler.MarshalToString(message)
}
