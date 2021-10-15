package serializer

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"reflect"
	"strings"
)

func PrintMethods (x interface{}) {

	fmt.Println("Reflecting")

	v := reflect.ValueOf(x)
	r := v.Type()
	y := reflect.TypeOf(x)

	fmt.Printf("type %s with %d methods - typeof %v\n",r, v.NumMethod(), y)

	for i := 0; i < v.NumMethod(); i ++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s \n", r, r.Method(i).Name, strings.TrimPrefix(methType.String(), "func"))
	}

}

func ProtobufToJSON (message proto.Message) (string, error) {

	PrintMethods(message)

	marshaler := jsonpb.Marshaler{
		EnumsAsInts:	false,
		EmitDefaults:	true,
		Indent:			"\t",
		OrigName: 		true,
	}

	return marshaler.MarshalToString(message)
}
