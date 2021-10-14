package serializer

import (
	"github.com/GuiMaron/gocourse/pcbook/sample"
	"github.com/GuiMaron/gocourse/pcbook/serializer"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFileSerializer (t *testing.T) {

	t.Parallel()

	binaryFile	:= "../tmp/laptop.bin"

	laptop1	:= sample.NewLaptop()
	err 	:= serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)

	require.NoError(t, err)

}
