package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON (message proto.Message) (string, error) {

	marshaler := jsonpb.Marshaler{
		EnumsAsInts:	false,
		EmitDefaults:	true,
		Indent:			"\t",
		OrigName: 		true,
	}

	return marshaler.MarshalToString(message)
}
