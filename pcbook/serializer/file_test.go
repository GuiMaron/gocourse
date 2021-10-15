package serializer_test

import (
	"fmt"
	pcbook "github.com/GuiMaron/gocourse/pcbook/pb/proto"
	"github.com/GuiMaron/gocourse/pcbook/sample"
	"github.com/GuiMaron/gocourse/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"reflect"
	"strings"
	"testing"
)

func TestFileSerializer (t *testing.T) {

	binaryFile	:= "../tmp/laptop.bin"

	laptop1	:= sample.NewLaptop()
	laptop2 := &pcbook.Laptop{}

	err 	:= serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)

	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)

}
