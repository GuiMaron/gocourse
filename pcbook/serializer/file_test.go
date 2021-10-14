package serializer_test

import (
	"github.com/GuiMaron/gocourse/pcbook/pb"
	"github.com/GuiMaron/gocourse/pcbook/sample"
	"github.com/GuiMaron/gocourse/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestFileSerializer (t *testing.T) {

	t.Parallel()

	binaryFile	:= "../tmp/laptop.bin"

	laptop1	:= sample.NewLaptop()
	err 	:= serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

	laptop2 := &pcbook.Laptop{}

	serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)

	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

}
