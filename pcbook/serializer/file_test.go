package serializer_test

import (
	"github.com/GuiMaron/gocourse/pcbook/pb/pcbook"
	"github.com/GuiMaron/gocourse/pcbook/sample"
	"github.com/GuiMaron/gocourse/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestFileSerializer (t *testing.T) {

	t.Parallel()

	binaryFile	:= "../tmp/laptop.bin"
	jsonFile	:= "../tmp/laptop.json"

	laptop1	:= sample.NewLaptop()
	laptop2 := &pcbook.Laptop{}

	err 	:= serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)

	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)

}
